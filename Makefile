build_spacemarket:
	go build -o ./builds/spacemarket ./cmd/spacemarket

build_instabase:
	go build -o ./builds/instabase ./cmd/instabase

build_noonde:
	go build -o ./builds/noonde ./cmd/noonde

build_deploy:
	go build -o ./builds/deploy ./cmd/deploy

migrate_up:
	./scripts/migrate up

migrate_down:
	./scripts/migrate down

db_create:
	./scripts/create_database

db_drop:
	./scripts/drop_database

db_reset:
	./scripts/drop_database
	./scripts/create_database
	./scripts/migrate up

install_migration_cli:
	./scripts/install_migration_cli

test:
	go test ./...

copy_build_to_production:
	scp ./builds/noonde_linux noonde:/home/ubuntu/app/noonde

elasticsearch_reset:
	./scripts/elastic_search_drop_index.sh noonde
	./scripts/elastic_search_create_index.sh noonde

dejavue_open_listings:
	./scripts/dejavu_start.sh noonde_listings

dejavue_open_search_requests:
	./scripts/dejavu_start.sh noonde_search_requests

dejavue_open_place_suggestions:
	./scripts/dejavu_start.sh noonde_place_suggestions

dejavue_open_users:
	./scripts/dejavu_start.sh noonde_users

dejavue_open_reviews:
	./scripts/dejavu_start.sh noonde_reviews

dejavue_stop:
	./scripts/dejavu_stop.sh
