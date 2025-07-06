import axios from 'axios';
import { useWalletStore } from '@/stores/account';
const baseURL = import.meta.env.VITE_APP_BASE_URL;
//VITE_APP_BASE_URL=http://10.108.10.51:7070

const Api = axios.create({
    baseURL: baseURL,
    headers: {
        'Authorization': useWalletStore().token,//这里直接使用useWalletStore()才能保证数据的动态性
        'Content-Type': 'application/json'
    }
});

/**登录请求 */
const login = axios.create({
    baseURL: baseURL,
    headers: {
        'Content-Type': 'application/json'
    }
})

export { login }
export default Api;