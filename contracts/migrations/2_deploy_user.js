const User = artifacts.require("User");
const Telco = artifacts.require("Telco");
const ERC20 = artifacts.require("ERC20");

module.exports = function(deployer) {
  // Use deployer to state migration tasks.
  console.log("Deploying...")

  let passwords = ["hi", "hello", "test", "what", "where", "when", "this", "awesome"]
  let hashed = []
  for(password of passwords) {
      hashed.push(web3.utils.sha3(web3.utils.asciiToHex(password)))
  }
  console.log(hashed)
  
	deployer.deploy(Telco).then(async () => {
		return deployer.deploy(User, "254777777777", "Bob", Telco.address, hashed);
	}).then(async () => {
    return deployer.deploy(ERC20, "TEST", "TEST", 1, User.address, "100000000")
	}).then(async () => {
    console.log("Deployed with addresses")
    console.log("User ", User.address)
    console.log("Telco ", Telco.address)
    console.log("ERC20 ", ERC20.address)
	});
};
