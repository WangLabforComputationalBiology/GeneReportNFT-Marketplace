<template>
   <div class="container" ref="container" @wheel="handleWheel">
      <div class="wrapper banner_img" style="">
         <div class="green-bar"></div>
         <div class="main-text">Bio-ifo <br>Sharing</div>
         <div class="sec-text">A platform delivering genetic information.</div>
         <div class="scroll-arrow"><span class="arrow-fadein">↓</span></div>
      </div>


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
   width: 80vw;
   margin: auto;
   background-color: #fff;
   scroll-snap-align: start;
   padding: 100px 0;
}

.banner_img {
   position: relative;
   margin: 0;
   padding: 0;
   width: 99vw;
   height: 95vh;
   background-size: cover;
   background-position: center;
   justify-content: center;
   align-items: center;


   /* 绿色条块 */
   .green-bar {
      position: absolute;
      top: 48%;
      left: 0;
      width: 0;
      /* 初始宽度为0 */
      height: 200px;

      /* 条块高度 */
      background: #169608;
      /* 绿色 */
      transform: translateY(-50%);
      animation: extend 0.8s ease-out forwards;
      /* 延长动画 */
      border-top-right-radius: 50px;
   }

   /* 延长动画 */
   @keyframes extend {
      0% {
         width: 0;
      }

      100% {
         width: 52%;
         /* 延长到容器宽度的一半 */
      }
   }

   /* 主体文字 */
   .main-text {
      position: absolute;
      top: 24%;
      left: 25%;
      transform: translate(-50%, -50%);

      line-height: 200px;
      font-size: 10rem;
      color: #333;
      opacity: 0;
      /* 初始不可见 */
      animation: fadeIn 0.6s ease-in forwards 0.8s;
      /* 延迟0.8s淡入 */
   }

   .sec-text {
      position: absolute;
      top: 50%;
      left: 52.5%;
      transform: translate(-50%, -50%);

      line-height: 200px;
      font-size: 2rem;
      color: #333;
      opacity: 0;
      /* 初始不可见 */
      animation: fadeIn 0.6s ease-in forwards 1.2s;
      /* 延迟0.8s淡入 */
   }



   /* 淡入动画 */
   @keyframes fadeIn {
      0% {
         opacity: 0;
         transform: translateY(20px);
         /* 轻微向下偏移 */
      }

      100% {
         opacity: 1;
         transform: translateY(0);
      }
   }
}


/* 跳跃箭头 */
.scroll-arrow {
   position: absolute;
   bottom: 10%;
   /* 位于容器中间偏下 */
   left: 50%;
   transform: translateX(-50%);
   /* 水平居中 */
   font-size: 2rem;
   color: #333;
   animation: bounce 2.5s infinite;
   /* 跳跃动画，无限循环 */

   .arrow-fadein {
      opacity: 0;
      animation: fadeIn 2.5s infinite;
      /* 箭头淡入动画 */
   }
}

/* 跳跃动画 */
@keyframes bounce {
   0%,
   20%,
   50%,
   80%,
   100% {
      transform: translate(-50%, 0);
      /* 初始位置 */
   }

   40% {
      transform: translate(-50%, -20px);
      /* 向上跳跃 */
   }

   60% {
      transform: translate(-50%, -10px);
      /* 轻微回落 */
   }
}

banner {
   display: flex;
   justify-self: center;

   .banner-left {

      h2 {
         font-size: 25px;
         color: #909399;
      }

      .getStartBTN {
         color: #333;
         width: 200px;
         height: 60px;
         background-color: #fff;
         box-shadow: 0 0 5px #E4E7ED;
         font-size: 24px;
         border-radius: 10px;
         text-align: center;
         align-content: center;

         margin-top: 40px;
         cursor: pointer;

         &:hover {
            box-shadow: 0 0 10px #E4E7ED;
         }
      }
   }

   img {
      margin: 220px 0 0 300px;
      width: 300px;
      height: 300px;
   }
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