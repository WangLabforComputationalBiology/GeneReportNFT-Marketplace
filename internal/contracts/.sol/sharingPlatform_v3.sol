// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import {GeneSharing} from "./GeneSharing_v3.sol";
import {Metadata} from "./Metadata.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract SharingPlatform is ERC20 {

    //管理员
    address private _admin;

    // Metadata合约地址
    address private _metadataContract;

    // GeneSharing合约地址
    address[] private _geneSharingContracts;

    // sharingID 到合约地址的映射
    mapping(address => bool) private _geneSharingContract;

    // 用户创建的GeneSharing的数目
    mapping(address => uint256) private _geneSharingAmount;

    // 机构认证 KYC 状态
    mapping(address => bool) private _isAuthed;

    // 新建GeneSharing合约事件
    event GeneSharingCreated(address indexed contractAddress, address creator, bool isOfficial);

    // 代理调用事件
    event ProxyCallExecuted(address indexed user, address indexed target, bytes data);

    // 从第三方构建事件
    event CreateAllFromThirdParty(address indexed user, address indexed targetGeneSharing);

    // 管理员约束
    modifier onlyAdmin() {
        require(msg.sender == _admin, "Only admin");
        _;
    }

    constructor(uint initialSupply) ERC20("GeneCoin", "GNC") payable {
        _admin = msg.sender;
        _mint(address(this), initialSupply * 10 ** 18);
    }

//    function _authorizeUpgrade(address newImplementation) internal override onlyAdmin {
//
//    }

    // 设置Metadata合约地址
    function setMetadataContract(address metadataContract) external onlyAdmin {
        _metadataContract = metadataContract;
    }

    // 获取Metadata合约地址
    function MetadataContract() external view returns (address) {
        return _metadataContract;
    }

    // 更新 KYC 状态
    function setUserAuthStatus(address user) external onlyAdmin {
        _isAuthed[user] = true;
    }

    //创建GeneSharing及Metadata（来源：第三方官方）
    function createAllFromThirdParty(address user, bytes32[] calldata dataHashs) external onlyAdmin returns (address) {

        GeneSharing newSharing = new GeneSharing(address(this), true);
        address contractAddress = address(newSharing);

        _geneSharingContracts.push(contractAddress);
        _geneSharingContract[contractAddress] = true;
        _geneSharingAmount[user]++;

        emit GeneSharingCreated(contractAddress, msg.sender, true);

        Metadata metadata = Metadata(_metadataContract);
        for (uint i = 0; i < dataHashs.length; i++) {
            metadata.newMetadata(dataHashs[i], msg.sender);
            require(metadata.isMetadataExist(dataHashs[i]), "Metadata not exist");
            newSharing.addMetadata(dataHashs[i]);
        }
        //奖励msg.sender
        _transfer(address(this), user, 1 * 10 ** 18);

        emit CreateAllFromThirdParty(user, address(newSharing));
        return contractAddress;
    }

    function createEmptyGeneSharingFromCreativeWorkSpace() external returns (address) {
        GeneSharing newSharing = new GeneSharing(address(this), false);
        address contractAddress = address(newSharing);
        _geneSharingContracts.push(contractAddress);
        _geneSharingContract[contractAddress] = true;
        emit GeneSharingCreated(contractAddress, msg.sender, false);
        _geneSharingAmount[msg.sender]++;
        return contractAddress;
    }

    // 添加多个Metadata
    function addMetadataBatchesFromCreativeWorkSpace(address geneSharingAddress, bytes32[] calldata dataHash) external {
        require(_isAuthed[msg.sender], "KYC not verified");
        require(_geneSharingContract[geneSharingAddress] == true, "Sharing not found");
        GeneSharing geneSharing = GeneSharing(geneSharingAddress);
        Metadata metadata = Metadata(_metadataContract);
        require(geneSharing.isOfficial() == false, "Official Sharing can't add Metadata");
        for (uint i = 0; i < dataHash.length; i++) {
            require(Metadata(_metadataContract).isMetadataExist(dataHash[i]), "Metadata Must exists");
            geneSharing.addMetadata(dataHash[i]);
            metadata.addToGeneSharing(dataHash[i], msg.sender, address(geneSharing));
        }

    }

    // 更新共享状态
    function updateMetadataSharingStatus(bytes32 dataHash, bool status) external {
        Metadata metadata = Metadata(_metadataContract);
        require(metadata.isMetadataExist(dataHash), "Metadata Must exists");
        require(metadata.owner(dataHash) == msg.sender, "Not the owner");
        metadata.updateSharingStatus(dataHash, status);
    }

    //新增查看权限
    function addViewAccess(address geneSharingAddress, bytes32 dataHash, string calldata remark) external {
        require(_geneSharingContract[geneSharingAddress] == true, "Sharing not found");
        GeneSharing geneSharing = GeneSharing(geneSharingAddress);
        Metadata metadata = Metadata(_metadataContract);
        require(geneSharing.isMetadataIn(dataHash), "Metadata Must In this GeneSharing");
        metadata.addViewAccess(dataHash, msg.sender, address(geneSharing), remark);

        //奖励creator
        _transfer(address(this), geneSharing.creator(), 1 * 10 ** 18);
        //奖励owner
        _transfer(address(this), metadata.owner(dataHash), 1 * 10 ** 18);

    }

    // 续约查看权限
    function renewalViewAccess(bytes32 dataHash, string calldata remark) external {
        Metadata metadata = Metadata(_metadataContract);
        //记录已续约次数
        uint256 renewalCount = metadata.getRenewalCount(dataHash, msg.sender);
        metadata.renewalViewAccess(dataHash, msg.sender, remark);

        //奖励owner
        uint256 reward = 1 * 10 ** 18 / (renewalCount ** 2);
        if (reward == 0) {
            _transfer(address(this), metadata.owner(dataHash), 1);
        } else {
            _transfer(address(this), metadata.owner(dataHash), reward);
        }
    }

    // 验证查看权限
    function verifyViewAccess(bytes32 dataHash) external view onlyAdmin returns (bool) {
        Metadata metadata = Metadata(_metadataContract);
        if (msg.sender == metadata.owner(dataHash)) {
            return true;
        }

        return metadata.verifyViewAccess(dataHash, msg.sender);
    }


}
