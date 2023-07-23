import { defineStore } from "pinia";

export const useSettingStore = defineStore("download", {
    persist: {
        storage: localStorage,
        paths: ['someState'],
    },
    state: () => ({
        storagePath: "/var/opt/video",
        maxDownload: 1,
        maxDownloadSpeed: 5,
        language: "en",
        theme: "light",
    }),
    getters: {
    },    
    actions:{
    },
});