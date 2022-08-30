package flags

import (
	"errors"
	"flag"
	"fmt"
	"github.com/billizzard/go-mysql-migration/configs"
	"github.com/billizzard/go-mysql-migration/internal/command/db"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Migration struct {
	version string
	file    string
	done    bool
}

type UpFlag struct {
	flag *flag.FlagSet
}

func NewUpFlag() *UpFlag {
	return &UpFlag{
		flag: flag.NewFlagSet("up", flag.ExitOnError),
	}
}

func (f *UpFlag) Run(args []string) error {
	err := f.flag.Parse(args[2:])
	if err != nil {
		return err
	}

	count := 0
	if len(args) > 2 {
		count, err = strconv.Atoi(args[2])
		if err != nil {
			return errors.New("Second argument need to be integer, now: " + args[3])
		}
	}

	migrations, errFiles := f.getMigrationsFromFiles()
	if errFiles != nil {
		return errFiles
	}

	return f.updateDbFromMigrations(migrations, count)
}

func (f *UpFlag) getMigrationsFromFiles() (map[string]*Migration, error) {
	files, err := ioutil.ReadDir(configs.PATH)
	if err != nil {
		return nil, err
	}

	migrations := make(map[string]*Migration)
	for _, file := range files {
		version, errName := f.getDateFromName(file.Name())
		if errName != nil {
			return nil, errName
		}
		migrations[version] = &Migration{version: version, file: file.Name(), done: false}
	}

	return migrations, nil
}

func (f *UpFlag) updateDbFromMigrations(files map[string]*Migration, count int) error {
	log.Println("Start migrate...")
	rows, err := db.DB.Query("SELECT migration FROM `migrations`;")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var migration string
		err := rows.Scan(&migration)
		if err != nil {
			return err
		}

		version, errName := f.getDateFromName(migration)
		if errName != nil {
			return errName
		}

		if _, ok := files[version]; ok {
			files[version].done = true
		}
	}

	// for iterate map by order
	versions := make([]string, 0, len(files))
	for version := range files {
		versions = append(versions, version)
	}

	sort.Strings(versions)
	stopCount := 0
	for _, version := range versions {
		if count > 0 && count <= stopCount {
			break
		}

		v := files[version]
		if v.done {
			continue
		}

		log.Println("Execute migration " + v.file)

		sqlContent, err := ioutil.ReadFile(configs.PATH + "/" + v.file)
		if err != nil {
			log.Fatal("Can not read migration " + v.file + ". " + err.Error())
		}

		_, err = db.DB.Exec(string(sqlContent))
		if err != nil {
			return errors.New(fmt.Sprintf("Error while run migration %s. %s", v.file, err.Error()))
		}

		timeNow := time.Now()
		_, err = db.DB.Exec(fmt.Sprintf(
			"INSERT INTO migrations (migration, executed_at) VALUES ('%s', '%s');", v.file,
			fmt.Sprintf("%d-%d-%d %d:%d:%d", timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second()),
		))

		if err != nil {
			panic("Error while save migration version. Err: " + err.Error())
		}

		stopCount++
	}

	log.Println("Successfully migrated")

	return nil
}

func (f *UpFlag) getDateFromName(name string) (string, error) {
	nameParts := strings.Split(name, "_")
	if len(nameParts) == 0 {
		return "", errors.New("Not valid name for migration file: " + name)
	}

	return nameParts[0], nil
}
