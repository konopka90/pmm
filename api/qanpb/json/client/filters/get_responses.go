// Code generated by go-swagger; DO NOT EDIT.

package filters

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
	"github.com/go-openapi/validate"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/* GetOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOK struct {
	Payload *GetOKBody
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/Filters/Get][%d] getOk  %+v", 200, o.Payload)
}
func (o *GetOK) GetPayload() *GetOKBody {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefault creates a GetDefault with default headers values
func NewGetDefault(code int) *GetDefault {
	return &GetDefault{
		_statusCode: code,
	}
}

/* GetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetDefault struct {
	_statusCode int

	Payload *GetDefaultBody
}

// Code gets the status code for the get default response
func (o *GetDefault) Code() int {
	return o._statusCode
}

func (o *GetDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/Filters/Get][%d] Get default  %+v", o._statusCode, o.Payload)
}
func (o *GetDefault) GetPayload() *GetDefaultBody {
	return o.Payload
}

func (o *GetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBody FiltersRequest contains period for which we need filters.
swagger:model GetBody
*/
type GetBody struct {

	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// main metric name
	MainMetricName string `json:"main_metric_name,omitempty"`

	// labels
	Labels []*GetParamsBodyLabelsItems0 `json:"labels"`
}

// Validate validates this get body
func (o *GetBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBody) validatePeriodStartFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetBody) validatePeriodStartTo(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetBody) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get body based on the context it is used
func (o *GetBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBody) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Labels); i++ {

		if o.Labels[i] != nil {
			if err := o.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBody) UnmarshalBinary(b []byte) error {
	var res GetBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDefaultBody get default body
swagger:model GetDefaultBody
*/
type GetDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get default body
func (o *GetDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Get default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Get default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get default body based on the context it is used
func (o *GetDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Get default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Get default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDefaultBodyDetailsItems0 get default body details items0
swagger:model GetDefaultBodyDetailsItems0
*/
type GetDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this get default body details items0
func (o *GetDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get default body details items0 based on context it is used
func (o *GetDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOKBody FiltersReply is map of labels for given period by key.
// Key is label's name and value is label's value and how many times it occur.
swagger:model GetOKBody
*/
type GetOKBody struct {

	// labels
	Labels map[string]GetOKBodyLabelsAnon `json:"labels,omitempty"`
}

// Validate validates this get OK body
func (o *GetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBody) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for k := range o.Labels {

		if swag.IsZero(o.Labels[k]) { // not required
			continue
		}
		if val, ok := o.Labels[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOk" + "." + "labels" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOk" + "." + "labels" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get OK body based on the context it is used
func (o *GetOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBody) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for k := range o.Labels {

		if val, ok := o.Labels[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOKBody) UnmarshalBinary(b []byte) error {
	var res GetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOKBodyLabelsAnon ListLabels is list of label's values: duplicates are impossible.
swagger:model GetOKBodyLabelsAnon
*/
type GetOKBodyLabelsAnon struct {

	// name
	Name []*GetOKBodyLabelsAnonNameItems0 `json:"name"`
}

// Validate validates this get OK body labels anon
func (o *GetOKBodyLabelsAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBodyLabelsAnon) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	for i := 0; i < len(o.Name); i++ {
		if swag.IsZero(o.Name[i]) { // not required
			continue
		}

		if o.Name[i] != nil {
			if err := o.Name[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("name" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get OK body labels anon based on the context it is used
func (o *GetOKBodyLabelsAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOKBodyLabelsAnon) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Name); i++ {

		if o.Name[i] != nil {
			if err := o.Name[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("name" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOKBodyLabelsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOKBodyLabelsAnon) UnmarshalBinary(b []byte) error {
	var res GetOKBodyLabelsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOKBodyLabelsAnonNameItems0 Values is label values and main metric percent and per second.
swagger:model GetOKBodyLabelsAnonNameItems0
*/
type GetOKBodyLabelsAnonNameItems0 struct {

	// value
	Value string `json:"value,omitempty"`

	// main metric percent
	MainMetricPercent float32 `json:"main_metric_percent,omitempty"`

	// main metric per sec
	MainMetricPerSec float32 `json:"main_metric_per_sec,omitempty"`
}

// Validate validates this get OK body labels anon name items0
func (o *GetOKBodyLabelsAnonNameItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get OK body labels anon name items0 based on context it is used
func (o *GetOKBodyLabelsAnonNameItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOKBodyLabelsAnonNameItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOKBodyLabelsAnonNameItems0) UnmarshalBinary(b []byte) error {
	var res GetOKBodyLabelsAnonNameItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetParamsBodyLabelsItems0 MapFieldEntry allows to pass labels/dimensions in form like {"server": ["db1", "db2"...]}.
swagger:model GetParamsBodyLabelsItems0
*/
type GetParamsBodyLabelsItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this get params body labels items0
func (o *GetParamsBodyLabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get params body labels items0 based on context it is used
func (o *GetParamsBodyLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetParamsBodyLabelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetParamsBodyLabelsItems0) UnmarshalBinary(b []byte) error {
	var res GetParamsBodyLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
