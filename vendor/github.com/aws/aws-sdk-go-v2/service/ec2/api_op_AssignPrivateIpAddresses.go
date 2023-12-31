// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns one or more secondary private IP addresses to the specified network
// interface. You can specify one or more specific secondary IP addresses, or you
// can specify the number of secondary IP addresses to be automatically assigned
// within the subnet's CIDR block range. The number of secondary IP addresses that
// you can assign to an instance varies by instance type. For information about
// instance types, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information about
// Elastic IP addresses, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide. When you move a secondary
// private IP address to another network interface, any Elastic IP address that is
// associated with the IP address is also moved. Remapping an IP address is an
// asynchronous operation. When you move an IP address from one network interface
// to another, check network/interfaces/macs/mac/local-ipv4s in the instance
// metadata to confirm that the remapping is complete. You must specify either the
// IP addresses or the IP address count in the request. You can optionally use
// Prefix Delegation on the network interface. You must specify either the IPv4
// Prefix Delegation prefixes, or the IPv4 Prefix Delegation count. For
// information, see Assigning prefixes to Amazon EC2 network interfaces (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) AssignPrivateIpAddresses(ctx context.Context, params *AssignPrivateIpAddressesInput, optFns ...func(*Options)) (*AssignPrivateIpAddressesOutput, error) {
	if params == nil {
		params = &AssignPrivateIpAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssignPrivateIpAddresses", params, optFns, c.addOperationAssignPrivateIpAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssignPrivateIpAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for AssignPrivateIpAddresses.
type AssignPrivateIpAddressesInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Indicates whether to allow an IP address that is already assigned to another
	// network interface or instance to be reassigned to the specified network
	// interface.
	AllowReassignment *bool

	// The number of IPv4 prefixes that Amazon Web Services automatically assigns to
	// the network interface. You cannot use this option if you use the Ipv4 Prefixes
	// option.
	Ipv4PrefixCount *int32

	// One or more IPv4 prefixes assigned to the network interface. You cannot use
	// this option if you use the Ipv4PrefixCount option.
	Ipv4Prefixes []string

	// The IP addresses to be assigned as a secondary private IP address to the
	// network interface. You can't specify this parameter when also specifying a
	// number of secondary IP addresses. If you don't specify an IP address, Amazon EC2
	// automatically selects an IP address within the subnet range.
	PrivateIpAddresses []string

	// The number of secondary IP addresses to assign to the network interface. You
	// can't specify this parameter when also specifying private IP addresses.
	SecondaryPrivateIpAddressCount *int32

	noSmithyDocumentSerde
}

type AssignPrivateIpAddressesOutput struct {

	// The IPv4 prefixes that are assigned to the network interface.
	AssignedIpv4Prefixes []types.Ipv4PrefixSpecification

	// The private IP addresses assigned to the network interface.
	AssignedPrivateIpAddresses []types.AssignedPrivateIpAddress

	// The ID of the network interface.
	NetworkInterfaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssignPrivateIpAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssignPrivateIpAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssignPrivateIpAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addAssignPrivateIpAddressesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssignPrivateIpAddressesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssignPrivateIpAddresses(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssignPrivateIpAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssignPrivateIpAddresses",
	}
}

type opAssignPrivateIpAddressesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAssignPrivateIpAddressesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAssignPrivateIpAddressesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addAssignPrivateIpAddressesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAssignPrivateIpAddressesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
