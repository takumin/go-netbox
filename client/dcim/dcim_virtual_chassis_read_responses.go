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

// DcimVirtualChassisReadReader is a Reader for the DcimVirtualChassisRead structure.
type DcimVirtualChassisReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimVirtualChassisReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimVirtualChassisReadOK creates a DcimVirtualChassisReadOK with default headers values
func NewDcimVirtualChassisReadOK() *DcimVirtualChassisReadOK {
	return &DcimVirtualChassisReadOK{}
}

/*DcimVirtualChassisReadOK handles this case with default header values.

DcimVirtualChassisReadOK dcim virtual chassis read o k
*/
type DcimVirtualChassisReadOK struct {
	Payload *models.VirtualChassis
}

func (o *DcimVirtualChassisReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisReadOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
