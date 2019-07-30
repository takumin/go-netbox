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

// IPAMPrefixesAvailableIpsReadReader is a Reader for the IPAMPrefixesAvailableIpsRead structure.
type IPAMPrefixesAvailableIpsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesAvailableIpsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMPrefixesAvailableIpsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesAvailableIpsReadOK creates a IPAMPrefixesAvailableIpsReadOK with default headers values
func NewIPAMPrefixesAvailableIpsReadOK() *IPAMPrefixesAvailableIpsReadOK {
	return &IPAMPrefixesAvailableIpsReadOK{}
}

/*IPAMPrefixesAvailableIpsReadOK handles this case with default header values.

IPAMPrefixesAvailableIpsReadOK ipam prefixes available ips read o k
*/
type IPAMPrefixesAvailableIpsReadOK struct {
	Payload *models.Prefix
}

func (o *IPAMPrefixesAvailableIpsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsReadOK  %+v", 200, o.Payload)
}

func (o *IPAMPrefixesAvailableIpsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
