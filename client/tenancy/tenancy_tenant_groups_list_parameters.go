// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// NewTenancyTenantGroupsListParams creates a new TenancyTenantGroupsListParams object
// with the default values initialized.
func NewTenancyTenantGroupsListParams() *TenancyTenantGroupsListParams {
	var ()
	return &TenancyTenantGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsListParamsWithTimeout creates a new TenancyTenantGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenancyTenantGroupsListParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsListParams {
	var ()
	return &TenancyTenantGroupsListParams{

		timeout: timeout,
	}
}

// NewTenancyTenantGroupsListParamsWithContext creates a new TenancyTenantGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenancyTenantGroupsListParamsWithContext(ctx context.Context) *TenancyTenantGroupsListParams {
	var ()
	return &TenancyTenantGroupsListParams{

		Context: ctx,
	}
}

// NewTenancyTenantGroupsListParamsWithHTTPClient creates a new TenancyTenantGroupsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenancyTenantGroupsListParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsListParams {
	var ()
	return &TenancyTenantGroupsListParams{
		HTTPClient: client,
	}
}

/*TenancyTenantGroupsListParams contains all the parameters to send to the API endpoint
for the tenancy tenant groups list operation typically these are written to a http.Request
*/
type TenancyTenantGroupsListParams struct {

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
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithContext(ctx context.Context) *TenancyTenantGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithID(id *string) *TenancyTenantGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithLimit(limit *int64) *TenancyTenantGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithName(name *string) *TenancyTenantGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithOffset(offset *int64) *TenancyTenantGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithQ(q *string) *TenancyTenantGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) WithSlug(slug *string) *TenancyTenantGroupsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the tenancy tenant groups list params
func (o *TenancyTenantGroupsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
