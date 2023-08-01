from typing import List

from script.model.videoInfo import VideoInfo
import script.utils.common as common
from script.config.config import Config
import yt_dlp
import json
import os
from typing import List

import script.utils.common as common
import script.api.request as request

from script.model.videoInfo import VideoInfo
from script.utils.video import generate_uuid_from_url
from script.config.config import Config
from script.utils.video import generate_uuid_from_url
from script.utils.ytdlp import extract_progress

class BaseDownloader():
    def progress_hook(self,d):
        self.video_info.set_status('downloading')
        self.video_info.set_percent(extract_progress(d['_percent_str']))
        request.updateVideoStatus(self.video_info)

    def __init__(self, next):
        self.next = next
    
    def isSupport(self,url):
        raise Exception("isSupport didn't implement")
    
    def isSupportWithVideoInfo(self,video_info: VideoInfo) -> bool:
        raise Exception("isSupportWithVideoInfo didn't implement")

    def _initVideoInfo(self, url:str)->VideoInfo:
        raise Exception("Init Video Info didn't implement")
    
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
            video_info.set_type("playlist")
            
            # fetch every children video info
            for p in range(1,video_info.get_length()+1):
                new_video_info = self.getVideoInfo(f"{video_info.url}?p={p}")[0]
                new_video_info.set_episode(p)
                new_video_info.set_parent(video_info.get_id())
                new_video_info.set_type("episode")
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
        print(f"yt-dlp --skip-download --write-info-json -o {Config.getTempPath()}/{video_info.get_id()} {video_info.get_url()}")
        common.waitFile(f"{Config.getTempPath()}/{video_info.get_id()}.info.json")

        with open(f"{Config.getTempPath()}/{video_info.get_id()}.info.json", "r") as f:
            return self._parseVideoInfo(video_info, f.read())
    
    def _downloadVideo(self, video_info: VideoInfo, output_dir: str):
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
    
    def getVideoInfo(self,url)->List[VideoInfo]: # 这是一个责任链模式
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
