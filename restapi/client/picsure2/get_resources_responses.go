// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetResourcesReader is a Reader for the GetResources structure.
type GetResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResourcesOK creates a GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {
	return &GetResourcesOK{}
}

/*GetResourcesOK handles this case with default header values.

PIC-SURE 2 list of resources.
*/
type GetResourcesOK struct {
	Payload []*GetResourcesOKBodyItems0
}

func (o *GetResourcesOK) Error() string {
	return fmt.Sprintf("[GET /resource][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesDefault creates a GetResourcesDefault with default headers values
func NewGetResourcesDefault(code int) *GetResourcesDefault {
	return &GetResourcesDefault{
		_statusCode: code,
	}
}

/*GetResourcesDefault handles this case with default header values.

Error response
*/
type GetResourcesDefault struct {
	_statusCode int

	Payload *GetResourcesDefaultBody
}

// Code gets the status code for the get resources default response
func (o *GetResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetResourcesDefault) Error() string {
	return fmt.Sprintf("[GET /resource][%d] getResources default  %+v", o._statusCode, o.Payload)
}

func (o *GetResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourcesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetResourcesDefaultBody get resources default body
swagger:model GetResourcesDefaultBody
*/
type GetResourcesDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get resources default body
func (o *GetResourcesDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetResourcesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourcesOKBodyItems0 get resources o k body items0
swagger:model GetResourcesOKBodyItems0
*/
type GetResourcesOKBodyItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource r s path
	ResourceRSPath string `json:"resourceRSPath,omitempty"`

	// target URL
	TargetURL string `json:"targetURL,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this get resources o k body items0
func (o *GetResourcesOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourcesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourcesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetResourcesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
