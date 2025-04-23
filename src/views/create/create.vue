<template>
   <div class="body">
      <div class="banner">
         <h1 class="banner-title">Create</h1>
      </div>
      <h2 style="padding-left:20px; color:#99a9bf">Your collections:</h2>
      <div class="CreateArea">
         <!-- wegene数据库 -->
         <el-card class="custom-card" v-for="(item, index) in cardList" :key="index">
            <template #header>
               <div class="card-header">
                  <span style="height: 30px;">Wegene</span>
               </div>
            </template>

            <div class="card-info">
               <p>Name: {{ item.name }}</p>
               <p>Date: {{ item.date }}</p>
               <p>Description: {{ item.description }}</p>
            </div>

            <template #footer>
               <div class="card-buttons">
                  <el-button type="primary" class="custom-btton">Create</el-button>
                  <el-button type="primary" class="custom-btton" @click="handleDownload">Download</el-button>
               </div>
            </template>
         </el-card>
      </div>
   </div>
</template>

<script>
export default {
   name: "create",
   data() {
      return {
         cardList: [
            {
               name: 'Card 1',
               date: '2023-01-01',
               description: 'First card description'
            },
            {
               name: 'Card 2',
               date: '2023-01-02',
               description: 'Second card description'
            },
            {
               name: 'Card 1',
               date: '2023-01-01',
               description: 'First card description'
            },
            {
               name: 'Card 2',
               date: '2023-01-02',
               description: 'Second card description'
            },
            {
               name: 'Card 1',
               date: '2023-01-01',
               description: 'First card description'
            },
            {
               name: 'Card 2',
               date: '2023-01-02',
               description: 'Second card description'
            },
            {
               name: 'Card 1',
               date: '2023-01-01',
               description: 'First card description'
            },
            {
               name: 'Card 2',
               date: '2023-01-02',
               description: 'Second card description'
            },
            {
               name: 'Card 1',
               date: '2023-01-01',
               description: 'First card description'
            },
            {
               name: 'Card 2',
               date: '2023-01-02',
               description: 'Second card description'
            },
         ]
      }
   }
}
</script>

<style lang="scss" scoped>
.body {
   margin: auto;
   width: 80vw;
   min-height: calc(100vh - 80px);
   min-width: 1200px;
}

/* 卡片区域 */
.CreateArea {
   display: flex;
   flex-wrap: wrap;
   gap: 40px;
   width: 100%;
   min-width: 1200px;
}

/* 卡片样式 */
:deep(.custom-card) {
   // margin: 10px 10px;
   width: 250px;
   height: 320px;
   box-shadow: none;
   border-radius: 10px;

   .card-info {
      width: 100%;
      height: 150px;

      p {
         margin: 2px;
         // color: #67C23A;
      }
   }

   .footer {
      display: flex;
      height: 40px;
   }
}

:deep(.custom-card):hover {
   box-shadow: 0px 0px 16px #DCDFE6;
}

:deep {
   .el-card__header {
      padding: auto 5px;
      // background-color: #F5F7FA;
      color: #67C23A;

   }
}

:deep(.el-card__body) {
   display: flex;

   position: relative;
   height: 180px;

}

.card-buttons {
   display: flex;
   margin: auto;
   width: 100%;
   justify-content: space-between;

}

:deep(.custom-btton) {
   width: 120px;
   height: 40px;
   line-height: 6px;
   background-color: #E6A23C;
   border: none;

   &:first-child {
      background-color: #67C23A;

      &:hover {
         background-color: #67C23Aee;
         cursor: pointer;
      }

   }

   &:hover {
      background-color: #e6a23cee;
      cursor: pointer;
   }
}

.banner-title {
   font-size: 70px;
   color: #67C23A;
}
</style>