// Code generated by go-swagger; DO NOT EDIT.

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

// IPAMIPAddressesCreateReader is a Reader for the IPAMIPAddressesCreate structure.
type IPAMIPAddressesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMIPAddressesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMIPAddressesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMIPAddressesCreateCreated creates a IPAMIPAddressesCreateCreated with default headers values
func NewIPAMIPAddressesCreateCreated() *IPAMIPAddressesCreateCreated {
	return &IPAMIPAddressesCreateCreated{}
}

/*IPAMIPAddressesCreateCreated handles this case with default header values.

IPAMIPAddressesCreateCreated ipam Ip addresses create created
*/
type IPAMIPAddressesCreateCreated struct {
	Payload *models.IPAddress
}

func (o *IPAMIPAddressesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-addresses/][%d] ipamIpAddressesCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMIPAddressesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
