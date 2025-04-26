<template>
    <div class="wrapper">
        <el-form :model="form" label-width="auto" style="max-width: 600px">
            <el-form-item label="Phone Number:">

                <el-input v-model="form.number" type="tel" oninput="value=value.replace(/[^0-9]/g,'')" maxlength="11">
                    <template #prepend>
                        <el-select v-model="form.country" placeholder="Select" style="width: 70px">
                            <el-option label="+86" value="cn" />
                            <el-option label="+1" value="us" />
                            <el-option label="+49" value="de" />
                        </el-select>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item label="Verification code:">
                <el-input v-model="form.code" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="getCode">Send Verification Code</el-button>
                <el-button>Verify</el-button>
            </el-form-item>
        </el-form>
        <img :src="'https://flagcdn.com/24x18/' + form.country + '.png'" :alt="form.country + '国旗'" class="mr-2" />
    </div>
</template>

<script>

export default {
    name: 'Verify',
    data() {
        return {
            form: {
                number: '',
                code: '',
                selectVal: '+86',
                country: 'cn',
            }
        };
    },
    methods: {
        async getCode() {
            await Api.post("/user/send_sms", {
                phone: this.form.number,
            }).then((res) => {
                console.log(res.data);
                if (res.data.code == 200) {
                    this.$message.success("Verification code sent successfully!");
                } else {
                    this.$message.error("Failed to send verification code.");
                }
            }).catch((error) => {
                console.error(error);
                this.$message.error("An error occurred while sending the verification code.");
            });
        },

        async verify() {
            await Api.post("/user/verify_sms", {
                phone: this.form.number,
                code: this.form.code,
            }).then((res) => {
                console.log(res.data);
                if (res.data.code == 200) {
                    this.$message.success("Verification successful!");
                } else {
                    this.$message.error("Verification failed.");
                }
            }).catch((error) => {
                console.error(error);
                this.$message.error("An error occurred during verification.");
            });
        }
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