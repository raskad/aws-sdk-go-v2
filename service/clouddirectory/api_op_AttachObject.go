// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AttachObjectInput struct {
	_ struct{} `type:"structure"`

	// The child object reference to be attached to the object.
	//
	// ChildReference is a required field
	ChildReference *ObjectReference `type:"structure" required:"true"`

	// Amazon Resource Name (ARN) that is associated with the Directory where both
	// objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The link name with which the child object is attached to the parent.
	//
	// LinkName is a required field
	LinkName *string `min:"1" type:"string" required:"true"`

	// The parent object reference.
	//
	// ParentReference is a required field
	ParentReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s AttachObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachObjectInput"}

	if s.ChildReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChildReference"))
	}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.LinkName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LinkName"))
	}
	if s.LinkName != nil && len(*s.LinkName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LinkName", 1))
	}

	if s.ParentReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachObjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChildReference != nil {
		v := s.ChildReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ChildReference", v, metadata)
	}
	if s.LinkName != nil {
		v := *s.LinkName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LinkName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentReference != nil {
		v := s.ParentReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ParentReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AttachObjectOutput struct {
	_ struct{} `type:"structure"`

	// The attached ObjectIdentifier, which is the child ObjectIdentifier.
	AttachedObjectIdentifier *string `type:"string"`
}

// String returns the string representation
func (s AttachObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AttachObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttachedObjectIdentifier != nil {
		v := *s.AttachedObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AttachedObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAttachObject = "AttachObject"

// AttachObjectRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches an existing object to another object. An object can be accessed
// in two ways:
//
// Using the path
//
// Using ObjectIdentifier
//
//    // Example sending a request using AttachObjectRequest.
//    req := client.AttachObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachObject
func (c *Client) AttachObjectRequest(input *AttachObjectInput) AttachObjectRequest {
	op := &aws.Operation{
		Name:       opAttachObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/attach",
	}

	if input == nil {
		input = &AttachObjectInput{}
	}

	req := c.newRequest(op, input, &AttachObjectOutput{})
	return AttachObjectRequest{Request: req, Input: input, Copy: c.AttachObjectRequest}
}

// AttachObjectRequest is the request type for the
// AttachObject API operation.
type AttachObjectRequest struct {
	*aws.Request
	Input *AttachObjectInput
	Copy  func(*AttachObjectInput) AttachObjectRequest
}

// Send marshals and sends the AttachObject API request.
func (r AttachObjectRequest) Send(ctx context.Context) (*AttachObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachObjectResponse{
		AttachObjectOutput: r.Request.Data.(*AttachObjectOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachObjectResponse is the response type for the
// AttachObject API operation.
type AttachObjectResponse struct {
	*AttachObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachObject request.
func (r *AttachObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
