import { defineStore } from "pinia";
import axios from "axios";
import type { DownloadHistory } from "src/types";

function convert(
    data: object
):DownloadHistory[]{
    return Object.values(data).map((item:any) => {
        return {
            id: item.id,
            title: item.title,
            url: item.url,
            status: item.status,
            percent: item.percent,
            size: item.size,
            type: item.type,
            // alreadyDownloadSize: item.AlreadyDownloadSize,
            // collectionId: item.CollectionId,
        } as DownloadHistory
    })
}

function convertOne(
    videoInfoString: string
):DownloadHistory{
    const videoInfo = JSON.parse(videoInfoString) 
    return {
        id: videoInfo.id,
        title: videoInfo.title,
        url: videoInfo.url,
        status: videoInfo.status,
        percent: videoInfo.percent,
        size: videoInfo.size,
        type: videoInfo.type,
        // alreadyDownloadSize: item.AlreadyDownloadSize,
        // collectionId: item.CollectionId,
    } as DownloadHistory;
}

export const useHistoryStore = defineStore("history", {
    state: () => ({
        historyData: [] as DownloadHistory[],
        loading: true,
        tab: "downloading",
    }),
    getters: {
        getVideoHistory: (state) => {
            return state.historyData.filter(
                (item) => item.status === state.tab
            );
        }
    },    
    actions:{
        async getVideoStatus() {
            this.loading = true;
            const res = (await axios.get("api/videos")).data;
            const historyData = convert(res.data);
            this.historyData = historyData;
            this.loading = false;

            return historyData;
        },
        updateVideoStatus(videoData: any) {
            const videoInfo = convertOne(videoData);
            const index = this.historyData.findIndex((video) => video.id == videoInfo.id);
            if (index == -1) {
                this.historyData.push(videoInfo);
            }else{
                this.historyData[index] = videoInfo;
            }
        },
    },
    persist: true,
});