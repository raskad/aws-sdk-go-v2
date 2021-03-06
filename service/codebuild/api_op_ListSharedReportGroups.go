// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSharedReportGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of paginated shared report groups per response. Use nextToken
	// to iterate pages in the list of returned ReportGroup objects. The default
	// value is 100.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// During a previous call, the maximum number of items that can be returned
	// is the value specified in maxResults. If there more items in the list, then
	// a unique string called a nextToken is returned. To get the next batch of
	// items in the list, call this operation again, adding the next token to the
	// call. To get all of the items in the list, keep calling this operation with
	// each subsequent next token that is returned, until no more next tokens are
	// returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The criterion to be used to list report groups shared with the current AWS
	// account or user. Valid values include:
	//
	//    * ARN: List based on the ARN.
	//
	//    * MODIFIED_TIME: List based on when information about the shared report
	//    group was last changed.
	SortBy SharedResourceSortByType `locationName:"sortBy" type:"string" enum:"true"`

	// The order in which to list shared report groups. Valid values include:
	//
	//    * ASCENDING: List in ascending order.
	//
	//    * DESCENDING: List in descending order.
	SortOrder SortOrderType `locationName:"sortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListSharedReportGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSharedReportGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSharedReportGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSharedReportGroupsOutput struct {
	_ struct{} `type:"structure"`

	// During a previous call, the maximum number of items that can be returned
	// is the value specified in maxResults. If there more items in the list, then
	// a unique string called a nextToken is returned. To get the next batch of
	// items in the list, call this operation again, adding the next token to the
	// call. To get all of the items in the list, keep calling this operation with
	// each subsequent next token that is returned, until no more next tokens are
	// returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of ARNs for the report groups shared with the current AWS account
	// or user.
	ReportGroups []string `locationName:"reportGroups" min:"1" type:"list"`
}

// String returns the string representation
func (s ListSharedReportGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSharedReportGroups = "ListSharedReportGroups"

// ListSharedReportGroupsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Gets a list of report groups that are shared with other AWS accounts or users.
//
//    // Example sending a request using ListSharedReportGroupsRequest.
//    req := client.ListSharedReportGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/ListSharedReportGroups
func (c *Client) ListSharedReportGroupsRequest(input *ListSharedReportGroupsInput) ListSharedReportGroupsRequest {
	op := &aws.Operation{
		Name:       opListSharedReportGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSharedReportGroupsInput{}
	}

	req := c.newRequest(op, input, &ListSharedReportGroupsOutput{})
	return ListSharedReportGroupsRequest{Request: req, Input: input, Copy: c.ListSharedReportGroupsRequest}
}

// ListSharedReportGroupsRequest is the request type for the
// ListSharedReportGroups API operation.
type ListSharedReportGroupsRequest struct {
	*aws.Request
	Input *ListSharedReportGroupsInput
	Copy  func(*ListSharedReportGroupsInput) ListSharedReportGroupsRequest
}

// Send marshals and sends the ListSharedReportGroups API request.
func (r ListSharedReportGroupsRequest) Send(ctx context.Context) (*ListSharedReportGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSharedReportGroupsResponse{
		ListSharedReportGroupsOutput: r.Request.Data.(*ListSharedReportGroupsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSharedReportGroupsResponse is the response type for the
// ListSharedReportGroups API operation.
type ListSharedReportGroupsResponse struct {
	*ListSharedReportGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSharedReportGroups request.
func (r *ListSharedReportGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
