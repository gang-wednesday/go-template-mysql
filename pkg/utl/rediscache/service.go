package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go-template/daos"
	"go-template/models"
	"go-template/pkg/utl/resultwrapper"

	"github.com/go-redis/redis/v8"
)

func SavePostInRedis(rdb *redis.Client, ctx context.Context, id int, p *models.Post) error {
	u, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, fmt.Sprintf("posts%d", id), u, time.Hour*6).Err()
	return err

}

func SaveUserInRedis(rdb *redis.Client, ctx context.Context, id int, a *models.Author) error {
	u, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, fmt.Sprintf("user%d", id), u, time.Hour*6).Err()
	return err
}

func GetAuthorById(rdb *redis.Client, ctx context.Context, id int) (*models.Author, error) {

	log.Println(rdb)
	bytes, err := rdb.Get(ctx, fmt.Sprintf("user%d", id)).Bytes()
	if err != nil {
		log.Println(err)
		if err == redis.Nil {
			a, err := daos.FindAuthorById(id, ctx)
			if err != nil {
				return nil, err
			}
			err = SaveUserInRedis(rdb, ctx, a.ID, a)
			if err != nil {
				log.Println(err)

			}
			return a, nil

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

// func GetAuthorByToken(rdb *redis.Client, ctx context.Context, token string) (*models.Author, error) {

// 	bytes, err := rdb.Get(ctx, fmt.Sprintf("user:Token:%s", token)).Bytes()
// 	if err != nil {
// 		if err == redis.Nil {
// 			return daos.FindAuthorByToken(token, ctx)
// 		}
// 		return nil, err
// 	}
// 	var author models.Author
// 	err = json.Unmarshal(bytes, &author)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &author, nil
// }

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
			post, err := daos.FindPostbyId(id, ctx)
			if err != nil {
				return nil, err
			}
			err = SavePostInRedis(rdb, ctx, post.ID, post)
			if err != nil {
				log.Println(err)
			}
			return post, nil
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
