// Code generated by go-swagger; DO NOT EDIT.

package admins

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
)

// NewGetOrganizationAdminsParams creates a new GetOrganizationAdminsParams object
// with the default values initialized.
func NewGetOrganizationAdminsParams() *GetOrganizationAdminsParams {
	var ()
	return &GetOrganizationAdminsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAdminsParamsWithTimeout creates a new GetOrganizationAdminsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationAdminsParamsWithTimeout(timeout time.Duration) *GetOrganizationAdminsParams {
	var ()
	return &GetOrganizationAdminsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationAdminsParamsWithContext creates a new GetOrganizationAdminsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationAdminsParamsWithContext(ctx context.Context) *GetOrganizationAdminsParams {
	var ()
	return &GetOrganizationAdminsParams{

		Context: ctx,
	}
}

// NewGetOrganizationAdminsParamsWithHTTPClient creates a new GetOrganizationAdminsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationAdminsParamsWithHTTPClient(client *http.Client) *GetOrganizationAdminsParams {
	var ()
	return &GetOrganizationAdminsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationAdminsParams contains all the parameters to send to the API endpoint
for the get organization admins operation typically these are written to a http.Request
*/
type GetOrganizationAdminsParams struct {

	/*OrganizationID*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization admins params
func (o *GetOrganizationAdminsParams) WithTimeout(timeout time.Duration) *GetOrganizationAdminsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization admins params
func (o *GetOrganizationAdminsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization admins params
func (o *GetOrganizationAdminsParams) WithContext(ctx context.Context) *GetOrganizationAdminsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization admins params
func (o *GetOrganizationAdminsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization admins params
func (o *GetOrganizationAdminsParams) WithHTTPClient(client *http.Client) *GetOrganizationAdminsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization admins params
func (o *GetOrganizationAdminsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization admins params
func (o *GetOrganizationAdminsParams) WithOrganizationID(organizationID string) *GetOrganizationAdminsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization admins params
func (o *GetOrganizationAdminsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAdminsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
