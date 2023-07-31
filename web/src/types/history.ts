export interface DownloadHistory {
    id: string;
    title: string;
    url: string;
    status: string;
    percent: number;
    size: number;
    type: string;
    children: string[];
    author: string;
    source: string;
    content: string;
    episode: number;
    parent: string;
    length: number;
    start_download_time: number;
}