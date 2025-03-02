<template>
    <div class="body">
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
    </div>

</template>

<script>
import { ethers } from 'ethers';

export default {
    methods: {
        // MetaMask连接并获取账户
        async connectWallet() {
            if (typeof window.ethereum !== 'undefined') {
                try {
                    // 创建 Web3Provider
                    const provider = new ethers.BrowserProvider(window.ethereum);
                    //获取账户
                    const accounts = await provider.send('eth_requestAccounts', []);
                    this.account = accounts[0];
                    console.log('Connected account:', this.account);
                } catch (error) {
                    console.error('User denied account access or error occurred:', error);
                    this.$message({
                        message: 'Please login in to MetaMask.',
                        type: 'error',
                        duration: 3000,
                    });
                }
            } else {
                console.log('MetaMask is not installed');
            }
            if (this.account) {
                this.$message({
                    message: 'Wallet have Connected: ' + this.account,
                    type: 'success',
                    duration: 2000,
                });
            }
        },
    }


}
</script>

<style lang="scss" scoped>
.body {
    margin: auto;
    width: 1400px;
}

.loginPage {
    width: 100%;
    height: 655px;
    display: flex;

    .Wrapper {
        width: 50%;
        height: 100%;
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