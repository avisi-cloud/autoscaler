// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package comprehendmedicaliface provides an interface to enable mocking the AWS Comprehend Medical service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package comprehendmedicaliface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/comprehendmedical"
)

// ComprehendMedicalAPI provides an interface to enable mocking the
// comprehendmedical.ComprehendMedical service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Comprehend Medical.
//    func myFunc(svc comprehendmedicaliface.ComprehendMedicalAPI) bool {
//        // Make svc.DescribeEntitiesDetectionV2Job request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := comprehendmedical.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComprehendMedicalClient struct {
//        comprehendmedicaliface.ComprehendMedicalAPI
//    }
//    func (m *mockComprehendMedicalClient) DescribeEntitiesDetectionV2Job(input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComprehendMedicalClient{}
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
type ComprehendMedicalAPI interface {
	DescribeEntitiesDetectionV2Job(*comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribeEntitiesDetectionV2JobWithContext(aws.Context, *comprehendmedical.DescribeEntitiesDetectionV2JobInput, ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribeEntitiesDetectionV2JobRequest(*comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.DescribeEntitiesDetectionV2JobOutput)

	DescribeICD10CMInferenceJob(*comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)
	DescribeICD10CMInferenceJobWithContext(aws.Context, *comprehendmedical.DescribeICD10CMInferenceJobInput, ...request.Option) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)
	DescribeICD10CMInferenceJobRequest(*comprehendmedical.DescribeICD10CMInferenceJobInput) (*request.Request, *comprehendmedical.DescribeICD10CMInferenceJobOutput)

	DescribePHIDetectionJob(*comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DescribePHIDetectionJobWithContext(aws.Context, *comprehendmedical.DescribePHIDetectionJobInput, ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DescribePHIDetectionJobRequest(*comprehendmedical.DescribePHIDetectionJobInput) (*request.Request, *comprehendmedical.DescribePHIDetectionJobOutput)

	DescribeRxNormInferenceJob(*comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)
	DescribeRxNormInferenceJobWithContext(aws.Context, *comprehendmedical.DescribeRxNormInferenceJobInput, ...request.Option) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)
	DescribeRxNormInferenceJobRequest(*comprehendmedical.DescribeRxNormInferenceJobInput) (*request.Request, *comprehendmedical.DescribeRxNormInferenceJobOutput)

	DescribeSNOMEDCTInferenceJob(*comprehendmedical.DescribeSNOMEDCTInferenceJobInput) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error)
	DescribeSNOMEDCTInferenceJobWithContext(aws.Context, *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, ...request.Option) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error)
	DescribeSNOMEDCTInferenceJobRequest(*comprehendmedical.DescribeSNOMEDCTInferenceJobInput) (*request.Request, *comprehendmedical.DescribeSNOMEDCTInferenceJobOutput)

	DetectEntities(*comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesWithContext(aws.Context, *comprehendmedical.DetectEntitiesInput, ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesRequest(*comprehendmedical.DetectEntitiesInput) (*request.Request, *comprehendmedical.DetectEntitiesOutput)

	DetectEntitiesV2(*comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectEntitiesV2WithContext(aws.Context, *comprehendmedical.DetectEntitiesV2Input, ...request.Option) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectEntitiesV2Request(*comprehendmedical.DetectEntitiesV2Input) (*request.Request, *comprehendmedical.DetectEntitiesV2Output)

	DetectPHI(*comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIWithContext(aws.Context, *comprehendmedical.DetectPHIInput, ...request.Option) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIRequest(*comprehendmedical.DetectPHIInput) (*request.Request, *comprehendmedical.DetectPHIOutput)

	InferICD10CM(*comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error)
	InferICD10CMWithContext(aws.Context, *comprehendmedical.InferICD10CMInput, ...request.Option) (*comprehendmedical.InferICD10CMOutput, error)
	InferICD10CMRequest(*comprehendmedical.InferICD10CMInput) (*request.Request, *comprehendmedical.InferICD10CMOutput)

	InferRxNorm(*comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error)
	InferRxNormWithContext(aws.Context, *comprehendmedical.InferRxNormInput, ...request.Option) (*comprehendmedical.InferRxNormOutput, error)
	InferRxNormRequest(*comprehendmedical.InferRxNormInput) (*request.Request, *comprehendmedical.InferRxNormOutput)

	InferSNOMEDCT(*comprehendmedical.InferSNOMEDCTInput) (*comprehendmedical.InferSNOMEDCTOutput, error)
	InferSNOMEDCTWithContext(aws.Context, *comprehendmedical.InferSNOMEDCTInput, ...request.Option) (*comprehendmedical.InferSNOMEDCTOutput, error)
	InferSNOMEDCTRequest(*comprehendmedical.InferSNOMEDCTInput) (*request.Request, *comprehendmedical.InferSNOMEDCTOutput)

	ListEntitiesDetectionV2Jobs(*comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListEntitiesDetectionV2JobsWithContext(aws.Context, *comprehendmedical.ListEntitiesDetectionV2JobsInput, ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListEntitiesDetectionV2JobsRequest(*comprehendmedical.ListEntitiesDetectionV2JobsInput) (*request.Request, *comprehendmedical.ListEntitiesDetectionV2JobsOutput)

	ListICD10CMInferenceJobs(*comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)
	ListICD10CMInferenceJobsWithContext(aws.Context, *comprehendmedical.ListICD10CMInferenceJobsInput, ...request.Option) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)
	ListICD10CMInferenceJobsRequest(*comprehendmedical.ListICD10CMInferenceJobsInput) (*request.Request, *comprehendmedical.ListICD10CMInferenceJobsOutput)

	ListPHIDetectionJobs(*comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	ListPHIDetectionJobsWithContext(aws.Context, *comprehendmedical.ListPHIDetectionJobsInput, ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	ListPHIDetectionJobsRequest(*comprehendmedical.ListPHIDetectionJobsInput) (*request.Request, *comprehendmedical.ListPHIDetectionJobsOutput)

	ListRxNormInferenceJobs(*comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)
	ListRxNormInferenceJobsWithContext(aws.Context, *comprehendmedical.ListRxNormInferenceJobsInput, ...request.Option) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)
	ListRxNormInferenceJobsRequest(*comprehendmedical.ListRxNormInferenceJobsInput) (*request.Request, *comprehendmedical.ListRxNormInferenceJobsOutput)

	ListSNOMEDCTInferenceJobs(*comprehendmedical.ListSNOMEDCTInferenceJobsInput) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error)
	ListSNOMEDCTInferenceJobsWithContext(aws.Context, *comprehendmedical.ListSNOMEDCTInferenceJobsInput, ...request.Option) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error)
	ListSNOMEDCTInferenceJobsRequest(*comprehendmedical.ListSNOMEDCTInferenceJobsInput) (*request.Request, *comprehendmedical.ListSNOMEDCTInferenceJobsOutput)

	StartEntitiesDetectionV2Job(*comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartEntitiesDetectionV2JobWithContext(aws.Context, *comprehendmedical.StartEntitiesDetectionV2JobInput, ...request.Option) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartEntitiesDetectionV2JobRequest(*comprehendmedical.StartEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.StartEntitiesDetectionV2JobOutput)

	StartICD10CMInferenceJob(*comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)
	StartICD10CMInferenceJobWithContext(aws.Context, *comprehendmedical.StartICD10CMInferenceJobInput, ...request.Option) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)
	StartICD10CMInferenceJobRequest(*comprehendmedical.StartICD10CMInferenceJobInput) (*request.Request, *comprehendmedical.StartICD10CMInferenceJobOutput)

	StartPHIDetectionJob(*comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StartPHIDetectionJobWithContext(aws.Context, *comprehendmedical.StartPHIDetectionJobInput, ...request.Option) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StartPHIDetectionJobRequest(*comprehendmedical.StartPHIDetectionJobInput) (*request.Request, *comprehendmedical.StartPHIDetectionJobOutput)

	StartRxNormInferenceJob(*comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error)
	StartRxNormInferenceJobWithContext(aws.Context, *comprehendmedical.StartRxNormInferenceJobInput, ...request.Option) (*comprehendmedical.StartRxNormInferenceJobOutput, error)
	StartRxNormInferenceJobRequest(*comprehendmedical.StartRxNormInferenceJobInput) (*request.Request, *comprehendmedical.StartRxNormInferenceJobOutput)

	StartSNOMEDCTInferenceJob(*comprehendmedical.StartSNOMEDCTInferenceJobInput) (*comprehendmedical.StartSNOMEDCTInferenceJobOutput, error)
	StartSNOMEDCTInferenceJobWithContext(aws.Context, *comprehendmedical.StartSNOMEDCTInferenceJobInput, ...request.Option) (*comprehendmedical.StartSNOMEDCTInferenceJobOutput, error)
	StartSNOMEDCTInferenceJobRequest(*comprehendmedical.StartSNOMEDCTInferenceJobInput) (*request.Request, *comprehendmedical.StartSNOMEDCTInferenceJobOutput)

	StopEntitiesDetectionV2Job(*comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopEntitiesDetectionV2JobWithContext(aws.Context, *comprehendmedical.StopEntitiesDetectionV2JobInput, ...request.Option) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopEntitiesDetectionV2JobRequest(*comprehendmedical.StopEntitiesDetectionV2JobInput) (*request.Request, *comprehendmedical.StopEntitiesDetectionV2JobOutput)

	StopICD10CMInferenceJob(*comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)
	StopICD10CMInferenceJobWithContext(aws.Context, *comprehendmedical.StopICD10CMInferenceJobInput, ...request.Option) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)
	StopICD10CMInferenceJobRequest(*comprehendmedical.StopICD10CMInferenceJobInput) (*request.Request, *comprehendmedical.StopICD10CMInferenceJobOutput)

	StopPHIDetectionJob(*comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error)
	StopPHIDetectionJobWithContext(aws.Context, *comprehendmedical.StopPHIDetectionJobInput, ...request.Option) (*comprehendmedical.StopPHIDetectionJobOutput, error)
	StopPHIDetectionJobRequest(*comprehendmedical.StopPHIDetectionJobInput) (*request.Request, *comprehendmedical.StopPHIDetectionJobOutput)

	StopRxNormInferenceJob(*comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error)
	StopRxNormInferenceJobWithContext(aws.Context, *comprehendmedical.StopRxNormInferenceJobInput, ...request.Option) (*comprehendmedical.StopRxNormInferenceJobOutput, error)
	StopRxNormInferenceJobRequest(*comprehendmedical.StopRxNormInferenceJobInput) (*request.Request, *comprehendmedical.StopRxNormInferenceJobOutput)

	StopSNOMEDCTInferenceJob(*comprehendmedical.StopSNOMEDCTInferenceJobInput) (*comprehendmedical.StopSNOMEDCTInferenceJobOutput, error)
	StopSNOMEDCTInferenceJobWithContext(aws.Context, *comprehendmedical.StopSNOMEDCTInferenceJobInput, ...request.Option) (*comprehendmedical.StopSNOMEDCTInferenceJobOutput, error)
	StopSNOMEDCTInferenceJobRequest(*comprehendmedical.StopSNOMEDCTInferenceJobInput) (*request.Request, *comprehendmedical.StopSNOMEDCTInferenceJobOutput)
}

var _ ComprehendMedicalAPI = (*comprehendmedical.ComprehendMedical)(nil)
