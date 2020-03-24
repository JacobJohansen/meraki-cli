// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkDevice updateNetworkDevice
//
// swagger:model updateNetworkDevice
type UpdateNetworkDevice struct {

	// The address of a device
	Address string `json:"address,omitempty"`

	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanID string `json:"floorPlanId,omitempty"`

	// The latitude of a device
	Lat float64 `json:"lat,omitempty"`

	// The longitude of a device
	Lng float64 `json:"lng,omitempty"`

	// Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	MoveMapMarker bool `json:"moveMapMarker,omitempty"`

	// The name of a device
	Name string `json:"name,omitempty"`

	// The notes for the device. String. Limited to 255 characters.
	Notes string `json:"notes,omitempty"`

	// The ID of a switch profile to bind to the device (for available switch profiles, see the 'Switch Profiles' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	SwitchProfileID string `json:"switchProfileId,omitempty"`

	// The tags of a device
	Tags string `json:"tags,omitempty"`
}

// Validate validates this update network device
func (m *UpdateNetworkDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkDevice) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}