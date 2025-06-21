<template>
    <div class="outerWrapper">
        <Bubbles />
        <div class="wrapper">
            <div class="useMeta">
                <p>3. Click to connect</p>
                <span class="meta" @click="connectWallet">
                    <img class="icon" src="@/icons/metalogo.png" />
                    <p style="color:#333;">MetaMask</p>
                </span>
            </div>
        </div>
        <div class="wrapper">
            <h1>Login to <span style="color:#333;">begin.</span></h1>
            <p>1. <a href="https://chromewebstore.google.com/search/metamask?utm_source=ext_sidebar"
                    target="_blank">Download MetaMask extension</a> in your Brower and setup your wallet.</p>
            <img class="step-tip-img" src="@/assets/imgs/step1.png" alt="step1">
            <p style="font-size: 16px;color:#333">*We do suggest that use Google Chrome.</p>
            <br>
            <p>2. Click the little earth icon, and switch to BioChainer network.</p>
            <img class="step-tip-img" src="@/assets/imgs/step2.jpg" alt="step2">
        </div>

    </div>


</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Bubbles from '@/views/components/bubbles.vue';
import { ethers, getAddress } from 'ethers';
import { useWalletStore } from '@/stores/account'
import Api from '../../axios/aixos'
import { ElLoading } from 'element-plus'

const fullscreenLoading = ref(false)
const walletStore = useWalletStore();
const router = useRouter();

const address = ref('');
const message = ref('');
const nonce = ref('');
const error = ref('');

function structureMessage(addressVal, nonceVal) {
    const template = `Welcome to GeneReport_platform!

Click to sign in and accept the OpenSeaTerms of Service and Privacy Policy.

This request will not trigger a blockchain transaction or cost any gas fees.

Wallet address:
${addressVal}

Nonce:
${nonceVal}`;

    return template;
}

// MetaMask连接并获取账户
async function connectWallet() {
    error.value = '';
    address.value = '';
    nonce.value = '';
    message.value = '';
    try {
        // 1. 检查MetaMask是否安装
        if (!window.ethereum) {
            throw new Error('Please install MetaMask');
        }

        // 2. 创建provider并请求账户
        const provider = new ethers.BrowserProvider(window.ethereum);
        const accounts = await provider.send("eth_requestAccounts", []);
        address.value = accounts[0];

        if (!error.value) {
            // 3. 获取nonce
            const nonceResponse = await Api.get(`/user/nonce/${address.value}`);
            nonce.value = nonceResponse.data.data.nonce;

            // 4. 构造消息
            message.value = structureMessage(address.value, nonce.value);

            // 5. 请求签名
            const signature = await window.ethereum.request({
                method: "personal_sign",
                params: [message.value, address.value],
            });

            // 6. 发送登录请求
            const loginResponse = await Api.post("/user/login", {
                user_address: address.value,
                signature: signature,
            });

            // console.log('Login success:', loginResponse.data);
            walletStore.setAddress(address.value);
            walletStore.setInstitution(loginResponse.data.data.user.institution);
            walletStore.setToken(loginResponse.data.data.access_token);
            walletStore.setEmail(loginResponse.data.data.user.email);
            // Use globalProperties for $message in script setup
            if (typeof window !== 'undefined' && window.__VUE_APP__ && window.__VUE_APP__.config.globalProperties.$message) {
                window.__VUE_APP__.config.globalProperties.$message.success('Login successful!');
            }

            toAccount();

        }

    } catch (err) {
        console.error('Error::', err.message, err.code);
        // Use globalProperties for $message in script setup
        const $message = typeof window !== 'undefined' && window.__VUE_APP__ && window.__VUE_APP__.config.globalProperties.$message
            ? window.__VUE_APP__.config.globalProperties.$message
            : null;
        if (err.message && err.message.includes('404')) {
            $message && $message.error('Network error.');
        }
        if (err.message && err.message.includes('-32002')) {
            $message && $message.error('MetaMask is not accessible.Please unlock your wallet and try again.');
        }
        if (err.code === 4001) {
            $message && $message.warning('You denied the wallet connection');
        }
        if (err.code === 500 || err.code === 501 || err.code === 502 || err.code === 503 || (err.message && err.message.includes('503'))) {
            alert('Server error, please try again later.');
        }
    }
}

const toAccount = () => {
    const loading = ElLoading.service({
        lock: true,
        text: 'Loading...',
        background: 'rgba(0, 0, 0, 0.7)',
        customClass: 'loading',
    })
    setTimeout(() => {
        loading.close()
        router.push('/account');

    }, 1500)
}
</script>

<style lang="scss" scoped>
.outerWrapper {
    display: flex;
    position: relative;
    height: 95vh;
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

    .wrapper {
        overflow: hidden;
        position: absolute;
        width: 600px;
        top: 40%;
        right: 55%;
        transform: translateY(-50%);


        h1 {
            text-align: center;
            font-size: 48px;
            color: #169608;
            margin-bottom: 20px;
        }

        p {
            text-align: left;
            font-size: 24px;
            color: #333;
        }

        a {
            color: #333;

            &:hover {
                text-decoration: underline;
                color: #169608;
            }
        }

        .step-tip-img {
            margin: 25px 0 0 50px;
            width: 64%;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }

        &:nth-child(2) {
            top: 40%;
            left: 50%;
            transform: translateY(-50%);
        }
    }



}

.useMeta {

    .meta {
        margin-top: 20px;
        display: flex;
        border-radius: 20px;
        border: 1px solid #E4E7ED;
        background-color: #fff;

        .icon {
            margin-left: 50px !important;
        }

        p {
            font-size: 36px;
            line-height: 120px;
            margin-left: 100px;
        }

        &:hover {
            cursor: pointer;
            box-shadow: 0 0 6px #DCDFE6;
        }
    }
}
</style>