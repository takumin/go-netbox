// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// IPAMRirsReadReader is a Reader for the IPAMRirsRead structure.
type IPAMRirsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMRirsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMRirsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMRirsReadOK creates a IPAMRirsReadOK with default headers values
func NewIPAMRirsReadOK() *IPAMRirsReadOK {
	return &IPAMRirsReadOK{}
}

/*IPAMRirsReadOK handles this case with default header values.

IPAMRirsReadOK ipam rirs read o k
*/
type IPAMRirsReadOK struct {
	Payload *models.RIR
}

func (o *IPAMRirsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/rirs/{id}/][%d] ipamRirsReadOK  %+v", 200, o.Payload)
}

func (o *IPAMRirsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
