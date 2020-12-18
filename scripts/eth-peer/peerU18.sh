#!/bin/bash

# TODO: FIX ME, by deleting command-line arguments not specific to a peer node.

geth \
--nodiscover \
--networkid $ETH_NETWORK_ID \
--datadir $DATADIR \
--allow-insecure-unlock \
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
--verbosity 0 \
--syncmode "full" \
--pprof \
--pprof.addr="0.0.0.0" \
--metrics \
--metrics.expensive \
--metrics.addr "0.0.0.0" \
console 2>> PeerEth.log