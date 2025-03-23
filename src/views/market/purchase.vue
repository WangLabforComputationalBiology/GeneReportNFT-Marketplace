<template>

    <body>
        <div class="purchase-page">
            <div class="info-wrapper">
                <h1 style="color: #169608;margin-bottom:50px">Confirm your order</h1>
                <h1>ID: #{{ purchaseId.id }} </h1>
                <p>you spend:</p>
                <div class="price">{{ purchaseId.price }} ETH</div>
            </div>
            <div class="info-wrapper">
                <div class="img">
                    <img src="../../assets/imgs/bioschains.svg" alt="">
                </div>
            </div>
            <el-button>Confirm</el-button>
        </div>
    </body>
</template>

<script>
export default {
    name: 'Purchase',
    computed: {
        purchaseId() {
            return this.$route.query;
        }
    }
}

</script>

<style lang="scss" scoped>
body {
    margin: auto;
    width: 1400px;
    overflow: visible;
}

.purchase-page {
    position: relative;
    display: flex;
    margin: auto;
    margin-top: 100px;
    margin-bottom: 100px;
    padding: 30px 100px;
    height: 620px;
    width: 1000px;
    border-radius: 15px;
    padding: 50px 100px;
    box-shadow: 0px 0px 16px #DCDFE6;

    .info-wrapper {
        width: 50%;
        display: inflex;
        height: 400px;
        padding: 50px 0px 50px 20px;

        h1 {
            margin: 15px 0;
            color: #67C23A;
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
                margin: 40px 0 0 80px;
            }
        }
    }

    :deep(.el-button) {
        left: 50%;
        transform: translateX(-50%);
        width: 760px;
        height: 50px;
        position: absolute;
        bottom: 100px;
        border-radius: 10px;
        line-height: 9px;
    }
}
</style>