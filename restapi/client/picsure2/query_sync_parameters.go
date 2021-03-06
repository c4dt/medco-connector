// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewQuerySyncParams creates a new QuerySyncParams object
// with the default values initialized.
func NewQuerySyncParams() *QuerySyncParams {
	var ()
	return &QuerySyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuerySyncParamsWithTimeout creates a new QuerySyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuerySyncParamsWithTimeout(timeout time.Duration) *QuerySyncParams {
	var ()
	return &QuerySyncParams{

		timeout: timeout,
	}
}

// NewQuerySyncParamsWithContext creates a new QuerySyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuerySyncParamsWithContext(ctx context.Context) *QuerySyncParams {
	var ()
	return &QuerySyncParams{

		Context: ctx,
	}
}

// NewQuerySyncParamsWithHTTPClient creates a new QuerySyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuerySyncParamsWithHTTPClient(client *http.Client) *QuerySyncParams {
	var ()
	return &QuerySyncParams{
		HTTPClient: client,
	}
}

/*QuerySyncParams contains all the parameters to send to the API endpoint
for the query sync operation typically these are written to a http.Request
*/
type QuerySyncParams struct {

	/*Body
	  Query request.

	*/
	Body QuerySyncBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query sync params
func (o *QuerySyncParams) WithTimeout(timeout time.Duration) *QuerySyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query sync params
func (o *QuerySyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query sync params
func (o *QuerySyncParams) WithContext(ctx context.Context) *QuerySyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query sync params
func (o *QuerySyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query sync params
func (o *QuerySyncParams) WithHTTPClient(client *http.Client) *QuerySyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query sync params
func (o *QuerySyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query sync params
func (o *QuerySyncParams) WithBody(body QuerySyncBody) *QuerySyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query sync params
func (o *QuerySyncParams) SetBody(body QuerySyncBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *QuerySyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
