#!/bin/bash

# setup for ubuntu 16
sudo apt-get install wget ca-certificates
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'

sudo apt-get update
sudo apt-get install postgresql postgresql-contrib

# Success. You can now start the database server using:
# /usr/lib/postgresql/11/bin/pg_ctl -D /var/lib/postgresql/11/main -l logfile start

# sudo su - postgres
# psql
# >
# CREATE USER ubuntu WITH PASSWORD 'ubuntu' CREATEDB;
# CREATE DATABASE ubuntu;

mkdir app app/configs app/tools

export SPACEMARKET_CRAWLER_ENV="production"
echo 'export SPACEMARKET_CRAWLER_ENV="production"' >>~/.bashrc

cat <<EOF >app/configs/production.yml
database:
  host: localhost
  port: 5432
  name: spacemarket_crawler_production
  password: $(whoami)
  username: $(whoami)
  url: postgres://$(whoami):$(whoami)@localhost:5432/spacemarket_crawler_production?sslmode=disable
EOF

scp ./scripts spacemarket_crawler:/home/ubuntu/app
scp -r ./migrations spacemarket_crawler:/home/ubuntu/app
scp ./Makefile spacemarket_crawler:/home/ubuntu/app

# Install make!
sudo apt install make

make install_migration_cli
make create_db
make migrate_up

scp ./builds/spacemarket_crawler_linux spacemarket_crawler:/home/ubuntu/app
