import axios from 'axios';
import { useWalletStore } from '@/stores/account';
const baseURL = import.meta.env.VITE_APP_BASE_URL;
//VITE_APP_BASE_URL=http://10.108.10.51:7070

const Api = axios.create({
    baseURL: baseURL,
    headers: {
        // 'Authorization': useWalletStore().token, //axios.create只会进行一次初始化，导致更换账户时token没有更新
        // 'Content-Type': 'application/json'
    }
});

/**添加请求拦截器，动态更新 token
 */
Api.interceptors.request.use((config) => {
    const walletStore = useWalletStore();
    if (walletStore.token) {
        config.headers.Authorization = walletStore.token; // 每次请求前重新设置
    }
    return config;
});

export default Api;