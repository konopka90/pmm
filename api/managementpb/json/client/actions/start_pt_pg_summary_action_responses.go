// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartPTPgSummaryActionReader is a Reader for the StartPTPgSummaryAction structure.
type StartPTPgSummaryActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPTPgSummaryActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPTPgSummaryActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPTPgSummaryActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPTPgSummaryActionOK creates a StartPTPgSummaryActionOK with default headers values
func NewStartPTPgSummaryActionOK() *StartPTPgSummaryActionOK {
	return &StartPTPgSummaryActionOK{}
}

/* StartPTPgSummaryActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartPTPgSummaryActionOK struct {
	Payload *StartPTPgSummaryActionOKBody
}

func (o *StartPTPgSummaryActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTPgSummary][%d] startPtPgSummaryActionOk  %+v", 200, o.Payload)
}

func (o *StartPTPgSummaryActionOK) GetPayload() *StartPTPgSummaryActionOKBody {
	return o.Payload
}

func (o *StartPTPgSummaryActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTPgSummaryActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPTPgSummaryActionDefault creates a StartPTPgSummaryActionDefault with default headers values
func NewStartPTPgSummaryActionDefault(code int) *StartPTPgSummaryActionDefault {
	return &StartPTPgSummaryActionDefault{
		_statusCode: code,
	}
}

/* StartPTPgSummaryActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartPTPgSummaryActionDefault struct {
	_statusCode int

	Payload *StartPTPgSummaryActionDefaultBody
}

// Code gets the status code for the start PT pg summary action default response
func (o *StartPTPgSummaryActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPTPgSummaryActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTPgSummary][%d] StartPTPgSummaryAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartPTPgSummaryActionDefault) GetPayload() *StartPTPgSummaryActionDefaultBody {
	return o.Payload
}

func (o *StartPTPgSummaryActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTPgSummaryActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPTPgSummaryActionBody Message to prepare pt-pg-summary data
swagger:model StartPTPgSummaryActionBody
*/
type StartPTPgSummaryActionBody struct {
	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this start PT pg summary action body
func (o *StartPTPgSummaryActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT pg summary action body based on context it is used
func (o *StartPTPgSummaryActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTPgSummaryActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTPgSummaryActionBody) UnmarshalBinary(b []byte) error {
	var res StartPTPgSummaryActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTPgSummaryActionDefaultBody start PT pg summary action default body
swagger:model StartPTPgSummaryActionDefaultBody
*/
type StartPTPgSummaryActionDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartPTPgSummaryActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start PT pg summary action default body
func (o *StartPTPgSummaryActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTPgSummaryActionDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPTPgSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTPgSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start PT pg summary action default body based on the context it is used
func (o *StartPTPgSummaryActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTPgSummaryActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPTPgSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTPgSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPTPgSummaryActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTPgSummaryActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPTPgSummaryActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTPgSummaryActionDefaultBodyDetailsItems0 start PT pg summary action default body details items0
swagger:model StartPTPgSummaryActionDefaultBodyDetailsItems0
*/
type StartPTPgSummaryActionDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this start PT pg summary action default body details items0
func (o *StartPTPgSummaryActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT pg summary action default body details items0 based on context it is used
func (o *StartPTPgSummaryActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTPgSummaryActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTPgSummaryActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartPTPgSummaryActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTPgSummaryActionOKBody Message to retrieve the prepared pt-pg-summary data
swagger:model StartPTPgSummaryActionOKBody
*/
type StartPTPgSummaryActionOKBody struct {
	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start PT pg summary action OK body
func (o *StartPTPgSummaryActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT pg summary action OK body based on context it is used
func (o *StartPTPgSummaryActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTPgSummaryActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTPgSummaryActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPTPgSummaryActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
