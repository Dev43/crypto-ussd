pragma solidity ^0.5.2;


// need ERC20 interactor
// need ERC721 interactor
// maybe MakerDao too

import 'openzeppelin-solidity/contracts/token/ERC20/IERC20.sol';
import "./User.sol";
import 'openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract Telco is Ownable {

  // mapping from User contract address to user
  mapping(address => OurUser) users;
  // mapping from phone number to user address
  mapping(string => address) fromPhoneNumberToAddress;

  struct OurUser {
    string phoneNumber;
    // keeps token balances
    mapping(address => uint256) balances;
  }

  constructor() public {
  }

  // Register is how the users register from their smart contract
  function register(string calldata _phoneNumber) external returns(bool) {
    // msg.sender is the user's contract here
    users[msg.sender] = OurUser({
      phoneNumber: _phoneNumber
    });
    fromPhoneNumberToAddress[_phoneNumber] = msg.sender;
    return true;
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
  // maybe a counter for locked balances in channels + all channels that belong to the user


  function userDirectWithdraw(address _tokenAddress, uint256 _amount) external {
    OurUser storage user = users[msg.sender];
    require(user.balances[_tokenAddress] >= _amount, "not enough funds to withdraw");
    IERC20 erc20 = IERC20(_tokenAddress);
    erc20.transfer(msg.sender, _amount);
  }

  function userInDirectWithdraw(address _tokenAddress, uint256 _amount, string calldata _phoneNumber, bytes calldata _password) external {
    require(checkPasswordValid(_phoneNumber, _password), "password not valid");
    address userAddress = fromPhoneNumberToAddress[_phoneNumber];
    IERC20 erc20 = IERC20(_tokenAddress);
    erc20.transfer(userAddress, _amount);
    // TODO need to remove the password here!
  }
}
