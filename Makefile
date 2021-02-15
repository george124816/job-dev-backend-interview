DOCKER=$(shell which docker)
GO=$(shell which go)

run:
	${GO} run cmd/api/main.go

mysql_create:
	${DOCKER} run -d --name mysql --network host -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=main mysql:8.0

mysql_delete:
	${DOCKER} rm -f mysql

mysql_start:
	${DOCKER} start mysql