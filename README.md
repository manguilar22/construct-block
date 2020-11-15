# ETH Blockchain 

### Dependencies 

* solc 
* geth 
* ~~abigen~~ 
* ~~puppeth~~ 
* ~~bootnode~~

### Starting Ethereum Node 

* config.sh
    * Environment Variables  
* priv.sh 
    * Configure Ethereum Node  
* start.sh 
    * Deploy Ethereum Node

#### Signing Address 
 
```
export PUBLIC_KEY="0x..."
```

#### Start Network 

``` 
geth --networkid 2020
--datadir ./node1/
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
console 2>> Eth.log
```

#### Ethereum Console  

```
> loadScript("./GethFunc.js")
```

### Bootnode 

#### Generate Bootkey 

```bash
bootnode -genkey boot.key 
``` 

#### Deploy Bootnode 


```bash 
bootnode -nodekey boot.key -verbosity 7 -addr "127.0.0.1:30310"
```
