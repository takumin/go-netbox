// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package extras

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

// NewExtrasGraphsReadParams creates a new ExtrasGraphsReadParams object
// with the default values initialized.
func NewExtrasGraphsReadParams() *ExtrasGraphsReadParams {
	var ()
	return &ExtrasGraphsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsReadParamsWithTimeout creates a new ExtrasGraphsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsReadParamsWithTimeout(timeout time.Duration) *ExtrasGraphsReadParams {
	var ()
	return &ExtrasGraphsReadParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsReadParamsWithContext creates a new ExtrasGraphsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsReadParamsWithContext(ctx context.Context) *ExtrasGraphsReadParams {
	var ()
	return &ExtrasGraphsReadParams{

		Context: ctx,
	}
}

// NewExtrasGraphsReadParamsWithHTTPClient creates a new ExtrasGraphsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsReadParamsWithHTTPClient(client *http.Client) *ExtrasGraphsReadParams {
	var ()
	return &ExtrasGraphsReadParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsReadParams contains all the parameters to send to the API endpoint
for the extras graphs read operation typically these are written to a http.Request
*/
type ExtrasGraphsReadParams struct {

	/*ID
	  A unique integer value identifying this graph.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs read params
func (o *ExtrasGraphsReadParams) WithTimeout(timeout time.Duration) *ExtrasGraphsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs read params
func (o *ExtrasGraphsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs read params
func (o *ExtrasGraphsReadParams) WithContext(ctx context.Context) *ExtrasGraphsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs read params
func (o *ExtrasGraphsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs read params
func (o *ExtrasGraphsReadParams) WithHTTPClient(client *http.Client) *ExtrasGraphsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs read params
func (o *ExtrasGraphsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras graphs read params
func (o *ExtrasGraphsReadParams) WithID(id int64) *ExtrasGraphsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphs read params
func (o *ExtrasGraphsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
