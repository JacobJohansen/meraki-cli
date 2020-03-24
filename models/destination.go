// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Destination Destination
//
// swagger:model Destination
type Destination struct {

	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`

	// Description of the testing destination. Optional, defaults to null
	Description string `json:"description,omitempty"`

	// The IP address to test connectivity with
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this destination
func (m *Destination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Destination) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Destination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destination) UnmarshalBinary(b []byte) error {
	var res Destination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}