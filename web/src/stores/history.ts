import { defineStore } from "pinia";
import axios from "axios";
import type { History } from "src/types";

function convert(
    data: object
):History[]{
    return Object.values(data).map((item:any) => {
        console.log(item)
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
        } as History
    })
}


export const useHistoryStore = defineStore("history", {
    state: () => ({
        historyData: [] as History[],
    }),
    getters: {
        getVideoHistory: (state) => state.historyData,
    },    
    actions:{
        async getVideoStatus() {
            const res = (await axios.get("api/video")).data;
            const historyData = convert(res.data);
            // console.log(historyData)
            this.historyData = historyData;
            return historyData;
        },
    },
});