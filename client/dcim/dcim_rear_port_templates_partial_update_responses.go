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

// DcimRearPortTemplatesPartialUpdateReader is a Reader for the DcimRearPortTemplatesPartialUpdate structure.
type DcimRearPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRearPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRearPortTemplatesPartialUpdateOK creates a DcimRearPortTemplatesPartialUpdateOK with default headers values
func NewDcimRearPortTemplatesPartialUpdateOK() *DcimRearPortTemplatesPartialUpdateOK {
	return &DcimRearPortTemplatesPartialUpdateOK{}
}

/*DcimRearPortTemplatesPartialUpdateOK handles this case with default header values.

DcimRearPortTemplatesPartialUpdateOK dcim rear port templates partial update o k
*/
type DcimRearPortTemplatesPartialUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
