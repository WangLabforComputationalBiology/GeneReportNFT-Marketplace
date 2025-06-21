import { defineStore } from "pinia";

export const useWalletStore = defineStore("account", {
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
            *不打印就是无法储存，是本项目最大最莫名其妙最逆天的bug
            *而且打印其他数据都无法储存  
            */
        },
        setEmail(email) {
            this.email = email;
        }
    },
    persist: true,

})