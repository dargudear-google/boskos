// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumegroupsPutReader is a Reader for the PcloudVolumegroupsPut structure.
type PcloudVolumegroupsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVolumegroupsPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumegroupsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumegroupsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudVolumegroupsPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVolumegroupsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}] pcloud.volumegroups.put", response, response.Code())
	}
}

// NewPcloudVolumegroupsPutAccepted creates a PcloudVolumegroupsPutAccepted with default headers values
func NewPcloudVolumegroupsPutAccepted() *PcloudVolumegroupsPutAccepted {
	return &PcloudVolumegroupsPutAccepted{}
}

/*
PcloudVolumegroupsPutAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVolumegroupsPutAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud volumegroups put accepted response has a 2xx status code
func (o *PcloudVolumegroupsPutAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumegroups put accepted response has a 3xx status code
func (o *PcloudVolumegroupsPutAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put accepted response has a 4xx status code
func (o *PcloudVolumegroupsPutAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups put accepted response has a 5xx status code
func (o *PcloudVolumegroupsPutAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put accepted response a status code equal to that given
func (o *PcloudVolumegroupsPutAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud volumegroups put accepted response
func (o *PcloudVolumegroupsPutAccepted) Code() int {
	return 202
}

func (o *PcloudVolumegroupsPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVolumegroupsPutAccepted) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVolumegroupsPutAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudVolumegroupsPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutBadRequest creates a PcloudVolumegroupsPutBadRequest with default headers values
func NewPcloudVolumegroupsPutBadRequest() *PcloudVolumegroupsPutBadRequest {
	return &PcloudVolumegroupsPutBadRequest{}
}

/*
PcloudVolumegroupsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put bad request response has a 2xx status code
func (o *PcloudVolumegroupsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put bad request response has a 3xx status code
func (o *PcloudVolumegroupsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put bad request response has a 4xx status code
func (o *PcloudVolumegroupsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put bad request response has a 5xx status code
func (o *PcloudVolumegroupsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put bad request response a status code equal to that given
func (o *PcloudVolumegroupsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud volumegroups put bad request response
func (o *PcloudVolumegroupsPutBadRequest) Code() int {
	return 400
}

func (o *PcloudVolumegroupsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutUnauthorized creates a PcloudVolumegroupsPutUnauthorized with default headers values
func NewPcloudVolumegroupsPutUnauthorized() *PcloudVolumegroupsPutUnauthorized {
	return &PcloudVolumegroupsPutUnauthorized{}
}

/*
PcloudVolumegroupsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumegroupsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put unauthorized response has a 2xx status code
func (o *PcloudVolumegroupsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put unauthorized response has a 3xx status code
func (o *PcloudVolumegroupsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put unauthorized response has a 4xx status code
func (o *PcloudVolumegroupsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put unauthorized response has a 5xx status code
func (o *PcloudVolumegroupsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put unauthorized response a status code equal to that given
func (o *PcloudVolumegroupsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud volumegroups put unauthorized response
func (o *PcloudVolumegroupsPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudVolumegroupsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutForbidden creates a PcloudVolumegroupsPutForbidden with default headers values
func NewPcloudVolumegroupsPutForbidden() *PcloudVolumegroupsPutForbidden {
	return &PcloudVolumegroupsPutForbidden{}
}

/*
PcloudVolumegroupsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put forbidden response has a 2xx status code
func (o *PcloudVolumegroupsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put forbidden response has a 3xx status code
func (o *PcloudVolumegroupsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put forbidden response has a 4xx status code
func (o *PcloudVolumegroupsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put forbidden response has a 5xx status code
func (o *PcloudVolumegroupsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put forbidden response a status code equal to that given
func (o *PcloudVolumegroupsPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud volumegroups put forbidden response
func (o *PcloudVolumegroupsPutForbidden) Code() int {
	return 403
}

func (o *PcloudVolumegroupsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutNotFound creates a PcloudVolumegroupsPutNotFound with default headers values
func NewPcloudVolumegroupsPutNotFound() *PcloudVolumegroupsPutNotFound {
	return &PcloudVolumegroupsPutNotFound{}
}

/*
PcloudVolumegroupsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumegroupsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put not found response has a 2xx status code
func (o *PcloudVolumegroupsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put not found response has a 3xx status code
func (o *PcloudVolumegroupsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put not found response has a 4xx status code
func (o *PcloudVolumegroupsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put not found response has a 5xx status code
func (o *PcloudVolumegroupsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put not found response a status code equal to that given
func (o *PcloudVolumegroupsPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud volumegroups put not found response
func (o *PcloudVolumegroupsPutNotFound) Code() int {
	return 404
}

func (o *PcloudVolumegroupsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVolumegroupsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVolumegroupsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutConflict creates a PcloudVolumegroupsPutConflict with default headers values
func NewPcloudVolumegroupsPutConflict() *PcloudVolumegroupsPutConflict {
	return &PcloudVolumegroupsPutConflict{}
}

/*
PcloudVolumegroupsPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudVolumegroupsPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put conflict response has a 2xx status code
func (o *PcloudVolumegroupsPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put conflict response has a 3xx status code
func (o *PcloudVolumegroupsPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put conflict response has a 4xx status code
func (o *PcloudVolumegroupsPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put conflict response has a 5xx status code
func (o *PcloudVolumegroupsPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put conflict response a status code equal to that given
func (o *PcloudVolumegroupsPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud volumegroups put conflict response
func (o *PcloudVolumegroupsPutConflict) Code() int {
	return 409
}

func (o *PcloudVolumegroupsPutConflict) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutConflict  %+v", 409, o.Payload)
}

func (o *PcloudVolumegroupsPutConflict) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutConflict  %+v", 409, o.Payload)
}

func (o *PcloudVolumegroupsPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutUnprocessableEntity creates a PcloudVolumegroupsPutUnprocessableEntity with default headers values
func NewPcloudVolumegroupsPutUnprocessableEntity() *PcloudVolumegroupsPutUnprocessableEntity {
	return &PcloudVolumegroupsPutUnprocessableEntity{}
}

/*
PcloudVolumegroupsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVolumegroupsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put unprocessable entity response has a 2xx status code
func (o *PcloudVolumegroupsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put unprocessable entity response has a 3xx status code
func (o *PcloudVolumegroupsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put unprocessable entity response has a 4xx status code
func (o *PcloudVolumegroupsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups put unprocessable entity response has a 5xx status code
func (o *PcloudVolumegroupsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups put unprocessable entity response a status code equal to that given
func (o *PcloudVolumegroupsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud volumegroups put unprocessable entity response
func (o *PcloudVolumegroupsPutUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVolumegroupsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVolumegroupsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVolumegroupsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsPutInternalServerError creates a PcloudVolumegroupsPutInternalServerError with default headers values
func NewPcloudVolumegroupsPutInternalServerError() *PcloudVolumegroupsPutInternalServerError {
	return &PcloudVolumegroupsPutInternalServerError{}
}

/*
PcloudVolumegroupsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups put internal server error response has a 2xx status code
func (o *PcloudVolumegroupsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups put internal server error response has a 3xx status code
func (o *PcloudVolumegroupsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups put internal server error response has a 4xx status code
func (o *PcloudVolumegroupsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups put internal server error response has a 5xx status code
func (o *PcloudVolumegroupsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumegroups put internal server error response a status code equal to that given
func (o *PcloudVolumegroupsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud volumegroups put internal server error response
func (o *PcloudVolumegroupsPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudVolumegroupsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}