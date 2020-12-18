#!/bin/bash

# Ethereum - Peer Node

geth --networkid 2020 --nodiscover \
--datadir ./node1 \
--syncmode "full" \
--rpc \
--rpcaddr "0.0.0.0" \
--rpcport 8545 \
--rpcapi "eth,admin,txpool,net,web3,personal,miner,debug" \
--rpccorsdomain "*" \
--ws \
--wsaddr "0.0.0.0" \
--wsport 8546 \
--wsorigins "*" \
--wsapi "eth,admin,txpool,net,web3,personal,miner,debug" \
--pprof \
--pprofaddr "0.0.0.0" \
--pprofport 6060 \
--metrics \
--metrics.expensive \
--verbosity 0 \
console