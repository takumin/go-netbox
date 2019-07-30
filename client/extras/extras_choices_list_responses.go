// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExtrasChoicesListReader is a Reader for the ExtrasChoicesList structure.
type ExtrasChoicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasChoicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasChoicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasChoicesListOK creates a ExtrasChoicesListOK with default headers values
func NewExtrasChoicesListOK() *ExtrasChoicesListOK {
	return &ExtrasChoicesListOK{}
}

/*ExtrasChoicesListOK handles this case with default header values.

ExtrasChoicesListOK extras choices list o k
*/
type ExtrasChoicesListOK struct {
}

func (o *ExtrasChoicesListOK) Error() string {
	return fmt.Sprintf("[GET /extras/_choices/][%d] extrasChoicesListOK ", 200)
}

func (o *ExtrasChoicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
