// Code generated by go-swagger; DO NOT EDIT.

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

// DcimConsoleServerPortTemplatesCreateReader is a Reader for the DcimConsoleServerPortTemplatesCreate structure.
type DcimConsoleServerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimConsoleServerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesCreateCreated creates a DcimConsoleServerPortTemplatesCreateCreated with default headers values
func NewDcimConsoleServerPortTemplatesCreateCreated() *DcimConsoleServerPortTemplatesCreateCreated {
	return &DcimConsoleServerPortTemplatesCreateCreated{}
}

/*DcimConsoleServerPortTemplatesCreateCreated handles this case with default header values.

DcimConsoleServerPortTemplatesCreateCreated dcim console server port templates create created
*/
type DcimConsoleServerPortTemplatesCreateCreated struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
