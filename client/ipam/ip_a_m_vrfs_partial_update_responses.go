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

// IPAMVrfsPartialUpdateReader is a Reader for the IPAMVrfsPartialUpdate structure.
type IPAMVrfsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVrfsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMVrfsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVrfsPartialUpdateOK creates a IPAMVrfsPartialUpdateOK with default headers values
func NewIPAMVrfsPartialUpdateOK() *IPAMVrfsPartialUpdateOK {
	return &IPAMVrfsPartialUpdateOK{}
}

/*IPAMVrfsPartialUpdateOK handles this case with default header values.

IPAMVrfsPartialUpdateOK ipam vrfs partial update o k
*/
type IPAMVrfsPartialUpdateOK struct {
	Payload *models.VRF
}

func (o *IPAMVrfsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMVrfsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
