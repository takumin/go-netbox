// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimChoicesReadReader is a Reader for the DcimChoicesRead structure.
type DcimChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimChoicesReadOK creates a DcimChoicesReadOK with default headers values
func NewDcimChoicesReadOK() *DcimChoicesReadOK {
	return &DcimChoicesReadOK{}
}

/*DcimChoicesReadOK handles this case with default header values.

DcimChoicesReadOK dcim choices read o k
*/
type DcimChoicesReadOK struct {
}

func (o *DcimChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/_choices/{id}/][%d] dcimChoicesReadOK ", 200)
}

func (o *DcimChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
