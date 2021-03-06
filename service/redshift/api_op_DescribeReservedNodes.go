// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReservedNodesInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeReservedNodes request exceed
	// the value specified in MaxRecords, AWS returns a value in the Marker field
	// of the response. You can retrieve the next set of response records by providing
	// the returned marker value in the Marker parameter and retrying the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// Identifier for the node reservation.
	ReservedNodeId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedNodesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeReservedNodesOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// The list of ReservedNode objects.
	ReservedNodes []ReservedNode `locationNameList:"ReservedNode" type:"list"`
}

// String returns the string representation
func (s DescribeReservedNodesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedNodes = "DescribeReservedNodes"

// DescribeReservedNodesRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns the descriptions of the reserved nodes.
//
//    // Example sending a request using DescribeReservedNodesRequest.
//    req := client.DescribeReservedNodesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeReservedNodes
func (c *Client) DescribeReservedNodesRequest(input *DescribeReservedNodesInput) DescribeReservedNodesRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedNodes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedNodesInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedNodesOutput{})
	return DescribeReservedNodesRequest{Request: req, Input: input, Copy: c.DescribeReservedNodesRequest}
}

// DescribeReservedNodesRequest is the request type for the
// DescribeReservedNodes API operation.
type DescribeReservedNodesRequest struct {
	*aws.Request
	Input *DescribeReservedNodesInput
	Copy  func(*DescribeReservedNodesInput) DescribeReservedNodesRequest
}

// Send marshals and sends the DescribeReservedNodes API request.
func (r DescribeReservedNodesRequest) Send(ctx context.Context) (*DescribeReservedNodesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedNodesResponse{
		DescribeReservedNodesOutput: r.Request.Data.(*DescribeReservedNodesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedNodesRequestPaginator returns a paginator for DescribeReservedNodes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedNodesRequest(input)
//   p := redshift.NewDescribeReservedNodesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedNodesPaginator(req DescribeReservedNodesRequest) DescribeReservedNodesPaginator {
	return DescribeReservedNodesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedNodesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeReservedNodesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedNodesPaginator struct {
	aws.Pager
}

func (p *DescribeReservedNodesPaginator) CurrentPage() *DescribeReservedNodesOutput {
	return p.Pager.CurrentPage().(*DescribeReservedNodesOutput)
}

// DescribeReservedNodesResponse is the response type for the
// DescribeReservedNodes API operation.
type DescribeReservedNodesResponse struct {
	*DescribeReservedNodesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedNodes request.
func (r *DescribeReservedNodesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
