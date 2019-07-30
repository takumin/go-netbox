// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// ExtrasConfigContextsPartialUpdateReader is a Reader for the ExtrasConfigContextsPartialUpdate structure.
type ExtrasConfigContextsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasConfigContextsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasConfigContextsPartialUpdateOK creates a ExtrasConfigContextsPartialUpdateOK with default headers values
func NewExtrasConfigContextsPartialUpdateOK() *ExtrasConfigContextsPartialUpdateOK {
	return &ExtrasConfigContextsPartialUpdateOK{}
}

/*ExtrasConfigContextsPartialUpdateOK handles this case with default header values.

ExtrasConfigContextsPartialUpdateOK extras config contexts partial update o k
*/
type ExtrasConfigContextsPartialUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/{id}/][%d] extrasConfigContextsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigContextsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
