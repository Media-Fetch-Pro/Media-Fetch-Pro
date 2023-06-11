import os
import yt_dlp
import api.request as request
from tool_utils.video import generate_uuid_from_url
import tool_utils.ytdlp as ytdlp

class BilibiliDownloader():
    def __init__(self, url, output_dir):
        self.url = url
        self.output_dir = output_dir
    
    def progress_hook(d):
        url = d['info_dict']['original_url']
        request.updateVideoStatus(generate_uuid_from_url(url),url,d['status'],"title",ytdlp.extract_progress(d['_percent_str']),1)


    def readNfo(self):
        temp_path = "/home/ctrdh/video/temp/"
        os.system(f"yt-dlp --skip-download --write-info-json -o {temp_path}temp {self.url}")
        os.system(f"ytdl-nfo {temp_path}temp.info.json")
        with open(f"{temp_path}temp.nfo", "r") as f:
            return f.read()

    def downloadVideo(self):
        ydl_opts =  {
            'outtmpl': dir +'/%(title)s.%(ext)s',
            'progress_hooks': [self.progress_hook]
        }
        with yt_dlp.YoutubeDL(ydl_opts) as ydl:
            ydl.download([self.url])


