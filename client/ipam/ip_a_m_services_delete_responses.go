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

// IPAMServicesDeleteReader is a Reader for the IPAMServicesDelete structure.
type IPAMServicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMServicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewIPAMServicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMServicesDeleteNoContent creates a IPAMServicesDeleteNoContent with default headers values
func NewIPAMServicesDeleteNoContent() *IPAMServicesDeleteNoContent {
	return &IPAMServicesDeleteNoContent{}
}

/*IPAMServicesDeleteNoContent handles this case with default header values.

IPAMServicesDeleteNoContent ipam services delete no content
*/
type IPAMServicesDeleteNoContent struct {
}

func (o *IPAMServicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/services/{id}/][%d] ipamServicesDeleteNoContent ", 204)
}

func (o *IPAMServicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
