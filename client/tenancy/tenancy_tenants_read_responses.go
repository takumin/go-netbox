// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// TenancyTenantsReadReader is a Reader for the TenancyTenantsRead structure.
type TenancyTenantsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTenancyTenantsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyTenantsReadOK creates a TenancyTenantsReadOK with default headers values
func NewTenancyTenantsReadOK() *TenancyTenantsReadOK {
	return &TenancyTenantsReadOK{}
}

/*TenancyTenantsReadOK handles this case with default header values.

TenancyTenantsReadOK tenancy tenants read o k
*/
type TenancyTenantsReadOK struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/tenants/{id}/][%d] tenancyTenantsReadOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
