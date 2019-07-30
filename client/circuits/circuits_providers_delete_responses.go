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

// CircuitsProvidersDeleteReader is a Reader for the CircuitsProvidersDelete structure.
type CircuitsProvidersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCircuitsProvidersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsProvidersDeleteNoContent creates a CircuitsProvidersDeleteNoContent with default headers values
func NewCircuitsProvidersDeleteNoContent() *CircuitsProvidersDeleteNoContent {
	return &CircuitsProvidersDeleteNoContent{}
}

/*CircuitsProvidersDeleteNoContent handles this case with default header values.

CircuitsProvidersDeleteNoContent circuits providers delete no content
*/
type CircuitsProvidersDeleteNoContent struct {
}

func (o *CircuitsProvidersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/{id}/][%d] circuitsProvidersDeleteNoContent ", 204)
}

func (o *CircuitsProvidersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
