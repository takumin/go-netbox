// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TenancyChoicesReadReader is a Reader for the TenancyChoicesRead structure.
type TenancyChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTenancyChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyChoicesReadOK creates a TenancyChoicesReadOK with default headers values
func NewTenancyChoicesReadOK() *TenancyChoicesReadOK {
	return &TenancyChoicesReadOK{}
}

/*TenancyChoicesReadOK handles this case with default header values.

TenancyChoicesReadOK tenancy choices read o k
*/
type TenancyChoicesReadOK struct {
}

func (o *TenancyChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/_choices/{id}/][%d] tenancyChoicesReadOK ", 200)
}

func (o *TenancyChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
