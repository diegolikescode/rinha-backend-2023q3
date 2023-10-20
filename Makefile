go:
	docker-compose down --volumes --remove-orphans
	docker-compose -f docker-compose.local.yml up --build -d
	sleep 20
	go run main.go

build:
	go build -o main

compose:
	docker build -t rinha2023q3 .
	docker-compose up --build

kill:
	docker-compose down --volumes --remove-orphans
	docker container prune -f
	# docker system prune -f
	docker ps -a

dethatched:
	docker build -t rinha2023q3 .
	docker-compose up -d --build
	docker-compose logs > "/home/camel/src/rinha-backend-2023q3/resultados/primeira-fase/diegolikescode/docker-compose.logs"

updown: 
	$(MAKE) kill 
	sleep 2
	$(MAKE) compose

my_pgadmin:
	docker run -d --name my_pgadmin --network rinha-network --health-cmd="pg_isready -U rinha" --health-interval=10s --health-timeout=5s --health-retries=3 -e PGADMIN_DEFAULT_EMAIL=admin@example.com -e PG




