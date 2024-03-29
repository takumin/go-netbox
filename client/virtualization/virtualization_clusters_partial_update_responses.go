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

// VirtualizationClustersPartialUpdateReader is a Reader for the VirtualizationClustersPartialUpdate structure.
type VirtualizationClustersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClustersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClustersPartialUpdateOK creates a VirtualizationClustersPartialUpdateOK with default headers values
func NewVirtualizationClustersPartialUpdateOK() *VirtualizationClustersPartialUpdateOK {
	return &VirtualizationClustersPartialUpdateOK{}
}

/*VirtualizationClustersPartialUpdateOK handles this case with default header values.

VirtualizationClustersPartialUpdateOK virtualization clusters partial update o k
*/
type VirtualizationClustersPartialUpdateOK struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
