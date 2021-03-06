// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProvisionNetworkClientsReader is a Reader for the ProvisionNetworkClients structure.
type ProvisionNetworkClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisionNetworkClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProvisionNetworkClientsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProvisionNetworkClientsCreated creates a ProvisionNetworkClientsCreated with default headers values
func NewProvisionNetworkClientsCreated() *ProvisionNetworkClientsCreated {
	return &ProvisionNetworkClientsCreated{}
}

/*ProvisionNetworkClientsCreated handles this case with default header values.

Successful operation
*/
type ProvisionNetworkClientsCreated struct {
	Payload interface{}
}

func (o *ProvisionNetworkClientsCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/clients/provision][%d] provisionNetworkClientsCreated  %+v", 201, o.Payload)
}

func (o *ProvisionNetworkClientsCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ProvisionNetworkClientsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
