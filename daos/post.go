package daos

import (
	"context"
	"database/sql"
	"fmt"
	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindPostbyId(id int, ctx context.Context) (*models.Post, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Posts(qm.Where(fmt.Sprintf("%s=?", models.PostColumns.ID), id)).One(ctx, contextExecutor)
}

func FindPostbyTitle(title string, ctx context.Context) (*models.Post, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Posts(qm.Where(fmt.Sprintf("%s=?", models.PostColumns.Title), title)).One(ctx, contextExecutor)
}

func FindPostByAuthorId(authorId int, ctx context.Context) (models.PostSlice, int64, error) {
	contextExecutor := getContextExecutor(nil)
	posts, err := models.Posts(qm.Where(fmt.Sprintf("%s=?", models.PostColumns.AuthorID), authorId)).All(ctx, contextExecutor)
	if err != nil {
		return nil, 0, err
	}

	count, err := models.Authors(qm.Where(fmt.Sprintf("%s=?", models.PostColumns.AuthorID), authorId)).Count(ctx, contextExecutor)

	return posts, count, err
}

func FindPosts(queryMods []qm.QueryMod, ctx context.Context) (models.PostSlice, int64, error) {
	contextExecutor := getContextExecutor(nil)
	posts, err := models.Posts(queryMods...).All(ctx, contextExecutor)
	if err != nil {
		return models.PostSlice{}, 0, err
	}
	queryMods = append(queryMods, qm.Offset(0))
	count, err := models.Authors(queryMods...).Count(ctx, contextExecutor)

	return posts, count, err
}

func CreatePost(post models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
	contextExecutor := getContextExecutor(tx)
	err := post.Insert(ctx, contextExecutor, boil.Infer())
	return post, err
}

func Updatepost(post models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
	contextExecutor := getContextExecutor(tx)
	_, err := post.Update(ctx, contextExecutor, boil.Infer())
	return post, err
}

func DeletePost(post models.Post, ctx context.Context) (int64, error) {
	contextExecutor := getContextExecutor(nil)
	return post.Delete(ctx, contextExecutor)
}
