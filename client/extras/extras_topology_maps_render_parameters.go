// Code generated by go-swagger; DO NOT EDIT.

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

// NewExtrasTopologyMapsRenderParams creates a new ExtrasTopologyMapsRenderParams object
// with the default values initialized.
func NewExtrasTopologyMapsRenderParams() *ExtrasTopologyMapsRenderParams {
	var ()
	return &ExtrasTopologyMapsRenderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsRenderParamsWithTimeout creates a new ExtrasTopologyMapsRenderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsRenderParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsRenderParams {
	var ()
	return &ExtrasTopologyMapsRenderParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsRenderParamsWithContext creates a new ExtrasTopologyMapsRenderParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsRenderParamsWithContext(ctx context.Context) *ExtrasTopologyMapsRenderParams {
	var ()
	return &ExtrasTopologyMapsRenderParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsRenderParamsWithHTTPClient creates a new ExtrasTopologyMapsRenderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsRenderParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsRenderParams {
	var ()
	return &ExtrasTopologyMapsRenderParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsRenderParams contains all the parameters to send to the API endpoint
for the extras topology maps render operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsRenderParams struct {

	/*ID
	  A unique integer value identifying this topology map.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsRenderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) WithContext(ctx context.Context) *ExtrasTopologyMapsRenderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsRenderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) WithID(id int64) *ExtrasTopologyMapsRenderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras topology maps render params
func (o *ExtrasTopologyMapsRenderParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsRenderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
