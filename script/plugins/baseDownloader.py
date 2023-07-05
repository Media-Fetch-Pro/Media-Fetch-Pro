from typing import List

from script.model.videoInfo import VideoInfo

class BaseDownloader():
    def __init__(self, next):
        self.next = next
    
    def isSupport(self,url):
        pass
    
    def getVideoInfo(self,url)->List[VideoInfo]: # 这是一个责任链模式
        if self.isSupport(url):
            pass
        else:
            return self.next.getVideoInfo(url)


    def downloadVideo(self, video_info: VideoInfo, output_dir: str):
        if self.isSupport(video_info.url):
            self._downloadVideo(video_info, output_dir)
        else:
            return self.next.downloadVideo(video_info, output_dir)


    def downloadPoster(self, video_info: VideoInfo, output_dir: str):
        if self.isSupport(video_info.url):
            self._downloadPoster(video_info, output_dir)
        else:
            return self.next.downloadPoster(video_info, output_dir)


    def downloadNfo(self, video_info: VideoInfo, output_dir: str):
        if self.isSupport(video_info.url):
            self.downloadNfo(video_info, output_dir)
        else:
            return self.next.downloadNfo(video_info, output_dir)

