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

// IPAMVrfsCreateReader is a Reader for the IPAMVrfsCreate structure.
type IPAMVrfsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVrfsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMVrfsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVrfsCreateCreated creates a IPAMVrfsCreateCreated with default headers values
func NewIPAMVrfsCreateCreated() *IPAMVrfsCreateCreated {
	return &IPAMVrfsCreateCreated{}
}

/*IPAMVrfsCreateCreated handles this case with default header values.

IPAMVrfsCreateCreated ipam vrfs create created
*/
type IPAMVrfsCreateCreated struct {
	Payload *models.VRF
}

func (o *IPAMVrfsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMVrfsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}