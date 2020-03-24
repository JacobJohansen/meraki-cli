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

// CreateNetworkSwitchSettingsQosRule createNetworkSwitchSettingsQosRule
//
// swagger:model createNetworkSwitchSettingsQosRule
type CreateNetworkSwitchSettingsQosRule struct {

	// DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp int32 `json:"dscp,omitempty"`

	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort int32 `json:"dstPort,omitempty"`

	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	DstPortRange string `json:"dstPortRange,omitempty"`

	// protocol
	Protocol Protocol8 `json:"protocol,omitempty"`

	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort int32 `json:"srcPort,omitempty"`

	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	SrcPortRange string `json:"srcPortRange,omitempty"`

	// The VLAN of the incoming packet. A null value will match any VLAN.
	// Required: true
	Vlan *int32 `json:"vlan"`
}

// Validate validates this create network switch settings qos rule
func (m *CreateNetworkSwitchSettingsQosRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateNetworkSwitchSettingsQosRule) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		}
		return err
	}

	return nil
}

func (m *CreateNetworkSwitchSettingsQosRule) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateNetworkSwitchSettingsQosRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateNetworkSwitchSettingsQosRule) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchSettingsQosRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}