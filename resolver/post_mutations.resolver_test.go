package resolver

import (
	"context"
	"go-template/gqlmodels"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)
	mock.ExpectExec(regexp.
		QuoteMeta("INSERT INTO `posts` (`title`,`content`,`author_id`," +
			"`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))
	resolver1 := Resolver{}
	t.Run("successfull role insert", func(t *testing.T) {
		mockPost := gqlmodels.PostCreateInput{
			Title:   &testutls.MockPost().Title,
			Content: &testutls.MockPost().Content,
		}
		_, err := resolver1.Mutation().CreatePost(context.Background(), mockPost)
		if err != nil {
			t.Fatal(err)
		}
	})

}
