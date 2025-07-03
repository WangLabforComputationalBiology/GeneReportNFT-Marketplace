<template>
    <div class="wrapper">
        <div class="banner">
            <span class="tip">
                <h1 style="margin-bottom: 30px;"><span style="color: #169608;">Wegene</span> Connected</h1>
                <h4>User:&nbsp;{{ userEmail }}</h4>
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

<script setup>
import { ref, watch, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Api from '@/axios/aixos';
import { ElMessage } from 'element-plus';

const route = useRoute();
const router = useRouter();

const userEmail = ref('');
const code = ref(route.params.uuid || 'default title');
const profiles = ref([]);
const selectedProfile = ref(null);
const showAlert = ref(false);

function back() {
    router.push('/publish');
}

/**请求profiles数据列表 */
async function fetchData() {
    try {
        const res = await Api.get(`/user/getProfileIds?code=${code.value}`);
        profiles.value = res.data.profiles || [];
        userEmail.value = res.data.email || '';
    } catch (error) {
        console.error('Error fetching data:', error);
        ElMessage.error('Error fetching data, please try again later');
    }
}

/**上传选定profile */
async function authorizeProfile() {
    if (selectedProfile.value == null) {
        ElMessage.error('Please select a profile');
        return;
    }
    try {
        const res = await Api.post(`/user/saveProfile`, {
            code: code.value,
            profileId: selectedProfile.value
        });
        if (res.data.code === 200) {
            showAlert.value = true;
            ElMessage.success('Profile authorized successfully');
            router.push(`/publish?authorized=true`);
        }
    } catch (error) {
        console.error('Error authorizing profile:', error);
    }
}

watch(
    () => route.params.uuid,
    (newSegment) => {
        code.value = newSegment || 'default title';
        fetchData();
    }
);

onMounted(() => {
    fetchData();
});
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
        margin-bottom: 30px;
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