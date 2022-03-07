// Code generated by go-swagger; DO NOT EDIT.

package kubectl_cluster_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/api/def/clients/sentry/models"
)

// KubectlClusterSettingsGetKubectlClusterSettingsReader is a Reader for the KubectlClusterSettingsGetKubectlClusterSettings structure.
type KubectlClusterSettingsGetKubectlClusterSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubectlClusterSettingsGetKubectlClusterSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubectlClusterSettingsGetKubectlClusterSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewKubectlClusterSettingsGetKubectlClusterSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubectlClusterSettingsGetKubectlClusterSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewKubectlClusterSettingsGetKubectlClusterSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubectlClusterSettingsGetKubectlClusterSettingsOK creates a KubectlClusterSettingsGetKubectlClusterSettingsOK with default headers values
func NewKubectlClusterSettingsGetKubectlClusterSettingsOK() *KubectlClusterSettingsGetKubectlClusterSettingsOK {
	return &KubectlClusterSettingsGetKubectlClusterSettingsOK{}
}

/* KubectlClusterSettingsGetKubectlClusterSettingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type KubectlClusterSettingsGetKubectlClusterSettingsOK struct {
	Payload *models.RPCGetKubectlClusterSettingsResponse
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/kubectl/{opts.urlScope}/settings][%d] kubectlClusterSettingsGetKubectlClusterSettingsOK  %+v", 200, o.Payload)
}
func (o *KubectlClusterSettingsGetKubectlClusterSettingsOK) GetPayload() *models.RPCGetKubectlClusterSettingsResponse {
	return o.Payload
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCGetKubectlClusterSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubectlClusterSettingsGetKubectlClusterSettingsForbidden creates a KubectlClusterSettingsGetKubectlClusterSettingsForbidden with default headers values
func NewKubectlClusterSettingsGetKubectlClusterSettingsForbidden() *KubectlClusterSettingsGetKubectlClusterSettingsForbidden {
	return &KubectlClusterSettingsGetKubectlClusterSettingsForbidden{}
}

/* KubectlClusterSettingsGetKubectlClusterSettingsForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type KubectlClusterSettingsGetKubectlClusterSettingsForbidden struct {
	Payload interface{}
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/kubectl/{opts.urlScope}/settings][%d] kubectlClusterSettingsGetKubectlClusterSettingsForbidden  %+v", 403, o.Payload)
}
func (o *KubectlClusterSettingsGetKubectlClusterSettingsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubectlClusterSettingsGetKubectlClusterSettingsNotFound creates a KubectlClusterSettingsGetKubectlClusterSettingsNotFound with default headers values
func NewKubectlClusterSettingsGetKubectlClusterSettingsNotFound() *KubectlClusterSettingsGetKubectlClusterSettingsNotFound {
	return &KubectlClusterSettingsGetKubectlClusterSettingsNotFound{}
}

/* KubectlClusterSettingsGetKubectlClusterSettingsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type KubectlClusterSettingsGetKubectlClusterSettingsNotFound struct {
	Payload string
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/kubectl/{opts.urlScope}/settings][%d] kubectlClusterSettingsGetKubectlClusterSettingsNotFound  %+v", 404, o.Payload)
}
func (o *KubectlClusterSettingsGetKubectlClusterSettingsNotFound) GetPayload() string {
	return o.Payload
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubectlClusterSettingsGetKubectlClusterSettingsDefault creates a KubectlClusterSettingsGetKubectlClusterSettingsDefault with default headers values
func NewKubectlClusterSettingsGetKubectlClusterSettingsDefault(code int) *KubectlClusterSettingsGetKubectlClusterSettingsDefault {
	return &KubectlClusterSettingsGetKubectlClusterSettingsDefault{
		_statusCode: code,
	}
}

/* KubectlClusterSettingsGetKubectlClusterSettingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KubectlClusterSettingsGetKubectlClusterSettingsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the kubectl cluster settings get kubectl cluster settings default response
func (o *KubectlClusterSettingsGetKubectlClusterSettingsDefault) Code() int {
	return o._statusCode
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/kubectl/{opts.urlScope}/settings][%d] KubectlClusterSettings_GetKubectlClusterSettings default  %+v", o._statusCode, o.Payload)
}
func (o *KubectlClusterSettingsGetKubectlClusterSettingsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *KubectlClusterSettingsGetKubectlClusterSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}