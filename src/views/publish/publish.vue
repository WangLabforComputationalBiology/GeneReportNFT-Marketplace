<template>

   <div class="wrapper">
      <Bubbles v-if="this.step === -1" />
      <div class="wrapper-center" v-if="this.step === -1">
         <h1 class="title">Data <span style="color: #333;">Publish</span></h1>
         <span class="tip1">Ready to upload your data and publish it.</span>
         <div class="start-btn" @click="turnToSteps">Start</div>
      </div>


      <!-- 短信验证阶段 0 -->
      <div class="wrapper-left" v-if="this.step === 0">
         <h1>Start your GNFT journey from here</h1>
         <div class="inside-step"><img src="../../icons/status_ing.svg" alt="status icon">
            <div class="step-tip" style="line-height: 80px; color:#67c23a;">Verify your region using email.
            </div>

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
               you may delete your GNFT; However, once listed, you will not be able to delete your GNFT.</div>

         </div>
      </div>

      <div class="wrapper-right" v-if="this.step === 0">
         <el-form :model="form" label-width="auto" style="max-width: 480px">
            <el-form-item label="Your email:">
               <el-input v-model="form.number">
                  <template #append>
                     <el-button type="primary" @click="getCode">Send</el-button>
                  </template>
               </el-input>
            </el-form-item>

            <SliderCheck :successFun="handleSuccessFun" :errorFun="handleErrorFun" style="margin-bottom: 22px;" />


            <el-form-item label="Sent code:">
               <el-input v-model="form.code" type="code" maxlength="6">
                  <template #prepend>
                     <img src="../../icons/verification_code.svg" alt="" style="width: 20px;margin: 5px 2px 0 2px;" />
                  </template>
                  <!-- <template #append>
                        
                     </template> -->
               </el-input>
            </el-form-item>
            <el-form-item>

               <el-button class="custom-botton" @click="next">Verify to continue</el-button>
            </el-form-item>
            <el-form-item>
               <el-button class="custom-botton" @click="back">back</el-button>
            </el-form-item>
         </el-form>
      </div>

      <!-- 选择平台阶段 1 -->
      <div class="wrapper-left" v-if="this.step === 1">
         <h1 style="margin: 10px; color:#909399">Start your GNFT journey from here</h1>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip" style="line-height: 80px; color:#67c23a;">Verify your region using phone number.
            </div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ing.svg" alt="status icon">
            <div class="step-tip">Allow us to access your genetic reports. We ensure that no additonal personal data
               will be stored by your platform.</div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
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

      <div class="wrapper-right" style="display: block;" v-if="this.step === 1 && profiles.length <= 0">
         <div class="platforms" @click="redirectToOAuth">
            <img src="../../icons/wegene_logo.svg" alt="wegene">
            <el-icon style="position: absolute; right:30px; top: 50px; transform: translateY(-50%); font-size: 30px;">
               <Right />
            </el-icon>
         </div>
         <div class="platforms"></div>
         <div class="platforms"></div>

         <div class="button-wrapper" style="display: flex; position:relative;height:200px">
            <el-button class="botton" @click="back" style="right: 240px;">back</el-button>
            <el-button class="botton" @click="next">Continue</el-button>
            <div class="tips">
               If your genetic testing platform is not listed, please to provide feedback, and we will address it
               promptly.
            </div>
         </div>
         <div class="wrapper-right" style="display: block;" v-if="this.step === 1 && profiles.length > 0">
            <p style="color: #99a9bf;">Please select a Profile:</p>
            <el-table :data="profiles" style="width: 50%" @selection-change="handleSelectionChange">
               <!-- Radio button column -->
               <el-table-column width="500" label="#Id">
                  <template #default="{ row }">
                     <el-radio v-model="selectedProfile" :label="row.id" />
                  </template>
               </el-table-column>

            </el-table>
         </div>
      </div>

      <!-- 添加描述阶段 1 -->
      <div class="wrapper-left" v-if="this.step === 2">
         <h1>Start your GNFT journey from here</h1>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip" style="line-height: 80px; color:#67c23a">Verify your region using phone number.
            </div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip-ok">Allow us to access your genetic reports. We ensure that no
               additonal
               personal data
               will be stored by your platform.</div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ing.svg" alt="status icon">
            <div class="step-tip-ok">Create your unique GNFT. Once your item is minted you will not be able to change
               any of its information.</div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ing_grey.svg" alt="status icon">
            <div class="step-tip" style="line-height: 28px;">Offer your GNFT for sale on the market. Prior listing,
               you may
               delete your GNFT; However, once listed, you will not be able to delete your GNFT.
            </div>

         </div>
      </div>

      <div class="wrapper-right" style="display: block; padding-top: 100px;" v-if="this.step === 2">
         <p>Report *</p>
         <div class="select">
            <div class="add">+</div>
            <div>Select the report as collection</div>
         </div>
         <p>Name *</p>
         <el-input class="name-input"></el-input>
         <p class="introduction" style="margin-bottom: 0;">Since there are several analysis files in the genetic
            report, your
            GNFT will be
            automatically given a unique name depending on the Collection name you provide.</p>
         <p class="introduction"><span class="click-here">Click here</span> to view an example.</p>
         <p>supply *</p>
         <el-input class="supply-input"></el-input>
         <p class="introduction">The Supply refers to the quantity of each NFT within the Collection.</p>
         <p>Description</p>
         <input type="text" class="Description-input"
            placeholder="Please enter a description of the collection"></input>
         <p class="introduction">The description will be included in every GNFT in the Collection</p>
         <p>Trait</p>
         <p class="introduction">Traits describe attributes of your item. They appear as filters inside your
            collection page
            and are also listed out inside your item page.</p>
         <div class="add-trait">+ Add trait</div>
         <div class="button-wrapper" style="margin-top: 50px;">
            <el-button class="botton" @click="back" style="right: 240px;">back</el-button>
            <el-button class="botton" @click="next">Create</el-button>

         </div>

      </div>


      <!-- 完成阶段 3 -->
      <div class="wrapper-left" v-if="this.step === 3">
         <h1>Start your GNFT journey from here</h1>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip" style="line-height: 80px; color:#67c23a">Verify your region using phone number.
            </div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip" style="color:#67c23a">Allow us to access your genetic reports. We ensure that no
               additonal
               personal data
               will be stored by your platform.</div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ok.svg" alt="status icon">
            <div class="step-tip">Create your unique GNFT. Once your item is minted you will not be able to change
               any of its information.</div>

         </div>
         <div class="line" style="background-color: #67c23a;"></div>
         <div class="inside-step"><img src="../../icons/status_ing.svg" alt="status icon">
            <div class="step-tip" style="line-height: 28px;">Offer your GNFT for sale on the market. Prior listing,
               you may
               delete your GNFT; However, once listed, you will not be able to delete your GNFT.
            </div>

         </div>
      </div>

      <div class="wrapper-right" style="display: block; padding-top: 100px;" v-if="this.step === 3">
         <h2>Select your collection</h2>
         <div class="button-wrapper" style="margin-top: 50px;">
            <el-button class="botton" @click="back" style="right: 240px; b">back</el-button>
            <el-button class="botton" @click="next">Create</el-button>

         </div>
      </div>



   </div>
