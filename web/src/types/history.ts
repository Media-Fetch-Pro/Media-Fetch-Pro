export interface VideoInfo {
    id: string;
    title: string;
    author: string;
    url: string;
    content: string;
    publish_time: number;
    thumbnail: string;
    tags: string;

    status: string;
    reason: string;
    percent: number;

    size: number;
    already_download_size: number;

    type: string;    
    source: string;

    parent: string;
    length: number;
    episode: number;
    children: string[];

    start_download_time: number;
    download_speed: number;
    end_download_time: number;
}