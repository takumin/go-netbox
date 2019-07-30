// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedVirtualChassis Virtual chassis
// swagger:model NestedVirtualChassis
type NestedVirtualChassis struct {

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// master
	// Required: true
	Master *NestedDevice `json:"master"`

	// Member count
	// Read Only: true
	MemberCount int64 `json:"member_count,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested virtual chassis
func (m *NestedVirtualChassis) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualChassis) validateMaster(formats strfmt.Registry) error {

	if err := validate.Required("master", "body", m.Master); err != nil {
		return err
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *NestedVirtualChassis) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualChassis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualChassis) UnmarshalBinary(b []byte) error {
	var res NestedVirtualChassis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}