</template>


<script lang="js">
import SliderCheck from '../verify/slidecheck.vue';
import Bubbles from '../components/bubbles.vue';
// import Api from '../../axios/aixos';


export default {
   name: 'Create',
   components: {
      SliderCheck,
      Bubbles
   },
   data() {
      return {
         step: -1,//0:进行手机验证
         form: {
            number: '',
            code: '',
            selectVal: '+86',
            country: 'cn',
         },
         code: this.$route.params.lastSegment || '默认标题',
         profiles: [],
         selectedProfile: null,
         showAlert: false // 添加状态变量
      }
   },
   mounted() {
      console.log('Mounted title:', this.code); // 打印挂载时的 title 值
      console.log('Environment variable VITE_APP_BASE_URL:', import.meta.env.VITE_APP_BASE_URL); // 打印环境变量
      this.fetchData(); // 在组件挂载时发起请求
   },
   watch: {
      // Watch for changes in the route parameter
      '$route.params.lastSegment': function (newSegment) {
         this.code = newSegment || '默认标题';
         this.fetchData(); // 重新获取数据
      }
   },
   methods: {
      turnToSteps() {
         this.step = 0;
      },
      next() {
         this.step++;
      },
      back() {
         this.step--;
      },

      redirectToOAuth() {
         window.location.href = import.meta.env.VITE_APP_BASE_URL + '/user/oauth2Wegene';
      },
      getCode() {
         // Add logic to send verification code
         console.log('Sending verification code...');
      },

      // 处理授权
      async fetchData() {
         try {
            //const response = await fetch(`http://127.0.0.1:4523/m1/4576706-4225408-default/user/getProfileIds`);
            const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/getProfileIds?code=${this.code}`);
            if (!response.ok) {
               throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            this.profiles = data.profiles || [];
            console.log('Fetched data:', this.profiles);
            // 你可以在这里更新 content 或其他数据
            this.content = data.content || '默认内容'; // 假设服务器返回的数据中有一个 content 字段
         } catch (error) {
            console.error('Error fetching data:', error);
            this.content = '请求失败，请重试'; // 更新 content 以显示错误信息
         }
      },
      async authorizeProfile() {
         // 在这里处理授权逻辑
         try {
            const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/saveProfile`, {
               method: 'POST',
               headers: {
                  'Content-Type': 'application/json'
               },
               body: JSON.stringify({
                  code: this.code,
                  profileId: this.selectedProfile
               })
            });

            if (response.ok) { // 检查响应状态码是否为 200
               this.showAlert = true; // 显示 alert
               console.log('Profile authorized successfully');
            } else {
               console.error('Failed to authorize profile:', response.statusText);
            }
         } catch (error) {
            console.error('Error authorizing profile:', error);
         }
      },
      // 滑块验证成功回调
      handleSuccessFun() {
         this.login_model.status = true
      },
      // 滑块验证失败回调
      handleErrorFun() { },
   }
}
</script>

