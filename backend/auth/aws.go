package auth

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	AppClientId string
	UserPoolId  string
	*cip.Client
}

// Init initializes a new instance of the CognitoClient.
//
// It loads the shared AWS configuration and returns a pointer to a CognitoClient
// with the provided app client ID, user pool ID, and AWS SDK configuration.
//
// Returns:
// - *CognitoClient: A pointer to the initialized CognitoClient.
func Init() *CognitoClient {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		os.Getenv("COGNITO_APP_CLIENT_ID"),
		os.Getenv("COGNITO_USER_POOL_ID"),
		cip.NewFromConfig(cfg),
	}
}
