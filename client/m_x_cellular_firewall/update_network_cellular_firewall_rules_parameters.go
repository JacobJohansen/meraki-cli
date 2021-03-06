// Code generated by go-swagger; DO NOT EDIT.

package m_x_cellular_firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkCellularFirewallRulesParams creates a new UpdateNetworkCellularFirewallRulesParams object
// with the default values initialized.
func NewUpdateNetworkCellularFirewallRulesParams() *UpdateNetworkCellularFirewallRulesParams {
	var ()
	return &UpdateNetworkCellularFirewallRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkCellularFirewallRulesParamsWithTimeout creates a new UpdateNetworkCellularFirewallRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkCellularFirewallRulesParamsWithTimeout(timeout time.Duration) *UpdateNetworkCellularFirewallRulesParams {
	var ()
	return &UpdateNetworkCellularFirewallRulesParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkCellularFirewallRulesParamsWithContext creates a new UpdateNetworkCellularFirewallRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkCellularFirewallRulesParamsWithContext(ctx context.Context) *UpdateNetworkCellularFirewallRulesParams {
	var ()
	return &UpdateNetworkCellularFirewallRulesParams{

		Context: ctx,
	}
}

// NewUpdateNetworkCellularFirewallRulesParamsWithHTTPClient creates a new UpdateNetworkCellularFirewallRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkCellularFirewallRulesParamsWithHTTPClient(client *http.Client) *UpdateNetworkCellularFirewallRulesParams {
	var ()
	return &UpdateNetworkCellularFirewallRulesParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkCellularFirewallRulesParams contains all the parameters to send to the API endpoint
for the update network cellular firewall rules operation typically these are written to a http.Request
*/
type UpdateNetworkCellularFirewallRulesParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkCellularFirewallRules*/
	UpdateNetworkCellularFirewallRules *models.UpdateNetworkCellularFirewallRules

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) WithTimeout(timeout time.Duration) *UpdateNetworkCellularFirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) WithContext(ctx context.Context) *UpdateNetworkCellularFirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) WithHTTPClient(client *http.Client) *UpdateNetworkCellularFirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) WithNetworkID(networkID string) *UpdateNetworkCellularFirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkCellularFirewallRules adds the updateNetworkCellularFirewallRules to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) WithUpdateNetworkCellularFirewallRules(updateNetworkCellularFirewallRules *models.UpdateNetworkCellularFirewallRules) *UpdateNetworkCellularFirewallRulesParams {
	o.SetUpdateNetworkCellularFirewallRules(updateNetworkCellularFirewallRules)
	return o
}

// SetUpdateNetworkCellularFirewallRules adds the updateNetworkCellularFirewallRules to the update network cellular firewall rules params
func (o *UpdateNetworkCellularFirewallRulesParams) SetUpdateNetworkCellularFirewallRules(updateNetworkCellularFirewallRules *models.UpdateNetworkCellularFirewallRules) {
	o.UpdateNetworkCellularFirewallRules = updateNetworkCellularFirewallRules
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkCellularFirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkCellularFirewallRules != nil {
		if err := r.SetBodyParam(o.UpdateNetworkCellularFirewallRules); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
