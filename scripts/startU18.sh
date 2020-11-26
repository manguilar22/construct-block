#!/bin/bash

#bash ./scripts/config.sh

geth \
--networkid $ETH_NETWORK_ID \
--datadir $DATADIR \
--jspath ./scripts/ \
--ws \
--ws.addr "0.0.0.0" \
--ws.port 8546 \
--ws.origins "*" \
--ws.api "eth,admin,rpc,txpool,net,web3,personal,miner,debug" \
--http \
--http.addr "0.0.0.0" \
--http.port 8545 \
--http.corsdomain "*" \
--rpcapi "eth,admin,rpc,txpool,net,web3,personal,miner,debug" \
--unlock $PUBLIC_KEY \
--password $PASSWORD_FILE \
--verbosity 0 \
--syncmode "full" \
--allow-insecure-unlock \
--pprof \
--pprof.addr="0.0.0.0" \
--mine \
--miner.threads 8 \
--metrics.addr "0.0.0.0" \
--metrics \
--metrics.expensive \
#console 2>> Eth.log