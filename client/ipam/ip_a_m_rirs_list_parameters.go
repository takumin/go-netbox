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

// NewIPAMRirsListParams creates a new IPAMRirsListParams object
// with the default values initialized.
func NewIPAMRirsListParams() *IPAMRirsListParams {
	var ()
	return &IPAMRirsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMRirsListParamsWithTimeout creates a new IPAMRirsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMRirsListParamsWithTimeout(timeout time.Duration) *IPAMRirsListParams {
	var ()
	return &IPAMRirsListParams{

		timeout: timeout,
	}
}

// NewIPAMRirsListParamsWithContext creates a new IPAMRirsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMRirsListParamsWithContext(ctx context.Context) *IPAMRirsListParams {
	var ()
	return &IPAMRirsListParams{

		Context: ctx,
	}
}

// NewIPAMRirsListParamsWithHTTPClient creates a new IPAMRirsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMRirsListParamsWithHTTPClient(client *http.Client) *IPAMRirsListParams {
	var ()
	return &IPAMRirsListParams{
		HTTPClient: client,
	}
}

/*IPAMRirsListParams contains all the parameters to send to the API endpoint
for the ipam rirs list operation typically these are written to a http.Request
*/
type IPAMRirsListParams struct {

	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*IsPrivate*/
	IsPrivate *string
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

// WithTimeout adds the timeout to the ipam rirs list params
func (o *IPAMRirsListParams) WithTimeout(timeout time.Duration) *IPAMRirsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs list params
func (o *IPAMRirsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs list params
func (o *IPAMRirsListParams) WithContext(ctx context.Context) *IPAMRirsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs list params
func (o *IPAMRirsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs list params
func (o *IPAMRirsListParams) WithHTTPClient(client *http.Client) *IPAMRirsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs list params
func (o *IPAMRirsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDIn adds the iDIn to the ipam rirs list params
func (o *IPAMRirsListParams) WithIDIn(iDIn *string) *IPAMRirsListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam rirs list params
func (o *IPAMRirsListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithIsPrivate adds the isPrivate to the ipam rirs list params
func (o *IPAMRirsListParams) WithIsPrivate(isPrivate *string) *IPAMRirsListParams {
	o.SetIsPrivate(isPrivate)
	return o
}

// SetIsPrivate adds the isPrivate to the ipam rirs list params
func (o *IPAMRirsListParams) SetIsPrivate(isPrivate *string) {
	o.IsPrivate = isPrivate
}

// WithLimit adds the limit to the ipam rirs list params
func (o *IPAMRirsListParams) WithLimit(limit *int64) *IPAMRirsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam rirs list params
func (o *IPAMRirsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam rirs list params
func (o *IPAMRirsListParams) WithName(name *string) *IPAMRirsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam rirs list params
func (o *IPAMRirsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam rirs list params
func (o *IPAMRirsListParams) WithOffset(offset *int64) *IPAMRirsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam rirs list params
func (o *IPAMRirsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam rirs list params
func (o *IPAMRirsListParams) WithQ(q *string) *IPAMRirsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam rirs list params
func (o *IPAMRirsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the ipam rirs list params
func (o *IPAMRirsListParams) WithSlug(slug *string) *IPAMRirsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the ipam rirs list params
func (o *IPAMRirsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMRirsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.IsPrivate != nil {

		// query param is_private
		var qrIsPrivate string
		if o.IsPrivate != nil {
			qrIsPrivate = *o.IsPrivate
		}
		qIsPrivate := qrIsPrivate
		if qIsPrivate != "" {
			if err := r.SetQueryParam("is_private", qIsPrivate); err != nil {
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
