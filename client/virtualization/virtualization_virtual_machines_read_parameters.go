// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package virtualization

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

// NewVirtualizationVirtualMachinesReadParams creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized.
func NewVirtualizationVirtualMachinesReadParams() *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithTimeout creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationVirtualMachinesReadParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithContext creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationVirtualMachinesReadParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{

		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithHTTPClient creates a new VirtualizationVirtualMachinesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationVirtualMachinesReadParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	var ()
	return &VirtualizationVirtualMachinesReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationVirtualMachinesReadParams contains all the parameters to send to the API endpoint
for the virtualization virtual machines read operation typically these are written to a http.Request
*/
type VirtualizationVirtualMachinesReadParams struct {

	/*ID
	  A unique integer value identifying this virtual machine.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithID(id int64) *VirtualizationVirtualMachinesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
