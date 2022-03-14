// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerAuthInfoUserReader is a Reader for the ServiceBrokerAuthInfoUser structure.
type ServiceBrokerAuthInfoUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthInfoUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthInfoUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewServiceBrokerAuthInfoUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerAuthInfoUserOK creates a ServiceBrokerAuthInfoUserOK with default headers values
func NewServiceBrokerAuthInfoUserOK() *ServiceBrokerAuthInfoUserOK {
	return &ServiceBrokerAuthInfoUserOK{}
}

/* ServiceBrokerAuthInfoUserOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthInfoUserOK struct {
	Payload *models.UserInfo
}

func (o *ServiceBrokerAuthInfoUserOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/user][%d] serviceBrokerAuthInfoUserOK  %+v", 200, o.Payload)
}
func (o *ServiceBrokerAuthInfoUserOK) GetPayload() *models.UserInfo {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoUserInternalServerError creates a ServiceBrokerAuthInfoUserInternalServerError with default headers values
func NewServiceBrokerAuthInfoUserInternalServerError() *ServiceBrokerAuthInfoUserInternalServerError {
	return &ServiceBrokerAuthInfoUserInternalServerError{}
}

/* ServiceBrokerAuthInfoUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthInfoUserInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthInfoUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/user][%d] serviceBrokerAuthInfoUserInternalServerError  %+v", 500, o.Payload)
}
func (o *ServiceBrokerAuthInfoUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthInfoUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
