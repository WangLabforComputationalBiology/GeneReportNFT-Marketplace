import axios from 'axios';
import { useWalletStore } from '@/stores/account';

const baseURL = import.meta.env.VITE_APP_BASE_URL;
const walletStore = useWalletStore();

const Api = axios.create({
    baseURL: baseURL,
    headers: {
        'Authorization': walletStore.access_token,
        'Content-Type': 'application/json'
    }
});

export default Api;