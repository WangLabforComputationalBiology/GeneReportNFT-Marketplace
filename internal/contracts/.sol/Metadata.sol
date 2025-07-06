// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Metadata {
    // 代理
    address private _proxy;

    // 查看许可
    struct ViewAccess {
        //访问者
        address viewer;

        //有效期
        uint256 expiry;

        //是否为有效的许可
        bool isValid;

        //次数
        uint256 count;
    }

    // 溯源元数据
    struct TraceData {
        //操作者是通过哪个sharing合约
        address SharingAddress;

        //目标用户
        address operator;

        //追溯事件状态:
        //0：被operator创建,
        //1：operator通过SharingAddress新增查看许可,
        //2：operator续约查看许可,
        //3：在工坊被operator添加到SharingAddress,
        //4：operator修改共享状态
        int16 operationType;

        //操作时间
        uint timestamp;

        //备注用途
        string remark;
    }

    modifier onlyProxy() {
        require(msg.sender == _proxy, "Only proxy allowed");
        _;
    }

    // 所有者
    mapping(bytes32 dataHash => address) private _owners;

    // 可共享状态
    mapping(bytes32 dataHash => bool) private _sharingStatus;

    // 访问许可
    mapping(bytes32 dataHash => mapping(address => ViewAccess)) private _viewerPermissions;

    // 溯源信息
    mapping(bytes32 dataHash => TraceData[]) private _traceDatas;

    // 默认7天有效期
    uint public DURATION;



    constructor(address proxy_){
        DURATION = 7 days;
        _proxy = proxy_;
    }


    function owner(bytes32 dataHash) public view returns (address) {
        return _owners[dataHash];
    }

    //续约查看许可
    function _renewalViewAccess(ViewAccess storage vs) internal returns(uint256){
        vs.expiry = block.timestamp + DURATION;
        vs.count++;
        return vs.expiry;
    }

    //添加溯源信息（不带备注）
    function _addTraceData(bytes32 dataHash, address operator_, address sharingAddress, int16 operatorType_) internal {
        TraceData[] storage traceDataArray = _traceDatas[dataHash];
        traceDataArray.push(TraceData({
            SharingAddress: sharingAddress,
            remark: "",
            operator: operator_,
            operationType: operatorType_,
            timestamp: block.timestamp
        }));
    }

    //添加溯源信息（不带备注）
    function _addTraceDataWithRemark(bytes32 dataHash, address operator_, address sharingAddress, int16 operatorType_, string memory remark) internal {
        TraceData[] storage traceDataArray = _traceDatas[dataHash];
        traceDataArray.push(TraceData({
            SharingAddress: sharingAddress,
            remark: remark,
            operator: operator_,
            operationType: operatorType_,
            timestamp: block.timestamp
        }));
    }

    //判断MetaData是否存在
    function isMetadataExist(bytes32 dataHash) public view returns (bool) {
        return _owners[dataHash] != address(0);
    }

    //新建Metadata
    function newMetadata(bytes32 dataHash, address owner_) external onlyProxy {
        require(_owners[dataHash] == address(0), "the data hash already exists");
        _owners[dataHash] = owner_;
        _addTraceData(dataHash, owner_, address(0), 0);
    }

    //指定Metadata续约查看许可
    function renewalViewAccess(bytes32 dataHash, address viewer, string memory remark) external onlyProxy returns(uint256){
        ViewAccess storage viewAccess = _viewerPermissions[dataHash][viewer];
        require(viewAccess.viewer == viewer, "the viewer has not been granted permission");
        require(viewAccess.isValid, "the viewer is not valid");

        uint256 expiry=_renewalViewAccess(viewAccess);
        _addTraceDataWithRemark(dataHash, viewer, address(0), 2, remark);
        return expiry;
    }

    //修改共享状态
    function updateSharingStatus(bytes32 dataHash, bool newStatus) external onlyProxy {
        require(msg.sender == owner(dataHash), "only the owner can update the sharing status");
        _sharingStatus[dataHash] = newStatus;

        _addTraceData(dataHash, _owners[dataHash], address(0), 4);
    }

    function addToGeneSharing(bytes32 dataHash, address operator, address sharingAddress) external onlyProxy {
        _addTraceData(dataHash, operator, sharingAddress, 3);
    }

    //验证查看许可
    function verifyViewAccess(bytes32 dataHash, address viewer) external view onlyProxy returns (int)  {
        if (isMetadataExist(dataHash) == false) {
            //the Metadata does not exist
            return 1;
        }

        ViewAccess storage viewAccess = _viewerPermissions[dataHash][viewer];

        if (viewAccess.viewer != viewer) {
            //"the viewer has not been granted permission"
            return 2;
        }

        if (viewAccess.expiry <= block.timestamp) {
            //"the viewer's permission has expired"
            return 3;
        }
        if (viewAccess.isValid == false) {
            return 4;
        }
        return 0;
    }

    //新增查看许可
    function addViewAccess(bytes32 dataHash, address viewer, address sharingAddress, string calldata remark) external onlyProxy returns(uint256){
        require(owner(dataHash) != viewer, "ur the owner of the Metadata, no need to add permission");
        require(isMetadataExist(dataHash), "the Metadata does not exist");
        ViewAccess storage viewAccess = _viewerPermissions[dataHash][viewer];
        require(viewAccess.viewer == address(0), "the viewer has already been granted permission");
        viewAccess.viewer = viewer;
        viewAccess.expiry = block.timestamp + DURATION;
        viewAccess.isValid = true;
        viewAccess.count = 1;

        _addTraceDataWithRemark(dataHash, viewer, sharingAddress, 1, remark);
        return viewAccess.expiry;
    }
    //返回某个Metadata的对于某个viewer的已续约次数
    function getRenewalCount(bytes32 dataHash, address viewer) public view returns (uint256) {
        return _viewerPermissions[dataHash][viewer].count;
    }

    //用户是否拥有access（过期或未过期）
    function isUserHaveAccess(address user, bytes32 dataHash) public view returns (bool) {
        return _viewerPermissions[dataHash][user].count > 0;
    }
}
