// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {GNC_20} from "GNC_20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {ERC721Burnable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import {ERC721Royalty} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Royalty.sol";
import {ERC2981} from "@openzeppelin/contracts/token/common/ERC2981.sol";

contract GNFT_721 is ERC721Burnable, ERC721Royalty, Ownable {
    GNC_20 public GNC;
    uint256 public rewardAmount; // 每次铸造时奖励的 GNC 数量

    // 事件：签名交易
    event MetaTransactionExecuted(address userAddress, address relayerAddress);
    event Minted(address indexed recipient, uint256 tokenId);

    // 构造函数
    constructor(address _gncAddress, uint256 _rewardAmount) ERC721("GNFT_721", "GNFT") Ownable(msg.sender) payable {
        GNC = GNC_20(_gncAddress);
        rewardAmount = _rewardAmount;
    }

    // 设置奖励数量的函数，只有合约拥有者可以调用
    function setRewardAmount(uint256 newAmount) external onlyOwner {
        rewardAmount = newAmount;
    }

    // 铸造新 NFT 并设置版税，同时给予 GNC 奖励
    function mintAndReward(address recipient, uint256 tokenId, uint96 royaltyValue) external onlyOwner {
        // 铸造 NFT 并设置版税
        _mint(recipient, tokenId);
        _setTokenRoyalty(tokenId, recipient, royaltyValue);

        // 转移奖励给用户
        require(GNC.transfer(recipient, rewardAmount), "Reward transfer failed");
    }

    // 重写 supportsInterface 函数，支持 ERC721 和 ERC2981
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721, ERC721Royalty) returns (bool) {
        return super.supportsInterface(interfaceId);
    }

    // 重写 _burn 函数，以支持版税清除
    function burn(uint256 tokenId) public virtual override(ERC721Burnable) {
        super.burn(tokenId);
        _resetTokenRoyalty(tokenId);
    }

    // 设置版税分母，用于计算版税比例
    function _feeDenominator() internal pure virtual override(ERC2981) returns (uint96) {
        return 1000; // 版税为百分比，分母为 1000 表示最高 10%
    }

    // 重写 approve 函数，把 tokenId 全部授权给 to 这个账户
    function approve(address to, uint256 tokenId) public override(ERC721) {
        address owner = ownerOf(tokenId);
        require(to != owner, "ERC721: approval to current owner");
        require(msg.sender == owner || isApprovedForAll(owner, msg.sender), "ERC721: approve caller is not owner nor approved for all");
        super.approve(to, tokenId);
    }

    // 签名验证：确保交易由用户签名并且未被篡改
    function isValidSignature(
        address userAddress,
        uint256 tokenId,
        uint256 price,
        bytes memory signature
    ) public view returns (bool) {
        bytes32 hash = getTransactionHash(userAddress, tokenId, price);
        bytes32 messageHash = ECDSA.toEthSignedMessageHash(hash);
        address signer = ECDSA.recover(messageHash, signature);
        return signer == userAddress;
    }

    // 生成交易哈希
    function getTransactionHash(
        address userAddress,
        uint256 tokenId,
        uint256 price
    ) public view returns (bytes32) {
        return keccak256(abi.encodePacked(userAddress, tokenId, price, nonces[userAddress]++));
    }

    // 执行 MetaTransaction，用户签署交易后，Relayer 调用此函数代为执行
    function executeMetaTransaction(
        address userAddress,
        uint256 tokenId,
        uint256 price,
        bytes memory signature
    ) public payable {
        // 验证签名
        require(isValidSignature(userAddress, tokenId, price, signature), "Invalid Signature");

        // 执行铸造操作（可以根据需求修改）
        _mint(userAddress, tokenId);

        // 支付 GNC 代币奖励（根据需求调整）
        GNC.transfer(userAddress, price);

        // 触发 MetaTransaction 执行事件
        emit MetaTransactionExecuted(userAddress, msg.sender);
    }

    // 判断 tokenId 是否存在，若有归属则一定存在
    function _exists(uint256 tokenId) internal view returns (bool) {
        address owner = ownerOf(tokenId);
        return owner != address(0);
    }

    // 重写 setApprovalForAll 函数
    function setApprovalForAll(address operator, bool approved) public override(ERC721) {
        require(operator != address(0), "ERC721: approve to the zero address");
        require(operator != msg.sender, "ERC721: approve to caller");
        super.setApprovalForAll(operator, approved);
    }

    // 重写 _isApprovedOrOwner 函数，判断 spender 是否被授权或是拥有者
    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view returns (bool) {
        require(_exists(tokenId), "ERC721: operator query for nonexistent token");
        address owner = ownerOf(tokenId);
        return (spender == owner || super.getApproved(tokenId) == spender || isApprovedForAll(owner, spender));
    }

    // transferFrom 实现转账
    function transferFrom(address from, address to, uint256 tokenId) public override(ERC721) {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        _transfer(from, to, tokenId);
    }

    // safeTransferFrom 实现安全转账
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public override(ERC721) {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        _safeTransfer(from, to, tokenId, _data);
    }
}
