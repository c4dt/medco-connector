// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Query query
// swagger:model query
type Query struct {

	// genomic annotations query (todo)
	GenomicAnnotations interface{} `json:"genomic-annotations,omitempty"`

	// i2b2 medco
	I2b2Medco *QueryI2b2Medco `json:"i2b2-medco,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this query
func (m *Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateI2b2Medco(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) validateI2b2Medco(formats strfmt.Registry) error {

	if swag.IsZero(m.I2b2Medco) { // not required
		return nil
	}

	if m.I2b2Medco != nil {
		if err := m.I2b2Medco.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("i2b2-medco")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Query) UnmarshalBinary(b []byte) error {
	var res Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryI2b2Medco i2b2-medco query
// swagger:model QueryI2b2Medco
type QueryI2b2Medco struct {

	// differential privacy
	DifferentialPrivacy *QueryI2b2MedcoDifferentialPrivacy `json:"differentialPrivacy,omitempty"`

	// i2b2 panels (linked by an AND)
	Panels []*QueryI2b2MedcoPanelsItems0 `json:"panels"`

	// query type
	QueryType QueryType `json:"queryType,omitempty"`

	// user public key
	UserPublicKey string `json:"userPublicKey,omitempty"`
}

// Validate validates this query i2b2 medco
func (m *QueryI2b2Medco) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDifferentialPrivacy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePanels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryI2b2Medco) validateDifferentialPrivacy(formats strfmt.Registry) error {

	if swag.IsZero(m.DifferentialPrivacy) { // not required
		return nil
	}

	if m.DifferentialPrivacy != nil {
		if err := m.DifferentialPrivacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("i2b2-medco" + "." + "differentialPrivacy")
			}
			return err
		}
	}

	return nil
}

func (m *QueryI2b2Medco) validatePanels(formats strfmt.Registry) error {

	if swag.IsZero(m.Panels) { // not required
		return nil
	}

	for i := 0; i < len(m.Panels); i++ {
		if swag.IsZero(m.Panels[i]) { // not required
			continue
		}

		if m.Panels[i] != nil {
			if err := m.Panels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("i2b2-medco" + "." + "panels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryI2b2Medco) validateQueryType(formats strfmt.Registry) error {

	if swag.IsZero(m.QueryType) { // not required
		return nil
	}

	if err := m.QueryType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("i2b2-medco" + "." + "queryType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryI2b2Medco) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryI2b2Medco) UnmarshalBinary(b []byte) error {
	var res QueryI2b2Medco
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryI2b2MedcoDifferentialPrivacy differential privacy query parameters (todo)
// swagger:model QueryI2b2MedcoDifferentialPrivacy
type QueryI2b2MedcoDifferentialPrivacy struct {

	// query budget
	QueryBudget float64 `json:"queryBudget,omitempty"`
}

// Validate validates this query i2b2 medco differential privacy
func (m *QueryI2b2MedcoDifferentialPrivacy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryI2b2MedcoDifferentialPrivacy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryI2b2MedcoDifferentialPrivacy) UnmarshalBinary(b []byte) error {
	var res QueryI2b2MedcoDifferentialPrivacy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryI2b2MedcoPanelsItems0 query i2b2 medco panels items0
// swagger:model QueryI2b2MedcoPanelsItems0
type QueryI2b2MedcoPanelsItems0 struct {

	// i2b2 items (linked by an OR)
	Items []*QueryI2b2MedcoPanelsItems0ItemsItems0 `json:"items"`

	// exclude the i2b2 panel
	// Required: true
	Not *bool `json:"not"`
}

// Validate validates this query i2b2 medco panels items0
func (m *QueryI2b2MedcoPanelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryI2b2MedcoPanelsItems0) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryI2b2MedcoPanelsItems0) validateNot(formats strfmt.Registry) error {

	if err := validate.Required("not", "body", m.Not); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryI2b2MedcoPanelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryI2b2MedcoPanelsItems0) UnmarshalBinary(b []byte) error {
	var res QueryI2b2MedcoPanelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryI2b2MedcoPanelsItems0ItemsItems0 query i2b2 medco panels items0 items items0
// swagger:model QueryI2b2MedcoPanelsItems0ItemsItems0
type QueryI2b2MedcoPanelsItems0ItemsItems0 struct {

	// encrypted
	// Required: true
	Encrypted *bool `json:"encrypted"`

	// operator
	// Enum: [exists equals]
	Operator string `json:"operator,omitempty"`

	// query term
	QueryTerm string `json:"queryTerm,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this query i2b2 medco panels items0 items items0
func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncrypted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) validateEncrypted(formats strfmt.Registry) error {

	if err := validate.Required("encrypted", "body", m.Encrypted); err != nil {
		return err
	}

	return nil
}

var queryI2b2MedcoPanelsItems0ItemsItems0TypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["exists","equals"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryI2b2MedcoPanelsItems0ItemsItems0TypeOperatorPropEnum = append(queryI2b2MedcoPanelsItems0ItemsItems0TypeOperatorPropEnum, v)
	}
}

const (

	// QueryI2b2MedcoPanelsItems0ItemsItems0OperatorExists captures enum value "exists"
	QueryI2b2MedcoPanelsItems0ItemsItems0OperatorExists string = "exists"

	// QueryI2b2MedcoPanelsItems0ItemsItems0OperatorEquals captures enum value "equals"
	QueryI2b2MedcoPanelsItems0ItemsItems0OperatorEquals string = "equals"
)

// prop value enum
func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) validateOperatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, queryI2b2MedcoPanelsItems0ItemsItems0TypeOperatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryI2b2MedcoPanelsItems0ItemsItems0) UnmarshalBinary(b []byte) error {
	var res QueryI2b2MedcoPanelsItems0ItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
