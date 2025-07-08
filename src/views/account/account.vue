<template>
    <div class="wrapper">
        <Bubbles />
        <div class="info">
            <!-- <h1 class="title"><span style="color: #169608;">Account</span> information</h1> -->
            <div ref="avatarContainer" />
            {{ useWalletStore().address ? `Address: ${useWalletStore().address}` : `No address connected` }}
            <br>
            {{ useWalletStore().insititution ? ` Institution: ${useWalletStore().insititution}.` : `Your institution
            accreditation is not verified.` }}
            <br>
            {{ useWalletStore().email ? `Email: ${useWalletStore().email}` : `No email verified.` }}
            <router-link to="/verify" style="color: #fff;">
                <el-button class="custom-button">
                    Verify
                </el-button>
            </router-link>
            <br>
            <el-button class="custom-button" @click="logout">Log out</el-button>
            <br>
            <el-button class="custom-button logs-btn" @click="viewActivity">Activity</el-button>
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
import { ElMessage } from 'element-plus';
import Bubbles from '@/views/components/bubbles.vue';
import jazzicon from 'jazzicon';
import Api from "@/axios/aixos";

/**
 * jazzicon 
 * 根据hash生成头像，有趣的库
 * 确保在组件挂载后操作 DOM
 */
const avatarContainer = ref(null);
const diameter = 100; // 头像直径

const router = useRouter();

/**登出 */
const logout = () => {
    // walletStore.$reset();//登出重置
    useWalletStore().reset();
    ElMessage.success('Log out successful!');
    setTimeout(() => {
        router.push('/login');
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
    display: flex;
    position: relative;
    width: 100vw;
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

.info {
    background-color: #ffffffaa;
    border-radius: 30px;
    box-shadow: 0 0 5px #ccc;
    position: absolute;
    width: 800px;
    padding: 60px;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
    font-size: 24px;
    color: #333;
    line-height: 50px;
}

:deep(.custom-button) {
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