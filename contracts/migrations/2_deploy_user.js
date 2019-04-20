const User = artifacts.require("User");
const Telco = artifacts.require("Telco");

module.exports = function(deployer) {
  // Use deployer to state migration tasks.
  console.log("Deploying...")
	deployer.deploy(User, "254777777777", "Bob").then(async () => {
		return deployer.deploy(Telco);
	}).then(async () => {
    console.log("Deployed with addresses")
    console.log("User ", User.address)
    console.log("Telco ", Telco.address)
	});
};
