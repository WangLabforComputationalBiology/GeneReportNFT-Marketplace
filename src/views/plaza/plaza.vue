<template>
    <div class="wrapper">
        <div class="banner">
            <h1 class="banner-title">Data&nbsp;<span style="color: #333;">Plaza</span></h1>
            <el-button class="obtain-btn btn-location" @click="getList">Refresh</el-button>

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
                    <div class="info">
                        <p>Name: <span class="info-item"> &nbsp; {{ item.name }}</span></p>
                        <p>Description:<span class="info-item"> &nbsp; {{ item.description }}</span>
                        </p>
                        <p>Format:<span class="info-item"> &nbsp; {{ item.format }}</span></p>
                        <div v-show="item.is_sharable" class="sharable">Sharable</div>
                    </div>
                </div>
            </div>
        </el-scrollbar>
    </div>

    <el-drawer v-model="drawerIsVisible" title="Detail" :direction="'rtl'" @closed="handleClosed" :size="'60%'"
        :show-close="false">
        <template #header>
            <div class="dt-page-top">
                <div class="icon" />
                <div class="info">
                    <p>Name: <span class="info-item"> &nbsp; {{ selectedData.name }}</span></p>
                    <p>Sex: <span class="info-item"> &nbsp; {{ selectedData.sex == true ? 'Male' : 'Female' }}</span>
                    </p>
                    <p>Format:<span class="info-item"> &nbsp; {{ selectedData.format }}</span></p>
                    <p>Date:<span class="info-item"> &nbsp; {{ selectedData.created_at }}</span></p>
                    <div class="category">
                        <p>Category:</p>
                        <el-select v-model="selectedData.category" placeholder="Category">
                            <el-option label="Psychology" value="Psychology" />
                            <el-option label="HealthyMetabolism" value="HealthyMetabolism" />
                            <el-option label="HealthyTraits" value="HealthyTraits" />
                            <el-option label="Skin" value="Skin" />
                            <el-option label="Athletigen" value="Athletigen" />
                            <el-option label="HealthyCarrier" value="HealthyCarrier" />
                            <el-option label="Risk" value="Risk" />
                        </el-select>
                    </div>
                </div>
                <div class="info-description">
                    <p>Description:<span class="info-item"> &nbsp; {{ selectedData.description }}</span></p>
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
            <div>
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
                <el-button class="obtain-btn" @click="writeContract">Confirm</el-button>
                <p style="margin-top: 30px;text-align: left; font-size: 14px;color:#888">If Metamask does not have a
                    pop-up, please open it to see if "Activity" is suspended, and confirm.</p>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { ethers } from "ethers";
import Api from "@/axios/aixos";
import { useWalletStore } from "@/stores/account";
import { ElLoading, ElMessage } from 'element-plus'
import abi from "@/assets/sharingPlatform.json";
const walletStore = useWalletStore();

/* 获取plaza卡片数据列表 */
const loadingForPlaza = ref(false);
const List = ref([]);   //获取的plaza卡片列表  [{}]
const page = ref(1);    //页数
const getList = async () => {
    try {
        List.value = [];
        loadingForPlaza.value = true;
        const res = await Api.get(`/plaza/${page.value}`);
        List.value = res.data.data.multi_metadata || [];
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
        if (error.response.status == 404) {
            ElMessage.error('Network error. Please try again later.');
            return;
        }
        loadingForPlaza.value = false;

    } finally {
        setTimeout(() => {
            loadingForPlaza.value = false;
        }, 1500)
    }
}
getList();  //获取列表

/* 获取详细数据 */
const selectedData = ref({});   //选中查看的卡片
const loadingForTable = ref(false);
const detailData = ref([]);     //根据选中卡片hash获取的报告详情 [{}]
const getDetailData = async () => {
    try {
        loadingForTable.value = true;
        if (!selectedData.value?.data_hash) return; //问号表示允许不存在
        const res = await Api.get(`/metadata/${selectedData.value.data_hash}?t=${Date.now()}`);
        if (res.data.data) {
            detailData.value = res.data.data.details;
            loadingForTable.value = false;
        }
    } catch (error) {
        console.error('Error fetching detail data:', error);
        loadingForTable.value = false;
        ElMessage.error('Error fetching detail data');
    }
}

