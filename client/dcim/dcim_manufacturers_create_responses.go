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

// DcimManufacturersCreateReader is a Reader for the DcimManufacturersCreate structure.
type DcimManufacturersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimManufacturersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimManufacturersCreateCreated creates a DcimManufacturersCreateCreated with default headers values
func NewDcimManufacturersCreateCreated() *DcimManufacturersCreateCreated {
	return &DcimManufacturersCreateCreated{}
}

/*DcimManufacturersCreateCreated handles this case with default header values.

DcimManufacturersCreateCreated dcim manufacturers create created
*/
type DcimManufacturersCreateCreated struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/manufacturers/][%d] dcimManufacturersCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimManufacturersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
