package graph

import (
	"crud_ql/auth"
	"crud_ql/repository"
)

type Resolver struct {
	DB   *repository.DB
	AUTH *auth.CognitoClient
}
