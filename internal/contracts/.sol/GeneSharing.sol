// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract GeneSharing {

    constructor(string memory sharingID_, address proxy_, bool isOfficial_){
        _sharingID = sharingID_;
        _proxy = proxy_;
        _isOfficial = isOfficial_;
        _creator = msg.sender;
        _createTime = block.timestamp;
    }
    //sharingID
    string private _sharingID;

    //代理
    address private _proxy;

    // 创建者
    address private _creator;

    // 创建时间
    uint private _createTime;

    // 合集中的Metadatas
    mapping(bytes32 => bool) private _MetadataKeyMapping;

    // Metadata列表
    bytes32[] private _MetadataKeyArray;

    // 是否通过官方授权的第三方构建
    bool private _isOfficial;

    // 代理约束
    modifier onlyProxy() {
        require(msg.sender == _proxy, "Only proxy allowed");
        _;
    }

    function sharingID() public view returns (string memory) {
        return _sharingID;
    }

    function creator() public view returns (address) {
        return _creator;
    }

    function isOfficial() public view returns (bool) {
        return _isOfficial;
    }

    //判断该GeneSharing是否包含该Metadata
    function isMetadataIn(bytes32 dataHash) public view returns (bool) {
        return _MetadataKeyMapping[dataHash];
    }

    //添加Metadata
    function addMetadata(bytes32 dataHash) external onlyProxy {
        _MetadataKeyMapping[dataHash] = true;
        _MetadataKeyArray.push(dataHash);
    }

}
