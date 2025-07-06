import { defineStore } from "pinia";

const useWalletStore = defineStore("account", {
    state: () => ({
        token: null,
        address: null,
        insititution: null,
        email: null,
    }),
    actions: {
        setToken(token) {
            this.token = token;
        },
        setWalletInfo(address, insititution, email) {
            this.address = address;
            this.insititution = insititution;
            this.email = email;
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
        storage: sessionStorage,
        /*使用 sessionStorage 而不是 localStorage。MetaMask是会话级通信，并不需要长久保存 */
    },

})


export { useWalletStore };
