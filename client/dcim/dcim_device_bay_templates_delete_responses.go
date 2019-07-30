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

// DcimDeviceBayTemplatesDeleteReader is a Reader for the DcimDeviceBayTemplatesDelete structure.
type DcimDeviceBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimDeviceBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesDeleteNoContent creates a DcimDeviceBayTemplatesDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesDeleteNoContent() *DcimDeviceBayTemplatesDeleteNoContent {
	return &DcimDeviceBayTemplatesDeleteNoContent{}
}

/*DcimDeviceBayTemplatesDeleteNoContent handles this case with default header values.

DcimDeviceBayTemplatesDeleteNoContent dcim device bay templates delete no content
*/
type DcimDeviceBayTemplatesDeleteNoContent struct {
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
