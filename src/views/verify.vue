<template>
    <div class="wrapper">
        <div class="tittle"><span style="color: #169608;">Institutional</span> Accreditation</div>
        <div class="input-wrapper">
            <el-input v-model="inputValue"></el-input>
            <el-button>Search</el-button>
        </div>

    </div>
</template>

<script setup>
import { ref } from 'vue';

const inputValue = ref('');
</script>

<style lang="scss" scoped>
.wrapper {
    width: 30vw;
    min-width: 500px;
    height: 95vh;
    margin: auto;
    display: flex;
    position: relative;
    text-align: center;
    justify-content: center;
    // flex-direction: column;

}

.tittle {
    position: absolute;
    top: 20%;
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
</style>

<style lang="scss" scoped>
:deep(.el-input__wrapper) {
    padding: none !important;
    height: 50px;
    width: 82%;
    font-size: 20px;
    border-radius: 10px;
    border: 1px solid #169608;
    position: absolute;
    top: 28%;
    box-shadow: none;


}

:deep(.is-focus) {
    box-shadow: 0 0 0 3px #eeeeee inset;
    ;
}

:deep(.el-input__inner) {
    height: 50px;
    font-size: 20px;
    border: none;
    outline: none;
    padding-left: 10px;
}

:deep(.el-button) {
    position: absolute;
    top: 25%;
    right: 0%;
    width: 15%;
    height: 50px;
    min-width: 50px;
    font-size: 20px;
    border-radius: 10px;
    background-color: #169608;
    color: #fff;
    border: none;
}
</style>