// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ToDoToken is ERC20, Ownable {

    mapping(address => bool) blacklist;

    constructor() ERC20("ToDo Token", "TDT") {
        _mint(msg.sender, 10**11 * 10**18);
    }

    function checkBlackList(address target) public view returns (bool) {
        return blacklist[target];
    }

    function setBlackList(address target, bool value) public onlyOwner {
        blacklist[target] = value;
    }

    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        require(blacklist[_msgSender()] == false, "TRANSFER: sender is in the blacklist");
        require(blacklist[recipient] == false, "TRANSFER: recipient is in the blacklist");
        _transfer(_msgSender(), recipient, amount);
        return true;
    }
}
