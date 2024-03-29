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

// DcimPowerPanelsPartialUpdateReader is a Reader for the DcimPowerPanelsPartialUpdate structure.
type DcimPowerPanelsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerPanelsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPanelsPartialUpdateOK creates a DcimPowerPanelsPartialUpdateOK with default headers values
func NewDcimPowerPanelsPartialUpdateOK() *DcimPowerPanelsPartialUpdateOK {
	return &DcimPowerPanelsPartialUpdateOK{}
}

/*DcimPowerPanelsPartialUpdateOK handles this case with default header values.

DcimPowerPanelsPartialUpdateOK dcim power panels partial update o k
*/
type DcimPowerPanelsPartialUpdateOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
