// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// CircuitsCircuitTerminationsCreateReader is a Reader for the CircuitsCircuitTerminationsCreate structure.
type CircuitsCircuitTerminationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCircuitsCircuitTerminationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsCreateCreated creates a CircuitsCircuitTerminationsCreateCreated with default headers values
func NewCircuitsCircuitTerminationsCreateCreated() *CircuitsCircuitTerminationsCreateCreated {
	return &CircuitsCircuitTerminationsCreateCreated{}
}

/*CircuitsCircuitTerminationsCreateCreated handles this case with default header values.

CircuitsCircuitTerminationsCreateCreated circuits circuit terminations create created
*/
type CircuitsCircuitTerminationsCreateCreated struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsCircuitTerminationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
