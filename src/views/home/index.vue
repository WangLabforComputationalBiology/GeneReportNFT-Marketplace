<template>
   <div class="container">
      <IndexBanner />

      <Intro1 />

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
import Intro1 from './intro_1.vue';
import { ref } from 'vue'

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