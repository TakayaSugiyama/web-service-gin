run:
	docker run -p 8080:8080 -it --rm --name gin-service gin-service
build:
	docker build -t gin-service .
exec:
	docker exec -it gin-service sh
rundb:
	docker run --rm --name gin-service-db -e MYSQL_USER=ginuser -e MYSQL_PASSWORD=password -e MYSQL_ROOT_PASSWORD=password -d  mysql:8.0
push:
	heroku container:push web
release:
	heroku container:release web
