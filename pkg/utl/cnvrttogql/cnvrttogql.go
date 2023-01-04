package cnvrttogql

import (
	"context"
	graphql "go-template/gqlmodels"
	"go-template/internal/constants"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// UsersToGraphQlUsers converts array of type models.Author into array of pointer type graphql.Author
func UsersToGraphQlUsers(u models.AuthorSlice, count int) []*graphql.Author {
	var r []*graphql.Author
	for _, e := range u {
		r = append(r, UserToGraphQlUser(e, count))
	}
	return r
}

// UserToGraphQlUser converts type models.Author into pointer type graphql.Author
func UserToGraphQlUser(u *models.Author, count int) *graphql.Author {
	count++
	if u == nil {
		return nil
	}
	var role *models.Role
	if count <= constants.MaxDepth {
		u.L.LoadRole(context.Background(), boil.GetContextDB(), true, u, nil) //nolint:errcheck
		if u.R != nil {
			role = u.R.Role
		}
	}

	return &graphql.Author{
		ID:       strconv.Itoa(u.ID),
		UserName: convert.NullDotStringToPointerString(u.Username),

		Email: convert.NullDotStringToPointerString(u.Email),

		Address: convert.NullDotStringToPointerString(u.AuthorAddress),
		Active:  convert.NullDotBoolToPointerBool(u.Active),
		Role:    RoleToGraphqlRole(role, count),
	}
}

func RoleToGraphqlRole(r *models.Role, count int) *graphql.Role {
	count++
	if r == nil {
		return nil
	}
	var users models.AuthorSlice
	if count <= constants.MaxDepth {
		r.L.LoadAuthors(context.Background(), boil.GetContextDB(), true, r, nil) //nolint:errcheck
		if r.R != nil {
			users = r.R.Authors
		}
	}

	return &graphql.Role{
		ID:          strconv.Itoa(r.ID),
		AccessLevel: r.AccessLevel,
		Name:        r.Name,
		UpdatedAt:   convert.NullDotTimeToPointerInt(r.UpdatedAt),
		CreatedAt:   convert.NullDotTimeToPointerInt(r.CreatedAt),
		DeletedAt:   convert.NullDotTimeToPointerInt(r.DeletedAt),
		Authors:     UsersToGraphQlUsers(users, count),
	}
}
