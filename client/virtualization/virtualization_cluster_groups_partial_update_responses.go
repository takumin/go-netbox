// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// VirtualizationClusterGroupsPartialUpdateReader is a Reader for the VirtualizationClusterGroupsPartialUpdate structure.
type VirtualizationClusterGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsPartialUpdateOK creates a VirtualizationClusterGroupsPartialUpdateOK with default headers values
func NewVirtualizationClusterGroupsPartialUpdateOK() *VirtualizationClusterGroupsPartialUpdateOK {
	return &VirtualizationClusterGroupsPartialUpdateOK{}
}

/*VirtualizationClusterGroupsPartialUpdateOK handles this case with default header values.

VirtualizationClusterGroupsPartialUpdateOK virtualization cluster groups partial update o k
*/
type VirtualizationClusterGroupsPartialUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
