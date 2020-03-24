// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AppliesTo AppliesTo
//
// Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
//
// swagger:model AppliesTo
type AppliesTo string

const (

	// AppliesToAllOrganizationAdmins captures enum value "All organization admins"
	AppliesToAllOrganizationAdmins AppliesTo = "All organization admins"

	// AppliesToAllEnterpriseAdmins captures enum value "All enterprise admins"
	AppliesToAllEnterpriseAdmins AppliesTo = "All enterprise admins"

	// AppliesToAllNetworkAdmins captures enum value "All network admins"
	AppliesToAllNetworkAdmins AppliesTo = "All network admins"

	// AppliesToAllAdminsOfNetworks captures enum value "All admins of networks..."
	AppliesToAllAdminsOfNetworks AppliesTo = "All admins of networks..."

	// AppliesToAllAdminsOfNetworksTagged captures enum value "All admins of networks tagged..."
	AppliesToAllAdminsOfNetworksTagged AppliesTo = "All admins of networks tagged..."

	// AppliesToSpecificAdmins captures enum value "Specific admins..."
	AppliesToSpecificAdmins AppliesTo = "Specific admins..."

	// AppliesToAllAdmins captures enum value "All admins"
	AppliesToAllAdmins AppliesTo = "All admins"

	// AppliesToAllSAMLAdmins captures enum value "All SAML admins"
	AppliesToAllSAMLAdmins AppliesTo = "All SAML admins"
)

// for schema
var appliesToEnum []interface{}

func init() {
	var res []AppliesTo
	if err := json.Unmarshal([]byte(`["All organization admins","All enterprise admins","All network admins","All admins of networks...","All admins of networks tagged...","Specific admins...","All admins","All SAML admins"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appliesToEnum = append(appliesToEnum, v)
	}
}

func (m AppliesTo) validateAppliesToEnum(path, location string, value AppliesTo) error {
	if err := validate.Enum(path, location, value, appliesToEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this applies to
func (m AppliesTo) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppliesToEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}