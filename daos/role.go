package daos

import (
	"context"
	"go-template/models"
)

func FindRoleByID(role int, ctx context.Context) (*models.Role, error) {
	return &models.Role{}, nil
}
