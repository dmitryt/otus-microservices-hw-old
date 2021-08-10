#!/usr/bin/env bash

RELEASE_NAME="prod"
PG_RELEASE_NAME="pg"
TESTS_PATH="pg"
CHART_DIR="k8s/charts/hw02-k8s-service"
TESTS_FILE_PATH="Users_API.postman_collection.json"

function usage {
    echo "Usage: hw02-k8s-service {start|stop|installDB|dropDB|test}"
    exit 1
}

function installDB {
    echo "Installing and starting Postgres...."
    helm upgrade --install $PG_RELEASE_NAME bitnami/postgresql -f $CHART_DIR/db-values.yaml
    echo "Installing and starting Postgres....Done"
}

function dropDB {
    echo "Dropping Postgres...."
    helm uninstall $PG_RELEASE_NAME
    echo "Dropping Postgres....Done"
}

function startService {
    echo "Starting service...."
    echo "Updating /etc/hosts...."
    sudo sh -c 'echo "$(minikube ip) arch.homework" >> /etc/hosts'
    echo "Updating /etc/hosts....Done"
    helm upgrade --install $RELEASE_NAME $CHART_DIR
    echo "Starting service....Done"
}

function stopService {
    echo "Stopping service...."
    echo "Updating /etc/hosts...."
    sudo sed -i '' '/arch\.homework$/d' /etc/hosts
    echo "Updating /etc/hosts....Done"
    helm uninstall $RELEASE_NAME
    echo "Stopping service....Done"
}

function runTests {
  newman run $TESTS_FILE_PATH
}

case $1 in
  installDB) installDB ;;
  dropDB) dropDB ;;
  start) startService ;;
  stop) stopService ;;
  test) runTests ;;
  *) usage ;;
esac
