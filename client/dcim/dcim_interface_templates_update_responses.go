// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// DcimInterfaceTemplatesUpdateReader is a Reader for the DcimInterfaceTemplatesUpdate structure.
type DcimInterfaceTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInterfaceTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesUpdateOK creates a DcimInterfaceTemplatesUpdateOK with default headers values
func NewDcimInterfaceTemplatesUpdateOK() *DcimInterfaceTemplatesUpdateOK {
	return &DcimInterfaceTemplatesUpdateOK{}
}

/*DcimInterfaceTemplatesUpdateOK handles this case with default header values.

DcimInterfaceTemplatesUpdateOK dcim interface templates update o k
*/
type DcimInterfaceTemplatesUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
