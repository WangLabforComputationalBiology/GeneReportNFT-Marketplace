// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract sharingPlatform  is ERC20 {

    constructor (uint256 initialSupply_, address admin_) ERC20("GeneCoin", "GNC") payable {
        _admin = admin_;
        _mint(address(this), initialSupply_);
    }

    //用户认证状态
    mapping(address => bool) public isAuthed;

    //后台管理员地址
    address private _admin;

    struct GeneSharing {
        // 基因报告id或数据hash，若从wegene构建则为profile_id，若直接新建则为数据hash
        string sharingID;

        // sharing合集创建者
        address creator;

        // 创建时间
        uint256 createTime;

        // 合集中的MetaDatas
        mapping(bytes32 => bool) MetaDataInSharing;

        // MetaData列表
        bytes32[] MetaDataKeys;
    }

    // Sharing hub
    mapping(string => GeneSharing) public GeneSharingRepository;

    //查看许可
    struct ViewAccess {
        //访问者
        address viewer;

        //有效期
        uint256 expiry;

        //是否为有效的许可
        bool isValid;
    }

    // 溯源元数据
    struct TraceData {
        //操作者是通过哪个sharing合约
        string targetSharingID;

        //目标用户
        address operator;

        //追溯事件状态，0：被operator创建，1：operator通过targetSharing查看，2：在工坊被operator引用到targetSharing，3：operator修改共享状态targetSharing
        int16 status;

        //操作时间
        uint timestamp;

        //备注用途
        string remark;
    }

    // MetaData基因元数据
    struct MetaData {
        // 基因数据哈希
        bytes32 dataHash;

        // 所有者
        address owner;

        // 可共享状态
        bool sharingStatus;

        // 默认7天有效期
        uint DURATION;

        // 访问许可
        mapping(address => ViewAccess) viewerPermission;

        // 所有溯源信息
        TraceData[] traceData;
    }

    // MetaData hub
    mapping(bytes32 => MetaData) public MetaDataRepository;

    // 创意工坊资源
    mapping(bytes32 => bool) public CreativeSpaceAsset;

    //新增traceData
    function _newTraceData(bytes32 dataHash, string memory targetSharingID, address operator, int16 status, string memory remark) internal {
        MetaData storage metaData = MetaDataRepository[dataHash];
        //新增traceData
        metaData.traceData.push(TraceData({
            targetSharingID: targetSharingID,
            operator: operator,
            status: status,
            timestamp: block.timestamp,
            remark: remark
        }));
    }

    //新增查看许可
    function _newViewAccess(address viewer, bytes32 dataHash) internal {
        MetaData storage metaData = MetaDataRepository[dataHash];
        metaData.viewerPermission[viewer] = ViewAccess({
            viewer: viewer,
            expiry: block.timestamp + 7 days,
            isValid: true
        });
    }

    //续约查看许可
    function _renewViewAccess(address viewer, bytes32 dataHash) internal {

    }

    // 新建一个Sharing
    function NewSharing(string memory sharingID) external {
        GeneSharing storage sharing = GeneSharingRepository[sharingID];
        sharing.sharingID = sharingID;
        sharing.creator = msg.sender;
        sharing.createTime = block.timestamp;
    }

    // 在Sharing中新建一个MetaData
    function NewMetaData(bytes32 dataHash) external {
        MetaData storage metaData = MetaDataRepository[dataHash];
        metaData.dataHash = dataHash;
        metaData.owner = msg.sender;
        metaData.sharingStatus = true;
        metaData.DURATION = 604800;// 7*24*60*60s

        //新增traceData
        metaData.traceData.push(TraceData({
            targetSharingID: "0",
            operator: msg.sender,
            status: 0,
            timestamp: block.timestamp,
            remark: "be created"
        }));
    }
    // 只允许sharing的creator调用
    modifier OnlyCreator(string memory sharingID){
        require(GeneSharingRepository[sharingID].creator == msg.sender, "only creator of GeneSharing allowed");
        _;
    }
    // 只允许MetaData的owner调用
    modifier OnlyOwner(bytes32 dataHash){
        require(MetaDataRepository[dataHash].owner == msg.sender, "only owner of MetaData allowed");
        _;
    }

    //指定的MetaData必须存在
    modifier MetaDataMustExist(bytes32 dataHash){
        require(MetaDataRepository[dataHash].owner != address(0), "the MetaData is not exists");
        _;
    }

    //指定的MetaData必须在CreativeSpaceAsset中
    modifier MetaDataMustInCreativeSpaceAsset(bytes32 dataHash){
        require(CreativeSpaceAsset[dataHash], "the MetaData must in CreativeSpaceAsset");
        _;
    }

    //指定的MetaData的共享状态必须是可共享的
    modifier MetaDataMustShareable(bytes32 dataHash){
        require(MetaDataRepository[dataHash].sharingStatus == true, "the MetaData must be shareable");
        _;
    }
    // 只允许后台调用
    modifier OnlyAdmin {
        require(msg.sender == _admin, "only admin allowed");
        _;
    }

    //在指定sharing加入MetaData
    function AddMetaDataToSharing(bytes32 dataHash, string memory sharingID) external
    MetaDataMustExist(dataHash)
    MetaDataMustInCreativeSpaceAsset(dataHash)
    OnlyCreator(sharingID) {
        GeneSharing storage geneSharing = GeneSharingRepository[sharingID];
        geneSharing.MetaDataInSharing[dataHash] = true;
    }

    //把指定MetaData上传到CreativeSpaceAsset
    function AddMetaDataToCreativeSpaceAsset(bytes32 dataHash) external OnlyOwner(dataHash) {
        CreativeSpaceAsset[dataHash] = true;
    }

    //把指定MetaData从CreativeSpaceAsset移除
    function RemoveMetaDataFromCreativeSpaceAsset(bytes32 dataHash) external OnlyOwner(dataHash) {
        MetaData storage metaData = MetaDataRepository[dataHash];
        require(metaData.owner == msg.sender, "only owner of MetaData allowed");

        CreativeSpaceAsset[dataHash] = false;
    }

    // 只允许非owner调用
    modifier OnlyNotOwner(bytes32 dataHash){
        require(MetaDataRepository[dataHash].owner != msg.sender, "only not owner allowed");
        _;
    }

    //从指定sharing新增指定MetaData查看许可
    function NewViewAccessFromSharing(string memory sharingID, bytes32 dataHash) external MetaDataMustExist(dataHash) OnlyNotOwner(dataHash) {
        GeneSharing storage geneSharing = GeneSharingRepository[sharingID];
        require(geneSharing.MetaDataInSharing[dataHash], "the MetaData is not in the Genesharing");

        // 给予creator奖励
        _transfer(address(this), geneSharing.creator, 1 * 10 ** 18);

        MetaData storage metaData = MetaDataRepository[dataHash];
        // 给予MetaData的owner奖励
        _transfer(address(this), metaData.owner, 1 * 10 ** 18);

        require(metaData.viewerPermission[msg.sender].viewer != msg.sender, "the viewer has already been granted permission");
        _newViewAccess(msg.sender, dataHash);
    }

    //指定sharing新增其下所有MetaData查看许可
    function NewViewAccessFromSharingAll(string memory sharingID) external {
        GeneSharing storage geneSharing = GeneSharingRepository[sharingID];
        for (uint256 i = 0; i < geneSharing.MetaDataKeys.length; i++) {
            bytes32 dataHash = geneSharing.MetaDataKeys[i];
            MetaData storage metaData=MetaDataRepository[dataHash];
            require(metaData.viewerPermission[msg.sender].viewer != msg.sender, "the viewer has already been granted permission");
            _newViewAccess(msg.sender, dataHash);
        }
    }


    //指定MetaData续约查看许可
    function RenewViewAccess(bytes32 dataHash) external MetaDataMustExist(dataHash) OnlyNotOwner(dataHash) {
        MetaData storage metaData = MetaDataRepository[dataHash];
        require(metaData.viewerPermission[msg.sender].viewer == msg.sender, "the viewer has not been granted permission");
        require(metaData.viewerPermission[msg.sender].isValid, "the viewer is not valid");

        //给予MetaData的owner奖励
        _transfer(address(this), metaData.owner, 5 * 10 ** 17);
        _renewViewAccess(msg.sender, dataHash);
    }

    //指定sharing续约其下所有MetaData的查看许可
    function RenewViewAccessBatchesFromSharing(string memory sharingID) external {

    }

    //验证某个用户是否有权限查看某个MetaData
    function isViewerAllowed(address viewer, bytes32 dataHash) external view MetaDataMustExist(dataHash) returns (bool) {
        MetaData storage metaData = MetaDataRepository[dataHash];
        if (metaData.owner == viewer) {
            return true;
        }
        if (metaData.viewerPermission[viewer].isValid && metaData.viewerPermission[viewer].expiry > block.timestamp) {
            return true;
        }
        return false;
    }

    //修改共享状态
    function UpdateSharingStatus(bytes32 dataHash, bool newStatus) external MetaDataMustExist(dataHash) OnlyOwner(dataHash) {
        MetaData storage metaData = MetaDataRepository[dataHash];
        metaData.sharingStatus = newStatus;
    }

    //登记某个用户的认证状态
    function SetUserAuthStatus(address user, bool status) external OnlyAdmin {

    }



}