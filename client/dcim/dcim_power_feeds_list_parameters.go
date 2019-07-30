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

// NewDcimPowerFeedsListParams creates a new DcimPowerFeedsListParams object
// with the default values initialized.
func NewDcimPowerFeedsListParams() *DcimPowerFeedsListParams {
	var ()
	return &DcimPowerFeedsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsListParamsWithTimeout creates a new DcimPowerFeedsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerFeedsListParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsListParams {
	var ()
	return &DcimPowerFeedsListParams{

		timeout: timeout,
	}
}

// NewDcimPowerFeedsListParamsWithContext creates a new DcimPowerFeedsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerFeedsListParamsWithContext(ctx context.Context) *DcimPowerFeedsListParams {
	var ()
	return &DcimPowerFeedsListParams{

		Context: ctx,
	}
}

// NewDcimPowerFeedsListParamsWithHTTPClient creates a new DcimPowerFeedsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerFeedsListParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsListParams {
	var ()
	return &DcimPowerFeedsListParams{
		HTTPClient: client,
	}
}

/*DcimPowerFeedsListParams contains all the parameters to send to the API endpoint
for the dcim power feeds list operation typically these are written to a http.Request
*/
type DcimPowerFeedsListParams struct {

	/*Amperage*/
	Amperage *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MaxUtilization*/
	MaxUtilization *string
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Phase*/
	Phase *string
	/*PowerPanelID*/
	PowerPanelID *string
	/*Q*/
	Q *string
	/*RackID*/
	RackID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Status*/
	Status *string
	/*Supply*/
	Supply *string
	/*Tag*/
	Tag *string
	/*Type*/
	Type *string
	/*Voltage*/
	Voltage *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithContext(ctx context.Context) *DcimPowerFeedsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmperage adds the amperage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperage(amperage *string) *DcimPowerFeedsListParams {
	o.SetAmperage(amperage)
	return o
}

// SetAmperage adds the amperage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperage(amperage *string) {
	o.Amperage = amperage
}

// WithIDIn adds the iDIn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDIn(iDIn *string) *DcimPowerFeedsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithLimit(limit *int64) *DcimPowerFeedsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxUtilization adds the maxUtilization to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilization(maxUtilization *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilization(maxUtilization)
	return o
}

// SetMaxUtilization adds the maxUtilization to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilization(maxUtilization *string) {
	o.MaxUtilization = maxUtilization
}

// WithName adds the name to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithName(name *string) *DcimPowerFeedsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithOffset(offset *int64) *DcimPowerFeedsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPhase adds the phase to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPhase(phase *string) *DcimPowerFeedsListParams {
	o.SetPhase(phase)
	return o
}

// SetPhase adds the phase to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPhase(phase *string) {
	o.Phase = phase
}

// WithPowerPanelID adds the powerPanelID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPowerPanelID(powerPanelID *string) *DcimPowerFeedsListParams {
	o.SetPowerPanelID(powerPanelID)
	return o
}

// SetPowerPanelID adds the powerPanelId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPowerPanelID(powerPanelID *string) {
	o.PowerPanelID = powerPanelID
}

// WithQ adds the q to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithQ(q *string) *DcimPowerFeedsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackID adds the rackID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRackID(rackID *string) *DcimPowerFeedsListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithSite adds the site to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSite(site *string) *DcimPowerFeedsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSiteID(siteID *string) *DcimPowerFeedsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithStatus(status *string) *DcimPowerFeedsListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetStatus(status *string) {
	o.Status = status
}

// WithSupply adds the supply to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSupply(supply *string) *DcimPowerFeedsListParams {
	o.SetSupply(supply)
	return o
}

// SetSupply adds the supply to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSupply(supply *string) {
	o.Supply = supply
}

// WithTag adds the tag to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTag(tag *string) *DcimPowerFeedsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithType adds the typeVar to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithType(typeVar *string) *DcimPowerFeedsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithVoltage adds the voltage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltage(voltage *string) *DcimPowerFeedsListParams {
	o.SetVoltage(voltage)
	return o
}

// SetVoltage adds the voltage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltage(voltage *string) {
	o.Voltage = voltage
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Amperage != nil {

		// query param amperage
		var qrAmperage string
		if o.Amperage != nil {
			qrAmperage = *o.Amperage
		}
		qAmperage := qrAmperage
		if qAmperage != "" {
			if err := r.SetQueryParam("amperage", qAmperage); err != nil {
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

	if o.MaxUtilization != nil {

		// query param max_utilization
		var qrMaxUtilization string
		if o.MaxUtilization != nil {
			qrMaxUtilization = *o.MaxUtilization
		}
		qMaxUtilization := qrMaxUtilization
		if qMaxUtilization != "" {
			if err := r.SetQueryParam("max_utilization", qMaxUtilization); err != nil {
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

	if o.Phase != nil {

		// query param phase
		var qrPhase string
		if o.Phase != nil {
			qrPhase = *o.Phase
		}
		qPhase := qrPhase
		if qPhase != "" {
			if err := r.SetQueryParam("phase", qPhase); err != nil {
				return err
			}
		}

	}

	if o.PowerPanelID != nil {

		// query param power_panel_id
		var qrPowerPanelID string
		if o.PowerPanelID != nil {
			qrPowerPanelID = *o.PowerPanelID
		}
		qPowerPanelID := qrPowerPanelID
		if qPowerPanelID != "" {
			if err := r.SetQueryParam("power_panel_id", qPowerPanelID); err != nil {
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

	if o.RackID != nil {

		// query param rack_id
		var qrRackID string
		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := qrRackID
		if qRackID != "" {
			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Supply != nil {

		// query param supply
		var qrSupply string
		if o.Supply != nil {
			qrSupply = *o.Supply
		}
		qSupply := qrSupply
		if qSupply != "" {
			if err := r.SetQueryParam("supply", qSupply); err != nil {
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

	if o.Voltage != nil {

		// query param voltage
		var qrVoltage string
		if o.Voltage != nil {
			qrVoltage = *o.Voltage
		}
		qVoltage := qrVoltage
		if qVoltage != "" {
			if err := r.SetQueryParam("voltage", qVoltage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
