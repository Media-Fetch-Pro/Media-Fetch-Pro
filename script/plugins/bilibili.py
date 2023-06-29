import os
import yt_dlp
import script.api.request as request
from script.utils.video import generate_uuid_from_url
import script.utils.ytdlp as ytdlp
import script.utils.common as common
import json
import datetime

from script.plugins import baseDownloader
from script.model import videoInfo

from script.config.config import Config

# class BilibiliDownloader():
#     def __init__(self, url, output_dir,temp_path):
#         self.temp_path = temp_path
#         self.url = url
#         self.id = generate_uuid_from_url(url)
#         self.output_dir = output_dir
    
#     def progress_hook(self,d):
#         url = d['info_dict']['original_url']
#         title = d['info_dict']['title']
#         self.title = title # it will tiger multi times. So we need to optimize it.
#         request.updateVideoStatus(generate_uuid_from_url(url),url,title,d['status'],ytdlp.extract_progress(d['_percent_str']),1)

#     def downloadPoster(self):
#         temp_path = self.output_dir
#         os.system(f"yt-dlp --skip-download --write-thumbnail -o {temp_path}/poster {self.url}")

#     def getNfo(self):
#         request.updateVideoStatus(generate_uuid_from_url(self.url),self.url,"title is fetching","fetching meta",0,1)
#         os.system(f"yt-dlp --skip-download --write-info-json -o {self.temp_path}/{self.id} {self.url}")
#         common.waitFile(f"{self.temp_path}/{self.id}.info.json")
#         os.system(f"ytdl-nfo {self.temp_path}/{self.id}.info.json")
#         common.waitFile(f"{self.temp_path}/{self.id}.nfo")
#         with open(f"{self.temp_path}/{self.id}.nfo", "r") as f:
#             return f.read()

#     def removeTemp(self):
#         temp_path = self.output_dir + "/temp"
#         os.system(f"rm -rf {temp_path}")

#     def downloadVideo(self):
#         ydl_opts =  {
#             'outtmpl': self.output_dir +'/%(title)s.%(ext)s',
#             'progress_hooks': [self.progress_hook]
#         }
#         with yt_dlp.YoutubeDL(ydl_opts) as ydl:
#             ydl.download([self.url])


class Bilibili():
    def downloadVideo(self):
        pass
    
    def downloadNfo(self):
        pass
    
    def isSupport(self,url):
        if "bilibili" in url:
            return True
        else:
            return False
    
    def _initVideoInfo(self, url:str)->'videoInfo':
        video_info = videoInfo.VideoInfo()
        video_info.set_url(url)
        video_info.set_title("title is fetching")
        video_info.set_status("fetching")
        video_info.set_percent(0)
        
        
        # get now unix timestamp as start download time
        presentDate = datetime.datetime.now()
        unix_timestamp = datetime.datetime.timestamp(presentDate)*1000
        video_info.set_start_download_time(unix_timestamp)
        return video_info
    
    def _parseVideoInfo(self, video_info: videoInfo, json_text: str)->'videoInfo':
        video_json = json.loads(json_text)
        video_info.set_title(video_json["title"])

        if video_json['_type'] == 'playlist':
            # if the url is playlist
            video_info.set_length(int(video_json["playlist_count"]))
            return video_info
        else:
            # if the url is video
            video_info.set_author(video_json["uploader"])
            video_info.set_content(video_json["description"])
            video_info.set_source("bilibili")
            video_info.set_type("video")
            video_info.set_size(video_json["filesize_approx"])
            return video_info


    def _fetchVideoInfo(self, video_info: videoInfo)->'videoInfo':
        os.system(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.info.json")

        with open(f"{Config.getTempPath()}/{video_info.get_id()}.info.json", "r") as f:
            return self._parseVideoInfo(video_info, f.read())
        
    def getVideoInfo(self,url)->'videoInfo': # 这是一个责任链模式
        if self.isSupport(url):
            video_info = self._initVideoInfo(url)
            video_info = self._fetchVideoInfo(video_info)
            return video_info
        else:
            return self.next.getVideoInfo(url)

