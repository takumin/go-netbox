// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIPAMVrfsListParams creates a new IPAMVrfsListParams object
// with the default values initialized.
func NewIPAMVrfsListParams() *IPAMVrfsListParams {
	var ()
	return &IPAMVrfsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVrfsListParamsWithTimeout creates a new IPAMVrfsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVrfsListParamsWithTimeout(timeout time.Duration) *IPAMVrfsListParams {
	var ()
	return &IPAMVrfsListParams{

		timeout: timeout,
	}
}

// NewIPAMVrfsListParamsWithContext creates a new IPAMVrfsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVrfsListParamsWithContext(ctx context.Context) *IPAMVrfsListParams {
	var ()
	return &IPAMVrfsListParams{

		Context: ctx,
	}
}

// NewIPAMVrfsListParamsWithHTTPClient creates a new IPAMVrfsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVrfsListParamsWithHTTPClient(client *http.Client) *IPAMVrfsListParams {
	var ()
	return &IPAMVrfsListParams{
		HTTPClient: client,
	}
}

/*IPAMVrfsListParams contains all the parameters to send to the API endpoint
for the ipam vrfs list operation typically these are written to a http.Request
*/
type IPAMVrfsListParams struct {

	/*EnforceUnique*/
	EnforceUnique *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
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
	/*Rd*/
	Rd *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *string
	/*TenantID*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTimeout(timeout time.Duration) *IPAMVrfsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithContext(ctx context.Context) *IPAMVrfsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithHTTPClient(client *http.Client) *IPAMVrfsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithEnforceUnique(enforceUnique *string) *IPAMVrfsListParams {
	o.SetEnforceUnique(enforceUnique)
	return o
}

// SetEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetEnforceUnique(enforceUnique *string) {
	o.EnforceUnique = enforceUnique
}

// WithIDIn adds the iDIn to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithIDIn(iDIn *string) *IPAMVrfsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithLimit(limit *int64) *IPAMVrfsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithName(name *string) *IPAMVrfsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithOffset(offset *int64) *IPAMVrfsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithQ(q *string) *IPAMVrfsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRd adds the rd to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithRd(rd *string) *IPAMVrfsListParams {
	o.SetRd(rd)
	return o
}

// SetRd adds the rd to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetRd(rd *string) {
	o.Rd = rd
}

// WithTag adds the tag to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTag(tag *string) *IPAMVrfsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTenant(tenant *string) *IPAMVrfsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTenantGroup(tenantGroup *string) *IPAMVrfsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTenantGroupID(tenantGroupID *string) *IPAMVrfsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the ipam vrfs list params
func (o *IPAMVrfsListParams) WithTenantID(tenantID *string) *IPAMVrfsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vrfs list params
func (o *IPAMVrfsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVrfsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnforceUnique != nil {

		// query param enforce_unique
		var qrEnforceUnique string
		if o.EnforceUnique != nil {
			qrEnforceUnique = *o.EnforceUnique
		}
		qEnforceUnique := qrEnforceUnique
		if qEnforceUnique != "" {
			if err := r.SetQueryParam("enforce_unique", qEnforceUnique); err != nil {
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

	if o.Rd != nil {

		// query param rd
		var qrRd string
		if o.Rd != nil {
			qrRd = *o.Rd
		}
		qRd := qrRd
		if qRd != "" {
			if err := r.SetQueryParam("rd", qRd); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string
		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {
			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}

	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string
		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {
			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
