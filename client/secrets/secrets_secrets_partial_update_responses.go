// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// SecretsSecretsPartialUpdateReader is a Reader for the SecretsSecretsPartialUpdate structure.
type SecretsSecretsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSecretsSecretsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecretsSecretsPartialUpdateOK creates a SecretsSecretsPartialUpdateOK with default headers values
func NewSecretsSecretsPartialUpdateOK() *SecretsSecretsPartialUpdateOK {
	return &SecretsSecretsPartialUpdateOK{}
}

/*SecretsSecretsPartialUpdateOK handles this case with default header values.

SecretsSecretsPartialUpdateOK secrets secrets partial update o k
*/
type SecretsSecretsPartialUpdateOK struct {
	Payload *models.Secret
}

func (o *SecretsSecretsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/secrets/{id}/][%d] secretsSecretsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
