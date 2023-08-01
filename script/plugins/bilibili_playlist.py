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
    def _handleMultiLine(self, content: str) -> str:
        return f"<![CDATA[{content}]]>"

    def _generateTvShowNfo(self, video_info: VideoInfo) -> str:
        content = f"""<?xml version="1.0" encoding="UTF-8"?>
<plot>{self._handleMultiLine(video_info.get_content())}</plot>
<outline>{self._handleMultiLine(video_info.get_content())}</outline>
<lockdata>false</lockdata>
<dateadded>2023-02-04 22:27:18</dateadded>
<title>{video_info.get_title()}</title>
<originaltitle>{video_info.get_title()}</originaltitle>
<director>{video_info.get_author()}</director>
<rating>6.6</rating>
<year>2022</year>
<premiered>2016-07-09</premiered>
<releasedate>2016-07-09</releasedate>
<genre>剧情</genre>
<season>-1</season>
<episode>-1</episode>
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

