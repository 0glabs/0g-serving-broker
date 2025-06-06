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

	"github.com/0glabs/0g-serving-broker/inference/zkclient/models"
)

// SignatureReader is a Reader for the Signature structure.
type SignatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSignatureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSignatureOK creates a SignatureOK with default headers values
func NewSignatureOK() *SignatureOK {
	return &SignatureOK{}
}

/*
SignatureOK describes a response with status code 200, with default header values.

OK
*/
type SignatureOK struct {
	Payload *SignatureOKBody
}

// IsSuccess returns true when this signature o k response has a 2xx status code
func (o *SignatureOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this signature o k response has a 3xx status code
func (o *SignatureOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this signature o k response has a 4xx status code
func (o *SignatureOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this signature o k response has a 5xx status code
func (o *SignatureOK) IsServerError() bool {
	return false
}

// IsCode returns true when this signature o k response a status code equal to that given
func (o *SignatureOK) IsCode(code int) bool {
	return code == 200
}

func (o *SignatureOK) Error() string {
	return fmt.Sprintf("[POST /signature][%d] signatureOK  %+v", 200, o.Payload)
}

func (o *SignatureOK) String() string {
	return fmt.Sprintf("[POST /signature][%d] signatureOK  %+v", 200, o.Payload)
}

func (o *SignatureOK) GetPayload() *SignatureOKBody {
	return o.Payload
}

func (o *SignatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SignatureOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignatureDefault creates a SignatureDefault with default headers values
func NewSignatureDefault(code int) *SignatureDefault {
	return &SignatureDefault{
		_statusCode: code,
	}
}

/*
SignatureDefault describes a response with status code -1, with default header values.

Error
*/
type SignatureDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the signature default response
func (o *SignatureDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this signature default response has a 2xx status code
func (o *SignatureDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this signature default response has a 3xx status code
func (o *SignatureDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this signature default response has a 4xx status code
func (o *SignatureDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this signature default response has a 5xx status code
func (o *SignatureDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this signature default response a status code equal to that given
func (o *SignatureDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SignatureDefault) Error() string {
	return fmt.Sprintf("[POST /signature][%d] signature default  %+v", o._statusCode, o.Payload)
}

func (o *SignatureDefault) String() string {
	return fmt.Sprintf("[POST /signature][%d] signature default  %+v", o._statusCode, o.Payload)
}

func (o *SignatureDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SignatureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SignatureBody signature body
swagger:model SignatureBody
*/
type SignatureBody struct {

	// priv key
	PrivKey models.PrivateKey `json:"privKey"`

	// requests
	Requests []*models.RequestResponse `json:"requests"`

	// sign response
	SignResponse bool `json:"signResponse,omitempty"`
}

// Validate validates this signature body
func (o *SignatureBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrivKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SignatureBody) validatePrivKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PrivKey) { // not required
		return nil
	}

	if err := o.PrivKey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "privKey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "privKey")
		}
		return err
	}

	return nil
}

func (o *SignatureBody) validateRequests(formats strfmt.Registry) error {
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

// ContextValidate validate this signature body based on the context it is used
func (o *SignatureBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrivKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SignatureBody) contextValidatePrivKey(ctx context.Context, formats strfmt.Registry) error {

	if err := o.PrivKey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "privKey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "privKey")
		}
		return err
	}

	return nil
}

func (o *SignatureBody) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *SignatureBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SignatureBody) UnmarshalBinary(b []byte) error {
	var res SignatureBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SignatureOKBody signature o k body
swagger:model SignatureOKBody
*/
type SignatureOKBody struct {

	// signatures
	Signatures [][]int64 `json:"signatures"`
}

// Validate validates this signature o k body
func (o *SignatureOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signature o k body based on context it is used
func (o *SignatureOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SignatureOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SignatureOKBody) UnmarshalBinary(b []byte) error {
	var res SignatureOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
