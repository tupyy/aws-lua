package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/tupyy/aws-lua/internal/lua"
)

type AwsProvider struct {
	config ClientConfiguration
}

func New(c ClientConfiguration) *AwsProvider {
	return &AwsProvider{c}
}

func (a *AwsProvider) Create(ctx context.Context, resource string, o lua.Object) (lua.Object, error) {
	var opFunc func(ctx context.Context, o lua.Object) (lua.Object, error)

	switch resource {
	case lua.AwsUser:
		opFunc = NewBuilder[iam.CreateUserInput, iam.CreateUserOutput](a.config).
			Type(IamClient).
			Op(CreateUser).
			TransformInputFunc(toCreateUserInput).
			TransformOutputFunc(fromCreateUserOutput).
			Build(ctx)
	case lua.AwsAccessKey:
		opFunc = NewBuilder[iam.CreateAccessKeyInput, iam.CreateAccessKeyOutput](a.config).
			Type(IamClient).
			Op(CreateAccessKeys).
			TransformInputFunc(toCreateAccessKeyInput).
			TransformOutputFunc(fromCreateAccessKeyOutput).
			Build(ctx)
	case lua.AwsSubnet:
		opFunc = NewBuilder[ec2.CreateSubnetInput, ec2.CreateSubnetOutput](a.config).
			Type(Ec2Client).
			Op(CreateSubnet).
			TransformInputFunc(toCreateSubnetInput).
			TransformOutputFunc(fromCreateSubnetOutput).
			Build(ctx)
	case lua.AwsVpc:
		opFunc = NewBuilder[ec2.CreateVpcInput, ec2.CreateVpcOutput](a.config).
			Type(Ec2Client).
			Op(CreateVpc).
			TransformInputFunc(toCreateVpcInput).
			TransformOutputFunc(fromCreateVpcOutput).
			Build(ctx)
	case lua.AwsInternetGateway:
		opFunc = NewBuilder[ec2.CreateInternetGatewayInput, ec2.CreateInternetGatewayOutput](a.config).
			Type(Ec2Client).
			Op(CreateIgw).
			TransformInputFunc(toCreateIgwInput).
			TransformOutputFunc(fromCreateIgwOutput).
			Build(ctx)
	case lua.AwsNatGateway:
		opFunc = NewBuilder[ec2.CreateNatGatewayInput, ec2.CreateNatGatewayOutput](a.config).
			Type(Ec2Client).
			Op(CreateNat).
			TransformInputFunc(toCreateNatInput).
			TransformOutputFunc(fromCreateNatOutput).
			Build(ctx)

	default:
		return lua.Object{}, fmt.Errorf("unknown resource")
	}

	return opFunc(ctx, o)
}

func (a *AwsProvider) Delete(ctx context.Context, resource string, o lua.Object) (lua.Object, error) {
	return lua.Object{}, fmt.Errorf("unknown resource")
}

func (a *AwsProvider) List(ctx context.Context, resource string, o lua.Object) (lua.Object, error) {
	var opFunc func(ctx context.Context, o lua.Object) (lua.Object, error)
	switch resource {
	case lua.AwsUser:
		opFunc = NewBuilder[iam.ListUsersInput, iam.ListUsersOutput](a.config).
			Type(IamClient).
			Op(ListUsers).
			TransformInputFunc(toListUserInput).
			TransformOutputFunc(fromListUserOutput).
			Build(ctx)
	case lua.AwsVpc:
		opFunc = NewBuilder[ec2.DescribeVpcsInput, ec2.DescribeVpcsOutput](a.config).
			Type(Ec2Client).
			Op(ListVpcs).
			TransformInputFunc(toDescribeVpcsInput).
			TransformOutputFunc(fromDescribeVpcsOutput).
			Build(ctx)
	case lua.AwsSubnet:
		opFunc = NewBuilder[ec2.DescribeSubnetsInput, ec2.DescribeSubnetsOutput](a.config).
			Type(Ec2Client).
			Op(ListSubnets).
			TransformInputFunc(toDescribeSubnetsInput).
			TransformOutputFunc(fromDescribeSubnetsOutput).
			Build(ctx)
	case lua.AwsAZs:
		opFunc = NewBuilder[ec2.DescribeAvailabilityZonesInput, ec2.DescribeAvailabilityZonesOutput](a.config).
			Type(Ec2Client).
			Op(ListAvailabilityZones).
			TransformInputFunc(toDescribeAZsInput).
			TransformOutputFunc(fromDescribeAZsOutput).
			Build(ctx)
	case lua.AwsInternetGateway:
		opFunc = NewBuilder[ec2.DescribeInternetGatewaysInput, ec2.DescribeInternetGatewaysOutput](a.config).
			Type(Ec2Client).
			Op(ListIgws).
			TransformInputFunc(toDescribeIgwInput).
			TransformOutputFunc(fromDescribeIgwOutput).
			Build(ctx)
	case lua.AwsNatGateway:
		opFunc = NewBuilder[ec2.DescribeNatGatewaysInput, ec2.DescribeNatGatewaysOutput](a.config).
			Type(Ec2Client).
			Op(ListNats).
			TransformInputFunc(toDescribeNatInput).
			TransformOutputFunc(fromDescribeNatOutput).
			Build(ctx)
	default:
		return lua.Object{}, fmt.Errorf("unknown resource")
	}

	return opFunc(ctx, o)

}
