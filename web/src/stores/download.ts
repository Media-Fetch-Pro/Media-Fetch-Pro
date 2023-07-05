import { defineStore } from "pinia";
import axios from "axios";

export const useDownloadStore = defineStore("download", {
    state: () => ({
        url: "",
        // selectStoragePath: "/var/opt/video",
        selectStoragePath: "/Users/ctrdh/Video ",
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