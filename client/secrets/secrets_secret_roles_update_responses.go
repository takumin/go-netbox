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

// SecretsSecretRolesUpdateReader is a Reader for the SecretsSecretRolesUpdate structure.
type SecretsSecretRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSecretsSecretRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecretsSecretRolesUpdateOK creates a SecretsSecretRolesUpdateOK with default headers values
func NewSecretsSecretRolesUpdateOK() *SecretsSecretRolesUpdateOK {
	return &SecretsSecretRolesUpdateOK{}
}

/*SecretsSecretRolesUpdateOK handles this case with default header values.

SecretsSecretRolesUpdateOK secrets secret roles update o k
*/
type SecretsSecretRolesUpdateOK struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/secret-roles/{id}/][%d] secretsSecretRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
