import { defineStore } from "pinia";

const useRoutesLoadingStore = defineStore("routesLoading", {
    state: () => ({
        isLoading: false,
    }),
    mutations: {
        
    },
    actions: {
        Loading(loading) {
            this.isLoading = loading;
        },
    }
})


export { useRoutesLoadingStore };