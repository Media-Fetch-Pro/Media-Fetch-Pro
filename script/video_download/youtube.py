import yt_dlp
import api.request as request
from tool_utils.video import generate_uuid_from_url
import tool_utils.ytdlp as ytdlp

def progress_hook(d):
    # print(d["original_url"])
    url = d['info_dict']['original_url']
    request.updateVideoStatus(generate_uuid_from_url(url),url,d['status'],"title",ytdlp.extract_progress(d['_percent_str']),1)

def downloadVideo(url,dir):
    ydl_opts = {
        'outtmpl': dir +'/%(title)s.%(ext)s',
        'progress_hooks': [progress_hook]
    }
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        ydl.download([url])
