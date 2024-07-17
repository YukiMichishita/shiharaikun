include .env.local
export $(shell sed 's/=.*//' .env.local)

SQL_FILE := db/migration/init_db.sql

.PHONY: all run_mysql init_db clean stop_mysql

run_mysql:
	@echo "Starting MySQL container..."
	@docker run --name $(MYSQL_CONTAINER_NAME) \
	    -e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
	    -e MYSQL_DATABASE=$(MYSQL_DATABASE) \
	    -e MYSQL_USER=$(MYSQL_USER) \
	    -e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	    -e MYSQLD_OPTS="--default-authentication-plugin=mysql_native_password" \
	    -p $(MYSQL_PORT):3306 \
	    -d mysql:latest

init_db:
	@echo "Initializing database..."
	@docker exec -i $(MYSQL_CONTAINER_NAME) \
	    mysql -u root -p$(MYSQL_ROOT_PASSWORD) $(MYSQL_DATABASE) < $(SQL_FILE)
	@echo "Database initialized."

stop_mysql:
	@echo "Stopping and removing MySQL container..."
	@docker stop $(MYSQL_CONTAINER_NAME)
	@docker rm $(MYSQL_CONTAINER_NAME)
	@echo "MySQL container stopped and removed."

gorm_gen:
	@cd ./cmd/gormgen && go run main.go

run_serer:
	@go run cmd/rest/main.go
