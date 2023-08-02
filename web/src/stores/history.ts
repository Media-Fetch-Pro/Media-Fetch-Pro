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
            children: item.children,
            author: item.author,
            source: item.source,
            content: item.content,
            episode: item.episode,
            parent: item.parent,
            length: item.length,
            start_download_time: item.start_download_time,
            // alreadyDownloadSize: item.AlreadyDownloadSize,
            // collectionId: item.CollectionId,
        } as DownloadHistory
    })
}
type Option<T> = T | undefined;

function convertOne(
    videoInfoString: string
):DownloadHistory{
    const item = JSON.parse(videoInfoString) 
    return {
        id: item.id,
        title: item.title,
        url: item.url,
        status: item.status,
        percent: item.percent,
        size: item.size,
        type: item.type,
        children: item.children,
        author: item.author,
        source: item.source,
        content: item.content,
        episode: item.episode,
        parent: item.parent,
        length: item.length,
        start_download_time: item.start_download_time,
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
        async updateVideoInfo(videoInfo: Option<DownloadHistory>) {
            const res = (await axios.post("api/update",videoInfo)).data;
            return res;
        },
        getVideoInfoById(id: string): DownloadHistory | undefined {
            return this.historyData.find((item) => item.id === id);
        },
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