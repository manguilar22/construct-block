#!/bin/bash

# Ethereum Config
export VERSION="0"                      # 0=host, 1=HTTP
## Wallet
export DATADIR=""                       # Ethereum Keystore
export KEYSTORE=""                      # Keystore File Directory
## Networking
export ETH_NETWORK_ID=""                # Chain ID
export HOSTNAME=""                      # HTTP, Web Socket, RPC
## Application Credentials
export PUBLIC_KEY=""                    # Signing Hex Address
export PRIVATE_KEY=""                   # Private Key Hex Address
export PASSWORD_FILE=""                 # KeyStore Password (Unlock Account)

# Ethereum - Swarm Config
export $SWARM_NODE_NAME=""
export IPC_FILE=""
export SWARM_DATADIR=""
export SWARM_ACCOUNT=""
export SWARM_ACCOUNT_PASSWORD=""