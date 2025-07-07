import { useWalletStore } from '@/stores/account';
import { EventSourcePolyfill } from 'event-source-polyfill';    //SSE服务完善包，方便携带token
const walletStore = useWalletStore();

export class SSEManager {
    constructor(dataArray) {
        this.connection = null;
        this.dataArray = dataArray;
    }

    connect() {
        this.connection = new EventSourcePolyfill(`${import.meta.env.VITE_APP_BASE_URL}/studio/getProfile/uncompleted`, {
            headers: {
                'Authorization': walletStore.token
            }
        });

        this.connection.onopen = () => {
            console.log('SSE连接已建立');
        };

        this.connection.addEventListener('message', (event) => {
            console.log(event.data);
        });

        //progress为返回进度的信息类型
        this.connection.addEventListener('progress', (event) => {
            try {
                const jsonData = JSON.parse(event.data);    //SSE返回的是文本格式，这里转为json
                console.log(jsonData);
                const exists = this.dataArray.value.some(item => item.taskID === jsonData.taskID);  //是否已经存在该元素
                console.log('是否已经存在该元素', exists);
                if (!exists) {
                    this.dataArray.value.push(jsonData);
                } else {
                    //更新已有的数据
                    const index = this.dataArray.value.findIndex(item => item.taskID === jsonData.taskID);
                    this.dataArray.value[index] = jsonData;
                    console.log('更新已有的数据', this.dataArray.value);
                }
            } catch (error) {
                console.log("Error(已断开连接):" + error);
                this.close();
            }
        });

        this.connection.addEventListener('keepalive', (event) => {
            console.log('keepalive: SSE heartbeat');//心跳检测
        });

        this.connection.addEventListener('error', (event) => {
            this.close();
        });

        this.connection.onerror = (event) => {
            this.close();
            console.log('SSE连接已关闭');
        };
    }

    close() {
        if (this.connection) {
            this.connection.close();
            this.connection = null;
            console.log('SSE连接已关闭');
        }
    }
}

