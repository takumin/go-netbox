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

// DcimPowerOutletTemplatesReadReader is a Reader for the DcimPowerOutletTemplatesRead structure.
type DcimPowerOutletTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerOutletTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesReadOK creates a DcimPowerOutletTemplatesReadOK with default headers values
func NewDcimPowerOutletTemplatesReadOK() *DcimPowerOutletTemplatesReadOK {
	return &DcimPowerOutletTemplatesReadOK{}
}

/*DcimPowerOutletTemplatesReadOK handles this case with default header values.

DcimPowerOutletTemplatesReadOK dcim power outlet templates read o k
*/
type DcimPowerOutletTemplatesReadOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
