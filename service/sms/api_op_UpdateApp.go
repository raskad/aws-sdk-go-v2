// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAppInput struct {
	_ struct{} `type:"structure"`

	// ID of the application to update.
	AppId *string `locationName:"appId" type:"string"`

	// New description of the application.
	Description *string `locationName:"description" type:"string"`

	// New name of the application.
	Name *string `locationName:"name" type:"string"`

	// Name of the service role in the customer's account used by AWS SMS.
	RoleName *string `locationName:"roleName" type:"string"`

	// List of server groups in the application to update.
	ServerGroups []ServerGroup `locationName:"serverGroups" type:"list"`

	// List of tags to associate with the application.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s UpdateAppInput) String() string {
	return awsutil.Prettify(s)
}

type UpdateAppOutput struct {
	_ struct{} `type:"structure"`

	// Summary description of the application.
	AppSummary *AppSummary `locationName:"appSummary" type:"structure"`

	// List of updated server groups in the application.
	ServerGroups []ServerGroup `locationName:"serverGroups" type:"list"`

	// List of tags associated with the application.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s UpdateAppOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApp = "UpdateApp"

// UpdateAppRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Updates an application.
//
//    // Example sending a request using UpdateAppRequest.
//    req := client.UpdateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/UpdateApp
func (c *Client) UpdateAppRequest(input *UpdateAppInput) UpdateAppRequest {
	op := &aws.Operation{
		Name:       opUpdateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAppInput{}
	}

	req := c.newRequest(op, input, &UpdateAppOutput{})
	return UpdateAppRequest{Request: req, Input: input, Copy: c.UpdateAppRequest}
}

// UpdateAppRequest is the request type for the
// UpdateApp API operation.
type UpdateAppRequest struct {
	*aws.Request
	Input *UpdateAppInput
	Copy  func(*UpdateAppInput) UpdateAppRequest
}

// Send marshals and sends the UpdateApp API request.
func (r UpdateAppRequest) Send(ctx context.Context) (*UpdateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAppResponse{
		UpdateAppOutput: r.Request.Data.(*UpdateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAppResponse is the response type for the
// UpdateApp API operation.
type UpdateAppResponse struct {
	*UpdateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApp request.
func (r *UpdateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
