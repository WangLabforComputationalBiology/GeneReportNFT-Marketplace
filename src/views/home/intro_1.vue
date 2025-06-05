<template>
    <div class="wrapper">

        <div class="wrapper-left">
            <h1>Sharing Gene code.</h1><br>
            <h2>A Biological alliance-chain based on <span class="click-link" @click="toFB">FISCO BCOS↗</span> from
                SZTU@Wanglab.
                <br>
                We advocate the creation of a blockchain-based biological information sharing chain to achieve reliable
                circulation of genetic information between universities and specific institutions, promote the
                circulation of genetic information and assist genetic research.
            </h2>
            <div class="Btn-wrapper">
                <div class="toBTN" @click="">Publication</div>
                <div class="toBTN" @click="">About us</div>
                <div class="toBTN" @click="">Git Hub↗</div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
    name: "Intro1",
    data() {
        return {};
    },
    methods: {
        toFB() {
            window.open('https://www.fisco.com.cn/', '_blank');
        }
    }
};
</script>

<style lang="scss" scoped>
a {
    
}

.wrapper {
    height: 95vh;
    background-color: #fff;
    scroll-snap-align: start;
    display: flex;
    position: relative;
}

.wrapper-left {
    // width: 40%;
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
        font-size: 50px;
        color: #169608;
        margin-bottom: 20px;
    }

    h2 {
        font-size: 24px;
        color: #333;
        margin-bottom: 40px;
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
</style>