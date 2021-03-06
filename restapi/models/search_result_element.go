// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchResultElement search result element
// swagger:model searchResultElement
type SearchResultElement struct {

	// code
	Code string `json:"code,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// leaf
	// Required: true
	Leaf *bool `json:"leaf"`

	// medco encryption
	MedcoEncryption *SearchResultElementMedcoEncryption `json:"medcoEncryption,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// type
	// Enum: [container concept concept_numeric concept_enum concept_text genomic_annotation]
	Type string `json:"type,omitempty"`
}

// Validate validates this search result element
func (m *SearchResultElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeaf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedcoEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResultElement) validateLeaf(formats strfmt.Registry) error {

	if err := validate.Required("leaf", "body", m.Leaf); err != nil {
		return err
	}

	return nil
}

func (m *SearchResultElement) validateMedcoEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.MedcoEncryption) { // not required
		return nil
	}

	if m.MedcoEncryption != nil {
		if err := m.MedcoEncryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("medcoEncryption")
			}
			return err
		}
	}

	return nil
}

var searchResultElementTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["container","concept","concept_numeric","concept_enum","concept_text","genomic_annotation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchResultElementTypeTypePropEnum = append(searchResultElementTypeTypePropEnum, v)
	}
}

const (

	// SearchResultElementTypeContainer captures enum value "container"
	SearchResultElementTypeContainer string = "container"

	// SearchResultElementTypeConcept captures enum value "concept"
	SearchResultElementTypeConcept string = "concept"

	// SearchResultElementTypeConceptNumeric captures enum value "concept_numeric"
	SearchResultElementTypeConceptNumeric string = "concept_numeric"

	// SearchResultElementTypeConceptEnum captures enum value "concept_enum"
	SearchResultElementTypeConceptEnum string = "concept_enum"

	// SearchResultElementTypeConceptText captures enum value "concept_text"
	SearchResultElementTypeConceptText string = "concept_text"

	// SearchResultElementTypeGenomicAnnotation captures enum value "genomic_annotation"
	SearchResultElementTypeGenomicAnnotation string = "genomic_annotation"
)

// prop value enum
func (m *SearchResultElement) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchResultElementTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchResultElement) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultElement) UnmarshalBinary(b []byte) error {
	var res SearchResultElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SearchResultElementMedcoEncryption search result element medco encryption
// swagger:model SearchResultElementMedcoEncryption
type SearchResultElementMedcoEncryption struct {

	// children ids
	ChildrenIds []int64 `json:"childrenIds"`

	// encrypted
	// Required: true
	Encrypted *bool `json:"encrypted"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this search result element medco encryption
func (m *SearchResultElementMedcoEncryption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncrypted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResultElementMedcoEncryption) validateEncrypted(formats strfmt.Registry) error {

	if err := validate.Required("medcoEncryption"+"."+"encrypted", "body", m.Encrypted); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultElementMedcoEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultElementMedcoEncryption) UnmarshalBinary(b []byte) error {
	var res SearchResultElementMedcoEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
