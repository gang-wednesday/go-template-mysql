package resolver

import (
	"context"
	"go-template/gqlmodels"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestPosts(t *testing.T) {

	cases := []struct {
		name     string
		wantErr  bool
		wantResp int
	}{
		{
			name:     "succesfully get all posts",
			wantErr:  false,
			wantResp: 1,
		},
	}

	err := godotenv.Load("../.env.local")
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range cases {

		resolver1 := Resolver{}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		rows := sqlmock.
			NewRows(
				[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
			).
			AddRow(testutls.MockID, "title", "content", testutls.MockID, "c", "u", "d")
		mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM `posts`;")).WithArgs().WillReturnRows(rows)
		rowCount := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `posts`;")).
			WithArgs().
			WillReturnRows(rowCount)
		boil.SetDB(db)
		t.Run(tt.name, func(t *testing.T) {
			response, err := resolver1.Query().Posts(context.Background(), &gqlmodels.PostPagination{})
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.wantResp, len(response.Posts))
		})

	}

}

func TestPostsById(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succesfully get all post by id",
			wantErr: false,
		},
	}

	err := godotenv.Load("../.env.local")
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range cases {
		resolver1 := Resolver{}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		boil.SetDB(db)

		rows := sqlmock.
			NewRows(
				[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
			).
			AddRow(testutls.MockID, "title", "content", testutls.MockID, "c", "u", "d")
		mock.ExpectQuery("SELECT `posts`.* FROM `posts` WHERE (id=?) LIMIT 1;").WithArgs().WillReturnRows(rows)
		t.Run(tt.name, func(t *testing.T) {
			_, err := resolver1.Query().PostByID(context.Background(), string(testutls.MockStringId))
			if err != nil {
				t.Fatal(err)
			}

		})
	}

}
