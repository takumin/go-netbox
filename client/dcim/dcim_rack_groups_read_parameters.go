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

// NewDcimRackGroupsReadParams creates a new DcimRackGroupsReadParams object
// with the default values initialized.
func NewDcimRackGroupsReadParams() *DcimRackGroupsReadParams {
	var ()
	return &DcimRackGroupsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsReadParamsWithTimeout creates a new DcimRackGroupsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsReadParamsWithTimeout(timeout time.Duration) *DcimRackGroupsReadParams {
	var ()
	return &DcimRackGroupsReadParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsReadParamsWithContext creates a new DcimRackGroupsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsReadParamsWithContext(ctx context.Context) *DcimRackGroupsReadParams {
	var ()
	return &DcimRackGroupsReadParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsReadParamsWithHTTPClient creates a new DcimRackGroupsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsReadParamsWithHTTPClient(client *http.Client) *DcimRackGroupsReadParams {
	var ()
	return &DcimRackGroupsReadParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsReadParams contains all the parameters to send to the API endpoint
for the dcim rack groups read operation typically these are written to a http.Request
*/
type DcimRackGroupsReadParams struct {

	/*ID
	  A unique integer value identifying this rack group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithTimeout(timeout time.Duration) *DcimRackGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithContext(ctx context.Context) *DcimRackGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithHTTPClient(client *http.Client) *DcimRackGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) WithID(id int64) *DcimRackGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups read params
func (o *DcimRackGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
