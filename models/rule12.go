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

// Rule12 Rule12
//
// swagger:model Rule12
type Rule12 struct {

	// Description of the rule (optional).
	Comment string `json:"comment,omitempty"`

	// Destination IP address (in IP or CIDR notation) or 'any'.
	// Required: true
	DstCidr *string `json:"dstCidr"`

	// Destination port. Must be in the range of 1-65535 or 'any'. Default is 'any'.
	DstPort string `json:"dstPort,omitempty"`

	// ip version
	IPVersion IPVersion `json:"ipVersion,omitempty"`

	// policy
	// Required: true
	Policy Policy7 `json:"policy"`

	// protocol
	// Required: true
	Protocol Protocol7 `json:"protocol"`

	// Source IP address (in IP or CIDR notation) or 'any'.
	// Required: true
	SrcCidr *string `json:"srcCidr"`

	// Source port. Must be in the range  of 1-65535 or 'any'. Default is 'any'.
	SrcPort string `json:"srcPort,omitempty"`

	// Incoming traffic VLAN. Must be in the range of 1-4095 or 'any'. Default is 'any'.
	Vlan string `json:"vlan,omitempty"`
}

// Validate validates this rule12
func (m *Rule12) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDstCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule12) validateDstCidr(formats strfmt.Registry) error {

	if err := validate.Required("dstCidr", "body", m.DstCidr); err != nil {
		return err
	}

	return nil
}

func (m *Rule12) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	if err := m.IPVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ipVersion")
		}
		return err
	}

	return nil
}

func (m *Rule12) validatePolicy(formats strfmt.Registry) error {

	if err := m.Policy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policy")
		}
		return err
	}

	return nil
}

func (m *Rule12) validateProtocol(formats strfmt.Registry) error {

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		}
		return err
	}

	return nil
}

func (m *Rule12) validateSrcCidr(formats strfmt.Registry) error {

	if err := validate.Required("srcCidr", "body", m.SrcCidr); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule12) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule12) UnmarshalBinary(b []byte) error {
	var res Rule12
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
