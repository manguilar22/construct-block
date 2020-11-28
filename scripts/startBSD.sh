#!/bin/csh

# Ethereum Config ( ~/.cshrc )
setenv ETH_NETWORK_ID 0
setenv DATADIR /usr/home/${USER}
setenv VERSION 0                                         # 0=hostname, 1=HTTP
setenv HOSTNAME 0                                        # HTTP, Web Socket, RPC
setenv PUBLIC_KEY 0x...                                  # Signing Hex Address
setenv KEYSTORE /usr/home/${USER}                        # Keystore File Directory
setenv PRIVATE_KEY 0                                     # Private Key Hex Address
setenv PASSWORD_FILE /usr/home/${USER}                   # KeyStore Password
--keystore ./bsdNode1/keystore/ \

# Standalone - Ethereum Node
geth \
--nodiscover \
--networkid $ETH_NETWORK_ID \
--datadir $DATADIR \
--jspath ./scripts/ \
--syncmode "full" \
--unlock $PUBLIC_KEY \
--password $PASSWORD_FILE \
--allow-insecure-unlock \
--rpc \
--rpcaddr "0.0.0.0" \
--rpcport 8545 \
--rpcapi "eth,admin,txpool,net,web3,personal,miner,debug" \
--rpccorsdomain "*" \
--ws \
--wsaddr "0.0.0.0" \
--wsport 8546 \
--wsorigins "*" \
--mine \
--miner.threads 2 \
--pprof \
--pprofaddr "0.0.0.0" \
--pprofport 6060 \
--metrics \
--metrics.expensive \
--verbosity 0 \
console


# Connect to Remote Signer - Ethereum Node
geth \
--networkid $ETH_NETWORK_ID \
--datadir $DATADIR \
--jspath ./scripts/ \
--syncmode "full" \
--signer $REMOTE_SIGNER_ETH_NODE \
--unlock $PUBLIC_KEY \
--password $PASSWORD_FILE \
--allow-insecure-unlock \
--rpc \
--rpcaddr "0.0.0.0" \
--rpcport 8545 \
--rpcapi "eth,admin,txpool,net,web3,personal,miner,debug" \
--rpccorsdomain "*" \
--ws \
--wsaddr "0.0.0.0" \
--wsport 8546 \
--wsorigins "*" \
--mine \
--miner.threads 2 \
--pprof \
--pprofaddr "0.0.0.0" \
--pprofport 6060 \
--metrics \
--metrics.expensive \
--verbosity 0 \
console
