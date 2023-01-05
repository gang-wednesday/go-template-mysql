package main

import (
	"go-template/cmd/seeder/utls"
)

func main() {

	_ = utls.SeedData("roles",
		`INSERT INTO roles(access_level, name) VALUES (100, 'SUPER_ADMIN');
		INSERT INTO roles(access_level, name) VALUES (200, 'USER');
		INSERT INTO roles(access_level, name) VALUES (110, 'ADMIN');`)
}
