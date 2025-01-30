#!/usr/bin/bash

docker system prune -f

GATLING_WORKSPACE=$(pwd)/stress-test
GATLING_BIN_DIR=$GATLING_WORKSPACE/bin
RESULTS_WORKSPACE=$GATLING_WORKSPACE/resultados
TEST_ARTIFACTS=$GATLING_WORKSPACE/test_artifacts

echo "iniciando e logando execução da API"
mkdir -p "$RESULTS_WORKSPACE"
docker compose up -d --build
docker compose logs > "$RESULTS_WORKSPACE/docker-compose.logs"
echo "pausa de 6 segundos pro start da API"
sleep 6
echo "iniciando teste"

$GATLING_BIN_DIR/gatling.sh \
    -rm local \
    -s RinhaBackendSimulation \
    -rd "DESCRIPTION"\
    -rf "$RESULTS_WORKSPACE" \
    -sf $GATLING_WORKSPACE/user-files/simulations\
    -rsf $GATLING_WORKSPACE/user-files/resources

echo "teste finalizado"
echo "fazendo request e salvando a contagem de pessoas"
SAVE_CONTAGEM="$RESULTS_WORKSPACE/contagem-pessoas-$(date).log"
curl -v "http://localhost:9999/contagem-pessoas" > "$SAVE_CONTAGEM"
echo "resultado da contagem em $SAVE_CONTAGEM"
echo "INSERTS: $(cat "$SAVE_CONTAGEM")"
echo "cleaning up do docker"
docker compose rm -f
docker compose down
