// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilCodeBindingExists uses the Schemas API operation
// DescribeCodeBinding to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilCodeBindingExists(ctx context.Context, input *DescribeCodeBindingInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilCodeBindingExists",
		MaxAttempts: 30,
		Delay:       aws.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_IN_PROGRESS",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_FAILED",
			},
			{
				State:    aws.FailureWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "NotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeCodeBindingInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeCodeBindingRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
