// Code generated by go-swagger; DO NOT EDIT.

package restore_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListRestoreHistoryReader is a Reader for the ListRestoreHistory structure.
type ListRestoreHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRestoreHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRestoreHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRestoreHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRestoreHistoryOK creates a ListRestoreHistoryOK with default headers values
func NewListRestoreHistoryOK() *ListRestoreHistoryOK {
	return &ListRestoreHistoryOK{}
}

/* ListRestoreHistoryOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRestoreHistoryOK struct {
	Payload *ListRestoreHistoryOKBody
}

func (o *ListRestoreHistoryOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/RestoreHistory/List][%d] listRestoreHistoryOk  %+v", 200, o.Payload)
}
func (o *ListRestoreHistoryOK) GetPayload() *ListRestoreHistoryOKBody {
	return o.Payload
}

func (o *ListRestoreHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListRestoreHistoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRestoreHistoryDefault creates a ListRestoreHistoryDefault with default headers values
func NewListRestoreHistoryDefault(code int) *ListRestoreHistoryDefault {
	return &ListRestoreHistoryDefault{
		_statusCode: code,
	}
}

/* ListRestoreHistoryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListRestoreHistoryDefault struct {
	_statusCode int

	Payload *ListRestoreHistoryDefaultBody
}

// Code gets the status code for the list restore history default response
func (o *ListRestoreHistoryDefault) Code() int {
	return o._statusCode
}

func (o *ListRestoreHistoryDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/RestoreHistory/List][%d] ListRestoreHistory default  %+v", o._statusCode, o.Payload)
}
func (o *ListRestoreHistoryDefault) GetPayload() *ListRestoreHistoryDefaultBody {
	return o.Payload
}

func (o *ListRestoreHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListRestoreHistoryDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListRestoreHistoryDefaultBody list restore history default body
swagger:model ListRestoreHistoryDefaultBody
*/
type ListRestoreHistoryDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListRestoreHistoryDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list restore history default body
func (o *ListRestoreHistoryDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListRestoreHistory default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRestoreHistory default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list restore history default body based on the context it is used
func (o *ListRestoreHistoryDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListRestoreHistory default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRestoreHistory default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListRestoreHistoryDefaultBodyDetailsItems0 list restore history default body details items0
swagger:model ListRestoreHistoryDefaultBodyDetailsItems0
*/
type ListRestoreHistoryDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this list restore history default body details items0
func (o *ListRestoreHistoryDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list restore history default body details items0 based on context it is used
func (o *ListRestoreHistoryDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListRestoreHistoryOKBody list restore history OK body
swagger:model ListRestoreHistoryOKBody
*/
type ListRestoreHistoryOKBody struct {

	// items
	Items []*ListRestoreHistoryOKBodyItemsItems0 `json:"items"`
}

// Validate validates this list restore history OK body
func (o *ListRestoreHistoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRestoreHistoryOk" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRestoreHistoryOk" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list restore history OK body based on the context it is used
func (o *ListRestoreHistoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoreHistoryOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRestoreHistoryOk" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRestoreHistoryOk" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryOKBody) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListRestoreHistoryOKBodyItemsItems0 RestoreHistoryItem represents single backup restore item.
swagger:model ListRestoreHistoryOKBodyItemsItems0
*/
type ListRestoreHistoryOKBodyItemsItems0 struct {

	// Machine-readable restore id.
	RestoreID string `json:"restore_id,omitempty"`

	// ID of the artifact used for restore.
	ArtifactID string `json:"artifact_id,omitempty"`

	// Artifact name used for restore.
	Name string `json:"name,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name.
	LocationName string `json:"location_name,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"service_id,omitempty"`

	// Service name.
	ServiceName string `json:"service_name,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// RestoreStatus shows the current status of execution of restore.
	// Enum: [RESTORE_STATUS_INVALID RESTORE_STATUS_IN_PROGRESS RESTORE_STATUS_SUCCESS RESTORE_STATUS_ERROR]
	Status *string `json:"status,omitempty"`

	// Restore start time.
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Restore finish time.
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`
}

// Validate validates this list restore history OK body items items0
func (o *ListRestoreHistoryOKBodyItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listRestoreHistoryOkBodyItemsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listRestoreHistoryOkBodyItemsItems0TypeDataModelPropEnum = append(listRestoreHistoryOkBodyItemsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ListRestoreHistoryOKBodyItemsItems0DataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ListRestoreHistoryOKBodyItemsItems0DataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ListRestoreHistoryOKBodyItemsItems0DataModelPHYSICAL captures enum value "PHYSICAL"
	ListRestoreHistoryOKBodyItemsItems0DataModelPHYSICAL string = "PHYSICAL"

	// ListRestoreHistoryOKBodyItemsItems0DataModelLOGICAL captures enum value "LOGICAL"
	ListRestoreHistoryOKBodyItemsItems0DataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ListRestoreHistoryOKBodyItemsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listRestoreHistoryOkBodyItemsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListRestoreHistoryOKBodyItemsItems0) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

var listRestoreHistoryOkBodyItemsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RESTORE_STATUS_INVALID","RESTORE_STATUS_IN_PROGRESS","RESTORE_STATUS_SUCCESS","RESTORE_STATUS_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listRestoreHistoryOkBodyItemsItems0TypeStatusPropEnum = append(listRestoreHistoryOkBodyItemsItems0TypeStatusPropEnum, v)
	}
}

const (

	// ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSINVALID captures enum value "RESTORE_STATUS_INVALID"
	ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSINVALID string = "RESTORE_STATUS_INVALID"

	// ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSINPROGRESS captures enum value "RESTORE_STATUS_IN_PROGRESS"
	ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSINPROGRESS string = "RESTORE_STATUS_IN_PROGRESS"

	// ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSSUCCESS captures enum value "RESTORE_STATUS_SUCCESS"
	ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSSUCCESS string = "RESTORE_STATUS_SUCCESS"

	// ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSERROR captures enum value "RESTORE_STATUS_ERROR"
	ListRestoreHistoryOKBodyItemsItems0StatusRESTORESTATUSERROR string = "RESTORE_STATUS_ERROR"
)

// prop value enum
func (o *ListRestoreHistoryOKBodyItemsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listRestoreHistoryOkBodyItemsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListRestoreHistoryOKBodyItemsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ListRestoreHistoryOKBodyItemsItems0) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", o.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListRestoreHistoryOKBodyItemsItems0) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", o.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list restore history OK body items items0 based on context it is used
func (o *ListRestoreHistoryOKBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoreHistoryOKBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoreHistoryOKBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res ListRestoreHistoryOKBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
