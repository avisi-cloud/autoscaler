// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53recoverycluster

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have sufficient permissions to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There was a conflict with this request. Try again.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeEndpointTemporarilyUnavailableException for service response error code
	// "EndpointTemporarilyUnavailableException".
	//
	// The cluster endpoint isn't available. Try another cluster endpoint.
	ErrCodeEndpointTemporarilyUnavailableException = "EndpointTemporarilyUnavailableException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// There was an unexpected error during processing of the request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The request references a routing control or control panel that was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceLimitExceededException for service response error code
	// "ServiceLimitExceededException".
	//
	// The request can't update that many routing control states at the same time.
	// Try again with fewer routing control states.
	ErrCodeServiceLimitExceededException = "ServiceLimitExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied because of request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// There was a validation error on the request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                   newErrorAccessDeniedException,
	"ConflictException":                       newErrorConflictException,
	"EndpointTemporarilyUnavailableException": newErrorEndpointTemporarilyUnavailableException,
	"InternalServerException":                 newErrorInternalServerException,
	"ResourceNotFoundException":               newErrorResourceNotFoundException,
	"ServiceLimitExceededException":           newErrorServiceLimitExceededException,
	"ThrottlingException":                     newErrorThrottlingException,
	"ValidationException":                     newErrorValidationException,
}
