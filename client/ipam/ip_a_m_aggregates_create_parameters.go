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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// NewIPAMAggregatesCreateParams creates a new IPAMAggregatesCreateParams object
// with the default values initialized.
func NewIPAMAggregatesCreateParams() *IPAMAggregatesCreateParams {
	var ()
	return &IPAMAggregatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMAggregatesCreateParamsWithTimeout creates a new IPAMAggregatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMAggregatesCreateParamsWithTimeout(timeout time.Duration) *IPAMAggregatesCreateParams {
	var ()
	return &IPAMAggregatesCreateParams{

		timeout: timeout,
	}
}

// NewIPAMAggregatesCreateParamsWithContext creates a new IPAMAggregatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMAggregatesCreateParamsWithContext(ctx context.Context) *IPAMAggregatesCreateParams {
	var ()
	return &IPAMAggregatesCreateParams{

		Context: ctx,
	}
}

// NewIPAMAggregatesCreateParamsWithHTTPClient creates a new IPAMAggregatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMAggregatesCreateParamsWithHTTPClient(client *http.Client) *IPAMAggregatesCreateParams {
	var ()
	return &IPAMAggregatesCreateParams{
		HTTPClient: client,
	}
}

/*IPAMAggregatesCreateParams contains all the parameters to send to the API endpoint
for the ipam aggregates create operation typically these are written to a http.Request
*/
type IPAMAggregatesCreateParams struct {

	/*Data*/
	Data *models.WritableAggregate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) WithTimeout(timeout time.Duration) *IPAMAggregatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) WithContext(ctx context.Context) *IPAMAggregatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) WithHTTPClient(client *http.Client) *IPAMAggregatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) WithData(data *models.WritableAggregate) *IPAMAggregatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates create params
func (o *IPAMAggregatesCreateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMAggregatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
