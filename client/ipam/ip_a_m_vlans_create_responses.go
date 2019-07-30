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

// IPAMVlansCreateReader is a Reader for the IPAMVlansCreate structure.
type IPAMVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVlansCreateCreated creates a IPAMVlansCreateCreated with default headers values
func NewIPAMVlansCreateCreated() *IPAMVlansCreateCreated {
	return &IPAMVlansCreateCreated{}
}

/*IPAMVlansCreateCreated handles this case with default header values.

IPAMVlansCreateCreated ipam vlans create created
*/
type IPAMVlansCreateCreated struct {
	Payload *models.VLAN
}

func (o *IPAMVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
