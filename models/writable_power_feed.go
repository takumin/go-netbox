// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerFeed writable power feed
// swagger:model WritablePowerFeed
type WritablePowerFeed struct {

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Phase
	// Enum: [1 3]
	Phase int64 `json:"phase,omitempty"`

	// Power panel
	// Required: true
	PowerPanel *int64 `json:"power_panel"`

	// Rack
	Rack *int64 `json:"rack,omitempty"`

	// Status
	// Enum: [1 0 2 4]
	Status int64 `json:"status,omitempty"`

	// Supply
	// Enum: [1 2]
	Supply int64 `json:"supply,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [1 2]
	Type int64 `json:"type,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: 1
	Voltage int64 `json:"voltage,omitempty"`
}

// Validate validates this writable power feed
func (m *WritablePowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerFeed) validateAmperage(formats strfmt.Registry) error {

	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", int64(m.Amperage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", int64(m.Amperage), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateMaxUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", int64(m.MaxUtilization), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", int64(m.MaxUtilization), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateName(formats strfmt.Registry) error {

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

var writablePowerFeedTypePhasePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypePhasePropEnum = append(writablePowerFeedTypePhasePropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerFeed) validatePhaseEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypePhasePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,0,2,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeStatusPropEnum = append(writablePowerFeedTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerFeed) validateStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeSupplyPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeSupplyPropEnum = append(writablePowerFeedTypeSupplyPropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerFeed) validateSupplyEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeSupplyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateSupply(formats strfmt.Registry) error {

	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupplyEnum("supply", "body", m.Supply); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateTags(formats strfmt.Registry) error {

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

var writablePowerFeedTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeTypePropEnum = append(writablePowerFeedTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerFeed) validateTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateVoltage(formats strfmt.Registry) error {

	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", int64(m.Voltage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", int64(m.Voltage), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerFeed) UnmarshalBinary(b []byte) error {
	var res WritablePowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
