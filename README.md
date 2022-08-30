### Migrations
To initialize migrations library:
```bash
go install github.com/billizzard/go-mysql-migration
```
For proper work need to set following env variables:
```
MYSQL_USER
MYSQL_PASS
MYSQL_HOST
MYSQL_PORT
MYSQL_DBNAME
MIGRATION_PATH (in our case migrations in folder internal/migrations)

export MYSQL_USER=user MYSQL_PASS=pass MYSQL_HOST=localhost MYSQL_PORT=3306 MYSQL_DBNAME=mydb MIGRATION_PATH=internal/migrations

```
To initialize database use:
```bash
migrations init
```
To generate migration template use (file will appear in folder db/migrations):
```bash
migrations generate name_migration
```
To run all not used migrations use:
```bash
migrations up
```
To run certain amount migrations use:
```bash
migrations up 2
```