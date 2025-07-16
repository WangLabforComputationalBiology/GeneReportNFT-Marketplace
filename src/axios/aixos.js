import axios from 'axios';
import { useWalletStore } from '@/stores/account';
const baseURL = import.meta.env.VITE_APP_BASE_URL;  //VITE_APP_BASE_URL=http://10.108.10.51:7070

const Api = axios.create({
    baseURL: baseURL,
    headers: {
        //在此处设置token有个弊端在于axios只初始化一次，导致更换账户时token残留
    }
});

/**请求拦截器*/
Api.interceptors.request.use((config) => {
    const walletStore = useWalletStore();
    if (walletStore.token) {
        config.headers.Authorization = walletStore.token; // 确保每次请求携带最新token
    } else {
        console.log('No token.');
    }
    return config;
});

// axios.interceptors.response.use(
//     response => response,
//     error => {
//         if (error.response.status === 401) {
//             window.location.href = '/login';
//         }
//         return Promise.reject(error);
//     }
// );

export default Api;