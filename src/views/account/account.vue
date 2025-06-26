<template>
    <div class="wrapper">
        <Bubbles />
        <div class="info">
            <h1 class="title"><span style="color: #169608;">Account</span> information</h1>
            {{ walletStore.address ? `Address: ${walletStore.address}` : 'No address connected' }}
            <br>
            {{ walletStore.insititution ? ` Institution: ${walletStore.insititution}.` : 'Your institution accreditation is not verified.' }}
            <br>
            {{ walletStore.email ? `Email: ${walletStore.email}` : 'No email verified.'}}
            <router-link to="/verify" style="color: #fff;">
                <el-button class="custom-button">
                    verify
                </el-button>
            </router-link>
            <br>
            <el-button class="custom-button" @click="logout">Log out</el-button>
        </div>
    </div>
</template>

<script setup>
import { useWalletStore } from '@/stores/account';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus'
import Bubbles from '@/views/components/bubbles.vue'

const walletStore = useWalletStore();
const router = useRouter();
const logout = () => {
    walletStore.$reset();//登出重置
    ElMessage.success('Log out successful!');
    setTimeout(() => {
        router.push('/login');
    }, 1500);
    
}

</script>

<style lang="scss" scoped>
.wrapper {
    display: flex;
    position: relative;
    width: 100vw;
    min-width: 800px;
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
    position: absolute;
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
        background-color: #fff;
        color: #169608;
    }

}
</style>