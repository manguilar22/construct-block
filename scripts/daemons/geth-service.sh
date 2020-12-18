#!/usr/bin/env sh

start() {
  ETH_NETWORK_ID=""
  DATADIR=""
  PUBLIC_KEY=""
  PASSWORD_FILE=""

}

stop() {
  echo "Stopping GETH in systemd."
}

case $1 in
  start|stop) "$1"
esac