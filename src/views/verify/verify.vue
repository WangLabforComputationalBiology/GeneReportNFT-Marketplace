<template>
    <Bubbles />
    <div class="wrapper">
        <div class="tittle"><span style="color: #169608;">Institutional</span> Accreditation</div>
        <div class="input-wrapper">
            <el-autocomplete v-model="institution" :fetch-suggestions="querySearch" clearable class="inline-input"
                placeholder="Please select your institution" @select="handleSelect" />
        </div>
        <el-input class="email-input" v-if="step > 0" v-model="emailFront" placeholder="Please enter email">
            <template #append>{{ emailEnd }}</template>
        </el-input>
        <Slidecheck v-if="step > 0" class="slidecheck" @ready="handleReady" />
        <el-input class="code" v-if="Ready" placeholder="Please enter verification code" type="text" v-model="code">
            <template #prepend>
            </template>
            <template #append>
                <div class="send-btn" @click="resEmail" :disabled="isSent && countdown > 0">{{ sendBTN }}</div>
            </template>
        </el-input>
        <el-button v-if="Ready" class="verify" @click="verify">Verify</el-button>
    </div>

</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import loadUniversities from '@/assets/universities.json';
import Bubbles from '@/views/components/bubbles.vue';
import Slidecheck from '@/views/verify/slidecheck.vue';
import Api from '@/axios/aixos';

/* 大学列表 && 检索 */
const institution = ref('')
const universities = ref([])
const querySearch = (queryString, cb) => {
    const results = queryString
        ? universities.value.filter(createFilter(queryString))
        : universities.value
    cb(results)
}
const createFilter = (queryString) => {
    return (restaurant) => {
        return (
            restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
        )
    }
}

const loadAll = () => {
    return loadUniversities.map(item => {
        return {
            value: item.value,
            email: item.email,
        }
    })
}

let emailEnd = ref('');

// 选择大学后，设置emailEnd为对应的email
let step = ref(0);
const handleSelect = (item) => {
    emailEnd.value = item.email;;
    step.value = 1;//显示人机验证
}

/* 人机验证
 * 通过滑动验证来确认用户是人类
 * 如果验证通过，设置Ready为true
 * 后续显示发送验证码页面
*/
const Ready = ref(false);
const handleReady = (isReady) => {
    if (isReady === true) {
        Ready.value = true;
    }
}

onMounted(() => {
    universities.value = loadAll()
})


//邮箱整合
const emailFront = ref('');
const fullEmail = computed(() => emailFront.value + emailEnd.value);

// 是否可验证
let isSent = ref(false);
/* 发送邮箱验证请求 */
const resEmail = async () => {
    if (!emailFront.value) {
        // 可以在这里添加提示：请输入完整的邮箱信息
        alert('Please enter your email address.');
        return;
    }
    try {
        const emailResponse = await Api.post("/user/send_email", {
            email: fullEmail.value,
            institution: institution.value,
        });
        if (emailResponse.data.code === 200) {
            // 成功提示
            alert('Email sent successfully!');
            isSent = true;
        } else {
            // 可败提示
            alert('Failed to send email: ' + (emailResponse.data.message || 'Unknown error'));
        }
    } catch (error) {
        if (error.response.status === 401) {
            alert('Token expired.Please log in again');
        }
        if (error.response.status === 404) {
            alert('Network error.');
        }
        console.error('Error:', error.response.status);
    }
}

/*
 * 发送验证码按钮状态
 * 如果按钮状态为true且倒计时大于0，则显示倒计时
 * 否则显示“发送”按钮
 */
let countdown = ref(60);
let timer = null;
const sendBTN = computed(() => {
    if (isSent.value === true && countdown.value > 0) {
        return `${countdown.value}s`;
    }
    return 'Send';
});
/* 倒计时函数 */
watch(isSent, (val) => {
    if (val) {
        countdown.value = 60;
        timer = setInterval(() => {
            if (countdown.value > 0) {
                countdown.value--;
            } else {
                isSent.value = false;
                clearInterval(timer);
                timer = null;
            }
        }, 1000);
    } else {
        if (timer) {
            clearInterval(timer);
            timer = null;
        }
    }
});


