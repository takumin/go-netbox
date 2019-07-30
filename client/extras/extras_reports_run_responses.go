// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExtrasReportsRunReader is a Reader for the ExtrasReportsRun structure.
type ExtrasReportsRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasReportsRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewExtrasReportsRunCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasReportsRunCreated creates a ExtrasReportsRunCreated with default headers values
func NewExtrasReportsRunCreated() *ExtrasReportsRunCreated {
	return &ExtrasReportsRunCreated{}
}

/*ExtrasReportsRunCreated handles this case with default header values.

ExtrasReportsRunCreated extras reports run created
*/
type ExtrasReportsRunCreated struct {
}

func (o *ExtrasReportsRunCreated) Error() string {
	return fmt.Sprintf("[POST /extras/reports/{id}/run/][%d] extrasReportsRunCreated ", 201)
}

func (o *ExtrasReportsRunCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
