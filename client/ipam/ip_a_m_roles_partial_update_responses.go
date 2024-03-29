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

// IPAMRolesPartialUpdateReader is a Reader for the IPAMRolesPartialUpdate structure.
type IPAMRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMRolesPartialUpdateOK creates a IPAMRolesPartialUpdateOK with default headers values
func NewIPAMRolesPartialUpdateOK() *IPAMRolesPartialUpdateOK {
	return &IPAMRolesPartialUpdateOK{}
}

/*IPAMRolesPartialUpdateOK handles this case with default header values.

IPAMRolesPartialUpdateOK ipam roles partial update o k
*/
type IPAMRolesPartialUpdateOK struct {
	Payload *models.Role
}

func (o *IPAMRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipamRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
