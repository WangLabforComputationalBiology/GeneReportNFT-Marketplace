import axios from 'axios';

const Api = axios.create({
    baseURL: 'http://10.108.10.51:7070',
    // timeout: 10000, // 可选
});

export default Api;