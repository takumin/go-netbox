// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TenancyTenantGroupsDeleteReader is a Reader for the TenancyTenantGroupsDelete structure.
type TenancyTenantGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTenancyTenantGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyTenantGroupsDeleteNoContent creates a TenancyTenantGroupsDeleteNoContent with default headers values
func NewTenancyTenantGroupsDeleteNoContent() *TenancyTenantGroupsDeleteNoContent {
	return &TenancyTenantGroupsDeleteNoContent{}
}

/*TenancyTenantGroupsDeleteNoContent handles this case with default header values.

TenancyTenantGroupsDeleteNoContent tenancy tenant groups delete no content
*/
type TenancyTenantGroupsDeleteNoContent struct {
}

func (o *TenancyTenantGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
