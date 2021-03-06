// Code generated by go-swagger; DO NOT EDIT.

package m_g_subnet_pool_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new m g subnet pool settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for m g subnet pool settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkCellularGatewaySettingsSubnetPool(params *GetNetworkCellularGatewaySettingsSubnetPoolParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkCellularGatewaySettingsSubnetPoolOK, error)

	UpdateNetworkCellularGatewaySettingsSubnetPool(params *UpdateNetworkCellularGatewaySettingsSubnetPoolParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkCellularGatewaySettingsSubnetPoolOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkCellularGatewaySettingsSubnetPool gets network cellular gateway settings subnet pool

  Return the subnet pool and mask configured for MGs in the network.
*/
func (a *Client) GetNetworkCellularGatewaySettingsSubnetPool(params *GetNetworkCellularGatewaySettingsSubnetPoolParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkCellularGatewaySettingsSubnetPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkCellularGatewaySettingsSubnetPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkCellularGatewaySettingsSubnetPool",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/cellularGateway/settings/subnetPool",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkCellularGatewaySettingsSubnetPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkCellularGatewaySettingsSubnetPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkCellularGatewaySettingsSubnetPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkCellularGatewaySettingsSubnetPool updates network cellular gateway settings subnet pool

  Update the subnet pool and mask configuration for MGs in the network.
*/
func (a *Client) UpdateNetworkCellularGatewaySettingsSubnetPool(params *UpdateNetworkCellularGatewaySettingsSubnetPoolParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkCellularGatewaySettingsSubnetPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkCellularGatewaySettingsSubnetPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkCellularGatewaySettingsSubnetPool",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/cellularGateway/settings/subnetPool",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkCellularGatewaySettingsSubnetPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkCellularGatewaySettingsSubnetPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkCellularGatewaySettingsSubnetPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
