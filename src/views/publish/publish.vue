<template>
   <div class="wrapper">
      <Bubbles v-show="showIndexPage && step === -1" />
      <div class="wrapper-center" v-if="showIndexPage">
         <h1 class="title">Data <span style="color: #333;">Publish</span></h1>
         <span class="tip1">Ready to upload your data and publish it.</span>
         <div class="start-btn" @click="turnToSteps">Start</div>
         <div class="start-btn" style="background-color: #E6A23C;" @click="getProfileStatus">Profile status</div>
      </div>

      <div class="wrapper-left" v-if="step >= 0">
         <h1><span style="color: #169608;">Start journey</span> from here</h1>
         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 0">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 0">
            <div :class="{ 'step-tip': true, 'active': step >= 0 ? true : false }">Verify your account. We will send you
               a
               organization by email.
            </div>
         </div>

         <div class="line" />

         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 1">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 1">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 1">
            <div :class="{ 'step-tip': true, 'active': step >= 1 ? true : false }">Authorize your genetic
               reports. We ensure
               that no additonal personal data
               will be stored by your platform.If authorized, you can continue. </div>
         </div>

         <div class="line" />

         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 2">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 2">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 2">
            <div :class="{ 'step-tip': true, 'active': step >= 2 ? true : false }">Create your unique data. Once your
               item is minted
               you will not be able to change
               any of its information.</div>
         </div>

         <div class="line" />

         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 3">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 3">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 3">
            <div :class="{ 'step-tip': true, 'active': step >= 3 ? true : false }" style="line-height: 28px;">Offer your
               data on the plaza. Prior listing,
               you may delete your data; However, once listed, you will not be able to delete your data.</div>
         </div>
      </div>

      <div class="wrapper-right" v-if="step === 0">
         <el-button @click="toVerify" class="custom-button" style="width: 400px;">
            Click to verify your organization
         </el-button>
      </div>

      <div class="wrapper-right" style="display: block;" v-if="step === 1">
         <div class="platforms" @click="redirectToOAuth">
            <img src="../../icons/wegene_logo.svg" alt="wegene">
         </div>
         <div class="platforms">
            <p>comming soon...</p>
         </div>
         <div class="platforms">
            <p>comming soon...</p>
         </div>
         <div class="button-wrapper" style="height:150px">
            <el-button class="button" @click="turnToIndexPage" style="right: 120px;">Back</el-button>
            <el-button class="button" @click="step++" style="right: 0;">Continue</el-button>
            <div class="tips">
               If your genetic testing platform is not listed, please to provide feedback, and we will address it
               promptly.
            </div>
         </div>
      </div>

      <div class="wrapper-right" v-if="step === 2">
         <div class="form-wrapper">
            <p>Report *</p>
            <div class="select" @click="getCompleted">
               <div class="add" v-if="!selectedProfile">+</div>
               <div v-if="!selectedProfile">Select the report as collection</div>
               <div v-if="selectedProfile">{{ selectedProfile }}</div>
            </div>
            <p>Name *</p>
            <el-input class="name-input" v-model="name"
               placeholder="Please enter the name of the collection"></el-input>
            <!-- <p class="introduction"><span class="click-here">Click here</span> to view an example.</p> -->
            <p>Description</p>
            <input type="text" class="Description-input" placeholder="Please enter a description of the collection"
               v-model="description"></input>
            <div class="button-wrapper" style="width: 100%; margin-top: 50px;">
               <el-button class="button" @click="step--" style="right: 120px;">Back</el-button>
               <el-button class="button" @click="createData" style="right: 0px;">Create</el-button>
            </div>
         </div>

      </div>

      <div class="wrapper-right" style="flex-direction: column;gap:30px;transform: translateY(-10%);" v-if="step === 3">
         <h1 style="color: #169608; font-size: 30px;">Success! Thank you for your support!</h1>
         <div class="upload-hash">Hash:{{ hash }}
            <div style="cursor: pointer;" @click="pasteHash">
               <img src="@/icons/paste.svg" alt="paste_icon">
            </div>
         </div>
         <p>Use this hash, you can view the block on the FISCO BCOS broswer.</p>
         <div class="button-wrapper" style="margin-top: 50px;">
            <el-button class="button" @click="step--">Back</el-button>
         </div>
      </div>

   </div>

   <el-drawer v-model="drawerIsVisible" :element-loading-svg="svg" title="Profile status" @close="handleClose"
      :show-close="false">
      <el-table v-loading="loadingForComplete" :data="profiles" @selection-change="handleSelectionChange">
         <el-table-column label="Completed reports">
            <template #default="{ row }">
               <el-radio v-model="selectedProfile" :label="row" /> 
            </template>

         </el-table-column>
      </el-table>
      <el-table :data="uploadProgress" :element-loading-svg="svg" max-height="65vh">
         <el-table-column prop="taskID" label="Task ID" width="380" />
         <el-table-column prop="progress" label="Progress(%)" width="130" style="color: #169608;"/>
      </el-table>
      <template #footer v-if="step === 2">
         <div style="flex: auto">
            <p class="selected-profile">Selected Profile: <span>{{ selectedProfile }}</span></p>
            <el-button @click="drawerIsVisible = false" class="confirm-button">Confirm</el-button>
         </div>
      </template>
   </el-drawer>
