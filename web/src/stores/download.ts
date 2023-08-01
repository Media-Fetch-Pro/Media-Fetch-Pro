import { defineStore } from 'pinia'
import { ref } from "vue";
import axios, { type AxiosResponse } from "axios";

export const useDownloadStore = defineStore("download",()=> {
    const url = ref("https://www.youtube.com/watch?v=9bZkp7q19f0");

    async function download(storagePath: string): Promise<AxiosResponse<any, any>>{
        const result = await axios.post("api/video", {
            url: url.value,
            storage: storagePath,
        }).catch((error) => {
            return error.response;
        });
        return result;
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
