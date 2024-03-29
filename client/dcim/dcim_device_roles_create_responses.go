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

// DcimDeviceRolesCreateReader is a Reader for the DcimDeviceRolesCreate structure.
type DcimDeviceRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimDeviceRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceRolesCreateCreated creates a DcimDeviceRolesCreateCreated with default headers values
func NewDcimDeviceRolesCreateCreated() *DcimDeviceRolesCreateCreated {
	return &DcimDeviceRolesCreateCreated{}
}

/*DcimDeviceRolesCreateCreated handles this case with default header values.

DcimDeviceRolesCreateCreated dcim device roles create created
*/
type DcimDeviceRolesCreateCreated struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-roles/][%d] dcimDeviceRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDeviceRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
