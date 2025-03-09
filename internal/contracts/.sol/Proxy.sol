// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract Proxy is Ownable {
    address public GNFT_721_implementation;
    address public GNC_20_implementation;
    address public Marketplace_implementation;
    address public MetaTransaction_implementation;

    event Upgraded(address indexed newImplementation, string contractType);

    constructor(
        address _GNFT_721,
        address _GNC_20,
        address _Marketplace,
        address _MetaTransaction
    ) {
        GNFT_721_implementation = _GNFT_721;
        GNC_20_implementation = _GNC_20;
        Marketplace_implementation = _Marketplace;
        MetaTransaction_implementation = _MetaTransaction;
    }

    // 升级合约
    function upgradeTo(address newImplementation, string calldata contractType) external onlyOwner {
        require(newImplementation != address(0), "Invalid address");

        if (keccak256(bytes(contractType)) == keccak256(bytes("GNFT_721"))) {
            GNFT_721_implementation = newImplementation;
        } else if (keccak256(bytes(contractType)) == keccak256(bytes("GNC_20"))) {
            GNC_20_implementation = newImplementation;
        } else if (keccak256(bytes(contractType)) == keccak256(bytes("Marketplace"))) {
            Marketplace_implementation = newImplementation;
        } else if (keccak256(bytes(contractType)) == keccak256(bytes("MetaTransaction"))) {
            MetaTransaction_implementation = newImplementation;
        } else {
            revert("Unknown contract type");
        }

        emit Upgraded(newImplementation, contractType);
    }

    // 代理所有调用
    fallback() external payable {
        address impl = _getImplementation(msg.sig);
        require(impl != address(0), "Implementation not found");

        _delegate(impl);
    }

    // 根据函数签名选择实现合约
    function _getImplementation(bytes4 sig) internal view returns (address) {
        if (sig == bytes4(keccak256("mint(address)"))) {
            return GNFT_721_implementation;
        } else if (sig == bytes4(keccak256("transfer(address,uint256)"))) {
            return GNC_20_implementation;
        } else if (sig == bytes4(keccak256("list(address,uint256,uint256)"))) {
            return Marketplace_implementation;
        } else if (sig == bytes4(keccak256("executeMetaTransaction(address,uint256,bytes)"))) {
            return MetaTransaction_implementation;
        } else {
            return address(0);
        }
    }

    // 转发调用到实现合约
    function _delegate(address impl) internal {
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), impl, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 { revert(0, returndatasize()) }
            default { return(0, returndatasize()) }
        }
    }
}
