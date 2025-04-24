<template>

    <body>
        <div class="banner">
            <h1 class="banner-title">Market</h1>
        </div>
        <div class="marketPage">
            <div class="nfts-wrapper">
                <el-tabs v-model="activeName">
                    <el-tab-pane label="classification1" name="first">
                        <el-table style="width:100%;" :title="NFTs" class="nft-table" :data="tableData">
                            <el-table-column type="expand" width="40px">
                                <template v-slot="props">
                                    <el-form label-position="left" class="demo-table-expand">
                                        <el-form-item label="ID:">
                                            <span>{{ props.row.id }}</span>
                                        </el-form-item>
                                        <el-form-item label="Name:">
                                            <span>{{ props.row.nft_name }}</span>
                                        </el-form-item>
                                        <el-form-item label="price:">
                                            <span>{{ props.row.price + " ETH" }}</span>
                                        </el-form-item>
                                        <el-form-item label="Sales:">
                                            <span>{{ props.row.sales }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Address:">
                                            <span>{{ props.row.address }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Description:">
                                            <span>{{ props.row.description }}</span>
                                            <span></span>
                                        </el-form-item>
                                    </el-form>
                                </template>

                            </el-table-column>
                            <el-table-column prop="id" label="#" width="80px"></el-table-column>
                            <el-table-column prop="nft_name" label="Collection" width="" align="left"></el-table-column>
                            <el-table-column prop="price" label="Price" width="180px" align="right"></el-table-column>
                            <el-table-column prop="sales" label="Sales" width="200px" align="right"></el-table-column>
                            <el-table-column prop="limit" label="Limit" width="100px" align="right"></el-table-column>
                            <el-table-column prop="address" label="Address" align="right" width="300px">
                            </el-table-column>
                            <el-table-column width="120px" align="right">
                                <template v-slot="scope">
                                    <el-button size="mini"
                                        @click="purchase(scope.row.id, scope.row.price)">purchase</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-tab-pane>


                    <el-tab-pane label="classification2" name="second">
                        <el-table style="width:100%;" :title="NFTs" class="nft-table" :data="tableData">
                            <el-table-column type="expand" width="40px">
                                <template v-slot="props">
                                    <el-form label-position="left" class="demo-table-expand">
                                        <el-form-item label="ID:">
                                            <span>{{ props.row.id }}</span>
                                        </el-form-item>
                                        <el-form-item label="Name:">
                                            <span>{{ props.row.nft_name }}</span>
                                        </el-form-item>
                                        <el-form-item label="price:">
                                            <span>{{ props.row.price + " ETH" }}</span>
                                        </el-form-item>
                                        <el-form-item label="Sales:">
                                            <span>{{ props.row.sales }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Address:">
                                            <span>{{ props.row.address }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Description:">
                                            <span>{{ props.row.description }}</span>
                                            <span></span>
                                        </el-form-item>
                                    </el-form>
                                </template>
                            </el-table-column>
                            <el-table-column prop="id" label="#" width="80px"></el-table-column>
                            <el-table-column prop="nft_name" label="Collection" width="" align="left"></el-table-column>
                            <el-table-column prop="price" label="Price" width="180px" align="right"></el-table-column>
                            <el-table-column prop="sales" label="Sales" width="200px" align="right"></el-table-column>
                            <el-table-column prop="address" label="Address" align="right" width="300px">
                            </el-table-column>
                            <el-table-column width="120px" align="right">
                                <el-button size="mini">purchase</el-button>
                            </el-table-column>
                        </el-table>
                    </el-tab-pane>


                    <el-tab-pane label="classification3" name="third">
                        <el-table style="width:100%;" :title="NFTs" class="nft-table" :data="tableData">
                            <el-table-column type="expand" width="40px">
                                <template v-slot="props">
                                    <el-form label-position="left" class="demo-table-expand">
                                        <el-form-item label="ID:">
                                            <span>{{ props.row.id }}</span>
                                        </el-form-item>
                                        <el-form-item label="Name:">
                                            <span>{{ props.row.nft_name }}</span>
                                        </el-form-item>
                                        <el-form-item label="price:">
                                            <span>{{ props.row.price + " ETH" }}</span>
                                        </el-form-item>
                                        <el-form-item label="Sales:">
                                            <span>{{ props.row.sales }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Address:">
                                            <span>{{ props.row.address }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Description:">
                                            <span>{{ props.row.description }}</span>
                                            <span></span>
                                        </el-form-item>
                                    </el-form>
                                </template>
                            </el-table-column>
                            <el-table-column prop="id" label="#" width="80px"></el-table-column>
                            <el-table-column prop="nft_name" label="Collection" width="" align="left"></el-table-column>
                            <el-table-column prop="price" label="Price" width="180px" align="right"></el-table-column>
                            <el-table-column prop="sales" label="Sales" width="200px" align="right"></el-table-column>
                            <el-table-column prop="address" label="Address" align="right" width="300px">
                            </el-table-column>
                            <template>
                                <el-table-column width="120px" align="right">
                                    <el-button size="mini">purchase</el-button>
                                </el-table-column>
                            </template>
                        </el-table>
                    </el-tab-pane>


                    <el-tab-pane label="classification4" name="fourth">
                        <el-table style="width:100%;" :title="NFTs" class="nft-table" :data="tableData">
                            <el-table-column type="expand" width="40px">
                                <template v-slot="props">
                                    <el-form label-position="left" class="demo-table-expand">
                                        <el-form-item label="ID:">
                                            <span>{{ props.row.id }}</span>
                                        </el-form-item>
                                        <el-form-item label="Name:">
                                            <span>{{ props.row.nft_name }}</span>
                                        </el-form-item>
                                        <el-form-item label="price:">
                                            <span>{{ props.row.price + " ETH" }}</span>
                                        </el-form-item>
                                        <el-form-item label="Sales:">
                                            <span>{{ props.row.sales }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Address:">
                                            <span>{{ props.row.address }}</span>
                                            <span></span>
                                        </el-form-item>
                                        <el-form-item label="Description:">
                                            <span>{{ props.row.description }}</span>
                                            <span></span>
                                        </el-form-item>
                                    </el-form>
                                </template>
                            </el-table-column>
                            <el-table-column prop="id" label="#" width="80px"></el-table-column>
                            <el-table-column prop="nft_name" label="Collection" width="" align="left"></el-table-column>
                            <el-table-column prop="price" label="Price" width="180px" align="right"></el-table-column>
                            <el-table-column prop="sales" label="Sales" width="200px" align="right"></el-table-column>
                            <el-table-column prop="address" label="Address" align="right" width="300px">
                            </el-table-column>

                            <el-table-column width="120px" align="right">
                                <el-button size="mini" @click="drawerVisible">purchase</el-button>
                            </el-table-column>
                        </el-table>
                    </el-tab-pane>
                </el-tabs>


            </div>
        </div>

    </body>
</template>

<script>
import { ethers } from "ethers";
import { useWalletStore } from "@/stores/account";
const wallet = useWalletStore();
export default {
    name: "market",
    data() {
        return {
            account: "",
            provider: null,
            signer: null,
            recipient: "",
            amount: "",
            txHash: "",
            balance: null,
            activeName: 'first',//默认选中第一个分类
            visible: false,//购买弹窗
            //表单数据
            tableData: [{
                id: 133,
                price: 0.02,
                nft_name: 'NFT1',
                sales: 0,
                limit: 10,
                address: '上海市普陀区金沙江路',
                description: "xxxxx"
            }, {
                id: 212,
                price: 0.02,
                nft_name: 'NFT2',
                sales: 0,
                limit: 10,
                address: '上海市普陀区金沙江路'
            }, {
                id: 331,
                price: 0.02,
                nft_name: 'NFT3',
                sales: 0,
                limit: 10,
                address: '上海市普陀区金沙江路'
            }, {
                id: 423,
                price: 0.02,
                nft_name: 'NFT4',
                sales: 0,
                limit: 10,
                address: '上海市普陀区金沙江路'
            }]
        };
    },

    async created() {

    },
    methods: {
        async connectWallet() {
            if (window.ethereum) {
                try {
                    // 创建 Web3Provider
                    const provider = new ethers.BrowserProvider(window.ethereum);
                    // 请求账户权限
                    const accounts = await provider.send("eth_requestAccounts", []);
                    this.account = accounts[0]; // 获取第一个账户
                    console.log("Account:", this.account);

                    this.signer = await provider.getSigner();// 获取签名者（用于发送交易）
                    console.log("Signer:", this.signer);


                    // const balanceWei = await provider.getBalance(accounts[0]);// 获取余额（返回值为 BigNumber，单位为 wei）
                    // this.balance = ethers.formatEther(balanceWei);// 将余额转换为 ETH 单位
                    console.log("Balance:", this.balance);

                } catch (error) {
                    console.error("连接失败:", error);
                }
            } else {
                alert("请安装 MetaMask");
            }
        },

        // 发送交易
        async sendTransaction() {
            if (!this.account || !this.signer) {
                alert("请先连接 MetaMask");
                return;
            }
            if (!ethers.isAddress(this.recipient)) {
                alert("请输入有效的接收地址！");
                return;
            }
            if (isNaN(this.amount) || parseFloat(this.amount) <= 0) {
                alert("请输入有效的 ETH 数量！");
                return;
            }
            if (!this.signer) {
                console.error("Signer 未初始化");
                return;
            }
            try {
                // 验证账户余额是否足够
                const requiredBalance = ethers.parseEther(this.amount);
                if (this.balance < requiredBalance) {
                    alert("账户余额不足！");
                    return;
                }
                console.log("发送交易...");

                // 获取账户
                const provider = new ethers.BrowserProvider(window.ethereum);
                const accounts = await provider.send("eth_requestAccounts", []);
                this.account = accounts[0];

                // 发送交易
                const signer = await provider.getSigner();
                const tx = await signer.sendTransaction({
                    to: this.recipient,
                    value: ethers.parseEther(this.amount), // 转换 ETH 单位
                });

                console.log("交易发送中...", tx);
                // alert("交易已提交，等待确认...");
                this.$message({
                    message: "交易已提交，等待确认...",
                    type: "success",
                });

                // 等待交易完成
                await tx.wait();
                this.txHash = tx.hash;
                this.$message({
                    message: "交易成功！",
                    type: "success",
                });

            }
            catch (error) {
                console.error("交易失败:", error);
                this.$message({
                    message: "交易失败！",
                    type: "error",
                });
            }

        },

        drawerVisible() {
            this.$nextTick(() => {
                this.visible = true;
            })
        },

        purchase(id, price) {
            this.$router.push({
                path: '/market/purchase',
                query: {
                    id,
                    price
                }
            });
        }
    },


}
</script>

<style lang="scss" scoped>

</style>

<style lang="scss" scoped>
.banner-title {
    font-size: 70px;
    color: #67C23A;
}

.marketPage {
    padding: 10px;
    min-height: 700px;
    overflow: hidden;
}

.nfts-wrapper {
    margin-top: 8px;
}

//下拉栏样式
:deep(.el-table__expanded-cell) {
    color: #67C23A;
}

//表格样式
:deep(.demo-table-expand) {
    margin-left: 50px;
}

:deep(.el-form-item--label-left) {
    margin: 0;
}

:deep(.el-button--mini) {
    border-radius: 10px;
}

:deep(.el-overlay) {
    background-color: rgba(0, 0, 0, 0.5);
    width: 100%;
    overflow: hidden;

}
</style>