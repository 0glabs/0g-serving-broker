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
)

// NewSolidityCalldataCombinedParams creates a new SolidityCalldataCombinedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSolidityCalldataCombinedParams() *SolidityCalldataCombinedParams {
	return &SolidityCalldataCombinedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSolidityCalldataCombinedParamsWithTimeout creates a new SolidityCalldataCombinedParams object
// with the ability to set a timeout on a request.
func NewSolidityCalldataCombinedParamsWithTimeout(timeout time.Duration) *SolidityCalldataCombinedParams {
	return &SolidityCalldataCombinedParams{
		timeout: timeout,
	}
}

// NewSolidityCalldataCombinedParamsWithContext creates a new SolidityCalldataCombinedParams object
// with the ability to set a context for a request.
func NewSolidityCalldataCombinedParamsWithContext(ctx context.Context) *SolidityCalldataCombinedParams {
	return &SolidityCalldataCombinedParams{
		Context: ctx,
	}
}

// NewSolidityCalldataCombinedParamsWithHTTPClient creates a new SolidityCalldataCombinedParams object
// with the ability to set a custom HTTPClient for a request.
func NewSolidityCalldataCombinedParamsWithHTTPClient(client *http.Client) *SolidityCalldataCombinedParams {
	return &SolidityCalldataCombinedParams{
		HTTPClient: client,
	}
}

/*
SolidityCalldataCombinedParams contains all the parameters to send to the API endpoint

	for the solidity calldata combined operation.

	Typically these are written to a http.Request.
*/
type SolidityCalldataCombinedParams struct {

	// Backend.
	Backend *string

	// Body.
	Body SolidityCalldataCombinedBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the solidity calldata combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SolidityCalldataCombinedParams) WithDefaults() *SolidityCalldataCombinedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the solidity calldata combined params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SolidityCalldataCombinedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) WithTimeout(timeout time.Duration) *SolidityCalldataCombinedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) WithContext(ctx context.Context) *SolidityCalldataCombinedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) WithHTTPClient(client *http.Client) *SolidityCalldataCombinedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) WithBackend(backend *string) *SolidityCalldataCombinedParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) SetBackend(backend *string) {
	o.Backend = backend
}

// WithBody adds the body to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) WithBody(body SolidityCalldataCombinedBody) *SolidityCalldataCombinedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the solidity calldata combined params
func (o *SolidityCalldataCombinedParams) SetBody(body SolidityCalldataCombinedBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SolidityCalldataCombinedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Backend != nil {

		// query param backend
		var qrBackend string

		if o.Backend != nil {
			qrBackend = *o.Backend
		}
		qBackend := qrBackend
		if qBackend != "" {

			if err := r.SetQueryParam("backend", qBackend); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
