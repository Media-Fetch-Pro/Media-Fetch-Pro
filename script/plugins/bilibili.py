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

class Bilibili(BaseDownloader):
    def progress_hook(self,d):
        self.video_info.set_status('downloading')
        self.video_info.set_percent(extract_progress(d['_percent_str']))
        request.updateVideoStatus(self.video_info)

    def isSupport(self,url):
        if "bilibili" in url:
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
        video_info.set_source("bilibili")

        # get now unix timestamp as start download time
        presentDate = datetime.datetime.now()
        unix_timestamp = datetime.datetime.timestamp(presentDate)*1000
        video_info.set_start_download_time(int(unix_timestamp))
        return video_info
    
    