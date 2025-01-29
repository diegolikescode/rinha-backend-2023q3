#!/usr/bin/bash

docker system prune -f

GATLING_BIN_DIR=$(pwd)/teste/gatling/deps/gatling/bin
GATLING_WORKSPACE=$(pwd)/stress-test
RESULTS_WORKSPACE=$(pwd)/resultados/primeira-fase
DIEGO=diegolikescode

echo "iniciando e logando execução da API"
mkdir "$RESULTS_WORKSPACE/$DIEGO"
docker compose up -d --build
docker compose logs > "$RESULTS_WORKSPACE/$DIEGO/docker-compose.logs"
echo "pausa de 6 segundos para startup pra API"
sleep 6
echo "iniciando teste"
sh $GATLING_BIN_DIR/gatling.sh -rm local -s RinhaBackendSimulation \
    -rd $DIEGO\
    -rf "$RESULTS_WORKSPACE/$DIEGO" \
    -sf $GATLING_WORKSPACE/user-files/simulations \
    -rsf $GATLING_WORKSPACE/user-files/resources
echo "teste finalizado"
echo "fazendo request e salvando a contagem de pessoas"
SAVE_CONTAGEM="$RESULTS_WORKSPACE/$DIEGO/contagem-pessoas-$(date).log"
curl -v "http://localhost:9999/contagem-pessoas" > "$SAVE_CONTAGEM"
echo "resultado da contagem em $SAVE_CONTAGEM"
cat "$SAVE_CONTAGEM"
echo "cleaning up do docker"
docker compose rm -f
docker compose down
touch "$RESULTS_WORKSPACE/$DIEGO/testado"
