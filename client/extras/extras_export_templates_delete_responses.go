// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExtrasExportTemplatesDeleteReader is a Reader for the ExtrasExportTemplatesDelete structure.
type ExtrasExportTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewExtrasExportTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesDeleteNoContent creates a ExtrasExportTemplatesDeleteNoContent with default headers values
func NewExtrasExportTemplatesDeleteNoContent() *ExtrasExportTemplatesDeleteNoContent {
	return &ExtrasExportTemplatesDeleteNoContent{}
}

/*ExtrasExportTemplatesDeleteNoContent handles this case with default header values.

ExtrasExportTemplatesDeleteNoContent extras export templates delete no content
*/
type ExtrasExportTemplatesDeleteNoContent struct {
}

func (o *ExtrasExportTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
