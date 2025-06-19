<template>
    <div class="wrapper">
        <div class="wrapper-left">
            <h1>How we do that?</h1><br>
            <h2>
                Through blockchain and smart contracts, we enable the recording and traceability of
                genetic information transfers. The process of logging transfer records on the chain is decentralized,
                while the storage of genetic data remains centralized—a decision made for regulatory compliance and
                security considerations.
                <br>
                We advocate for the participation of universities and legitimate institutions,
                with the goal of advancing academic research.
                <br>
                <br>
                <span class="tags">#Coinless blockchain</span>
                <span class="tags">#Traceable</span>
            </h2>



        </div>
        <div class="wrapper-right">
            <h2>Framework</h2>
            <img src="@/assets/imgs/framework.svg" alt="">
        </div>
    </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
.wrapper {
    height: 95vh;
    min-width: 1200px;
    scroll-snap-align: start;
    display: flex;
    position: relative;
}

.wrapper-left {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    position: absolute;
    width: 50%;
    height: 95vh;
    border-bottom: #fff 1px solid;
    left: 0%;
    // top: 50%;
    background-color: #169608;
    // transform: translate(0%, -60%);


    h1 {
        font-size: 6vh;
        font-weight: 500;
        color: #fff;
        margin-bottom: 20px;
    }

    h2 {
        font-size: 2.5vh;
        color: #fff;
        margin-bottom: 30px;
    }

    .tags {
        font-size: 18px;
        color: #333;
        font-weight: 700;
        background-color: #f0f0f0;
        padding: 5px 10px;
        border-radius: 12px;
        margin: 10px;
    }

}

.wrapper-right {
    position: absolute;
    left: 50%;
    // top: 50%;
    width: 50%;
    height: 95vh;
    background-color: #fff;

    // transform: translate(0%, -70%);
    img {
        position: absolute;
        top: 50%;
        left: 25%;
        transform: translate(0%, -50%);
        width: 25vw;
    }

    h2 {
        position: absolute;
        top: 25vh;
        left: 25vw;
        transform: translate(-50%, -50%);
        color: #169608;
        font-weight: 700;
        font-size: 30px;
    }
}
</style>