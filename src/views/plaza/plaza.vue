<template>
    <div class="wrapper">
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
            <el-button class="obtain-btn btn-location" @click="writeContract">Refresh</el-button>

            <div class="download">&gt;Download decrypt tool</div>
            <div class="page">
                <div class="previous" v-show="page > 1" @click="page--">&lt;Previous</div>
                <div style="color: #E6A23C;font-size: 18px;font-weight: 700;">Page:{{ page }}</div>
                <div class="next" @click="page++">Next&gt;</div>
            </div>
        </div>

        <el-scrollbar max-height="80vh">
            <el-empty description="Empty...Maybe try it again?" v-if="List == null || List == '' || List == undefined"
                v-loading="loadingForPlaza" :element-loading-svg="svg">
                <el-button class="obtain-btn" @click="getList">Refresh</el-button>
            </el-empty>
            <div class="plaza-page" v-else>
                <div class="card" v-for="(item, index) in List" :key="index" @click="isVisible(item)">
                    <div class="icon" />
                    <div class="ifo">
                        <p>Name: <span class="ifo-item"> &nbsp; {{ item.name }}</span></p>
                        <p>Category: <span class="ifo-item"> &nbsp; {{ item.category }}</span></p>
                        <p>Description:<span class="ifo-item"> &nbsp; {{ item.description }}</span></p>
                        <p>Format:<span class="ifo-item"> &nbsp; {{ item.format }}</span></p>
                        <p>Date:<span class="ifo-item"> &nbsp; {{ item.created_at }}</span></p>
                        <span v-show="item.is_sharable" class="sharable">Sharable</span>
                        <!-- <p>Sharable:<span class="ifo-item"> &nbsp; {{ item.is_sharable }}</span></p> -->
                    </div>
                </div>
            </div>
        </el-scrollbar>
    </div>

    <el-drawer v-model="drawerIsVisible" title="Detail" :direction="'rtl'" @closed="handleClosed" :size="'55%'"
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
                <el-table-column v-for="col in column(selectedData.category)" :key="col.prop" :prop="col.prop"
                    :label="col.label" />
            </el-table>
        </div>
        <template #footer>
            <div style="flex: auto">
                <el-button type="primary" @click="ObtainClick" class="obtain-btn"
                    v-if="walletStore.address && walletStore.email && walletStore.insititution">Obtain</el-button>
                <span v-else style="color: #169608;">Login and verify to obtain</span>
            </div>
        </template>
    </el-drawer>

    <el-dialog v-model="dialogIsVisible" title="Obtaining" width="25%" :show-close="false">
        <el-form :model="form" label-width="auto">
            <el-form-item label="Account: ">
                <span style="color: #169608;">{{ walletStore.address }}</span>
            </el-form-item>
            <el-form-item label="Email: ">{{ walletStore.email }}</el-form-item>
            <el-form-item label="Insititution: ">{{ walletStore.insititution }}</el-form-item>
            <el-form-item label="Purpose:">
                <el-select placeholder="please select one" v-model="purpose">
                    <el-option label="Academic research" value="Academic research" />
                    <el-option label="Bioinformatics and Big Data" value="Bioinformatics and Big Data" />
                    <el-option label="Biopharmaceuticals" value="Biopharmaceuticals" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogIsVisible = false" class="obtain-btn"
                    style="background-color: #fff; color: #169608;">Cancel</el-button>
                <el-button class="obtain-btn" @click="uploadForm">Confirm</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { ethers } from "ethers";
import Api from "@/axios/aixos";
import { useWalletStore } from "@/stores/account";
import { ElMessage } from 'element-plus'
import abi from "@/assets/sharingPlatform.json";
const walletStore = useWalletStore();

/* 获取plaza卡片数据列表 */
const loadingForPlaza = ref(false);
const List = ref([]);
const page = ref(1);
const getList = async () => {
    try {
        List.value = [];
        loadingForPlaza.value = true;
        const res = await Api.get(`/plaza/${page.value}`);
        List.value = res.data.data.multi_metadata;
        if (List.value == null || List.value == '' || List.value == undefined) {
            setTimeout(() => {
                loadingForPlaza.value = false;
            }, 1500)
            return;
        };
        List.value.forEach(item => {
            item.created_at = item.created_at.slice(0, 10);
        })
    } catch (error) {
        if (error) {
            loadingForPlaza.value = false;
        }
    } finally {
        setTimeout(() => {
            loadingForPlaza.value = false;
        }, 1500)
    }
}

/* 翻页 */
watch(page, () => {
    getList();
})

