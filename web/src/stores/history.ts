import { defineStore } from "pinia";
import axios from "axios";
import type { DownloadHistory } from "src/types";

function convert(
    data: object
):DownloadHistory[]{
    return Object.values(data).map((item:any) => {
        return {
            id: item.Id,
            title: item.Title,
            url: item.Url,
            status: item.Status,
            percent: item.Percent,
            size: item.Size,
            type: item.Type,
            alreadyDownloadSize: item.AlreadyDownloadSize,
            collectionId: item.CollectionId,
        } as DownloadHistory
    })
}


export const useHistoryStore = defineStore("history", {
    state: () => ({
        historyData: [] as DownloadHistory[],
        loading: true,
    }),
    getters: {
        getVideoHistory: (state) => state.historyData,
    },    
    actions:{
        async getVideoStatus() {
            this.loading = true;
            const res = (await axios.get("api/videos")).data;
            const historyData = convert(res.data);
            // console.log(historyData)
            this.historyData = historyData;
            this.loading = false;
            return historyData;
        },
    },
});