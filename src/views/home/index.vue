<template>
   <div class="container" ref="container" @wheel="handleWheel">
      <IndexBanner />

      <div class="wrapper">
         <banner>
            <div class="banner-left">
               <h1 style="margin-top: 150px">Gain & Share <br>Gene code.</h1><br>
               <h2>A Biological alliance-chain based on FISCO BCOS.
                  <br>
                  Unlocking the power of genetic data.
               </h2>

               <div class="getStartBTN" @click="">Git Hub↗</div>

            </div>
            <img src="../assets/imgs/bioschains.svg">
         </banner>
      </div>

      <!-- 介绍  -->
      <div class="wrapper">
         <div class="intro">
            <h1>Introduction</h1>
         </div>
      </div>

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


   </div>

</template>

<script setup>
import IndexBanner from './banner.vue';
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
* {
   /* 隐藏默认的滚动条样式 */
   scrollbar-width: none;
   /* Firefox */
   -ms-overflow-style: none;
   /* IE and Edge */
}

h1 {
   font-size: 50px;
   color: #169608;
}

.container {
   height: 95vh;
   width: 100vw;
   overflow-y: scroll;
   scroll-snap-type: y mandatory;
}

.wrapper {
   height: 95vh;
   margin: auto;
   background-color: #fff;
   scroll-snap-align: start;
   padding: 100px 0;
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