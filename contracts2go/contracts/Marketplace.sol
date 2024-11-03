// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import {GNFT_721} from "./GNFT_721.sol";
import {GNC_20} from "./GNC_20.sol";

contract Marketplace {
    GNFT_721 public GNFT;
    GNC_20 public GNC;
    constructor(address _erc20Address, address _erc721Address){
        GNC= GNC_20(_erc20Address);
        GNFT = GNFT_721(_erc721Address);
    }

    }
}
