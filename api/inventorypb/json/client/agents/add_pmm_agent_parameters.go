// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAddPMMAgentParams creates a new AddPMMAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPMMAgentParams() *AddPMMAgentParams {
	return &AddPMMAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPMMAgentParamsWithTimeout creates a new AddPMMAgentParams object
// with the ability to set a timeout on a request.
func NewAddPMMAgentParamsWithTimeout(timeout time.Duration) *AddPMMAgentParams {
	return &AddPMMAgentParams{
		timeout: timeout,
	}
}

// NewAddPMMAgentParamsWithContext creates a new AddPMMAgentParams object
// with the ability to set a context for a request.
func NewAddPMMAgentParamsWithContext(ctx context.Context) *AddPMMAgentParams {
	return &AddPMMAgentParams{
		Context: ctx,
	}
}

// NewAddPMMAgentParamsWithHTTPClient creates a new AddPMMAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPMMAgentParamsWithHTTPClient(client *http.Client) *AddPMMAgentParams {
	return &AddPMMAgentParams{
		HTTPClient: client,
	}
}

/* AddPMMAgentParams contains all the parameters to send to the API endpoint
   for the add PMM agent operation.

   Typically these are written to a http.Request.
*/
type AddPMMAgentParams struct {

	// Body.
	Body AddPMMAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add PMM agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPMMAgentParams) WithDefaults() *AddPMMAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add PMM agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPMMAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add PMM agent params
func (o *AddPMMAgentParams) WithTimeout(timeout time.Duration) *AddPMMAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add PMM agent params
func (o *AddPMMAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add PMM agent params
func (o *AddPMMAgentParams) WithContext(ctx context.Context) *AddPMMAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add PMM agent params
func (o *AddPMMAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add PMM agent params
func (o *AddPMMAgentParams) WithHTTPClient(client *http.Client) *AddPMMAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add PMM agent params
func (o *AddPMMAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add PMM agent params
func (o *AddPMMAgentParams) WithBody(body AddPMMAgentBody) *AddPMMAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add PMM agent params
func (o *AddPMMAgentParams) SetBody(body AddPMMAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddPMMAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
