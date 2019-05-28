// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lca1/medco-connector/restapi/models"
)

// QueryResultOKCode is the HTTP code returned for type QueryResultOK
const QueryResultOKCode int = 200

/*QueryResultOK Query result.

swagger:response queryResultOK
*/
type QueryResultOK struct {

	/*
	  In: Body
	*/
	Payload *models.QueryResultElement `json:"body,omitempty"`
}

// NewQueryResultOK creates QueryResultOK with default headers values
func NewQueryResultOK() *QueryResultOK {

	return &QueryResultOK{}
}

// WithPayload adds the payload to the query result o k response
func (o *QueryResultOK) WithPayload(payload *models.QueryResultElement) *QueryResultOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query result o k response
func (o *QueryResultOK) SetPayload(payload *models.QueryResultElement) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryResultOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*QueryResultDefault Error response

swagger:response queryResultDefault
*/
type QueryResultDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *QueryResultDefaultBody `json:"body,omitempty"`
}

// NewQueryResultDefault creates QueryResultDefault with default headers values
func NewQueryResultDefault(code int) *QueryResultDefault {
	if code <= 0 {
		code = 500
	}

	return &QueryResultDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the query result default response
func (o *QueryResultDefault) WithStatusCode(code int) *QueryResultDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the query result default response
func (o *QueryResultDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the query result default response
func (o *QueryResultDefault) WithPayload(payload *QueryResultDefaultBody) *QueryResultDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query result default response
func (o *QueryResultDefault) SetPayload(payload *QueryResultDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryResultDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}