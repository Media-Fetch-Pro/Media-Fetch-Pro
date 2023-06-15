import os
import yt_dlp
import api.request as request
from tool_utils.video import generate_uuid_from_url
import tool_utils.ytdlp as ytdlp
import tool_utils.common as common
class BilibiliDownloader():
    def __init__(self, url, output_dir,temp_path):
        self.temp_path = temp_path
        self.url = url
        self.id = generate_uuid_from_url(url)
        self.output_dir = output_dir
    
    def progress_hook(self,d):
        url = d['info_dict']['original_url']
        title = d['info_dict']['title']
        self.title = title # it will tiger multi times. So we need to optimize it.
        request.updateVideoStatus(generate_uuid_from_url(url),url,title,d['status'],ytdlp.extract_progress(d['_percent_str']),1)

    def downloadPoster(self):
        temp_path = self.output_dir
        os.system(f"yt-dlp --skip-download --write-thumbnail -o {temp_path}/poster {self.url}")

    def getNfo(self):
        request.updateVideoStatus(generate_uuid_from_url(self.url),self.url,"title is fetching","fetching meta",0,1)
        os.system(f"yt-dlp --skip-download --write-info-json -o {self.temp_path}/{self.id} {self.url}")
        common.waitFile(f"{self.temp_path}/{self.id}.info.json")
        os.system(f"ytdl-nfo {self.temp_path}/{self.id}.info.json")
        common.waitFile(f"{self.temp_path}/{self.id}.nfo")
        with open(f"{self.temp_path}/{self.id}.nfo", "r") as f:
            return f.read()

    def removeTemp(self):
        temp_path = self.output_dir + "/temp"
        os.system(f"rm -rf {temp_path}")

    def downloadVideo(self):
        ydl_opts =  {
            'outtmpl': self.output_dir +'/%(title)s.%(ext)s',
            'progress_hooks': [self.progress_hook]
        }
        with yt_dlp.YoutubeDL(ydl_opts) as ydl:
            ydl.download([self.url])


