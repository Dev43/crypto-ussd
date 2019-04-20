pragma solidity ^0.5.2;


// need ERC20 interactor
// need ERC721 interactor
// maybe MakerDao too

import './IERC20.sol';
import './Ownable.sol';

contract User is Ownable {

    string public phoneNumber;
    string public name;
    uint256 public totalPasswordsLeft;

    // keeps one time passwords
    mapping(bytes32 => bool) public passwords;

    mapping(address => bool) public trustedEntities;

  constructor(string memory _phoneNumber, string memory _name, address _trustedEntity, bytes32[] memory _hashedPasswords) public {
      phoneNumber = _phoneNumber;
      name = _name;
      addPasswords(_hashedPasswords);
      addTrustedEntity(_trustedEntity);
  }

  function addPasswords(bytes32[] memory _hashedPasswords) public onlyOwner {
    for(uint256 i = 0; i < _hashedPasswords.length; i++) {
      passwords[_hashedPasswords[i]] = true;
    }
    totalPasswordsLeft += _hashedPasswords.length;
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
      totalPasswordsLeft -= 1;
      address(uint160(_to)).transfer(_amount);
  }

  function retirePassword(bytes memory _password) public {
    delete passwords[keccak256(_password)];
    totalPasswordsLeft -= 1;
  }

  function trustedEntityTransferERC20(address _to, address _erc20Address, uint256 _amount, bytes memory _password) public  {
      require(trustedEntities[msg.sender], "not trusted entity");
      require(passwords[keccak256(_password)], "password is not valid");
      // delete the password
      delete passwords[keccak256(_password)];
      totalPasswordsLeft -= 1;

      IERC20 erc20 = IERC20(_erc20Address);
      require(erc20.transfer(_to, _amount), "problem sending ERC20");
  }

}
