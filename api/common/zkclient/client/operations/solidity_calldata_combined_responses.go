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

	"github.com/0glabs/0g-serving-agent/common/zkclient/models"
)

// SolidityCalldataCombinedReader is a Reader for the SolidityCalldataCombined structure.
type SolidityCalldataCombinedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SolidityCalldataCombinedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSolidityCalldataCombinedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSolidityCalldataCombinedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSolidityCalldataCombinedOK creates a SolidityCalldataCombinedOK with default headers values
func NewSolidityCalldataCombinedOK() *SolidityCalldataCombinedOK {
	return &SolidityCalldataCombinedOK{}
}

/*
SolidityCalldataCombinedOK describes a response with status code 200, with default header values.

OK
*/
type SolidityCalldataCombinedOK struct {
	Payload *SolidityCalldataCombinedOKBody
}

// IsSuccess returns true when this solidity calldata combined o k response has a 2xx status code
func (o *SolidityCalldataCombinedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this solidity calldata combined o k response has a 3xx status code
func (o *SolidityCalldataCombinedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this solidity calldata combined o k response has a 4xx status code
func (o *SolidityCalldataCombinedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this solidity calldata combined o k response has a 5xx status code
func (o *SolidityCalldataCombinedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this solidity calldata combined o k response a status code equal to that given
func (o *SolidityCalldataCombinedOK) IsCode(code int) bool {
	return code == 200
}

func (o *SolidityCalldataCombinedOK) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] solidityCalldataCombinedOK  %+v", 200, o.Payload)
}

func (o *SolidityCalldataCombinedOK) String() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] solidityCalldataCombinedOK  %+v", 200, o.Payload)
}

func (o *SolidityCalldataCombinedOK) GetPayload() *SolidityCalldataCombinedOKBody {
	return o.Payload
}

func (o *SolidityCalldataCombinedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SolidityCalldataCombinedOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSolidityCalldataCombinedDefault creates a SolidityCalldataCombinedDefault with default headers values
func NewSolidityCalldataCombinedDefault(code int) *SolidityCalldataCombinedDefault {
	return &SolidityCalldataCombinedDefault{
		_statusCode: code,
	}
}

/*
SolidityCalldataCombinedDefault describes a response with status code -1, with default header values.

Error
*/
type SolidityCalldataCombinedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the solidity calldata combined default response
func (o *SolidityCalldataCombinedDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this solidity calldata combined default response has a 2xx status code
func (o *SolidityCalldataCombinedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this solidity calldata combined default response has a 3xx status code
func (o *SolidityCalldataCombinedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this solidity calldata combined default response has a 4xx status code
func (o *SolidityCalldataCombinedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this solidity calldata combined default response has a 5xx status code
func (o *SolidityCalldataCombinedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this solidity calldata combined default response a status code equal to that given
func (o *SolidityCalldataCombinedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SolidityCalldataCombinedDefault) Error() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] solidityCalldataCombined default  %+v", o._statusCode, o.Payload)
}

func (o *SolidityCalldataCombinedDefault) String() string {
	return fmt.Sprintf("[POST /solidity-calldata-combined][%d] solidityCalldataCombined default  %+v", o._statusCode, o.Payload)
}

func (o *SolidityCalldataCombinedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SolidityCalldataCombinedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SolidityCalldataCombinedBody solidity calldata combined body
swagger:model SolidityCalldataCombinedBody
*/
type SolidityCalldataCombinedBody struct {

	// l
	L int64 `json:"l,omitempty"`

	// pubkey
	Pubkey models.PublicKey `json:"pubkey"`

	// requests
	Requests []*models.Request `json:"requests"`

	// signatures
	Signatures models.Signatures `json:"signatures"`
}

// Validate validates this solidity calldata combined body
func (o *SolidityCalldataCombinedBody) Validate(formats strfmt.Registry) error {
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

func (o *SolidityCalldataCombinedBody) validatePubkey(formats strfmt.Registry) error {
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

func (o *SolidityCalldataCombinedBody) validateRequests(formats strfmt.Registry) error {
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

func (o *SolidityCalldataCombinedBody) validateSignatures(formats strfmt.Registry) error {
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

// ContextValidate validate this solidity calldata combined body based on the context it is used
func (o *SolidityCalldataCombinedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *SolidityCalldataCombinedBody) contextValidatePubkey(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SolidityCalldataCombinedBody) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SolidityCalldataCombinedBody) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SolidityCalldataCombinedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SolidityCalldataCombinedBody) UnmarshalBinary(b []byte) error {
	var res SolidityCalldataCombinedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SolidityCalldataCombinedOKBody solidity calldata combined o k body
swagger:model SolidityCalldataCombinedOKBody
*/
type SolidityCalldataCombinedOKBody struct {

	// p a
	PA []string `json:"pA"`

	// p b
	PB [][]string `json:"pB"`

	// p c
	PC []string `json:"pC"`

	// pub inputs
	PubInputs []string `json:"pubInputs"`
}

// Validate validates this solidity calldata combined o k body
func (o *SolidityCalldataCombinedOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this solidity calldata combined o k body based on context it is used
func (o *SolidityCalldataCombinedOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SolidityCalldataCombinedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SolidityCalldataCombinedOKBody) UnmarshalBinary(b []byte) error {
	var res SolidityCalldataCombinedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
