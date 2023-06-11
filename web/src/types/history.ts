import type { type } from "os";

export interface History {
    id: string;
    title: string;
    url: string;
    status: string;
    percent: number;
    size: number;
    type: string;
    alreadyDownloadSize: number;
    collectionId: string;
}