<template>
   <div class="container" ref="container" @wheel="handleWheel">

      <div class="wrapper">
         <banner>
            <div class="banner-left">
               <h1 style="margin-top: 150px">Gain & Share <br>Gene code.</h1><br>
               <h2>A bio-alliance chain based on FISCO BCOS.
                  <br>
                  Unlocking the power of genetic data.
               </h2>
               <div class="BTN-wrapper" style="display: flex;">
                  <div class="getStartBTN" @click="scrollToSection">GET STARTED</div>
                  <div class="getStartBTN" @click="scrollToSection">Git Hub</div>
               </div>

            </div>
            <img src="../assets/imgs/bioschains.svg">
         </banner>
      </div>

      <!-- 跳转参考线 -->
      <hr style="border:none; " ref="targetSection">

      <!-- 引导 -->
      <div class="wrapper">
         <div class="howToDo">
            <h1>How it works</h1>
            <div class="howTodo-wrapper">
               <div class="steps">
                  <div class="content-wrapper">
                     <h3>Setup your wallet</h3>
                     <p>Set up your Metamask wallet. Connect it to us by clicking the wallet icon in the top
                        right corner.</p>
                  </div>
               </div>
               <div class="steps">
                  <div class="content-wrapper">
                     <h3>Create Collection</h3>
                     <p>Upload your Wegene data and setup your collection. Add a description, social links and floor
                        price.
                     </p>
                  </div>
               </div>

               <div class="steps">
                  <div class="content-wrapper">
                     <h3>Start Sharing</h3>
                     <p>Choose fixed-price listings. Get revenue from your sharing or trading
                        others.</p>
                     <div class="content-warpper">
                     </div>
                  </div>
               </div>
            </div>
         </div>
      </div>

      <!-- 介绍  -->
      <div class="wrapper">
         <div class="intro">
            <h1>Introduction</h1>
            <el-carousel height="650px" interval="0">
               <el-carousel-item v-for="item in 4" :key="item">
                  <h3 class="small">{{ item }}</h3>
               </el-carousel-item>
            </el-carousel>
         </div>
      </div>
   </div>

</template>

<script setup>
import { ref } from 'vue'
// 滚动状态
const isScrolling = ref(false);
const activeSection = ref(0);
const container = ref(null);

// 处理滚轮事件
const handleWheel = () => {
   if (isScrolling.value) return;
   isScrolling.value = true;

   setTimeout(() => {
      isScrolling.value = false;
      console.log('滚动结束');
   }, 100);

   let nextIndex = activeSection.value;

   scrollToSection(nextIndex);
};

// 滚动到指定section
const scrollToSection = (index) => {
   const section = document.getElementById(`section${index + 1}`);
   activeSection.value = index;
};

// 监听滚动更新导航点
const updateNavDots = () => {
   const scrollTop = container.value.scrollTop;
   let currentIndex = 0;

   sections.forEach((_, index) => {
      const section = document.getElementById(`section${index + 1}`);
      const sectionTop = section.offsetTop;
      if (scrollTop >= sectionTop - window.innerHeight / 2) {
         currentIndex = index;
      }
   });

   activeSection.value = currentIndex;
};

// 绑定滚动事件
container.value?.addEventListener('scroll', updateNavDots);
</script>



<style lang="scss" scoped>
.container {
   height: 100vh;
   overflow-y: scroll;
   scroll-snap-type: y mandatory;
}

.wrapper {
   margin: 60px auto;
   width: 100vw;
   background-color: #fff;
   height: 100vh;


   font-size: 2rem;
   scroll-snap-align: start;


}

h1 {
   font-size: 70px;
   color: #67C23A;
}

banner {
   display: flex;

   .banner-left {

      h2 {
         font-size: 25px;
         color: #909399;
      }

      .getStartBTN {
         width: 200px;
         height: 60px;
         background-color: #fff;
         color: #67C23A;
         box-shadow: 0 0 5px #E4E7ED;
         font-size: 24px;
         border-radius: 10px;
         text-align: center;
         align-content: center;

         margin-top: 40px;
         cursor: pointer;

         &:hover {
            background-color: #67C23Add;
         }
      }
   }

   img {
      margin: 220px 0 0 300px;
      width: 300px;
      height: 300px;
   }
}

.intro {

   h1 {
      margin-top: 50px;
      font-size: 70px;
      color: #67C23A;
   }
}

//走马灯样式
.el-carousel__item h3 {
   color: #475669;
   font-size: 14px;
   opacity: 0.75;
   line-height: 150px;
   margin: 0;
}

.el-carousel__item:nth-child(2n) {
   background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
   background-color: #d3dce6;
}

.howToDo {
   .howTodo-wrapper {
      width: 100%;
      display: flex;
      justify-content: space-around;
      margin: auto;

      .steps {
         margin-top: 100px;
         width: 300px;
         height: 400px;
         background-color: #fff;
         box-shadow: 0 0 5px #E4E7ED;
         border-radius: 10px;
         background-image: url(../assets/imgs/steps.png);
         background-repeat: no-repeat;
         background-position: center 0;

         .content-wrapper {
            margin-top: 200px;
            width: 100%;
            text-align: center;
            padding: 10px 10px;

            h3 {
               font-size: 24px;
               color: #169608;
            }

            p {
               margin-top: 20px;
               color: #67C23A;
               padding: 10px 10px;
            }
         }
      }

   }
}
</style>