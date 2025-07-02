<template>
    <div class="wrapper">
        <div class="banner">
            <span class="tip">
                <h1 style="margin-bottom: 30px;"><span style="color: #169608;">Wegene</span> Connected</h1>
            </span>
        </div>

        <div class="card-body">
            <div v-if="profiles.length > 0">
                <p style="color: #E6A23C;">Please select a Profile:</p>
                <el-table :data="profiles" @selection-change="handleSelectionChange">
                    <el-table-column>
                        <template #default="{ row }">
                            <el-radio v-model="selectedProfile" :label="row.id" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <div v-else>
                <p>No Profiles.Please request again.</p>
            </div>
            <p class="selected-profile">Selected Profile: {{ selectedProfile }}</p>
        </div>

        <div class="btn-wrapper">
            <el-button @click="back">Back</el-button>
            <el-button @click="authorizeProfile" :disable="selectedProfile == null">Confirm</el-button>
        </div>
    </div>
</template>

<script>
import { useWalletStore } from '@/stores/account';
const walletStore = useWalletStore();
export default {
    props: {
        content: {
            type: String,
            default: '内容'
        }
    },
    data() {
        return {
            code: this.$route.params.lastSegment || 'default title',
            profiles: [],
            selectedProfile: null,
            showAlert: false // 添加状态变量
        };
    },
    watch: {
        '$route.params.lastSegment': function (newSegment) {
            this.code = newSegment || 'default title';
            this.fetchData(); // 重新获取数据
        }
    },
    mounted() {
        this.fetchData(); // 在组件挂载时发起请求
    },
    methods: {
        back() {
            this.$router.push('/publish')
        },
        async fetchData() {
            try {
                const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/getProfileIds?code=${this.code}`, {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': walletStore.access_token
                    },
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                this.profiles = data.profiles || [];
                // 更新 content 或其他数据
                this.content = data.content || 'default content'; // 假设服务器返回的数据中有一个 content 字段
            } catch (error) {
                console.error('Error fetching data:', error);
                this.content = 'Error fetching data, please try again later'; // 更新 content 以显示错误信息
            }
        },
        async authorizeProfile() {
            // 在这里处理授权逻辑
            if (this.selectedProfile == null) {
                this.$message.error('Please select a profile');
                return
            }
            try {
                const response = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/user/saveProfile`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': walletStore.access_token
                    },
                    body: JSON.stringify({
                        code: this.code,
                        profileId: this.selectedProfile
                    })
                });
                if (response.ok) { // 检查响应状态码是否为 200
                    this.showAlert = true; // 显示 alert
                    this.$message.success('Profile authorized successfully');
                    this.$router.push('/publish?authorized=true')
                    this.$emit('profileAuthorized', true); // 触发父组件的自定义事件
                } else {
                    console.error('Failed to authorize profile:', response.statusText);
                }
            } catch (error) {
                console.error('Error authorizing profile:', error);
            }
        }
    }
}
</script>

<style lang="scss" scoped>
p {
    font-size: 20px;
}

.wrapper {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 40%;
    min-width: 410px;
    height: 50vh;
}

.user {
    font-size: 26px;
    width: 100%;
    color: #99a9bf;
    text-align: center;
    margin-bottom: 30px;
}

.selected-profile {
    color: #99a9bf;
    height: 50px;
    line-height: 50px;
}

.banner {
    display: flex;

    .banner-title {
        font-size: 70px;
        color: #67C23A;
    }

    .tip {
        width: 100%;
        font-size: 30px;
        color: #333;
        text-align: center;
    }
}

:deep(.el-button) {
    width: 80px;
    height: 40px;
    border-radius: 15px;
    font-size: 18px;

    &:hover {
        color: #67C23A;
        background-color: #fff;
        border: #67c23ab6 1px solid;
    }
}

:deep(.el-radio__label) {
    font-size: 18px;
}

.btn-wrapper {
    position: absolute;
    right: 0%;
}
</style>