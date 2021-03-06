// Code generated by go-swagger; DO NOT EDIT.

package wireless_health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkConnectionStatsReader is a Reader for the GetNetworkConnectionStats structure.
type GetNetworkConnectionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkConnectionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkConnectionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkConnectionStatsOK creates a GetNetworkConnectionStatsOK with default headers values
func NewGetNetworkConnectionStatsOK() *GetNetworkConnectionStatsOK {
	return &GetNetworkConnectionStatsOK{}
}

/*GetNetworkConnectionStatsOK handles this case with default header values.

Successful operation
*/
type GetNetworkConnectionStatsOK struct {
	Payload interface{}
}

func (o *GetNetworkConnectionStatsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/connectionStats][%d] getNetworkConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkConnectionStatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkConnectionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
