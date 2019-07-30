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

// DcimRacksCreateReader is a Reader for the DcimRacksCreate structure.
type DcimRacksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimRacksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksCreateCreated creates a DcimRacksCreateCreated with default headers values
func NewDcimRacksCreateCreated() *DcimRacksCreateCreated {
	return &DcimRacksCreateCreated{}
}

/*DcimRacksCreateCreated handles this case with default header values.

DcimRacksCreateCreated dcim racks create created
*/
type DcimRacksCreateCreated struct {
	Payload *models.Rack
}

func (o *DcimRacksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/racks/][%d] dcimRacksCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRacksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
