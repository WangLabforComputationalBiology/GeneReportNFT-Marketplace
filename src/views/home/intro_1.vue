<template>
    <div class="wrapper">
        <div class="wrapper-left">
            <div class="inner-wrapper">
                <h1>Sharing <span style="color: #333;"> Gene code.</span></h1><br>
                <h2>
                    <div class="des-left">
                        <span style="color: #169608;">Bio</span>Chainer, a Biological alliance-chain based on <span
                            class="click-link" @click="toFB">FISCO
                            BCOS.</span>
                        <br>
                        We advocate the creation of a blockchain-based biological information sharing chain to achieve
                        reliable
                        circulation of genetic information, promote the
                        circulation of genetic information and assist genetic research.
                    </div>
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

        <div class="wrapper-right">
            <div class="card">
                <img class="FB" src="@/icons/FISCO_BCOS_logo.svg" alt="Fisco_bcos_logo" @click="toFB">
                <div class="des des1">Rely on FISCO BCOS to build a currency-free blockchain, and build a shared
                    platform by
                    interacting with the metamask wallet.</div>
            </div>
            <div class="card">
                <img class="MM" src="@/icons/metalogo.svg" alt="Metamask_logo" @click="toMM">
                <div class="des">MetaMask, an interactive tool which is a powerful wallet manager that manages
                    account addresses and accesses the BioChainer network.</div>
            </div>

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
    width: 100%;
    min-width: 1200px;
    min-height: 800px;
    position: relative;
}

.wrapper-left {
    display: flex;
    flex-direction: column;
    width: 60%;
    height: 60%;
    min-height: 460px;
    justify-content: center;
    text-align: center;
    position: absolute;
    border-radius: 25px;
    left: -2%;
    top: 18%;
    background-color: #fff;
    @include shadow;

    .inner-wrapper {
        position: relative;
        width: 80%;
        left: 15%;
    }

    h1 {
        text-align: center;
        font-size: 56px;
        font-weight: 500;
        color: $color-primary;
    }

    h2 {
        font-size: 20px;
        color: #333;
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
    margin-top: 30px;
    justify-content: center;
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


.wrapper-right {
    height: 60%;
    width: 45%;
    padding-left: 40px;
    border-radius: 20px;
    min-height: 460px;
    position: relative;
    left: 60%;
    top: 18%;
    background-color: $color-primary;
    @include shadow;


    .card {
        margin: 20px 0;
        width: 500px;
        height: 25vh;
        min-height: 210px;
        position: absolute;
        border-radius: 15px;
        @include shadow;
        background-color: #fff;
        // @include blur;

        .FB {
            margin: 15px 80px;
            width: 340px;
        }


        &:nth-child(2) {
            bottom: 0%;
            display: flex;
            align-items: center;

            .MM {
                margin: 0 10px 0 50px;
                width: 115px;
            }
        }


    }
}

.des1 {
    margin: 0 50px !important;
    text-align: center;
}

.des {
    font-size: 20px;
    color: #333;
    margin: 0 50px 0 10px;

}

.des-left {
    margin: 0 30px;
}
</style>