// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra-client-go/models"
)

// DeleteOAuth2TokenReader is a Reader for the DeleteOAuth2Token structure.
type DeleteOAuth2TokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOAuth2TokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOAuth2TokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOAuth2TokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOAuth2TokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOAuth2TokenNoContent creates a DeleteOAuth2TokenNoContent with default headers values
func NewDeleteOAuth2TokenNoContent() *DeleteOAuth2TokenNoContent {
	return &DeleteOAuth2TokenNoContent{}
}

/*DeleteOAuth2TokenNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteOAuth2TokenNoContent struct {
}

func (o *DeleteOAuth2TokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2TokenNoContent ", 204)
}

func (o *DeleteOAuth2TokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOAuth2TokenUnauthorized creates a DeleteOAuth2TokenUnauthorized with default headers values
func NewDeleteOAuth2TokenUnauthorized() *DeleteOAuth2TokenUnauthorized {
	return &DeleteOAuth2TokenUnauthorized{}
}

/*DeleteOAuth2TokenUnauthorized handles this case with default header values.

genericError
*/
type DeleteOAuth2TokenUnauthorized struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2TokenUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2TokenUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOAuth2TokenUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteOAuth2TokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOAuth2TokenInternalServerError creates a DeleteOAuth2TokenInternalServerError with default headers values
func NewDeleteOAuth2TokenInternalServerError() *DeleteOAuth2TokenInternalServerError {
	return &DeleteOAuth2TokenInternalServerError{}
}

/*DeleteOAuth2TokenInternalServerError handles this case with default header values.

genericError
*/
type DeleteOAuth2TokenInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteOAuth2TokenInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2TokenInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOAuth2TokenInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteOAuth2TokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
