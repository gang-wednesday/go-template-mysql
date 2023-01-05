package main

import (
	"context"
	"fmt"

	"go-template/cmd/seeder/utls"
	"go-template/internal/mysql"
	"go-template/models"
	"go-template/pkg/utl/secure"
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

// to generate query for a particular role
func query(sec *secure.Service, roleId int) string {
	c := "INSERT into authors(username,email,password,role_id) VALUES"
	return fmt.Sprintf(c+" (%s, %s, '%s', %d);", randSeq(10), randSeq(5)+"@gmail.com",
		sec.Hash("adminuser"), roleId)
}
func main() {
	rand.Seed(time.Now().UnixNano())

	sec := secure.New(1, nil)
	db, _ := mysql.Connect()
	roles, err := models.Roles(qm.OrderBy("id ASC")).All(context.Background(), db)
	var insertQuery string
	for _, val := range roles {
		insertQuery = insertQuery + query(sec, val.ID)
	}
	if err != nil {
		zaplog.Logger.Error("error while seeding", err)
	}

	_ = utls.SeedData("authors", insertQuery)
}
