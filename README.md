# ETH Blockchain 

### Dependencies 

* solc 
* geth 
* ~~abigen~~ 
* ~~puppeth~~ 
* ~~bootnode~~

### Creating an Ethereum Wallet 

```bash 
geth --datadir ./myNode1 new account  
```

**Save Credentials** 

#### Generate Genesis Block (Cleque)

```bash 
puppeth
```

#### Start Genesis Block 

```bash 
geth --datadir ./myNode1 init genesis_block.json 
```

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
geth 
--networkid 2020
--datadir ./node1/
--jspath ./scripts/
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
--pprof.addr="0.0.0.0"                                    
console 2>> Eth.log
```

#### Ethereum Console  

```
> loadScript("./GethFunc.js")
```


