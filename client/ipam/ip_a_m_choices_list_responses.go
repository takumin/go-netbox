// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IPAMChoicesListReader is a Reader for the IPAMChoicesList structure.
type IPAMChoicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMChoicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMChoicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMChoicesListOK creates a IPAMChoicesListOK with default headers values
func NewIPAMChoicesListOK() *IPAMChoicesListOK {
	return &IPAMChoicesListOK{}
}

/*IPAMChoicesListOK handles this case with default header values.

IPAMChoicesListOK ipam choices list o k
*/
type IPAMChoicesListOK struct {
}

func (o *IPAMChoicesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/_choices/][%d] ipamChoicesListOK ", 200)
}

func (o *IPAMChoicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
