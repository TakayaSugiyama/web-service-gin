run:
	docker run -p 8080:8080 -it --rm --name gin-service gin-service
build:
	docker build -t gin-service .
exec:
	docker exec -it gin-service sh
rundb:
	docker run --rm --name gin-service-db -e MYSQL_USER=ginuser \
		-e MYSQL_PASSWORD=password -e MYSQL_ROOT_PASSWORD=password \
		-e MYSQL_DATABASE=gin-service \
		-p 3307:3306 \
		-d  mysql:8.0
stopdb:
	docker stop gin-service-db
execdb:
	docker exec -it gin-service-db bash
migrateup:
	migrate -database="mysql://ginuser:password@tcp(localhost:3307)/gin-service"\
		-path "./db/migrations"\
		--verbose up
migratedown:
	migrate -database="mysql://ginuser:password@tcp(localhost:3307)/gin-service"\
		-path "./db/migrations"\
		--verbose down
push:
	heroku container:push web
release:
	heroku container:release web
