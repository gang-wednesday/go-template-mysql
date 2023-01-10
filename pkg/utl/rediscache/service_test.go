package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-template/testutls"
	"regexp"
	"testing"
	"time"

	redisutil "go-template/pkg/utl/redisUtil"

	redismock "github.com/go-redis/redismock/v8"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-redis/redis/v8"

	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestSavePostInRedis(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succesfully cache post in redis",
			wantErr: false,
		},
	}

	for _, tt := range cases {
		redisdb, rdmock := redismock.NewClientMock()
		p, _ := json.Marshal(testutls.MockPost())
		rdmock.ExpectSet(fmt.Sprintf("posts%d", testutls.MockID), p, time.Hour*6).SetVal("31")

		t.Run(tt.name, func(t *testing.T) {
			err := SavePostInRedis(redisdb, context.Background(), testutls.MockID, testutls.MockPost())
			if err != nil && tt.wantErr == false {
				t.Fatal(err)
			}
		})
	}
}

func TestSaveAuthorInRedis(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succesfully cache author in redis",
			wantErr: false,
		},
	}

	for _, tt := range cases {
		redisdb, rdmock := redismock.NewClientMock()
		p, _ := json.Marshal(testutls.MockAuthor())
		rdmock.ExpectSet(fmt.Sprintf("user%d", testutls.MockID), p, time.Hour*6).SetVal("31")

		t.Run(tt.name, func(t *testing.T) {
			err := SaveAuthorInRedis(redisdb, context.Background(), testutls.MockID, testutls.MockAuthor())
			if err != nil && tt.wantErr == false {
				t.Fatal(err)
			}
		})
	}
}

func TestGetAuthorById(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succefully getting author from cache",
			wantErr: false,
		},
		//, {
		// 	name: "succefully getting author with cache miss",
		// },
	}

	a, err := json.Marshal(testutls.MockAuthor())
	if err != nil {
		t.Fatal(err)
	}

	err = godotenv.Load("./../../../../gtmsql/.env.local")
	if err != nil {
		t.Fatal(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)
	redisdb, rdmock := redismock.NewClientMock()

	for _, tt := range cases {

		if tt.name == "succefully getting author from cache" {
			//mocking redis
			rdmock.ExpectGet(fmt.Sprintf("user%d", testutls.MockID)).SetVal(string(a))
		} else {
			// mocking redis for failure
			rdmock.ExpectGet(fmt.Sprintf("user%d", testutls.MockID)).SetErr(redis.Nil)
			// mocking sql database
			rows := sqlmock.
				NewRows(
					[]string{"id", "username", "email", "roleId", "createdAt", "updatedAt", "deletedAt"},
				).
				AddRow(testutls.MockID, testutls.MockUsername, testutls.MockEmail, testutls.MockID, "c", "u", "d")
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (id=?) LIMIT 1")).
				WithArgs().WillReturnRows(rows)
		}
		//appying go monkey patching to redisutil's get client to return my own custom mock client
		gomonkey.ApplyFunc(redisutil.GetClient, func() *redis.Client {

			rdmock.ExpectGet(fmt.Sprintf("user%d", testutls.MockID)).SetVal(string(a))
			return redisdb
		})
		//funtion that calls redisutil.GetClient
		ath, err := GetAuthorById(redisdb, context.Background(), testutls.MockID)
		if err != nil {
			t.Fatal(err)
		}
		if ath == nil {
			t.Fatal("unexpected value of author")
		}
	}

}

func TestPostByID(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succefully getting author from cache",
			wantErr: false,
		},
		//, {
		// 	name: "succefully getting author with cache miss",
		// },
	}

	a, err := json.Marshal(testutls.MockPost())
	if err != nil {
		t.Fatal(err)
	}

	err = godotenv.Load("./../../../../gtmsql/.env.local")
	if err != nil {
		t.Fatal(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)
	redisdb, rdmock := redismock.NewClientMock()

	for _, tt := range cases {

		if tt.name == "succefully getting author from cache" {
			//mocking redis
			rdmock.ExpectGet(fmt.Sprintf("posts%d", testutls.MockID)).SetVal(string(a))
		} else {
			// mocking redis for failure
			rdmock.ExpectGet(fmt.Sprintf("posts%d", testutls.MockID)).SetErr(redis.Nil)
			// mocking sql database
			rows := sqlmock.
				NewRows(
					[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
				).
				AddRow(testutls.MockID, testutls.MockTitle, testutls.MockContent, testutls.MockID, "c", "u", "d")
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM `posts` WHERE (id=?) LIMIT 1")).
				WithArgs().WillReturnRows(rows)
		}
		//appying go monkey patching to redisutil's get client to return my own custom mock client

		//funtion that calls redisutil.GetClient
		ath, err := PostById(redisdb, context.Background(), testutls.MockID)
		if err != nil {
			t.Fatal(err)
		}
		if ath == nil {
			t.Fatal("unexpected value of post")
		}
	}
}