</template>


<script lang="js" setup>
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import Bubbles from '../components/bubbles.vue';
import Api from '../../axios/aixos';
import { useWalletStore } from '@/stores/account';
import { ElLoading, ElMessage } from 'element-plus';
import { SSEManager } from './SSE.js';
const walletStore = useWalletStore();
const router = useRouter();
const route = useRoute()

const step = ref(-1); // -1:首页
const showIndexPage = ref(true);    //是否展示首页，因为通过step维护的状态流转wrapper并不是全覆盖

/**离开首页 */
function turnToSteps() {
   showIndexPage.value = false;
   if (showIndexPage.value === false && walletStore.insititution) {
      step.value += 1;
   }
   step.value += 1;
}
/**返回首页 */
function turnToIndexPage() {
   showIndexPage.value = true;
   step.value = -1;
}

function toVerify() {
   router.push('/verify');
}
/**重定向请求 */
function redirectToOAuth() {
   window.location.href = `${import.meta.env.VITE_APP_BASE_URL}/user/oauth2Wegene?t=${Date.now()}`;
}

/**
 * 获取已完成的报告 && 上传进度
 * @param {string} profiles   获取的已完成报告
 * @param {bool}  drawerIsVisible   进度抽屉是否展示，主页和create页复用同一个drawer
 * @param {bool}  loadingForProgress
 */
const profiles = ref([]);
const drawerIsVisible = ref(false);
const loadingForComplete = ref(false);
const getCompleted = async () => {
   drawerIsVisible.value = true;
   loadingForComplete.value = true;
   try {
      const res = await Api.get('/studio/getProfile/completed');
      if (res.data.code === 200) {
         profiles.value = res.data.data.profile_ids
         loadingForComplete.value = false;
      }
   } catch (error) {
      loadingForComplete.value = false;
      alert("Network error. Please try again later");
      console.error(error);
   }
}

/**
 * 获取上传进度 && 已上传好数据列表
 * @param {bool} drawerIsVisible 上传进度抽屉是否展示
 * @param {function} getCompleted 获取已上传好数据列表，无参无返
 * @param {function} getUncompleted 获取正在上传进度，SSE通道，无参返回sseRequest对象
 * @param {array} uploadProgress 上传进度
 * @param {object} sseConnection  SSE连接,使用模块级变量存储连接，便于卸载管理
 */
const loadingForProgress = ref(false);
const getProfileStatus = () => {
   drawerIsVisible.value = true;
   getCompleted();
   getUncompleted();
}
const uploadProgress = ref([]);
let sseConnection = new SSEManager(uploadProgress);   //传入一个ref
function getUncompleted() {
   loadingForProgress.value = true;
   sseConnection.connect();
}
//drawer关闭时关闭SSE连接，handleClose为drawer绑定的清理事件
function handleClose() {
   sseConnection.close();
}

/**
 * create创建sharing数据步骤
 * @param {string} selectedProfile 选择的报告
 * @param {string} name 报告名称
 * @param {string} description 报告描述
 * @param {string} tags 报告标签
 */
