// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/0glabs/0g-serving-broker/inference/zkclient/models"
)

// GenerateSolidityCalldataReader is a Reader for the GenerateSolidityCalldata structure.
type GenerateSolidityCalldataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateSolidityCalldataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateSolidityCalldataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateSolidityCalldataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateSolidityCalldataOK creates a GenerateSolidityCalldataOK with default headers values
func NewGenerateSolidityCalldataOK() *GenerateSolidityCalldataOK {
	return &GenerateSolidityCalldataOK{}
}

/*
GenerateSolidityCalldataOK describes a response with status code 200, with default header values.

OK
*/
type GenerateSolidityCalldataOK struct {
	Payload *GenerateSolidityCalldataOKBody
}

// IsSuccess returns true when this generate solidity calldata o k response has a 2xx status code
func (o *GenerateSolidityCalldataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate solidity calldata o k response has a 3xx status code
func (o *GenerateSolidityCalldataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate solidity calldata o k response has a 4xx status code
func (o *GenerateSolidityCalldataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate solidity calldata o k response has a 5xx status code
func (o *GenerateSolidityCalldataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate solidity calldata o k response a status code equal to that given
func (o *GenerateSolidityCalldataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateSolidityCalldataOK) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata][%d] generateSolidityCalldataOK  %+v", 200, o.Payload)
}

func (o *GenerateSolidityCalldataOK) String() string {
	return fmt.Sprintf("[POST /solidity-calldata][%d] generateSolidityCalldataOK  %+v", 200, o.Payload)
}

func (o *GenerateSolidityCalldataOK) GetPayload() *GenerateSolidityCalldataOKBody {
	return o.Payload
}

func (o *GenerateSolidityCalldataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GenerateSolidityCalldataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateSolidityCalldataDefault creates a GenerateSolidityCalldataDefault with default headers values
func NewGenerateSolidityCalldataDefault(code int) *GenerateSolidityCalldataDefault {
	return &GenerateSolidityCalldataDefault{
		_statusCode: code,
	}
}

/*
GenerateSolidityCalldataDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateSolidityCalldataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the generate solidity calldata default response
func (o *GenerateSolidityCalldataDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this generate solidity calldata default response has a 2xx status code
func (o *GenerateSolidityCalldataDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate solidity calldata default response has a 3xx status code
func (o *GenerateSolidityCalldataDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate solidity calldata default response has a 4xx status code
func (o *GenerateSolidityCalldataDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate solidity calldata default response has a 5xx status code
func (o *GenerateSolidityCalldataDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate solidity calldata default response a status code equal to that given
func (o *GenerateSolidityCalldataDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GenerateSolidityCalldataDefault) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata][%d] generateSolidityCalldata default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateSolidityCalldataDefault) String() string {
	return fmt.Sprintf("[POST /solidity-calldata][%d] generateSolidityCalldata default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateSolidityCalldataDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateSolidityCalldataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GenerateSolidityCalldataOKBody generate solidity calldata o k body
swagger:model GenerateSolidityCalldataOKBody
*/
type GenerateSolidityCalldataOKBody struct {

	// p a
	PA []string `json:"pA"`

	// p b
	PB [][]string `json:"pB"`

	// p c
	PC []string `json:"pC"`

	// pub inputs
	PubInputs []string `json:"pubInputs"`
}

// Validate validates this generate solidity calldata o k body
func (o *GenerateSolidityCalldataOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generate solidity calldata o k body based on context it is used
func (o *GenerateSolidityCalldataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenerateSolidityCalldataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateSolidityCalldataOKBody) UnmarshalBinary(b []byte) error {
	var res GenerateSolidityCalldataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}