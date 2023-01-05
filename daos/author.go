package daos

import (
	"context"
	"database/sql"
	"fmt"
	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindAuthorByEmail(email string, ctx context.Context) (*models.Author, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Authors(qm.Where(fmt.Sprintf("%s=?", models.AuthorColumns.Token), email)).One(ctx, contextExecutor)
}
func FindAuthorById(id int, ctx context.Context) (*models.Author, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Authors(qm.Where(fmt.Sprintf("%s=?", models.AuthorColumns.ID), id)).One(ctx, contextExecutor)
}

func FindAuthorByToken(token string, ctx context.Context) (*models.Author, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Authors(qm.Where(fmt.Sprintf("%s=?", models.AuthorColumns.Token), token)).One(ctx, contextExecutor)
}
func FindAuthorByUsername(username string, ctx context.Context) (*models.Author, error) {
	contextExecutor := getContextExecutor(nil)
	return models.Authors(qm.Where(fmt.Sprintf("%s=?", models.AuthorColumns.Username), username)).One(ctx, contextExecutor)
}

func CreateAuthorTx(author models.Author, ctx context.Context, tx *sql.Tx) (*models.Author, error) {
	contextExecutor := getContextExecutor(tx)

	err := author.Insert(ctx, contextExecutor, boil.Infer())
	return &author, err
}

func UpdateAuthorTx(author models.Author, ctx context.Context, tx *sql.Tx) (models.Author, error) {
	contextExecutor := getContextExecutor(tx)
	_, err := author.Update(ctx, contextExecutor, boil.Infer())

	return author, err
}

func DeleteAuthor(author models.Author, ctx context.Context) (int64, error) {
	contextExecutor := getContextExecutor(nil)

	return author.Delete(ctx, contextExecutor)
}

func FindAllAuthorWithCount(queryMods []qm.QueryMod, ctx context.Context) (models.AuthorSlice, int64, error) {
	contextExecutor := getContextExecutor(nil)
	authors, err := models.Authors(queryMods...).All(ctx, contextExecutor)
	if err != nil {
		return models.AuthorSlice{}, 0, err
	}
	return authors, int64(len(authors)), nil

}

func GetAuthor(id int, ctx context.Context) (models.Author, error) {
	return models.Author{}, nil
}
