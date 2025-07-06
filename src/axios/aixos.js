import axios from 'axios';
import { useWalletStore } from '@/stores/account';
const baseURL = import.meta.env.VITE_APP_BASE_URL;

const Api = axios.create({
    baseURL: baseURL,
    headers: {
        'Authorization': useWalletStore().token,//这里直接使用useWalletStore()才能保证数据的动态性
        'Content-Type': 'application/json'
    }
});

export default Api;