#!/bin/bash

#bash ./scripts/config.sh

geth \
--nodiscover \
--networkid $ETH_NETWORK_ID \
--datadir $DATADIR \
--syncmode "full" \
--preload ./scripts/GethFunc.js \
--http \
--http.addr "0.0.0.0" \
--http.port 8545 \
--http.corsdomain "*" \
--http.api "admin,clique,debug,eth,miner,net,personal,rpc,txpool,web3" \
--ws \
--ws.addr "0.0.0.0" \
--ws.port 8546 \
--ws.origins "*" \
--ws.api "admin,clique,debug,eth,miner,net,personal,rpc,txpool,web3" \
--rpc \
--rpcaddr "0.0.0.0" \
--rpcport 8545 \
--rpccorsdomain "*" \
--rpcapi "admin,clique,debug,eth,miner,net,personal,rpc,txpool,web3" \
--unlock $PUBLIC_KEY \
--password $PASSWORD_FILE \
--allow-insecure-unlock \
--mine \
--miner.threads 8 \
--verbosity 0 \
--metrics \
--metrics.expensive \
--metrics.addr "0.0.0.0" \
--pprof \
--pprof.addr="0.0.0.0" \
console 2>> Eth.log
