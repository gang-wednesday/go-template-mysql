package main

import (
	"go-template/cmd/seeder/utls"
	"log"
	"os"
)

func main() {
	base, _ := os.Getwd()
	log.Println(base)

	_ = utls.SeedData("roles",
		`INSERT INTO roles(access_level, name) VALUES (100, 'SUPER_ADMIN');
		INSERT INTO roles(access_level, name) VALUES (200, 'USER');
		INSERT INTO roles(access_level, name) VALUES (110, 'ADMIN');`)
}
