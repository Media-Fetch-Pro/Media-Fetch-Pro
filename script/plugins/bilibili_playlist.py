from script.plugins.baseDownloader import BaseDownloader
from script.model.videoInfo import VideoInfo

# this downloader is for download bilibili playlist nfo. not for video🤔
class BilibiliPlayList(BaseDownloader):
    def isSupport(self,url):
        if "bilibili" in url:
            return True
        else:
            return False

    def isSupportWithVideoInfo(self,video_info: VideoInfo) -> bool:
        return self.isSupport(video_info.url) and video_info.type == "playlist" # string type should to be enum type


    def _downloadNfo(self, video_info: VideoInfo, output_dir: str):
        with open(f"{output_dir}/tvshow.nfo", "w") as f:
                f.write("tvshow.nfo")


    def downloadNfo(self, video_info: VideoInfo, output_dir: str):
        if self.isSupportWithVideoInfo(video_info):
            self._downloadNfo(video_info, output_dir)
        else:
            return self.next.downloadNfo(video_info, output_dir)

