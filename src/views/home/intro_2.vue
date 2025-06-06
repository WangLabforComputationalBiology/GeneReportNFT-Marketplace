<template>
    <div class="wrapper">
        <div class="wrapper-left">
            <h1>How we do that?</h1><br>
            <h2>
                We leverage a non-cryptocurrency blockchain for implementation, with FISCO BCOS providing robust
                technical support. Through blockchain and smart contracts, we enable the recording and traceability of
                genetic information transfers. The process of logging transfer records on the chain is decentralized,
                while the storage of genetic data remains centralized—a decision made for regulatory compliance and
                security considerations. We advocate for the participation of universities and legitimate institutions,
                with the goal of advancing academic research.
                <br>
                <br>
                <span class="tags">#Coinless blockchain</span>
                <span class="tags">#Traceable</span>
                <!-- <span class="tags">#Academic</span> -->
            </h2>



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
    left: 50%;
    top: 50%;
    transform: translate(-50%, -75%);


    h1 {
        font-size: 6vh;
        font-weight: 500;
        color: #169608;
        margin-bottom: 20px;
    }

    h2 {
        font-size: 2.5vh;
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
</style>