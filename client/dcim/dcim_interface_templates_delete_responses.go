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

// DcimInterfaceTemplatesDeleteReader is a Reader for the DcimInterfaceTemplatesDelete structure.
type DcimInterfaceTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimInterfaceTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesDeleteNoContent creates a DcimInterfaceTemplatesDeleteNoContent with default headers values
func NewDcimInterfaceTemplatesDeleteNoContent() *DcimInterfaceTemplatesDeleteNoContent {
	return &DcimInterfaceTemplatesDeleteNoContent{}
}

/*DcimInterfaceTemplatesDeleteNoContent handles this case with default header values.

DcimInterfaceTemplatesDeleteNoContent dcim interface templates delete no content
*/
type DcimInterfaceTemplatesDeleteNoContent struct {
}

func (o *DcimInterfaceTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesDeleteNoContent ", 204)
}

func (o *DcimInterfaceTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
