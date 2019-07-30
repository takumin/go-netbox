// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

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

// NewIPAMPrefixesListParams creates a new IPAMPrefixesListParams object
// with the default values initialized.
func NewIPAMPrefixesListParams() *IPAMPrefixesListParams {
	var ()
	return &IPAMPrefixesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesListParamsWithTimeout creates a new IPAMPrefixesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesListParamsWithTimeout(timeout time.Duration) *IPAMPrefixesListParams {
	var ()
	return &IPAMPrefixesListParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesListParamsWithContext creates a new IPAMPrefixesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesListParamsWithContext(ctx context.Context) *IPAMPrefixesListParams {
	var ()
	return &IPAMPrefixesListParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesListParamsWithHTTPClient creates a new IPAMPrefixesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesListParamsWithHTTPClient(client *http.Client) *IPAMPrefixesListParams {
	var ()
	return &IPAMPrefixesListParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesListParams contains all the parameters to send to the API endpoint
for the ipam prefixes list operation typically these are written to a http.Request
*/
type IPAMPrefixesListParams struct {

	/*Contains*/
	Contains *string
	/*Family*/
	Family *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*IsPool*/
	IsPool *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MaskLength*/
	MaskLength *float64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Prefix*/
	Prefix *string
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Status*/
	Status *string
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
	/*VlanID*/
	VlanID *string
	/*VlanVid*/
	VlanVid *float64
	/*Vrf*/
	Vrf *string
	/*VrfID*/
	VrfID *string
	/*Within*/
	Within *string
	/*WithinInclude*/
	WithinInclude *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTimeout(timeout time.Duration) *IPAMPrefixesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithContext(ctx context.Context) *IPAMPrefixesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithHTTPClient(client *http.Client) *IPAMPrefixesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContains adds the contains to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithContains(contains *string) *IPAMPrefixesListParams {
	o.SetContains(contains)
	return o
}

// SetContains adds the contains to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetContains(contains *string) {
	o.Contains = contains
}

// WithFamily adds the family to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithFamily(family *string) *IPAMPrefixesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetFamily(family *string) {
	o.Family = family
}

// WithIDIn adds the iDIn to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithIDIn(iDIn *string) *IPAMPrefixesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithIsPool adds the isPool to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithIsPool(isPool *string) *IPAMPrefixesListParams {
	o.SetIsPool(isPool)
	return o
}

// SetIsPool adds the isPool to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetIsPool(isPool *string) {
	o.IsPool = isPool
}

// WithLimit adds the limit to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithLimit(limit *int64) *IPAMPrefixesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaskLength adds the maskLength to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithMaskLength(maskLength *float64) *IPAMPrefixesListParams {
	o.SetMaskLength(maskLength)
	return o
}

// SetMaskLength adds the maskLength to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetMaskLength(maskLength *float64) {
	o.MaskLength = maskLength
}

// WithOffset adds the offset to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithOffset(offset *int64) *IPAMPrefixesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPrefix adds the prefix to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithPrefix(prefix *string) *IPAMPrefixesListParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WithQ adds the q to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithQ(q *string) *IPAMPrefixesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithRole(role *string) *IPAMPrefixesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithRoleID(roleID *string) *IPAMPrefixesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSite adds the site to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithSite(site *string) *IPAMPrefixesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithSiteID(siteID *string) *IPAMPrefixesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithStatus(status *string) *IPAMPrefixesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTag(tag *string) *IPAMPrefixesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTenant(tenant *string) *IPAMPrefixesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTenantGroup(tenantGroup *string) *IPAMPrefixesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTenantGroupID(tenantGroupID *string) *IPAMPrefixesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithTenantID(tenantID *string) *IPAMPrefixesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithVlanID adds the vlanID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithVlanID(vlanID *string) *IPAMPrefixesListParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetVlanID(vlanID *string) {
	o.VlanID = vlanID
}

// WithVlanVid adds the vlanVid to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithVlanVid(vlanVid *float64) *IPAMPrefixesListParams {
	o.SetVlanVid(vlanVid)
	return o
}

