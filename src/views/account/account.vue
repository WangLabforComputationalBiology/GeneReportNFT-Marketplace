<template>
    <div class="wrapper">
        <div class="bar">
            <p class="dec">BioChainer</p>
            <p class="dec">BioChainer</p>
        </div>
        <div class="info">
            <div ref="avatarContainer" />
            <div class="info-address">
                <span>
                    {{ useWalletStore().address.slice(0, 8) + '...' + useWalletStore().address.slice(-6) }}
                </span>
            </div>
            {{ useWalletStore().insititution ? ` Institution: ${useWalletStore().insititution}.` : `Your institution
            accreditation is not verified.` }}
            <br>
            {{ useWalletStore().email ? `Email: ${useWalletStore().email}` : `No email verified.` }}
            <router-link to="/verify" style="color: #fff;">
                <el-button class="custom-button">
                    Verify
                </el-button>
            </router-link>
            <br><br><br><br><br>

            <el-button class="custom-button" @click="logout">Log out</el-button>
            <br>
            <el-button class="custom-button logs-btn" @click="viewActivity">Activity</el-button>
            <br>
            <!-- <el-button class="custom-button logs-btn" @click="addERC20TokenToMetaMask">Add Token</el-button> -->
            <!-- <el-button class="custom-button logs-btn" @click="addNetworkToMetaMask">Add Network</el-button> -->
        </div>
    </div>

    <el-drawer :direction="'btt'" title="Activity" v-model="ActivityIsVisible" :size="'60%'" @closed="drawerClosed">
        <el-table :data="activityData" style="width: 100%" :table-layout="'fixed'">
            <el-table-column prop="time" label="Time" width="160" />
            <el-table-column prop="expiry" label="Expiry Time" width="160" />
            <el-table-column prop="event" label="Event" width="100" />
            <el-table-column prop="id" label="Id" />
            <el-table-column prop="geneSharing" label="Report Address" />
            <el-table-column prop="metadata" label="Metadata Address" />
            <el-table-column prop="from" label="From" />
        </el-table>
    </el-drawer>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useWalletStore } from '@/stores/account';
import { useRouter } from 'vue-router';
import { ElMessage, ElLoading } from 'element-plus';
import Bubbles from '@/views/components/bubbles.vue';
import jazzicon from 'jazzicon';
import Api from "@/axios/aixos";
import { contractAddress } from '@/contract/contractConfig.js';

/**
 * jazzicon 
 * 根据hash生成头像，有趣的库
 * 确保在组件挂载后操作 DOM
 */
const avatarContainer = ref(null);
const diameter = 100; // 头像直径

const router = useRouter();

/**自动添加ERC20代币（Token） */
async function addERC20TokenToMetaMask() {
    if (!window.ethereum) {
        ElMessage.warning('MetaMask missing.');
        return
    }
    const tokenAddress = contractAddress; // 同合约地址
    const tokenSymbol = 'GNC'; // 币符号，必填
    const tokenDecimals = 18; // 小数点位数必填（通常 18，但 USDT 是 6）

    try {
        // 调用 MetaMask 的 wallet_watchAsset 方法
        const wasAdded = await window.ethereum.request({
            method: 'wallet_watchAsset',
            params: {
                type: 'ERC20', // 代币类型
                options: {
                    address: tokenAddress,
                    symbol: tokenSymbol,
                    decimals: tokenDecimals,
                },
            },
        });

        if (wasAdded) {
            console.log('代币已成功添加到 MetaMask!');
            ElMessage.success('Token added successfully!');
        } else {
            console.log('用户拒绝了代币添加请求。');
            ElMessage.warning('Token added failed!');
        }
    } catch (error) {
        if (error.code === 4001 || error.message === "User rejected the request.") {
            ElMessage.warning('Token added failed! User denied.');
        }
    }
}

/**自动添加网络 */
async function addNetworkToMetaMask() {
    if (!window.ethereum) {
        ElMessage.warning('MetaMask missing.');
        return
    }

    // 要添加的网络配置（以 BSC 主网为例）
    const networkConfig = {
        chainId: "0x4EE8", // 20200 的 16 进制
        chainName: "BioChainer",
        nativeCurrency: {
            name: "FBC",
            symbol: "FBC",
            decimals: 18,
        },
        rpcUrls: ["http://10.108.10.51:8545"],  // RPC 节点，必须要https协议
        blockExplorerUrls: [],  // 区块链浏览器,目前没有  
    };

    try {
        // 调用 MetaMask 的 wallet_addEthereumChain 方法
        await window.ethereum.request({
            method: "wallet_addEthereumChain",
            params: [networkConfig],
        });
        console.log("网络已成功添加到 MetaMask！");
    } catch (error) {
        console.error("添加网络失败:", error);
        if (error.code === 4001) {
            console.log("用户拒绝了网络添加请求。");
        }
    }
}

