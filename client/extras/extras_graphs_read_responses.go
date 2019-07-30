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

// ExtrasGraphsReadReader is a Reader for the ExtrasGraphsRead structure.
type ExtrasGraphsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasGraphsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasGraphsReadOK creates a ExtrasGraphsReadOK with default headers values
func NewExtrasGraphsReadOK() *ExtrasGraphsReadOK {
	return &ExtrasGraphsReadOK{}
}

/*ExtrasGraphsReadOK handles this case with default header values.

ExtrasGraphsReadOK extras graphs read o k
*/
type ExtrasGraphsReadOK struct {
	Payload *models.Graph
}

func (o *ExtrasGraphsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/graphs/{id}/][%d] extrasGraphsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasGraphsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Graph)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
