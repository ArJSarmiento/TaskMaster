package auth

import (
	"context"
	"crud_ql/graph/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func SignUp(ctx context.Context, cognitoClient CognitoClient, input model.CreateUserInput) error {
	// Build a signup request
	params := &cip.SignUpInput{
		ClientId: aws.String(cognitoClient.AppClientId),
		Username: aws.String(input.Username),
		Password: aws.String(input.Password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(input.Email),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(input.Phone),
			},
		},
	}

	_, err := cognitoClient.SignUp(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
