<template>
   <div class="body">
      <div class="banner">
         <h1 class="banner-title">Create</h1>
      </div>
      <div class="CreateArea">
         <div class="uploadDataBlock">
            <img src="../assets/imgs/upload.png" alt="upload">
            <p>Click here to upload your gene data</p>
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
body {
   background-color: #fff !important;
}

.body {
   margin: auto;
   width: 80%;
   // background-color: #fff;
}

.CreateArea {
   width: 100%;
   height: 600px;
   // background-color: grey;
   overflow: hidden;

   .uploadDataBlock {
      width: 400px;
      height: 350px;
      margin: 100px 0 0 300px;
      border-radius: 10px;
      border: 1px solid #E4E7ED;
      background-color: #fff;

      &:hover {
         cursor: pointer;
         box-shadow: 0px 0px 6px #DCDFE6;
         transition: 0.3s;
      }

      img {
         display: block;
         margin: 30px auto;
         color: #fff;


      }

      p {
         color: #C0C4CC;
         text-align: center;
      }
   }


}

.banner-title {
   font-size: 70px;
   color: #67C23A;
}
</style>