/* 详情抽屉 */
const drawerIsVisible = ref(false);
function isVisible(item) {
    selectedData.value = item || [];
    getDetailData();
    drawerIsVisible.value = true;
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

function column(category) {
    if (category === 'undefined' || category === 'Unknown' || category === '' || category === null) {
        ElMessage.error('Data error. Category missed.');
        return [];
    }
    return columnConfigs.value[category] || [];   //根据指定category返回列配置
}

/* 释放表单内容 */
const handleClosed = () => {
    detailData.value = [];
    selectedData.value = {};
    loadingForTable.value = false;
};

/* 获取前数据采集 */
let dialogIsVisible = ref(false);
const purpose = ref('');
const ObtainClick = async () => {
    try {
        const res = await Api.get(`/gene_type/${selectedData.value.data_hash}`, {
            data_hash: selectedData.value.data_hash,
        });
        console.log(res.data.data)
        // if (res.data.data.access_status == false) {
        //     dialogIsVisible.value = true; //如果返回值为false，则显示弹窗
        // }
    } catch (error) {
        console.error('No rights(没有权限，将进行获权操作). Error: ', error);
        if (error.status == 403 && error.response.data.code == 403) {
            dialogIsVisible.value = true; //如果返回值为false，则显示弹窗
        }
    }
}

/* 下载 */
const download = async () => {

}


/**调用合约_test */
// import abi_test from '@/assets/test.json'
// const test_contract = abi_test;
// const test_address = '0x0958817F161D6c9Ee7974Bff07f354E410632Eb1';//测试合约地址
// const test_hash = '89c792eed9551d2b477e5b300b6dfc26c11ab4ccd72a3d44899c5b1b69a52123';

/**调用合约 */
const contractABI = abi; // 合约ABI(固定，每个合约对应一个abi)
const contractAddress = '0xcde2e7e1483716b491B9f38F7908161414A6260d'; // 合约地址(固定，每次部署合约生成一个)
async function writeContract() {
    if (purpose.value == '') {
        ElMessage.error('Please select a purpose');
        return
    }
    const loading = ElLoading.service({
        lock: true,
        text: 'Loading...',
        background: 'rgba(0, 0, 0, 0.7)'
    })
    try {
        // 连接MetaMask提供者
        const provider = new ethers.BrowserProvider(window.ethereum);
        // 获取签名者
        const signer = await provider.getSigner();
        // 创建合约实例（合约地址，合约ABI，签名者）
        const contract = new ethers.Contract(contractAddress, contractABI, signer);
        const rawValue = "0x" + selectedData.value.data_hash;//调用合约要求0x开头的byteslike类型，实际上这里字符串直接拼接也可以
        // 发送transaction
        ElMessage.warning('Please confirm the transaction in MetaMask.');
        const tx = await contract.obtainViewAccess(selectedData.value.geneSharing_contract_address, rawValue, purpose.value);

        /**合约测试 */
        // if (!ethers.isAddress(selectedData.value.geneSharing_contract_address)) {
        //     throw new Error("无效的地址");
        // }
        // console.log('合约地址：',selectedData.value.geneSharing_contract_address);
        // const rawValue = "0x" + test_hash;//调用合约要求0x开头的byteslike类型
        // const BytesLikeDataHash = ethers.keccak256(rawValue) //如果看到让你用ethers.utils.keccak256()，这是v5的写法，v6直接顶层调用
        // console.log("哈希值：",BytesLikeDataHash,"原始值：",rawValue);
        // console.log(purpose.value)
        // const tx = await contract.setString(purpose.value);
        /** */

        // 等待交易确认后接受回执receipt
        const receipt = await tx.wait();
        if (receipt.status === 1) {
            ElMessage.success('Submission success!');
            dialogIsVisible.value = false;
            const receivedHash = receipt.hash || '';   //通过日志获取返回的交易哈希
            console.log('Transaction hash(交易哈希):', receipt.hash);
            uploadReceivedHash(receivedHash);   //返回交易哈希校验
        }
    } catch (error) {
        console.error('Error writing to contract:', error);
        //取消操作反馈
        if (error.code == 'ACTION_REJECTED' || error.reason == 'rejected') {
            ElMessage.error('User action denied.');
        }
    } finally {
        loading.close();
    }
}
//上传哈希
async function uploadReceivedHash(receivedHash) {
    try {
        const res = await Api.post('/gene_type/obtainAccess', {
            tx_hash: receivedHash
        })
        if (res.data.data.code == 200) {
            ElMessage.success('Submission success!');
            console.log('uploadReceivedHash success!');
        }
    } catch (error) {
        console.error('uploadReceivedHash error:', error);
        // ElMessage.error('Submission failed!');
        if(error.status == 400)
        {
            ElMessage.error('BAD REQUEST.Please try again later!');
        }
    }
}

watch(page, () => {
    getList();  //页数改变时，获取新的列表
})

onMounted(() => {
    // getList();  //页面刷新即加载list,似乎更慢,改为在js中加载
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
        height: 230px;
        background-color: #fff;
        box-shadow: 0 0 3px #ccc;
        border-radius: 20px;
        align-items: center;

        &:hover {
            box-shadow: 0 0 6px #ccc;
            cursor: pointer;
        }

        .info {
            display: flex;
            flex-direction: column;
            position: relative;
            font-size: 18px;
            margin-left: 40px;
            width: 60%;
            border-left: #ddd 2px solid;
            padding: 0 0 0 20px;
            height: 70%;

            p {
                line-height: 26px;
                color: #333;
            }

            .info-item {
                color: #67C23A;
            }

            .sharable {
                position: absolute;
                bottom: 2%;
                font-size: 16px;
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
    width: 150px;
    min-width: 140px;
    height: 150px;
    border: #ddd 1px solid;
    border-radius: 15px;
    background-image: url('@/icons/dna_icon.jpg');
    background-size: cover;
    background-position: center;
}

.dt-page-top {
    display: flex;

    .icon {
        padding: 0 10px;
        width: 180px;
        height: 180px;
    }

    .info {
        min-width: 200px;
        width: 300px;
        display: flex;
        flex-direction: column;
        margin-left: 60px;
    }

    p {
        color: #333;
        font-size: 20px;
        line-height: 30px;
    }

    .info-item {
        color: #67C23A;
    }

    .info-description {
        overflow: auto;
        width: 50%;
        margin-left: 60px;
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
    margin-right: 20px;
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