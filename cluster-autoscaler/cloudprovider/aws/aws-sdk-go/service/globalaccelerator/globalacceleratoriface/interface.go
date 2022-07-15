// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package globalacceleratoriface provides an interface to enable mocking the AWS Global Accelerator service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package globalacceleratoriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/globalaccelerator"
)

// GlobalAcceleratorAPI provides an interface to enable mocking the
// globalaccelerator.GlobalAccelerator service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Global Accelerator.
//    func myFunc(svc globalacceleratoriface.GlobalAcceleratorAPI) bool {
//        // Make svc.AddCustomRoutingEndpoints request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := globalaccelerator.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGlobalAcceleratorClient struct {
//        globalacceleratoriface.GlobalAcceleratorAPI
//    }
//    func (m *mockGlobalAcceleratorClient) AddCustomRoutingEndpoints(input *globalaccelerator.AddCustomRoutingEndpointsInput) (*globalaccelerator.AddCustomRoutingEndpointsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGlobalAcceleratorClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type GlobalAcceleratorAPI interface {
	AddCustomRoutingEndpoints(*globalaccelerator.AddCustomRoutingEndpointsInput) (*globalaccelerator.AddCustomRoutingEndpointsOutput, error)
	AddCustomRoutingEndpointsWithContext(aws.Context, *globalaccelerator.AddCustomRoutingEndpointsInput, ...request.Option) (*globalaccelerator.AddCustomRoutingEndpointsOutput, error)
	AddCustomRoutingEndpointsRequest(*globalaccelerator.AddCustomRoutingEndpointsInput) (*request.Request, *globalaccelerator.AddCustomRoutingEndpointsOutput)

	AdvertiseByoipCidr(*globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error)
	AdvertiseByoipCidrWithContext(aws.Context, *globalaccelerator.AdvertiseByoipCidrInput, ...request.Option) (*globalaccelerator.AdvertiseByoipCidrOutput, error)
	AdvertiseByoipCidrRequest(*globalaccelerator.AdvertiseByoipCidrInput) (*request.Request, *globalaccelerator.AdvertiseByoipCidrOutput)

	AllowCustomRoutingTraffic(*globalaccelerator.AllowCustomRoutingTrafficInput) (*globalaccelerator.AllowCustomRoutingTrafficOutput, error)
	AllowCustomRoutingTrafficWithContext(aws.Context, *globalaccelerator.AllowCustomRoutingTrafficInput, ...request.Option) (*globalaccelerator.AllowCustomRoutingTrafficOutput, error)
	AllowCustomRoutingTrafficRequest(*globalaccelerator.AllowCustomRoutingTrafficInput) (*request.Request, *globalaccelerator.AllowCustomRoutingTrafficOutput)

	CreateAccelerator(*globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorWithContext(aws.Context, *globalaccelerator.CreateAcceleratorInput, ...request.Option) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorRequest(*globalaccelerator.CreateAcceleratorInput) (*request.Request, *globalaccelerator.CreateAcceleratorOutput)

	CreateCustomRoutingAccelerator(*globalaccelerator.CreateCustomRoutingAcceleratorInput) (*globalaccelerator.CreateCustomRoutingAcceleratorOutput, error)
	CreateCustomRoutingAcceleratorWithContext(aws.Context, *globalaccelerator.CreateCustomRoutingAcceleratorInput, ...request.Option) (*globalaccelerator.CreateCustomRoutingAcceleratorOutput, error)
	CreateCustomRoutingAcceleratorRequest(*globalaccelerator.CreateCustomRoutingAcceleratorInput) (*request.Request, *globalaccelerator.CreateCustomRoutingAcceleratorOutput)

	CreateCustomRoutingEndpointGroup(*globalaccelerator.CreateCustomRoutingEndpointGroupInput) (*globalaccelerator.CreateCustomRoutingEndpointGroupOutput, error)
	CreateCustomRoutingEndpointGroupWithContext(aws.Context, *globalaccelerator.CreateCustomRoutingEndpointGroupInput, ...request.Option) (*globalaccelerator.CreateCustomRoutingEndpointGroupOutput, error)
	CreateCustomRoutingEndpointGroupRequest(*globalaccelerator.CreateCustomRoutingEndpointGroupInput) (*request.Request, *globalaccelerator.CreateCustomRoutingEndpointGroupOutput)

	CreateCustomRoutingListener(*globalaccelerator.CreateCustomRoutingListenerInput) (*globalaccelerator.CreateCustomRoutingListenerOutput, error)
	CreateCustomRoutingListenerWithContext(aws.Context, *globalaccelerator.CreateCustomRoutingListenerInput, ...request.Option) (*globalaccelerator.CreateCustomRoutingListenerOutput, error)
	CreateCustomRoutingListenerRequest(*globalaccelerator.CreateCustomRoutingListenerInput) (*request.Request, *globalaccelerator.CreateCustomRoutingListenerOutput)

	CreateEndpointGroup(*globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupWithContext(aws.Context, *globalaccelerator.CreateEndpointGroupInput, ...request.Option) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupRequest(*globalaccelerator.CreateEndpointGroupInput) (*request.Request, *globalaccelerator.CreateEndpointGroupOutput)

	CreateListener(*globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerWithContext(aws.Context, *globalaccelerator.CreateListenerInput, ...request.Option) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerRequest(*globalaccelerator.CreateListenerInput) (*request.Request, *globalaccelerator.CreateListenerOutput)

	DeleteAccelerator(*globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorWithContext(aws.Context, *globalaccelerator.DeleteAcceleratorInput, ...request.Option) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorRequest(*globalaccelerator.DeleteAcceleratorInput) (*request.Request, *globalaccelerator.DeleteAcceleratorOutput)

	DeleteCustomRoutingAccelerator(*globalaccelerator.DeleteCustomRoutingAcceleratorInput) (*globalaccelerator.DeleteCustomRoutingAcceleratorOutput, error)
	DeleteCustomRoutingAcceleratorWithContext(aws.Context, *globalaccelerator.DeleteCustomRoutingAcceleratorInput, ...request.Option) (*globalaccelerator.DeleteCustomRoutingAcceleratorOutput, error)
	DeleteCustomRoutingAcceleratorRequest(*globalaccelerator.DeleteCustomRoutingAcceleratorInput) (*request.Request, *globalaccelerator.DeleteCustomRoutingAcceleratorOutput)

	DeleteCustomRoutingEndpointGroup(*globalaccelerator.DeleteCustomRoutingEndpointGroupInput) (*globalaccelerator.DeleteCustomRoutingEndpointGroupOutput, error)
	DeleteCustomRoutingEndpointGroupWithContext(aws.Context, *globalaccelerator.DeleteCustomRoutingEndpointGroupInput, ...request.Option) (*globalaccelerator.DeleteCustomRoutingEndpointGroupOutput, error)
	DeleteCustomRoutingEndpointGroupRequest(*globalaccelerator.DeleteCustomRoutingEndpointGroupInput) (*request.Request, *globalaccelerator.DeleteCustomRoutingEndpointGroupOutput)

	DeleteCustomRoutingListener(*globalaccelerator.DeleteCustomRoutingListenerInput) (*globalaccelerator.DeleteCustomRoutingListenerOutput, error)
	DeleteCustomRoutingListenerWithContext(aws.Context, *globalaccelerator.DeleteCustomRoutingListenerInput, ...request.Option) (*globalaccelerator.DeleteCustomRoutingListenerOutput, error)
	DeleteCustomRoutingListenerRequest(*globalaccelerator.DeleteCustomRoutingListenerInput) (*request.Request, *globalaccelerator.DeleteCustomRoutingListenerOutput)

	DeleteEndpointGroup(*globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupWithContext(aws.Context, *globalaccelerator.DeleteEndpointGroupInput, ...request.Option) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupRequest(*globalaccelerator.DeleteEndpointGroupInput) (*request.Request, *globalaccelerator.DeleteEndpointGroupOutput)

	DeleteListener(*globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerWithContext(aws.Context, *globalaccelerator.DeleteListenerInput, ...request.Option) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerRequest(*globalaccelerator.DeleteListenerInput) (*request.Request, *globalaccelerator.DeleteListenerOutput)

	DenyCustomRoutingTraffic(*globalaccelerator.DenyCustomRoutingTrafficInput) (*globalaccelerator.DenyCustomRoutingTrafficOutput, error)
	DenyCustomRoutingTrafficWithContext(aws.Context, *globalaccelerator.DenyCustomRoutingTrafficInput, ...request.Option) (*globalaccelerator.DenyCustomRoutingTrafficOutput, error)
	DenyCustomRoutingTrafficRequest(*globalaccelerator.DenyCustomRoutingTrafficInput) (*request.Request, *globalaccelerator.DenyCustomRoutingTrafficOutput)

	DeprovisionByoipCidr(*globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error)
	DeprovisionByoipCidrWithContext(aws.Context, *globalaccelerator.DeprovisionByoipCidrInput, ...request.Option) (*globalaccelerator.DeprovisionByoipCidrOutput, error)
	DeprovisionByoipCidrRequest(*globalaccelerator.DeprovisionByoipCidrInput) (*request.Request, *globalaccelerator.DeprovisionByoipCidrOutput)

	DescribeAccelerator(*globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorWithContext(aws.Context, *globalaccelerator.DescribeAcceleratorInput, ...request.Option) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorRequest(*globalaccelerator.DescribeAcceleratorInput) (*request.Request, *globalaccelerator.DescribeAcceleratorOutput)

	DescribeAcceleratorAttributes(*globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.DescribeAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesRequest(*globalaccelerator.DescribeAcceleratorAttributesInput) (*request.Request, *globalaccelerator.DescribeAcceleratorAttributesOutput)

	DescribeCustomRoutingAccelerator(*globalaccelerator.DescribeCustomRoutingAcceleratorInput) (*globalaccelerator.DescribeCustomRoutingAcceleratorOutput, error)
	DescribeCustomRoutingAcceleratorWithContext(aws.Context, *globalaccelerator.DescribeCustomRoutingAcceleratorInput, ...request.Option) (*globalaccelerator.DescribeCustomRoutingAcceleratorOutput, error)
	DescribeCustomRoutingAcceleratorRequest(*globalaccelerator.DescribeCustomRoutingAcceleratorInput) (*request.Request, *globalaccelerator.DescribeCustomRoutingAcceleratorOutput)

	DescribeCustomRoutingAcceleratorAttributes(*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput) (*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput, error)
	DescribeCustomRoutingAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput, error)
	DescribeCustomRoutingAcceleratorAttributesRequest(*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput) (*request.Request, *globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput)

	DescribeCustomRoutingEndpointGroup(*globalaccelerator.DescribeCustomRoutingEndpointGroupInput) (*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput, error)
	DescribeCustomRoutingEndpointGroupWithContext(aws.Context, *globalaccelerator.DescribeCustomRoutingEndpointGroupInput, ...request.Option) (*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput, error)
	DescribeCustomRoutingEndpointGroupRequest(*globalaccelerator.DescribeCustomRoutingEndpointGroupInput) (*request.Request, *globalaccelerator.DescribeCustomRoutingEndpointGroupOutput)

	DescribeCustomRoutingListener(*globalaccelerator.DescribeCustomRoutingListenerInput) (*globalaccelerator.DescribeCustomRoutingListenerOutput, error)
	DescribeCustomRoutingListenerWithContext(aws.Context, *globalaccelerator.DescribeCustomRoutingListenerInput, ...request.Option) (*globalaccelerator.DescribeCustomRoutingListenerOutput, error)
	DescribeCustomRoutingListenerRequest(*globalaccelerator.DescribeCustomRoutingListenerInput) (*request.Request, *globalaccelerator.DescribeCustomRoutingListenerOutput)

	DescribeEndpointGroup(*globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupWithContext(aws.Context, *globalaccelerator.DescribeEndpointGroupInput, ...request.Option) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupRequest(*globalaccelerator.DescribeEndpointGroupInput) (*request.Request, *globalaccelerator.DescribeEndpointGroupOutput)

	DescribeListener(*globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerWithContext(aws.Context, *globalaccelerator.DescribeListenerInput, ...request.Option) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerRequest(*globalaccelerator.DescribeListenerInput) (*request.Request, *globalaccelerator.DescribeListenerOutput)

	ListAccelerators(*globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsWithContext(aws.Context, *globalaccelerator.ListAcceleratorsInput, ...request.Option) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsRequest(*globalaccelerator.ListAcceleratorsInput) (*request.Request, *globalaccelerator.ListAcceleratorsOutput)

	ListAcceleratorsPages(*globalaccelerator.ListAcceleratorsInput, func(*globalaccelerator.ListAcceleratorsOutput, bool) bool) error
	ListAcceleratorsPagesWithContext(aws.Context, *globalaccelerator.ListAcceleratorsInput, func(*globalaccelerator.ListAcceleratorsOutput, bool) bool, ...request.Option) error

	ListByoipCidrs(*globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error)
	ListByoipCidrsWithContext(aws.Context, *globalaccelerator.ListByoipCidrsInput, ...request.Option) (*globalaccelerator.ListByoipCidrsOutput, error)
	ListByoipCidrsRequest(*globalaccelerator.ListByoipCidrsInput) (*request.Request, *globalaccelerator.ListByoipCidrsOutput)

	ListByoipCidrsPages(*globalaccelerator.ListByoipCidrsInput, func(*globalaccelerator.ListByoipCidrsOutput, bool) bool) error
	ListByoipCidrsPagesWithContext(aws.Context, *globalaccelerator.ListByoipCidrsInput, func(*globalaccelerator.ListByoipCidrsOutput, bool) bool, ...request.Option) error

	ListCustomRoutingAccelerators(*globalaccelerator.ListCustomRoutingAcceleratorsInput) (*globalaccelerator.ListCustomRoutingAcceleratorsOutput, error)
	ListCustomRoutingAcceleratorsWithContext(aws.Context, *globalaccelerator.ListCustomRoutingAcceleratorsInput, ...request.Option) (*globalaccelerator.ListCustomRoutingAcceleratorsOutput, error)
	ListCustomRoutingAcceleratorsRequest(*globalaccelerator.ListCustomRoutingAcceleratorsInput) (*request.Request, *globalaccelerator.ListCustomRoutingAcceleratorsOutput)

	ListCustomRoutingAcceleratorsPages(*globalaccelerator.ListCustomRoutingAcceleratorsInput, func(*globalaccelerator.ListCustomRoutingAcceleratorsOutput, bool) bool) error
	ListCustomRoutingAcceleratorsPagesWithContext(aws.Context, *globalaccelerator.ListCustomRoutingAcceleratorsInput, func(*globalaccelerator.ListCustomRoutingAcceleratorsOutput, bool) bool, ...request.Option) error

	ListCustomRoutingEndpointGroups(*globalaccelerator.ListCustomRoutingEndpointGroupsInput) (*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, error)
	ListCustomRoutingEndpointGroupsWithContext(aws.Context, *globalaccelerator.ListCustomRoutingEndpointGroupsInput, ...request.Option) (*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, error)
	ListCustomRoutingEndpointGroupsRequest(*globalaccelerator.ListCustomRoutingEndpointGroupsInput) (*request.Request, *globalaccelerator.ListCustomRoutingEndpointGroupsOutput)

	ListCustomRoutingEndpointGroupsPages(*globalaccelerator.ListCustomRoutingEndpointGroupsInput, func(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, bool) bool) error
	ListCustomRoutingEndpointGroupsPagesWithContext(aws.Context, *globalaccelerator.ListCustomRoutingEndpointGroupsInput, func(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, bool) bool, ...request.Option) error

	ListCustomRoutingListeners(*globalaccelerator.ListCustomRoutingListenersInput) (*globalaccelerator.ListCustomRoutingListenersOutput, error)
	ListCustomRoutingListenersWithContext(aws.Context, *globalaccelerator.ListCustomRoutingListenersInput, ...request.Option) (*globalaccelerator.ListCustomRoutingListenersOutput, error)
	ListCustomRoutingListenersRequest(*globalaccelerator.ListCustomRoutingListenersInput) (*request.Request, *globalaccelerator.ListCustomRoutingListenersOutput)

	ListCustomRoutingListenersPages(*globalaccelerator.ListCustomRoutingListenersInput, func(*globalaccelerator.ListCustomRoutingListenersOutput, bool) bool) error
	ListCustomRoutingListenersPagesWithContext(aws.Context, *globalaccelerator.ListCustomRoutingListenersInput, func(*globalaccelerator.ListCustomRoutingListenersOutput, bool) bool, ...request.Option) error

	ListCustomRoutingPortMappings(*globalaccelerator.ListCustomRoutingPortMappingsInput) (*globalaccelerator.ListCustomRoutingPortMappingsOutput, error)
	ListCustomRoutingPortMappingsWithContext(aws.Context, *globalaccelerator.ListCustomRoutingPortMappingsInput, ...request.Option) (*globalaccelerator.ListCustomRoutingPortMappingsOutput, error)
	ListCustomRoutingPortMappingsRequest(*globalaccelerator.ListCustomRoutingPortMappingsInput) (*request.Request, *globalaccelerator.ListCustomRoutingPortMappingsOutput)

	ListCustomRoutingPortMappingsPages(*globalaccelerator.ListCustomRoutingPortMappingsInput, func(*globalaccelerator.ListCustomRoutingPortMappingsOutput, bool) bool) error
	ListCustomRoutingPortMappingsPagesWithContext(aws.Context, *globalaccelerator.ListCustomRoutingPortMappingsInput, func(*globalaccelerator.ListCustomRoutingPortMappingsOutput, bool) bool, ...request.Option) error

	ListCustomRoutingPortMappingsByDestination(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput) (*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, error)
	ListCustomRoutingPortMappingsByDestinationWithContext(aws.Context, *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, ...request.Option) (*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, error)
	ListCustomRoutingPortMappingsByDestinationRequest(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput) (*request.Request, *globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput)

	ListCustomRoutingPortMappingsByDestinationPages(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, func(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, bool) bool) error
	ListCustomRoutingPortMappingsByDestinationPagesWithContext(aws.Context, *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, func(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, bool) bool, ...request.Option) error

	ListEndpointGroups(*globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsWithContext(aws.Context, *globalaccelerator.ListEndpointGroupsInput, ...request.Option) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsRequest(*globalaccelerator.ListEndpointGroupsInput) (*request.Request, *globalaccelerator.ListEndpointGroupsOutput)

	ListEndpointGroupsPages(*globalaccelerator.ListEndpointGroupsInput, func(*globalaccelerator.ListEndpointGroupsOutput, bool) bool) error
	ListEndpointGroupsPagesWithContext(aws.Context, *globalaccelerator.ListEndpointGroupsInput, func(*globalaccelerator.ListEndpointGroupsOutput, bool) bool, ...request.Option) error

	ListListeners(*globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error)
	ListListenersWithContext(aws.Context, *globalaccelerator.ListListenersInput, ...request.Option) (*globalaccelerator.ListListenersOutput, error)
	ListListenersRequest(*globalaccelerator.ListListenersInput) (*request.Request, *globalaccelerator.ListListenersOutput)

	ListListenersPages(*globalaccelerator.ListListenersInput, func(*globalaccelerator.ListListenersOutput, bool) bool) error
	ListListenersPagesWithContext(aws.Context, *globalaccelerator.ListListenersInput, func(*globalaccelerator.ListListenersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *globalaccelerator.ListTagsForResourceInput, ...request.Option) (*globalaccelerator.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*globalaccelerator.ListTagsForResourceInput) (*request.Request, *globalaccelerator.ListTagsForResourceOutput)

	ProvisionByoipCidr(*globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error)
	ProvisionByoipCidrWithContext(aws.Context, *globalaccelerator.ProvisionByoipCidrInput, ...request.Option) (*globalaccelerator.ProvisionByoipCidrOutput, error)
	ProvisionByoipCidrRequest(*globalaccelerator.ProvisionByoipCidrInput) (*request.Request, *globalaccelerator.ProvisionByoipCidrOutput)

	RemoveCustomRoutingEndpoints(*globalaccelerator.RemoveCustomRoutingEndpointsInput) (*globalaccelerator.RemoveCustomRoutingEndpointsOutput, error)
	RemoveCustomRoutingEndpointsWithContext(aws.Context, *globalaccelerator.RemoveCustomRoutingEndpointsInput, ...request.Option) (*globalaccelerator.RemoveCustomRoutingEndpointsOutput, error)
	RemoveCustomRoutingEndpointsRequest(*globalaccelerator.RemoveCustomRoutingEndpointsInput) (*request.Request, *globalaccelerator.RemoveCustomRoutingEndpointsOutput)

	TagResource(*globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *globalaccelerator.TagResourceInput, ...request.Option) (*globalaccelerator.TagResourceOutput, error)
	TagResourceRequest(*globalaccelerator.TagResourceInput) (*request.Request, *globalaccelerator.TagResourceOutput)

	UntagResource(*globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *globalaccelerator.UntagResourceInput, ...request.Option) (*globalaccelerator.UntagResourceOutput, error)
	UntagResourceRequest(*globalaccelerator.UntagResourceInput) (*request.Request, *globalaccelerator.UntagResourceOutput)

	UpdateAccelerator(*globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorWithContext(aws.Context, *globalaccelerator.UpdateAcceleratorInput, ...request.Option) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorRequest(*globalaccelerator.UpdateAcceleratorInput) (*request.Request, *globalaccelerator.UpdateAcceleratorOutput)

	UpdateAcceleratorAttributes(*globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.UpdateAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesRequest(*globalaccelerator.UpdateAcceleratorAttributesInput) (*request.Request, *globalaccelerator.UpdateAcceleratorAttributesOutput)

	UpdateCustomRoutingAccelerator(*globalaccelerator.UpdateCustomRoutingAcceleratorInput) (*globalaccelerator.UpdateCustomRoutingAcceleratorOutput, error)
	UpdateCustomRoutingAcceleratorWithContext(aws.Context, *globalaccelerator.UpdateCustomRoutingAcceleratorInput, ...request.Option) (*globalaccelerator.UpdateCustomRoutingAcceleratorOutput, error)
	UpdateCustomRoutingAcceleratorRequest(*globalaccelerator.UpdateCustomRoutingAcceleratorInput) (*request.Request, *globalaccelerator.UpdateCustomRoutingAcceleratorOutput)

	UpdateCustomRoutingAcceleratorAttributes(*globalaccelerator.UpdateCustomRoutingAcceleratorAttributesInput) (*globalaccelerator.UpdateCustomRoutingAcceleratorAttributesOutput, error)
	UpdateCustomRoutingAcceleratorAttributesWithContext(aws.Context, *globalaccelerator.UpdateCustomRoutingAcceleratorAttributesInput, ...request.Option) (*globalaccelerator.UpdateCustomRoutingAcceleratorAttributesOutput, error)
	UpdateCustomRoutingAcceleratorAttributesRequest(*globalaccelerator.UpdateCustomRoutingAcceleratorAttributesInput) (*request.Request, *globalaccelerator.UpdateCustomRoutingAcceleratorAttributesOutput)

	UpdateCustomRoutingListener(*globalaccelerator.UpdateCustomRoutingListenerInput) (*globalaccelerator.UpdateCustomRoutingListenerOutput, error)
	UpdateCustomRoutingListenerWithContext(aws.Context, *globalaccelerator.UpdateCustomRoutingListenerInput, ...request.Option) (*globalaccelerator.UpdateCustomRoutingListenerOutput, error)
	UpdateCustomRoutingListenerRequest(*globalaccelerator.UpdateCustomRoutingListenerInput) (*request.Request, *globalaccelerator.UpdateCustomRoutingListenerOutput)

	UpdateEndpointGroup(*globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupWithContext(aws.Context, *globalaccelerator.UpdateEndpointGroupInput, ...request.Option) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupRequest(*globalaccelerator.UpdateEndpointGroupInput) (*request.Request, *globalaccelerator.UpdateEndpointGroupOutput)

	UpdateListener(*globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerWithContext(aws.Context, *globalaccelerator.UpdateListenerInput, ...request.Option) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerRequest(*globalaccelerator.UpdateListenerInput) (*request.Request, *globalaccelerator.UpdateListenerOutput)

	WithdrawByoipCidr(*globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error)
	WithdrawByoipCidrWithContext(aws.Context, *globalaccelerator.WithdrawByoipCidrInput, ...request.Option) (*globalaccelerator.WithdrawByoipCidrOutput, error)
	WithdrawByoipCidrRequest(*globalaccelerator.WithdrawByoipCidrInput) (*request.Request, *globalaccelerator.WithdrawByoipCidrOutput)
}

var _ GlobalAcceleratorAPI = (*globalaccelerator.GlobalAccelerator)(nil)
