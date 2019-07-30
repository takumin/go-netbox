// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

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

// ExtrasImageAttachmentsCreateReader is a Reader for the ExtrasImageAttachmentsCreate structure.
type ExtrasImageAttachmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewExtrasImageAttachmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasImageAttachmentsCreateCreated creates a ExtrasImageAttachmentsCreateCreated with default headers values
func NewExtrasImageAttachmentsCreateCreated() *ExtrasImageAttachmentsCreateCreated {
	return &ExtrasImageAttachmentsCreateCreated{}
}

/*ExtrasImageAttachmentsCreateCreated handles this case with default header values.

ExtrasImageAttachmentsCreateCreated extras image attachments create created
*/
type ExtrasImageAttachmentsCreateCreated struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/image-attachments/][%d] extrasImageAttachmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasImageAttachmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
