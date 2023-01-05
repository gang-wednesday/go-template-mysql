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

func FindPostbyTitle(title string, ctx context.Context) (models.PostSlice, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Posts(qm.Where(fmt.Sprintf("%s=?", models.PostColumns.Title), title)).All(ctx, contextExecutor)
}

func FindPostByAuthor(authorId int, ctx context.Context) (models.PostSlice, error) {
	return models.PostSlice{}, nil
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
