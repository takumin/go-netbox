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

// TenancyTenantGroupsPartialUpdateReader is a Reader for the TenancyTenantGroupsPartialUpdate structure.
type TenancyTenantGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTenancyTenantGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyTenantGroupsPartialUpdateOK creates a TenancyTenantGroupsPartialUpdateOK with default headers values
func NewTenancyTenantGroupsPartialUpdateOK() *TenancyTenantGroupsPartialUpdateOK {
	return &TenancyTenantGroupsPartialUpdateOK{}
}

/*TenancyTenantGroupsPartialUpdateOK handles this case with default header values.

TenancyTenantGroupsPartialUpdateOK tenancy tenant groups partial update o k
*/
type TenancyTenantGroupsPartialUpdateOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
