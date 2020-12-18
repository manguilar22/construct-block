#!/bin/bash

bee start \
--data-dir /home/user/eth/node1/ \
--swap-enable \
--swap-endpoint http://localhost:8545 \
--debug-api-enable \
--debug-api-addr=localhost:1635 \
--clef-signer-enable \
--clef-signer-endpoint /home/user/eth/node1/clef.ipc