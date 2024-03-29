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

// IPAMAggregatesPartialUpdateReader is a Reader for the IPAMAggregatesPartialUpdate structure.
type IPAMAggregatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMAggregatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMAggregatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMAggregatesPartialUpdateOK creates a IPAMAggregatesPartialUpdateOK with default headers values
func NewIPAMAggregatesPartialUpdateOK() *IPAMAggregatesPartialUpdateOK {
	return &IPAMAggregatesPartialUpdateOK{}
}

/*IPAMAggregatesPartialUpdateOK handles this case with default header values.

IPAMAggregatesPartialUpdateOK ipam aggregates partial update o k
*/
type IPAMAggregatesPartialUpdateOK struct {
	Payload *models.Aggregate
}

func (o *IPAMAggregatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipamAggregatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMAggregatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
