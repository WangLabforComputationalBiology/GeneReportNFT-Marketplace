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
                <el-button @click="verify">Verify</el-button>
            </el-form-item>
        </el-form>
        <img :src="'https://flagcdn.com/24x18/' + form.country + '.png'" :alt="form.country + '国旗'" class="mr-2" />
    </div>
</template>

<script>
import { reactive, computed } from 'vue';
import Api from '../axios/aixos'
// 国家代码映射
const countryCodeMap = {
    cn: '+86',
    us: '+1',
    de: '+49',
};

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
    computed: {
        phoneNumber() {
            if (!this.form.country || !this.form.number) {
                return '';
            }
            return `${countryCodeMap[this.form.country]}${this.form.number}`;
        },
    }
    ,
    methods: {
        async getCode() {
            try {
                // 验证手机号格式（假设需要）
                if (!this.phoneNumber) {
                    this.$message.error("Please enter a valid phone number.");
                    return;
                }

                const response = await Api.post("/user/send_sms", {
                    phone: this.phoneNumber,
                });

                const { data } = response;

                // 确保响应数据结构正确
                if (!data || typeof data.code === "undefined") {
                    throw new Error("Invalid response format from server.");
                }

                // 处理不同的状态码
                if (data.code === 200) {
                    this.$message.success(this.$t("sms.sent_success")); // 使用国际化
                    return data; // 显式返回数据
                } else if (data.code === 429) {
                    this.$message.error
                    this.$message.error(this.$t("sms.rate_limit_exceeded")); // 频率限制
                    return;
                } else {
                    this.$message.error(data.message || this.$t("sms.send_failed")); // 使用后端返回的错误信息
                    return;
                }
            } catch (error) {
                console.error("Error sending verification code:", error);
                // 提取后端返回的错误信息（假设使用 Axios）
                const errorMessage =
                    error.response?.data?.message ||
                    error.message ||
                    this.$t("sms.error");
                this.$message.error(errorMessage);
                throw error; // 抛出错误以便调用方处理
            }
        },

        async verify() {
            try {
                // 验证输入
                if (!this.form.number || !/^\d{10,}$/.test(this.form.number)) {
                    this.$message.error(this.$t("verify.invalid_phone")); // 使用国际化
                    return;
                }
                if (!this.form.code || !/^\d{4,6}$/.test(this.form.code)) {
                    this.$message.error(this.$t("verify.invalid_code")); // 假设验证码为4-6位数字
                    return;
                }

                const response = await Api.post(
                    "/user/verify_sms",
                    {
                        phone: this.form.number,
                        code: this.form.code,
                    },
                    { timeout: 10000 } // 设置请求超时
                );

                const { data } = response;

                // 确保响应数据结构正确
                if (!data || typeof data.code === "undefined") {
                    throw new Error("Invalid response format from server.");
                }

                // 处理不同的状态码
                if (data.code === 200) {
                    this.$message.success(this.$t("verify.success"));
                    return data; // 显式返回数据
                } else if (data.code === 400) {
                    this.$message.error(this.$t("verify.invalid_code_error")); // 验证码错误
                    return;
                } else if (data.code === 429) {
                    this.$message.error(this.$t("verify.rate_limit_exceeded")); // 频率限制
                    return;
                } else {
                    this.$message.error(data.message || this.$t("verify.failed")); // 使用后端错误信息
                    return;
                }
            } catch (error) {
                console.error("Error verifying code:", error);
                // 提取后端返回的错误信息（假设使用 Axios）
                const errorMessage =
                    error.response?.data?.message ||
                    error.message ||
                    this.$t("verify.error");
                this.$message.error(errorMessage);
                throw error; // 抛出错误以便调用方处理
            }
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