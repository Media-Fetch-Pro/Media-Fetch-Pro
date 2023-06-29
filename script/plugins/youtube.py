import script.api.request as request
from script.utils.video import generate_uuid_from_url
import script.utils.ytdlp as ytdlp
import script.utils.common as common
import json
import datetime

from typing import List
from script.plugins.baseDownloader import BaseDownloader
from script.model.videoInfo import VideoInfo

from script.config.config import Config
from script.utils.video import generate_uuid_from_url

# class YoutubeDownloader():
#     def __init__(self, url, output_dir,temp_path):
#         self.temp_path = temp_path
#         self.id = generate_uuid_from_url(url)
#         self.url = url
#         self.output_dir = output_dir
        
#     def progress_hook(self,d):
#         # print(d["original_url"])
#         url = d['info_dict']['original_url']
#         title = d['info_dict']['title']
#         self.title = title # it will tiger multi times. So we need to optimize it.
#         request.updateVideoStatus(generate_uuid_from_url(url),url,title,d['status'],ytdlp.extract_progress(d['_percent_str']),1)

#     def getNfo(self):
#         request.updateVideoStatus(generate_uuid_from_url(self.url),self.url,"title is fetching","fetching meta",0,1)
#         os.system(f"yt-dlp --skip-download --write-info-json -o {self.temp_path}/{self.id} {self.url}")
#         common.waitFile(f"{self.temp_path}/{self.id}.info.json")
#         os.system(f"ytdl-nfo {self.temp_path}/{self.id}.info.json")
#         common.waitFile(f"{self.temp_path}/{self.id}.nfo")
#         with open(f"{self.temp_path}/{self.id}.nfo", "r") as f:
#             return f.read()
    
#     def downloadPoster(self):
#         temp_path = self.output_dir
#         os.system(f"yt-dlp --skip-download --write-thumbnail -o {temp_path}/poster {self.url}")
    
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


class Youtube(BaseDownloader):
    def downloadVideo(self):
        pass
    
    def downloadNfo(self):
        pass
    
    def isSupport(self,url):
        if "youtube" in url:
            return True
        else:
            return False
    
    def _initVideoInfo(self, url:str)->VideoInfo:
        video_info = VideoInfo()
        video_info.set_url(url)
        video_info.set_title("title is fetching")
        video_info.set_status("fetching")
        video_info.set_percent(0)
        video_info.set_source("youtube")

        
        # get now unix timestamp as start download time
        presentDate = datetime.datetime.now()
        unix_timestamp = datetime.datetime.timestamp(presentDate)*1000
        video_info.set_start_download_time(unix_timestamp)
        return video_info
    
    def _parseVideoInfo(self, video_info: VideoInfo, json_text: str)->List[VideoInfo]:
        video_json = json.loads(json_text)
        video_info.set_title(video_json["title"])
        p_video_array = []
        if video_json['_type'] == 'playlist':
            # if the url is playlist
            video_info.set_length(int(video_json["playlist_count"]))
            video_info.set_children(
                list(
                    map(
                        lambda p:generate_uuid_from_url(f"{video_info.url}?p={p}"),
                        range(1,video_info.get_length()+1)
                    )
                )
            )    
            # fetch every children video info
            for p in (1,video_info.get_length()):
                new_video_info = self.getVideoInfo(f"{video_info.url}?p={p}")[0]
                new_video_info.set_episode(p)
                new_video_info.set_parent(video_info.get_id())
                p_video_array.append(new_video_info)
                
                # copy info from children to parent
                # TODO: only execute once
                video_info.set_author(new_video_info.get_author())
                video_info.set_content(new_video_info.get_content())
            
            p_video_array.insert(0,video_info)
            return p_video_array
        
        else:
            # if the url is video
            video_info.set_author(video_json["uploader"])
            video_info.set_content(video_json["description"])
            video_info.set_type("video")
            video_info.set_size(video_json["filesize_approx"])
            return [video_info]


    def _fetchVideoInfo(self, video_info: VideoInfo)->List[VideoInfo]:
        common.runShell(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.info.json")

        with open(f"{Config.getTempPath()}/{video_info.get_id()}.info.json", "r") as f:
            return self._parseVideoInfo(video_info, f.read())
        
    def getVideoInfo(self,url)->List[VideoInfo]: # 这是一个责任链模式
        if self.isSupport(url):
            video_info = self._initVideoInfo(url)
            video_info_array = self._fetchVideoInfo(video_info)
            return video_info_array
        else:
            return self.next.getVideoInfo(url)

