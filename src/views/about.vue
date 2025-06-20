<template>
    <div class="wrapper">
        <div class="top">
            <h1><span style="color: #169608;">Powered</span> by <img src="@/assets/imgs/FISCO_BCOS_logo.svg"
                    alt="Fisco bcos logo" @click="toFB"></h1>
        </div>
        <div class="center">
            <p>WangLab@Shenzhen Technology University</p>
            <p>Email: wangxin@sztu.edu.cn</p>
            <p class="pub">Publication</p>
            <p class="pub">Github</p>
        </div>

        <div class="bottom">
            <p>© 2025 BioChainer Team. All rights reserved.</p>
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
    top: 25%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 600px;
    height: 100px;
    color: #333;
    line-height: 80px;
    text-align: center;

    h1 {
        cursor: default;
    }

    img {
        cursor: pointer;
    }
}

.center {
    position: absolute;
    top: 60%;
    left: 50%;
    transform: translate(-50%, -40%);
    width: 600px;
    height: 400px;

    P {
        font-size: 24px;
        color: #333;
        margin: 20px 0;
        text-align: center;
    }
    .pub {
        font-size: 20px;
        color: #333;
        text-decoration: underline;
        cursor: pointer;

        &:hover {
            color: #169608;
            
        }
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
    pointer-events: none;

    p {
        color: #333;
        text-align: center;
        margin-top: 20px;
    }
}
</style>