package daos

import (
	"context"
	"fmt"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestPostById(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "successfully finding post by id",
			wantErr: false,
		},
		{
			name:    "failing finding post by id",
			wantErr: true,
		},
	}
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	for _, tt := range cases {
		if tt.name == "successfully finding post by id" {
			rows := sqlmock.
				NewRows(
					[]string{"id", "title", "content", "authorId", "createdAt", "updatedAt", "deletedAt"},
				).
				AddRow(testutls.MockID, "title", "content", testutls.MockID, testutls.MockAuthor().
					CreatedAt, testutls.MockAuthor().UpdatedAt, testutls.MockAuthor().DeletedAt)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM `posts` WHERE (id=?) LIMIT 1;")).WithArgs().WillReturnRows(rows)
		} else {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `posts`.* FROM" +
				" `posts` WHERE (id=?) LIMIT 1;")).WithArgs().WillReturnError(fmt.Errorf("connection error"))

		}
		_, err := FindPostbyId(testutls.MockID, context.Background())
		if err != nil && tt.wantErr == false {
			t.Fatal(err)
		}

		if err == nil && tt.wantErr == true {
			t.Fatal("expected an error")
		}

	}
}
