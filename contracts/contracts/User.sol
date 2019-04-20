pragma solidity ^0.5.2;


// need ERC20 interactor
// need ERC721 interactor
// maybe MakerDao too

import 'openzeppelin-solidity/contracts/token/ERC20/IERC20.sol';
import 'openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract User is Ownable {

    string phoneNumber;
    string name;

    // keeps one time passwords
    mapping(bytes32 => bool) passwords;
    // keeps token balances
    mapping(bytes32 => uint256) balances;

    mapping(address => bool) trustedEntities;

  constructor(string memory _phoneNumber, string memory _name) public {
      phoneNumber = _phoneNumber;
      name = _name;
  }

  function addPasswords(bytes32[] calldata _hashedPasswords) external onlyOwner {
    for(uint256 i = 0; i < _hashedPasswords.length; i++) {
      passwords[_hashedPasswords[i]] = true;
    }
  }

  function isPasswordValid( bytes memory _password) public view returns(bool) {
    return passwords[keccak256(_password)];
  }

  function transferEth(address _to, uint256 _amount) public payable onlyOwner {
      address(uint160(_to)).transfer(_amount);
  }

  function transferERC20(address _to, address _erc20Address, uint256 _amount) public onlyOwner {
      IERC20 erc20 = IERC20(_erc20Address);
      require(erc20.transfer(_to, _amount), "problem sending ERC20");
  }

  function addTrustedEntity(address _trustedEntity) public onlyOwner {
      trustedEntities[_trustedEntity] = true;
  }

  function removeTrustedEntity(address _trustedEntity) public onlyOwner {
      trustedEntities[_trustedEntity] = false;
  }

  function trustedEntityTransferEth(address _to, uint256 _amount, bytes memory _password) public payable  {
      require(trustedEntities[msg.sender], "not trusted entity");
      require(passwords[keccak256(_password)], "password is not valid");
      // delete the password
      delete passwords[keccak256(_password)];
      address(uint160(_to)).transfer(_amount);
  }

  function trustedEntityTransferERC20(address _to, address _erc20Address, uint256 _amount, bytes memory _password) public  {
      require(trustedEntities[msg.sender], "not trusted entity");
      require(passwords[keccak256(_password)], "password is not valid");
      // delete the password
      delete passwords[keccak256(_password)];

      IERC20 erc20 = IERC20(_erc20Address);
      require(erc20.transfer(_to, _amount), "problem sending ERC20");
  }

}

// A user needs the ability to:
// Add passwords
// Change Ownership
// Approve a telco?
// Withdraw funds from our telco
// Receive tokens and move them
// Allow a trusted telco to transfer tokens (erc20 or other) with a password
