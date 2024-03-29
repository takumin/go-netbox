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

// DcimDevicesCreateReader is a Reader for the DcimDevicesCreate structure.
type DcimDevicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimDevicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDevicesCreateCreated creates a DcimDevicesCreateCreated with default headers values
func NewDcimDevicesCreateCreated() *DcimDevicesCreateCreated {
	return &DcimDevicesCreateCreated{}
}

/*DcimDevicesCreateCreated handles this case with default header values.

DcimDevicesCreateCreated dcim devices create created
*/
type DcimDevicesCreateCreated struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/devices/][%d] dcimDevicesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDevicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
