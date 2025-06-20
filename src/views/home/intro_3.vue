<template>
    <div class="wrapper">
        <h1>What <span style="color: #333">we want to do?</span></h1>
        <div class="page">
            <div class="des ch1">
                <img src="@/icons/transfer.svg" alt="" />Circulation, processing and research of
                genetic information
            </div>
            <div class="des ch2">
                <img src="@/icons/transparent.svg" alt="" />Transparent and traceable information
                circulation
            </div>
            <div class="des ch3">
                <img src="@/icons/animals.svg" alt="" />More species support
            </div>
            <div class="des ch4">
                <img src="@/icons/academic.svg" alt="" />Public welfare and academic
            </div>
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
    text-shadow: 0 5px 5px #ccc;
}

.page {
    position: absolute;
    left: 50%;
    top: 55%;
    transform: translate(-50%, -50%);
    width: 100%;
    height: 50vh;
    display: flex;

    .des {
        position: absolute;
        display: flex;
        width: 20%;
        min-width: 400px;
        height: 16vh;
        font-size: 2vh;
        color: #333;
        font-weight: 700;
        padding: 20px;
        background-color: #fff;
        border-radius: 20px;
        align-items: center;
        text-align: left;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);

        &:hover {
            transform: translateY(-3%);
            transition: all 0.5s;
        }

        img {
            height: 10vh;
            margin-right: 30px;
        }
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
