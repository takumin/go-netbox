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

// NewDcimConsoleServerPortsReadParams creates a new DcimConsoleServerPortsReadParams object
// with the default values initialized.
func NewDcimConsoleServerPortsReadParams() *DcimConsoleServerPortsReadParams {
	var ()
	return &DcimConsoleServerPortsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsReadParamsWithTimeout creates a new DcimConsoleServerPortsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortsReadParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsReadParams {
	var ()
	return &DcimConsoleServerPortsReadParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsReadParamsWithContext creates a new DcimConsoleServerPortsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortsReadParamsWithContext(ctx context.Context) *DcimConsoleServerPortsReadParams {
	var ()
	return &DcimConsoleServerPortsReadParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortsReadParamsWithHTTPClient creates a new DcimConsoleServerPortsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortsReadParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsReadParams {
	var ()
	return &DcimConsoleServerPortsReadParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortsReadParams contains all the parameters to send to the API endpoint
for the dcim console server ports read operation typically these are written to a http.Request
*/
type DcimConsoleServerPortsReadParams struct {

	/*ID
	  A unique integer value identifying this console server port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) WithContext(ctx context.Context) *DcimConsoleServerPortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) WithID(id int64) *DcimConsoleServerPortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server ports read params
func (o *DcimConsoleServerPortsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
