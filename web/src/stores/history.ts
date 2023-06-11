import { defineStore } from "pinia";
import axios from "axios";
import type { History } from "src/types";

function convert(
    data: any,
):History[]{
    return data.map((item:any) => {
        return {
            id: item.id,
            title: item.title,
            url: item.url,
            status: item.status,
            percent: item.percent,
            size: item.size,
            type: item.type,
            alreadyDownloadSize: item.alreadyDownloadSize,
            collectionId: item.collectionId,
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
            this.historyData = historyData;
        },
    },
});