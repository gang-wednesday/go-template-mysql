package resolver

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/internal/middleware/auth"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// did not need to use monkey patching
func TestCreatePost(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successfull post insert",
			wantErr: false,
		}, {
			name:    "failed post insert",
			wantErr: true,
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)
	defer func() {
		db.Close()
	}()
	resolver1 := Resolver{}
	for _, tt := range cases {
		if tt.name == "successfull post insert" {
			mock.ExpectExec(regexp.
				QuoteMeta("INSERT INTO `posts` (`title`,`content`,`author_id`," +
					"`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)")).
				WithArgs().
				WillReturnResult(sqlmock.NewResult(1, 1))
		} else {
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts`")).
				WithArgs().
				WillReturnError(fmt.Errorf("connection error :could not insert into database"))
		}
		t.Run(tt.name, func(t *testing.T) {
			mockPost := gqlmodels.PostCreateInput{
				Title:   &testutls.MockPost().Title,
				Content: &testutls.MockPost().Content,
			}
			_, err := resolver1.Mutation().CreatePost(context.Background(), mockPost)
			if err != nil {
				if !tt.wantErr {
					t.Fatal(err)
				}
			}
		})

	}

}

func TestUpdatePost(t *testing.T) {
	cases := []struct {
		name        string
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "Succesfully post update",
			wantErr:     false,
			expectedErr: "",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)
	defer func() {
		db.Close()
	}()
	//mocking select query to get row
	rows := sqlmock.
		NewRows(
			[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
		).
		AddRow(testutls.MockID, "title", "content", testutls.MockID, "c", "u", "d")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM `posts` WHERE (id=?) LIMIT 1;")).WithArgs().WillReturnRows(rows)
	//mocking update query to update row
	mock.ExpectExec(regexp.
		QuoteMeta("UPDATE `posts` SET `title`=?,`content`=?,`author_id`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))
	r := Resolver{}
	for _, tt := range cases {

		//setting the user details in the context
		c := context.WithValue(context.Background(), auth.UserCtxKey, testutls.MockAuthor())

		//checking the updatePost function
		t.Run(tt.name, func(t *testing.T) {
			_, err := r.Mutation().
				UpdatePost(c, gqlmodels.PostUpdateInput{Title: &testutls.MockPost().
					Title, Content: &testutls.MockPost().Title, ID: fmt.Sprint(testutls.MockID)})
			if err != nil {
				if tt.wantErr == false {
					t.Fatal(err)
				}
			}
			if tt.wantErr == true {
				t.Fatal("expected an error")
			}
		})
	}

}

func TestDeletePost(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "succesfull delte",
			wantErr: false,
		},
	}
	for _, tt := range cases {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		boil.SetDB(db)
		defer func() {
			db.Close()
		}()
		rows := sqlmock.
			NewRows(
				[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
			).
			AddRow(testutls.MockID, "title", "content", testutls.MockID, "c", "u", "d")
		mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM `posts` WHERE (id=?) LIMIT 1;")).WithArgs().WillReturnRows(rows)
		mock.ExpectExec(regexp.
			QuoteMeta(regexp.QuoteMeta("DELETE FROM `posts` WHERE `id`="))).
			WithArgs().
			WillReturnResult(sqlmock.NewResult(1, 1))
		r := Resolver{}
		t.Run(tt.name, func(t *testing.T) {
			_, err := r.Mutation().DeletePost(context.Background(), gqlmodels.PostDeleteInput{ID: "200"})
			if err != nil && tt.wantErr == false {
				if tt.wantErr == false {
					t.Fatal(err)
				}

			}
		})
	}

}
