// Code generated by go-swagger; DO NOT EDIT.

package m_x_l3_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new m x l3 firewall API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for m x l3 firewall API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNetworkL3FirewallRules(params *GetNetworkL3FirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkL3FirewallRulesOK, error)

	UpdateNetworkL3FirewallRules(params *UpdateNetworkL3FirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkL3FirewallRulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNetworkL3FirewallRules gets network l3 firewall rules

  Return the L3 firewall rules for an MX network
*/
func (a *Client) GetNetworkL3FirewallRules(params *GetNetworkL3FirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkL3FirewallRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkL3FirewallRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkL3FirewallRules",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/l3FirewallRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkL3FirewallRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkL3FirewallRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkL3FirewallRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkL3FirewallRules updates network l3 firewall rules

  Update the L3 firewall rules of an MX network
*/
func (a *Client) UpdateNetworkL3FirewallRules(params *UpdateNetworkL3FirewallRulesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkL3FirewallRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkL3FirewallRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkL3FirewallRules",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/l3FirewallRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkL3FirewallRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkL3FirewallRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkL3FirewallRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
