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

// DcimConsoleServerPortsPartialUpdateReader is a Reader for the DcimConsoleServerPortsPartialUpdate structure.
type DcimConsoleServerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimConsoleServerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortsPartialUpdateOK creates a DcimConsoleServerPortsPartialUpdateOK with default headers values
func NewDcimConsoleServerPortsPartialUpdateOK() *DcimConsoleServerPortsPartialUpdateOK {
	return &DcimConsoleServerPortsPartialUpdateOK{}
}

/*DcimConsoleServerPortsPartialUpdateOK handles this case with default header values.

DcimConsoleServerPortsPartialUpdateOK dcim console server ports partial update o k
*/
type DcimConsoleServerPortsPartialUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
