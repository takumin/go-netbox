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

// DcimPowerOutletTemplatesUpdateReader is a Reader for the DcimPowerOutletTemplatesUpdate structure.
type DcimPowerOutletTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerOutletTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesUpdateOK creates a DcimPowerOutletTemplatesUpdateOK with default headers values
func NewDcimPowerOutletTemplatesUpdateOK() *DcimPowerOutletTemplatesUpdateOK {
	return &DcimPowerOutletTemplatesUpdateOK{}
}

/*DcimPowerOutletTemplatesUpdateOK handles this case with default header values.

DcimPowerOutletTemplatesUpdateOK dcim power outlet templates update o k
*/
type DcimPowerOutletTemplatesUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
