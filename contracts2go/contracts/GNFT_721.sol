// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721Burnable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import {ERC721Royalty} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Royalty.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract GNFT_721 is ERC721, ERC721Burnable, ERC721Royalty, ERC721Holder,Ownable {

    constructor(string memory name, string memory symbol, address royaltyRecipient, uint96 royaltyValue)
    ERC721(name, symbol) {
        // 设置默认的版税接收者和比例
        _setDefaultRoyalty(royaltyRecipient, royaltyValue);
    }

    // 覆写 supportsInterface 来处理 ERC721 和 ERC721Royalty
    function supportsInterface(bytes4 interfaceId)
    public
    view
    virtual
    override(ERC721, ERC721Royalty)
    returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    // 铸造新的代币并设置版税
    function mintWithRoyalty(
        address recipient,
        uint256 tokenId,
        address royaltyRecipient,
        uint96 royaltyValue
    ) public {
        _mint(recipient, tokenId);
        _setTokenRoyalty(tokenId, royaltyRecipient, royaltyValue); // 为特定代币设置版税

    }

    // 重写 approve 函数,把 tokenId 全部授权给 to 这个账户
    function approve(address to, uint256 tokenId) public override(ERC721, ERC721Royalty) {
        address owner = ownerOf(tokenId);
        require(to != owner, "ERC721: approval to current owner");
        require(msg.sender == owner || isApprovedForAll(owner, msg.sender), "ERC721: approve caller is not owner nor approved for all");
        super.approve(to, tokenId);
    }

    // 重写 getApproved 函数,查看授权情况
    function getApproved(address to, uint256 tokenId) public view override(ERC721, ERC721Royalty) returns (address) {
        return super.getApproved(tokenId);
    }

    //判断 tokenId 是否存在，若有归属则一定存在
    function _exists(uint256 tokenId) internal view override(ERC721, ERC721Royalty) returns (bool) {
        address owner = _owners[tokenId];
        return owner != address(0);
    }

    // 重写 setApprovalForAll 函数
    function setApprovalForAll(address operator, bool approved) public override(ERC721, ERC721Royalty) {
        require(operator != address(0), "ERC721: approve to the zero address");
        require(operator != msg.sender, "ERC721: approve to caller");
        super.setApprovalForAll(operator, approved);
    }

    // 重写 isApprovedForAll 函数,判断 operator 是否被授权
    function isApprovedForAll(address owner, address operator) public view override(ERC721, ERC721Royalty) returns (bool) {
        return super.isApprovedForAll(owner, operator);
    }

    // 重写 _isApprovedOrOwner 函数,判断 operator 是否被授权或是拥有者
    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view override(ERC721, ERC721Royalty) returns (bool) {
        require(_exists(tokenId), "ERC721: operator query for nonexistent token");
        address owner = ownerOf(tokenId);
        return (spender == owner || getApproved(tokenId) == spender || isApprovedForAll(owner, spender));
    }


    // transferFrom 实现转账
    function transferFrom(address from,address to,uint256 tokenId) public override(ERC721, ERC721Royalty) {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        _transfer(from, to, tokenId);
    }

    // safeTransferFrom 实现安全转账
    function safeTransferFrom(address from,address to,uint256 tokenId,bytes memory _data) public override(ERC721, ERC721Royalty) {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "ERC721: transfer caller is not owner nor approved");
        _safeTransfer(from, to, tokenId, _data);
    }

}

