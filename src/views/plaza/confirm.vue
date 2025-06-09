<template>
    <div class="wrapper">
        <Bubbles />
        <div class="confirm-page">
            <div class="info-wrapper">
                <h1 style="color: #169608;margin-bottom:10px">Confirm imformation</h1>
                <h1>ID: #{{ purchaseId.id }} </h1>
                <span class="address-info">Your address: <span style="color: #67C23A; margin: 0 20px;">{{ address
                        }}</span>
                </span><br>
                <span class="address-info"> Destination address: <span style="color: #67C23A;margin: 0 20px;">{{
                        }}0x121724567f7d4373b589dacaac567815c1c6d904</span>
                </span><br>
            </div>
            <div class="info-wrapper">
                <div class="img">
                    <img src="../../assets/imgs/bioschains.svg" alt="">
                </div>
            </div>
            <div class="btn-wrapper">
                <el-button @click="goBack">Back</el-button>
                <el-button>Confirm</el-button>
            </div>


        </div>
    </div>
</template>

<script setup>
import { useWalletStore } from "@/stores/account";
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import Bubbles from "../components/bubbles.vue";

const wallet = useWalletStore();
const address = ref(wallet.address);
const route = useRoute();
const purchaseId = computed(() => route.query);

const router = useRouter();
const goBack = () => {
    router.go(-1); // 返回上一页
};

</script>

<style lang="scss" scoped>
.wrapper {
    height: 95vh;
}

.confirm-page {
    z-index: 999;
    background-color: #ffffffe6;
    position: absolute;
    display: flex;
    top: 50%;
    left: 50%;
    height: 580px;
    width: 900px;
    border-radius: 25px;
    padding: 50px 50px;
    box-shadow: 0px 0px 16px #DCDFE6;
    animation: fadeUp 0.5s ease-out forwards;

    .info-wrapper {
        width: 50%;
        display: inflex;
        padding: 20px 0px 40px 20px;

        h1 {
            margin: 15px 0;
            color: #333;
            font-size: 40px;
        }

        p {
            font-size: larger;
            font-weight: bold;
            color: #A8ABB2;
            margin: 20px 0 20px 0;
            font-size: 24px;
        }

        .price {
            color: #E6A23C;
            font-size: 34px;
            margin-bottom: 80px;
        }

        .img {
            justify-content: center;

            img {
                width: 250px;
                margin: 50px 0 0 90px;
            }
        }

        .address-info {
            line-height: 36px;
            font-size: 18px;
            color: #A8ABB2;

            &:last-child {
                line-height: 60px;
            }
        }
    }


}

.btn-wrapper {
    display: flex;
    position: absolute;
    bottom: 80px;
    left: 70px;
    justify-content: space-between;

    :deep(.el-button) {
        width: 210px;
        height: 50px;
        border-radius: 15px;
        line-height: 9px;
        font-size: 18px;
        color: #169608;

        &:hover {
            background-color: #fff;
            color: #169608;
            border: 1px solid #169608;
        }

        &:nth-child(2) {
            color: #fff;
            background-color: #169608;
        }
    }
}

@keyframes fadeUp {
    from {
        opacity: 0;
        transform: translate(-50%, -52%);
    }

    to {
        opacity: 1;
        transform: translate(-50%, -50%);
    }
}
</style>