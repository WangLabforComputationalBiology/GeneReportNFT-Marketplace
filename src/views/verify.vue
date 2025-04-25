<template>
    <div class="wrapper">
        <el-form :model="form" label-width="auto" style="max-width: 600px">
            <el-form-item label="Phone Number:">

                <el-input v-model="form.number" type="tel" oninput="value=value.replace(/[^0-9]/g,'')" maxlength="11">
                    <template #prepend>
                        <el-select v-model="form.selectVal" placeholder="Select" style="width: 70px">
                            <el-option label="+86" value="1" />
                            <el-option label="+88" value="2" />
                            <el-option label="+89" value="3" />
                        </el-select>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item label="Verification code:">
                <el-input v-model="form.Vcode" />
            </el-form-item>
            <!-- <el-form-item label="Activity zone">
                <el-select v-model="form.region" placeholder="please select your zone">
                    <el-option label="Zone one" value="shanghai" />
                    <el-option label="Zone two" value="beijing" />
                </el-select>
            </el-form-item> -->
            <el-form-item>
                <el-button type="primary" @click="onSubmit">Create</el-button>
                <el-button>Cancel</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
    name: 'Verify',
    data() {
        return {
            form: {
                number: '',
                Vcode: '',
                selectVal: '+86',
            }
        };
    },
    methods: {

    }
};
</script>

<style lang="scss" scoped>
.wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #fff;
}
</style>

<style lang="scss" scoped>
:deep(.el-form-item__label) {
    font-size: 16px;
    color: #67C23A;
}

:deep(.el-input__wrapper) {

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
    border: none;
    background-color: #fff;

}

// 下拉框选择器
:deep(.el-tooltip__trigger) {
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
:deep(.el-select-dropdown) {
    color: #67C23A !important;
}

</style>