<template>
    <div class="wrapper">

        <div class="how">
            <h1>How it works?</h1>
            <div class="inner-wrapper">
                <div class="steps">
                    <div class="step">1</div>
                    <div class="content-wrapper">
                        <h3>Setup your wallet</h3>
                        <p>Set up your Metamask wallet. Download MetaMask from the browser extension</p>
                    </div>
                </div>
                <div class="arrow">
                    <img src="../../assets/imgs/arrow-drop-right-line.svg" alt="">
                </div>
                <div class="steps">
                    <div class="step">2</div>
                    <div class="content-wrapper">
                        <h3>Verfiy youself</h3>
                        <p>Please let us make sure your organization is legal.
                        </p>
                    </div>
                </div>
                <div class="arrow">
                    <img src="../../assets/imgs/arrow-drop-right-line.svg" alt="">
                </div>
                <div class="steps">
                    <div class="step">3</div>
                    <div class="content-wrapper">
                        <h3>Start Sharing</h3>
                        <p>Select the data that may help you in Plaza and get it at any time.
                            You can also upload your data to promote community development.</p>
                        <div class="content-warpper">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.wrapper {
    height: 95vh;
    background-color: #fff;
    scroll-snap-align: start;
    padding: 100px 0;
    position: relative;
}

.how {
    position: relative;

    h1 {
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 50%);
        font-size: 60px;
        color: #169608;
        margin-bottom: 50px;
    }

    .inner-wrapper {
        position: absolute;
        left: 50%;
        top: 50%;
        transform: translate(-50%, 30%);
        gap: 80px;
        display: flex;
        justify-content: space-around;
        
        .arrow {
            width: 50px;
            height: 300px;
            display: flex;
            justify-content: center;
            position: relative; 
            top: 50%;

            img {
                width: 200%;
                height: 200%;
            }
        }

        .steps {
            margin-top: 100px;
            width: 300px;
            height: 400px;
            background-color: #fff;
            box-shadow: 0 0 15px #E4E7ED;
            border-radius: 20px;
            background-repeat: no-repeat;
            background-position: center 0;
            position: relative;

            .step {
                width: 160px;
                height: 160px;
                background-color: #169608;
                color: #fff;
                font-size: 54px;
                font-weight: bold;
                border-radius: 50%;
                display: flex;
                justify-content: center;
                align-items: center;
                position: absolute;
                top: -30px;
                left: 50%;
                transform: translateX(-50%);

            }

            .content-wrapper {
                margin-top: 180px;
                width: 100%;
                text-align: center;
                padding: 10px 10px;

                h3 {
                    font-size: 24px;
                    color: #333;
                    font-weight: 700;
                }

                p {
                    margin-top: 20px;
                    color: #666;
                    padding: 10px 10px;
                }
            }
        }

    }
}
</style>