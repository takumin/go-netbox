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

// IPAddressInterface Interface
// swagger:model IPAddressInterface
type IPAddressInterface struct {

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// virtual machine
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine,omitempty"`
}

// Validate validates this IP address interface
func (m *IPAddressInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddressInterface) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddressInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *IPAddressInterface) validateVirtualMachine(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachine) { // not required
		return nil
	}

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressInterface) UnmarshalBinary(b []byte) error {
	var res IPAddressInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
