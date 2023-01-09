package resolver

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestAuthors(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successfully get all authors",
			wantErr: false,
		},
	}
	err := godotenv.Load("../.env.local")
	if err != nil {
		t.Fatal(err)
	}
	resolver1 := Resolver{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	rows := sqlmock.
		NewRows(
			[]string{"id", "username", "email", "roleId", "createdAt", "updatedAt", "deletedAt"},
		).
		AddRow(testutls.MockID, "username", "email", testutls.MockID, "c", "u", "d")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors`;")).WithArgs().WillReturnRows(rows)
	rowCount := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM `authors`;")).
		WithArgs().
		WillReturnRows(rowCount)
	boil.SetDB(db)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := resolver1.Query().Authors(context.Background(), &gqlmodels.AuthorPagination{})
			if err != nil {
				t.Fatal(err)
			}
			if result.Total != 1 {
				t.Fatal("unexpected total")
			}
		})
	}
}

func TestMe(t *testing.T) {

	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successfully get information about yourself",
			wantErr: false,
		},
		{
			name:    "failed getting info about yourself",
			wantErr: true,
		},
	}
	err := godotenv.Load("../.env.local")
	if err != nil {
		t.Fatal(err)
	}
	resolver1 := Resolver{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	boil.SetDB(db)

	for _, tt := range cases {
		if tt.name == "successfully get information about yourself" {
			rows := sqlmock.
				NewRows(
					[]string{"id", "username", "email", "roleId", "createdAt", "updatedAt", "deletedAt"},
				).
				AddRow(testutls.MockID, "username", "email", testutls.MockID, "c", "u", "d")
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (id=?) LIMIT 1")).
				WithArgs().WillReturnRows(rows)

		} else {
			mock.ExpectQuery(regexp.QuoteMeta(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (id=?) LIMIT 1;"))).
				WithArgs().WillReturnError(fmt.Errorf("connection error"))
		}
		_, err := resolver1.Query().Me(context.Background())
		if err != nil && tt.wantErr == false {
			t.Fatal(err)
		}
		if err == nil && tt.wantErr == true {
			t.Fatal(err)
		}

	}

}
