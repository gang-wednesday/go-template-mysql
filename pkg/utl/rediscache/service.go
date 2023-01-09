package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-template/daos"
	"go-template/models"
	"go-template/pkg/utl/resultwrapper"

	"github.com/go-redis/redis/v8"
)

func GetAuthorById(rdb *redis.Client, ctx context.Context, id int) (*models.Author, error) {

	log.Println(rdb)
	bytes, err := rdb.Get(ctx, fmt.Sprintf("user%d", id)).Bytes()
	if err != nil {
		log.Println(err)
		if err == redis.Nil {
			return daos.FindAuthorById(id, ctx)
		}
		return nil, err
	}
	var author models.Author
	err = json.Unmarshal(bytes, &author)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func GetAuthorByToken(rdb *redis.Client, ctx context.Context, token string) (*models.Author, error) {

	bytes, err := rdb.Get(ctx, fmt.Sprintf("user:Token:%s", token)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return daos.FindAuthorByToken(token, ctx)
		}
		return nil, err
	}
	var author models.Author
	err = json.Unmarshal(bytes, &author)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func GetRole(roleID int, ctx context.Context) (*models.Role, error) {
	// get role cache key

	role, err := daos.FindRoleByID(roleID, ctx)
	if err != nil {
		return nil, resultwrapper.ResolverSQLError(err, "data")
	}
	// setting role cache key

	return role, nil
}

func PostById(rdb *redis.Client, ctx context.Context, id int) (*models.Post, error) {

	bytes, err := rdb.Get(ctx, fmt.Sprintf("posts%d", id)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return daos.FindPostbyId(id, ctx)
		}
		return nil, err
	}
	var post models.Post
	err = json.Unmarshal(bytes, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
