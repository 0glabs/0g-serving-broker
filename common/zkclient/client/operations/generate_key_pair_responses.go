// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/0glabs/0g-serving-agent/common/zkclient/models"
)

// GenerateKeyPairReader is a Reader for the GenerateKeyPair structure.
type GenerateKeyPairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateKeyPairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateKeyPairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateKeyPairDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateKeyPairOK creates a GenerateKeyPairOK with default headers values
func NewGenerateKeyPairOK() *GenerateKeyPairOK {
	return &GenerateKeyPairOK{}
}

/*
GenerateKeyPairOK describes a response with status code 200, with default header values.

OK
*/
type GenerateKeyPairOK struct {
	Payload *GenerateKeyPairOKBody
}

// IsSuccess returns true when this generate key pair o k response has a 2xx status code
func (o *GenerateKeyPairOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate key pair o k response has a 3xx status code
func (o *GenerateKeyPairOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate key pair o k response has a 4xx status code
func (o *GenerateKeyPairOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate key pair o k response has a 5xx status code
func (o *GenerateKeyPairOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate key pair o k response a status code equal to that given
func (o *GenerateKeyPairOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate key pair o k response
func (o *GenerateKeyPairOK) Code() int {
	return 200
}

func (o *GenerateKeyPairOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /sign-keypair][%d] generateKeyPairOK %s", 200, payload)
}

func (o *GenerateKeyPairOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /sign-keypair][%d] generateKeyPairOK %s", 200, payload)
}

func (o *GenerateKeyPairOK) GetPayload() *GenerateKeyPairOKBody {
	return o.Payload
}

func (o *GenerateKeyPairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GenerateKeyPairOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateKeyPairDefault creates a GenerateKeyPairDefault with default headers values
func NewGenerateKeyPairDefault(code int) *GenerateKeyPairDefault {
	return &GenerateKeyPairDefault{
		_statusCode: code,
	}
}

/*
GenerateKeyPairDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateKeyPairDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this generate key pair default response has a 2xx status code
func (o *GenerateKeyPairDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate key pair default response has a 3xx status code
func (o *GenerateKeyPairDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate key pair default response has a 4xx status code
func (o *GenerateKeyPairDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate key pair default response has a 5xx status code
func (o *GenerateKeyPairDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate key pair default response a status code equal to that given
func (o *GenerateKeyPairDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the generate key pair default response
func (o *GenerateKeyPairDefault) Code() int {
	return o._statusCode
}

func (o *GenerateKeyPairDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /sign-keypair][%d] generateKeyPair default %s", o._statusCode, payload)
}

func (o *GenerateKeyPairDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /sign-keypair][%d] generateKeyPair default %s", o._statusCode, payload)
}

func (o *GenerateKeyPairDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateKeyPairDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GenerateKeyPairOKBody generate key pair o k body
swagger:model GenerateKeyPairOKBody
*/
type GenerateKeyPairOKBody struct {

	// privkey
	Privkey []int64 `json:"privkey"`

	// pubkey
	Pubkey models.Pubkey `json:"pubkey"`
}

// Validate validates this generate key pair o k body
func (o *GenerateKeyPairOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubkey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateKeyPairOKBody) validatePubkey(formats strfmt.Registry) error {
	if swag.IsZero(o.Pubkey) { // not required
		return nil
	}

	if err := o.Pubkey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("generateKeyPairOK" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("generateKeyPairOK" + "." + "pubkey")
		}
		return err
	}

	return nil
}

// ContextValidate validate this generate key pair o k body based on the context it is used
func (o *GenerateKeyPairOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubkey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateKeyPairOKBody) contextValidatePubkey(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Pubkey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("generateKeyPairOK" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("generateKeyPairOK" + "." + "pubkey")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GenerateKeyPairOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateKeyPairOKBody) UnmarshalBinary(b []byte) error {
	var res GenerateKeyPairOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}