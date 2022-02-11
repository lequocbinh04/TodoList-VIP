// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TodoClaim is Ownable {

    ERC20 TokenClaim;

    constructor() {
        TokenClaim = ERC20(0x708C0BC45208776B32382983561ce9d38bD5F778);
    }

    function claimToken(address creator, uint256 amount) external onlyOwner {
        require(amount > 0, "CLAIM: Invalid amount");
        require(amount < TokenClaim.balanceOf(address(this)), "CLAIM: Invalid amount");
        require(TokenClaim.transfer(creator, amount), "CLAIM: Failed to transfer");
    }

    function tokenClaimAddress() public view returns (ERC20) {
        return TokenClaim;
    }

    function setTokenClaim(ERC20 new_token) external onlyOwner {
        TokenClaim = new_token;
    }

    function endClaim() external onlyOwner {
        address admin = owner();        
        TokenClaim.transfer(admin, TokenClaim.balanceOf(address(this)));
    }

}
