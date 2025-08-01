# 🧬 Genetic Data Sharing Platform

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-green.svg)
![FISCO BCOS](https://img.shields.io/badge/blockchain-FISCO%20BCOS-orange.svg)

## 🌟 Project Overview

The Genetic Data Sharing Platform is a decentralized solution based on a consortium blockchain, designed to enable **secure storage**, **authorized sharing**, and **privacy protection** of personal genetic data. Built on the [FISCO BCOS](https://fisco-bcos-doc.readthedocs.io/zh-cn/latest/) consortium blockchain, the platform integrates MetaMask wallets, OAuth2 protocols, encryption technologies, and an ERC20 token incentive mechanism to provide secure, efficient, and traceable data sharing services for data providers and institutional users (e.g., universities, pharmaceutical companies, and research institutes).

Through **tiered access control**, **data encryption**, and **watermark tracing**, the platform ensures the privacy and controllability of genomic data while incentivizing user participation with token rewards. Data providers earn rewards for sharing data, and institutional users access genomic data via email verification, supporting applications in genetic research, drug development, and more.

---

## 🚀 Features

- **🔐 User Authentication**: Secure login and identity management via MetaMask wallet.
- **📥 Data Acquisition and Classification**: Obtains user-authorized data from WeGene via OAuth2 protocol, categorizing data (e.g., skin, fitness) and generating hash credentials.
- **🔒 Access Control**:
    - **Regular Users**: Upload genetic data and earn token rewards.
    - **Institutional Users**: Access genomic data upon email verification.
- **🛡️ Data Privacy Protection**:
    - Genomic data encrypted with HCE (Hybrid CP-ABE with ECIES Encryption).
    - Embedded public key watermarks for tracing secondary data distribution.
    - Offline decryption tool enables data preview, prohibits local downloads, and enforces data expiration.
- **💰 Incentive Mechanism**: Data providers receive ERC20 tokens when their data is accessed, redeemable for data analysis services or university-branded merchandise.
- **🔍 Data Traceability**: Watermarks and smart contracts track unauthorized data sharing, enabling automatic banning of violators.

---

## 🏗️ Technical Architecture

The platform is built with the following technology stack, combining blockchain and encryption technologies for secure and efficient data sharing:

| Module              | Technology Implementation                     |
|---------------------|---------------------------------------------|
| **Blockchain**      | FISCO BCOS consortium chain for storing data hashes, access records, and token transactions |
| **Authentication**  | MetaMask wallet for Ethereum-compatible address management |
| **Data Integration**| OAuth2 protocol for integration with platforms like WeGene |
| **Data Storage**    | On-chain: Data hashes, access records; Off-chain: Distributed database (e.g., IPFS) |
| **Encryption**      | CP-ABE, ECIES                              |
| **Smart Contracts** | Access control, token distribution, and violation banning |
| **Local Decryption Tool** | Local application for data preview and expiration control |

---

## 🛠️ Installation and Deployment

### Prerequisites
- **MetaMask**: Browser extension
- **FISCO BCOS**: Consortium blockchain node
- **Database**: Support for distributed storage (e.g., IPFS or MySQL)

---

## 📖 Usage Instructions

1. **User Login**
    - Install the MetaMask browser extension and connect to the FISCO BCOS network.
    - Log in through the platform interface and authorize data access from WeGene.

2. **Data Upload**
    - Regular users upload genetic data, which is automatically classified into phenotypic and genomic data, with hashes generated for genomic data.
    - Data is stored off-chain, with hashes recorded on the FISCO BCOS blockchain.

3. **Data Access**
    - Institutional users submit email verification (e.g., `user@university.edu`).
    - Upon verification, request access to genomic data and receive encrypted data packages.
    - Use the local decryption tool with a MetaMask private key to preview data.

4. **Reward Acquisition**
    - Data providers automatically receive ERC20 tokens when their data is accessed.
    - Redeem tokens for university-branded merchandise or data analysis services on the platform's redemption page.

5. **Violation Handling**
    - The platform monitors secondary data distribution via watermarks.
    - Smart contracts automatically ban users upon detection of violations.

---

*Building a secure and trusted future for genetic data sharing!*