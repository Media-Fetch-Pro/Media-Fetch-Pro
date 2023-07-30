import { defineStore } from "pinia";

export const useSettingStore = defineStore("setting", {
    state: () => ({
        // these setting will be saved in backend in future
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
    persist: true,
});