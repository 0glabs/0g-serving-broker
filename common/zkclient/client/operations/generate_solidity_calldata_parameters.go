// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/0glabs/0g-serving-agent/common/zkclient/models"
)

// NewGenerateSolidityCalldataParams creates a new GenerateSolidityCalldataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateSolidityCalldataParams() *GenerateSolidityCalldataParams {
	return &GenerateSolidityCalldataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateSolidityCalldataParamsWithTimeout creates a new GenerateSolidityCalldataParams object
// with the ability to set a timeout on a request.
func NewGenerateSolidityCalldataParamsWithTimeout(timeout time.Duration) *GenerateSolidityCalldataParams {
	return &GenerateSolidityCalldataParams{
		timeout: timeout,
	}
}

// NewGenerateSolidityCalldataParamsWithContext creates a new GenerateSolidityCalldataParams object
// with the ability to set a context for a request.
func NewGenerateSolidityCalldataParamsWithContext(ctx context.Context) *GenerateSolidityCalldataParams {
	return &GenerateSolidityCalldataParams{
		Context: ctx,
	}
}

// NewGenerateSolidityCalldataParamsWithHTTPClient creates a new GenerateSolidityCalldataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateSolidityCalldataParamsWithHTTPClient(client *http.Client) *GenerateSolidityCalldataParams {
	return &GenerateSolidityCalldataParams{
		HTTPClient: client,
	}
}

/*
GenerateSolidityCalldataParams contains all the parameters to send to the API endpoint

	for the generate solidity calldata operation.

	Typically these are written to a http.Request.
*/
type GenerateSolidityCalldataParams struct {

	// Body.
	Body models.AdditionalProperties

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate solidity calldata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateSolidityCalldataParams) WithDefaults() *GenerateSolidityCalldataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate solidity calldata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateSolidityCalldataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) WithTimeout(timeout time.Duration) *GenerateSolidityCalldataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) WithContext(ctx context.Context) *GenerateSolidityCalldataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) WithHTTPClient(client *http.Client) *GenerateSolidityCalldataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) WithBody(body models.AdditionalProperties) *GenerateSolidityCalldataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate solidity calldata params
func (o *GenerateSolidityCalldataParams) SetBody(body models.AdditionalProperties) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateSolidityCalldataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
