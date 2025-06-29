<template>
    <div class="wrapper">
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
        </div>
        <div class="plaza-page" v-loading="loadingForPlaza" :element-loading-svg="svg">
            <div class="card" v-for="(item, index) in List" :key="index" @click="isVisible(item)">
                <div class="icon" />
                <div class="ifo">
                    <p>Name: <span class="ifo-item"> &nbsp; {{ item.name }}</span></p>
                    <p>Category: <span class="ifo-item"> &nbsp; {{ item.category }}</span></p>
                    <p>Description:<span class="ifo-item"> &nbsp; {{ item.description }}</span></p>
                    <p>Format:<span class="ifo-item"> &nbsp; {{ item.format }}</span></p>
                    <p>Date:<span class="ifo-item"> &nbsp; {{ item.created_at }}</span></p>
                </div>
            </div>
        </div>
    </div>

    <el-drawer v-model="drawerIsVisible" title="Detail" :direction="'rtl'" @closed="handleClosed" :size="'50%'"
        :show-close="false">
        <template #header>
            <div class="dt-page-top">
                <div class="icon" />
                <div class="ifo">
                    <p>Name: <span class="ifo-item"> &nbsp; {{ selectedData.name }}</span></p>
                    <p>Category: <span class="ifo-item"> &nbsp; {{ selectedData.category }}</span></p>
                    <p>Description:<span class="ifo-item"> &nbsp; {{ selectedData.description }}</span></p>
                    <p>Format:<span class="ifo-item"> &nbsp; {{ selectedData.format }}</span></p>
                    <p>Date:<span class="ifo-item"> &nbsp; {{ selectedData.created_at }}</span></p>
                </div>
            </div>
        </template>

        <div class="dt-page-bottom">
            <el-table v-loading="loadingForTable" :element-loading-svg="svg" max-height="65vh" :table-layout="'fixed'"
                :data="detailData">
                <el-table-column prop="ReportId" label="ReportId" />
                <el-table-column prop="description" label="Description" />
                <el-table-column prop="score" label="Score" />
                <el-table-column prop="rank" label="Rank" />
            </el-table>
        </div>

        <template #footer>
            <div style="flex: auto">
                <el-button type="primary" @click="confirmClick" class="obtain-btn">Obtain</el-button>
            </div>
        </template>
    </el-drawer>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Api from "@/axios/aixos";

/* 获取plaza卡片数据列表 */
const loadingForPlaza = ref(false);
const List = ref([]);
const getList = async () => {
    try {
        loadingForPlaza.value = true;
        const res = await Api.get('/plaza/getData');
        List.value = res.data.data.multi_metadata;
        List.value.forEach(item => {
            item.created_at = item.created_at.slice(0, 10);
        })
    } finally {
        loadingForPlaza.value = false;
    }
}

const selectedData = ref([]); 
const loadingForTable = ref(false);

/* 获取详细数据 */
const detailData = ref([]);
const getDetailData = async () => {
    try {
        loadingForTable.value = true;
        if (!selectedData.value?.data_hash) return; //问号表示不一定存在
        const res = await Api.get(`/metadata/${selectedData.value.data_hash}`);
        if (res.data.data) {
            detailData.value = res.data.data.details;
        }
    } catch (error) {
        console.error('Error fetching detail data:', error);
    } finally {
        loadingForTable.value = false;
    }
}

/* 详情抽屉 */
const drawerIsVisible = ref(false);
const isVisible = (item) => {
    selectedData.value = item;
    drawerIsVisible.value = true;
    getDetailData();
};

/* 释放表单内容 */
const handleClosed = () => {
    detailData.value = [];
    selectedData.value = null;
    loadingForTable.value = false;
};

/* 挂载时加载list */
onMounted(() => {
    getList();
});
</script>

<style lang="scss" scoped>
.wrapper {
    height: 95vh;
    width: 80vw;
    min-width: 1200px;
    margin: auto;
}

.banner {
    height: 150px;
    display: flex;
    position: relative;
    border-bottom: #ddd 3px solid;
    background-color: #ffffff;
}

.banner-title {
    position: absolute;
    bottom: 0%;
    display: flex;
    font-size: 70px;
    color: #169608;
    animation: slideIn 0.4s ease-in-out forwards;
}

@keyframes slideIn {
    from {
        transform: translateX(-5%);
        opacity: 0;
    }

    to {
        transform: translateX(0);
        opacity: 1;
    }
}

.plaza-page {
    width: 100%;
    padding: 25px 0;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    .card {
        display: flex;
        position: relative;
        flex: 0 0 calc(33.33% - 10px);
        box-sizing: border-box;
        height: 200px;
        background-color: #fff;
        box-shadow: 0 0 5px #ccc;
        border-radius: 30px;
        align-items: center;

        &:hover {
            box-shadow: 0 0 15px #ccc;
            cursor: pointer;
        }

        .ifo {
            display: flex;
            flex-direction: column;
            justify-content: center;
            font-size: 18px;
            margin-left: 50px;
            width: 70%;
            border-left: #ddd 2px solid;
            padding: 0 0 0 25px;
            height: 70%;

            p {
                line-height: 28px;
                color: #333;
            }

            .ifo-item {
                color: #67C23A;
            }
        }
    }

}

.icon {
    position: relative;
    left: 25px;
    width: 140px;
    min-width: 140px;
    height: 140px;
    border: #ddd 1px solid;
    border-radius: 15px;
    background-image: url('@/icons/dna_icon.jpg');
    background-size: cover;
    background-position: center;
}

.dt-page-top {
    display: flex;
    align-items: center;

    .icon {
        width: 180px;
        height: 180px;
    }

    .ifo {
        margin-left: 50px;
    }

    p {
        color: #333;
        font-size: 20px;
        line-height: 28px;
    }

    .ifo-item {
        color: #67C23A;
    }
}

.dt-page-bottom {
    padding: 20px;
}

:deep(.el-popup-parent--hidden) {
    width: 100% !important;
}

:deep(.obtain-btn) {
    font-size: 20px;
    border: #169608 1px solid;
    background-color: #169608;
    box-shadow: none;
}

// :deep(.el-loading-mask){
//     z-index: 3000 !important;
// }</style>