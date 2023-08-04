import { defineStore } from 'pinia'
import { reactive, ref } from "vue";
import { getCleanUrl } from '@/utils/bilibili';
import axios, { type AxiosResponse } from "axios";

export const useDownloadStore = defineStore("download",()=> {
    const url = ref("https://www.bilibili.com/video/BV1ah4y117G3/?spm_id_from=333.999.0.0");

    const website_status = reactive({
        "bilibili_status": false,
        "youtube_status": false
    });
    async function download(storagePath: string): Promise<AxiosResponse<any, any>>{
        const result = await axios.post("api/video", {
            url: getCleanUrl(url.value),
            storage: storagePath,
        }).catch((error) => {
            return error.response;
        });
        return result;
    }
    async function getWebSiteConnectionStatus(){
        const result = await axios.get("api/status");
        const data = result.data;
        website_status.bilibili_status = data.data.bilibili_status;
        website_status.youtube_status = data.data.youtube_status;
    }

    return { url, website_status, download, getWebSiteConnectionStatus }
},{
    persist: true,
}
);

