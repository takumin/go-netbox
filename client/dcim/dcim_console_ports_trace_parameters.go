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

// NewDcimConsolePortsTraceParams creates a new DcimConsolePortsTraceParams object
// with the default values initialized.
func NewDcimConsolePortsTraceParams() *DcimConsolePortsTraceParams {
	var ()
	return &DcimConsolePortsTraceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsTraceParamsWithTimeout creates a new DcimConsolePortsTraceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsTraceParamsWithTimeout(timeout time.Duration) *DcimConsolePortsTraceParams {
	var ()
	return &DcimConsolePortsTraceParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsTraceParamsWithContext creates a new DcimConsolePortsTraceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsTraceParamsWithContext(ctx context.Context) *DcimConsolePortsTraceParams {
	var ()
	return &DcimConsolePortsTraceParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsTraceParamsWithHTTPClient creates a new DcimConsolePortsTraceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsTraceParamsWithHTTPClient(client *http.Client) *DcimConsolePortsTraceParams {
	var ()
	return &DcimConsolePortsTraceParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsTraceParams contains all the parameters to send to the API endpoint
for the dcim console ports trace operation typically these are written to a http.Request
*/
type DcimConsolePortsTraceParams struct {

	/*ID
	  A unique integer value identifying this console port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithTimeout(timeout time.Duration) *DcimConsolePortsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithContext(ctx context.Context) *DcimConsolePortsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithHTTPClient(client *http.Client) *DcimConsolePortsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) WithID(id int64) *DcimConsolePortsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports trace params
func (o *DcimConsolePortsTraceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
