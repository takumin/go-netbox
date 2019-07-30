// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package ipam

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

// NewIPAMPrefixesDeleteParams creates a new IPAMPrefixesDeleteParams object
// with the default values initialized.
func NewIPAMPrefixesDeleteParams() *IPAMPrefixesDeleteParams {
	var ()
	return &IPAMPrefixesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesDeleteParamsWithTimeout creates a new IPAMPrefixesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesDeleteParamsWithTimeout(timeout time.Duration) *IPAMPrefixesDeleteParams {
	var ()
	return &IPAMPrefixesDeleteParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesDeleteParamsWithContext creates a new IPAMPrefixesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesDeleteParamsWithContext(ctx context.Context) *IPAMPrefixesDeleteParams {
	var ()
	return &IPAMPrefixesDeleteParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesDeleteParamsWithHTTPClient creates a new IPAMPrefixesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesDeleteParamsWithHTTPClient(client *http.Client) *IPAMPrefixesDeleteParams {
	var ()
	return &IPAMPrefixesDeleteParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesDeleteParams contains all the parameters to send to the API endpoint
for the ipam prefixes delete operation typically these are written to a http.Request
*/
type IPAMPrefixesDeleteParams struct {

	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) WithTimeout(timeout time.Duration) *IPAMPrefixesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) WithContext(ctx context.Context) *IPAMPrefixesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) WithHTTPClient(client *http.Client) *IPAMPrefixesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) WithID(id int64) *IPAMPrefixesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes delete params
func (o *IPAMPrefixesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
