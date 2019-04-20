const UserContract = artifacts.require("./User.sol")
const TelcoContract = artifacts.require("./Telco.sol")
const ERC20Contract = artifacts.require("./ERC20.sol")


contract("User", function (accounts) {
    let alice = accounts[0];
    let bob = "0x40058579f9D68ebebBe6E6F45c1995D5143F26AC";
    let userContract
    let telcoContract
    let erc20Contract
    console.log(UserContract.address)
    it("should assert true", async () => {
        userContract = await UserContract.deployed();
        assert(userContract.address, "address not defined");
        telcoContract = await TelcoContract.deployed();
        assert(telcoContract.address, "address not defined");
        assert.isTrue(true);
        erc20Contract = await ERC20Contract.deployed()
        assert(erc20Contract.address, "address not defined");
        assert.isTrue(true);
    });

    // it("should add passwords in user contract", async () => {
    //     let passwords = ["hi", "hello", "test"]
    //     let hashed = []
    //     for(password of passwords) {
    //         hashed.push(web3.utils.sha3(web3.utils.asciiToHex(password)))
    //     }
    //     console.log(hashed)

    //     await userContract.addPasswords(hashed)
    //     assert.equal(await userContract.isPasswordValid(web3.utils.asciiToHex("hi")), true)
    //     try {
    //         assert.equal(await userContract.isPasswordValid(web3.utils.asciiToHex("his")), false)

    //     } catch(e) {
    //         console.log(e)
    //     }
    // });

    // it("should add trusted entity in the contract in user contract", async () => {

    //     await userContract.addTrustedEntity(telcoContract.address)
    //     assert.equal(await userContract.trustedEntities(telcoContract.address), true)
    //     assert.equal(await userContract.trustedEntities("0x0000000000000000000000000000000000000000"), false)
    // });


    // Create an ERC20 token and add it to the contract balance (user already gets it filled)

    it("should register the user in the telco", async () => {

        // now the telco should register the phoneNumber of the user
    try {
            await telcoContract.registerNewUser("254777777777", userContract.address)
        } catch (e) {
            console.log(e)
        }
        // telco gets a portion of the user's token using a password
        assert.equal(await telcoContract.users(userContract.address), "254777777777")
    });

    it("should let the telco get 100 tokens from the user", async () => {

        let balanceBefore = await erc20Contract.balanceOf(userContract.address)
        // now the telco should register the phoneNumber of the user
        try {
            await telcoContract.getUserTokens(erc20Contract.address, "100", "254777777777", web3.utils.asciiToHex("hi"))
        } catch (e) {
            console.log(e)
        }
        let balanceAfter = await erc20Contract.balanceOf(userContract.address)
        console.log(balanceBefore.toString(), balanceAfter.toString())
        // telco gets a portion of the user's token using a password
        assert.equal((balanceBefore.sub(balanceAfter)).toString(), "100")
        let balanceTelco = await erc20Contract.balanceOf(telcoContract.address)
        assert.equal((balanceTelco).toString(), "100")

        // asset it is retired
        assert.equal(await userContract.isPasswordValid(web3.utils.asciiToHex("hi")), false)
        
        // assert the user has 100 tokens
        assert.equal((await telcoContract.getUserBalance(userContract.address, erc20Contract.address)).toString(), "100")
        // assert.equal((await userContract.totalPasswordsLeft()).toString(), "2")
        
    });

    // now authtransfer from the user with a password

    it("should let the telco transfer user tokens with password", async () => {

        // need to register a new address
        let balanceBefore = (await telcoContract.getUserBalance(userContract.address, erc20Contract.address))
        // now the telco should register the phoneNumber of the user

        await telcoContract.registerNewUser("254777777778", bob)
        try {
            await telcoContract.authTransfer(erc20Contract.address,  "254777777778", "50", "254777777777", web3.utils.asciiToHex("hello"))
        } catch (e) {
            console.log(e)
        }
        let balanceAfter = (await telcoContract.getUserBalance(userContract.address, erc20Contract.address))
        
        // assert the user has 50 tokens
        assert.equal((await telcoContract.getUserBalance(userContract.address, erc20Contract.address)).toString(), "50")
        // assert.equal((await userContract.totalPasswordsLeft()).toString(), "1")
        
    });

});