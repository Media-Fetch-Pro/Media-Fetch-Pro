import yt_dlp
import json
import os
import datetime
from typing import List

import script.utils.common as common
import script.api.request as request

from script.plugins.baseDownloader import BaseDownloader
from script.model.videoInfo import VideoInfo
from script.utils.video import generate_uuid_from_url
from script.config.config import Config
from script.utils.video import generate_uuid_from_url
from script.utils.ytdlp import extract_progress

class Youtube(BaseDownloader):
    def progress_hook(self,d):
        self.video_info.set_status('downloading')
        self.video_info.set_percent(extract_progress(d['_percent_str']))
        request.updateVideoStatus(self.video_info)

    def isSupport(self,url):
        if "youtube" in url:
            return True
        else:
            return False
        
    def isSupportWithVideoInfo(self,video_info: VideoInfo) -> bool:
        return self.isSupport(video_info.url) and video_info.type == "video" # string type should to be enum type

    def _initVideoInfo(self, url:str)->VideoInfo:
        video_info = VideoInfo()
        video_info.set_url(url)
        video_info.set_title("title is fetching")
        video_info.set_status("pending")
        video_info.set_percent(0)
        video_info.set_source("youtube")

        # get now unix timestamp as start download time
        presentDate = datetime.datetime.now()
        unix_timestamp = datetime.datetime.timestamp(presentDate)*1000
        video_info.set_start_download_time(int(unix_timestamp))
        return video_info
    
    def _parseVideoInfo(self, video_info: VideoInfo, json_text: str)->List[VideoInfo]:
        video_json = json.loads(json_text)
        video_info.set_title(video_json["title"])
        p_video_array = []
        if video_json['_type'] == 'playlist':
            print(video_json)
            video_info.set_author(video_json["uploader"])
            video_info.set_content(video_json["description"])
            video_info.set_type("video")
            video_info.set_length(int(video_json["playlist_count"]))

            # TODO How to fetch playlist video info in Youtube
            return [video_info]
        else:
            # if the url is video
            video_info.set_author(video_json["uploader"])
            video_info.set_content(video_json["description"])
            video_info.set_type("video")
            video_info.set_size(video_json["filesize_approx"])
            return [video_info]


    def _fetchVideoInfo(self, video_info: VideoInfo)->List[VideoInfo]:
        common.runShell(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        print(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.info.json")

        with open(f"{Config.getTempPath()}/{video_info.get_id()}.info.json", "r") as f:
            return self._parseVideoInfo(video_info, f.read())

    def _downloadVideo(self, video_info: VideoInfo, output_dir: str, cookies_file_path: str = None):
        self.video_info = video_info
        print("download path",output_dir +'/%(title)s.%(ext)s')
        ydl_opts =  {
            'outtmpl': output_dir +'/%(title)s.%(ext)s',
            'progress_hooks': [self.progress_hook]
        }
        with yt_dlp.YoutubeDL(ydl_opts) as ydl:
            ydl.download([video_info.get_url()])

    def _downloadPoster(self, video_info: VideoInfo, output_dir: str):
        common.runShell(f"yt-dlp --skip-download --write-thumbnail -o {output_dir}/poster {video_info.get_url()}")

    def _downloadNfo(self, video_info: VideoInfo, output_dir: str):
        # why does this not directly write nfo to output_dir?
        # because this is a history problem, I will fix it in the future
        # in old version. the function is not download nfo. it generate nfo and write in other place.
        
        request.updateVideoStatus(video_info)
        os.system(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.info.json")
        os.system(f"ytdl-nfo {Config.getTempPath()}/{video_info.get_id()}.info.json")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.nfo")
        with open(f"{Config.getTempPath()}/{video_info.get_id()}.nfo", "r") as f:
            content = f.read()
            
            with open(f"{output_dir}/movie.nfo", "w") as f:
                f.write(content)


    def getVideoInfo(self,url)->List[VideoInfo]: 
        if self.isSupport(url):
            video_info = self._initVideoInfo(url)
            video_info_array = self._fetchVideoInfo(video_info)
            return video_info_array
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
        if self.isSupportWithVideoInfo(video_info):
            self._downloadNfo(video_info, output_dir)
        else:
            return self.next.downloadNfo(video_info, output_dir)

