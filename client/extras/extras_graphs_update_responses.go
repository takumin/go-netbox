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

// ExtrasGraphsUpdateReader is a Reader for the ExtrasGraphsUpdate structure.
type ExtrasGraphsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasGraphsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasGraphsUpdateOK creates a ExtrasGraphsUpdateOK with default headers values
func NewExtrasGraphsUpdateOK() *ExtrasGraphsUpdateOK {
	return &ExtrasGraphsUpdateOK{}
}

/*ExtrasGraphsUpdateOK handles this case with default header values.

ExtrasGraphsUpdateOK extras graphs update o k
*/
type ExtrasGraphsUpdateOK struct {
	Payload *models.Graph
}

func (o *ExtrasGraphsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/graphs/{id}/][%d] extrasGraphsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasGraphsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Graph)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}