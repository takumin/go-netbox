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

// DcimFrontPortsPartialUpdateReader is a Reader for the DcimFrontPortsPartialUpdate structure.
type DcimFrontPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimFrontPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimFrontPortsPartialUpdateOK creates a DcimFrontPortsPartialUpdateOK with default headers values
func NewDcimFrontPortsPartialUpdateOK() *DcimFrontPortsPartialUpdateOK {
	return &DcimFrontPortsPartialUpdateOK{}
}

/*DcimFrontPortsPartialUpdateOK handles this case with default header values.

DcimFrontPortsPartialUpdateOK dcim front ports partial update o k
*/
type DcimFrontPortsPartialUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcimFrontPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
