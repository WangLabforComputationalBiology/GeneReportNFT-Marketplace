import axios from 'axios';

const Api = axios.create({
    baseURL: 'http://120.24.168.132:8080',
    // timeout: 10000, // 可选
});

export default Api;