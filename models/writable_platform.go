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

// WritablePlatform writable platform
// swagger:model WritablePlatform
type WritablePlatform struct {

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Manufacturer
	//
	// Optionally limit this platform to devices of a certain manufacturer
	Manufacturer *int64 `json:"manufacturer,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// NAPALM arguments
	//
	// Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs *string `json:"napalm_args,omitempty"`

	// NAPALM driver
	//
	// The name of the NAPALM driver to use when interacting with devices
	// Max Length: 50
	NapalmDriver string `json:"napalm_driver,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Virtualmachine count
	// Read Only: true
	VirtualmachineCount int64 `json:"virtualmachine_count,omitempty"`
}

// Validate validates this writable platform
func (m *WritablePlatform) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNapalmDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePlatform) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritablePlatform) validateNapalmDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.NapalmDriver) { // not required
		return nil
	}

	if err := validate.MaxLength("napalm_driver", "body", string(m.NapalmDriver), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritablePlatform) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", string(*m.Slug), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePlatform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePlatform) UnmarshalBinary(b []byte) error {
	var res WritablePlatform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
