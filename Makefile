run:
	docker run -p 8080:8080 -it --rm --name gin-service\
		-v ${PWD}/:/app\
		--network gin-service-network\
		--hostname gin_server\
		gin-service
build:
	docker build -t gin-service -f dockers/Dockerfile .
exec:
	docker exec -it gin-service sh
rundb:
	docker run --rm --name gin-service-db -e MYSQL_USER=ginuser \
		-e MYSQL_PASSWORD=password -e MYSQL_ROOT_PASSWORD=password \
		-e MYSQL_DATABASE=gin-service \
		-p 3307:3306 \
		--network gin-service-network\
		--hostname gin_service_db\
		-d  mysql:8.0
test:
	docker exec gin-service go test -v
buildcliant:
	docker build -f dockers/Dockerfile.cliant -t gin-service-cliant ./cliant/
runcliant:
	docker run -p 3000:3000 -v ${PWD}/cliant:/usr/src/app\
		--name gin-service-cliant  \
		--rm \
		--network gin-service-network\
		--hostname gin_service_cliant\
		gin-service-cliant
stopdb:
	docker stop gin-service-db
createnetwork:
	docker network create gin-service-network
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
