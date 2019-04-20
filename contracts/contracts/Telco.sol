pragma solidity ^0.5.2;

// THIS IS A HACKATHON CONTRACT >> CARE WAS NOT TAKEN INTO SECURITY CONSIDERATIONS

import './IERC20.sol';
import "./User.sol";
import './Ownable.sol';

contract Telco is Ownable {

  // mapping from User contract address to user
  mapping(address => OurUser) public users;
  // mapping from phone number to user address
  mapping(string => address)  fromPhoneNumberToAddress;

  struct OurUser {
    string phoneNumber;
    // keeps token balances
    mapping(address => uint256) balances;
  }

  constructor() public {
  }

  // Register is how the users register from their smart contract
  function registerNewUser(string calldata _phoneNumber, address userAddress) external returns(bool) {
    users[userAddress] = OurUser({
      phoneNumber: _phoneNumber
    });
    fromPhoneNumberToAddress[_phoneNumber] = userAddress;
    return true;
  }

  function getUserPhoneNumber(address _addr) view public returns(string memory) {
    return users[_addr].phoneNumber;
  }

  function getFromPhoneNumberToAddress(string calldata _phoneNumber) external view returns(address) {
    return fromPhoneNumberToAddress[_phoneNumber];
  }
  
  function checkPasswordValid(string memory _phoneNumber, bytes memory _password) public view returns(bool) {
    address userAddress = fromPhoneNumberToAddress[_phoneNumber];
    require(userAddress != address(0), "user does not exist");

    User userContract = User(userAddress);
    return userContract.isPasswordValid(_password); 
  }

  function getUserTokens(address _tokenAddress, uint256 _amount, string memory _phoneNumber, bytes memory _password) public onlyOwner {
    require(checkPasswordValid(_phoneNumber, _password), "password not valid");
    address userAddress = fromPhoneNumberToAddress[_phoneNumber];
    OurUser storage user = users[userAddress];

    User userContract = User(userAddress);
    userContract.trustedEntityTransferERC20(address(this), _tokenAddress, _amount, _password);
    user.balances[_tokenAddress] += _amount;
  }

  function userDirectWithdraw(address _tokenAddress, uint256 _amount) external {
    OurUser storage user = users[msg.sender];
    require(user.balances[_tokenAddress] >= _amount, "not enough funds to withdraw");
    IERC20 erc20 = IERC20(_tokenAddress);
    user.balances[_tokenAddress] -= _amount;
    erc20.transfer(msg.sender, _amount);
  }

  function userInDirectWithdraw(address _tokenAddress, uint256 _amount, string calldata _phoneNumber, bytes calldata _password) external {
    require(checkPasswordValid(_phoneNumber, _password), "password not valid");
    address userAddress = fromPhoneNumberToAddress[_phoneNumber];
    IERC20 erc20 = IERC20(_tokenAddress);
    users[userAddress].balances[_tokenAddress] -= _amount;
    erc20.transfer(userAddress, _amount);
    // TODO need to remove the password here!
  }

  function getUserBalance(address _addr, address _tokenAddress) public view returns(uint256) {
    return users[_addr].balances[_tokenAddress];
  }

  // add an authenticated transfer  (all it does is transfer tokens internally)
  function authTransfer(address _tokenAddress, string calldata _toPhoneNumber, uint256 _amount, string calldata _phoneNumber, bytes calldata _password) external {
    require(checkPasswordValid(_phoneNumber, _password), "password not valid");
    address userAddress = fromPhoneNumberToAddress[_phoneNumber];
    OurUser storage user = users[userAddress];
    User userContract = User(userAddress);
    userContract.retirePassword(_password);
    require(user.balances[_tokenAddress] >= _amount, "not enough balance");
    // remove from us
    user.balances[_tokenAddress] -= _amount;
    // add to other user
    address toAddress = fromPhoneNumberToAddress[_toPhoneNumber];
    users[toAddress].balances[_tokenAddress] += _amount;
  }

}

// Phase 1
// user creates a smart contract w/ deposit, withdraw, register and passwords
// user registers with the telco
// telco takes coins from the user (adds it to their "balance" mapping) 
// any transfer from the telco in the contract needs the blessing of the user (the password)

// Need to deploy a User contract, get its address 0x4A31ECe693fB614935eFB337034F9C79efEC03B5
  //  Add in all password -- create script to create 20 passwords hashed
// Deploy a Telco contract, get its address 0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11
// ABIGEN the contract (to interact with)
// Create logic in the USSD app for transfers etc





// user requests telco to create a channel with another phone number
  // maybe a counter for locked balances in channels + all channels that belong to the user
// telco creates the channel with himsel
