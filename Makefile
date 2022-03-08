run:
	docker run -p 8080:8080 -it --rm --name gin-service gin-service
build:
	docker build -t gin-service .
exec:
	docker exec -it gin-service sh
push:
	heroku container:push web
release:
	heroku container:release web
