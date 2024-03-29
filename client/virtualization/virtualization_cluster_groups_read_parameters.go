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

// NewVirtualizationClusterGroupsReadParams creates a new VirtualizationClusterGroupsReadParams object
// with the default values initialized.
func NewVirtualizationClusterGroupsReadParams() *VirtualizationClusterGroupsReadParams {
	var ()
	return &VirtualizationClusterGroupsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsReadParamsWithTimeout creates a new VirtualizationClusterGroupsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterGroupsReadParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsReadParams {
	var ()
	return &VirtualizationClusterGroupsReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsReadParamsWithContext creates a new VirtualizationClusterGroupsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterGroupsReadParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsReadParams {
	var ()
	return &VirtualizationClusterGroupsReadParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsReadParamsWithHTTPClient creates a new VirtualizationClusterGroupsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterGroupsReadParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsReadParams {
	var ()
	return &VirtualizationClusterGroupsReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterGroupsReadParams contains all the parameters to send to the API endpoint
for the virtualization cluster groups read operation typically these are written to a http.Request
*/
type VirtualizationClusterGroupsReadParams struct {

	/*ID
	  A unique integer value identifying this cluster group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) WithID(id int64) *VirtualizationClusterGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups read params
func (o *VirtualizationClusterGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
