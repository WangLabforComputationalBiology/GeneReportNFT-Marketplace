# Vue3 Webpage for Biochainer

Author: Tang Jiaming@SZTU BDI   202200201002@stumail.sztu.edu.cn    
Date: 2025.2 - 2025.8  

## 项目介绍-Introduction

基于 Vue3 的 BioChainer 前端页面，提供账户验证、合约交互功能以及数据共享功能模块。合约通过 MetaMask 钱包交互。

## **目录**-Directory

- [项目技术栈-Technology Stack](#项目技术栈-Technology-Stack)
- [项目启动-Startup](#项目启动-Startup)
- [部署说明-Deployment Notes](#部署说明-Deployment-Notes)

## 项目技术栈-Technology Stack

- Vue3
- Vite
- Vue Router
- Pinia
- Element Plus
- Axios
- Ethers.js

## 项目启动-Startup

```bash
# 安装依赖
npm install

# 运行项目
npm run dev

# 构建项目
npm run build

```

## 部署说明-Deployment Notes

区块链前端的区别在于需要在前端进行对合约的操作，如合约部署、合约调用等。

- 合约部署：
  - 合约 ABI：参考 `contractABI.json` 文件，通常为.abi 或.json 格式，描述合约的接口，**需要根据实际更改**。
  - 合约地址：参考 `contractConfig.js` 文件，通常为合约部署后的地址，目前为硬编码，不同的合约地址不同，**需要根据实际更改**。
- 前端配置：
  - 请确保已安装 MetaMask 浏览器插件并配置好钱包账户，**根据 MetaMask 官方描述，目前仅支持 Chrome、Edge、Firefox、Opera 以及 Brave 浏览器**，详情请参考[MetaMask 官方](https://metamask.io/)。
  - 请确保 MetaMask 正确连接到 BioChainer 的网络；BioChainer 网络为测试网，需自行部署网络与合约。
- 关于WeBASE
  - FISCO BCOS提供官方WeBASE区块链中间件平台，详情：[WeBASE官方文档](https://fisco-bcos-documentation.readthedocs.io/zh-cn/latest/docs/webase/webase.html)
