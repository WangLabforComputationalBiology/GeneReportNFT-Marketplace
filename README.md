# 🧬 BioLedger - Genomic Data Sharing Platform

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-green.svg)
![FISCO BCOS](https://img.shields.io/badge/blockchain-FISCO%20BCOS-orange.svg)

## 🌟 Project Overview

**BioLedger** is a decentralized genomic data sharing platform built on a consortium blockchain, dedicated to enabling **secure storage**, **authorized sharing**, and **privacy protection** of personal genomic data. Based on [WeBank FISCO BCOS](https://fisco-bcos-doc.readthedocs.io/zh-cn/latest/), BioLedger integrates MetaMask wallet, OAuth2 protocol, encryption technologies, and an ERC20 token incentive mechanism to provide secure, efficient, and traceable data sharing services for data providers (individuals) and institutional users (e.g., universities, pharmaceutical companies, research institutions).

Through **tiered access control**, **data encryption**, and **watermark tracing**, BioLedger ensures the privacy and controllability of genotypic data while incentivizing user participation with token rewards. Data providers earn rewards for sharing data, while institutional users access genotypic data via email verification, supporting applications in genetic research, drug development, and more.

---

## 🚀 Features

- **🔐 User Authentication**: Secure login and identity management via MetaMask wallet.
- **📥 Data Acquisition and Classification**: Retrieve user-authorized data from WeGene via OAuth2 protocol, classify by type (e.g., skin, fitness), and generate hash credentials.
- **🔒 Access Control**:
    - **Individual Users**: Upload genomic data and earn token rewards.
    - **Institutional Users**: Access genotypic data after email verification.
- **🛡️ Data Privacy Protection**:
    - Genotypic data encrypted using HCE (Hybrid CP-ABE with ECIES Encryption).
    - Embedded public key watermarks for tracing secondary data distribution.
    - Offline decryption tool for data preview, with local download prevention and data expiration mechanisms.
- **💰 Incentive Mechanism**: Data providers receive ERC20 token rewards when their data is accessed, redeemable for data analysis services or university-branded merchandise.
- **🔍 Data Traceability**: Watermarks and smart contracts track unauthorized data sharing, with automatic bans for violators.

---

## 🏗️ Technical Architecture

BioLedger is built with the following technology stack, combining blockchain and encryption for secure and efficient data sharing:

| Module             | Technology                          |
|--------------------|------------------------------------|
| **Blockchain**     | FISCO BCOS consortium chain for storing data hashes, access records, and token transactions |
| **Authentication** | MetaMask wallet for Ethereum-compatible address management |
| **Data Integration** | OAuth2 protocol for integration with platforms like WeGene |
| **Data Storage**   | On-chain: Data hashes, access records; Off-chain: Distributed database (e.g., IPFS) |
| **Encryption**     | CP-ABE, ECIES                     |
| **Smart Contracts**| Access control, token distribution, violation bans |
| **Local Decryption Tool** | Local program for data preview and expiration control |

---

## 🛠️ Installation and Deployment

### Requirements
- **MetaMask**: Browser extension
- **FISCO BCOS**: Consortium chain node
- **Database**: Support for distributed storage (e.g., IPFS or MySQL)

### Deployment Steps
1. Configure FISCO BCOS consortium chain nodes to ensure proper operation.
2. Install MetaMask extension and connect to BioLedger’s FISCO BCOS network.
3. Set up a distributed database (e.g., IPFS) for off-chain data storage.
4. Deploy smart contracts to configure access control and token distribution logic.
5. Launch the local decryption tool to enable data preview functionality.

---

## 📖 Usage Guide

1. **User Login**
    - Install MetaMask extension and connect to BioLedger’s FISCO BCOS network.
    - Log in via the BioLedger platform interface and authorize data access from WeGene.

2. **Data Upload**
    - Individual users upload genomic data; BioLedger automatically classifies phenotypic and genotypic data and generates hash credentials for genotypic data.
    - Data is stored off-chain, with hashes recorded on the FISCO BCOS chain.

3. **Data Access**
    - Institutional users submit email verification (e.g., `user@university.edu`).
    - Upon verification, request access to genotypic data and receive encrypted data packages.
    - Use BioLedger’s local decryption tool with MetaMask private key for data preview.

4. **Reward Collection**
    - Data providers automatically receive ERC20 tokens when their data is accessed.
    - Redeem tokens for university-branded merchandise or data analysis services via the BioLedger platform.

5. **Violation Handling**
    - BioLedger monitors secondary data distribution via watermarks.
    - Smart contracts automatically ban users for detected violations.

---

## 🌍 Vision

**BioLedger** aims to build a secure and trusted ecosystem for genomic data sharing, empowering genetic research and personalized medicine for the future!

---