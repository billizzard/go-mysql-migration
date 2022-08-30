package main

import (
	"flag"
	"github.com/billizzard/go-mysql-migration/internal/command/db"
	"github.com/billizzard/go-mysql-migration/internal/command/flags"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		log.Fatal("Command is required")
	}

	generate := flags.NewGenerateFlag()
	init := flags.NewInitFlag()
	up := flags.NewUpFlag()
	flag.Parse()

	if os.Args[1] != "init" {
		db.ConnectDb()
	}

	switch os.Args[1] {
	case "up":
		err = up.Run(os.Args)
	case "init":
		err = init.Run(os.Args)
	case "generate":
		err = generate.Run(os.Args)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
