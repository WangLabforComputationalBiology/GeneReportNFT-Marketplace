import axios from 'axios';
import { useWalletStore } from '@/stores/account';
const walletStore = useWalletStore();

const Api = axios.create({
    baseURL: 'http://10.108.10.51:7070',
});


/*请求头 */
Api.defaults.headers.common['Authorization'] = walletStore.access_token;
Api.defaults.headers.post['Content-Type'] = 'application/json';
Api.defaults.headers.put['Content-Type'] = 'application/json';
export default Api;