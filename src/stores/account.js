import { defineStore } from "pinia";

export const useWalletStore = defineStore("Wallet", {
    state: () => ({
        address: null,
        insitution: null,
    }),
    mutations: {
        
    },
    actions: {
        setAddress(address) {
            this.address = address;
        },
        setInsitution(insitution) {
            this.insitution = insitution;
        },
    },
    persist: true,
})