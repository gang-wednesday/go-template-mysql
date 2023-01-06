package daos

import (
	"context"
	"fmt"
	"go-template/models"
	"go-template/testutls"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestCreateAuthor(t *testing.T) {
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `authors`")).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		name string
		args models.Author
		err  error
	}{
		{
			name: "Passing user type value",
			args: models.Author{
				ID:       testutls.MockAuthor().ID,
				Email:    testutls.MockAuthor().Email,
				Password: testutls.MockAuthor().Password,
			},
			err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.Email)
			_, err := CreateAuthor(tt.args, context.Background())
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestUpdateAuthor(t *testing.T) {
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `authors` SET `username`=?,`email`=?,`password`=?,`active`=?," +
		"`author_address`=?,`last_login`=?,`last_password_change`=?,`token`=?,`role_id`=?,`updated_at`=?" +
		",`deleted_at`=? WHERE `id`=?")).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))
	cases := []struct {
		name string
		args models.Author
		err  error
	}{
		{
			name: "Passing user type value",
			args: models.Author{ID: testutls.MockID},
			err:  nil},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UpdateAuthor(tt.args, context.Background())
			assert.Equal(t, err, tt.err)
		})

	}
}

func TestDeleteAuthor(t *testing.T) {
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `authors` WHERE `id`=?")).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	cases := []struct {
		name string
		args models.Author
		err  error
	}{
		{
			name: "Passing user type value",
			args: models.Author{ID: testutls.MockID},
			err:  nil},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeleteAuthor(tt.args, context.Background())
			assert.Equal(t, err, tt.err)
		})

	}

}

func TestFindAuthorbyId(t *testing.T) {
	cases := []struct {
		name string
		args int
		err  error
	}{
		{
			name: "Passing a user_id",
			args: 1,
			err:  nil,
		},
	}

	for _, tt := range cases {
		mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
		oldDB := boil.GetDB()
		defer func() {
			db.Close()
			boil.SetDB(oldDB)
		}()
		boil.SetDB(db)

		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (id=?) LIMIT 1;")).
			WithArgs().
			WillReturnRows(rows)

		t.Run(tt.name, func(t *testing.T) {
			_, err := FindAuthorById(tt.args, context.Background())
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestFindAuthorByEmail(t *testing.T) {
	type args struct {
		email string
	}
	cases := []struct {
		name string
		req  args
		err  error
	}{
		// {
		// 	name: "Fail on finding user",
		// 	req:  args{email: "abc"},
		// 	err:  fmt.Errorf("sql: no rows in sql"),
		// },
		{
			name: "Passing an email",
			req:  args{email: "l@gmail.com"},
			err:  nil,
		},
	}
	for _, tt := range cases {
		mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
		oldDB := boil.GetDB()
		defer func() {
			db.Close()
			boil.SetDB(oldDB)
		}()
		boil.SetDB(db)

		if tt.name == "Fail on finding user" {
			mock.ExpectQuery(regexp.QuoteMeta(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (email=?) LIMIT 1;"))).
				WithArgs().
				WillReturnError(fmt.Errorf("sql: no rows in sql")).WillReturnRows()
		} else {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (email=?) LIMIT 1;")).
				WithArgs().
				WillReturnRows(rows)
		}

		t.Run(tt.name, func(t *testing.T) {
			_, err := FindAuthorByEmail(tt.req.email, context.Background())

			assert.Equal(t, err, tt.err)

		})
	}

}

func TestFindAuthorByToken(t *testing.T) {
	cases := []struct {
		name string
		args string
		err  error
	}{
		{
			name: "Passing a token",
			args: testutls.MockToken,
			err:  nil,
		},
	}
	for _, tt := range cases {
		mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{EnvFileLocation: "../.env.local"})
		oldDB := boil.GetDB()
		defer func() {
			db.Close()
			boil.SetDB(oldDB)
		}()
		boil.SetDB(db)
		rows := sqlmock.NewRows([]string{"token"}).AddRow(1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors` WHERE (token=?) LIMIT 1;")).
			WithArgs().WillReturnRows(rows)

		t.Run(tt.name, func(t *testing.T) {
			_, err := FindAuthorByToken(tt.args, context.Background())
			assert.Equal(t, err, tt.err)

		})

	}
}

func TestFindAllAuthorsWithCount(t *testing.T) {

	oldDB := boil.GetDB()
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
	defer func() {
		boil.SetDB(oldDB)
		db.Close()
	}()

	cases := []struct {
		name      string
		err       error
		dbQueries []testutls.QueryData
	}{
		{
			name: "Failed to find any author",
			err:  fmt.Errorf("sql: no rows in sql"),
		},
		{
			name: "Successfully find all authors with count",
			err:  nil,
			dbQueries: []testutls.QueryData{
				{
					Query: regexp.QuoteMeta("SELECT `authors`.* FROM `authors`;"),
					DbResponse: sqlmock.NewRows([]string{"id", "email", "token"}).AddRow(
						testutls.MockID,
						testutls.MockEmail,
						testutls.MockToken),
				},
				{
					Query:      regexp.QuoteMeta("SELECT COUNT(*) FROM `authors`;"),
					DbResponse: sqlmock.NewRows([]string{"count"}).AddRow(testutls.MockCount),
				},
			},
		},
	}

	for _, tt := range cases {

		if tt.err != nil {
			mock.ExpectQuery("SELECT `authors`.* FROM `authors`;").
				WithArgs().
				WillReturnError(fmt.Errorf("sql: no rows in sql"))
		} else {
			for _, dbQuery := range tt.dbQueries {
				mock.ExpectQuery(dbQuery.Query).
					WithArgs().
					WillReturnRows(dbQuery.DbResponse)
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			res, c, err := FindAllAuthorWithCount([]qm.QueryMod{}, context.Background())
			if err != nil {
				assert.Equal(t, true, tt.err != nil)
			} else {
				assert.Equal(t, err, tt.err)
				assert.Equal(t, testutls.MockCount, c)
				assert.Equal(t, res[0].Email, null.StringFrom(testutls.MockEmail))
				assert.Equal(t, res[0].Token, null.StringFrom(testutls.MockToken))
				assert.Equal(t, res[0].ID, int(testutls.MockID))

			}
		})
	}

}
