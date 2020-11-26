//  JS modules: admin:1.0 clique:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

function makeBallots(accountPassword, accountNumber) {
    for (var i = 0; i < accountNumber; i++) {
        personal.newAccount(accountPassword);
    }
}

function makeVote(weiValue, weiUnit) {
    // 1 * 10^28 wei = 1 ether
    var genesis = eth.accounts[0];
    var l = eth.accounts.length;
    for (var i = 1; i < l; i++) {
        console.log("Sending wei from genesis block (" + genesis + ") to account " + eth.accounts[i]);
        web3.eth.sendTransaction({from:genesis,to: eth.accounts[i], value: web3.toWei(weiValue,weiUnit)})
    }
}

function checkAllBalances(weiUnit) {
    var totalBal = 0;
    for (var acctNum in eth.accounts) {
        var acct = eth.accounts[acctNum];
        var acctBal = web3.fromWei(eth.getBalance(acct), weiUnit);
        totalBal += parseFloat(acctBal);
        console.log("eth.accounts[" + acctNum + "]: \t" + acct + " \tbalance: " + acctBal + " " + weiUnit);
    }
    console.log("Total balance: " + totalBal + " ether");
}

function unlockAllAccounts(password) {
    for(var e in eth.accounts) {
        var accountStatus = personal.unlockAccount(eth.accounts[e],password);
        console.log(accountStatus);
    }
}

function unlockAllAccounts(passwordString) {
    for(var e in eth.accounts) {
        var accountStatus = personal.unlockAccount(eth.accounts[e],passwordString);
        console.log(accountStatus);
    }
}

function sendTransaction(fromAccountHex, toAccountHex, weiValue, weiUnit) {
    eth.sendTransaction({
        from:fromAccountHex,
        to:toAccountHex,
        value: web3.toWei(weiValue, weiUnit)
    })
}

function getWei(accountHex) {
    var currentWei = eth.getBalance(accountHex);
    return currentWei
}

function checkBalance(accountHex, weiUnit) {
    return web3.fromWei(eth.getBalance(accountHex), weiUnit)
}