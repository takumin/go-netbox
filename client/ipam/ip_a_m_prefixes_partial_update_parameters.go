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

	models "github.com/takumin/go-netbox/models"
)

// NewIPAMPrefixesPartialUpdateParams creates a new IPAMPrefixesPartialUpdateParams object
// with the default values initialized.
func NewIPAMPrefixesPartialUpdateParams() *IPAMPrefixesPartialUpdateParams {
	var ()
	return &IPAMPrefixesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesPartialUpdateParamsWithTimeout creates a new IPAMPrefixesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesPartialUpdateParamsWithTimeout(timeout time.Duration) *IPAMPrefixesPartialUpdateParams {
	var ()
	return &IPAMPrefixesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesPartialUpdateParamsWithContext creates a new IPAMPrefixesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesPartialUpdateParamsWithContext(ctx context.Context) *IPAMPrefixesPartialUpdateParams {
	var ()
	return &IPAMPrefixesPartialUpdateParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesPartialUpdateParamsWithHTTPClient creates a new IPAMPrefixesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesPartialUpdateParamsWithHTTPClient(client *http.Client) *IPAMPrefixesPartialUpdateParams {
	var ()
	return &IPAMPrefixesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam prefixes partial update operation typically these are written to a http.Request
*/
type IPAMPrefixesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritablePrefix
	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) WithTimeout(timeout time.Duration) *IPAMPrefixesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) WithContext(ctx context.Context) *IPAMPrefixesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) WithHTTPClient(client *http.Client) *IPAMPrefixesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) WithData(data *models.WritablePrefix) *IPAMPrefixesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) WithID(id int64) *IPAMPrefixesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes partial update params
func (o *IPAMPrefixesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
