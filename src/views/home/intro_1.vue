<template>
    <div class="wrapper">
        <Bubbles :balls="balls" />
        <div class="wrapper-left">
            <h1>Sharing Gene code.</h1><br>
            <h2>A Biological alliance-chain based on <span class="click-link" @click="toFB">FISCO BCOS↗</span> from
                SZTU@Wanglab.
                <br>
                We advocate the creation of a blockchain-based biological information sharing chain to achieve reliable
                circulation of genetic information between universities and specific institutions, promote the
                circulation of genetic information and assist genetic research.
                <br>
                <br>
                <span class="tags">#Non-profit</span>
                <span class="tags">#Volunteer</span>
                <span class="tags">#Academic</span>
            </h2>

            <div class="Btn-wrapper">
                <div class="toBTN" @click="">Publication</div>
                <div class="toBTN" @click="">Git Hub↗</div>
            </div>


        </div>
    </div>
</template>

<script setup>
import Bubbles from './bubbles.vue';
</script>

<style lang="scss" scoped>


.wrapper {
    height: 95vh;
    background-color: #fff;
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
    left: 50%;
    top: 50%;
    transform: translate(-50%, -70%);


    h1 {
        font-size: 60px;
        color: #169608;
        margin-bottom: 20px;
    }

    h2 {
        font-size: 24px;
        color: #333;
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

.Btn-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20px;

    .toBTN {
        width: 200px;
        height: 50px;
        background-color: #169608;
        color: #fff;
        border-radius: 25px;
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
        font-size: 18px;


        &:hover {
            box-shadow: 0 0 20px #e0e0e0;
        }

        &:nth-child(1) {
            background-color: #169608;
        }

        &:nth-child(2) {
            background-color: #E6A23C;
        }

        &:nth-child(3) {
            background-color: #67C23A;
        }


    }
}

.click-link {
    cursor: pointer;
    text-decoration: underline;
}

.ball {
    position: absolute;
    border-radius: 50%;
    opacity: 0.7;
    pointer-events: none;
}
</style>