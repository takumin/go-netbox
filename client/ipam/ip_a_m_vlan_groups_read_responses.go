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

// IPAMVlanGroupsReadReader is a Reader for the IPAMVlanGroupsRead structure.
type IPAMVlanGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVlanGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMVlanGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVlanGroupsReadOK creates a IPAMVlanGroupsReadOK with default headers values
func NewIPAMVlanGroupsReadOK() *IPAMVlanGroupsReadOK {
	return &IPAMVlanGroupsReadOK{}
}

/*IPAMVlanGroupsReadOK handles this case with default header values.

IPAMVlanGroupsReadOK ipam vlan groups read o k
*/
type IPAMVlanGroupsReadOK struct {
	Payload *models.VLANGroup
}

func (o *IPAMVlanGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *IPAMVlanGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
