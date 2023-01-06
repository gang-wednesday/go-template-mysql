package rediscache

import (
	"context"

	"go-template/daos"
	"go-template/models"
	resultwrapper "go-template/pkg/utl/resultwrapper"
)

// Service ...
type Service interface {
	GetAuthor(id int, ctx context.Context) (models.Author, error)
	GetRole(id int, ctx context.Context) (models.Role, error)
}

// GetUser gets user from redis, if present, else from the database
func GetUser(userID int, ctx context.Context) (*models.Author, error) {
	// get user cache key

	user, err := daos.FindAuthorById(userID, ctx)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	// setting user cache key

	return user, nil
}

// GetRole gets role from redis, if present, else from the database
func GetRole(roleID int, ctx context.Context) (*models.Role, error) {
	// get role cache key

	role, err := daos.FindRoleByID(roleID, ctx)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	// setting role cache key

	return role, nil
}
