// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package circuits

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

// NewCircuitsCircuitTypesDeleteParams creates a new CircuitsCircuitTypesDeleteParams object
// with the default values initialized.
func NewCircuitsCircuitTypesDeleteParams() *CircuitsCircuitTypesDeleteParams {
	var ()
	return &CircuitsCircuitTypesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithTimeout creates a new CircuitsCircuitTypesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesDeleteParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesDeleteParams {
	var ()
	return &CircuitsCircuitTypesDeleteParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithContext creates a new CircuitsCircuitTypesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesDeleteParamsWithContext(ctx context.Context) *CircuitsCircuitTypesDeleteParams {
	var ()
	return &CircuitsCircuitTypesDeleteParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithHTTPClient creates a new CircuitsCircuitTypesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesDeleteParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesDeleteParams {
	var ()
	return &CircuitsCircuitTypesDeleteParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesDeleteParams contains all the parameters to send to the API endpoint
for the circuits circuit types delete operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesDeleteParams struct {

	/*ID
	  A unique integer value identifying this circuit type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithContext(ctx context.Context) *CircuitsCircuitTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithID(id int64) *CircuitsCircuitTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
