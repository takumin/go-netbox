// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceType device type
// swagger:model DeviceType
type DeviceType struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is full depth
	//
	// Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// manufacturer
	// Required: true
	Manufacturer *NestedManufacturer `json:"manufacturer"`

	// Model
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Model *string `json:"model"`

	// Part number
	//
	// Discrete part number (optional)
	// Max Length: 50
	PartNumber string `json:"part_number,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// subdevice role
	SubdeviceRole *DeviceTypeSubdeviceRole `json:"subdevice_role,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Height (U)
	// Maximum: 32767
	// Minimum: 0
	UHeight *int64 `json:"u_height,omitempty"`
}

// Validate validates this device type
func (m *DeviceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceType) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MinLength("model", "body", string(*m.Model), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", string(*m.Model), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validatePartNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.PartNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("part_number", "body", string(m.PartNumber), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateSlug(formats strfmt.Registry) error {

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

func (m *DeviceType) validateSubdeviceRole(formats strfmt.Registry) error {

	if swag.IsZero(m.SubdeviceRole) { // not required
		return nil
	}

	if m.SubdeviceRole != nil {
		if err := m.SubdeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subdevice_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *DeviceType) validateUHeight(formats strfmt.Registry) error {

	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", int64(*m.UHeight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", int64(*m.UHeight), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceType) UnmarshalBinary(b []byte) error {
	var res DeviceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceTypeSubdeviceRole Subdevice role
// swagger:model DeviceTypeSubdeviceRole
type DeviceTypeSubdeviceRole struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this device type subdevice role
func (m *DeviceTypeSubdeviceRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceTypeSubdeviceRole) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("subdevice_role"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *DeviceTypeSubdeviceRole) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("subdevice_role"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceTypeSubdeviceRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceTypeSubdeviceRole) UnmarshalBinary(b []byte) error {
	var res DeviceTypeSubdeviceRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
