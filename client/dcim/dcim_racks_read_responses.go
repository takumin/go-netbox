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

// DcimRacksReadReader is a Reader for the DcimRacksRead structure.
type DcimRacksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRacksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksReadOK creates a DcimRacksReadOK with default headers values
func NewDcimRacksReadOK() *DcimRacksReadOK {
	return &DcimRacksReadOK{}
}

/*DcimRacksReadOK handles this case with default header values.

DcimRacksReadOK dcim racks read o k
*/
type DcimRacksReadOK struct {
	Payload *models.Rack
}

func (o *DcimRacksReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcimRacksReadOK  %+v", 200, o.Payload)
}

func (o *DcimRacksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
