// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeregisterClusterReader is a Reader for the DeregisterCluster structure.
type DeregisterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeregisterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeregisterClusterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeregisterClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeregisterClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeregisterClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeregisterClusterNoContent creates a DeregisterClusterNoContent with default headers values
func NewDeregisterClusterNoContent() *DeregisterClusterNoContent {
	return &DeregisterClusterNoContent{}
}

/*DeregisterClusterNoContent handles this case with default header values.

Cluster deregistered
*/
type DeregisterClusterNoContent struct {
}

func (o *DeregisterClusterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}][%d] deregisterClusterNoContent ", 204)
}

func (o *DeregisterClusterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterClusterNotFound creates a DeregisterClusterNotFound with default headers values
func NewDeregisterClusterNotFound() *DeregisterClusterNotFound {
	return &DeregisterClusterNotFound{}
}

/*DeregisterClusterNotFound handles this case with default header values.

Cluster not found
*/
type DeregisterClusterNotFound struct {
}

func (o *DeregisterClusterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}][%d] deregisterClusterNotFound ", 404)
}

func (o *DeregisterClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterClusterConflict creates a DeregisterClusterConflict with default headers values
func NewDeregisterClusterConflict() *DeregisterClusterConflict {
	return &DeregisterClusterConflict{}
}

/*DeregisterClusterConflict handles this case with default header values.

Invalid state
*/
type DeregisterClusterConflict struct {
}

func (o *DeregisterClusterConflict) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}][%d] deregisterClusterConflict ", 409)
}

func (o *DeregisterClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterClusterInternalServerError creates a DeregisterClusterInternalServerError with default headers values
func NewDeregisterClusterInternalServerError() *DeregisterClusterInternalServerError {
	return &DeregisterClusterInternalServerError{}
}

/*DeregisterClusterInternalServerError handles this case with default header values.

Internal server error
*/
type DeregisterClusterInternalServerError struct {
}

func (o *DeregisterClusterInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{cluster_id}][%d] deregisterClusterInternalServerError ", 500)
}

func (o *DeregisterClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
