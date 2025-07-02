<template>
    <div class="wrapper">
        <div class="wrapper-left">
            <h1>Sharing <span style="color: #333;"> Gene code.</span></h1><br>
            <h2>
                <div class="des">
                    BioChainer, a Biological alliance-chain based on <span class="click-link" @click="toFB">FISCO
                        BCOS.</span>
                    <br>
                    We advocate the creation of a blockchain-based biological information sharing chain to achieve
                    reliable
                    circulation of genetic information, promote the
                    circulation of genetic information and assist genetic research.
                </div>
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

        <div class="card">
            <img class="MM" src="@/icons/metalogo.svg" alt="Metamask_logo" @click="toMM">
            <div class="des">MetaMask, an interactive tool which is a powerful wallet manager that manages
                account addresses and accesses the BioChainer network.</div>
        </div>
        <div class="card">
            <img class="FB" src="@/icons/FISCO_BCOS_logo.svg" alt="Fisco_bcos_logo" @click="toFB">
            <div class="des des1">Rely on FISCO BCOS to build a currency-free blockchain, and build a shared platform by
                interacting with the metamask wallet.</div>
        </div>

        <div class="wrapper-right">
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
    height: 95vh;
    min-width: 1200px;
    min-height: 800px;
    display: flex;
    position: relative;
}

.wrapper-left {
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-align: center;
    position: absolute;
    width: 40%;
    left: 15%;
    top: 55%;
    transform: translate(-0%, -75%);

    h1 {
        text-align: center;
        font-size: 60px;
        font-weight: 500;
        color: #169608;
        margin-bottom: 20px;
        text-shadow: 0 5px 5px #ccc;
    }

    h2 {
        height: 15vh;
        font-size: 2.5vh;
        color: #333;
        margin-bottom: 30px;
        text-shadow: 0 5px 5px #ccc;
    }

    .tags {
        font-size: 18px;
        color: #333;
        font-weight: 700;
        background-color: #f0f0f0;
        padding: 5px 10px;
        border-radius: 12px;
        margin: 0 10px;
    }

}

.Btn-wrapper {
    display: flex;
    margin-top: 100px;
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

.card {
    width: 500px;
    height: 220px;
    background-color: #ffffff;
    position: absolute;
    left: 60%;
    top: 18%;
    border-radius: 25px;
    border: #ccc;
    box-shadow: #ccc 0 0 10px;

    .FB {
        margin:15px 80px;
        width: 340px;
    }


    &:nth-child(2) {
        top: 46%;
        left: 60%;
        display: flex;
        align-items: center;


        .MM {
            margin: 0 10px 0 35px;
            width: 125px;
        }
    }

    .des1 {
        margin: 0 50px !important;
        text-align: center;
    }

    .des {
        font-size: 20px;
        color: #333;
        margin: 0 30px;

    }
}
</style>