.DEFAULT_GOAL:=migration_up

.PHONY = migration_create migration_up migration_down migration_fix

VERSION = 1
MIGRATION_N = 1
MIGRATIONS_DIR = migrations
DATABASE_URL = 'postgres://test:test@localhost:5432/test?sslmode=disable'

migration_version:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) version

migration_drop:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) drop
	
migration_create:
	@migrate create -dir $(MIGRATIONS_DIR) -ext sql -seq users

migration_down_all:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) -verbose down

migration_down:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) -verbose down $(MIGRATION_N)

migration_up_all:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) -verbose up

migration_up:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) -verbose up $(MIGRATION_N)

migration_fix:
	@migrate -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) -verbose force $(VERSION)

