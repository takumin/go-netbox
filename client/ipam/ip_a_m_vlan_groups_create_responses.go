// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// IPAMVlanGroupsCreateReader is a Reader for the IPAMVlanGroupsCreate structure.
type IPAMVlanGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVlanGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMVlanGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVlanGroupsCreateCreated creates a IPAMVlanGroupsCreateCreated with default headers values
func NewIPAMVlanGroupsCreateCreated() *IPAMVlanGroupsCreateCreated {
	return &IPAMVlanGroupsCreateCreated{}
}

/*IPAMVlanGroupsCreateCreated handles this case with default header values.

IPAMVlanGroupsCreateCreated ipam vlan groups create created
*/
type IPAMVlanGroupsCreateCreated struct {
	Payload *models.VLANGroup
}

func (o *IPAMVlanGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipamVlanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMVlanGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
