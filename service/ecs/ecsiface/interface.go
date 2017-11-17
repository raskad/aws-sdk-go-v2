// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecsiface provides an interface to enable mocking the Amazon EC2 Container Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecsiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// ECSAPI provides an interface to enable mocking the
// ecs.ECS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EC2 Container Service.
//    func myFunc(svc ecsiface.ECSAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := ecs.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockECSClient struct {
//        ecsiface.ECSAPI
//    }
//    func (m *mockECSClient) CreateCluster(input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockECSClient{}
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
type ECSAPI interface {
	CreateClusterRequest(*ecs.CreateClusterInput) ecs.CreateClusterRequest

	CreateServiceRequest(*ecs.CreateServiceInput) ecs.CreateServiceRequest

	DeleteAttributesRequest(*ecs.DeleteAttributesInput) ecs.DeleteAttributesRequest

	DeleteClusterRequest(*ecs.DeleteClusterInput) ecs.DeleteClusterRequest

	DeleteServiceRequest(*ecs.DeleteServiceInput) ecs.DeleteServiceRequest

	DeregisterContainerInstanceRequest(*ecs.DeregisterContainerInstanceInput) ecs.DeregisterContainerInstanceRequest

	DeregisterTaskDefinitionRequest(*ecs.DeregisterTaskDefinitionInput) ecs.DeregisterTaskDefinitionRequest

	DescribeClustersRequest(*ecs.DescribeClustersInput) ecs.DescribeClustersRequest

	DescribeContainerInstancesRequest(*ecs.DescribeContainerInstancesInput) ecs.DescribeContainerInstancesRequest

	DescribeServicesRequest(*ecs.DescribeServicesInput) ecs.DescribeServicesRequest

	DescribeTaskDefinitionRequest(*ecs.DescribeTaskDefinitionInput) ecs.DescribeTaskDefinitionRequest

	DescribeTasksRequest(*ecs.DescribeTasksInput) ecs.DescribeTasksRequest

	DiscoverPollEndpointRequest(*ecs.DiscoverPollEndpointInput) ecs.DiscoverPollEndpointRequest

	ListAttributesRequest(*ecs.ListAttributesInput) ecs.ListAttributesRequest

	ListClustersRequest(*ecs.ListClustersInput) ecs.ListClustersRequest

	ListClustersPages(*ecs.ListClustersInput, func(*ecs.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *ecs.ListClustersInput, func(*ecs.ListClustersOutput, bool) bool, ...aws.Option) error

	ListContainerInstancesRequest(*ecs.ListContainerInstancesInput) ecs.ListContainerInstancesRequest

	ListContainerInstancesPages(*ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool) error
	ListContainerInstancesPagesWithContext(aws.Context, *ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool, ...aws.Option) error

	ListServicesRequest(*ecs.ListServicesInput) ecs.ListServicesRequest

	ListServicesPages(*ecs.ListServicesInput, func(*ecs.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *ecs.ListServicesInput, func(*ecs.ListServicesOutput, bool) bool, ...aws.Option) error

	ListTaskDefinitionFamiliesRequest(*ecs.ListTaskDefinitionFamiliesInput) ecs.ListTaskDefinitionFamiliesRequest

	ListTaskDefinitionFamiliesPages(*ecs.ListTaskDefinitionFamiliesInput, func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool) error
	ListTaskDefinitionFamiliesPagesWithContext(aws.Context, *ecs.ListTaskDefinitionFamiliesInput, func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool, ...aws.Option) error

	ListTaskDefinitionsRequest(*ecs.ListTaskDefinitionsInput) ecs.ListTaskDefinitionsRequest

	ListTaskDefinitionsPages(*ecs.ListTaskDefinitionsInput, func(*ecs.ListTaskDefinitionsOutput, bool) bool) error
	ListTaskDefinitionsPagesWithContext(aws.Context, *ecs.ListTaskDefinitionsInput, func(*ecs.ListTaskDefinitionsOutput, bool) bool, ...aws.Option) error

	ListTasksRequest(*ecs.ListTasksInput) ecs.ListTasksRequest

	ListTasksPages(*ecs.ListTasksInput, func(*ecs.ListTasksOutput, bool) bool) error
	ListTasksPagesWithContext(aws.Context, *ecs.ListTasksInput, func(*ecs.ListTasksOutput, bool) bool, ...aws.Option) error

	PutAttributesRequest(*ecs.PutAttributesInput) ecs.PutAttributesRequest

	RegisterContainerInstanceRequest(*ecs.RegisterContainerInstanceInput) ecs.RegisterContainerInstanceRequest

	RegisterTaskDefinitionRequest(*ecs.RegisterTaskDefinitionInput) ecs.RegisterTaskDefinitionRequest

	RunTaskRequest(*ecs.RunTaskInput) ecs.RunTaskRequest

	StartTaskRequest(*ecs.StartTaskInput) ecs.StartTaskRequest

	StopTaskRequest(*ecs.StopTaskInput) ecs.StopTaskRequest

	SubmitContainerStateChangeRequest(*ecs.SubmitContainerStateChangeInput) ecs.SubmitContainerStateChangeRequest

	SubmitTaskStateChangeRequest(*ecs.SubmitTaskStateChangeInput) ecs.SubmitTaskStateChangeRequest

	UpdateContainerAgentRequest(*ecs.UpdateContainerAgentInput) ecs.UpdateContainerAgentRequest

	UpdateContainerInstancesStateRequest(*ecs.UpdateContainerInstancesStateInput) ecs.UpdateContainerInstancesStateRequest

	UpdateServiceRequest(*ecs.UpdateServiceInput) ecs.UpdateServiceRequest

	WaitUntilServicesInactive(*ecs.DescribeServicesInput) error
	WaitUntilServicesInactiveWithContext(aws.Context, *ecs.DescribeServicesInput, ...aws.WaiterOption) error

	WaitUntilServicesStable(*ecs.DescribeServicesInput) error
	WaitUntilServicesStableWithContext(aws.Context, *ecs.DescribeServicesInput, ...aws.WaiterOption) error

	WaitUntilTasksRunning(*ecs.DescribeTasksInput) error
	WaitUntilTasksRunningWithContext(aws.Context, *ecs.DescribeTasksInput, ...aws.WaiterOption) error

	WaitUntilTasksStopped(*ecs.DescribeTasksInput) error
	WaitUntilTasksStoppedWithContext(aws.Context, *ecs.DescribeTasksInput, ...aws.WaiterOption) error
}

var _ ECSAPI = (*ecs.ECS)(nil)