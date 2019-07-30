// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package extras

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

// NewExtrasTopologyMapsListParams creates a new ExtrasTopologyMapsListParams object
// with the default values initialized.
func NewExtrasTopologyMapsListParams() *ExtrasTopologyMapsListParams {
	var ()
	return &ExtrasTopologyMapsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsListParamsWithTimeout creates a new ExtrasTopologyMapsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsListParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsListParams {
	var ()
	return &ExtrasTopologyMapsListParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsListParamsWithContext creates a new ExtrasTopologyMapsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsListParamsWithContext(ctx context.Context) *ExtrasTopologyMapsListParams {
	var ()
	return &ExtrasTopologyMapsListParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsListParamsWithHTTPClient creates a new ExtrasTopologyMapsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsListParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsListParams {
	var ()
	return &ExtrasTopologyMapsListParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsListParams contains all the parameters to send to the API endpoint
for the extras topology maps list operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsListParams struct {

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
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithContext(ctx context.Context) *ExtrasTopologyMapsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithLimit(limit *int64) *ExtrasTopologyMapsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithName(name *string) *ExtrasTopologyMapsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithOffset(offset *int64) *ExtrasTopologyMapsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSite adds the site to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithSite(site *string) *ExtrasTopologyMapsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithSiteID(siteID *string) *ExtrasTopologyMapsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSlug adds the slug to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) WithSlug(slug *string) *ExtrasTopologyMapsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras topology maps list params
func (o *ExtrasTopologyMapsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
