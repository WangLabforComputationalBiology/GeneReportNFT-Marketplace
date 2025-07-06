<template>
    <div class="outerWrapper">
        <Bubbles />
        <div class="wrapper">
            <div class="useMeta">
                <h2>3. Click to connect</h2>
                <span class="meta" @click="connectWallet">
                    <img class="icon" src="@/icons/metalogo.png" />
                    <p style="color:#333;">MetaMask</p>
                </span>
            </div>
            <p style="color:#333;font-size: 16px;margin-top: 20px;">*Today's version only supports one account at a time, so please do not switch account in MetaMask after logging in.</p>
        </div>
        <div class="wrapper">
            <h1>Login to <span style="color:#333;">begin.</span></h1>
            <h2>1. <span class="download" @click="downloadMetamask()">Download MetaMask extension</span> in your Brower
                and setup your wallet.</h2>
            <img class="step-tip-img" src="@/assets/imgs/step1.png" alt="step1">
            <p style="font-size: 16px;color:#333">*We do suggest that use Google Chrome. You are using <span
                    style="color: #333;font-weight: 700;">{{ browser }}</span>.
            </p>
            <br>
            <h2>2. Click the little earth icon, and switch to BioChainer network.</h2>
            <img class="step-tip-img" src="@/assets/imgs/step2.jpg" alt="step2">
            <div class="setup-network" @click="isVisible = true">->How to setup BioChainer network?</div>
        </div>
    </div>

    <el-drawer :direction="'ltr'" v-model="isVisible" title="How to setup BioChainer network?"
        header-class="drawer-header">
        <p class="setup-tips">1.Already installed MetaMask and setup your wallet, click the little
            earth icon, and click "Add Network".</p>
        <img class="step-tip-img" src="@/assets/imgs/setup1.jpg" alt="">
        <p class="setup-tips">2."Add a custom network"</p>
        <img class="step-tip-img" src="@/assets/imgs/setup2.jpg" alt="">
        <p class="setup-tips">3.Add the information as shown in the figure above and "Save"</p>
        <p class="setup-tips">4.Back, click the little earth icon, and switch to BioChainer network.</p>
        <p class="setup-tips">5.If seen "MetaMask isn’t connected to this site", turn to "3.Click to connect", connect
            Metamask to BioChainer site.</p>
    </el-drawer>

</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import Bubbles from '@/views/components/bubbles.vue';
import { ethers } from 'ethers';
import { useWalletStore } from '@/stores/account'
import Api from '../../axios/aixos'
import { login } from '../../axios/aixos'
import { ElLoading, ElMessage } from 'element-plus'

const router = useRouter();

const address = ref('');
const message = ref('');
const nonce = ref('');

/*  构造message 格式较严格 勿随意改动 */
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

/* MetaMask连接并获取账户 */
async function connectWallet() {
    address.value = '';
    nonce.value = '';
    message.value = '';
    const loading = ElLoading.service({
        lock: true,
        text: 'Loading...',
        background: 'rgba(0, 0, 0, 0.7)',
        customClass: 'loading',
    })
    try {
        // 1. 检查MetaMask是否安装
        if (!window.ethereum) {
            ElMessage.error('Please install MetaMask');
        }

        // 2. 创建provider并请求账户
        const provider = new ethers.BrowserProvider(window.ethereum);
        const accounts = await provider.send("eth_requestAccounts", []);
        address.value = accounts[0];


        // 3. 获取nonce
        const nonceResponse = await Api.get(`/user/nonce/${address.value}`);
        nonce.value = nonceResponse.data.data.nonce;

        // 4. 构造message消息
        message.value = structureMessage(address.value, nonce.value);

        // 5. 请求签名
        const signature = await window.ethereum.request({
            method: "personal_sign",
            params: [message.value, address.value],
        });

        // 6. 发送登录请求
        const loginResponse = await login.post(`/user/login`, {
            user_address: address.value,
            signature: signature,
        });

        // 7. 保存登录信息
        useWalletStore().setToken(loginResponse.data.data.access_token);
        useWalletStore().setWalletInfo(address.value, loginResponse.data.data.user.institution, loginResponse.data.data.user.email);

        ElMessage.success('Login successful!');
        router.replace(`/account?t=${Date.now()}`);    // 添加时间戳，防止浏览器缓存

    } catch (error) {
        loading.close();
        console.log(error);
        if (error.code == 404) {
            ElMessage.error('Network error.');
        }
        if (error.code == -32002) {
            ElMessage.error('MetaMask is not accessible.Please unlock your wallet and try again.');
        }
        if (error.code === 4001) {
            ElMessage.warning('You denied the wallet connection');
        }
        if (error.code === 500 || error.code === 501 || error.code === 502 || error.code === 503) {
            alert('Server error, please try again later.');
        }
    } finally {
        loading.close()
    }
}

/* network部署提示 */
let isVisible = ref(false)

/* 判断浏览器 */
const browser = computed(() => {
    const ua = navigator.userAgent.toLowerCase();
    if (ua.includes('qqbrowser')) {
        return 'QQBrowser';
    } else if (ua.includes('chrome') && !ua.includes('edg')) {
        return 'Chrome';
    } else if (ua.includes('firefox')) {
        return 'Firefox';
    } else if (ua.includes('safari') && !ua.includes('chrome')) {
        return 'Safari';
    } else if (ua.includes('edg')) {
        return 'Edge';
    } else if (ua.includes('opr')) {
        return 'Opera';
    } else if (ua.includes('msie') || ua.includes('trident')) {
        return 'Internet Explorer';
    }
    return 'Unknown browser';
})
function downloadMetamask() {
    switch (browser.value) {
        case 'Chrome':
            window.open('https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn');
            break;
        case 'Firefox':
            window.open('https://addons.mozilla.org/en-US/firefox/addon/ether-metamask/');
            break;
        case 'Edge':
            window.open('https://microsoftedge.microsoft.com/addons/detail/metamask/ejbalbakoplchlghecdalmeeeajnimhm');
            break;
        case 'Opera':
            window.open('https://addons.opera.com/en/extensions/details/metamask/');
            break;
        default:
            window.open('https://metamask.io/download.html');
    }
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

        .setup-network {
            cursor: pointer;
            width: 300px;

            &:hover {
                text-decoration: underline;
            }
        }

    }

}

.download {
    text-decoration: underline;
    cursor: pointer;

    &:hover {
        color: #169608;
    }
}

.setup-tips {
    font-size: 20px;
    color: #333;
    margin: 10px 0;
}

.useMeta {

    .meta {
        display: flex;
        border-radius: 20px;
        border: 1px solid #E4E7ED;
        background-color: #fff;

        .icon {
            width: 140px;
            margin-left: 50px !important;
        }

        p {
            font-size: 36px;
            line-height: 140px;
            margin-left: 80px;
        }

        &:hover {
            cursor: pointer;
            box-shadow: 0 0 6px #DCDFE6;
        }
    }
}
</style>