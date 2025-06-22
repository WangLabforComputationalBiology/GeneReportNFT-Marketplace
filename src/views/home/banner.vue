<template>
    <div class="banner">
        <div class="green-bar"></div>
        <div class="main-text">
            <img class="logo" src="@/assets/logo.svg" alt="">
            <span style="color: #169608;">Bio</span>Chainer <br><span style="font-weight: 600;">Shares</span>
        </div>
        <div class="sec-text">A platform delivering genetic information.</div>
        <div class="scroll-arrow"><span class="arrow-fadein">↓</span></div>
    </div>

</template>

<style lang="scss" scoped>
.banner {
    font-size: 16px;
    z-index: 999;
    position: relative;
    height: 95vh;
    min-width: 1200px;
    overflow: hidden;
    background-size: cover;
    background-position: center;
    justify-content: center;
    align-items: center;
    scroll-snap-align: start;
    pointer-events: none;
    background-color: #fff;

    /* logo */
    .logo {
        height: 10.7vh;
    }

    /* 绿色条块 */
    .green-bar {
        position: absolute;
        top: 52%;

        /* 初始宽度为0 */
        height: 20vh;

        /* 条块高度 */
        background: #169608;
        /* 绿色 */
        transform: translateY(-50%);
        animation: extend 0.6s ease-out forwards;
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
        top: 24.9%;
        left: 25%;
        transform: translate(-50%, -50%);
        font-size: 15vh;
        font-weight: 700;
        color: #333;
        opacity: 0;
        /* 初始不可见 */
        animation: fadeIn 0.6s ease-in forwards 0.1s; 
        /* 延迟0.8s淡入 */
    }

    .sec-text {
        // min-width: 400px;
        position: absolute;
        top: 62%;
        left: 52.2%;
        transform: translate(-50%, -50%);
        font-size: 4vh;
        color: #333;
        opacity: 0;
        /* 初始不可见 */
        animation: fadeIn 0.4s ease-in forwards 0.6s;
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