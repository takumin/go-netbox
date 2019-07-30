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

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasCustomFieldChoicesListParams creates a new ExtrasCustomFieldChoicesListParams object
// with the default values initialized.
func NewExtrasCustomFieldChoicesListParams() *ExtrasCustomFieldChoicesListParams {

	return &ExtrasCustomFieldChoicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithTimeout creates a new ExtrasCustomFieldChoicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasCustomFieldChoicesListParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesListParams {

	return &ExtrasCustomFieldChoicesListParams{

		timeout: timeout,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithContext creates a new ExtrasCustomFieldChoicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasCustomFieldChoicesListParamsWithContext(ctx context.Context) *ExtrasCustomFieldChoicesListParams {

	return &ExtrasCustomFieldChoicesListParams{

		Context: ctx,
	}
}

// NewExtrasCustomFieldChoicesListParamsWithHTTPClient creates a new ExtrasCustomFieldChoicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasCustomFieldChoicesListParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesListParams {

	return &ExtrasCustomFieldChoicesListParams{
		HTTPClient: client,
	}
}

/*ExtrasCustomFieldChoicesListParams contains all the parameters to send to the API endpoint
for the extras custom field choices list operation typically these are written to a http.Request
*/
type ExtrasCustomFieldChoicesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldChoicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithContext(ctx context.Context) *ExtrasCustomFieldChoicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldChoicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom field choices list params
func (o *ExtrasCustomFieldChoicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldChoicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
