// Code generated by go-swagger; DO NOT EDIT.

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

// NewDcimDeviceBayTemplatesDeleteParams creates a new DcimDeviceBayTemplatesDeleteParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesDeleteParams() *DcimDeviceBayTemplatesDeleteParams {
	var ()
	return &DcimDeviceBayTemplatesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithTimeout creates a new DcimDeviceBayTemplatesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesDeleteParams {
	var ()
	return &DcimDeviceBayTemplatesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithContext creates a new DcimDeviceBayTemplatesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesDeleteParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesDeleteParams {
	var ()
	return &DcimDeviceBayTemplatesDeleteParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesDeleteParamsWithHTTPClient creates a new DcimDeviceBayTemplatesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesDeleteParams {
	var ()
	return &DcimDeviceBayTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesDeleteParams contains all the parameters to send to the API endpoint
for the dcim device bay templates delete operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesDeleteParams struct {

	/*ID
	  A unique integer value identifying this device bay template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) WithID(id int64) *DcimDeviceBayTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates delete params
func (o *DcimDeviceBayTemplatesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
