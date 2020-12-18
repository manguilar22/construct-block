function OnSignerStartup() {
    return "Approve"
}
function OnApprovedTx() {
    return "Approve"
}
function ApproveListing() {
    return "Approve"
}
function ApproveTx() {
    return "Approve"
}
function ApproveSignData() {
    return "Approve"
}

// sha256sum rules.js
// clef --configdir $DATADIR attest <HASH>

// * Make CLEF masterseed.json
// clef init --confidir $DATADIR
// cat $DATADIR/masterseed.json | jq  (Copy Ciphertext > signer-password.txt)
// sha256sum rules.js >> signer-password.txt