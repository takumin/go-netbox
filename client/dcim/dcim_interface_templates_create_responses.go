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

// DcimInterfaceTemplatesCreateReader is a Reader for the DcimInterfaceTemplatesCreate structure.
type DcimInterfaceTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimInterfaceTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesCreateCreated creates a DcimInterfaceTemplatesCreateCreated with default headers values
func NewDcimInterfaceTemplatesCreateCreated() *DcimInterfaceTemplatesCreateCreated {
	return &DcimInterfaceTemplatesCreateCreated{}
}

/*DcimInterfaceTemplatesCreateCreated handles this case with default header values.

DcimInterfaceTemplatesCreateCreated dcim interface templates create created
*/
type DcimInterfaceTemplatesCreateCreated struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
