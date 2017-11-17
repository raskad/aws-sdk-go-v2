// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package configserviceiface provides an interface to enable mocking the AWS Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package configserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

// ConfigServiceAPI provides an interface to enable mocking the
// configservice.ConfigService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Config.
//    func myFunc(svc configserviceiface.ConfigServiceAPI) bool {
//        // Make svc.DeleteConfigRule request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := configservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConfigServiceClient struct {
//        configserviceiface.ConfigServiceAPI
//    }
//    func (m *mockConfigServiceClient) DeleteConfigRule(input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConfigServiceClient{}
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
type ConfigServiceAPI interface {
	DeleteConfigRuleRequest(*configservice.DeleteConfigRuleInput) configservice.DeleteConfigRuleRequest

	DeleteConfigurationRecorderRequest(*configservice.DeleteConfigurationRecorderInput) configservice.DeleteConfigurationRecorderRequest

	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) configservice.DeleteDeliveryChannelRequest

	DeleteEvaluationResultsRequest(*configservice.DeleteEvaluationResultsInput) configservice.DeleteEvaluationResultsRequest

	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) configservice.DeliverConfigSnapshotRequest

	DescribeComplianceByConfigRuleRequest(*configservice.DescribeComplianceByConfigRuleInput) configservice.DescribeComplianceByConfigRuleRequest

	DescribeComplianceByResourceRequest(*configservice.DescribeComplianceByResourceInput) configservice.DescribeComplianceByResourceRequest

	DescribeConfigRuleEvaluationStatusRequest(*configservice.DescribeConfigRuleEvaluationStatusInput) configservice.DescribeConfigRuleEvaluationStatusRequest

	DescribeConfigRulesRequest(*configservice.DescribeConfigRulesInput) configservice.DescribeConfigRulesRequest

	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) configservice.DescribeConfigurationRecorderStatusRequest

	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) configservice.DescribeConfigurationRecordersRequest

	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) configservice.DescribeDeliveryChannelStatusRequest

	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) configservice.DescribeDeliveryChannelsRequest

	GetComplianceDetailsByConfigRuleRequest(*configservice.GetComplianceDetailsByConfigRuleInput) configservice.GetComplianceDetailsByConfigRuleRequest

	GetComplianceDetailsByResourceRequest(*configservice.GetComplianceDetailsByResourceInput) configservice.GetComplianceDetailsByResourceRequest

	GetComplianceSummaryByConfigRuleRequest(*configservice.GetComplianceSummaryByConfigRuleInput) configservice.GetComplianceSummaryByConfigRuleRequest

	GetComplianceSummaryByResourceTypeRequest(*configservice.GetComplianceSummaryByResourceTypeInput) configservice.GetComplianceSummaryByResourceTypeRequest

	GetDiscoveredResourceCountsRequest(*configservice.GetDiscoveredResourceCountsInput) configservice.GetDiscoveredResourceCountsRequest

	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) configservice.GetResourceConfigHistoryRequest

	GetResourceConfigHistoryPages(*configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error
	GetResourceConfigHistoryPagesWithContext(aws.Context, *configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool, ...aws.Option) error

	ListDiscoveredResourcesRequest(*configservice.ListDiscoveredResourcesInput) configservice.ListDiscoveredResourcesRequest

	PutConfigRuleRequest(*configservice.PutConfigRuleInput) configservice.PutConfigRuleRequest

	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) configservice.PutConfigurationRecorderRequest

	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) configservice.PutDeliveryChannelRequest

	PutEvaluationsRequest(*configservice.PutEvaluationsInput) configservice.PutEvaluationsRequest

	StartConfigRulesEvaluationRequest(*configservice.StartConfigRulesEvaluationInput) configservice.StartConfigRulesEvaluationRequest

	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) configservice.StartConfigurationRecorderRequest

	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) configservice.StopConfigurationRecorderRequest
}

var _ ConfigServiceAPI = (*configservice.ConfigService)(nil)