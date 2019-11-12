// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExploreQueryParams creates a new ExploreQueryParams object
// with the default values initialized.
func NewExploreQueryParams() ExploreQueryParams {

	var (
		// initialize parameters with default values

		syncDefault = bool(true)
	)

	return ExploreQueryParams{
		Sync: &syncDefault,
	}
}

// ExploreQueryParams contains all the bound params for the explore query operation
// typically these are obtained from a http.Request
//
// swagger:parameters exploreQuery
type ExploreQueryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*MedCo-Explore query request.
	  Required: true
	  In: body
	*/
	QueryRequest ExploreQueryBody
	/*Request synchronous query (defaults to true).
	  In: query
	  Default: true
	*/
	Sync *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewExploreQueryParams() beforehand.
func (o *ExploreQueryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ExploreQueryBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("queryRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("queryRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.QueryRequest = body
			}
		}
	} else {
		res = append(res, errors.Required("queryRequest", "body"))
	}
	qSync, qhkSync, _ := qs.GetOK("sync")
	if err := o.bindSync(qSync, qhkSync, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSync binds and validates parameter Sync from query.
func (o *ExploreQueryParams) bindSync(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewExploreQueryParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("sync", "query", "bool", raw)
	}
	o.Sync = &value

	return nil
}