<style lang="scss" scoped>
.wrapper {
   height: 95vh;
   min-width: 1200px;
   overflow: hidden;
}

.active {
   p {
      color: #67C23A;
   }

   img {
      color: #67C23A;
   }
}

.wrapper {
   display: flex;
   position: relative;

}

.wrapper-center {
   height: 95vh;
   width: 100vw;
   display: flex;
   position: absolute;
   flex-direction: column;
   gap: 30px;
   justify-content: center;
   align-items: center;

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


.wrapper-left {
   overflow: visible;
   position: absolute;
   left: 0;
   width: 50%;
   padding: 150px 80px 100px 160px;
   height: 95vh;

   h1 {
      margin: 10px;
      color: #909399;
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
         color: #A8ABB2;
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
   padding: 100px 160px;
   height: 95vh;

   .platforms {
      position: relative;
      margin: 50px auto;
      width: 60%;
      height: 100px;
      border: #e0e0e0 1px solid;
      border-radius: 10px;
      box-shadow: 0px 0px 2px 0px #e0e0e0 !important;

      &:hover {
         box-shadow: 0px 0px 15px 1px #e0e0e0 !important;
         cursor: pointer;
      }

      img {
         margin: 25px 40px;
         height: 50px;
      }
   }

   .button-wrapper {
      display: flex;
      position: relative;
      height: 100px;

      :deep(.botton) {
         position: absolute;
         right: 132px;

         width: 100px;
         height: 42px;
         background-color: #67C23A;
         color: #fff;
         font-size: 16px;
         border-radius: 10px;
         border: none !important;
         text-align: center;
         box-shadow: 0px 0px 5px 0px #e0e0e0 !important;

         &:hover {
            color: #67C23A;
            background-color: #fff;
            box-shadow: 0px 0px 10px 1px #e0e0e0 !important;
            cursor: pointer;
         }
      }

      .tips {
         position: absolute;
         bottom: 0;
      }
   }

}




.select {
   padding: 10px;
   display: flex;
   align-items: center;
   margin: 10px 0;
   width: 400px;
   height: 60px;
   background-color: #F5F5F5;
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