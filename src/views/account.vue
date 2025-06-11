<template>
    <div class="wrapper">
        <Bubbles />
        <div class="info">
            <h1 class="title"><span style="color: #169608;">Your</span> information</h1>
            {{ address ? `Address: ${address}` : 'No address connected' }}
            <br>

            {{ insititutionAccreditation ? ` You are from: ${insititution}.`  : 'Your institution accreditation is not verified.' }}

            <router-link to="/verify" style="color: #fff;">
            
                <el-button>
                    verify
                </el-button>
            </router-link>
        </div>
    </div>
</template>

<script setup>
import { useWalletStore } from '@/stores/account';
import Bubbles from './components/bubbles.vue';
const wallet = useWalletStore();
const address = wallet.address;

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
</style>