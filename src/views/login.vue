<template>

    <body>
        <div class="loginPage">
            <div class="Wrapper">
                <h1>How to begin?</h1>
                <p>1.Download MetaMask in your Brower and login.</p>
                <img style="margin:25px 0 0 50px;width: 64%;border-radius:10px;" src="../assets/imgs/step1.png" alt="">

            </div>
            <div class="Wrapper">
                <div class="useMeta">
                    <p>2.Try this:</p>
                    <span class="meta" @click="connectWallet">
                        <img class="icon" src="../icons/metalogo.png" />
                        <p style="color:#C0C4CC;">MetaMask</p>
                    </span>
                </div>
            </div>
        </div>
    </body>

</template>

<script>
import { ethers, getAddress } from 'ethers';
import { useWalletStore } from '@/stores/account'
import axios from 'axios';
const wallet = useWalletStore();

export default {
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
        // console.log("++++++"+ wallet.address);
        // this.account = wallet.address; // 获取store中的账户
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
            try {
                // 1. 检查MetaMask是否安装
                if (!window.ethereum) {
                    throw new Error('Please install MetaMask');
                }

                // 2. 创建provider并请求账户
                const provider = new ethers.BrowserProvider(window.ethereum);
                const accounts = await provider.send("eth_requestAccounts", []);
                this.address = accounts[0];

                // 3. 获取nonce
                const nonceResponse = await axios.get(`http://120.24.168.132:8080/user/nonce/${this.address}`);//等待异步完成避免拿不到nonce
                this.nonce = nonceResponse.data.data.nonce;

                // 4. 构造消息
                this.message = this.structureMessage(this.address, this.nonce);
                console.log('Message to sign:', this.message);

                // 5. 请求签名
                const signature = await window.ethereum.request({
                    method: "personal_sign",
                    params: [this.message, this.address],
                });

                // 6. 发送登录请求
                const loginResponse = await axios.post("http://120.24.168.132:8080/user/login", {
                    user_address: this.address,
                    signature: signature,
                });

                console.log('Login success:', loginResponse.data);
                this.$message.success('Login successful!');

            } catch (error) {
                console.error('Error:', error);

                if (error.code === 4001) {
                    this.error = 'You denied the wallet connection';
                } else if (error.response) {
                    this.error = error.response.data.message || 'Login failed';
                } else {
                    this.error = error.message || 'Unknown error occurred';
                }

                this.$message.error(this.error);
            }
        }
    }


}

</script>

<style lang="scss" scoped>
body {
    margin: auto;
    width: 1400px;
}

.loginPage {
    height: 700px;
    display: flex;

    .Wrapper {
        width: 700px;
        background-color: #fff;
        overflow: hidden;

        h1 {
            text-align: center;
            margin-top: 80px;
            font-size: 48px;
            color: #169608;
            margin-bottom: 20px;
        }

        p {
            margin-left: 50px;
            text-align: left;
            ;
            font-size: 24px;
            color: #67C23A;
        }
    }

    .useMeta {
        margin-top: 175px;
        margin-left: 80px;

        p {
            text-align: left;
            margin-left: 0;
            color: #E6A23C;
        }

        .meta {
            margin-top: 30px;
            display: flex;
            border-radius: 10px;
            border: 1px solid #E4E7ED;


            p {
                font-size: 36px;
                line-height: 120px;
                margin-left: 100px;
            }

            &:hover {
                cursor: pointer;
                box-shadow: 0 0 6px #DCDFE6;
                transition: 200ms;
            }
        }
    }

}
</style>