let selectedProfile = ref(null);
let name = ref('');
let description = ref('');
let tags = ref('');
const hash = ref();
const createData = async () => {
   let loading = ElLoading.service({
      lock: true,
      text: 'Loading...',
      background: 'rgba(0, 0, 0, 0.7)',
      customClass: 'loading',
   })
   try {
      if (name.value === '' || name.value === null || name.value === undefined) {
         alert('Please enter the name of the collection');
         return;
      }
      if (selectedProfile.value === null || selectedProfile.value === '' || selectedProfile.value === undefined) {
         alert('Please select the report as collection');
         return;
      }
      if (description.value === '' || description.value === null || description.value === undefined) {
         alert('Please enter a description of the collection');
         return;
      }

      const res = await Api.post('/studio/createFromThirdParty', {
         profile_id: selectedProfile.value,
         name: name.value,
         description: description.value,
         tags: tags.value
      });
      if (res.data.code === 200) {
         ElMessage.success('Data created successfully');
         loading.close();
         hash.value = res.data.data.transaction_hash;
         step.value += 1;
         pasteHash();
      }
   } catch (error) {
      console.error('Error creating data:', error);
      ElMessage.error('Error creating data');
   } finally {
      loading.close();
   }
};

/**
 * 复制交易hash
 * @param {string} hash MetaMask交易成功后获取的交易hash，用于校验
 */
const pasteHash = () => {
   navigator.clipboard.writeText(hash.value);
   ElMessage.success('Hash copied successfully');
}

onMounted(() => {
})
</script>

<style lang="scss" scoped>
.wrapper {
   height: 95vh;
   width: 100vw;
   position: relative;
   min-width: 1200px !important;
   overflow: hidden;
   gap: 1px;
}

.active {
   color: #67C23A !important;
}

.wrapper-center {
   height: 95vh;
   width: 100vw;
   min-width: 1200px;
   display: flex;
   position: absolute;
   flex-direction: column;
   gap: 30px;
   justify-content: center;
   align-items: center;
   animation: fadeIn 0.2s ease-in-out 0s forwards;
   transform: translateY(-5%);

   @keyframes fadeIn {
      0% {
         opacity: 0;
      }

      100% {
         opacity: 1;
      }
   }

   .title {
      font-size: 70px;
      color: #169608;
   }

   .tip1 {
      font-size: 20px;
      color: #333;
   }

   .start-btn {
      width: 240px;
      height: 60px;
      line-height: 60px;
      background-color: #169608;
      color: #fff;
      font-size: 24px;
      border-radius: 50px;
      text-align: center;
      cursor: pointer;


   }
}


h1 {
   margin: 10px;
   color: #333;
   font-size: 36px;
}

.wrapper-left {
   position: absolute;
   left: 10%;
   top: 15%;
   width: 40%;
   min-width: 500px;
   padding: 20px;
   height: 65vh;
   animation: fadeIn 0.2s ease-in-out 0s forwards;

   @keyframes fadeIn {
      0% {
         opacity: 0;
      }

      100% {
         opacity: 1;
      }
   }

   .title {
      font-size: 70px;
      color: #169608;
      margin-bottom: 20px;
   }

   .inside-step {
      display: flex;
      width: 100%;
      height: 80px;

      img {
         width: 60px;
      }

      .step-tip {
         display: flex;
         align-self: center;
         font-size: 20px;
         margin-left: 25px;
         color: #A8ABB2;
      }

      .step-tip-ok {
         display: flex;
         align-self: center;
         font-size: 20px;
         margin-left: 25px;
         color: #67C23A;
      }
   }

   .line {
      margin-left: 28px;
      width: 3px;
      border-radius: 5px;
      height: 30px;
      background-color: #A8ABB2;
   }
}

.wrapper-right {
   position: absolute;
   overflow: auto;
   left: 50%;
   top: 15%;
   width: 30%;
   min-width: 600px !important;
   height: 75vh !important;
   overflow: auto;
   animation: fadeIn 0.2s ease-in-out 0s forwards;

   @keyframes fadeIn {
      0% {
         opacity: 0;
      }

      100% {
         opacity: 1;
      }
   }

   .platforms {
      margin: 50px auto;
      display: flex;
      width: 60%;
      min-width: 240px;
      height: 100px;
      border: #e0e0e0 1px solid;
      border-radius: 20px;
      align-items: center;
      box-shadow: 0px 0px 2px 0px #e0e0e0 !important;

      &:hover {
         box-shadow: 0px 0px 15px 1px #e0e0e0 !important;
         cursor: pointer;
      }

      img {
         margin: 0 auto;
         height: 60px;
      }

      p {
         line-height: 24px;
         width: 100%;
         text-align: center;
         color: #cdcdcd;
         font-size: 24px;
      }
   }

   .form-wrapper {
      position: relative;
      width: 60%;
      height: 80%;
      display: flex;
      flex-direction: column;
      margin: 0px auto;

      p {
         margin-top: 15px;
         line-height: 20px;
      }
   }

   .button-wrapper {
      display: flex;
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 60%;

      .tips {
         position: absolute;
         bottom: 0;
      }
   }

   .upload-hash {
      color: #888;
      display: flex;

      img {
         width: 16px;
      }
   }

}

