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

// NewVirtualizationClustersDeleteParams creates a new VirtualizationClustersDeleteParams object
// with the default values initialized.
func NewVirtualizationClustersDeleteParams() *VirtualizationClustersDeleteParams {
	var ()
	return &VirtualizationClustersDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersDeleteParamsWithTimeout creates a new VirtualizationClustersDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClustersDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClustersDeleteParams {
	var ()
	return &VirtualizationClustersDeleteParams{

		timeout: timeout,
	}
}

// NewVirtualizationClustersDeleteParamsWithContext creates a new VirtualizationClustersDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClustersDeleteParamsWithContext(ctx context.Context) *VirtualizationClustersDeleteParams {
	var ()
	return &VirtualizationClustersDeleteParams{

		Context: ctx,
	}
}

// NewVirtualizationClustersDeleteParamsWithHTTPClient creates a new VirtualizationClustersDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClustersDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClustersDeleteParams {
	var ()
	return &VirtualizationClustersDeleteParams{
		HTTPClient: client,
	}
}

/*VirtualizationClustersDeleteParams contains all the parameters to send to the API endpoint
for the virtualization clusters delete operation typically these are written to a http.Request
*/
type VirtualizationClustersDeleteParams struct {

	/*ID
	  A unique integer value identifying this cluster.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClustersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithContext(ctx context.Context) *VirtualizationClustersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClustersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithID(id int64) *VirtualizationClustersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
