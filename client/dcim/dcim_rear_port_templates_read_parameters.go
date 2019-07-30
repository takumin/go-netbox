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

// NewDcimRearPortTemplatesReadParams creates a new DcimRearPortTemplatesReadParams object
// with the default values initialized.
func NewDcimRearPortTemplatesReadParams() *DcimRearPortTemplatesReadParams {
	var ()
	return &DcimRearPortTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesReadParamsWithTimeout creates a new DcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRearPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesReadParams {
	var ()
	return &DcimRearPortTemplatesReadParams{

		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesReadParamsWithContext creates a new DcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRearPortTemplatesReadParamsWithContext(ctx context.Context) *DcimRearPortTemplatesReadParams {
	var ()
	return &DcimRearPortTemplatesReadParams{

		Context: ctx,
	}
}

// NewDcimRearPortTemplatesReadParamsWithHTTPClient creates a new DcimRearPortTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRearPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesReadParams {
	var ()
	return &DcimRearPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/*DcimRearPortTemplatesReadParams contains all the parameters to send to the API endpoint
for the dcim rear port templates read operation typically these are written to a http.Request
*/
type DcimRearPortTemplatesReadParams struct {

	/*ID
	  A unique integer value identifying this rear port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithContext(ctx context.Context) *DcimRearPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) WithID(id int64) *DcimRearPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates read params
func (o *DcimRearPortTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
