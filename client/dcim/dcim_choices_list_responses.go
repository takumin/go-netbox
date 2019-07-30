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

// DcimChoicesListReader is a Reader for the DcimChoicesList structure.
type DcimChoicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimChoicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimChoicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimChoicesListOK creates a DcimChoicesListOK with default headers values
func NewDcimChoicesListOK() *DcimChoicesListOK {
	return &DcimChoicesListOK{}
}

/*DcimChoicesListOK handles this case with default header values.

DcimChoicesListOK dcim choices list o k
*/
type DcimChoicesListOK struct {
}

func (o *DcimChoicesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/_choices/][%d] dcimChoicesListOK ", 200)
}

func (o *DcimChoicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
