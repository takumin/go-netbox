// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

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

// NewDcimDeviceRolesDeleteParams creates a new DcimDeviceRolesDeleteParams object
// with the default values initialized.
func NewDcimDeviceRolesDeleteParams() *DcimDeviceRolesDeleteParams {
	var ()
	return &DcimDeviceRolesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesDeleteParamsWithTimeout creates a new DcimDeviceRolesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceRolesDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesDeleteParams {
	var ()
	return &DcimDeviceRolesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimDeviceRolesDeleteParamsWithContext creates a new DcimDeviceRolesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceRolesDeleteParamsWithContext(ctx context.Context) *DcimDeviceRolesDeleteParams {
	var ()
	return &DcimDeviceRolesDeleteParams{

		Context: ctx,
	}
}

// NewDcimDeviceRolesDeleteParamsWithHTTPClient creates a new DcimDeviceRolesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceRolesDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesDeleteParams {
	var ()
	return &DcimDeviceRolesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimDeviceRolesDeleteParams contains all the parameters to send to the API endpoint
for the dcim device roles delete operation typically these are written to a http.Request
*/
type DcimDeviceRolesDeleteParams struct {

	/*ID
	  A unique integer value identifying this device role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithContext(ctx context.Context) *DcimDeviceRolesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) WithID(id int64) *DcimDeviceRolesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device roles delete params
func (o *DcimDeviceRolesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
