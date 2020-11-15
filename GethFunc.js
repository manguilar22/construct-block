function checkAllBalances() {
    var totalBal = 0;
    for (var acctNum in eth.accounts) {
        var acct = eth.accounts[acctNum];
        var acctBal = web3.fromWei(eth.getBalance(acct), "ether");
        totalBal += parseFloat(acctBal);
        console.log("eth.accounts[" + acctNum + "]: \t" + acct + " \tbalance: " + acctBal + " ether");
    }
    console.log("Total balance: " + totalBal + " ether");
}

function unlockAllAccounts() {
    for(var e in eth.accounts) {
        var accountStatus = personal.unlockAccount(eth.accounts[e],"123456");
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