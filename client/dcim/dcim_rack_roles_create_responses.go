// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// DcimRackRolesCreateReader is a Reader for the DcimRackRolesCreate structure.
type DcimRackRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimRackRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackRolesCreateCreated creates a DcimRackRolesCreateCreated with default headers values
func NewDcimRackRolesCreateCreated() *DcimRackRolesCreateCreated {
	return &DcimRackRolesCreateCreated{}
}

/*DcimRackRolesCreateCreated handles this case with default header values.

DcimRackRolesCreateCreated dcim rack roles create created
*/
type DcimRackRolesCreateCreated struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRackRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
