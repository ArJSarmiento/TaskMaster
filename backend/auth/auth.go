package auth

import (
	"context"
	"crud_ql/graph/model"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func SignUp(ctx context.Context, cognitoClient CognitoClient, input model.CreateUserInput) (*string, error) {
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

	user, err := cognitoClient.SignUp(ctx, params)
	if err != nil {
		return nil, err
	}

	confirmInput := &cip.AdminConfirmSignUpInput{
		UserPoolId: aws.String(cognitoClient.UserPoolId),
		Username:   aws.String(input.Username),
	}

	// auto confirm all users.
	_, err = cognitoClient.AdminConfirmSignUp(ctx, confirmInput)
	if err != nil {
		return nil, err
	}

	return user.UserSub, err
}

func SignIn(ctx context.Context, cognitoClient CognitoClient, input model.SignInRequest) (*model.SignInResponse, error) {
	signInInput := &cip.AdminInitiateAuthInput{
		AuthFlow:       "ADMIN_USER_PASSWORD_AUTH",
		ClientId:       aws.String(cognitoClient.AppClientId),
		UserPoolId:     aws.String(cognitoClient.UserPoolId),
		AuthParameters: map[string]string{"USERNAME": input.Username, "PASSWORD": input.Password},
	}

	output, err := cognitoClient.AdminInitiateAuth(ctx, signInInput)
	if err != nil {
		return nil, err
	}

	if output.AuthenticationResult == nil {
		return nil, errors.New("unexpected nil in output or output.AuthenticationResult")
	}

	expires_in := strconv.FormatInt(int64(output.AuthenticationResult.ExpiresIn), 10)
	res := model.SignInResponse{
		AccessToken:  output.AuthenticationResult.AccessToken,
		ExpiresIn:    &expires_in,
		IDToken:      output.AuthenticationResult.IdToken,
		RefreshToken: output.AuthenticationResult.RefreshToken,
		TokenType:    output.AuthenticationResult.TokenType,
	}

	return &res, nil
}

func Logout(ctx context.Context, cognitoClient CognitoClient, input model.LogoutRequest) error {
	signOutInput := &cip.GlobalSignOutInput{
		AccessToken: aws.String(input.AccessToken),
	}

	_, err := cognitoClient.GlobalSignOut(ctx, signOutInput)
	if err != nil {
		return err
	}

	return nil
}
