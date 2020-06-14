# Db migrations usage
Db migrations using tool https://github.com/golang-migrate/migrate/

# Installation
## Windows(PowerShell)
```PowerShell
go get -u -d github.com/golang-migrate/migrate/cmd/migrate
cd $Env:GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
git checkout $TAG  # e.g. v4.1.0
go build -tags 'postgres' -ldflags="-X main.Version=$TAG" -o $Env:GOPATH/bin/migrate $Env:GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
```

## Linux
```bash
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate
$ cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ git checkout $TAG  # e.g. v4.1.0
$ go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
```

# Usage
See Makefile -> migrate
```PowerShell
migrate -path ./db/migrations/ -database postgres://postgres:example@localhost:5432/postgres?sslmode=disable up
```

# Migrations
Folder migrations contains files what applies to database
Namig convention is simple NNN_NAME.[up|down].sql
where 
* NNN - order of migration
* NAME - simple human readable name