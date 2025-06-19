<template>
    <div class="wrapper">
        <h1>What we want to do?</h1>
        <div class="page">
            <div class="des ch1">有利于基因信息的流通、处理与研究</div>
            <div class="des ch2">信息流通透明、可溯源</div>
            <div class="des ch3">支持其他物种</div>
            <div class="des ch4">公益性、学术性</div>
        </div>
    </div>

</template>

<style scoped>
.wrapper {
    height: 95vh;
    min-width: 1200px;
    scroll-snap-align: start;
    display: flex;
    position: relative;
}

h1 {
    position: absolute;
    left: 50%;
    top: 20%;
    transform: translate(-50%, -50%);
    font-size: 6vh;
    color: #169608;
    margin-bottom: 20px;
    font-weight: 500;
}

.page {
    position: absolute;
    left: 50%;
    top: 55%;
    transform: translate(-50%, -50%);
    width: 100%;
    height: 50vh;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.1);
    display: flex;
    background-color: #169608;

  

    .des {
        position: absolute;

        width: 20%;
        height: 160px;
        font-size: 1.5vh;
        color: #333;
        text-align: center;
        padding: 20px;
        line-height: 1.6;
        background-color: #fff;
        border-radius: 20px;
    }

    .ch1 {
        top: 15%;
        right: 52%;
    }

    .ch2 {
        top: 15%;
        left: 52%;
    }

    .ch3 {
        bottom: 15%;
        right: 52%;
    }

    .ch4 {
        bottom: 15%;
        left: 52%;
    }




}
</style>