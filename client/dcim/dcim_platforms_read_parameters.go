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
)

// NewDcimPlatformsReadParams creates a new DcimPlatformsReadParams object
// with the default values initialized.
func NewDcimPlatformsReadParams() *DcimPlatformsReadParams {
	var ()
	return &DcimPlatformsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsReadParamsWithTimeout creates a new DcimPlatformsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPlatformsReadParamsWithTimeout(timeout time.Duration) *DcimPlatformsReadParams {
	var ()
	return &DcimPlatformsReadParams{

		timeout: timeout,
	}
}

// NewDcimPlatformsReadParamsWithContext creates a new DcimPlatformsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPlatformsReadParamsWithContext(ctx context.Context) *DcimPlatformsReadParams {
	var ()
	return &DcimPlatformsReadParams{

		Context: ctx,
	}
}

// NewDcimPlatformsReadParamsWithHTTPClient creates a new DcimPlatformsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPlatformsReadParamsWithHTTPClient(client *http.Client) *DcimPlatformsReadParams {
	var ()
	return &DcimPlatformsReadParams{
		HTTPClient: client,
	}
}

/*DcimPlatformsReadParams contains all the parameters to send to the API endpoint
for the dcim platforms read operation typically these are written to a http.Request
*/
type DcimPlatformsReadParams struct {

	/*ID
	  A unique integer value identifying this platform.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithTimeout(timeout time.Duration) *DcimPlatformsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithContext(ctx context.Context) *DcimPlatformsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithHTTPClient(client *http.Client) *DcimPlatformsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithID(id int64) *DcimPlatformsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
