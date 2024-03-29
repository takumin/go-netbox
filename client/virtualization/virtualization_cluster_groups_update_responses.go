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

// VirtualizationClusterGroupsUpdateReader is a Reader for the VirtualizationClusterGroupsUpdate structure.
type VirtualizationClusterGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsUpdateOK creates a VirtualizationClusterGroupsUpdateOK with default headers values
func NewVirtualizationClusterGroupsUpdateOK() *VirtualizationClusterGroupsUpdateOK {
	return &VirtualizationClusterGroupsUpdateOK{}
}

/*VirtualizationClusterGroupsUpdateOK handles this case with default header values.

VirtualizationClusterGroupsUpdateOK virtualization cluster groups update o k
*/
type VirtualizationClusterGroupsUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