/* 验证码输入 */
let code = ref('');
const verify = async () => {
    if (emailFront.value === false || isSent.value === false) {
        alert('Please enter your email and select your institution.');
        return;
    }
    if (code.value === false) {
        alert('Please enter the verification code.');
        return;
    }
    try {
        const response = await Api.post("/user/verify_email", {
            email: fullEmail.value,
            code: code.value
        });
        if (response.data.code === 200) {
            alert('Verification successful!');
            setTimeout(() => {
                window.location.href = '/account' // 跳转到账户页面
            }, 1500);
        }
    } catch (error) {
        if (error.response.status === 401) {
            alert('Token expired, please log in again.');
            setTimeout(() => {
                window.location.href = '/login' // 跳转到登录页面
            }, 1500);
        } else if (error.response.status === 403) {
            alert('Verification code incorrect, please try again.');
        } else if (error.response.status === 404) {
            alert('Network error, please try again later.');
        }
        else {
            alert('Verification failed, please try again.');
        }
        console.error('Error verifying email:', AxiosError.response.status);
    }
}
</script>

<style lang="scss" scoped>
.wrapper {
    width: 20vw;
    min-width: 450px;
    height: 95vh;
    margin: auto;
    position: relative;
    flex-direction: column-reverse;
}

.tittle {
    position: absolute;
    width: 100%;
    top: 20%;
    left: 50%;
    transform: translateX(-50%);
    text-align: center;
    font-size: 36px;
}

.input-wrapper {
    position: absolute;
    top: 28%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.slidecheck {
    position: absolute;
    top: 43%;
    width: 100%;
    left: 50%;
    transform: translateX(-50%);
}

.send-btn {
    width: 80px;
    color: #fff;

    font-size: 18px;
    text-align: center;
    cursor: pointer;
    height: 28px;
    line-height: 28px;
    background-color: #169608;
    border: 1px solid #169608;
    border-radius: 10px;
}
</style>

<style lang="scss" scoped>
:deep(.el-input__wrapper) {
    height: 40px;
    padding: none !important;
    width: 100%;
    font-size: 20px;
    border-radius: 10px;
    border: 1px solid #ccc;
    position: absolute;
    top: 28%;
    box-shadow: none;
}

:deep(.inline-input) {
    height: 65px;
}

:deep(.el-autocomplete-suggestion) {
    margin: 0 !important;
    border: none !important;
    box-shadow: none !important;
    border-radius: 20px !important;
}

:deep(.email-input) {
    position: absolute;
    top: 36%;
    width: 100%;

}

:deep(.el-input-group__append) {
    position: absolute;
    right: 0%;
    height: 40px;
    background-color: #fff;
    border: 1px solid #ccc;
    color: #169608;
    box-shadow: none;
    border-top-right-radius: 10px;
    border-bottom-right-radius: 10px;
    font-size: 18px;
    font-weight: 700;
    border-left: none;
}



:deep(.el-input__inner) {
    height: 40px;
    font-size: 20px;
    border: none;
    outline: none;
    padding-left: 10px;
}

:deep(.el-button) {
    position: absolute;
    top: 38%;
    width: 100%;
    height: 50px;
    min-width: 50px;
    font-size: 20px;
    border-radius: 10px;
    background-color: #169608;
    color: #fff;
    border: none;
}

:deep(.code) {
    position: absolute;
    top: 50%;
    width: 100%;
}

:deep(.verify) {
    position: absolute;
    top: 56%;
    width: 100%;
    height: 40px;
    font-size: 20px;
    border-radius: 10px;
    background-color: #169608;
    color: #fff;
    border: none;

}
</style>