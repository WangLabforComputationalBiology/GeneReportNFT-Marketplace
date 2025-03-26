<template>
    <div class="body">
        <div class="banner">
            <h1 class="banner-title">User</h1>
        </div>
        <span>Wallet address: </span> <span>{{ wallet.address }}</span><br>
        <span>Wallet balance: </span> <span>{{ wallet.balance }}</span>
        <div class="dropPage">
        </div>
    </div>

</template>

<script>
import { ethers } from 'ethers';
import { useWalletStore } from '@/stores/account';
const wallet = useWalletStore();
export default {
    name: "drop",
    data() {
        return {
            dropList: [],
            wallet: wallet,
        }
    },

    created() {
        this.getWalletBalance();
    },

    methods: {
        async getWalletBalance() {
            const provider = new ethers.BrowserProvider(window.ethereum);
            const accounts = await provider.send("eth_requestAccounts", []);
            const account = accounts[0]; 
            const balanceWei = await provider.getBalance(account);// 获取余额（返回值为 BigNumber，单位为 wei）
            wallet.setBalance(ethers.formatEther(balanceWei))
            // this.balance = ethers.formatEther(balanceWei);// 将余额转换为 ETH 单位
            console.log("Balance:", wallet.balance);
        }
    },
}
</script>

<style lang="scss" scoped>
.body {
    margin: auto;
    width: 1400px;
    background-color: #fff;
}

.banner {
    display: flex;
    // border-bottom: 1px solid #E4E7ED;

    .banner-title {
        font-size: 70px;
        color: #67C23A;
    }

}

.dropPage {
    height: 700px;
}
</style>