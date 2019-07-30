// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableInterfaceTemplate writable interface template
// swagger:model WritableInterfaceTemplate
type WritableInterfaceTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Form factor
	// Read Only: true
	FormFactor string `json:"form_factor,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Management only
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Type
	// Enum: [0 200 800 1000 1120 1130 1150 1170 1050 1100 1200 1300 1310 1320 1350 1400 1420 1500 1510 1650 1520 1550 1600 1700 1750 2600 2610 2620 2630 2640 2810 2820 2830 6100 6200 6300 6400 6500 6600 6700 3010 3020 3040 3080 3160 3320 3400 4000 4010 4040 4050 5000 5050 5100 5150 5200 5300 5310 5320 5330 32767]
	Type int64 `json:"type,omitempty"`
}

// Validate validates this writable interface template
func (m *WritableInterfaceTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableInterfaceTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateName(formats strfmt.Registry) error {

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

var writableInterfaceTemplateTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,200,800,1000,1120,1130,1150,1170,1050,1100,1200,1300,1310,1320,1350,1400,1420,1500,1510,1650,1520,1550,1600,1700,1750,2600,2610,2620,2630,2640,2810,2820,2830,6100,6200,6300,6400,6500,6600,6700,3010,3020,3040,3080,3160,3320,3400,4000,4010,4040,4050,5000,5050,5100,5150,5200,5300,5310,5320,5330,32767]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTemplateTypeTypePropEnum = append(writableInterfaceTemplateTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableInterfaceTemplate) validateTypeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableInterfaceTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterfaceTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInterfaceTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInterfaceTemplate) UnmarshalBinary(b []byte) error {
	var res WritableInterfaceTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
