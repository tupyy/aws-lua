// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the instance profiles that have the specified associated IAM role. If
// there are none, the operation returns an empty list. For more information about
// instance profiles, go to Using instance profiles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html)
// in the IAM User Guide. You can paginate the results using the MaxItems and
// Marker parameters.
func (c *Client) ListInstanceProfilesForRole(ctx context.Context, params *ListInstanceProfilesForRoleInput, optFns ...func(*Options)) (*ListInstanceProfilesForRoleOutput, error) {
	if params == nil {
		params = &ListInstanceProfilesForRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceProfilesForRole", params, optFns, c.addOperationListInstanceProfilesForRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceProfilesForRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceProfilesForRoleInput struct {

	// The name of the role to list instance profiles for. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	noSmithyDocumentSerde
}

// Contains the response to a successful ListInstanceProfilesForRole request.
type ListInstanceProfilesForRoleOutput struct {

	// A list of instance profiles.
	//
	// This member is required.
	InstanceProfiles []types.InstanceProfile

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstanceProfilesForRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListInstanceProfilesForRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListInstanceProfilesForRole{}, middleware.After)
	if err != nil {
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
	if err = addOpListInstanceProfilesForRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceProfilesForRole(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListInstanceProfilesForRoleAPIClient is a client that implements the
// ListInstanceProfilesForRole operation.
type ListInstanceProfilesForRoleAPIClient interface {
	ListInstanceProfilesForRole(context.Context, *ListInstanceProfilesForRoleInput, ...func(*Options)) (*ListInstanceProfilesForRoleOutput, error)
}

var _ ListInstanceProfilesForRoleAPIClient = (*Client)(nil)

// ListInstanceProfilesForRolePaginatorOptions is the paginator options for
// ListInstanceProfilesForRole
type ListInstanceProfilesForRolePaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceProfilesForRolePaginator is a paginator for
// ListInstanceProfilesForRole
type ListInstanceProfilesForRolePaginator struct {
	options   ListInstanceProfilesForRolePaginatorOptions
	client    ListInstanceProfilesForRoleAPIClient
	params    *ListInstanceProfilesForRoleInput
	nextToken *string
	firstPage bool
}

// NewListInstanceProfilesForRolePaginator returns a new
// ListInstanceProfilesForRolePaginator
func NewListInstanceProfilesForRolePaginator(client ListInstanceProfilesForRoleAPIClient, params *ListInstanceProfilesForRoleInput, optFns ...func(*ListInstanceProfilesForRolePaginatorOptions)) *ListInstanceProfilesForRolePaginator {
	if params == nil {
		params = &ListInstanceProfilesForRoleInput{}
	}

	options := ListInstanceProfilesForRolePaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstanceProfilesForRolePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceProfilesForRolePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstanceProfilesForRole page.
func (p *ListInstanceProfilesForRolePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceProfilesForRoleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListInstanceProfilesForRole(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInstanceProfilesForRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListInstanceProfilesForRole",
	}
}
