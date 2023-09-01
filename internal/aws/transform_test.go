package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	. "github.com/onsi/gomega"
)

type b struct {
	BFoo string
}
type a struct {
	Foo     *bool
	Bar     string
	MyArray []string
	B       b
	BB      []b
}

func TestFrom(t *testing.T) {
	RegisterTestingT(t)
	aa := a{aws.Bool(true), "test", []string{"1", "2"}, b{"bfoo"}, []b{
		{"test"}, {"test2"},
	}}

	o := toLua(aa)
	fmt.Printf("%+v\n", o)
	Expect(o.GetBool("Foo")).To(BeTrue())
	Expect(o["Bar"]).To(Equal("test"))

	arr, ok := o["MyArray"].([]interface{})
	Expect(ok).To(BeTrue())
	Expect(len(arr)).To(Equal(2))
}

func TestFrom2(t *testing.T) {
	RegisterTestingT(t)
	a := ec2.CreateSubnetOutput{
		Subnet: &types.Subnet{
			AssignIpv6AddressOnCreation: aws.Bool(false),
			AvailabilityZone:            aws.String("az"),
			CidrBlock:                   aws.String("10.0.0.0/24"),
			SubnetId:                    aws.String("owner"),
			Ipv6CidrBlockAssociationSet: []types.SubnetIpv6CidrBlockAssociation{
				{
					AssociationId: aws.String("test"),
					Ipv6CidrBlock: aws.String("Ipv6CidrBlock"),
					Ipv6CidrBlockState: &types.SubnetCidrBlockState{
						State: types.SubnetCidrBlockStateCodeAssociated,
					},
				},
			},
			Tags: []types.Tag{
				{
					Key:   aws.String("tag1"),
					Value: aws.String("value2"),
				},
			},
		},
	}
	_ = toLua(a)
}

func TestFrom3(t *testing.T) {
	RegisterTestingT(t)
	a := ec2.DescribeAvailabilityZonesOutput{
		AvailabilityZones: []types.AvailabilityZone{
			{
				GroupName:          aws.String("group1"),
				NetworkBorderGroup: aws.String("network-group1"),
				ZoneName:           aws.String("zone1"),
			},
			{
				GroupName:          aws.String("group2"),
				NetworkBorderGroup: aws.String("network-group2"),
				ZoneName:           aws.String("zone2"),
			},
			{
				GroupName:          aws.String("group2"),
				NetworkBorderGroup: aws.String("network-group2"),
				ZoneName:           aws.String("zone2"),
			},
		},
	}

	o := toLua(a)
	fmt.Printf("%+v\n", o)
}
