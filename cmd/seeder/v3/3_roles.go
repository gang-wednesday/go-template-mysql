package main

import (
	"context"
	"fmt"

	"go-template/cmd/seeder/utls"
	"go-template/internal/mysql"
	"go-template/models"
	"go-template/pkg/utl/zaplog"

	"math/rand"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// to generate random sequence of alphabets
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// to generate one query for a particular author
func query(authorId int) string {
	c := "insert into posts(title,content,author_id) values"
	return fmt.Sprintf(c+" (%s, %s, %d);", randSeq(10), randSeq(30), authorId)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	db, _ := mysql.Connect()
	roles, err := models.Authors(qm.OrderBy("id ASC")).All(context.Background(), db)
	var insertQuery string
	for _, val := range roles {
		insertQuery = insertQuery + query(val.ID)
	}
	if err != nil {
		zaplog.Logger.Error("error while seeding", err)
	}

	_ = utls.SeedData("authors", insertQuery)
}
