import { defineStore } from "pinia";

const useWalletStore = defineStore("account", {
    state: () => ({
        token: null,
        address: null,  //钱包地址
        insititution: null,      //机构
        email: null,
    }),
    actions: {
        setWalletInfo(address, insititution, email) {
            this.address = address;
            this.insititution = insititution;
            this.email = email;
        },
        setToken(token) {
            this.token = token;
        },
        setAddress(address) {
            this.address = address;
        },
        setInstitution(insititution) {
            this.insititution = insititution;
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
        //使用 sessionStorage 而不是 localStorage。MetaMask设计是会话级通信，不需要长久保存
        storage: sessionStorage,
    },

})

export { useWalletStore };
