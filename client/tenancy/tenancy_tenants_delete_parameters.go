// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTenancyTenantsDeleteParams creates a new TenancyTenantsDeleteParams object
// with the default values initialized.
func NewTenancyTenantsDeleteParams() *TenancyTenantsDeleteParams {
	var ()
	return &TenancyTenantsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsDeleteParamsWithTimeout creates a new TenancyTenantsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenancyTenantsDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	var ()
	return &TenancyTenantsDeleteParams{

		timeout: timeout,
	}
}

// NewTenancyTenantsDeleteParamsWithContext creates a new TenancyTenantsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenancyTenantsDeleteParamsWithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	var ()
	return &TenancyTenantsDeleteParams{

		Context: ctx,
	}
}

// NewTenancyTenantsDeleteParamsWithHTTPClient creates a new TenancyTenantsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenancyTenantsDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	var ()
	return &TenancyTenantsDeleteParams{
		HTTPClient: client,
	}
}

/*TenancyTenantsDeleteParams contains all the parameters to send to the API endpoint
for the tenancy tenants delete operation typically these are written to a http.Request
*/
type TenancyTenantsDeleteParams struct {

	/*ID
	  A unique integer value identifying this tenant.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithID(id int64) *TenancyTenantsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
