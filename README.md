# Ethereum Blockchain Development  

### Environment Variables 

```
# Ethereum Config
export VERSION="0"                      # 0=host, 1=HTTP
## Wallet
export DATADIR=""                       # Ethereum Keystore
export KEYSTORE=$DATADIR""              # Keystore File Directory
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
```

### REST-API 

```
~$ go run node.go 
```

* / 
* /blockchain 
* /accounts 
* /accounts/id/*:num* 
* /accounts/*:account* 
* /accounts/*:account*/balance 
* /accounts/balance 
* /eth 
* /eth/blocks 
* /eth/blocks/tx 

### Creating an Ethereum Account 

```bash 
geth --datadir ./myNode1 new account  
```

**Save Public Key** 

#### Generate Genesis Block

```bash 
puppeth
```

**Consensus Algorithm**: Cleque

#### Initialize Genesis Block 

```bash 
geth --datadir ./myNode1 init genesis_block.json 
```

### Ethereum Node Setup 

* scripts/config.sh
    * Environment variables  
* scripts/startU18.sh 
    * Start Ethereum node on Ubuntu 
* scripts/startBSD.sh 
    * Start Ethereum node on FreeBSD 
* scripts/GethFunc.js 
    * Use Ethereum management API's to perform operations on the blockchain.

#### Signing Address 
 
```
export PUBLIC_KEY="0x..."
```

#### Ethereum Node Options 

*construct-block/scripts/startU18.sh* 

``` 
geth 
--networkid 2020
--datadir ./node1/
--preload ./scripts/GethFunc.js
--http
--http.addr "0.0.0.0" 
--http.port 8545
--rpccorsdomain "*"
--rpcaddr "0.0.0.0"
--rpcapi "eth,admin,rpc,txpool,net,web3,personal,miner,debug"
--unlock $PUBLIC_KEY
--password $PASSWORD_FILE
--verbosity 7
--syncmode "full"
--allow-insecure-unlock
--mine
--miner.threads 8
--metrics                                                 # http://localhost:6060/debug/metrics 
--metrics.expensive                                       # http://localhost:6060/debug/metrics/prometheus
--metrics.addr "0.0.0.0" 
--pprof                                                   # http://localhost:6060/debug/pprof  
--pprof.addr "0.0.0.0"                                    
console 2>> Eth.log
```



