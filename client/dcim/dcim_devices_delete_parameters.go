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

// NewDcimDevicesDeleteParams creates a new DcimDevicesDeleteParams object
// with the default values initialized.
func NewDcimDevicesDeleteParams() *DcimDevicesDeleteParams {
	var ()
	return &DcimDevicesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesDeleteParamsWithTimeout creates a new DcimDevicesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDevicesDeleteParamsWithTimeout(timeout time.Duration) *DcimDevicesDeleteParams {
	var ()
	return &DcimDevicesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimDevicesDeleteParamsWithContext creates a new DcimDevicesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDevicesDeleteParamsWithContext(ctx context.Context) *DcimDevicesDeleteParams {
	var ()
	return &DcimDevicesDeleteParams{

		Context: ctx,
	}
}

// NewDcimDevicesDeleteParamsWithHTTPClient creates a new DcimDevicesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDevicesDeleteParamsWithHTTPClient(client *http.Client) *DcimDevicesDeleteParams {
	var ()
	return &DcimDevicesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimDevicesDeleteParams contains all the parameters to send to the API endpoint
for the dcim devices delete operation typically these are written to a http.Request
*/
type DcimDevicesDeleteParams struct {

	/*ID
	  A unique integer value identifying this device.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithTimeout(timeout time.Duration) *DcimDevicesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithContext(ctx context.Context) *DcimDevicesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithHTTPClient(client *http.Client) *DcimDevicesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim devices delete params
func (o *DcimDevicesDeleteParams) WithID(id int64) *DcimDevicesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices delete params
func (o *DcimDevicesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
