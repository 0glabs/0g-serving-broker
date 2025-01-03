// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/0glabs/0g-serving-broker/inference-router/zkclient/models"
)

// GenerateSolidityCalldataCombinedReader is a Reader for the GenerateSolidityCalldataCombined structure.
type GenerateSolidityCalldataCombinedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateSolidityCalldataCombinedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateSolidityCalldataCombinedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateSolidityCalldataCombinedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateSolidityCalldataCombinedOK creates a GenerateSolidityCalldataCombinedOK with default headers values
func NewGenerateSolidityCalldataCombinedOK() *GenerateSolidityCalldataCombinedOK {
	return &GenerateSolidityCalldataCombinedOK{}
}

/*
GenerateSolidityCalldataCombinedOK describes a response with status code 200, with default header values.

OK
*/
type GenerateSolidityCalldataCombinedOK struct {
	Payload *GenerateSolidityCalldataCombinedOKBody
}

// IsSuccess returns true when this generate solidity calldata combined o k response has a 2xx status code
func (o *GenerateSolidityCalldataCombinedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate solidity calldata combined o k response has a 3xx status code
func (o *GenerateSolidityCalldataCombinedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate solidity calldata combined o k response has a 4xx status code
func (o *GenerateSolidityCalldataCombinedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate solidity calldata combined o k response has a 5xx status code
func (o *GenerateSolidityCalldataCombinedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate solidity calldata combined o k response a status code equal to that given
func (o *GenerateSolidityCalldataCombinedOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateSolidityCalldataCombinedOK) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] generateSolidityCalldataCombinedOK  %+v", 200, o.Payload)
}

func (o *GenerateSolidityCalldataCombinedOK) String() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] generateSolidityCalldataCombinedOK  %+v", 200, o.Payload)
}

func (o *GenerateSolidityCalldataCombinedOK) GetPayload() *GenerateSolidityCalldataCombinedOKBody {
	return o.Payload
}

func (o *GenerateSolidityCalldataCombinedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GenerateSolidityCalldataCombinedOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateSolidityCalldataCombinedDefault creates a GenerateSolidityCalldataCombinedDefault with default headers values
func NewGenerateSolidityCalldataCombinedDefault(code int) *GenerateSolidityCalldataCombinedDefault {
	return &GenerateSolidityCalldataCombinedDefault{
		_statusCode: code,
	}
}

/*
GenerateSolidityCalldataCombinedDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateSolidityCalldataCombinedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the generate solidity calldata combined default response
func (o *GenerateSolidityCalldataCombinedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this generate solidity calldata combined default response has a 2xx status code
func (o *GenerateSolidityCalldataCombinedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate solidity calldata combined default response has a 3xx status code
func (o *GenerateSolidityCalldataCombinedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate solidity calldata combined default response has a 4xx status code
func (o *GenerateSolidityCalldataCombinedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate solidity calldata combined default response has a 5xx status code
func (o *GenerateSolidityCalldataCombinedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate solidity calldata combined default response a status code equal to that given
func (o *GenerateSolidityCalldataCombinedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GenerateSolidityCalldataCombinedDefault) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] generateSolidityCalldataCombined default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateSolidityCalldataCombinedDefault) String() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] generateSolidityCalldataCombined default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateSolidityCalldataCombinedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateSolidityCalldataCombinedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GenerateSolidityCalldataCombinedBody generate solidity calldata combined body
swagger:model GenerateSolidityCalldataCombinedBody
*/
type GenerateSolidityCalldataCombinedBody struct {

	// l
	L int64 `json:"l,omitempty"`

	// pubkey
	Pubkey models.PublicKey `json:"pubkey"`

	// requests
	Requests []*models.Request `json:"requests"`

	// signatures
	Signatures models.Signatures `json:"signatures"`
}

// Validate validates this generate solidity calldata combined body
func (o *GenerateSolidityCalldataCombinedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubkey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) validatePubkey(formats strfmt.Registry) error {
	if swag.IsZero(o.Pubkey) { // not required
		return nil
	}

	if err := o.Pubkey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "pubkey")
		}
		return err
	}

	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) validateRequests(formats strfmt.Registry) error {
	if swag.IsZero(o.Requests) { // not required
		return nil
	}

	for i := 0; i < len(o.Requests); i++ {
		if swag.IsZero(o.Requests[i]) { // not required
			continue
		}

		if o.Requests[i] != nil {
			if err := o.Requests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(o.Signatures) { // not required
		return nil
	}

	if err := o.Signatures.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signatures")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signatures")
		}
		return err
	}

	return nil
}

// ContextValidate validate this generate solidity calldata combined body based on the context it is used
func (o *GenerateSolidityCalldataCombinedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubkey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) contextValidatePubkey(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Pubkey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "pubkey")
		}
		return err
	}

	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Requests); i++ {

		if o.Requests[i] != nil {
			if err := o.Requests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GenerateSolidityCalldataCombinedBody) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Signatures.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signatures")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signatures")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GenerateSolidityCalldataCombinedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateSolidityCalldataCombinedBody) UnmarshalBinary(b []byte) error {
	var res GenerateSolidityCalldataCombinedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GenerateSolidityCalldataCombinedOKBody generate solidity calldata combined o k body
swagger:model GenerateSolidityCalldataCombinedOKBody
*/
type GenerateSolidityCalldataCombinedOKBody struct {

	// p a
	PA []string `json:"pA"`

	// p b
	PB [][]string `json:"pB"`

	// p c
	PC []string `json:"pC"`

	// pub inputs
	PubInputs []string `json:"pubInputs"`
}

// Validate validates this generate solidity calldata combined o k body
func (o *GenerateSolidityCalldataCombinedOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generate solidity calldata combined o k body based on context it is used
func (o *GenerateSolidityCalldataCombinedOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenerateSolidityCalldataCombinedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateSolidityCalldataCombinedOKBody) UnmarshalBinary(b []byte) error {
	var res GenerateSolidityCalldataCombinedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}