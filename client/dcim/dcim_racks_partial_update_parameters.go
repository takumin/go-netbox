// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

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

// NewDcimRacksPartialUpdateParams creates a new DcimRacksPartialUpdateParams object
// with the default values initialized.
func NewDcimRacksPartialUpdateParams() *DcimRacksPartialUpdateParams {
	var ()
	return &DcimRacksPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksPartialUpdateParamsWithTimeout creates a new DcimRacksPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksPartialUpdateParams {
	var ()
	return &DcimRacksPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimRacksPartialUpdateParamsWithContext creates a new DcimRacksPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksPartialUpdateParamsWithContext(ctx context.Context) *DcimRacksPartialUpdateParams {
	var ()
	return &DcimRacksPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimRacksPartialUpdateParamsWithHTTPClient creates a new DcimRacksPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksPartialUpdateParams {
	var ()
	return &DcimRacksPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimRacksPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim racks partial update operation typically these are written to a http.Request
*/
type DcimRacksPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableRack
	/*ID
	  A unique integer value identifying this rack.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithContext(ctx context.Context) *DcimRacksPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithData(data *models.WritableRack) *DcimRacksPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WithID adds the id to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithID(id int64) *DcimRacksPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
