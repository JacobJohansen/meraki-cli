// Code generated by go-swagger; DO NOT EDIT.

package net_flow_settings

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

// NewUpdateNetworkNetflowSettingsParams creates a new UpdateNetworkNetflowSettingsParams object
// with the default values initialized.
func NewUpdateNetworkNetflowSettingsParams() *UpdateNetworkNetflowSettingsParams {
	var ()
	return &UpdateNetworkNetflowSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkNetflowSettingsParamsWithTimeout creates a new UpdateNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkNetflowSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkNetflowSettingsParams {
	var ()
	return &UpdateNetworkNetflowSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkNetflowSettingsParamsWithContext creates a new UpdateNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkNetflowSettingsParamsWithContext(ctx context.Context) *UpdateNetworkNetflowSettingsParams {
	var ()
	return &UpdateNetworkNetflowSettingsParams{

		Context: ctx,
	}
}

// NewUpdateNetworkNetflowSettingsParamsWithHTTPClient creates a new UpdateNetworkNetflowSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkNetflowSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkNetflowSettingsParams {
	var ()
	return &UpdateNetworkNetflowSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkNetflowSettingsParams contains all the parameters to send to the API endpoint
for the update network netflow settings operation typically these are written to a http.Request
*/
type UpdateNetworkNetflowSettingsParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkNetflowSettings*/
	UpdateNetworkNetflowSettings *models.UpdateNetworkNetflowSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkNetflowSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) WithContext(ctx context.Context) *UpdateNetworkNetflowSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkNetflowSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) WithNetworkID(networkID string) *UpdateNetworkNetflowSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkNetflowSettings adds the updateNetworkNetflowSettings to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) WithUpdateNetworkNetflowSettings(updateNetworkNetflowSettings *models.UpdateNetworkNetflowSettings) *UpdateNetworkNetflowSettingsParams {
	o.SetUpdateNetworkNetflowSettings(updateNetworkNetflowSettings)
	return o
}

// SetUpdateNetworkNetflowSettings adds the updateNetworkNetflowSettings to the update network netflow settings params
func (o *UpdateNetworkNetflowSettingsParams) SetUpdateNetworkNetflowSettings(updateNetworkNetflowSettings *models.UpdateNetworkNetflowSettings) {
	o.UpdateNetworkNetflowSettings = updateNetworkNetflowSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkNetflowSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkNetflowSettings != nil {
		if err := r.SetBodyParam(o.UpdateNetworkNetflowSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
