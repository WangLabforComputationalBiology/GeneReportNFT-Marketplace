<template>
   <!-- banner区域 -->
   <body>
      
   
   <div class="wapper">
      <banner>
         <div class="banner-left">
            <h1 style="margin-top: 150px">Explore genetic code &<br> Collect NFTs</h1><br>
            <h2>NFT Marketplace UI create with Anima for Figma.<br> Collect, Buy and sell art from more than 20k NFT
               Artists</h2>
            <div class="getStartBTN" @click="scrollToSection">GET STARTED</div>
         </div>
         <img src="../assets/imgs/bioschains.svg">
      </banner>
   </div>

   <!-- 跳转参考线 -->
   <hr style="border:none; " ref="targetSection">

   <!-- 引导 -->
   <div class="wapper">
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
                  <p>Upload your Wegene data and setup your collection. Add a description, social links and floor price.
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
   <div class="wapper">
      <div class="intro">
         <h1>Introduction</h1>
         <el-carousel height="650px" interval="0">
            <el-carousel-item v-for="item in 4" :key="item">
               <h3 class="small">{{ item }}</h3>
            </el-carousel-item>
         </el-carousel>
      </div>
   </div>


</body>
</template>

<script>
import { ref } from 'vue'

// 使用 ref 绑定目标元素
const targetSection = ref(null)
export default {
   setup() {
      // 使用 ref 定义 targetSection
      const targetSection = ref(null)

      // 滚动方法
      const scrollToSection = () => {
         if (targetSection.value) {
            targetSection.value.scrollIntoView({ behavior: 'smooth' })
         }
      }

      // 将 ref 和方法暴露给模板
      return {
         targetSection,
         scrollToSection
      }
   }
}
</script>



<style lang="scss" scoped>
body{
   overflow: auto;
}
.wapper {
   margin: 60px auto;
   width: 1400px;
   background-color: #fff;
   height: calc(100vh - 60px);
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
         background-color: #E6A23C;
         color: #fff;
         font-size: 24px;
         border-radius: 10px;
         text-align: center;
         align-content: center;

         margin-top: 40px;
         cursor: pointer;

         &:hover {
            background-color: #e6a23cdd;
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