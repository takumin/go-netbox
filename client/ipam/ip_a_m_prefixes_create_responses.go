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

// IPAMPrefixesCreateReader is a Reader for the IPAMPrefixesCreate structure.
type IPAMPrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMPrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesCreateCreated creates a IPAMPrefixesCreateCreated with default headers values
func NewIPAMPrefixesCreateCreated() *IPAMPrefixesCreateCreated {
	return &IPAMPrefixesCreateCreated{}
}

/*IPAMPrefixesCreateCreated handles this case with default header values.

IPAMPrefixesCreateCreated ipam prefixes create created
*/
type IPAMPrefixesCreateCreated struct {
	Payload *models.Prefix
}

func (o *IPAMPrefixesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipamPrefixesCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMPrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
