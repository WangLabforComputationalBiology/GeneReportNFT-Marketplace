


-- TODO linux和macos对库的大小写敏感！对表的大小写敏感性取决于数据库的初始化参数和存储引擎

-- drop database GeneReport_platform;
create database if not exists GeneReport_platform;




use GeneReport_platform;



show tables ;


# ================================  数据库表结构  ================================
-- 创建用户列表表
CREATE TABLE `users` (
`id` int NOT NULL PRIMARY KEY COMMENT '自增id',
`address` varchar(64) NOT NULL COMMENT '钱包账户地址',
`name` varchar(32) NOT NULL COMMENT '用户名',
`picture` varchar(1024) NOT NULL DEFAULT '默认头像path' COMMENT '头像地址/文件',
`create_time` datetime NOT NULL COMMENT '账户注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户列表';

-- 创建基因报告基础信息表
CREATE TABLE `GNFT_basic` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`profile_format` varchar(40) NOT NULL COMMENT '基因检测厂商',
`profile_sex` int(8) NOT NULL COMMENT '基因报告检测者对应性别'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='基因报告基础信息';

-- 创建GNFT_athletigen表
CREATE TABLE `GNFT_athletigen` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT 'case_id',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类',
`score` int(8) NOT NULL COMMENT '该类得分',
`rank` varchar(10) NOT NULL COMMENT '等级',
`genetypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_athletigen';

-- 创建GNFT_skin表
CREATE TABLE `GNFT_skin` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '皮肤特性项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类',
`score` int(8) NOT NULL COMMENT '该类得分',
`rank` varchar(10) NOT NULL COMMENT '等级',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_skin';


-- 创建GNFT_psychology_report_id表
CREATE TABLE `GNFT_psychology_report_id` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '心理特质项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类',
`score` int(8) NOT NULL COMMENT '该类得分',
`rank` varchar(10) NOT NULL COMMENT '等级',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_psychology_report_id';


-- 创建GNFT_risk表
CREATE TABLE `GNFT_risk` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '疾病风险项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类或项目名称',
`risk` float NOT NULL COMMENT '用户在该项目上的患病风险倍数',
`percent` float NOT NULL COMMENT '用户在该项目所有用户中的百分比排名',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_risk';


-- 创建GNFT_health_carrier_report_id表
CREATE TABLE `GNFT_health_carrier_report_id` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '遗传性疾病项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类或项目名称',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_health_carrier_report_id';

-- 创建GNFT_health_drug表
CREATE TABLE `GNFT_health_drug` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '药物反应项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类或项目名称',
`tsummary` varchar(255) NOT NULL COMMENT '药物指南或遗传特征的用户报告结论',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_health_drug';

-- 创建GNFT_health_traits表
CREATE TABLE `GNFT_health_traits` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '遗传特征项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类或项目名称',
`tsummary` varchar(255) NOT NULL COMMENT '遗传特征的用户报告结论',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_health_traits';

-- 创建GNFT_health_metabolism表
CREATE TABLE `GNFT_health_metabolism` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`case_id` int NOT NULL COMMENT '营养代谢项目的唯一识别码',
`description` varchar(50) NOT NULL COMMENT '分析报告所属分类或项目名称',
`rank` varchar(10) NOT NULL COMMENT '用户在该项目中的代谢能力等级',
`tsummary` varchar(255) NOT NULL COMMENT '营养代谢类项目的用户报告结论',
`genotypes` varchar(512) NOT NULL COMMENT '基因型详情'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_health_metabolism';

-- 创建GNFT_ancestry表
CREATE TABLE `GNFT_ancestry` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`data_hash` varchar(64) NOT NULL COMMENT '数据哈希值',
`else_asia` float NOT NULL COMMENT '亚洲以外成分比例',
`southeast_asia` float NOT NULL COMMENT '东南亚成分比例',
`chinese_nation` float NOT NULL COMMENT '中华民族成分比例',
`african` float NOT NULL COMMENT '非洲成分比例',
`northeast_asia_america` float NOT NULL COMMENT '东北亚及美洲成分比例',
`europe` float NOT NULL COMMENT '欧洲成分比例',
`area` varchar(512) NOT NULL COMMENT '地区成分详情',
`update_time` int(10) NOT NULL COMMENT '更新时间戳'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_ancestry';

-- 创建GNFT_haplogroups表
CREATE TABLE `GNFT_haplogroups` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`y_haplo` varchar(32) NOT NULL COMMENT 'Y染色体单倍群',
`mt_haplo` varchar(32) NOT NULL COMMENT '线粒体DNA单倍群'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_haplogroups';

-- 创建GNFT_demographics表
CREATE TABLE `GNFT_demographics` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`report_id` int NOT NULL  COMMENT '分析报告类别id',
`surname` varchar(50) NOT NULL COMMENT '用户姓氏',
`population` varchar(50) NOT NULL COMMENT '所属民族',
`native_city` varchar(50) NOT NULL COMMENT '籍贯城市',
`native_province` varchar(50) NOT NULL COMMENT '籍贯省份'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='GNFT_demographics';



-- 创建GNFT_web_auth_profile_id表
CREATE TABLE `GNFT_web_auth_profile_id` (
`profile_id` varchar(32) NOT NULL COMMENT '基因报告id',
`token` varchar(255) NOT NULL COMMENT '临时访问密码',
`expires_in` int(10) NOT NULL DEFAULT 1
)
