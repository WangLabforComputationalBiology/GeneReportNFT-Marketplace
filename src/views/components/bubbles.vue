<template>
    <div class="ball-container">
        <div v-for="(ball, index) in balls" :key="index" class="ball" :style="{
            width: `${ball.size}px`,
            height: `${ball.size}px`,
            background: ball.color,
            left: `${ball.x}px`,
            top: `${ball.y}px`
        }"></div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

const balls = ref([]);

// 创建小球
function createBall() {
    const size = Math.random() * 200 + 20; // 随机大小
    const originX = Math.random() * (window.innerWidth - size - 20);
    const originY = Math.random() * (window.innerHeight - size - 20);
    return {
        size,
        color: `hsl(${Math.random() * 260}, 70%, 50%)`, // 随机颜色
        originX,
        originY,
        x: originX,
        y: originY,
        angle: Math.random() * 9 * Math.PI, // 随机初始角度
        // maxOffset: 15, // 最大偏移量
        // speed: 0.1 // 浮动速度
    };
}

// 小球数量
for (let i = 0; i < 11; i++) {
    balls.value.push(createBall());
}

// // 动画更新
// let animationFrameId = null;
// function update() {
//     balls.value.forEach(ball => {
//         // ball.angle += ball.speed;
//         ball.x = ball.originX + Math.sin(ball.angle) * ball.maxOffset;
//         ball.y = ball.originY + Math.cos(ball.angle) * ball.maxOffset;
//     });
//     animationFrameId = requestAnimationFrame(update);
// }

// // 启动动画
// onMounted(() => {
//     animationFrameId = requestAnimationFrame(update);
// });

//清理动画
// onUnmounted(() => {
//     if (animationFrameId) {
//         cancelAnimationFrame(animationFrameId);
//     }
// });

// 窗口大小变化时调整小球位置
// window.addEventListener('resize', () => {
//     balls.value.forEach(ball => {
//         const size = ball.size;
//         ball.originX = Math.min(ball.originX, window.innerWidth - size);
//         ball.originY = Math.min(ball.originY, window.innerHeight - size);
//         ball.x = ball.originX + Math.sin(ball.angle) * ball.maxOffset;
//         ball.y = ball.originY + Math.cos(ball.angle) * ball.maxOffset;
//     });
// });
</script>

<style scoped>
.ball-container {
    pointer-events: none;

}

.ball {
    position: absolute;
    border-radius: 10%;
    opacity: 0.4;
    transform: rotate(45deg);
}
</style>