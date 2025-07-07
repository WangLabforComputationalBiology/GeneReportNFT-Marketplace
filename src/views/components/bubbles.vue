<template>
    <div class="ball-container">
        <div v-for="(ball, index) in balls" :key="index" class="ball" :style="{
            width: `${ball.size}px`,
            height: `${ball.size}px`,
            background: ball.color,
            left: `${ball.x}px`,
            top: `${ball.y}px`
        }">
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const balls = ref([]);

/* 初始化小球 */
function createBall() {
    const size = Math.random() * 180 + 30; // 随机大小
    const originX = Math.random() * (window.innerWidth - size - 20);
    const originY = Math.random() * (window.innerHeight - size - 200);//随机位置
    return {
        size,
        color: `hsl(${Math.random() * 260}, 70%, 50%)`, // 随机颜色
        originX,
        originY,
        x: originX,
        y: originY,
    };
}

/* 小球数量 */
for (let i = 0; i < 11; i++) {
    balls.value.push(createBall());
}
</script>

<style scoped>
.ball-container {
    /*  防止鼠标事件穿透  */
    pointer-events: none;
}

.ball {
    position: absolute;
    border-radius: 10%;
    opacity: 0.4;
    /* 旋转45度 */
    transform: rotate(45deg);
}
</style>