go:
	docker compose down --volumes --remove-orphans
	docker compose -f docker-compose.yml up db --build -d
	go run main.go

t:
	# F in the chat
	docker compose down --remove-orphans
	docker system prune -f
	docker compose -f docker-compose.monitor.yml down --remove-orphans

	docker compose -f docker-compose.monitor.yml up -d

	echo "monitor has started composing, starting stress test"
	./stress-it.sh

build:
	go build -o main

compose:
	docker build -t rinha2023q3 .
	docker compose up --build

dethatched:
	docker build -t rinha2023q3 .
	docker compose up -d --build
	docker compose logs > "/home/camel/src/rinha-backend-2023q3/resultados/primeira-fase/diegolikescode/docker-compose.logs"

kill:
	docker compose down --volumes --remove-orphans
	docker ps -a

slp:
	sleep 2

updown: kill slp compose

my_pgadmin:
	docker run -d --name my_pgadmin --network rinha-network --health-cmd="pg_isready -U rinha" --health-interval=10s --health-timeout=5s --health-retries=3 -e PGADMIN_DEFAULT_EMAIL=admin@example.com -e PG
