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

// WritableVirtualMachineWithConfigContext writable virtual machine with config context
// swagger:model WritableVirtualMachineWithConfigContext
type WritableVirtualMachineWithConfigContext struct {

	// Cluster
	// Required: true
	Cluster *int64 `json:"cluster"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Disk (GB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Disk *int64 `json:"disk,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// Memory (MB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Memory *int64 `json:"memory,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Platform
	Platform *int64 `json:"platform,omitempty"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	PrimaryIp4 *int64 `json:"primary_ip4,omitempty"`

	// Primary IPv6
	PrimaryIp6 *int64 `json:"primary_ip6,omitempty"`

	// Role
	Role *int64 `json:"role,omitempty"`

	// Site
	// Read Only: true
	Site string `json:"site,omitempty"`

	// Status
	// Enum: [1 0 3]
	Status int64 `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// VCPUs
	// Maximum: 32767
	// Minimum: 0
	Vcpus *int64 `json:"vcpus,omitempty"`
}

// Validate validates this writable virtual machine with config context
func (m *WritableVirtualMachineWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if err := validate.MinimumInt("disk", "body", int64(*m.Disk), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("disk", "body", int64(*m.Disk), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory", "body", int64(*m.Memory), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory", "body", int64(*m.Memory), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateName(formats strfmt.Registry) error {

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

var writableVirtualMachineWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,0,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineWithConfigContextTypeStatusPropEnum = append(writableVirtualMachineWithConfigContextTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableVirtualMachineWithConfigContext) validateStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableVirtualMachineWithConfigContextTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineWithConfigContext) validateTags(formats strfmt.Registry) error {

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

func (m *WritableVirtualMachineWithConfigContext) validateVcpus(formats strfmt.Registry) error {

	if swag.IsZero(m.Vcpus) { // not required
		return nil
	}

	if err := validate.MinimumInt("vcpus", "body", int64(*m.Vcpus), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vcpus", "body", int64(*m.Vcpus), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVirtualMachineWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVirtualMachineWithConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableVirtualMachineWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
