// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceBayTemplate writable device bay template
// swagger:model WritableDeviceBayTemplate
type WritableDeviceBayTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this writable device bay template
func (m *WritableDeviceBayTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceBayTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceBayTemplate) validateName(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *WritableDeviceBayTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceBayTemplate) UnmarshalBinary(b []byte) error {
	var res WritableDeviceBayTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
