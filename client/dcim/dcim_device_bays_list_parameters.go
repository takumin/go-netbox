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

// NewDcimDeviceBaysListParams creates a new DcimDeviceBaysListParams object
// with the default values initialized.
func NewDcimDeviceBaysListParams() *DcimDeviceBaysListParams {
	var ()
	return &DcimDeviceBaysListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysListParamsWithTimeout creates a new DcimDeviceBaysListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBaysListParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysListParams {
	var ()
	return &DcimDeviceBaysListParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBaysListParamsWithContext creates a new DcimDeviceBaysListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBaysListParamsWithContext(ctx context.Context) *DcimDeviceBaysListParams {
	var ()
	return &DcimDeviceBaysListParams{

		Context: ctx,
	}
}

// NewDcimDeviceBaysListParamsWithHTTPClient creates a new DcimDeviceBaysListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBaysListParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysListParams {
	var ()
	return &DcimDeviceBaysListParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBaysListParams contains all the parameters to send to the API endpoint
for the dcim device bays list operation typically these are written to a http.Request
*/
type DcimDeviceBaysListParams struct {

	/*Description*/
	Description *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*ID*/
	ID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Tag*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithContext(ctx context.Context) *DcimDeviceBaysListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescription(description *string) *DcimDeviceBaysListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDevice adds the device to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDevice(device *string) *DcimDeviceBaysListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDeviceID(deviceID *string) *DcimDeviceBaysListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithID(id *string) *DcimDeviceBaysListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithLimit(limit *int64) *DcimDeviceBaysListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithName(name *string) *DcimDeviceBaysListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithOffset(offset *int64) *DcimDeviceBaysListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithQ(q *string) *DcimDeviceBaysListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithTag(tag *string) *DcimDeviceBaysListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
