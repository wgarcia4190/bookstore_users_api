package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wgarcia4190/bookstore_users_api/internal/datasources/mysql/users_db"
	"github.com/wgarcia4190/bookstore_users_api/internal/utils/schema"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

func run() error {
	flags := os.Args[1:]

	if len(flags) > 0 {
		flag := strings.Split(flags[0], " ")[0]

		var err error

		switch flag {
		case "migrate":
			err = migrate()
		}

		return err
	}

	return errors.New("must specify a command")
}

func migrate() error {
	if err := schema.Migrate(users_db.Client); err != nil {
		fmt.Println("error migrating db")
		return err
	}

	fmt.Println("Migrations complete")
	return nil
}