/**登出 */
const logout = () => {
    // walletStore.$reset();//登出重置
    let loading = ElLoading.service({
        lock: true,
        text: 'Loading',
        background: 'rgba(0, 0, 0, 0.7)'
    });
    setTimeout(() => {
        useWalletStore().reset();
        router.push('/login');
        loading.close();
        ElMessage.success('Log out successful!');
    }, 1500);

}

/**
 * 查看日志
 * @param {bool} logsIsVisible  日志抽屉开关
 * @param {array} activityData  日志数组 [{}]
 */
const ActivityIsVisible = ref(false);
const activityData = ref([]);
function viewActivity() {
    ActivityIsVisible.value = true;
    getActivity();
}
async function getActivity() {
    if (!useWalletStore().insititution || !useWalletStore().email) {
        ElMessage.warning('Please verify your email and institution before.');
        return
    }
    if (!useWalletStore().token) {
        ElMessage.warning('Token Error! Please login again.');
        return
    }
    try {
        const eventMap = {
            'NewViewAccess': 'Create',
            'Metadata': 'Metadata',
        }
        const res = await Api.get('/user/activity');
        if (res.data.code === 200) {
            activityData.value = res.data.data.multi_activities;
            activityData.value.forEach((item) => {
                item.time = new Date(item.time).toLocaleString();
                item.expiry = new Date(item.expiry).toLocaleString();
                item.id = item.id.slice(0, 10) + '...' + item.id.slice(-10);
                item.geneSharing = item.geneSharing.slice(0, 10) + '...' + item.geneSharing.slice(-10);
                item.metadata = item.metadata.slice(0, 10) + '...' + item.metadata.slice(-10);
                item.from = item.from.slice(0, 10) + '...' + item.from.slice(-10);
                item.event = eventMap[item.event] || item.event;
            })
        }
    } catch (error) {
        console.log("Error: " + error)
        ElMessage.error('Get activity failed!');
    }
}

/**回收 */
function drawerClosed() {
    activityData.value = [];
}

onMounted(() => {
    if (useWalletStore().address && avatarContainer.value) {
        // 生成头像
        const identicon = jazzicon(
            diameter,
            parseInt(useWalletStore().address.slice(2, 10), 16) // 使用地址部分作为种子
        );
        avatarContainer.value.appendChild(identicon);
    }
})
</script>

<style lang="scss" scoped>
.wrapper {
    position: relative;
    width: 100%;
    min-width: 1200px;
    height: 95vh;
    margin: auto;
    overflow: hidden;
    animation: fadeIn 0.2s ease-in-out 0s forwards;

    @keyframes fadeIn {
        0% {
            opacity: 0;
        }

        100% {
            opacity: 1;
        }
    }

    .title {
        font-size: 70px;
        color: #333;
        margin-bottom: 20px;
    }

}

.bar {
    position: absolute;
    top: 45%;
    transform: translate(-2%, -50%);
    width: 56%;
    height: 600px;
    border-radius: 20px;
    box-shadow: #00000010 0 0 10px;
    background-color: $color-primary;
    pointer-events: none;
    overflow: hidden;

    .dec {
        position: absolute;
        font-size: 280px;
        font-weight: 700;
        color: #fff;
        transform: rotate(-48deg) translate(15%, -60%);
        top: 30%;


        &:nth-child(1) {
            transform: rotate(-48deg) translate(-20%, 0%);
            top: 50%;
        }
    }
}

.info {
    border-radius: 20px;
    position: absolute;
    width: 45%;
    height: 600px;
    padding: 60px;
    top: 45%;
    left: 80%;
    transform: translate(-50%, -50%);
    font-size: 18px;
    color: #333;
    line-height: 35px;
    @include shadow;

    .info-address {
        color: $color-primary;
    }
}

:deep(.custom-button) {
    margin: 5px 0;
    font-size: 18px;
    background-color: #169608;
    color: #fff;
    width: 200px;
    height: 40px;
    box-shadow: none !important;
    border: #169608 solid 1px;
    border-radius: 10px;

    &:hover {
        box-shadow: 0 0 0 5px #ccc;
    }

    &:nth-child(1) {
        width: 60px;
        height: 34px;
        background-color: #fff;
        color: #169608;
    }

}

.logs-btn {
    background-color: #fff;
    color: #169608;
}
</style>