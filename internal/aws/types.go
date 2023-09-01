package aws

import "context"

type ClientType int

const (
	IamClient ClientType = iota
	Ec2Client
)

type OpType int

const (
	// iam op types
	CreateUser OpType = iota
	DeleteUser
	GetUser
	ListUsers
	CreateAccessKeys
	DeleteAccessKeys
	GetAccessKeys
	ListAccessKeys

	// ec2 op types
	CreateSubnet
	DeleteSubnet
	ListSubnets
	CreateVpc
	DeleteVpc
	ListVpcs
	// IGW
	CreateIgw
	DeleteIgw
	ListIgws
	// NAT
	CreateNat
	DeleteNat
	ListNats
	// AZ
	ListAvailabilityZones
)

type opFunc = func(ctx context.Context, input interface{}) (interface{}, error)
