<template>
    <div class="wrapper">
        <!-- <Bubbles /> -->
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
        </div>
        <div class="plaza-page">
            <div class="card" v-for="(item, index) in List" :key="index" @click="isVisible(item)">
                <div class="icon" />
                <div class="ifo">
                    <p>Name: <span class="ifo-item"> &nbsp; {{ item.name }}</span></p>
                    <p>Category: <span class="ifo-item"> &nbsp; {{ item.category }}</span></p>
                    <p>Description:<span class="ifo-item"> &nbsp; {{ item.description }}</span></p>
                    <p>Format:<span class="ifo-item"> &nbsp; {{ item.format }}</span></p>
                    <p>Date:<span class="ifo-item"> &nbsp; {{ item.created_at }}</span></p>
                </div>
            </div>
        </div>
    </div>

    <el-drawer v-model="drawer" title="Detail" :direction="'rtl'" :before-close="handleClose" :size="'40%'">
        <div class="dt-page-top">
            <div class="icon" />
            <div class="ifo">
                <p>Name: <span class="ifo-item"> &nbsp; {{ selectedData.name }}</span></p>
                <p>Category: <span class="ifo-item"> &nbsp; {{ selectedData.category }}</span></p>
                <p>Description:<span class="ifo-item"> &nbsp; {{ selectedData.description }}</span></p>
                <p>Format:<span class="ifo-item"> &nbsp; {{ selectedData.format }}</span></p>
                <p>Date:<span class="ifo-item"> &nbsp; {{ selectedData.created_at }}</span></p>
            </div>
        </div>

        <div class="dt-page-bottom">
            <el-table v-loading="loading" :element-loading-svg="svg" class="custom-loading-svg"
                element-loading-svg-view-box="-10, -10, 50, 50" :data="detailData">
                <el-table-column prop="desciption" label="Item" width="180">
                {{ detailData.ID }}
                </el-table-column>
                <el-table-column prop="score" label="Score" width="180" />
                <el-table-column prop="rank" label="Rank" />
            </el-table>
        </div>

        <template #footer>
            <div style="flex: auto">
                <el-button type="primary" @click="confirmClick" class="obtain-btn">Obtain</el-button>
            </div>
        </template>
    </el-drawer>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { ethers } from "ethers";
import { useWalletStore } from "@/stores/account";
import Api from "@/axios/aixos";
import { ElLoading } from "element-plus";

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
    const res = await Api.get('/plaza', {

    });
    List.value = res.data.data.multi_metadata;
}

let selectedData = ref(null);
var loading = ref(true);

/* 获取详细数据 */
let detailData = ref(null);
const getDetailData = async () => {
    try {
        const res = await Api.get(`/metadata/${selectedData.data_hash}`, {
            // params: {
            //     data_hash: selectedData.data_hash,
            // }
        })
        if (res.data.code == 200) {
            loading = false;
        }
        // if (!error) {
        //     loading = false;
        //     detailData.value = res.data.data;
        //     drawer.value = true;
        // }
    } catch (error) {
        // drawer.value = false;
        if (error) {
            loading = false;
            // alert('Get detail data failed. Please try again later.');
        }
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

const drawer = ref(false);
const isVisible = (item) => {
    selectedData = item;
    drawer.value = true;
    getDetailData();
};

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
    border-bottom: #ddd 3px solid;
    background-color: #ffffff;
}

.banner-title {
    position: absolute;
    bottom: 0%;
    display: flex;
    font-size: 70px;
    color: #169608;
    animation: slideIn 0.4s ease-in-out forwards;
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
    padding: 25px 0;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    .card {
        display: flex;
        position: relative;
        flex: 0 0 calc(33.33% - 10px);
        box-sizing: border-box;
        height: 200px;
        background-color: #fff;
        box-shadow: 0 0 5px #ccc;
        border-radius: 30px;
        align-items: center;

        &:hover {
            box-shadow: 0 0 15px #ccc;
            cursor: pointer;
        }

        .ifo {
            display: flex;
            flex-direction: column;
            justify-content: center;
            font-size: 18px;
            margin-left: 50px;
            width: 70%;
            border-left: #ddd 2px solid;
            padding: 0 0 0 25px;
            height: 70%;

            p {
                line-height: 24px;
                color: #333;
            }

            .ifo-item {
                color: #67C23A;
            }
        }
    }

}

.icon {
    position: relative;
    left: 25px;
    width: 140px;
    min-width: 140px;
    height: 140px;
    border: #ddd 1px solid;
    border-radius: 15px;
    background-image: url('@/icons/dna_icon.jpg');
    background-size: cover;
    background-position: center;
}

.dt-page-top {
    display: flex;
    align-items: center;

    .icon {
        width: 180px;
        height: 180px;
    }

    .ifo {
        margin-left: 50px;
    }

    p {
        color: #333;
        font-size: 20px;
        line-height: 28px;
    }

    .ifo-item {
        color: #67C23A;
    }
}

.dt-page-bottom {
    padding: 20px;
}

:deep(.el-popup-parent--hidden) {
    width: 100% !important;
}

:deep(.obtain-btn) {
    font-size: 20px;
    border: #169608 1px solid;
    background-color: #169608;
    box-shadow: none;
}

// :deep(.el-loading-mask){
//     z-index: 3000 !important;
// }</style>