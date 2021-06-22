// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packetbrokeragent

import (
	"context"
	"fmt"
	"time"

	pbtypes "github.com/gogo/protobuf/types"
	mappingpb "go.packetbroker.org/api/mapping/v2"
	packetbroker "go.packetbroker.org/api/v3"
	clusterauth "go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"google.golang.org/grpc"
)

const onlineTTL = 10 * time.Minute

type messageEncrypter interface {
	encryptUplink(context.Context, *packetbroker.UplinkMessage) error
}

type gsPbaServer struct {
	netID             types.NetID
	clusterID         string
	config            ForwarderConfig
	messageEncrypter  messageEncrypter
	contextDecoupler  contextDecoupler
	tenantIDExtractor TenantIDExtractor
	upstreamCh        chan *uplinkMessage
	mapperConn        *grpc.ClientConn
}

var errForwarderDisabled = errors.DefineFailedPrecondition("forwarder_disabled", "Forwarder is disabled")

// PublishUplink is called by the Gateway Server when an uplink message arrives and needs to get forwarded to Packet Broker.
func (s *gsPbaServer) PublishUplink(ctx context.Context, up *ttnpb.GatewayUplinkMessage) (*pbtypes.Empty, error) {
	if err := clusterauth.Authorized(ctx); err != nil {
		return nil, err
	}

	if s.upstreamCh == nil {
		return nil, errForwarderDisabled.New()
	}

	ctx = events.ContextWithCorrelationID(ctx, append(
		up.CorrelationIDs,
		fmt.Sprintf("pba:uplink:%s", events.NewCorrelationID()),
	)...)
	up.CorrelationIDs = events.CorrelationIDsFromContext(ctx)

	msg, err := toPBUplink(ctx, up, s.config)
	if err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to convert outgoing uplink message")
		return nil, err
	}
	if err := s.messageEncrypter.encryptUplink(ctx, msg); err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to encrypt outgoing uplink message")
		return nil, err
	}

	ctxMsg := &uplinkMessage{
		Context:       s.contextDecoupler.FromRequestContext(ctx),
		UplinkMessage: msg,
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case s.upstreamCh <- ctxMsg:
		return ttnpb.Empty, nil
	}
}

var errNoGatewayID = errors.DefineFailedPrecondition("no_gateway_id", "no gateway identifier provided or included in configuration")

// UpdateGateway is called by Gateway Server to update a gateway.
func (s *gsPbaServer) UpdateGateway(ctx context.Context, req *ttnpb.UpdatePacketBrokerGatewayRequest) (*ttnpb.UpdatePacketBrokerGatewayResponse, error) {
	if err := clusterauth.Authorized(ctx); err != nil {
		return nil, err
	}

	id := toPBGatewayIdentifier(req.Gateway.GatewayIdentifiers, s.config)
	if id == nil {
		return nil, errNoGatewayID.New()
	}
	updateReq := &mappingpb.UpdateGatewayRequest{
		ForwarderNetId:     s.netID.MarshalNumber(),
		ForwarderTenantId:  s.tenantIDExtractor(ctx),
		ForwarderClusterId: s.clusterID,
		ForwarderGatewayId: id,
		RxRate:             req.RxRate,
		TxRate:             req.TxRate,
	}

	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "location_public") {
		updateReq.GatewayLocation = &packetbroker.GatewayLocationValue{}
		if req.Gateway.LocationPublic && ttnpb.HasAnyField(req.FieldMask.GetPaths(), "antennas") && len(req.Gateway.Antennas) > 0 {
			val := &packetbroker.GatewayLocation_Terrestrial{
				AntennaCount: &pbtypes.UInt32Value{
					Value: uint32(len(req.Gateway.Antennas)),
				},
			}
			if loc := req.Gateway.Antennas[0].Location; loc.Latitude != 0 || loc.Longitude != 0 || loc.Altitude != 0 {
				val.Location = toPBLocation(&loc)
			}
			updateReq.GatewayLocation.Location = &packetbroker.GatewayLocation{
				Value: &packetbroker.GatewayLocation_Terrestrial_{
					Terrestrial: val,
				},
			}
		}
	}

	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "status_public") {
		updateReq.Online = &pbtypes.BoolValue{}
		if req.Gateway.StatusPublic && req.Online {
			updateReq.Online.Value = true
			updateReq.OnlineTtl = pbtypes.DurationProto(onlineTTL)
		}
	}

	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "contact_info") {
		adminContact, techContact := toPBContactInfo(req.Gateway.GetContactInfo())
		updateReq.AdministrativeContact = &packetbroker.ContactInfoValue{
			Value: adminContact,
		}
		updateReq.TechnicalContact = &packetbroker.ContactInfoValue{
			Value: techContact,
		}
	}

	// TODO: frequency_plan_ids

	_, err := mappingpb.NewMapperClient(s.mapperConn).UpdateGateway(ctx, updateReq)
	if err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to update gateway")
		return nil, err
	}

	res := &ttnpb.UpdatePacketBrokerGatewayResponse{}
	if updateReq.Online.GetValue() {
		res.OnlineTtl = pbtypes.DurationProto(onlineTTL)
	}
	return res, nil
}
