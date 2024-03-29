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

// DcimCablesDeleteReader is a Reader for the DcimCablesDelete structure.
type DcimCablesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimCablesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimCablesDeleteNoContent creates a DcimCablesDeleteNoContent with default headers values
func NewDcimCablesDeleteNoContent() *DcimCablesDeleteNoContent {
	return &DcimCablesDeleteNoContent{}
}

/*DcimCablesDeleteNoContent handles this case with default header values.

DcimCablesDeleteNoContent dcim cables delete no content
*/
type DcimCablesDeleteNoContent struct {
}

func (o *DcimCablesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcimCablesDeleteNoContent ", 204)
}

func (o *DcimCablesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
