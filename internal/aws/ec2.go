package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/tupyy/aws-lua/internal/lua"
)

func createEc2Client(ctx context.Context, accessKey, secretKey, region string) (*ec2.Client, error) {
	optFn := func(opts *awsConfig.LoadOptions) error {
		opts.Region = region
		opts.Credentials = getAwsCredentials(ctx, accessKey, secretKey)
		return nil
	}
	awsConfig, err := awsConfig.LoadDefaultConfig(ctx, optFn)
	if err != nil {
		return nil, fmt.Errorf("failed to create aws config for account: %w", err)
	}
	return ec2.NewFromConfig(awsConfig), nil
}

func ec2GetOpFunc(opType OpType, c ClientConfiguration) opFunc {
	switch opType {
	case CreateVpc:
		return createVpc(c)
	case ListVpcs:
		return listVpcs(c)
	case DeleteVpc:
		return deleteVpc(c)
	case CreateSubnet:
		return createSubnet(c)
	case ListSubnets:
		return listSubnets(c)
	case DeleteSubnet:
		return deleteSubnet(c)
	case CreateIgw:
		return createIgw(c)
	case DeleteIgw:
		return deleteIgw(c)
	case ListIgws:
		return listIgws(c)
	case ListAvailabilityZones:
		return listAZs(c)
	case CreateNat:
		return createNat(c)
	case DeleteNat:
		return deleteNat(c)
	case ListNats:
		return listNats(c)
	default:
		return func(ctx context.Context, input interface{}) (interface{}, error) {
			return nil, errors.New("unknown op type for ec2 client")
		}
	}
}

/**
 Op functions for VPC
**/

