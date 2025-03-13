<template>
    <div class="body">
        <div class="banner">
            <h1 class="banner-title">Market</h1>
        </div>
        <p>trading_test</p>
        <div class="marketPage">
            <!-- 交易表单 -->
            <div v-if="account">
                <input type="text" style="width: 100%;" v-model="recipient" placeholder="接收地址 (0x...)" />
                <input type="text" v-model="amount" placeholder="ETH 数量" />
                <button @click="sendTransaction">发送交易</button>
            </div>
            <p>you got: {{ balance }}</p>
            <p v-if="txHash">交易成功！哈希:{{ txHash }}</p>
        </div>
    </div>
</template>

<script>
import { ethers } from "ethers";
export default {
    name: "market",
    data() {
        return {
            account: "",
            provider: null,
            signer: null,
            recipient: "",
            amount: "",
            txHash: "",
            balance: "",
        };
    },

    async created() {
        await this.connectWallet(); // 确保 created 生命周期钩子调用异步方法
    },
    methods: {
        async connectWallet() {
            if (window.ethereum) {
                try {
                    // 创建 Web3Provider
                    const provider = new ethers.BrowserProvider(window.ethereum);
                    // 请求账户权限
                    const accounts = await provider.send("eth_requestAccounts", []);
                    this.account = accounts[0]; // 获取第一个账户
                    console.log("Account:", this.account);

                    this.signer = await provider.getSigner();// 获取签名者（用于发送交易）
                    console.log("Signer:", this.signer);

                    // 获取余额（返回值为 BigNumber，单位为 wei）
                    const balanceWei = await provider.getBalance(accounts[0]);
                    // 将余额转换为 ETH 单位
                    this.balance = ethers.formatEther(balanceWei);
                    console.log("Balance:", this.balance);

                } catch (error) {
                    console.error("连接失败:", error);
                }
            } else {
                alert("请安装 MetaMask");
            }
        },

        // 发送交易
        async sendTransaction() {
            if (!this.account || !this.signer) {
                alert("请先连接 MetaMask");
                return;
            }
            if (!ethers.isAddress(this.recipient)) {
                alert("请输入有效的接收地址！");
                return;
            }
            if (isNaN(this.amount) || parseFloat(this.amount) <= 0) {
                alert("请输入有效的 ETH 数量！");
                return;
            }
            if (!this.signer) {
                console.error("Signer 未初始化");
                return;
            }
            try {
                // const balance = await this.signer.provider.getBalance(this.account);
                // console.log("账户余额:", ethers.formatEther(balance));
                // 验证账户余额是否足够
                const requiredBalance = ethers.parseEther(this.amount);
                // const balance = this.balance;
                if (this.balance < requiredBalance) {
                    alert("账户余额不足！");
                    return;
                }
                console.log("发送交易...");

                // 获取账户
                const provider = new ethers.BrowserProvider(window.ethereum);
                const accounts = await provider.send("eth_requestAccounts", []);
                this.account = accounts[0];

                // 发送交易
                const signer = await provider.getSigner();
                const tx = await signer.sendTransaction({
                    to: this.recipient,
                    value: ethers.parseEther(this.amount), // 转换 ETH 单位
                });

                console.log("交易发送中...", tx);
                // alert("交易已提交，等待确认...");
                this.$message({
                    message: "交易已提交，等待确认...",
                    type: "success",
                });

                // 等待交易完成
                await tx.wait();
                this.txHash = tx.hash;
                // alert(`交易成功！交易哈希: ${tx.hash}`);
                this.$message({
                    message: "交易成功！",
                    type: "success",
                });

            }
            catch (error) {
                console.error("交易失败:", error);
                // alert("交易失败，请检查错误日志！");
            }

            //更新余额
            const balanceWei = await provider.getBalance(accounts[0]);
            this.balance = ethers.formatEther(balanceWei);
        },
    },


}
</script>

<style lang="scss" scoped>
.body {
    margin: auto;
    width: 1400px;
    background-color: #fff;
}

.banner-title {
    font-size: 70px;
    color: #67C23A;
}

.marketPage {
    height: 500px;
}
</style>