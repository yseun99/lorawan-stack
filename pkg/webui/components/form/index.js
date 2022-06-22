// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

/* eslint-disable react/sort-prop-types */
import React, { useCallback, useEffect } from 'react'
import { Formik, yupToFormErrors, useFormikContext, validateYupSchema } from 'formik'
import scrollIntoView from 'scroll-into-view-if-needed'
import { defineMessages } from 'react-intl'

import Notification from '@ttn-lw/components/notification'
import ErrorNotification from '@ttn-lw/components/error-notification'

import PropTypes from '@ttn-lw/lib/prop-types'
import { ingestError } from '@ttn-lw/lib/errors/utils'

import FormContext from './context'
import FormField from './field'
import FormInfoField from './field/info'
import FormSubmit from './submit'
import FormCollapseSection from './section'
import FormSubTitle from './sub-title'
import FormFieldContainer from './field/container'

const m = defineMessages({
  submitFailed: 'Submit failed',
})

const InnerForm = props => {
  const {
    formError,
    isSubmitting,
    isValid,
    className,
    children,
    formErrorTitle,
    formInfo,
    formInfoTitle,
    handleSubmit,
    id,
    ...rest
  } = props
  const notificationRef = React.useRef()

  useEffect(() => {
    // Scroll form notification into view if needed.
    if (formError) {
      scrollIntoView(notificationRef.current, { behavior: 'smooth' })
      notificationRef.current.focus({ preventScroll: true })
    }

    // Scroll invalid fields into view if needed and focus them.
    if (!isSubmitting && !isValid) {
      const firstErrorNode = document.querySelectorAll('[data-needs-focus="true"]')[0]
      if (firstErrorNode) {
        scrollIntoView(firstErrorNode, { behavior: 'smooth' })
        firstErrorNode.querySelector('input,textarea,canvas,video').focus({ preventScroll: true })
      }
    }
  }, [formError, isSubmitting, isValid])

  return (
    <form className={className} onSubmit={handleSubmit} id={id}>
      {(formError || formInfo) && (
        <div style={{ outline: 'none' }} ref={notificationRef} tabIndex="-1">
          {formError && <ErrorNotification content={formError} title={formErrorTitle} small />}
          {formInfo && <Notification content={formInfo} title={formInfoTitle} info small />}
        </div>
      )}
      <FormContext.Provider
        value={{
          formError,
          ...rest,
        }}
      >
        {children}
      </FormContext.Provider>
    </form>
  )
}
InnerForm.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  id: PropTypes.string,
  formError: PropTypes.error,
  formErrorTitle: PropTypes.message,
  formInfo: PropTypes.message,
  formInfoTitle: PropTypes.message,
  handleSubmit: PropTypes.func.isRequired,
  isSubmitting: PropTypes.bool.isRequired,
  isValid: PropTypes.bool.isRequired,
}

InnerForm.defaultProps = {
  className: undefined,
  id: undefined,
  formInfo: undefined,
  formInfoTitle: undefined,
  formError: undefined,
  formErrorTitle: m.submitFailed,
}

const Form = props => {
  const {
    children,
    className,
    disabled,
    enableReinitialize,
    error,
    errorTitle,
    formikRef,
    id,
    info,
    infoTitle,
    initialValues,
    onReset,
    onSubmit,
    validateOnBlur,
    validateOnChange,
    validateOnMount,
    validateSync,
    validationContext,
    validationSchema,
  } = props

  const handleSubmit = useCallback(
    async (...args) => {
      try {
        return await onSubmit(...args)
      } catch (error) {
        // Make sure all unhandled exceptions during submit are ingested.
        ingestError(error, { ingestedBy: 'FormSubmit' })

        throw error
      }
    },
    [onSubmit],
  )

  const validate = useEffect(
    values => {
      if (!validationSchema) {
        return {}
      }

      if (validateSync) {
        try {
          validateYupSchema(values, validationSchema, validateSync, validationContext)

          return {}
        } catch (err) {
          if (err.name === 'ValidationError') {
            return yupToFormErrors(err)
          }

          throw error
        }
      }

      return new Promise((resolve, reject) => {
        validateYupSchema(values, validationSchema, validateSync, validationContext).then(
          () => {
            resolve({})
          },
          err => {
            // Resolve yup errors, see https://jaredpalmer.com/formik/docs/migrating-v2#validate.
            if (err.name === 'ValidationError') {
              resolve(yupToFormErrors(err))
            } else {
              // Throw any other errors as it is not related to the validation process.
              reject(err)
            }
          },
        )
      })
    },
    [validationSchema, validateSync, validationContext, error],
  )

  return (
    <Formik
      innerRef={formikRef}
      validate={validate}
      onSubmit={handleSubmit}
      onReset={onReset}
      validateOnMount={validateOnMount}
      initialValues={initialValues}
      validateOnBlur={validateOnBlur}
      validateSync={validateSync}
      validateOnChange={validateOnChange}
      enableReinitialize={enableReinitialize}
    >
      {({ handleSubmit, ...restFormikProps }) => (
        <InnerForm
          className={className}
          formError={error}
          formErrorTitle={errorTitle}
          formInfo={info}
          formInfoTitle={infoTitle}
          handleSubmit={handleSubmit}
          disabled={disabled}
          id={id}
          {...restFormikProps}
        >
          {children}
        </InnerForm>
      )}
    </Formik>
  )
}

Form.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  disabled: PropTypes.bool,
  enableReinitialize: PropTypes.bool,
  error: PropTypes.error,
  errorTitle: PropTypes.message,
  info: PropTypes.message,
  infoTitle: PropTypes.message,
  formikRef: PropTypes.shape({ current: PropTypes.shape({}) }),
  id: PropTypes.string,
  initialValues: PropTypes.shape({}),
  onReset: PropTypes.func,
  onSubmit: PropTypes.func,
  validateOnBlur: PropTypes.bool,
  validateOnChange: PropTypes.bool,
  validateOnMount: PropTypes.bool,
  validateSync: PropTypes.bool,
  validationContext: PropTypes.shape({}),
  validationSchema: PropTypes.oneOfType([PropTypes.shape({}), PropTypes.func]),
}

Form.defaultProps = {
  className: undefined,
  disabled: false,
  enableReinitialize: false,
  error: undefined,
  errorTitle: m.submitFailed,
  info: undefined,
  infoTitle: undefined,
  formikRef: undefined,
  id: undefined,
  initialValues: undefined,
  onReset: () => null,
  onSubmit: () => null,
  validateOnBlur: true,
  validateOnChange: false,
  validateOnMount: false,
  validateSync: true,
  validationContext: {},
  validationSchema: undefined,
}

Form.Field = FormField
Form.InfoField = FormInfoField
Form.Submit = FormSubmit
Form.CollapseSection = FormCollapseSection
Form.SubTitle = FormSubTitle
Form.FieldContainer = FormFieldContainer

export { Form as default, useFormikContext as useFormContext }
