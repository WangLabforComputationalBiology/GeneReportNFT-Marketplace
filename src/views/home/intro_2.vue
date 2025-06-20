<template>
  <div class="wrapper">
    <div class="wrapper-left">
      <h1>How <span style="color: #333">we do that?</span></h1>
      <br />
      <h2>
        Through blockchain and smart contracts, we enable the recording and traceability
        of genetic information transfers. The process of logging transfer records on the
        chain is decentralized, while the storage of genetic data remains centralized—a
        decision made for regulatory compliance and security considerations.
        <br />
        We advocate for the participation of universities and legitimate institutions,
        with the goal of advancing academic research.
        <br />
        <br />
        <span class="tags">#Coinless blockchain</span>
        <span class="tags">#Traceable</span>
      </h2>
    </div>


    <div class="wrapper-right">
      <div class="des-btn-wrapper">
        <div class="des-btn" @mouseover="isVisible1 = true" @mouseleave="isVisible1 = false">1</div>
        <div class="des-btn" @mouseover="isVisible2 = true" @mouseleave="isVisible2 = false">2</div>
      </div>
      <img src="@/assets/imgs/framework.svg" alt="" />
      <div class="des" v-if="isVisible1">1. The data is packaged and put into the database by institution, and the
        process will be recorded on the blockchain to achieve information traceability.</div>
      <div class="des" v-if="isVisible2">2. Verified users can submit applications for obtaining data to the blockchain,
        and then use the chain information as a credential to obtain data from the database.</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const isVisible1 = ref(false);
const isVisible2 = ref(false);
</script>

<style lang="scss" scoped>
.wrapper {
  height: 95vh;
  min-width: 1200px;
  scroll-snap-align: start;
  display: flex;
  position: relative;
}

.fullpage-swiper {
  pointer-events: auto !important;
  width: 100%;
  height: 100%;
}

.swiper-slide {
  width: 100%;
  height: 100%;
}

.wrapper-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  position: absolute;
  width: 50%;
  min-width: 600px;
  height: 95vh;
  border-bottom: #fff 1px solid;
  left: 0%;
  background-color: #169608;


  h1 {
    font-size: 6vh;
    font-weight: 500;
    color: #fff;
    margin-bottom: 20px;
  }

  h2 {
    min-width: 445px;
    width: 35vw;
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
  width: 50%;
  min-width: 600px;
  height: 95vh;
  background-color: #fff;
  pointer-events: auto;

  img {
    min-width: 350px;
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

  .des-btn-wrapper {
    display: flex;
    position: absolute;
    // width: 120px;
    height: 5vh;
    bottom: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
    gap: 20px;

    .des-btn {
      width: 5vh;
      font-size: 20px;
      font-weight: 700;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 50%;
      color: #67c23a;
      background-color: #fff;
      border: #333333dd 3px solid;
      cursor: pointer;

      &:hover {
        box-shadow: 0 0 15px #ccc;
      }

      &:nth-child(2) {
        color: #E6A23C;
        background-color: #fff;
      }
    }
  }

  .des {
    position: absolute;
    bottom: 2%;
    left: 50%;
    transform: translate(-50%, 0%);
    font-size: 1.8vh;
    font-weight: 700;
    color: #333;
    width: 25vw;
    height: 18vh;
  }
}
</style>
