## Running the crawler:

Make sure you have postgres installed and access to `psql` and related commands `createdb`, etc.

```bash
# install golang-migration cli:
make install_migration_cli
```

```bash
# setup config file
$ cat <<EOF > configs/development.yml
database:
  host: localhost
  port: 5432
  name: spacemarket_crawler_development
  password: $(whoami)
  username: $(whoami)
  url: postgres://$(whoami):$(whoami)@localhost:5432/spacemarket_crawler_development?sslmode=disable

EOF
```

```bash
# create database
make create_db

# migrate database
make migrate_up
```

```bash
# install dependencies
go mod download
```

```bash
# make and run
make build && ./builds/spacemarket_crawler
```

## Managing dependencies:

This project is using go modules.

```bash
go help mod
```

***Adding a dependency***:
```bash
go get -u github.com/foo/bar
```

***Removing a dependency***:
```bash
go clean -i github.com/foo/bar
```



####

Docker

docker tag 518a41981a6a yanshiyason/spacemarket
docker push yanshiyason/spacemarket

