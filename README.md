# ETH Blockchain 

### Dependencies 

* solc 
* geth 
* ~~abigen~~ 
* ~~puppeth~~ 

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
--password ./password.txt
--verbosity 7
--syncmode "full"
--allow-insecure-unlock
--mine
--miner.threads 8
console 2>> Eth.log
```

#### Console  

```
> loadScript("./GethFunc.js")
```