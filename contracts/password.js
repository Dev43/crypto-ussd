const Web3 = require('web3');

const web3 = new Web3('http://localhost:8545');

let r = []
for (let i = 2; i < process.argv.length; i++) {
    let arg = process.argv[i]
    let hash = web3.utils.sha3(arg)
    r.push(hash)

}