import { defineStore } from "pinia";
import axios from "axios";
import type { VideoInfo } from "src/types";

function convert(
    data: object
):VideoInfo[]{
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
        } as VideoInfo
    })
}
type Option<T> = T | undefined;

function convertOne(
    videoInfoString: string
):VideoInfo{
    const item = JSON.parse(videoInfoString) 
    return {
        id: item.id,
        title: item.title,
        author: item.author,
        url: item.url,
        content: item.content,
        publish_time: item.publish_time,
        thumbnail: item.thumbnail,
        tags: item.tags,

        status: item.status,
        reason: item.reason,
        percent: item.percent,

        size: item.size,
        already_download_size: item.already_download_size,

        type: item.type,
        source: item.source,
        
        parent: item.parent,
        length: item.length,
        episode: item.episode,
        children: item.children,

        start_download_time: item.start_download_time,
        download_speed: item.download_speed,
        end_download_time: item.end_download_time,
    } as VideoInfo;
}

export const useHistoryStore = defineStore("history", {
    state: () => ({
        historyData: [] as VideoInfo[],
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
        async updateVideoInfo(videoInfo: Option<VideoInfo>) {
            const res = (await axios.post("api/update",videoInfo)).data;
            return res;
        },
        getVideoInfoById(id: string): VideoInfo | undefined {
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
            console.log("update ")
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