// Script file

// Page elements
let walletAmount = document.getElementById('wallet_amount');
let reloadWalletButton = document.getElementById('reload_wallet');
let reloadWalletSpinner = document.getElementById('reload_wallet_spinner');
let publicKey = document.getElementById('public_key');
let privateKey = document.getElementById('private_key');
let blockchainAddress = document.getElementById('blockchain_address');
let recipientBlockchainAddress = document.getElementById('recipient_blockchain_address');
let sendAmount = document.getElementById('send_amount');
let sendMoneyButton = document.getElementById('send_money_button');
let sendMoneySpinner = document.getElementById('send_money_spinner');


// Event Subscribing
window.addEventListener('DOMContentLoaded', init);
reloadWalletButton.addEventListener('click', reloadWallet);
recipientBlockchainAddress.addEventListener('keyup', verifySendData);
sendAmount.addEventListener('keyup', verifySendData);
sendMoneyButton.addEventListener('click', sendMoney);


/////////////// Event Handlers //////////////////////////////
function init() {
    loadWalletInfo();
    setInterval(function () {
        updateWallet();
    }, 10000);
}

// loadWalletInfo loads the wallet's data used in transactions
async function loadWalletInfo() {
    let options = {
        method: 'POST',
    }

    try {
        let resp = await fetch('/wallet', options);
        let data = await resp.json();
        if (publicKey.value != data.public_key) {
            fadeInElement(publicKey);
            publicKey.value = data.public_key;
        }
        if (privateKey.value != data.private_key) {
            fadeInElement(privateKey);
            privateKey.value = data.private_key;
        }
        if (blockchainAddress.value != data.blockchain_address) {
            fadeInElement(blockchainAddress);
            blockchainAddress.value = data.blockchain_address;
        }
    } catch (error) {
        console.error(error);
    }
}

// reloadWallet reload the wallet's total amount
async function reloadWallet() {
    reloadWalletButton.disabled = true;
    reloadWalletSpinner.classList.remove('visually-hidden');
    // add delay for more realistic effect
    setTimeout(function () {
        updateWallet();
    }, 1500);
}

// updateWallet gets the wallet's total amount
async function updateWallet() {
    try {
        let address = blockchainAddress.value;
        let resp = await fetch('/wallet/amount?blockchain_address=' + address);
        let data = await resp.json();
        if (walletAmount.innerText != data.amount) {
            fadeInElement(walletAmount);
            walletAmount.innerText = data.amount;
        }
    } catch (error) {
        console.error(error);
    } finally {
        reloadWalletButton.disabled = false;
        reloadWalletSpinner.classList.add('visually-hidden');
    }
}

// verifySendData check if both send fields have data and enables the sendMoneyButton
function verifySendData() {
    if (recipientBlockchainAddress.value != "" && sendAmount.value != "") {
        sendMoneyButton.disabled = false;
    } else {
        sendMoneyButton.disabled = true;
    }
}

// sendMoney confirms send intention and calls sendTransaction
async function sendMoney() {
    let confirmText = 'Are you sure?';
    let confirmResult = confirm(confirmText);
    if (confirmResult !== true) {
        alert('canceled');
        return;
    }

    sendMoneyButton.disabled = true;
    sendMoneySpinner.classList.remove('visually-hidden');
    // add delay for more realistic effect
    setTimeout(function () {
        sendTransaction();
    }, 1500);
}

// sendTransaction sends blockchain transaction
async function sendTransaction() {
    let transactionData = {
        'sender_private_key': privateKey.value,
        'sender_public_key': publicKey.value,
        'sender_blockchain_address': blockchainAddress.value,
        'recipient_blockchain_address': recipientBlockchainAddress.value,
        'value': sendAmount.value,
    }

    try {
        let options = {
            method: 'POST',
            body: JSON.stringify(transactionData),
            headers: {
                'Content-Type': 'application/json'
            }
        }
        let resp = await fetch('/transaction', options);
        let data = await resp.json();
    } catch (error) {
        console.log(error);
    } finally {
        recipientBlockchainAddress.value = '';
        sendAmount.value = '';
        sendMoneyButton.disabled = true;
        sendMoneySpinner.classList.add('visually-hidden');
    }
}

// fadeInElement simulates a fade in using css class fade-in
function fadeInElement(elem) {
    elem.classList.remove('fade-in');
    setTimeout(function () {
        elem.classList.add('fade-in');
    }, 100);
}