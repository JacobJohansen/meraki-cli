// Code generated by go-swagger; DO NOT EDIT.

package s_m

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

// NewCheckinNetworkSmDevicesParams creates a new CheckinNetworkSmDevicesParams object
// with the default values initialized.
func NewCheckinNetworkSmDevicesParams() *CheckinNetworkSmDevicesParams {
	var ()
	return &CheckinNetworkSmDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckinNetworkSmDevicesParamsWithTimeout creates a new CheckinNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckinNetworkSmDevicesParamsWithTimeout(timeout time.Duration) *CheckinNetworkSmDevicesParams {
	var ()
	return &CheckinNetworkSmDevicesParams{

		timeout: timeout,
	}
}

// NewCheckinNetworkSmDevicesParamsWithContext creates a new CheckinNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckinNetworkSmDevicesParamsWithContext(ctx context.Context) *CheckinNetworkSmDevicesParams {
	var ()
	return &CheckinNetworkSmDevicesParams{

		Context: ctx,
	}
}

// NewCheckinNetworkSmDevicesParamsWithHTTPClient creates a new CheckinNetworkSmDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckinNetworkSmDevicesParamsWithHTTPClient(client *http.Client) *CheckinNetworkSmDevicesParams {
	var ()
	return &CheckinNetworkSmDevicesParams{
		HTTPClient: client,
	}
}

/*CheckinNetworkSmDevicesParams contains all the parameters to send to the API endpoint
for the checkin network sm devices operation typically these are written to a http.Request
*/
type CheckinNetworkSmDevicesParams struct {

	/*CheckinNetworkSmDevices*/
	CheckinNetworkSmDevices *models.CheckinNetworkSmDevices
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) WithTimeout(timeout time.Duration) *CheckinNetworkSmDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) WithContext(ctx context.Context) *CheckinNetworkSmDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) WithHTTPClient(client *http.Client) *CheckinNetworkSmDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckinNetworkSmDevices adds the checkinNetworkSmDevices to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) WithCheckinNetworkSmDevices(checkinNetworkSmDevices *models.CheckinNetworkSmDevices) *CheckinNetworkSmDevicesParams {
	o.SetCheckinNetworkSmDevices(checkinNetworkSmDevices)
	return o
}

// SetCheckinNetworkSmDevices adds the checkinNetworkSmDevices to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) SetCheckinNetworkSmDevices(checkinNetworkSmDevices *models.CheckinNetworkSmDevices) {
	o.CheckinNetworkSmDevices = checkinNetworkSmDevices
}

// WithNetworkID adds the networkID to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) WithNetworkID(networkID string) *CheckinNetworkSmDevicesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the checkin network sm devices params
func (o *CheckinNetworkSmDevicesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckinNetworkSmDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CheckinNetworkSmDevices != nil {
		if err := r.SetBodyParam(o.CheckinNetworkSmDevices); err != nil {
			return err
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
