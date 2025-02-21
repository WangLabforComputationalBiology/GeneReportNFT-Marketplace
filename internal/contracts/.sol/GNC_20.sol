// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract GNC_20 is ERC20, ERC20Permit {
    constructor(uint256 initialSupply) ERC20("GeneCoin", "GNC") ERC20Permit("GeneCoin")  payable {
        _mint(msg.sender, initialSupply);
    }
}
