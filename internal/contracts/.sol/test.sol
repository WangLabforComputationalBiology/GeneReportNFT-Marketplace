// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract Byte32Contract {
    mapping(bytes32 => uint256) private dataStore;

    event DataUpdated(bytes32 indexed key, uint256 value);

    function storeData(bytes32 key, uint256 value) external {
        dataStore[key] = value;
        emit DataUpdated(key, value);
    }

    function getData(bytes32 key) external view returns (uint256) {
        return dataStore[key];
    }
}