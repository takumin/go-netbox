// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/takumin/go-netbox/models"
)

// DcimInventoryItemsCreateReader is a Reader for the DcimInventoryItemsCreate structure.
type DcimInventoryItemsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimInventoryItemsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInventoryItemsCreateCreated creates a DcimInventoryItemsCreateCreated with default headers values
func NewDcimInventoryItemsCreateCreated() *DcimInventoryItemsCreateCreated {
	return &DcimInventoryItemsCreateCreated{}
}

/*DcimInventoryItemsCreateCreated handles this case with default header values.

DcimInventoryItemsCreateCreated dcim inventory items create created
*/
type DcimInventoryItemsCreateCreated struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInventoryItemsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
