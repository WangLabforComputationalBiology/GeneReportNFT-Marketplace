<template>
    <div class="wrapper">
        <Bubbles />
        <div class="info">
            <h1 class="title"><span style="color: #169608;">Account</span> information</h1>
            <div ref="avatarContainer" />
            {{ useWalletStore().address ? `Address: ${useWalletStore().address}` : 'No address connected' }}
            <br>
            {{ useWalletStore().insititution ? ` Institution: ${useWalletStore().insititution}.` : 'Your institution accreditation is not verified.' }}
            <br>
            {{ useWalletStore().email ? `Email: ${useWalletStore().email}` : 'No email verified.'}}
            <router-link to="/verify" style="color: #fff;">
                <el-button class="custom-button">
                    Verify
                </el-button>
            </router-link>
            <br>
            <el-button class="custom-button" @click="logout">Log out</el-button>
        </div>
    </div>
</template>

<script setup>
import { ref,onMounted } from 'vue';
import { useWalletStore } from '@/stores/account';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus'
import Bubbles from '@/views/components/bubbles.vue'
import jazzicon from 'jazzicon'

/**jazzicon 
 * 根据hash生成头像，有趣的库
 * 确保在组件挂载后操作 DOM
 */
const avatarContainer = ref(null); 
const diameter = 100; // 头像直径

const router = useRouter();
const logout = () => {
    // walletStore.$reset();//登出重置
    useWalletStore().reset();
    ElMessage.success('Log out successful!');
    setTimeout(() => {
        router.push('/login');
    }, 1500);
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
        margin-bottom: 40px;
    }

}

.info {
    background-color: #ffffffaa;
    border-radius: 30px;
    box-shadow: 0 0 5px #ccc;
    position: absolute;
    width: 800px;
    padding: 40px;
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
    border-radius: 15px;

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
</style>