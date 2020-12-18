#!/usr/bin/env sh

start() {

}

stop() {
  echo "Stopping SWARM in systemd."
}

case $1 in
  start|stop) "$1"
esac