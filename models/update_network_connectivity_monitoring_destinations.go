// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateNetworkConnectivityMonitoringDestinations updateNetworkConnectivityMonitoringDestinations
//
// swagger:model updateNetworkConnectivityMonitoringDestinations
type UpdateNetworkConnectivityMonitoringDestinations struct {

	// The list of connectivity monitoring destinations
	Destinations []*Destination `json:"destinations"`
}

// Validate validates this update network connectivity monitoring destinations
func (m *UpdateNetworkConnectivityMonitoringDestinations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkConnectivityMonitoringDestinations) validateDestinations(formats strfmt.Registry) error {

	if swag.IsZero(m.Destinations) { // not required
		return nil
	}

	for i := 0; i < len(m.Destinations); i++ {
		if swag.IsZero(m.Destinations[i]) { // not required
			continue
		}

		if m.Destinations[i] != nil {
			if err := m.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkConnectivityMonitoringDestinations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkConnectivityMonitoringDestinations) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkConnectivityMonitoringDestinations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
