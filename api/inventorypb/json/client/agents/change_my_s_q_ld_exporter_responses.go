// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

// ChangeMySQLdExporterReader is a Reader for the ChangeMySQLdExporter structure.
type ChangeMySQLdExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeMySQLdExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeMySQLdExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeMySQLdExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeMySQLdExporterOK creates a ChangeMySQLdExporterOK with default headers values
func NewChangeMySQLdExporterOK() *ChangeMySQLdExporterOK {
	return &ChangeMySQLdExporterOK{}
}

/*ChangeMySQLdExporterOK handles this case with default header values.

A successful response.
*/
type ChangeMySQLdExporterOK struct {
	Payload *ChangeMySQLdExporterOKBody
}

func (o *ChangeMySQLdExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeMySQLdExporter][%d] changeMySQLdExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeMySQLdExporterOK) GetPayload() *ChangeMySQLdExporterOKBody {
	return o.Payload
}

func (o *ChangeMySQLdExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMySQLdExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeMySQLdExporterDefault creates a ChangeMySQLdExporterDefault with default headers values
func NewChangeMySQLdExporterDefault(code int) *ChangeMySQLdExporterDefault {
	return &ChangeMySQLdExporterDefault{
		_statusCode: code,
	}
}

/*ChangeMySQLdExporterDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeMySQLdExporterDefault struct {
	_statusCode int

	Payload *ChangeMySQLdExporterDefaultBody
}

// Code gets the status code for the change my s q ld exporter default response
func (o *ChangeMySQLdExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeMySQLdExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeMySQLdExporter][%d] ChangeMySQLdExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeMySQLdExporterDefault) GetPayload() *ChangeMySQLdExporterDefaultBody {
	return o.Payload
}

func (o *ChangeMySQLdExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMySQLdExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeMySQLdExporterBody change my s q ld exporter body
swagger:model ChangeMySQLdExporterBody
*/
type ChangeMySQLdExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeMySQLdExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change my s q ld exporter body
func (o *ChangeMySQLdExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMySQLdExporterBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterDefaultBody change my s q ld exporter default body
swagger:model ChangeMySQLdExporterDefaultBody
*/
type ChangeMySQLdExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change my s q ld exporter default body
func (o *ChangeMySQLdExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMySQLdExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeMySQLdExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterOKBody change my s q ld exporter OK body
swagger:model ChangeMySQLdExporterOKBody
*/
type ChangeMySQLdExporterOKBody struct {

	// mysqld exporter
	MysqldExporter *ChangeMySQLdExporterOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`
}

// Validate validates this change my s q ld exporter OK body
func (o *ChangeMySQLdExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMySQLdExporterOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeMySQLdExporterOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model ChangeMySQLdExporterOKBodyMysqldExporter
*/
type ChangeMySQLdExporterOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// Tablestats group collectors are disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if tablestats group collectors are currently disabled.
	TablestatsGroupDisabled bool `json:"tablestats_group_disabled,omitempty"`
}

// Validate validates this change my s q ld exporter OK body mysqld exporter
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum = append(changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// ChangeMySQLdExporterOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	ChangeMySQLdExporterOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeMySQLdExporterOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeMySQLdExporterOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMySQLdExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeMySQLdExporterParamsBodyCommon
*/
type ChangeMySQLdExporterParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change my s q ld exporter params body common
func (o *ChangeMySQLdExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMySQLdExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMySQLdExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeMySQLdExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
