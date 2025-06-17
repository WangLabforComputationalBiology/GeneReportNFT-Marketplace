<template>
    <div class="wrapper">
        <Bubbles />
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
        </div>
        <div class="plaza-page">
            <div class="card" v-for="(item, index) in List" :key="index">
                <div class="icon" />
                <div class="ifo">
                    <p>Id: <span class="ifo-item"> &nbsp; {{ item.ID }}</span></p>
                    <p>Description:<span class="ifo-item"> &nbsp; {{ item.description }}</span></p>
                    <p>Score:<span class="ifo-item"> &nbsp; {{ item.score }}</span></p>
                    <p>Rank:<span class="ifo-item"> &nbsp; {{ item.rank }}</span></p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { ethers } from "ethers";
import { useWalletStore } from "@/stores/account";
import Api from "@/axios/aixos";
import Bubbles from "../components/bubbles.vue";

const wallet = useWalletStore();
const account = ref("");
const provider = ref(null);
const signer = ref(null);
const recipient = ref("");
const amount = ref("");
const txHash = ref("");

/* 获取plaza卡片数据列表 */
let List = ref([]);
async function getList() {
    const res = await Api.get('/user/getData/Skin');
    List.value = res.data;
    console.log(List.value)
}



async function connectWallet() {
    if (window.ethereum) {
        try {
            // 创建 Web3Provider
            const _provider = new ethers.BrowserProvider(window.ethereum);
            provider.value = _provider;
            // 请求账户权限
            const accounts = await _provider.send("eth_requestAccounts", []);
            account.value = accounts[0]; // 获取第一个账户
            console.log("Account:", account.value);

            signer.value = await _provider.getSigner();// 获取签名者（用于发送交易）
            console.log("Signer:", signer.value);
            console.log("Balance:", balance.value);

        } catch (error) {
            console.error("连接失败:", error);
        }
    } else {
        alert("请安装 MetaMask");
    }
}

// 发送交易
async function sendTransaction() {
    if (!account.value || !signer.value) {
        alert("请先连接 MetaMask");
        return;
    }
    if (!ethers.isAddress(recipient.value)) {
        alert("请输入有效的接收地址！");
        return;
    }
    if (isNaN(amount.value) || parseFloat(amount.value) <= 0) {
        alert("请输入有效的 ETH 数量！");
        return;
    }
    if (!signer.value) {
        console.error("Signer 未初始化");
        return;
    }
    try {
        // 验证账户余额是否足够
        const requiredBalance = ethers.parseEther(amount.value);
        if (balance.value < requiredBalance) {
            alert("账户余额不足！");
            return;
        }
        console.log("发送交易...");

        // 获取账户
        const _provider = new ethers.BrowserProvider(window.ethereum);
        const accounts = await _provider.send("eth_requestAccounts", []);
        account.value = accounts[0];

        // 发送交易
        const _signer = await _provider.getSigner();
        const tx = await _signer.sendTransaction({
            to: recipient.value,
            value: ethers.parseEther(amount.value), // 转换 ETH 单位
        });

        console.log("交易发送中...", tx);
        // alert("交易已提交，等待确认...");
        if (typeof window !== "undefined" && window.$message) {
            window.$message({
                message: "交易已提交，等待确认...",
                type: "success",
            });
        }

        // 等待交易完成
        await tx.wait();
        txHash.value = tx.hash;
        if (typeof window !== "undefined" && window.$message) {
            window.$message({
                message: "交易成功！",
                type: "success",
            });
        }

    }
    catch (error) {
        console.error("交易失败:", error);
        if (typeof window !== "undefined" && window.$message) {
            window.$message({
                message: "交易失败！",
                type: "error",
            });
        }
    }
}


function purchase(id, price) {
    router.push({
        path: '/plaza/confirm',
        query: {
            id
        }
    });
}
onMounted(() => {
    getList();
});
</script>

<style lang="scss" scoped>
.wrapper {
    height: 95vh;
    width: 80vw;
    min-width: 1200px;
    margin: auto;
}

.banner {
    height: 150px;
    display: flex;
    position: relative;
    border-bottom: #E6A23C 3px solid;
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
    width: 100%;
    padding: 15px 0;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    .card {
        display: flex;
        position: relative;
        flex: 0 0 calc(33.33% - 10px);
        box-sizing: border-box;
        // width: 31%;
        height: 200px;
        background-color: #fff;
        box-shadow: 0 0 0 1px #eee;
        border-radius: 20px;
        align-items: center;

        &:hover {
            box-shadow: 0 0 5px #ddd;

        }

        .icon {
            position: relative;
            left: 20px;
            width: 140px;
            min-width: 140px;
            height: 70%;
            border: #ddd 1px solid;
            border-radius: 15px;
            background-image: url('@/icons/dna_icon.jpg');
            background-size: cover;
            background-position: center;
        }

        .ifo {
            margin-left: 50px;
            width: 60%;
            border-left: #ddd 2px solid;
            padding: 10px 0 0 25px;
            height: 70%;

            p{
                color: #333;
            }
            .ifo-item {
                color: #67C23A;
            }
        }
    }

}
</style>