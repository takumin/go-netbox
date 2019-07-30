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

// VirtualizationInterfacesPartialUpdateReader is a Reader for the VirtualizationInterfacesPartialUpdate structure.
type VirtualizationInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationInterfacesPartialUpdateOK creates a VirtualizationInterfacesPartialUpdateOK with default headers values
func NewVirtualizationInterfacesPartialUpdateOK() *VirtualizationInterfacesPartialUpdateOK {
	return &VirtualizationInterfacesPartialUpdateOK{}
}

/*VirtualizationInterfacesPartialUpdateOK handles this case with default header values.

VirtualizationInterfacesPartialUpdateOK virtualization interfaces partial update o k
*/
type VirtualizationInterfacesPartialUpdateOK struct {
	Payload *models.VirtualMachineInterface
}

func (o *VirtualizationInterfacesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualizationInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
