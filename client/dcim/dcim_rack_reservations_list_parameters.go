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

// NewDcimRackReservationsListParams creates a new DcimRackReservationsListParams object
// with the default values initialized.
func NewDcimRackReservationsListParams() *DcimRackReservationsListParams {
	var ()
	return &DcimRackReservationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsListParamsWithTimeout creates a new DcimRackReservationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackReservationsListParamsWithTimeout(timeout time.Duration) *DcimRackReservationsListParams {
	var ()
	return &DcimRackReservationsListParams{

		timeout: timeout,
	}
}

// NewDcimRackReservationsListParamsWithContext creates a new DcimRackReservationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackReservationsListParamsWithContext(ctx context.Context) *DcimRackReservationsListParams {
	var ()
	return &DcimRackReservationsListParams{

		Context: ctx,
	}
}

// NewDcimRackReservationsListParamsWithHTTPClient creates a new DcimRackReservationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackReservationsListParamsWithHTTPClient(client *http.Client) *DcimRackReservationsListParams {
	var ()
	return &DcimRackReservationsListParams{
		HTTPClient: client,
	}
}

/*DcimRackReservationsListParams contains all the parameters to send to the API endpoint
for the dcim rack reservations list operation typically these are written to a http.Request
*/
type DcimRackReservationsListParams struct {

	/*Created*/
	Created *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*RackID*/
	RackID *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Tenant*/
	Tenant *string
	/*TenantGroup*/
	TenantGroup *string
	/*TenantGroupID*/
	TenantGroupID *string
	/*TenantID*/
	TenantID *string
	/*User*/
	User *string
	/*UserID*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTimeout(timeout time.Duration) *DcimRackReservationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithContext(ctx context.Context) *DcimRackReservationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithHTTPClient(client *http.Client) *DcimRackReservationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreated(created *string) *DcimRackReservationsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithGroup adds the group to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroup(group *string) *DcimRackReservationsListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroupID(groupID *string) *DcimRackReservationsListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithIDIn adds the iDIn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDIn(iDIn *string) *DcimRackReservationsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithLimit(limit *int64) *DcimRackReservationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithOffset(offset *int64) *DcimRackReservationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithQ(q *string) *DcimRackReservationsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackID adds the rackID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithRackID(rackID *string) *DcimRackReservationsListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithSite adds the site to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSite(site *string) *DcimRackReservationsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSiteID(siteID *string) *DcimRackReservationsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTenant adds the tenant to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenant(tenant *string) *DcimRackReservationsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantGroup adds the tenantGroup to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroup(tenantGroup *string) *DcimRackReservationsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupID adds the tenantGroupID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroupID(tenantGroupID *string) *DcimRackReservationsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantID adds the tenantID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantID(tenantID *string) *DcimRackReservationsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithUser adds the user to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUser(user *string) *DcimRackReservationsListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUser(user *string) {
	o.User = user
}

// WithUserID adds the userID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUserID(userID *string) *DcimRackReservationsListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
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

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
