.PHONY: test seed db_create db_drop db_reset dev coverage
.DEFAULT_GOAL := test

test:
	GO_ENV=test buffalo test

seed:
	buffalo task db:seed

db_create:
	buffalo db create
	mysql -uroot < scripts/create_db_users.sql

db_drop:
	buffalo db drop

db_migrate:
	buffalo db migrate

db_reset: db_drop db_create db_migrate

run_dev:
	buffalo dev

build:
	buffalo build

# mocks:
# 	${GOPATH}/bin/mockery -recursive=true -name=SearchIndexer

coverage:
	GO_ENV=test buffalo test -coverprofile fmt ./...
