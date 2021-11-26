// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'

import PageTitle from '@ttn-lw/components/page-title'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import { ApiKeyCreateForm } from '@console/components/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

const ApplicationApiKeyAdd = props => {
  const { appId, rights, pseudoRights, createApplicationApiKey, navigateToList } = props

  useBreadcrumbs(
    'apps.single.api-keys.add',
    <Breadcrumb path={`/applications/${appId}/api-keys/add`} content={sharedMessages.add} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.addApiKey} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyCreateForm
            rights={rights}
            pseudoRights={pseudoRights}
            onCreate={createApplicationApiKey}
            onCreateSuccess={navigateToList}
          />
        </Col>
      </Row>
    </Container>
  )
}

ApplicationApiKeyAdd.propTypes = {
  appId: PropTypes.string.isRequired,
  createApplicationApiKey: PropTypes.func.isRequired,
  navigateToList: PropTypes.func.isRequired,
  pseudoRights: PropTypes.rights.isRequired,
  rights: PropTypes.rights.isRequired,
}

export default ApplicationApiKeyAdd
