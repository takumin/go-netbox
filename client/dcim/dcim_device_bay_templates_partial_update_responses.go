// Code generated by go-swagger; DO NOT EDIT.

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

// DcimDeviceBayTemplatesPartialUpdateReader is a Reader for the DcimDeviceBayTemplatesPartialUpdate structure.
type DcimDeviceBayTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimDeviceBayTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateOK creates a DcimDeviceBayTemplatesPartialUpdateOK with default headers values
func NewDcimDeviceBayTemplatesPartialUpdateOK() *DcimDeviceBayTemplatesPartialUpdateOK {
	return &DcimDeviceBayTemplatesPartialUpdateOK{}
}

/*DcimDeviceBayTemplatesPartialUpdateOK handles this case with default header values.

DcimDeviceBayTemplatesPartialUpdateOK dcim device bay templates partial update o k
*/
type DcimDeviceBayTemplatesPartialUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
