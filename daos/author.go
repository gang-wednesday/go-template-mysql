package daos

import (
	"context"
	"go-template/models"
)

func FindAuthorByEmail(email string, ctx context.Context) (models.Author, error) {
	return models.Author{}, nil
}
func FindAuthorById(id int, ctx context.Context) (*models.Author, error) {
	return &models.Author{}, nil
}

func GetAuthor(id int, ctx context.Context) (models.Author, error) {
	return models.Author{}, nil
}