/* 获取详细数据 */
const selectedData = ref([]);
const loadingForTable = ref(false);
const detailData = ref([]);
const getDetailData = async () => {
    try {
        loadingForTable.value = true;
        if (!selectedData.value?.data_hash) return; //问号表示允许不存在
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

/* 表格列封装 */
const columnConfigs = ref({
    'Psychology': [
        { label: 'Score', prop: 'score' },
        { label: 'Rank', prop: 'rank' },
    ],
    'Risk': [
        { label: 'Percent', prop: 'percent' },
        { label: 'Risk', prop: 'risk' },
    ],
    'Athletigen': [
        { label: 'Score', prop: 'score' },
        { label: 'Rank', prop: 'rank' },
    ],
    'Skin': [
        { label: 'Score', prop: 'score' },
        { label: 'Rank', prop: 'rank' },
    ],
    'HealthyMetabolism': [
        { label: 'Description_en', prop: 'description_en' },
        { label: 'Score', prop: 'score' },
        { label: 'Rank', prop: 'rank' },
    ],
    'HealthyCarrier': [
        { label: 'Description_en', prop: 'description_en' },
        { label: 'Mag', prop: 'mag' },
        { label: 'Summary', prop: 'tsummary' },
    ],
    'HealthyTraits': [
        { label: 'Description_en', prop: 'description_en' },
        { label: 'Mag', prop: 'mag' },
        { label: 'Summary', prop: 'tsummary' },
    ],

});
const hasWarned = ref(false);   //避免因为表格渲染重复报错
function column(category) {
    if (!category || category === 'undefined' || category === 'null' || category === 'Unknown' || category === '') {
        if (!hasWarned.value) {
            ElMessage.error('Data error. Category missed.');
            hasWarned.value = true;
        }
        return [];
    }
    return columnConfigs.value[category] || [];   //根据指定category返回列配置
}

/* 释放表单内容 */
const handleClosed = () => {
    detailData.value = [];
    selectedData.value = null;
    loadingForTable.value = false;
    hasWarned.value = false;
};

/* 获取前数据采集 */
let dialogIsVisible = ref(false);
const purpose = ref('');
const ObtainClick = async () => {
    try {
        const res = await Api.get(`/gene_type/${selectedData.value.data_hash}`, {
            data_hash: selectedData.value.data_hash,
        });
        console.log(res)
        if (res.data.data.access_status == false && res.data.message == 'success') {
            dialogIsVisible.value = !res.data.data.access_status; //如果返回值为false，则显示弹窗
        } else {

        }
    } catch (error) {
        console.error('Error fetching detail data:', error);
    } finally {
    }
}

const uploadForm = async () => {
    try {
        const res = await Api.post(`/gene_type/newAccess`, {
            data_hash: selectedData.value.data_hash,
            label: purpose.value,
        });
        if (res.data.message == 'success') {
            dialogIsVisible.value = false; //如果返回值为false，则显示弹窗
            ElMessage.success('Submitted successfully');
        } else {
            ElMessage.error('Submit failed');
        }
    } catch (error) {
        console.error('Error fetching detail data:', error);
    }
}

/* 下载 */
const download = async () => {

}

/**调用合约 */
const contractABI = abi; // 合约ABI
const contractAddress = '0x8c451bbd4b60C6811Ea3E2B98A510fBE83d333eF'; // 合约地址
async function writeContract() {
    // 连接MetaMask提供者
    const provider = new ethers.BrowserProvider(window.ethereum);
    // 获取签名者
    const signer = await provider.getSigner();
    // 创建合约实例
    const contract = new ethers.Contract(contractAddress, contractABI, signer);
    try {
        // 发送交易
        const tx = await contract.addViewAccess();
        // 等待交易确认
        const receipt = await tx.wait();
        if (receipt.status === 1) {
            console.log("交易成功");
            ElMessage.success('Submission success!');
            console.log('Transaction confirmed:', receipt.logs);
        } else {
            ElMessage.error('Submission failed!');
            console.log("交易失败");
        }
    } catch (error) {
        console.error('Error writing to contract:', error);
        throw error;
    }
}
//上传哈希
const uploadHash = async () => {
    
}

/* 加载list */
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
    width: 80vw;
    min-width: 1200px;
    height: 15vh;
    min-height: 100px;
    display: flex;
    position: relative;
    border-bottom: #ddd 1px solid;
    background-color: #ffffff;
    background: url('@/assets/imgs/biochain.svg') no-repeat;
    background-position: 85% -15%;

    .download {
        position: absolute;
        right: 0%;
        top: 50%;
        color: #169608;

        &:hover {
            cursor: pointer;
            text-decoration: underline;
        }
    }

    .page {
        position: absolute;
        display: flex;
        right: 0%;
        bottom: 0%;
        gap: 25px;

        .previous,
        .next {
            cursor: pointer;
            color: #333;
            font-size: 18px;
            font-weight: 700;

            &:hover {
                color: #169608;
            }
        }
    }
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
    padding: 20px 12px;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    .card {
        display: flex;
        position: relative;
        flex: 0 0 calc(33.33% - 10px);
        box-sizing: border-box;
        height: 250px;
        background-color: #fff;
        box-shadow: 0 0 3px #ccc;
        border-radius: 20px;
        align-items: center;

        &:hover {
            box-shadow: 0 0 6px #ccc;
            cursor: pointer;
        }

        .ifo {
            display: flex;
            flex-direction: column;
            justify-content: center;
            font-size: 18px;
            margin-left: 40px;
            width: 75%;
            border-left: #ddd 2px solid;
            padding: 0 0 0 20px;
            height: 70%;

            p {
                line-height: 24px;
                color: #333;
            }

            .ifo-item {
                color: #67C23A;
            }

            .sharable {
                font-size: 16px;
                margin-top: 5px;
                width: 88px;
                line-height: 24px;
                text-align: center;
                color: #E6A23C;
                background-color: #fff;
                border: #E6A23C 1px solid;
                border-radius: 15px;
            }
        }
    }

}

.icon {
    position: relative;
    left: 20px;
    width: 160px;
    min-width: 140px;
    height: 160px;
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
    padding: 0 20px;
}

:deep(.el-popup-parent--hidden) {
    width: 100% !important;
}

:deep(.obtain-btn) {
    font-size: 16px;
    border: #169608 1px solid;
    background-color: #169608;
    color: #fff;
    box-shadow: none;
}

.btn-location {
    position: absolute;
    left: 360px;
    bottom: 10%;
}

:deep(.el-scrollbar) {
    height: 80vh;
}
</style>