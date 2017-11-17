// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To create a new file system
//
// This operation creates a new file system with the default generalpurpose performance
// mode.
func ExampleEFS_CreateFileSystemRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.CreateFileSystemInput{
		CreationToken:   aws.String("tokenstring"),
		PerformanceMode: efs.PerformanceModeGeneralPurpose,
	}

	req := svc.CreateFileSystemRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemAlreadyExists:
				fmt.Println(efs.ErrCodeFileSystemAlreadyExists, aerr.Error())
			case efs.ErrCodeFileSystemLimitExceeded:
				fmt.Println(efs.ErrCodeFileSystemLimitExceeded, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new mount target
//
// This operation creates a new mount target for an EFS file system.
func ExampleEFS_CreateMountTargetRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.CreateMountTargetInput{
		FileSystemId: aws.String("fs-01234567"),
		SubnetId:     aws.String("subnet-1234abcd"),
	}

	req := svc.CreateMountTargetRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			case efs.ErrCodeIncorrectFileSystemLifeCycleState:
				fmt.Println(efs.ErrCodeIncorrectFileSystemLifeCycleState, aerr.Error())
			case efs.ErrCodeMountTargetConflict:
				fmt.Println(efs.ErrCodeMountTargetConflict, aerr.Error())
			case efs.ErrCodeSubnetNotFound:
				fmt.Println(efs.ErrCodeSubnetNotFound, aerr.Error())
			case efs.ErrCodeNoFreeAddressesInSubnet:
				fmt.Println(efs.ErrCodeNoFreeAddressesInSubnet, aerr.Error())
			case efs.ErrCodeIpAddressInUse:
				fmt.Println(efs.ErrCodeIpAddressInUse, aerr.Error())
			case efs.ErrCodeNetworkInterfaceLimitExceeded:
				fmt.Println(efs.ErrCodeNetworkInterfaceLimitExceeded, aerr.Error())
			case efs.ErrCodeSecurityGroupLimitExceeded:
				fmt.Println(efs.ErrCodeSecurityGroupLimitExceeded, aerr.Error())
			case efs.ErrCodeSecurityGroupNotFound:
				fmt.Println(efs.ErrCodeSecurityGroupNotFound, aerr.Error())
			case efs.ErrCodeUnsupportedAvailabilityZone:
				fmt.Println(efs.ErrCodeUnsupportedAvailabilityZone, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new tag
//
// This operation creates a new tag for an EFS file system.
func ExampleEFS_CreateTagsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.CreateTagsInput{
		FileSystemId: aws.String("fs-01234567"),
		Tags: []*efs.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
	}

	req := svc.CreateTagsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a file system
//
// This operation deletes an EFS file system.
func ExampleEFS_DeleteFileSystemRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DeleteFileSystemInput{
		FileSystemId: aws.String("fs-01234567"),
	}

	req := svc.DeleteFileSystemRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			case efs.ErrCodeFileSystemInUse:
				fmt.Println(efs.ErrCodeFileSystemInUse, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a mount target
//
// This operation deletes a mount target.
func ExampleEFS_DeleteMountTargetRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DeleteMountTargetInput{
		MountTargetId: aws.String("fsmt-12340abc"),
	}

	req := svc.DeleteMountTargetRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeDependencyTimeout:
				fmt.Println(efs.ErrCodeDependencyTimeout, aerr.Error())
			case efs.ErrCodeMountTargetNotFound:
				fmt.Println(efs.ErrCodeMountTargetNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete tags for an EFS file system
//
// This operation deletes tags for an EFS file system.
func ExampleEFS_DeleteTagsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DeleteTagsInput{
		FileSystemId: aws.String("fs-01234567"),
		TagKeys: []*string{
			aws.String("Name"),
		},
	}

	req := svc.DeleteTagsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe an EFS file system
//
// This operation describes all of the EFS file systems in an account.
func ExampleEFS_DescribeFileSystemsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DescribeFileSystemsInput{}

	req := svc.DescribeFileSystemsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe the security groups for a mount target
//
// This operation describes all of the security groups for a file system's mount target.
func ExampleEFS_DescribeMountTargetSecurityGroupsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DescribeMountTargetSecurityGroupsInput{
		MountTargetId: aws.String("fsmt-12340abc"),
	}

	req := svc.DescribeMountTargetSecurityGroupsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeMountTargetNotFound:
				fmt.Println(efs.ErrCodeMountTargetNotFound, aerr.Error())
			case efs.ErrCodeIncorrectMountTargetState:
				fmt.Println(efs.ErrCodeIncorrectMountTargetState, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe the mount targets for a file system
//
// This operation describes all of a file system's mount targets.
func ExampleEFS_DescribeMountTargetsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DescribeMountTargetsInput{
		FileSystemId: aws.String("fs-01234567"),
	}

	req := svc.DescribeMountTargetsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			case efs.ErrCodeMountTargetNotFound:
				fmt.Println(efs.ErrCodeMountTargetNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe the tags for a file system
//
// This operation describes all of a file system's tags.
func ExampleEFS_DescribeTagsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.DescribeTagsInput{
		FileSystemId: aws.String("fs-01234567"),
	}

	req := svc.DescribeTagsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemNotFound:
				fmt.Println(efs.ErrCodeFileSystemNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To modify the security groups associated with a mount target for a file system
//
// This operation modifies the security groups associated with a mount target for a
// file system.
func ExampleEFS_ModifyMountTargetSecurityGroupsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := efs.New(cfg)
	input := &efs.ModifyMountTargetSecurityGroupsInput{
		MountTargetId: aws.String("fsmt-12340abc"),
		SecurityGroups: []*string{
			aws.String("sg-abcd1234"),
		},
	}

	req := svc.ModifyMountTargetSecurityGroupsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeMountTargetNotFound:
				fmt.Println(efs.ErrCodeMountTargetNotFound, aerr.Error())
			case efs.ErrCodeIncorrectMountTargetState:
				fmt.Println(efs.ErrCodeIncorrectMountTargetState, aerr.Error())
			case efs.ErrCodeSecurityGroupLimitExceeded:
				fmt.Println(efs.ErrCodeSecurityGroupLimitExceeded, aerr.Error())
			case efs.ErrCodeSecurityGroupNotFound:
				fmt.Println(efs.ErrCodeSecurityGroupNotFound, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}