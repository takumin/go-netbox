// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimManufacturersDeleteReader is a Reader for the DcimManufacturersDelete structure.
type DcimManufacturersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimManufacturersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimManufacturersDeleteNoContent creates a DcimManufacturersDeleteNoContent with default headers values
func NewDcimManufacturersDeleteNoContent() *DcimManufacturersDeleteNoContent {
	return &DcimManufacturersDeleteNoContent{}
}

/*DcimManufacturersDeleteNoContent handles this case with default header values.

DcimManufacturersDeleteNoContent dcim manufacturers delete no content
*/
type DcimManufacturersDeleteNoContent struct {
}

func (o *DcimManufacturersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/manufacturers/{id}/][%d] dcimManufacturersDeleteNoContent ", 204)
}

func (o *DcimManufacturersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
