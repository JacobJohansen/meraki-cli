// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmDeviceCommandLogsReader is a Reader for the GetNetworkSmDeviceCommandLogs structure.
type GetNetworkSmDeviceCommandLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceCommandLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceCommandLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSmDeviceCommandLogsOK creates a GetNetworkSmDeviceCommandLogsOK with default headers values
func NewGetNetworkSmDeviceCommandLogsOK() *GetNetworkSmDeviceCommandLogsOK {
	return &GetNetworkSmDeviceCommandLogsOK{}
}

/*GetNetworkSmDeviceCommandLogsOK handles this case with default header values.

Successful operation
*/
type GetNetworkSmDeviceCommandLogsOK struct {
	Payload interface{}
}

func (o *GetNetworkSmDeviceCommandLogsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/sm/{id}/deviceCommandLogs][%d] getNetworkSmDeviceCommandLogsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceCommandLogsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmDeviceCommandLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
