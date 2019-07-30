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

// DcimRearPortTemplatesReadReader is a Reader for the DcimRearPortTemplatesRead structure.
type DcimRearPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRearPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRearPortTemplatesReadOK creates a DcimRearPortTemplatesReadOK with default headers values
func NewDcimRearPortTemplatesReadOK() *DcimRearPortTemplatesReadOK {
	return &DcimRearPortTemplatesReadOK{}
}

/*DcimRearPortTemplatesReadOK handles this case with default header values.

DcimRearPortTemplatesReadOK dcim rear port templates read o k
*/
type DcimRearPortTemplatesReadOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
