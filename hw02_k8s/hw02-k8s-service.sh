#!/usr/bin/env bash

RELEASE_NAME="prod"
PG_RELEASE_NAME="pg"
CHART_DIR="k8s/charts/hw02-k8s-service"

function usage {
    echo "Usage: hw02-k8s-service {start|stop|installDB|dropDB}"
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
    helm upgrade --install $RELEASE_NAME $CHART_DIR
    echo "Starting service....Done"
}

function stopService {
    echo "Stopping service...."
    helm uninstall $RELEASE_NAME
    echo "Stopping service....Done"
}

case $1 in
  installDB) installDB ;;
  dropDB) dropDB ;;
  start) startService ;;
  stop) stopService ;;
  *) usage ;;
esac
