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

// DcimInterfacesDeleteReader is a Reader for the DcimInterfacesDelete structure.
type DcimInterfacesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimInterfacesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfacesDeleteNoContent creates a DcimInterfacesDeleteNoContent with default headers values
func NewDcimInterfacesDeleteNoContent() *DcimInterfacesDeleteNoContent {
	return &DcimInterfacesDeleteNoContent{}
}

/*DcimInterfacesDeleteNoContent handles this case with default header values.

DcimInterfacesDeleteNoContent dcim interfaces delete no content
*/
type DcimInterfacesDeleteNoContent struct {
}

func (o *DcimInterfacesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/{id}/][%d] dcimInterfacesDeleteNoContent ", 204)
}

func (o *DcimInterfacesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
