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

// DcimFrontPortTemplatesCreateReader is a Reader for the DcimFrontPortTemplatesCreate structure.
type DcimFrontPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimFrontPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesCreateCreated creates a DcimFrontPortTemplatesCreateCreated with default headers values
func NewDcimFrontPortTemplatesCreateCreated() *DcimFrontPortTemplatesCreateCreated {
	return &DcimFrontPortTemplatesCreateCreated{}
}

/*DcimFrontPortTemplatesCreateCreated handles this case with default header values.

DcimFrontPortTemplatesCreateCreated dcim front port templates create created
*/
type DcimFrontPortTemplatesCreateCreated struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/front-port-templates/][%d] dcimFrontPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimFrontPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
