import { defineStore } from "pinia";

const useWalletStore = defineStore("account", {
    state: () => ({
        address: null,
        insititution: null,
        access_token: null,
        email: null,
    }),
    mutations: {

    },
    actions: {
        setAddress(address) {
            this.address = address;
        },
        setInstitution(insititution) {
            this.insititution = insititution;
        },
        setToken(token) {
            this.access_token = token;
            console.log(this.access_token);
            /*
            *不打印就是无法储存access_token，是本项目最大最莫名其妙最逆天的bug
            *并且打印其他数据也无法储存  
            */
        },
        setEmail(email) {
            this.email = email;
        }
    },
    persist: {
        storage: sessionStorage,
        /*使用 sessionStorage 而不是 localStorage。MetaMask是会话级通信，并不需要长久保存 */
    },

})


export { useWalletStore };
