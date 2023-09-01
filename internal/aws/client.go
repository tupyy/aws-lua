package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ClientConfiguration struct {
	AccessKey string
	SecretKey string
	Region    string
}

// getAwsCredentials returns a CredentialsProviderFunc to be used to create aws config.
func getAwsCredentials(ctx context.Context, accessKey, secretKey string) aws.CredentialsProviderFunc {
	return func(ctx context.Context) (aws.Credentials, error) {
		creds := aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}
		return creds, nil
	}
}
