// Code generated by go-swagger; DO NOT EDIT.

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

// VirtualizationClustersReadReader is a Reader for the VirtualizationClustersRead structure.
type VirtualizationClustersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClustersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClustersReadOK creates a VirtualizationClustersReadOK with default headers values
func NewVirtualizationClustersReadOK() *VirtualizationClustersReadOK {
	return &VirtualizationClustersReadOK{}
}

/*VirtualizationClustersReadOK handles this case with default header values.

VirtualizationClustersReadOK virtualization clusters read o k
*/
type VirtualizationClustersReadOK struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
