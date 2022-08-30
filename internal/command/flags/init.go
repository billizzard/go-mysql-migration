package flags

import (
	"flag"
	"fmt"
	"github.com/billizzard/go-mysql-migration/configs"
	db2 "github.com/billizzard/go-mysql-migration/internal/command/db"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type InitFlag struct {
	flag *flag.FlagSet
}

func NewInitFlag() *InitFlag {
	return &InitFlag{
		flag: flag.NewFlagSet("init", flag.ExitOnError),
	}
}

func (f *InitFlag) Run(args []string) error {
	err := f.flag.Parse(args[2:])
	if err != nil {
		return err
	}

	return f.initDb()
}

func (f *InitFlag) initDb() error {
	log.Println("Initialize migration tools...")
	db := db2.ConnectMysql()

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + configs.DB_NAME)
	if err != nil {
		return err
	}

	_, err = db.Exec("USE " + configs.DB_NAME)
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (migration varchar(255), executed_at datetime)", configs.MIGRATION_TABLE))
	if err != nil {
		return err
	}

	log.Println("Create table " + configs.MIGRATION_TABLE)

	return nil
}
