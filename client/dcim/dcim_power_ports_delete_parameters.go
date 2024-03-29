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

// NewDcimPowerPortsDeleteParams creates a new DcimPowerPortsDeleteParams object
// with the default values initialized.
func NewDcimPowerPortsDeleteParams() *DcimPowerPortsDeleteParams {
	var ()
	return &DcimPowerPortsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsDeleteParamsWithTimeout creates a new DcimPowerPortsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPortsDeleteParams {
	var ()
	return &DcimPowerPortsDeleteParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortsDeleteParamsWithContext creates a new DcimPowerPortsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortsDeleteParamsWithContext(ctx context.Context) *DcimPowerPortsDeleteParams {
	var ()
	return &DcimPowerPortsDeleteParams{

		Context: ctx,
	}
}

// NewDcimPowerPortsDeleteParamsWithHTTPClient creates a new DcimPowerPortsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPortsDeleteParams {
	var ()
	return &DcimPowerPortsDeleteParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortsDeleteParams contains all the parameters to send to the API endpoint
for the dcim power ports delete operation typically these are written to a http.Request
*/
type DcimPowerPortsDeleteParams struct {

	/*ID
	  A unique integer value identifying this power port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithContext(ctx context.Context) *DcimPowerPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithID(id int64) *DcimPowerPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
