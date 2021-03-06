// Code generated by go-swagger; DO NOT EDIT.

package v_l_a_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkVlansReader is a Reader for the GetNetworkVlans structure.
type GetNetworkVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkVlansOK creates a GetNetworkVlansOK with default headers values
func NewGetNetworkVlansOK() *GetNetworkVlansOK {
	return &GetNetworkVlansOK{}
}

/*GetNetworkVlansOK handles this case with default header values.

Successful operation
*/
type GetNetworkVlansOK struct {
	Payload interface{}
}

func (o *GetNetworkVlansOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/vlans][%d] getNetworkVlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkVlansOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
