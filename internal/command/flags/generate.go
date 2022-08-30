package flags

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/billizzard/go-mysql-migration/configs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"text/template"
	"time"
)

type GenerateFlag struct {
	flag *flag.FlagSet
}

func NewGenerateFlag() *GenerateFlag {
	return &GenerateFlag{
		flag: flag.NewFlagSet("generate", flag.ExitOnError),
	}
}

func (f *GenerateFlag) Run(args []string) error {
	err := f.flag.Parse(args[2:])
	if err != nil {
		return err
	}

	if len(args) < 3 {
		return errors.New("Need specify name for generaged migration: migrate my_name")
	}

	return f.createMigrationFile(args[2])
}

func (f *GenerateFlag) createMigrationFile(name string) error {
	log.Println("Generating migration file...")
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	//t := template.Must(template.ParseFiles(config.TEMPLATE_FILE_PATH))
	t, _ := template.New("migrations").Parse("-- +migrate Up\nCREATE TABLE example (id int);")
	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template. Error:" + err.Error())
	}

	filePath := fmt.Sprintf("%s/%s_%s.sql", configs.PATH, version, name)
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("Unable to create migration file: " + filePath + " Error: " + err.Error())
	}
	defer file.Close()

	if _, err := file.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file: " + filePath + " Error: " + err.Error())
	}

	log.Println("Migration file '" + file.Name() + "' successfully generated")

	return nil
}
