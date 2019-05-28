// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to generate the CNAME records needed to set up Easy
// DKIM with Amazon SES. For more information about setting up Easy DKIM, see
// the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyDomainDkimRequest
type VerifyDomainDkimInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to be verified for Easy DKIM signing.
	//
	// Domain is a required field
	Domain *string `type:"string" required:"true"`
}

// String returns the string representation
func (s VerifyDomainDkimInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyDomainDkimInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifyDomainDkimInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns CNAME records that you must publish to the DNS server of your domain
// to set up Easy DKIM with Amazon SES.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyDomainDkimResponse
type VerifyDomainDkimOutput struct {
	_ struct{} `type:"structure"`

	// A set of character strings that represent the domain's identity. If the identity
	// is an email address, the tokens represent the domain of that address.
	//
	// Using these tokens, you will need to create DNS CNAME records that point
	// to DKIM public keys hosted by Amazon SES. Amazon Web Services will eventually
	// detect that you have updated your DNS records; this detection process may
	// take up to 72 hours. Upon successful detection, Amazon SES will be able to
	// DKIM-sign emails originating from that domain.
	//
	// For more information about creating DNS records using DKIM tokens, go to
	// the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
	//
	// DkimTokens is a required field
	DkimTokens []string `type:"list" required:"true"`
}

// String returns the string representation
func (s VerifyDomainDkimOutput) String() string {
	return awsutil.Prettify(s)
}

const opVerifyDomainDkim = "VerifyDomainDkim"

// VerifyDomainDkimRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns a set of DKIM tokens for a domain. DKIM tokens are character strings
// that represent your domain's identity. Using these tokens, you will need
// to create DNS CNAME records that point to DKIM public keys hosted by Amazon
// SES. Amazon Web Services will eventually detect that you have updated your
// DNS records; this detection process may take up to 72 hours. Upon successful
// detection, Amazon SES will be able to DKIM-sign email originating from that
// domain.
//
// You can execute this operation no more than once per second.
//
// To enable or disable Easy DKIM signing for a domain, use the SetIdentityDkimEnabled
// operation.
//
// For more information about creating DNS records using DKIM tokens, go to
// the Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim-dns-records.html).
//
//    // Example sending a request using VerifyDomainDkimRequest.
//    req := client.VerifyDomainDkimRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/VerifyDomainDkim
func (c *Client) VerifyDomainDkimRequest(input *VerifyDomainDkimInput) VerifyDomainDkimRequest {
	op := &aws.Operation{
		Name:       opVerifyDomainDkim,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyDomainDkimInput{}
	}

	req := c.newRequest(op, input, &VerifyDomainDkimOutput{})
	return VerifyDomainDkimRequest{Request: req, Input: input, Copy: c.VerifyDomainDkimRequest}
}

// VerifyDomainDkimRequest is the request type for the
// VerifyDomainDkim API operation.
type VerifyDomainDkimRequest struct {
	*aws.Request
	Input *VerifyDomainDkimInput
	Copy  func(*VerifyDomainDkimInput) VerifyDomainDkimRequest
}

// Send marshals and sends the VerifyDomainDkim API request.
func (r VerifyDomainDkimRequest) Send(ctx context.Context) (*VerifyDomainDkimResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifyDomainDkimResponse{
		VerifyDomainDkimOutput: r.Request.Data.(*VerifyDomainDkimOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifyDomainDkimResponse is the response type for the
// VerifyDomainDkim API operation.
type VerifyDomainDkimResponse struct {
	*VerifyDomainDkimOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifyDomainDkim request.
func (r *VerifyDomainDkimResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}