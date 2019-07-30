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

// DcimFrontPortTemplatesUpdateReader is a Reader for the DcimFrontPortTemplatesUpdate structure.
type DcimFrontPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimFrontPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesUpdateOK creates a DcimFrontPortTemplatesUpdateOK with default headers values
func NewDcimFrontPortTemplatesUpdateOK() *DcimFrontPortTemplatesUpdateOK {
	return &DcimFrontPortTemplatesUpdateOK{}
}

/*DcimFrontPortTemplatesUpdateOK handles this case with default header values.

DcimFrontPortTemplatesUpdateOK dcim front port templates update o k
*/
type DcimFrontPortTemplatesUpdateOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
