<template>
    <Bubbles />
    <div class="wrapper">

        <div class="tittle"><span style="color: #169608;">Institutional</span> Accreditation</div>
        <div class="input-wrapper">
            <el-autocomplete v-model="state" :fetch-suggestions="querySearch" clearable class="inline-input"
                placeholder="Please select your institution" @select="handleSelect" />
        </div>
        <!-- <el-button type="primary" @click="handleContinue">Continue</el-button> -->
        <el-input class="email-input" v-if="step > 0" v-model="emailFront" placeholder="Please enter your email">
            <template #append>{{ emailEnd }}</template>
        </el-input>
        <Slidecheck v-if="step > 0" class="slidecheck" @ready="handleReady" />
        <el-input class="code" v-if="step > 1" placeholder="Please enter verification code" type="text">
            <template #prepend>
            </template>
            <template #append>
                <div class="send-btn">Send</div>
            </template>
        </el-input>
        <el-button v-if="step > 1" class="verify">Verify</el-button>
    </div>

</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import loadUniversities from '@/assets/universities.json';
import Bubbles from '@/views/components/bubbles.vue';
import Slidecheck from '@/views/verify/slidecheck.vue';


const state = ref('')
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
let step = ref(0);

const handleSelect = (item) => {
    emailEnd.value = item.email;;
    step.value += 1;
}

const handleReady = (isReady) => {
    if (isReady === true) {
        step.value += 1;
    }
}

onMounted(() => {
    universities.value = loadAll()
})

const emailFront = ref('');
//邮箱整合
const fullEmail = computed(() => emailFront.value + emailEnd.value);

</script>

<style lang="scss" scoped>
.wrapper {
    width: 20vw;
    min-width: 450px;
    height: 95vh;
    min-height: 750px;
    margin: auto;
    display: flex;
    position: relative;
    flex-direction: column;
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