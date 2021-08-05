import Wallet from 'ethereumjs-wallet'

var Accounts = require('web3-eth-accounts');
var EthUtil = require('ethereumjs-util');

export default {
    generateAccount() {
        const accounts = new Accounts()
        return accounts.create()
    },
    
    getPrivateKeyFromAccount(account) {
        return account.privateKey
    },
    
    getPublicKey(privateKey) {
        if (privateKey === undefined) {
            return
        }
    
        const privateKeyBuffer = EthUtil.toBuffer(privateKey);
        const wallet = Wallet.fromPrivateKey(privateKeyBuffer);
    
        // Get a public key
        const publicKey = wallet.getPublicKeyString(); 
        return publicKey
    },
    
    encryptText(account, safePass) {
        let keystore = account.encrypt(safePass);
        return keystore
    },
    
    signText(account, data) {
        let privateKeyBuffer = EthUtil.toBuffer(account.privateKey)
    
        let msgBuffer = Buffer.from(data);
        let hashBuffer = EthUtil.keccak256(msgBuffer);
    
        let signature = EthUtil.ecsign(hashBuffer, privateKeyBuffer)
        let sigStr = EthUtil.toRpcSig(signature.v, signature.r, signature.s)
        sigStr = sigStr.substr(0, sigStr.length - 2)
    
        return sigStr
    },
    
    getAccountFromKeystore(keystore, safePass) {
        let accounts = new Accounts();
        let account = null
        try {
            account = accounts.decrypt(keystore, safePass)
        } catch(err) {
            throw { name : 'error', message: '金钥验证失败，请输入正确金钥'}
        }
        return account
    },
}
