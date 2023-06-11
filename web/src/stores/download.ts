import { defineStore } from "pinia";
import axios from "axios";

export const useDownloadStore = defineStore("download", {
    state: () => ({
        url: "",
        selectStoragePath: "/home/ctrdh/video",
        StoragePathOptions: [],
    }),
    getters: {
    },    
    actions:{
        async download() {
            axios.post("api/video", {
                url: this.url,
                storage: this.selectStoragePath,
            })
        },
    },
});