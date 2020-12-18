#!/bin/bash


swarm --identity $SWARM_NODE_NAME \
--ens-api $IPC_FILE \
--datadir $SWARM_DATADIR \
--port 30399 \
--httpaddr "0.0.0.0" \
--bzzapi http://0.0.0.0:8500 \
--corsdomain "*" \
--bzzaccount $SWARM_ACCOUNT \
--password $SWARM_ACCOUNT_PASSWORD