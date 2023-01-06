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
func AuthorssToGraphQlAuthors(u models.AuthorSlice, count int) []*graphql.Author {
	var r []*graphql.Author
	for _, e := range u {
		r = append(r, AuthorToGraphQlAuthor(e, count))
	}
	return r
}

// UserToGraphQlUser converts type models.Author into pointer type graphql.Author
func AuthorToGraphQlAuthor(u *models.Author, count int) *graphql.Author {
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
	var posts []*models.Post
	if count <= constants.MaxDepth {
		u.L.LoadPosts(context.Background(), boil.GetContextDB(), true, u, nil) //nolint:errcheck
		if u.R != nil {
			posts = u.R.Posts
		}
	}
	return &graphql.Author{
		ID:        strconv.Itoa(u.ID),
		UserName:  convert.NullDotStringToPointerString(u.Username),
		Active:    convert.NullDotBoolToPointerBool(u.Active),
		Email:     convert.NullDotStringToPointerString(u.Email),
		Token:     convert.NullDotStringToPointerString(u.Token),
		CreatedAt: convert.NullDotTimeToPointerInt(u.CreatedAt),
		UpdatedAt: convert.NullDotTimeToPointerInt(u.UpdatedAt),
		DeletedAt: convert.NullDotTimeToPointerInt(u.DeletedAt),
		Address:   convert.NullDotStringToPointerString(u.AuthorAddress),
		Posts:     PostsToGraphqlPosts(posts, count),
		Role:      RoleToGraphqlRole(role, count),
	}
}

func PostsToGraphqlPosts(a models.PostSlice, count int) []*graphql.Post {
	var r []*graphql.Post
	for _, e := range a {
		r = append(r, PostToGraphQlPost(e, count))
	}
	return r
}

func PostToGraphQlPost(p *models.Post, count int) *graphql.Post {
	count++
	if p == nil {
		return nil
	}
	var author *models.Author
	if count <= constants.MaxDepth {
		p.L.LoadAuthor(context.Background(), boil.GetContextDB(), true, p, nil) //nolint:errcheck
		if p.R != nil {
			author = p.R.Author
		}
	}
	return &graphql.Post{
		ID:        strconv.Itoa(p.ID),
		Title:     convert.StringToPointerString(p.Title),
		Content:   convert.StringToPointerString(p.Content),
		CreatedAt: convert.NullDotTimeToPointerInt(p.CreatedAt),
		UpdatedAt: convert.NullDotTimeToPointerInt(p.UpdatedAt),
		DeletedAt: convert.NullDotTimeToPointerInt(p.DeletedAt),
		Author:    AuthorToGraphQlAuthor(author, count),
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
		Authors:     AuthorssToGraphQlAuthors(users, count),
	}
}
