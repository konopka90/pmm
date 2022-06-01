// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewAddRemoteAzureDatabaseNodeParams creates a new AddRemoteAzureDatabaseNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRemoteAzureDatabaseNodeParams() *AddRemoteAzureDatabaseNodeParams {
	return &AddRemoteAzureDatabaseNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRemoteAzureDatabaseNodeParamsWithTimeout creates a new AddRemoteAzureDatabaseNodeParams object
// with the ability to set a timeout on a request.
func NewAddRemoteAzureDatabaseNodeParamsWithTimeout(timeout time.Duration) *AddRemoteAzureDatabaseNodeParams {
	return &AddRemoteAzureDatabaseNodeParams{
		timeout: timeout,
	}
}

// NewAddRemoteAzureDatabaseNodeParamsWithContext creates a new AddRemoteAzureDatabaseNodeParams object
// with the ability to set a context for a request.
func NewAddRemoteAzureDatabaseNodeParamsWithContext(ctx context.Context) *AddRemoteAzureDatabaseNodeParams {
	return &AddRemoteAzureDatabaseNodeParams{
		Context: ctx,
	}
}

// NewAddRemoteAzureDatabaseNodeParamsWithHTTPClient creates a new AddRemoteAzureDatabaseNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRemoteAzureDatabaseNodeParamsWithHTTPClient(client *http.Client) *AddRemoteAzureDatabaseNodeParams {
	return &AddRemoteAzureDatabaseNodeParams{
		HTTPClient: client,
	}
}

/* AddRemoteAzureDatabaseNodeParams contains all the parameters to send to the API endpoint
   for the add remote azure database node operation.

   Typically these are written to a http.Request.
*/
type AddRemoteAzureDatabaseNodeParams struct {
	// Body.
	Body AddRemoteAzureDatabaseNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add remote azure database node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRemoteAzureDatabaseNodeParams) WithDefaults() *AddRemoteAzureDatabaseNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add remote azure database node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRemoteAzureDatabaseNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) WithTimeout(timeout time.Duration) *AddRemoteAzureDatabaseNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) WithContext(ctx context.Context) *AddRemoteAzureDatabaseNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) WithHTTPClient(client *http.Client) *AddRemoteAzureDatabaseNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) WithBody(body AddRemoteAzureDatabaseNodeBody) *AddRemoteAzureDatabaseNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add remote azure database node params
func (o *AddRemoteAzureDatabaseNodeParams) SetBody(body AddRemoteAzureDatabaseNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRemoteAzureDatabaseNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
