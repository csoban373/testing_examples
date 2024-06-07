.PHONY: test-db-create
test-db-create:
	mysql -u root -e "CREATE DATABASE IF NOT EXISTS test;" && \
		mysql -u root test < db/schema.sql