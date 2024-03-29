// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CircuitsCircuitTerminationsDeleteReader is a Reader for the CircuitsCircuitTerminationsDelete structure.
type CircuitsCircuitTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCircuitsCircuitTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsDeleteNoContent creates a CircuitsCircuitTerminationsDeleteNoContent with default headers values
func NewCircuitsCircuitTerminationsDeleteNoContent() *CircuitsCircuitTerminationsDeleteNoContent {
	return &CircuitsCircuitTerminationsDeleteNoContent{}
}

/*CircuitsCircuitTerminationsDeleteNoContent handles this case with default header values.

CircuitsCircuitTerminationsDeleteNoContent circuits circuit terminations delete no content
*/
type CircuitsCircuitTerminationsDeleteNoContent struct {
}

func (o *CircuitsCircuitTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
