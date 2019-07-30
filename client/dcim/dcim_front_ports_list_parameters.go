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

// NewDcimFrontPortsListParams creates a new DcimFrontPortsListParams object
// with the default values initialized.
func NewDcimFrontPortsListParams() *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsListParamsWithTimeout creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsListParamsWithTimeout(timeout time.Duration) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsListParamsWithContext creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsListParamsWithContext(ctx context.Context) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsListParamsWithHTTPClient creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsListParamsWithHTTPClient(client *http.Client) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsListParams contains all the parameters to send to the API endpoint
for the dcim front ports list operation typically these are written to a http.Request
*/
type DcimFrontPortsListParams struct {

	/*Cabled*/
	Cabled *string
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
	/*Type*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTimeout(timeout time.Duration) *DcimFrontPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithContext(ctx context.Context) *DcimFrontPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithHTTPClient(client *http.Client) *DcimFrontPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithCabled(cabled *string) *DcimFrontPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithDescription adds the description to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescription(description *string) *DcimFrontPortsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDevice adds the device to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDevice(device *string) *DcimFrontPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDeviceID(deviceID *string) *DcimFrontPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithID(id *string) *DcimFrontPortsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithLimit(limit *int64) *DcimFrontPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithName(name *string) *DcimFrontPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithOffset(offset *int64) *DcimFrontPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithQ(q *string) *DcimFrontPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTag(tag *string) *DcimFrontPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithType adds the typeVar to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithType(typeVar *string) *DcimFrontPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
