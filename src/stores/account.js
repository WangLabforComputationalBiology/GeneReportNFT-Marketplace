import { defineStore } from "pinia";

export const useWalletStore = defineStore("Wallet", {
    state: () => ({
        address: null,
        balance: 0,
        country: null,
    }),
    mutations: {
        
    },
    actions: {
        setAddress(address) {
            this.address = address;
        },
        setBalance(balance) {
            this.balance = balance;
        },
    },
    persist: true,
})