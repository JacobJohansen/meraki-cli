// Code generated by go-swagger; DO NOT EDIT.

package traffic_analysis_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkTrafficAnalysisSettingsReader is a Reader for the GetNetworkTrafficAnalysisSettings structure.
type GetNetworkTrafficAnalysisSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkTrafficAnalysisSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkTrafficAnalysisSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkTrafficAnalysisSettingsOK creates a GetNetworkTrafficAnalysisSettingsOK with default headers values
func NewGetNetworkTrafficAnalysisSettingsOK() *GetNetworkTrafficAnalysisSettingsOK {
	return &GetNetworkTrafficAnalysisSettingsOK{}
}

/*GetNetworkTrafficAnalysisSettingsOK handles this case with default header values.

Successful operation
*/
type GetNetworkTrafficAnalysisSettingsOK struct {
	Payload interface{}
}

func (o *GetNetworkTrafficAnalysisSettingsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/trafficAnalysisSettings][%d] getNetworkTrafficAnalysisSettingsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkTrafficAnalysisSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkTrafficAnalysisSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
