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

// VirtualizationClusterGroupsCreateReader is a Reader for the VirtualizationClusterGroupsCreate structure.
type VirtualizationClusterGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewVirtualizationClusterGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsCreateCreated creates a VirtualizationClusterGroupsCreateCreated with default headers values
func NewVirtualizationClusterGroupsCreateCreated() *VirtualizationClusterGroupsCreateCreated {
	return &VirtualizationClusterGroupsCreateCreated{}
}

/*VirtualizationClusterGroupsCreateCreated handles this case with default header values.

VirtualizationClusterGroupsCreateCreated virtualization cluster groups create created
*/
type VirtualizationClusterGroupsCreateCreated struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-groups/][%d] virtualizationClusterGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *VirtualizationClusterGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
