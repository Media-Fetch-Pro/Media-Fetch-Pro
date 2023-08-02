import { defineStore } from "pinia";
import type { downloadPath } from "src/types/setting";

export const useSettingStore = defineStore("setting", {
    state: () => ({
        // these setting will be saved in backend in future
        storagePath: "/var/opt/video",
        downloadPath: [
            {
                path: "/var/opt/video",
                name: "default",
            } as downloadPath,
            {
                path: "/var/opt/video2",
                name: "Auto",
            } as downloadPath,
            {
                path: "/Users/ctrdh/video",
                name: "DIY",
            } as downloadPath,
        ],
        maxDownload: 1,
        maxDownloadSpeed: 5,
        language: "en",
        theme: "light",
    }),
    getters: {
    },    
    actions:{
        setStoragePath(path: string) {
            this.storagePath = path;
        },
        newDownloadPath(name: string,path: string): boolean {
            name = name.trim();
            path = path.trim();
            // to check if the name is already exist
            for (let i = 0; i < this.downloadPath.length; i++) {
                if (this.downloadPath[i].name == name) {
                    return false; // TODO how to return error?
                }
            }
            this.downloadPath.push({
                path: path,
                name: name,
            });
            return true;
        },
        removeDownloadPath(name: string) {
            for (let i = 0; i < this.downloadPath.length; i++) {
                if (this.downloadPath[i].name == name) {
                    this.downloadPath.splice(i, 1);
                    return;
                }
            }
        }
    },
    persist: true,
});