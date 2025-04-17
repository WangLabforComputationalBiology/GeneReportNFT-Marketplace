<template>
    <div class="body">
        <div class="banner">
            <h1 class="banner-title">User</h1>
        </div>
        <span>Wallet address: </span> <span>{{ address }}</span><br>
        <span>Wallet balance: </span> <span>{{ balance }}</span>
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
            address: wallet.address,
            balance: wallet.balance,
        }
    },

    created() {
        // this.getWalletBalance();
    },

    methods: {
        
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