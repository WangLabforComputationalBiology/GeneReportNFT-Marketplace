<template>
   <div class="wrapper">
      <Bubbles v-if="showIndexPage" />
      <div class="wrapper-center" v-if="showIndexPage">
         <h1 class="title">Data <span style="color: #333;">Publish</span></h1>
         <span class="tip1">Ready to upload your data and publish it.</span>
         <div class="start-btn" @click="turnToSteps">Start</div>
      </div>

      <div class="wrapper-left" v-if="step >= 0">
         <h1><span style="color: #169608;">Start your journey</span> from here</h1>
         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 0">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 0">
            <div :class="{ 'step-tip': true, 'active': step >= 0 ? true : false }">Verify your account. We will send you
               a
               organization by email.
            </div>

         </div>
         <div class="line"></div>
         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 1">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 1">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 1">
            <div :class="{ 'step-tip': true, 'active': step >= 1 ? true : false }">Allow us to access your genetic
               reports. We ensure
               that no additonal personal data
               will be stored by your platform.</div>

         </div>
         <div class="line"></div>
         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 2">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 2">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 2">
            <div :class="{ 'step-tip': true, 'active': step >= 2 ? true : false }">Create your unique data. Once your
               item is minted
               you will not be able to change
               any of its information.</div>

         </div>
         <div class="line"></div>
         <div class="inside-step">
            <img src="../../icons/status_ing.svg" alt="status icon" v-if="step == 3">
            <img src="../../icons/status_ing_grey.svg" alt="status icon" v-if="step < 3">
            <img src="../../icons/status_ok.svg" alt="status icon" v-if="step > 3">
            <div :class="{ 'step-tip': true, 'active': step >= 3 ? true : false }" style="line-height: 28px;">Offer your
               data for sale
               on the plaza. Prior listing,
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
         <div class="platforms"><p>comming soon...</p></div>
         <div class="platforms"><p>comming soon...</p></div>

         <div class="button-wrapper" style="height:150px">
            <el-button class="button" @click="next" style="right: 0;">Continue</el-button>
            <div class="tips">
               If your genetic testing platform is not listed, please to provide feedback, and we will address it
               promptly.
            </div>
         </div>
      </div>


      <div class="wrapper-right" style="display: block; padding: 200px 200px;" v-if="step === 2">
         <p>Report *</p>
         <div class="select" @click="getProfileIds">
            <div class="add" v-if="!selectedProfile">+</div>
            <div v-if="!selectedProfile">Select the report as collection</div>
            <div v-if="selectedProfile">{{ selectedProfile }}</div>
         </div>
         <p>Name *</p>
         <el-input class="name-input" v-model="name" placeholder="Please enter the name of the collection"></el-input>
         <!-- <p class="introduction" style="margin-bottom: 0;">Since there are several analysis files in the genetic
            report, your
            GNFT will be
            automatically given a unique name depending on the Collection name you provide.</p> -->
         <p class="introduction"><span class="click-here">Click here</span> to view an example.</p>
         <p>Description</p>
         <input type="text" class="Description-input" placeholder="Please enter a description of the collection"
            v-model="description"></input>
         <div class="button-wrapper" style="margin-top: 50px;">
            <el-button class="button" @click="back" style="right: 120px;">back</el-button>
            <el-button class="button" @click="createData" style="right: 0px;">Create</el-button>

         </div>

      </div>

      <div class="wrapper-right" style="display: block; padding-top: 100px;" v-if="step === 3">
         <h2>Select your collection</h2>
         <div class="button-wrapper" style="margin-top: 50px;">
            <el-button class="button" @click="back">back</el-button>
            <el-button class="button" @click="next">Create</el-button>
         </div>
      </div>

   </div>
   <el-drawer v-model="isVisible" title="Please select a profile" :direction="ltr" :before-close="handleClose">
      <el-table :data="profiles" @selection-change="handleSelectionChange">
         <el-table-column>
            <template #default="{ row }">
               <el-radio v-model="selectedProfile" :label="row" />
            </template>
         </el-table-column>
      </el-table>
      <p class="selected-profile">Selected Profile: {{ selectedProfile }}</p>
      <template #footer>
         <div style="flex: auto">
            <el-button @click="confirmClick">Confirm</el-button>
         </div>
      </template>
   </el-drawer>
</template>


<script lang="js" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Bubbles from '../components/bubbles.vue';
import Api from '../../axios/aixos';
import { useWalletStore } from '@/stores/account';
import { ElLoading } from 'element-plus';

