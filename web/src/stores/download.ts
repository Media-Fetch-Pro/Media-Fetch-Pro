import { defineStore } from 'pinia'
import { ref } from "vue";
import axios from "axios";

export const useDownloadStore = defineStore("download",()=> {
    const url = ref("https://www.youtube.com/watch?v=9bZkp7q19f0");
    const selectStoragePath = ref([]);

    function download(storagePath: string){
        axios.post("api/video", {
            url: url.value,
            storage: storagePath,
        })
    }
    return { url, download }
},{
    persist: true,
}
);

// state: () => ({
//     url: "",
//     // selectStoragePath: "/var/opt/video",
//     // selectStoragePath: "/Users/ctrdh/video",
//     StoragePathOptions: [],
// }),
// getters: {
// },    
// actions: {
//     download(storagePath: string) {
//         axios.post("api/video", {
//             url: this.url,
//             storage: storagePath,
//         })
//     },
// },