// SetVlanVid adds the vlanVid to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetVlanVid(vlanVid *float64) {
	o.VlanVid = vlanVid
}

// WithVrf adds the vrf to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithVrf(vrf *string) *IPAMPrefixesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfID adds the vrfID to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithVrfID(vrfID *string) *IPAMPrefixesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetVrfID(vrfID *string) {
	o.VrfID = vrfID
}

// WithWithin adds the within to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithWithin(within *string) *IPAMPrefixesListParams {
	o.SetWithin(within)
	return o
}

// SetWithin adds the within to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetWithin(within *string) {
	o.Within = within
}

// WithWithinInclude adds the withinInclude to the ipam prefixes list params
func (o *IPAMPrefixesListParams) WithWithinInclude(withinInclude *string) *IPAMPrefixesListParams {
	o.SetWithinInclude(withinInclude)
	return o
}

// SetWithinInclude adds the withinInclude to the ipam prefixes list params
func (o *IPAMPrefixesListParams) SetWithinInclude(withinInclude *string) {
	o.WithinInclude = withinInclude
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contains != nil {

		// query param contains
		var qrContains string
		if o.Contains != nil {
			qrContains = *o.Contains
		}
		qContains := qrContains
		if qContains != "" {
			if err := r.SetQueryParam("contains", qContains); err != nil {
				return err
			}
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
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

	if o.IsPool != nil {

		// query param is_pool
		var qrIsPool string
		if o.IsPool != nil {
			qrIsPool = *o.IsPool
		}
		qIsPool := qrIsPool
		if qIsPool != "" {
			if err := r.SetQueryParam("is_pool", qIsPool); err != nil {
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

	if o.MaskLength != nil {

		// query param mask_length
		var qrMaskLength float64
		if o.MaskLength != nil {
			qrMaskLength = *o.MaskLength
		}
		qMaskLength := swag.FormatFloat64(qrMaskLength)
		if qMaskLength != "" {
			if err := r.SetQueryParam("mask_length", qMaskLength); err != nil {
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

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string
		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {
			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
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

	if o.VlanID != nil {

		// query param vlan_id
		var qrVlanID string
		if o.VlanID != nil {
			qrVlanID = *o.VlanID
		}
		qVlanID := qrVlanID
		if qVlanID != "" {
			if err := r.SetQueryParam("vlan_id", qVlanID); err != nil {
				return err
			}
		}

	}

	if o.VlanVid != nil {

		// query param vlan_vid
		var qrVlanVid float64
		if o.VlanVid != nil {
			qrVlanVid = *o.VlanVid
		}
		qVlanVid := swag.FormatFloat64(qrVlanVid)
		if qVlanVid != "" {
			if err := r.SetQueryParam("vlan_vid", qVlanVid); err != nil {
				return err
			}
		}

	}

	if o.Vrf != nil {

		// query param vrf
		var qrVrf string
		if o.Vrf != nil {
			qrVrf = *o.Vrf
		}
		qVrf := qrVrf
		if qVrf != "" {
			if err := r.SetQueryParam("vrf", qVrf); err != nil {
				return err
			}
		}

	}

	if o.VrfID != nil {

		// query param vrf_id
		var qrVrfID string
		if o.VrfID != nil {
			qrVrfID = *o.VrfID
		}
		qVrfID := qrVrfID
		if qVrfID != "" {
			if err := r.SetQueryParam("vrf_id", qVrfID); err != nil {
				return err
			}
		}

	}

	if o.Within != nil {

		// query param within
		var qrWithin string
		if o.Within != nil {
			qrWithin = *o.Within
		}
		qWithin := qrWithin
		if qWithin != "" {
			if err := r.SetQueryParam("within", qWithin); err != nil {
				return err
			}
		}

	}

	if o.WithinInclude != nil {

		// query param within_include
		var qrWithinInclude string
		if o.WithinInclude != nil {
			qrWithinInclude = *o.WithinInclude
		}
		qWithinInclude := qrWithinInclude
		if qWithinInclude != "" {
			if err := r.SetQueryParam("within_include", qWithinInclude); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
