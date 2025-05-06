<template>

   <body>
      <div class="wrapper">
         <div class="wrapper-left" v-if="this.step === null">
            <h1 class="title">Create</h1>
            <div class="create-button" @click="turnToSteps">
               <div class="button-wrapper">
                  <div class="button-title">
                     <h2>Collectin or item</h2>
                     <p>
                        create a new NFT collection or add an NFT to existing one.
                     </p>
                  </div>
                  <el-icon style="font-size: 30px; right: 20px; position: absolute;transform: translateY(-50%);">
                     <Right />
                  </el-icon>
               </div>
            </div>
         </div>


         <div class="wrapper-right" v-if="this.step === null">

         </div>


         <div class="wrapper-left" v-if="this.step === 0">
            <h1 style="margin: 10px; color:#909399">Start your GNFT journey from here</h1>
            <div class="inside-step"><img src="../../icons/status_ing.svg" alt="status icon">
               <div class="step-tip" style="line-height: 80px; color:#67c23a">Verify your region using phone number.</div>

            </div>
            <div class="line" style="background-color: #67c23a;"></div>
            <div class="inside-step"><img src="../../icons/status_ing_grey.svg" alt="status icon">
               <div class="step-tip">Allow us to access your genetic reports. We ensure that no additonal personal data
                  will be stored by your platform.</div>

            </div>
            <div class="line"></div>
            <div class="inside-step"><img src="../../icons/status_ing_grey.svg" alt="status icon">
               <div class="step-tip">Create your unique GNFT. Once your item is minted you will not be able to change
                  any of its information.</div>

            </div>
            <div class="line"></div>
            <div class="inside-step"><img src="../../icons/status_ing_grey.svg" alt="status icon">
               <div class="step-tip" style="line-height: 28px;">Offer your GNFT for sale on the market. Prior listing,
                  you may delete your GNFT; However, once listed, you will not be able to delete your GNFT。</div>

            </div>
         </div>

         <div class="wrapper-right" v-if="this.step === 0">
            <el-form :model="form" label-width="auto" style="max-width: 420px">
               <el-form-item label="Phone number:">
                  <el-input v-model="form.number" type="tel" oninput="value=value.replace(/[^0-9]/g,'')" maxlength="11">
                     <template #prepend>
                        <el-select v-model="form.country" placeholder="Select" style="width: 64px">
                           <el-option label="+86" value="cn" />
                           <el-option label="+1" value="us" />
                           <el-option label="+7" value="de" />
                           <el-option label="+44" value="gb" />
                           <el-option label="+33" value="fr" />
                        </el-select>
                     </template>
                     <template #append>
                        <img :src="'https://flagcdn.com/32x24/' + form.country + '.png'" :alt="form.country + '国旗'"
                           class="flags" />
                     </template>
                  </el-input>
               </el-form-item>

               <!-- <gocaptcha-rotate :config="{}" :data="{}" :events="{}" ref="domRef" /> -->

               <el-form-item label="Verification code:">
                  <el-input v-model="form.code" type="code" maxlength="6">
                     <template #prepend>
                        <img src="../../icons/verification_code.svg" alt=""
                           style="width: 20px;margin: 5px 2px 0 2px;" />
                     </template>
                     <template #append>
                        <el-button type="primary" @click="getCode">Send</el-button>
                     </template>
                  </el-input>
               </el-form-item>
               <el-form-item>

                  <el-button class="custom-botton" @click="verify">Verify</el-button>
               </el-form-item>
            </el-form>
         </div>
      </div>
   </body>
</template>


<script>
import Api from '../../axios/aixos'
import { ref } from 'vue'
// 国家代码映射
const countryCodeMap = {
   cn: '+86',
   us: '+1',
   de: '+49',
   fr: '+33',
   gb: '+44',
};

export default {
   name: "create",
   data() {
      return {
         step: null,//0:进行手机验证
         form: {
            number: '',
            code: '',
            selectVal: '+86',
            country: 'cn',
         },

      }
   },
   mounted() {

   },
   methods: {
      turnToSteps() {
         this.step = 0;
      }
   }
}
</script>

<style lang="scss" scoped>
.active{
   p{
      color: #67C23A;
   }
   img{
      color: #67C23A;
   }
}

.wrapper {
   display: flex;
   position: relative;

   .wrapper-left {
      position: absolute;
      left: 0;
      width: 50%;
      padding: 150px 80px 100px 160px;
      height: calc(100vh - 60px);

      .title {
         font-size: 70px;
         color: #67C23A;
         margin-bottom: 20px;
      }

      .create-button {
         display: flex;
         position: relative;
         align-items: center;
         width: 100%;
         height: 130px;
         box-shadow: 0px 0px 5px 0px #e0e0e0 !important;
         border-radius: 10px;

         &:hover {
            box-shadow: 0px 0px 10px 1px #e0e0e0 !important;
            cursor: pointer;
         }

         .button-title {
            font-size: 16px;
            left: 30px;
            transform: translateY(-60%);
            position: absolute;
         }
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
      }

      .line {
         margin-left: 28px;
         width: 3px;
         border-radius: 5px;
         height: 30px;
         background-color: #A8ABB2;
      }
   }
}

.wrapper-right {
   position: absolute;
   right: 0;
   width: 50%;
   height: calc(100vh - 60px);
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

   &:focus-within {
      box-shadow: 0px 0px 0px 1px inset #67C23A !important;
   }
}

:deep(.el-input__inner) {
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

:deep(.custom-botton) {
   width: 100%;
   height: 42px;
   background-color: #67C23A;
   color: #fff;
   font-size: 16px;
   border-radius: 5px;
   border: none !important;
   box-shadow: 0px 0px 5px 0px #e0e0e0 !important;

   &:hover {
      color: #67C23A;
      background-color: #fff;
      box-shadow: 0px 0px 10px 1px #e0e0e0 !important;
      cursor: pointer;
   }
}
</style>