const walletStore = useWalletStore();
const router = useRouter();

const step = ref(-1); // 0:进行手机验证

const showIndexPage = ref(true);


function turnToSteps() {
   showIndexPage.value = false;
   if (showIndexPage.value === false && walletStore.insititution) {
      step.value += 1;
   }
   step.value += 1;
}

function toVerify() {
   router.push('/verify');
}

function next() {
   step.value++;
}

function back() {
   step.value--;
}

function redirectToOAuth() {
   window.location.href = import.meta.env.VITE_APP_BASE_URL + '/user/oauth2Wegene';
}

let profiles = ref([]);
let isVisible = ref(false);
let profileName = ref('');
const getProfileIds = async () => {
   isVisible.value = true;
   try {
      const res = await Api.get('/studio/getProfileIds');
      var loading = ElLoading.service({
         lock: true,
         text: 'Loading...',
         background: 'rgba(0, 0, 0, 0.7)',
         customClass: 'loading',
      })
      if (res.data.code === 200) {
         loading.close();
         profiles.value = res.data.data.profile_ids
         profileName.value = res.data.data.profile_name
         console.log(profiles);
      }
   } catch (error) {
      loading.close();
      alert("Network error. Please try again later");
      console.error(error);
   }
}

/* 表单提交 */
let selectedProfile = ref(null);
let name = ref('');
let description = ref('');
let tags = ref('');

const confirmClick = () => {
   isVisible.value = false;
}

const createData = async () => {
   try {
      const res = await Api.post('/studio/createFromThirdParty', {
         profile_id: selectedProfile.value,
         name: name.value,
         description: description.value,
         tags: tags.value
      });
      if (res.data.code === 200) {
         console.log('Data created successfully:', res.data);
         // next();
         console.log('Data created successfully:', res.data);
      } else {
         console.error('Error creating data:', res.data);
      }
   } catch (error) {
      console.error('Error creating data:', error);
   }
};

const handleSelectionChange = (val) => {
}
</script>

<style lang="scss" scoped>
.wrapper {
   height: 95vh;
   width: 100vw !important;
   overflow: hidden;
   min-width: 1200px !important;
   animation: fadeIn 0.2s ease-in-out 0s forwards;

   @keyframes fadeIn {
      0% {
         opacity: 0;
      }

      100% {
         opacity: 1;
      }
   }
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
      width: 280px;
      height: 70px;
      line-height: 70px;
      background-color: #169608;
      color: #fff;
      font-size: 26px;
      border-radius: 50px;
      text-align: center;

      &:hover {
         box-shadow: 0px 0px 20px 1px #e7e4e4 !important;
         cursor: pointer;
      }
   }
}


h1 {
   margin: 10px;
   color: #333;
}

.wrapper-left {
   overflow: visible;
   position: absolute;
   left: 0;
   width: 50vw;
   min-width: 600px;
   padding: 150px 80px 100px 160px;
   height: 95vh;
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
   overflow: auto;
   right: 0;
   width: 50%;
   min-width: 600px !important;
   padding: 100px 160px;
   height: 95vh !important;
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
      position: relative;
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
         margin: 25px 40px;
         height: 50px;
      }

      p{
         width: 100%;
         text-align: center;
         color: #cdcdcd;
         font-size: 24px;
      }
   }

   .button-wrapper {
      display: flex;
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 60%;
      position: relative;
      // height: 100px;



      .tips {
         position: absolute;
         bottom: 0;
      }
   }

}

:deep(.button) {
   position: absolute;
   width: 100px;
   height: 42px;
   background-color: #67C23A;
   color: #fff;
   font-size: 18px;
   border-radius: 10px;
   border: none !important;

   &:hover {
      color: #67C23A;
      background-color: #fff;
      border: #67C23A 1px solid !important;
      cursor: pointer;
   }
}


.select {
   padding: 10px;
   display: flex;
   align-items: center;
   margin: 10px 0;
   width: 400px;
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
   width: 400px;
   height: 50px;
   // border-radius: 15px;
}

:deep(.supply-input) {
   margin: 10px 0;
   width: 200px;
   height: 50px;
   // border-radius: 15px;
}

.Description-input {
   margin: 10px 0;
   width: 400px;
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
</style>

// 电话验证部分
<style lang="scss" scoped>
.wrapper-right {
   display: flex;
   position: absolute;
   right: 0;
   width: 50%;
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
</style>