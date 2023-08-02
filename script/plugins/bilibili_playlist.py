from script.plugins.baseDownloader import BaseDownloader
from script.model.videoInfo import VideoInfo
import datetime
# this downloader is for download bilibili playlist nfo. not for video🤔
class BilibiliPlayList(BaseDownloader):
    def isSupport(self,url):
        if "bilibili" in url:
            return True
        else:
            return False

    def isSupportWithVideoInfo(self,video_info: VideoInfo) -> bool:
        return self.isSupport(video_info.url) and video_info.type == "playlist" # string type should to be enum type
    def _handleMultiLine(self, content: str) -> str:
        return f"<![CDATA[{content}]]>"

    def _generateTvShowNfo(self, video_info: VideoInfo) -> str:
        content = f"""<?xml version="1.0" encoding="UTF-8"?>
<plot>{self._handleMultiLine(video_info.get_content())}</plot>
<outline>{self._handleMultiLine(video_info.get_content())}</outline>
<lockdata>false</lockdata>
<dateadded>{str(datetime.datetime.fromtimestamp(video_info.get_start_download_time()))}</dateadded>
<title>{video_info.get_title()}</title>
<originaltitle>{video_info.get_title()}</originaltitle>
<director>{video_info.get_author()}</director>
<season>1</season>
<episode>{video_info.get_length()}</episode>
</tvshow>"""
        return content

    def _downloadNfo(self, video_info: VideoInfo, output_dir: str):
        with open(f"{output_dir}/tvshow.nfo", "w") as f:
            f.write(self._generateTvShowNfo(video_info))

    def downloadNfo(self, video_info: VideoInfo, output_dir: str):
        if self.isSupportWithVideoInfo(video_info):
            self._downloadNfo(video_info, output_dir)
        else:
            return self.next.downloadNfo(video_info, output_dir)

