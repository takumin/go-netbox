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

// NewDcimDeviceTypesListParams creates a new DcimDeviceTypesListParams object
// with the default values initialized.
func NewDcimDeviceTypesListParams() *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesListParamsWithTimeout creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceTypesListParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		timeout: timeout,
	}
}

// NewDcimDeviceTypesListParamsWithContext creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceTypesListParamsWithContext(ctx context.Context) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		Context: ctx,
	}
}

// NewDcimDeviceTypesListParamsWithHTTPClient creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceTypesListParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{
		HTTPClient: client,
	}
}

/*DcimDeviceTypesListParams contains all the parameters to send to the API endpoint
for the dcim device types list operation typically these are written to a http.Request
*/
type DcimDeviceTypesListParams struct {

	/*ConsolePorts*/
	ConsolePorts *string
	/*ConsoleServerPorts*/
	ConsoleServerPorts *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Interfaces*/
	Interfaces *string
	/*IsFullDepth*/
	IsFullDepth *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *string
	/*Model*/
	Model *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*PartNumber*/
	PartNumber *string
	/*PassThroughPorts*/
	PassThroughPorts *string
	/*PowerOutlets*/
	PowerOutlets *string
	/*PowerPorts*/
	PowerPorts *string
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string
	/*SubdeviceRole*/
	SubdeviceRole *string
	/*Tag*/
	Tag *string
	/*UHeight*/
	UHeight *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithContext(ctx context.Context) *DcimDeviceTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsolePorts adds the consolePorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithConsolePorts(consolePorts *string) *DcimDeviceTypesListParams {
	o.SetConsolePorts(consolePorts)
	return o
}

// SetConsolePorts adds the consolePorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetConsolePorts(consolePorts *string) {
	o.ConsolePorts = consolePorts
}

// WithConsoleServerPorts adds the consoleServerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithConsoleServerPorts(consoleServerPorts *string) *DcimDeviceTypesListParams {
	o.SetConsoleServerPorts(consoleServerPorts)
	return o
}

// SetConsoleServerPorts adds the consoleServerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetConsoleServerPorts(consoleServerPorts *string) {
	o.ConsoleServerPorts = consoleServerPorts
}

// WithIDIn adds the iDIn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIn(iDIn *string) *DcimDeviceTypesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithInterfaces adds the interfaces to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithInterfaces(interfaces *string) *DcimDeviceTypesListParams {
	o.SetInterfaces(interfaces)
	return o
}

// SetInterfaces adds the interfaces to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetInterfaces(interfaces *string) {
	o.Interfaces = interfaces
}

// WithIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsFullDepth(isFullDepth *string) *DcimDeviceTypesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLimit(limit *int64) *DcimDeviceTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturer(manufacturer *string) *DcimDeviceTypesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturerID(manufacturerID *string) *DcimDeviceTypesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModel(model *string) *DcimDeviceTypesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModel(model *string) {
	o.Model = model
}

// WithOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithOffset(offset *int64) *DcimDeviceTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumber(partNumber *string) *DcimDeviceTypesListParams {
	o.SetPartNumber(partNumber)
	return o
}

// SetPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumber(partNumber *string) {
	o.PartNumber = partNumber
}

// WithPassThroughPorts adds the passThroughPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPassThroughPorts(passThroughPorts *string) *DcimDeviceTypesListParams {
	o.SetPassThroughPorts(passThroughPorts)
	return o
}

// SetPassThroughPorts adds the passThroughPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPassThroughPorts(passThroughPorts *string) {
	o.PassThroughPorts = passThroughPorts
}

// WithPowerOutlets adds the powerOutlets to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPowerOutlets(powerOutlets *string) *DcimDeviceTypesListParams {
	o.SetPowerOutlets(powerOutlets)
	return o
}

// SetPowerOutlets adds the powerOutlets to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPowerOutlets(powerOutlets *string) {
	o.PowerOutlets = powerOutlets
}

// WithPowerPorts adds the powerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPowerPorts(powerPorts *string) *DcimDeviceTypesListParams {
	o.SetPowerPorts(powerPorts)
	return o
}