func createVpc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.CreateVpcInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.CreateVpc(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listVpcs(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DescribeVpcsInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DescribeVpcs(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteVpc(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DeleteVpcInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DeleteVpc(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listAZs(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DescribeAvailabilityZonesInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DescribeAvailabilityZones(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
 Op functions for subnets
**/

func createSubnet(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.CreateSubnetInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.CreateSubnet(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteSubnet(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DeleteSubnetInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DeleteSubnet(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listSubnets(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DescribeSubnetsInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DescribeSubnets(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
 Op for IGW
**/
func createIgw(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.CreateInternetGatewayInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.CreateInternetGateway(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteIgw(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DeleteInternetGatewayInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DeleteInternetGateway(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listIgws(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DescribeInternetGatewaysInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DescribeInternetGateways(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
 Op for NAT
**/
func createNat(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.CreateNatGatewayInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.CreateNatGateway(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func deleteNat(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DeleteNatGatewayInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DeleteNatGateway(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

func listNats(config ClientConfiguration) opFunc {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		i := input.(ec2.DescribeNatGatewaysInput)
		ec2Client, err := createEc2Client(ctx, config.AccessKey, config.SecretKey, config.Region)
		if err != nil {
			return nil, err
		}
		o, err := ec2Client.DescribeNatGateways(ctx, &i)
		if err != nil {
			return nil, err
		}
		return *o, nil
	}
}

/**
 transform functions
	All these functions do not return error if required data is missing just return empty respose struct.
	Calling aws with empty input struct will fail.
**/
func toCreateVpcInput(o lua.Object) ec2.CreateVpcInput {
	input := ec2.CreateVpcInput{}
	cidr := o.GetString("cidr")
	if cidr == "" {
		return input
	}

	input.CidrBlock = aws.String(cidr)
	awsTags := createTags(getObject(o, "tags"))
	if len(awsTags) > 0 {
		input.TagSpecifications = []types.TagSpecification{{
			ResourceType: types.ResourceTypeVpc,
			Tags:         awsTags,
		}}
	}

	return input
}

func fromCreateVpcOutput(o ec2.CreateVpcOutput) lua.Object {
	return toLua(o)
}

func toDescribeVpcsInput(o lua.Object) ec2.DescribeVpcsInput {
	return ec2.DescribeVpcsInput{}
}

func fromDescribeVpcsOutput(o ec2.DescribeVpcsOutput) lua.Object {
	return toLua(o)
}

func toCreateSubnetInput(o lua.Object) ec2.CreateSubnetInput {
	input := ec2.CreateSubnetInput{}

	cidr := o.GetString("cidr")
	vpc_id := o.GetString("vpc_id")
	if cidr == "" || vpc_id == "" {
		return input
	}
	input.VpcId = aws.String(vpc_id)
	input.CidrBlock = aws.String(cidr)

	az := o.GetString("availability_zone_id")
	if az != "" {
		input.AvailabilityZoneId = aws.String(az)
	}

	awsTags := createTags(getObject(o, "tags"))
	if len(awsTags) > 0 {
		input.TagSpecifications = []types.TagSpecification{{
			ResourceType: types.ResourceTypeSubnet,
			Tags:         awsTags,
		}}
	}
	return input
}

func fromCreateSubnetOutput(o ec2.CreateSubnetOutput) lua.Object {
	return toLua(o)
}

func toDescribeSubnetsInput(o lua.Object) ec2.DescribeSubnetsInput {
	input := ec2.DescribeSubnetsInput{}

	awsFilters := createFilters(o.GetList("filters"))
	if len(awsFilters) > 0 {
		input.Filters = awsFilters
	}
	return input
}

func fromDescribeSubnetsOutput(o ec2.DescribeSubnetsOutput) lua.Object {
	return toLua(o)
}

func toDescribeAZsInput(o lua.Object) ec2.DescribeAvailabilityZonesInput {
	return ec2.DescribeAvailabilityZonesInput{}
}

func fromDescribeAZsOutput(o ec2.DescribeAvailabilityZonesOutput) lua.Object {
	return toLua(o)
}

func toCreateIgwInput(o lua.Object) ec2.CreateInternetGatewayInput {
	// TODO impl
	return ec2.CreateInternetGatewayInput{}
}

func fromCreateIgwOutput(o ec2.CreateInternetGatewayOutput) lua.Object {
	return toLua(o)
}

func toDeleteIgwInput(o lua.Object) ec2.DeleteInternetGatewayInput {
	return ec2.DeleteInternetGatewayInput{}
}

func fromDeleteIgwOutput(o ec2.DeleteInternetGatewayOutput) lua.Object {
	return toLua(o)
}

func toDescribeIgwInput(o lua.Object) ec2.DescribeInternetGatewaysInput {
	// TODO
	return ec2.DescribeInternetGatewaysInput{}
}

func fromDescribeIgwOutput(o ec2.DescribeInternetGatewaysOutput) lua.Object {
	return toLua(o)
}

func toCreateNatInput(o lua.Object) ec2.CreateNatGatewayInput {
	// TODO
	return ec2.CreateNatGatewayInput{}
}

func fromCreateNatOutput(o ec2.CreateNatGatewayOutput) lua.Object {
	return toLua(o)
}

func toDeleteNatInput(o lua.Object) ec2.DeleteNatGatewayInput {
	// TODO
	return ec2.DeleteNatGatewayInput{}
}

func fromDeleteNatOutput(o ec2.DeleteNatGatewayOutput) lua.Object {
	return toLua(o)
}

func toDescribeNatInput(o lua.Object) ec2.DescribeNatGatewaysInput {
	// TODO
	return ec2.DescribeNatGatewaysInput{}
}

func fromDescribeNatOutput(o ec2.DescribeNatGatewaysOutput) lua.Object {
	return toLua(o)
}

/**
 Util functions
**/
func getObject(o lua.Object, key string) map[string]string {
	m := o.GetObject(key)
	obj := make(map[string]string)
	for k, v := range m {
		if s, ok := v.(string); ok {
			obj[k] = s
		}
	}
	return obj
}

func createTags(tags map[string]string) []types.Tag {
	awsTags := make([]types.Tag, 0, len(tags))
	for k, v := range tags {
		awsTags = append(awsTags, types.Tag{Key: aws.String(k), Value: aws.String(v)})
	}
	return awsTags
}

func createFilters(filters []interface{}) []types.Filter {
	awsFilters := make([]types.Filter, 0, len(filters))
	for _, filter := range filters {
		o, ok := filter.(lua.Object)
		if !ok {
			continue
		}
		name := o.GetString("Name")
		values := o.GetList("Values")
		svalues := make([]string, 0, len(values))
		for _, v := range values {
			svalues = append(svalues, v.(string))
		}
		awsFilters = append(awsFilters, types.Filter{Name: aws.String(name), Values: svalues})
	}
	return awsFilters
}
