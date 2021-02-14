DOCKER=$(shell which docker)

mysql_create:
	${DOCKER} run -d --name mysql --network host -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=main mysql:8.0

mysql_delete:
	${DOCKER} rm -f mysql