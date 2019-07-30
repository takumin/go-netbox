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

// DcimFrontPortsUpdateReader is a Reader for the DcimFrontPortsUpdate structure.
type DcimFrontPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimFrontPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimFrontPortsUpdateOK creates a DcimFrontPortsUpdateOK with default headers values
func NewDcimFrontPortsUpdateOK() *DcimFrontPortsUpdateOK {
	return &DcimFrontPortsUpdateOK{}
}

/*DcimFrontPortsUpdateOK handles this case with default header values.

DcimFrontPortsUpdateOK dcim front ports update o k
*/
type DcimFrontPortsUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/{id}/][%d] dcimFrontPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
