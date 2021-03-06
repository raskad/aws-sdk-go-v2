// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetGatewayGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the gateway group to get.
	//
	// GatewayGroupArn is a required field
	GatewayGroupArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetGatewayGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGatewayGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGatewayGroupInput"}

	if s.GatewayGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayGroupArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetGatewayGroupOutput struct {
	_ struct{} `type:"structure"`

	// The details of the gateway group.
	GatewayGroup *GatewayGroup `type:"structure"`
}

// String returns the string representation
func (s GetGatewayGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetGatewayGroup = "GetGatewayGroup"

// GetGatewayGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Retrieves the details of a gateway group.
//
//    // Example sending a request using GetGatewayGroupRequest.
//    req := client.GetGatewayGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetGatewayGroup
func (c *Client) GetGatewayGroupRequest(input *GetGatewayGroupInput) GetGatewayGroupRequest {
	op := &aws.Operation{
		Name:       opGetGatewayGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetGatewayGroupInput{}
	}

	req := c.newRequest(op, input, &GetGatewayGroupOutput{})
	return GetGatewayGroupRequest{Request: req, Input: input, Copy: c.GetGatewayGroupRequest}
}

// GetGatewayGroupRequest is the request type for the
// GetGatewayGroup API operation.
type GetGatewayGroupRequest struct {
	*aws.Request
	Input *GetGatewayGroupInput
	Copy  func(*GetGatewayGroupInput) GetGatewayGroupRequest
}

// Send marshals and sends the GetGatewayGroup API request.
func (r GetGatewayGroupRequest) Send(ctx context.Context) (*GetGatewayGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGatewayGroupResponse{
		GetGatewayGroupOutput: r.Request.Data.(*GetGatewayGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGatewayGroupResponse is the response type for the
// GetGatewayGroup API operation.
type GetGatewayGroupResponse struct {
	*GetGatewayGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGatewayGroup request.
func (r *GetGatewayGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
