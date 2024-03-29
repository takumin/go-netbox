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

// DcimPlatformsDeleteReader is a Reader for the DcimPlatformsDelete structure.
type DcimPlatformsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimPlatformsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPlatformsDeleteNoContent creates a DcimPlatformsDeleteNoContent with default headers values
func NewDcimPlatformsDeleteNoContent() *DcimPlatformsDeleteNoContent {
	return &DcimPlatformsDeleteNoContent{}
}

/*DcimPlatformsDeleteNoContent handles this case with default header values.

DcimPlatformsDeleteNoContent dcim platforms delete no content
*/
type DcimPlatformsDeleteNoContent struct {
}

func (o *DcimPlatformsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcimPlatformsDeleteNoContent ", 204)
}

func (o *DcimPlatformsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