:deep(.button) {
   position: absolute;
   width: 100px;
   height: 42px;
   background-color: #169608;
   color: #fff;
   font-size: 18px;
   border-radius: 10px;
   border: none !important;

   &:hover {
      color: #169608;
      background-color: #fff;
      border: #169608 1px solid !important;
      cursor: pointer;
   }
}

.select {
   padding: 10px;
   display: flex;
   align-items: center;
   margin: 10px 0;
   width: 100%;
   height: 60px;
   background-color: #f9f8f8;
   border: #e0e0e0 1px solid;
   border-radius: 10px;
   box-shadow: 0px 0px 2px 0px #e0e0e0 !important;

   &:hover {
      cursor: pointer;
   }

   .add {
      margin-right: 20px;
      display: flex;
      align-items: center;
      justify-content: center;
      width: 50px;
      height: 50px;
      border: #cdcdcd 2px solid;
      border-radius: 10px;
   }

}

.introduction {
   color: #666666;
   margin: 5px 0 20px 5px;
}

:deep(.name-input) {
   margin: 10px 0;
   width: 100%;
   height: 50px;
}

:deep(.supply-input) {
   margin: 10px 0;
   width: 200px;
   height: 50px;
}

.Description-input {
   margin: 10px 0;
   width: 100%;
   height: 100px;
   border: #e0e0e0 1px solid;
   border-radius: 10px;

}

.click-here {
   color: #67C23A;
   cursor: pointer;

   &:hover {
      text-decoration: underline;
      font-size: 16px;
   }

}

.add-trait {
   color: #67C23A;

   &:hover {
      cursor: pointer;

   }
}

.selected-profile {
   text-align: left;
   color: #333;

   span {
      color: #169608;
   }
}
</style>

<style lang="scss" scoped>
.wrapper-right {
   display: flex;
   position: absolute;
   right: 0;
   width: 50%;
   min-width: 600px;
   height: calc(100vh - 60px);
   justify-content: center;
   align-items: center;
   background-color: #fff;
}

.flags {
   width: 24px;
   height: 18px;
   margin-top: 5px;
}


:deep(.el-form-item__label) {
   font-size: 16px;
   color: #909399;
}


:deep(.el-input__wrapper) {
   width: 100%;
   border-radius: 10px;

   &:focus-within {
      box-shadow: 0px 0px 0px 1px inset #67C23A !important;
   }
}

:deep(.el-input__inner) {
   height: 40px;
   padding: 0;
   border: none;

   &:focus {
      color: #67C23A;
   }
}

:deep(.el-input-group__prepend) {
   width: 10px;
   border: none;
   background-color: #fff;

}

:deep(.el-select__placeholder) {
   width: 40px;
}

:deep(.el-input-group__append) {
   border: none;

}


// 下拉框选择器
:deep(.el-tooltip__trigger) {
   padding: 10px;
   height: 42px;
   border: none !important;
   box-shadow: none !important;

   &:hover {
      border: none;
      box-shadow: none;
   }
}

:deep(.el-select__selected-item) {
   color: #000;
   line-height: 70px;


}

:deep(.el-button--primary) {
   padding: 0;
   width: 64px;
   height: 42px;
   background-color: #fff;
}

:deep(.custom-button) {
   width: 80%;
   height: 42px;
   background-color: #169608;
   color: #fff;
   font-size: 18px;
   border-radius: 5px;
   border: none !important;
   // box-shadow: 0px 0px 5px 0px #e0e0e0 !important;

   &:hover {
      background-color: #169608;
      color: #fff;
      box-shadow: 0px 0px 10px 5px #e0e0e0 !important;
      cursor: pointer;
   }
}

:deep(.el-table__row) {
   height: 40px;

   .el-radio__inner {
      background: #fff;
      border-color: #169608;

      &::after {
         background-color: #169608;
      }
   }

   .el-radio__label {
      color: #169608;
      font-size: 16px;
   }
}

:deep(.confirm-button) {
   border: none !important;
   background-color: #169608 !important;
   color: #fff !important;
   font-size: 16px;
}
</style>