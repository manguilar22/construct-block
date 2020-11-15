#!/bin/bash

# Start Ethereum Node
# ws
#--shh  (Whisper)
geth --networkid 2020
--datadir ./node1/
--ws
--ws.origins "*"
--rpc
--rpcport 8545
--rpccorsdomain "*"
--rpcaddr "0.0.0.0"
--rpcapi "eth,admin,rpc,txpool,net,web3,personal,miner,debug"
--unlock $PUBLIC_KEY
--password $PASSWORD
--verbosity 7
--syncmode "full"
--allow-insecure-unlock
--mine
--miner.threads 8
--metrics
--metrics.expensive
console 2>> Eth.log

geth --networkid $ETH_NETWORK_ID --datadir ./node1/ --ws --ws.origins "*" --rpc --rpcport 8545 --rpccorsdomain "*" --rpcaddr "0.0.0.0" --rpcapi "eth,admin,rpc,txpool,net,web3,personal,miner,debug" --unlock $PUBLIC_KEY --password $PASSWORD --verbosity 0 --syncmode "full" --allow-insecure-unlock --mine --miner.threads 8 --metrics --metrics.expensive console 2>> Eth.log
