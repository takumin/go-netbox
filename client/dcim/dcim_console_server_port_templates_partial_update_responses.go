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

// DcimConsoleServerPortTemplatesPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesPartialUpdate structure.
type DcimConsoleServerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateOK creates a DcimConsoleServerPortTemplatesPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateOK() *DcimConsoleServerPortTemplatesPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesPartialUpdateOK{}
}

/*DcimConsoleServerPortTemplatesPartialUpdateOK handles this case with default header values.

DcimConsoleServerPortTemplatesPartialUpdateOK dcim console server port templates partial update o k
*/
type DcimConsoleServerPortTemplatesPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
