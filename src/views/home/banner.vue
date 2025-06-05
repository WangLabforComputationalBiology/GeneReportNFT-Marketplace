<template>
    <div class="banner">
        <div class="green-bar"></div>
        <div class="main-text"><span style="color: #169608;">Bio</span>Chainer <br>Sharing</div>
        <div class="sec-text">A platform delivering genetic information.</div>
        <div class="scroll-arrow"><span class="arrow-fadein">↓</span></div>
    </div>

</template>

<script setup>

</script>

<style lang="scss" scoped>
.banner {
    z-index: 999;
    position: relative;
    margin: 0;
    padding: 0;
    width: 99vw;
    height: 95vh;
    background-size: cover;
    background-position: center;
    justify-content: center;
    align-items: center;
    scroll-snap-align: start;
    pointer-events: none;


    /* 绿色条块 */
    .green-bar {
        position: absolute;
        top: 50.3%;
        left: 0;
        width: 0;
        /* 初始宽度为0 */
        height: 206px;

        /* 条块高度 */
        background: #169608;
        /* 绿色 */
        transform: translateY(-50%);
        animation: extend 0.8s ease-out forwards;
        /* 延长动画 */
        border-top-right-radius: 48px;
    }

    /* 延长动画 */
    @keyframes extend {
        0% {
            width: 0;
        }

        100% {
            width: 52%;
            /* 延长到容器宽度的一半 */
        }
    }

    /* 主体文字 */
    .main-text {
        position: absolute;
        top: 24%;
        left: 25%;
        transform: translate(-50%, -50%);

        line-height: 200px;
        font-size: 10rem;
        color: #333;
        opacity: 0;
        /* 初始不可见 */
        animation: fadeIn 0.6s ease-in forwards 0.8s;
        /* 延迟0.8s淡入 */
    }

    .sec-text {
        position: absolute;
        top: 52%;
        left: 52.5%;
        transform: translate(-50%, -50%);
        line-height: 200px;
        font-size: 2rem;
        color: #333;
        opacity: 0;
        /* 初始不可见 */
        animation: fadeIn 0.6s ease-in forwards 1.2s;
        /* 延迟0.8s淡入 */
    }

    /* 淡入动画 */
    @keyframes fadeIn {
        0% {
            opacity: 0;
            transform: translateY(20px);
            /* 轻微向下偏移 */
        }

        100% {
            opacity: 1;
            transform: translateY(0);
        }
    }
}


/* 跳跃箭头 */
.scroll-arrow {
    position: absolute;
    bottom: 10%;
    /* 位于容器中间偏下 */
    left: 50%;
    transform: translateX(-50%);
    /* 水平居中 */
    font-size: 2rem;
    color: #333;
    animation: bounce 2.5s infinite;
    /* 跳跃动画，无限循环 */

    .arrow-fadein {
        animation: fadeIn 2.5s infinite;
        /* 箭头淡入动画 */
    }
}

/* 跳跃动画 */
@keyframes bounce {

    0%,
    20%,
    50%,
    80%,
    100% {
        transform: translate(-50%, 0);
        /* 初始位置 */
    }

    40% {
        transform: translate(-50%, -20px);
        /* 向上跳跃 */
    }

    60% {
        transform: translate(-50%, -10px);
        /* 轻微回落 */
    }
}
</style>