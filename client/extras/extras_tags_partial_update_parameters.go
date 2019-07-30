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

	models "github.com/takumin/go-netbox/models"
)

// NewExtrasTagsPartialUpdateParams creates a new ExtrasTagsPartialUpdateParams object
// with the default values initialized.
func NewExtrasTagsPartialUpdateParams() *ExtrasTagsPartialUpdateParams {
	var ()
	return &ExtrasTagsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsPartialUpdateParamsWithTimeout creates a new ExtrasTagsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTagsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsPartialUpdateParams {
	var ()
	return &ExtrasTagsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasTagsPartialUpdateParamsWithContext creates a new ExtrasTagsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTagsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasTagsPartialUpdateParams {
	var ()
	return &ExtrasTagsPartialUpdateParams{

		Context: ctx,
	}
}

// NewExtrasTagsPartialUpdateParamsWithHTTPClient creates a new ExtrasTagsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTagsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsPartialUpdateParams {
	var ()
	return &ExtrasTagsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasTagsPartialUpdateParams contains all the parameters to send to the API endpoint
for the extras tags partial update operation typically these are written to a http.Request
*/
type ExtrasTagsPartialUpdateParams struct {

	/*Data*/
	Data *models.Tag
	/*ID
	  A unique integer value identifying this tag.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasTagsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithData(data *models.Tag) *ExtrasTagsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WithID adds the id to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithID(id int64) *ExtrasTagsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
