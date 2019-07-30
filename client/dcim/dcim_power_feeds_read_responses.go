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

// DcimPowerFeedsReadReader is a Reader for the DcimPowerFeedsRead structure.
type DcimPowerFeedsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerFeedsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerFeedsReadOK creates a DcimPowerFeedsReadOK with default headers values
func NewDcimPowerFeedsReadOK() *DcimPowerFeedsReadOK {
	return &DcimPowerFeedsReadOK{}
}

/*DcimPowerFeedsReadOK handles this case with default header values.

DcimPowerFeedsReadOK dcim power feeds read o k
*/
type DcimPowerFeedsReadOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
