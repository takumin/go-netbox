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

// IPAMRirsDeleteReader is a Reader for the IPAMRirsDelete structure.
type IPAMRirsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMRirsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewIPAMRirsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMRirsDeleteNoContent creates a IPAMRirsDeleteNoContent with default headers values
func NewIPAMRirsDeleteNoContent() *IPAMRirsDeleteNoContent {
	return &IPAMRirsDeleteNoContent{}
}

/*IPAMRirsDeleteNoContent handles this case with default header values.

IPAMRirsDeleteNoContent ipam rirs delete no content
*/
type IPAMRirsDeleteNoContent struct {
}

func (o *IPAMRirsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipamRirsDeleteNoContent ", 204)
}

func (o *IPAMRirsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
