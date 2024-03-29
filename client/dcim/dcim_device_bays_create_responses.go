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

// DcimDeviceBaysCreateReader is a Reader for the DcimDeviceBaysCreate structure.
type DcimDeviceBaysCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimDeviceBaysCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceBaysCreateCreated creates a DcimDeviceBaysCreateCreated with default headers values
func NewDcimDeviceBaysCreateCreated() *DcimDeviceBaysCreateCreated {
	return &DcimDeviceBaysCreateCreated{}
}

/*DcimDeviceBaysCreateCreated handles this case with default header values.

DcimDeviceBaysCreateCreated dcim device bays create created
*/
type DcimDeviceBaysCreateCreated struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-bays/][%d] dcimDeviceBaysCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDeviceBaysCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
