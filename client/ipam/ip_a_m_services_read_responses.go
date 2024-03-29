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

// IPAMServicesReadReader is a Reader for the IPAMServicesRead structure.
type IPAMServicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMServicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMServicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMServicesReadOK creates a IPAMServicesReadOK with default headers values
func NewIPAMServicesReadOK() *IPAMServicesReadOK {
	return &IPAMServicesReadOK{}
}

/*IPAMServicesReadOK handles this case with default header values.

IPAMServicesReadOK ipam services read o k
*/
type IPAMServicesReadOK struct {
	Payload *models.Service
}

func (o *IPAMServicesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipamServicesReadOK  %+v", 200, o.Payload)
}

func (o *IPAMServicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
