package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-template/testutls"
	"regexp"
	"testing"

	redisutil "go-template/pkg/utl/redisUtil"

	redismock "github.com/go-redis/redismock/v8"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-redis/redis/v8"

	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
)

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
