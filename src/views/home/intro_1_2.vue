<template>
    <div class="wrapper">
        <div class="inner-wrapper">
            <div class="card">
                <div class="logo"><img src="@/icons/metalogo.svg" alt=""></div>
                <div class="des">Use MetaMask as an interactive tool, which is a powerful wallet manager that manages
                    account addresses and accesses the BioChainer network.</div>
            </div>
            <div class="card">
                <div class="logo FB"><img class="FB" src="@/icons/FISCO_BCOS_logo.svg" alt="Fisco bcos logo"
                        @click="toFB"></div>
                <div class="des">Rely on FISCO BCOS to build a currency-free blockchain, and build a shared platform by
                    interacting with the metamask wallet.</div>
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
    min-width: 1200px;
    display: flex;
    position: relative;
    justify-content: center;
    position: relative;

}

.inner-wrapper {
    position: absolute;
    top: 15%;
    width: 50%;
    height: 50%;
    align-items: center;
    border-radius: 40px;
    box-shadow: 0 0 15px #E4E7ED;
    background-color: #ffffffee;
    display: flex;
    gap: 20px;

    .card {
        width: 48%;
        height: 70%;

        .logo {
            margin-top: 50px;
            margin: auto;
            width: 150px;
            height: 150px;

            img {
                width: 150px;
            }
        }

        .FB {
            width: 400px;

            img {
                margin: 20px 100px 0 0;
                width: 400px;
            }
        }

        .des {
            margin: auto;
            padding: 20px;
            width: 350px;
            height: 250px;
            font-size: 24px;
            text-align: center;
        }
    }
}
</style>