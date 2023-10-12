#!/usr/bin/bash

docker system prune -f

# /home/camel/src/rinha-backend-2023q3/teste/gatling/deps/gatling/bin
GATLING_BIN_DIR=$(pwd)/teste/gatling/deps/gatling/bin
GATLING_WORKSPACE=$(pwd)/stress-test
RESULTS_WORKSPACE=$(pwd)/resultados/primeira-fase
DIEGO=diegolikescode

echo "iniciando e logando execução da API"
mkdir "$RESULTS_WORKSPACE/$DIEGO"
docker-compose up -d --build
docker-compose logs > "$RESULTS_WORKSPACE/$DIEGO/docker-compose.logs"
echo "pausa de 10 segundos para startup pra API"
sleep 45
echo "iniciando teste"
sh $GATLING_BIN_DIR/gatling.sh -rm local -s RinhaBackendSimulation \
    -rd $DIEGO\
    -rf "$RESULTS_WORKSPACE/$DIEGO" \
    -sf $GATLING_WORKSPACE/user-files/simulations \
    -rsf $GATLING_WORKSPACE/user-files/resources
echo "teste finalizado"
echo "fazendo request e salvando a contagem de pessoas"
curl -v "http://localhost:9999/contagem-pessoas" > "$RESULTS_WORKSPACE/$DIEGO/contagem-pessoas.log"
echo "cleaning up do docker"
docker-compose rm -f
docker-compose down
touch "$RESULTS_WORKSPACE/$DIEGO/testado"
