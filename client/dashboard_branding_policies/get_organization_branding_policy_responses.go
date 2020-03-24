// Code generated by go-swagger; DO NOT EDIT.

package dashboard_branding_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationBrandingPolicyReader is a Reader for the GetOrganizationBrandingPolicy structure.
type GetOrganizationBrandingPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationBrandingPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationBrandingPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationBrandingPolicyOK creates a GetOrganizationBrandingPolicyOK with default headers values
func NewGetOrganizationBrandingPolicyOK() *GetOrganizationBrandingPolicyOK {
	return &GetOrganizationBrandingPolicyOK{}
}

/*GetOrganizationBrandingPolicyOK handles this case with default header values.

Successful operation
*/
type GetOrganizationBrandingPolicyOK struct {
	Payload interface{}
}

func (o *GetOrganizationBrandingPolicyOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/brandingPolicies/{brandingPolicyId}][%d] getOrganizationBrandingPolicyOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationBrandingPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationBrandingPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}