package resolver

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestCreateAuthor(t *testing.T) {
	cases := []struct {
		name         string
		wantErr      bool
		err          error
		wantResponse *gqlmodels.Author
	}{
		{
			name:    "succesfully create author",
			wantErr: false,
			err:     nil,
			wantResponse: &gqlmodels.Author{
				UserName: &testutls.MockUsername,
				Email:    &testutls.MockEmail,
				ID:       "100",
			}},
		{
			name:         "failed creating author",
			wantErr:      true,
			err:          fmt.Errorf("connection error"),
			wantResponse: nil,
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
	r := Resolver{}
	testutls.SetupEnv("../.env.local")

	for _, tt := range cases {
		if tt.name == "succesfully create author" {
			mock.ExpectExec(regexp.
				QuoteMeta("INSERT INTO `authors` (`username`,`email`,`password`,`active`,`author_address`,`last_login`," +
					"`last_password_change`,`token`,`role_id`,`created_at`,`updated_at`,`deleted_at`) " +
					"VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
		} else {
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `authors`")).
				WithArgs().
				WillReturnError(fmt.Errorf("connection error"))
		}
		_, err := r.Mutation().
			CreateAuthor(context.Background(),
				gqlmodels.AuthorCreateInput{UserName: testutls.MockUsername, Email: testutls.MockEmail})
		if err != nil && tt.wantErr == false {
			t.Fatal(err)
		}
		if err == nil && tt.wantErr == true {
			t.Fatal("expected an error")
		}

	}

}
