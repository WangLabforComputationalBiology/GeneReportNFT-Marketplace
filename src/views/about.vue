<template>
    <div class="wrapper">
        <div class="top">
            <h1><span style="color: #169608;">Powered</span> by <img src="@/assets/imgs/FISCO_BCOS_logo.svg"
                    alt="Fisco bcos logo" @click="toFB"></h1>
        </div>
        <div class="center">
            <p><span style="color: #169608;">Bio</span>Chainer Team</p>
            <p>WangLab@Shenzhen Technology University</p>
            <p>Email: wangxin@sztu.edu.cn</p>
        </div>

        <div class="bottom">
            <p style="color: #333; text-align: center; margin-top: 20px;">© 2025 BioChainer Team. All rights reserved.</p>
        </div>
    </div>
</template>

<script setup>
function toFB() {
    window.open('https://www.fisco.com.cn/fisco_20.html', '_blank');
}
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
}

.top {
    position: absolute;
    top: 15%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 600px;
    height: 100px;
    color: #333;
    line-height: 80px;
    text-align: center;

    img {
        cursor: pointer;
    }
}

.center {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -40%);
    width: 600px;
    height: 400px;

    P{
        font-size: 24px;
        color: #333;
        margin: 20px 0;
        text-align: center;
    }
}

.bottom {
    position: absolute;
    bottom: 20px;
    left: 50%;
    transform: translateX(-50%);
    width: 100%;
    text-align: center;
    color: #333;
}
</style>