// SetPowerPorts adds the powerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPowerPorts(powerPorts *string) {
	o.PowerPorts = powerPorts
}

// WithQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithQ(q *string) *DcimDeviceTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlug(slug *string) *DcimDeviceTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSubdeviceRole(subdeviceRole *string) *DcimDeviceTypesListParams {
	o.SetSubdeviceRole(subdeviceRole)
	return o
}

// SetSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSubdeviceRole(subdeviceRole *string) {
	o.SubdeviceRole = subdeviceRole
}

// WithTag adds the tag to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTag(tag *string) *DcimDeviceTypesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeight(uHeight *string) *DcimDeviceTypesListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeight(uHeight *string) {
	o.UHeight = uHeight
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConsolePorts != nil {

		// query param console_ports
		var qrConsolePorts string
		if o.ConsolePorts != nil {
			qrConsolePorts = *o.ConsolePorts
		}
		qConsolePorts := qrConsolePorts
		if qConsolePorts != "" {
			if err := r.SetQueryParam("console_ports", qConsolePorts); err != nil {
				return err
			}
		}

	}

	if o.ConsoleServerPorts != nil {

		// query param console_server_ports
		var qrConsoleServerPorts string
		if o.ConsoleServerPorts != nil {
			qrConsoleServerPorts = *o.ConsoleServerPorts
		}
		qConsoleServerPorts := qrConsoleServerPorts
		if qConsoleServerPorts != "" {
			if err := r.SetQueryParam("console_server_ports", qConsoleServerPorts); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.Interfaces != nil {

		// query param interfaces
		var qrInterfaces string
		if o.Interfaces != nil {
			qrInterfaces = *o.Interfaces
		}
		qInterfaces := qrInterfaces
		if qInterfaces != "" {
			if err := r.SetQueryParam("interfaces", qInterfaces); err != nil {
				return err
			}
		}

	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string
		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {
			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
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

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID string
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := qrManufacturerID
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}

	}

	if o.Model != nil {

		// query param model
		var qrModel string
		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {
			if err := r.SetQueryParam("model", qModel); err != nil {
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

	if o.PartNumber != nil {

		// query param part_number
		var qrPartNumber string
		if o.PartNumber != nil {
			qrPartNumber = *o.PartNumber
		}
		qPartNumber := qrPartNumber
		if qPartNumber != "" {
			if err := r.SetQueryParam("part_number", qPartNumber); err != nil {
				return err
			}
		}

	}

	if o.PassThroughPorts != nil {

		// query param pass_through_ports
		var qrPassThroughPorts string
		if o.PassThroughPorts != nil {
			qrPassThroughPorts = *o.PassThroughPorts
		}
		qPassThroughPorts := qrPassThroughPorts
		if qPassThroughPorts != "" {
			if err := r.SetQueryParam("pass_through_ports", qPassThroughPorts); err != nil {
				return err
			}
		}

	}

	if o.PowerOutlets != nil {

		// query param power_outlets
		var qrPowerOutlets string
		if o.PowerOutlets != nil {
			qrPowerOutlets = *o.PowerOutlets
		}
		qPowerOutlets := qrPowerOutlets
		if qPowerOutlets != "" {
			if err := r.SetQueryParam("power_outlets", qPowerOutlets); err != nil {
				return err
			}
		}

	}

	if o.PowerPorts != nil {

		// query param power_ports
		var qrPowerPorts string
		if o.PowerPorts != nil {
			qrPowerPorts = *o.PowerPorts
		}
		qPowerPorts := qrPowerPorts
		if qPowerPorts != "" {
			if err := r.SetQueryParam("power_ports", qPowerPorts); err != nil {
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

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SubdeviceRole != nil {

		// query param subdevice_role
		var qrSubdeviceRole string
		if o.SubdeviceRole != nil {
			qrSubdeviceRole = *o.SubdeviceRole
		}
		qSubdeviceRole := qrSubdeviceRole
		if qSubdeviceRole != "" {
			if err := r.SetQueryParam("subdevice_role", qSubdeviceRole); err != nil {
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

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight string
		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := qrUHeight
		if qUHeight != "" {
			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
