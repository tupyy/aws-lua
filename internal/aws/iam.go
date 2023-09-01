package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/tupyy/aws-lua/internal/lua"
)

func createIamClient(ctx context.Context, accessKey, secretKey, region string) (*iam.Client, error) {
	optFn := func(opts *awsConfig.LoadOptions) error {
		opts.Region = region
		opts.Credentials = getAwsCredentials(ctx, accessKey, secretKey)
		return nil
	}
	awsConfig, err := awsConfig.LoadDefaultConfig(ctx, optFn)
	if err != nil {
		return nil, fmt.Errorf("failed to create aws config for account: %w", err)
	}
	return iam.NewFromConfig(awsConfig), nil
}

func iamGetOpFunc(opType OpType, c ClientConfiguration) opFunc {
	switch opType {
	case CreateUser:
		return createUserFunc(c)
	case GetUser:
		return deleteUserFunc(c)
	case DeleteUser:
		return deleteAccessKeyFunc(c)
	case ListUsers:
		return listAccessKeysFunc(c)
	case CreateAccessKeys:
		return createAccessKeyFunc(c)
	case ListAccessKeys:
		return listAccessKeysFunc(c)
	case DeleteAccessKeys:
		return deleteAccessKeyFunc(c)
	default:
		return func(ctx context.Context, input interface{}) (interface{}, error) {
			return nil, errors.New("unknown op type for iam client")
		}
	}
}

/**
	op functions for User resource
**/
func createUserFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.CreateUserInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		output, err := iamClient.CreateUser(ctx, &i)
		return *output, err
	}
}

func getUserFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.GetUserInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.GetUser(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteUserFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.DeleteUserInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.DeleteUser(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listUserFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.ListUsersInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.ListUsers(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
 op functions for AccessKey resource
**/

func createAccessKeyFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.CreateAccessKeyInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.CreateAccessKey(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listAccessKeysFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.ListAccessKeysInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.ListAccessKeys(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteAccessKeyFunc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(iam.DeleteAccessKeyInput)
		iamClient, err := createIamClient(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := iamClient.DeleteAccessKey(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
	Transformation function
**/

func toCreateUserInput(o lua.Object) iam.CreateUserInput {
	username := o.GetString("username")
	if username == "" {
		return iam.CreateUserInput{}
	}
	return iam.CreateUserInput{UserName: aws.String(username)}
}

func fromCreateUserOutput(o iam.CreateUserOutput) lua.Object {
	if o.User == nil {
		return lua.Object{}
	}
	return lua.Object{
		"username": o.User.UserName,
		"arn":      o.User.Arn,
	}
}

func toListUserInput(o lua.Object) iam.ListUsersInput {
	return iam.ListUsersInput{}
}

func fromListUserOutput(o iam.ListUsersOutput) lua.Object {
	return toLua(o)
}

func toGetUserInput(o lua.Object) iam.GetUserInput {
	username := o.GetString("username")
	if username == "" {
		return iam.GetUserInput{}
	}
	return iam.GetUserInput{
		UserName: aws.String(username),
	}
}

func fromGetUserOutput(o iam.GetUserOutput) lua.Object {
	if o.User == nil {
		return lua.Object{}
	}
	l := lua.Object{
		"username": o.User.UserName,
		"arn":      o.User.Arn,
	}
	if len(o.User.Tags) > 0 {
		tags := make(map[string]string)
		for _, t := range o.User.Tags {
			tags[*t.Key] = *t.Value
		}
		l["tags"] = tags
	}
	return l
}

func toCreateAccessKeyInput(o lua.Object) iam.CreateAccessKeyInput {
	username := o.GetString("username")
	if username == "" {
		return iam.CreateAccessKeyInput{}
	}
	return iam.CreateAccessKeyInput{
		UserName: aws.String(username),
	}
}

func fromCreateAccessKeyOutput(o iam.CreateAccessKeyOutput) lua.Object {
	if o.AccessKey == nil {
		return lua.Object{}
	}
	return lua.Object{
		"access_key":        o.AccessKey.AccessKeyId,
		"secret_access_key": o.AccessKey.SecretAccessKey,
	}
}
