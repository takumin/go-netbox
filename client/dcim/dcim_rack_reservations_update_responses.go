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

// DcimRackReservationsUpdateReader is a Reader for the DcimRackReservationsUpdate structure.
type DcimRackReservationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackReservationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackReservationsUpdateOK creates a DcimRackReservationsUpdateOK with default headers values
func NewDcimRackReservationsUpdateOK() *DcimRackReservationsUpdateOK {
	return &DcimRackReservationsUpdateOK{}
}

/*DcimRackReservationsUpdateOK handles this case with default header values.

DcimRackReservationsUpdateOK dcim rack reservations update o k
*/
type DcimRackReservationsUpdateOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-reservations/{id}/][%d] dcimRackReservationsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackReservationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
