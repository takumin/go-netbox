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

// IPAMRolesReadReader is a Reader for the IPAMRolesRead structure.
type IPAMRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMRolesReadOK creates a IPAMRolesReadOK with default headers values
func NewIPAMRolesReadOK() *IPAMRolesReadOK {
	return &IPAMRolesReadOK{}
}

/*IPAMRolesReadOK handles this case with default header values.

IPAMRolesReadOK ipam roles read o k
*/
type IPAMRolesReadOK struct {
	Payload *models.Role
}

func (o *IPAMRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK  %+v", 200, o.Payload)
}

func (o *IPAMRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
