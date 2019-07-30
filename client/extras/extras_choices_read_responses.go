// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExtrasChoicesReadReader is a Reader for the ExtrasChoicesRead structure.
type ExtrasChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasChoicesReadOK creates a ExtrasChoicesReadOK with default headers values
func NewExtrasChoicesReadOK() *ExtrasChoicesReadOK {
	return &ExtrasChoicesReadOK{}
}

/*ExtrasChoicesReadOK handles this case with default header values.

ExtrasChoicesReadOK extras choices read o k
*/
type ExtrasChoicesReadOK struct {
}

func (o *ExtrasChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/_choices/{id}/][%d] extrasChoicesReadOK ", 200)
}

func (o *ExtrasChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
