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

<script>
import Bubbles from '@/views/components/bubbles.vue';
import { ethers, getAddress } from 'ethers';
import { useWalletStore } from '@/stores/account'
import Api from '../../axios/aixos'
const walletStore = useWalletStore();

export default {
    name: 'Login',
    components: {
        Bubbles
    },
    data() {
        return {
            address: '',
            balance: '',
            message: '',
            nonce: '',
            error: ''
        }
    },
    created() {
    },

    methods: {
        //验签信息构造函数
        structureMessage(address, nonce) {
            const template = `Welcome to GeneReport_platform!

Click to sign in and accept the OpenSeaTerms of Service and Privacy Policy.

This request will not trigger a blockchain transaction or cost any gas fees.

Wallet address:
${address}

Nonce:
${nonce}`;

            return template;
        },

        // MetaMask连接并获取账户
        async connectWallet() {
            this.error = '';
            this.address = '';
            this.nonce = '';
            this.message = '';
            try {
                // 1. 检查MetaMask是否安装
                if (!window.ethereum) {
                    throw new Error('Please install MetaMask');
                }

                // 2. 创建provider并请求账户
                const provider = new ethers.BrowserProvider(window.ethereum);
                const accounts = await provider.send("eth_requestAccounts", []);
                this.address = accounts[0];

                if (!this.error) {
                    // 3. 获取nonce
                    const nonceResponse = await Api.get(`/user/nonce/${this.address}`);
                    this.nonce = nonceResponse.data.data.nonce;

                    // 4. 构造消息
                    this.message = this.structureMessage(this.address, this.nonce);

                    // 5. 请求签名
                    const signature = await window.ethereum.request({
                        method: "personal_sign",
                        params: [this.message, this.address],
                    });

                    // 6. 发送登录请求
                    const loginResponse = await Api.post("/user/login", {
                        user_address: this.address,
                        signature: signature,
                    });

                    console.log('Login success:', loginResponse.data);
                    walletStore.setAddress(this.address);
                    walletStore.setInstitution(loginResponse.data.data.user.institution);
                    walletStore.setToken(loginResponse.data.data.access_token);
                    walletStore.setEmail(loginResponse.data.data.user.email);
                    this.$message.success('Login successful!');


                    setTimeout(() => {
                        this.$router.push('/account');
                    }, 1500);
                }


            } catch (error) {
                console.error('Error::', error.message, error.code);
                if (error.message.includes('404')) {
                    this.$message.error('Network error.');

                }
                if (error.message.includes('-32002')) {
                    this.$message.error('MetaMask is not accessible.Please unlock your wallet and try again.');
                }
                if (error.code === 4001) {
                    this.$message.warning('You denied the wallet connection');
                }
                if (error.code === 500 || error.code === 501 || error.code === 502 || error.code === 503 || error.message.includes('503')) {
                    alert('Server error, please try again later.');
                }
            }
        }
    }


}

</script>

<style lang="scss" scoped>
.outerWrapper {
    display: flex;
    position: relative;
    height: 95vh;
    overflow: hidden;

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