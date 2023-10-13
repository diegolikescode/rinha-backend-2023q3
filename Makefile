go:
	go run main.go

build:
	go build -o main

compose:
	docker build -t rinha2023q3 .
	docker-compose up --build

kill:
	docker-compose down --volumes --remove-orphans
	docker ps -a

dethatched:
	# docker build -t rinha2023q3 .
	docker-compose up -d --build
	docker-compose logs > "/home/camel/src/rinha-backend-2023q3/resultados/primeira-fase/diegolikescode/docker-compose.logs"

updown: kill compose
