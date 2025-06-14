<template>
    <div class="wrapper">
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
        </div>
        <div class="plaza-page">
            <div class="card">
                <div class="icon">
                </div>
            </div>
            <div class="card"></div>
            <div class="card"></div>
            <div class="card"></div>
            <div class="card"></div>
        </div>
    </div>
</template>

<script>
import { ethers } from "ethers";
import { useWalletStore } from "@/stores/account";
const wallet = useWalletStore();
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
            balance: null,
            activeName: 'first',//默认选中第一个分类
            visible: false,//购买弹窗
            //表单数据
            tableData: [{
                id: 133,
                nft_name: 'Name1',
                sales: 0,
                limit: 10,
                address: '上海大学',
                description: "xxxxx"
            }, {
                id: 212,
                nft_name: 'Name1',
                sales: 0,
                limit: 10,
                address: '上海大学'
            }, {
                id: 331,
                nft_name: 'Name1',
                sales: 0,
                limit: 10,
                address: '上海大学'
            }, {
                id: 423,
                nft_name: 'Name1',
                sales: 0,
                limit: 10,
                address: '上海大学'
            }]
        };
    },

    async created() {

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


                    // const balanceWei = await provider.getBalance(accounts[0]);// 获取余额（返回值为 BigNumber，单位为 wei）
                    // this.balance = ethers.formatEther(balanceWei);// 将余额转换为 ETH 单位
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
                // 验证账户余额是否足够
                const requiredBalance = ethers.parseEther(this.amount);
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
                this.$message({
                    message: "交易成功！",
                    type: "success",
                });

            }
            catch (error) {
                console.error("交易失败:", error);
                this.$message({
                    message: "交易失败！",
                    type: "error",
                });
            }

        },

        drawerVisible() {
            this.$nextTick(() => {
                this.visible = true;
            })
        },

        purchase(id, price) {
            this.$router.push({
                path: '/plaza/confirm',
                query: {
                    id
                }
            });
        }
    },


}
</script>

<style lang="scss" scoped></style>

<style lang="scss" scoped>
.wrapper {
    width: 80vw;
    min-width: 1200px;
    margin: auto;
}

.banner {
    height: 150px;
    display: flex;
    position: relative;
    border-bottom: #169608 3px solid;
}

.banner-title {
    position: absolute;
    bottom: 0%;
    display: flex;
    font-size: 70px;
    color: #169608;
    animation: slideIn 0.8s ease-in-out forwards;
}

@keyframes slideIn {
    from {
        transform: translateX(-5%);
        opacity: 0;
    }

    to {
        transform: translateX(0);
        opacity: 1;
    }
}

.plaza-page {
    padding: 15px 0;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    .card {
        display: flex;
        position: relative;
        flex: 0 0 calc(33.33% - 10px);
        box-sizing: border-box;
        width: 31%;
        height: 200px;
        background-color: #fff;
        box-shadow: 0 0 0 1px #eee;
        border-radius: 20px;

        &:hover {
            box-shadow: 0 0 5px #ddd;

        }

        .icon{
            position: relative;
            top: 50%;
            transform: translateY(-50%);
            left: 20px;
            width: 140px;
            height: 70%;
            border: #169608 1px solid;
            border-radius: 15px;
        }
    }

}
</style>