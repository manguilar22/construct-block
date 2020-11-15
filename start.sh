#!/bin/bash

#bash ./scrips/config.sh
#bash ./scripts/priv.sh

geth \
--networkid $ETH_NETWORK_ID \
--datadir ./node1/ \
--ws.origins "*" \
--rpc \
--rpcport 8545 \
--rpccorsdomain "*" \
--rpcaddr "0.0.0.0" \
--rpcapi "eth,admin,rpc,txpool,net,web3,personal,miner,debug" \
--unlock $PUBLIC_KEY \
--password $PASSWORD \
--verbosity 0 \
--syncmode "full" \
--allow-insecure-unlock \
--mine \
--miner.threads 8 \
console 2>> Eth.log
