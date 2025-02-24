<template>
   <!-- banner区域 -->
   <div class="body">
      <div class="banner">
         <div class="banner-left">
            <h1 class="banner-title" style="margin-top: 200px">Discover Digital Art &<br> Collect NFTs</h1><br>
            <h2>NFT Marketplace UI create with Anima for Figma.<br> Collect, Buy and sell art from more than 20k NFT
               Artists</h2>
            <div class="getStartBTN">GET STARTED</div>
         </div>
         <img src="../assets/imgs/bioschains.png">
      </div>
   </div>

   <!-- 介绍  -->
   <div class="body">
      <div class="intro">
         <h1>Introduction</h1>
      </div>

   </div>

   <!-- 引导 -->
   <div class="body">
      <div class="howToDo">
         <h1>How it works</h1>
         <div class="howTodo-wrapper">
            <div class="steps">
               <div class="content-wrapper">
                  <h3>Setup your wallet</h3>
                  <p>Set up your wallet of choice. Connect it to the Animarket by clicking the wallet icon in the top
                     right corner.</p>
               </div>
            </div>
            <div class="steps">
               <div class="content-wrapper">
                  <h3>Create Collection</h3>
                  <p>Upload your work and setup your collection. Add a description, social links and floor price.</p>
               </div>
            </div>

            <div class="steps">
               <div class="content-wrapper">
                  <h3>Start Earning</h3>
                  <p>Choose between auctions and fixed-price listings. Start earning by selling your NFTs or trading others.</p>
                  <div class="content-warpper">
                  </div>
               </div>
            </div>
         </div>
      </div>
   </div>

</template>

<script>
export default {
   name: "create"
}
</script>

<style lang="scss" scoped>
.body {
   margin: auto;
   width: 70%;
   background-color: #fff;
   height: 780px;
   border-bottom: 1px solid #E4E7ED;
}

h1 {
   font-size: 70px;
   color: #67C23A;
}

.banner {
   display: flex;
   height: 700px;
   overflow: hidden;

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
      }
   }

   img {
      margin: 250px 0 0 300px;
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
            h3{
               font-size: 24px;
               color: #169608;
            }

            p{
               margin-top: 20px;
               color: #67C23A;
               padding: 10px 10px;
            }
         }
      }


   }
}
</style>