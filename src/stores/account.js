import { defineStore } from "pinia";

const useWalletStore = defineStore("account", {
    state: () => ({
        token: null,
        address: null,
        insititution: null,
        email: null,
    }),
    actions: {
        setAddress(address) {
            this.address = address;
        },
        setInstitution(insititution) {
            this.insititution = insititution;
        },
        setToken(token) {
            this.token = token;
            // console.log(this.token);
            /*
            *不打印就是无法储存token，是本项目最大最莫名其妙最逆天的bug
            *并且打印其他数据也无法储存  
            */
        },
        setEmail(email) {
            this.email = email;
        },
        reset() {
            this.address = null;
            this.insititution = null;
            this.token = null;
            this.email = null;
        },
    },
    persist: {
        storage: sessionStorage,
        /*使用 sessionStorage 而不是 localStorage。MetaMask是会话级通信，并不需要长久保存 */
    },

})


export { useWalletStore };
