#!/bin/bash

clef --configdir $DATADIR \
--keystore $KEYSTORE \
--nousb \
--chainid $ETH_NETWORK_ID \
--http \
--http.addr "0.0.0.0" \
--http.port 8550 \
--rules rules.js  \
--signersecret ./clef-sign.txt