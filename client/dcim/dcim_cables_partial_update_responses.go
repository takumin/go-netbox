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

// DcimCablesPartialUpdateReader is a Reader for the DcimCablesPartialUpdate structure.
type DcimCablesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimCablesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimCablesPartialUpdateOK creates a DcimCablesPartialUpdateOK with default headers values
func NewDcimCablesPartialUpdateOK() *DcimCablesPartialUpdateOK {
	return &DcimCablesPartialUpdateOK{}
}

/*DcimCablesPartialUpdateOK handles this case with default header values.

DcimCablesPartialUpdateOK dcim cables partial update o k
*/
type DcimCablesPartialUpdateOK struct {
	Payload *models.Cable
}

func (o *DcimCablesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcimCablesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCablesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
