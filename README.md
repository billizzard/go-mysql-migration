### GO Mysql migrations
Very simple implementation of migrations using sql files. Transactions are written in files if needed

To initialize migrations library:
```bash
go install github.com/billizzard/go-mysql-migration@latest
```
If you want, you can rename binary from "go-mysql-migration" file to "migrations"
```
mv $GOPATH/bin/go-mysql-migration $GOPATH/bin/migrations
```
For proper work need to set following env variables:
```
MYSQL_USER
MYSQL_PASS
MYSQL_HOST
MYSQL_PORT
MYSQL_DBNAME
MIGRATION_PATH (in this repo examples sql migrations in folder internal/migrations)

Example:
export MYSQL_USER=user MYSQL_PASS=pass MYSQL_HOST=localhost MYSQL_PORT=3306 MYSQL_DBNAME=mydb MIGRATION_PATH=internal/migrations
```
All commands run from folder witch is parent for MIGRATION_PATH folder, in our repo it's root folder, because "internal/migrations" contains in root folder

To initialize database use (create migrations table), need run only once:
```bash
migrations init
```
To generate migration template use (file will appear in folder MIGRATION_PATH):
```bash
migrations generate name_migration
```
Fill generated file as you need

To run all not used migrations use:
```bash
migrations up
```
To run certain amount migrations use:
```bash
migrations up 2
```

Migrations "down" not implemented. Just create new migrations with rollback