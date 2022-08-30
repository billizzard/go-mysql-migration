package configs

import "os"

const MIGRATION_TABLE = "migrations"

var DB_USER = os.Getenv("MYSQL_USER")
var DB_PASS = os.Getenv("MYSQL_PASS")
var DB_HOST = os.Getenv("MYSQL_HOST")
var DB_PORT = os.Getenv("MYSQL_PORT")
var DB_NAME = os.Getenv("MYSQL_DBNAME")
var PATH = os.Getenv("MIGRATION_PATH")
