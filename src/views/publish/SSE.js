import { useWalletStore } from '@/stores/account';
import { EventSourcePolyfill } from 'event-source-polyfill';    //SSE服务完善包，方便携带token
const walletStore = useWalletStore();

export class SSEManager {
    constructor(dataArray) {
        this.connection = null;
        this.dataArray = dataArray;//传入ref对象（数组）接收数据
    }

    connect() {
        this.connection = new EventSourcePolyfill(`${import.meta.env.VITE_APP_BASE_URL}/studio/getProfile/uncompleted`, {
            headers: {
                'Authorization': walletStore.token
            },
            //重连控制
            initialRetryDelay: 1000,    //初始重试延迟时间
            maxRetryDelay: 10000,    //最大重试延迟时间
            backoff: {
                strategy: 'exponential' // 使用指数退避策略
            },
            heartbeatTimeout: 30000, // 30秒无数据视为超时
            // forcePolyfill: true, // 强制使用 EventSourcePolyfill
        });

        this.connection.onopen = () => {
            console.log('SSE连接已建立');
        };

        this.connection.addEventListener('message', (event) => {
            // console.log(event.data);
        });

        //progress为返回进度的信息类型
        this.connection.addEventListener('progress', (event) => {
            try {
                const jsonData = JSON.parse(event.data);    //SSE返回的是文本格式，这里转为json
                const exists = this.dataArray.value.some(item => item.taskID === jsonData.taskID);  //是否已经存在该元素
                if (!exists) {
                    this.dataArray.value.push(jsonData);
                } else {
                    //更新已有的数据
                    const index = this.dataArray.value.findIndex(item => item.taskID === jsonData.taskID);
                    this.dataArray.value[index] = jsonData;
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

