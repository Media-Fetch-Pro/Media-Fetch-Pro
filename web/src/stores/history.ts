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
    data: any
):DownloadHistory{
    return {
        id: data.id,
        title: data.title,
        url: data.url,
        status: data.status,
        percent: data.percent,
        size: data.size,
        type: data.type,
        // alreadyDownloadSize: data.AlreadyDownloadSize,
        // collectionId: data.CollectionId,
    } as DownloadHistory
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
        async updateVideoStatus(object: any) {
            const videoInfo = convertOne(object);
            this.historyData = this.historyData.map((item) => {
                if (item.id === videoInfo.id) {
                    return videoInfo;
                }
                return item;
            });
        },
    },
});