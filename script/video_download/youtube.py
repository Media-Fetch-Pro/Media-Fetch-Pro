import yt_dlp
import api.request as request
import tool_utils.ytdlp as ytdlp

def progress_hook(d):
    # print(d["original_url"])
    request.updateVideoStatus(d['info_dict']['original_url'],d['status'],ytdlp.extract_progress(d['_percent_str']))

def downloadVideo(url,dir):
    ydl_opts = {
        'outtmpl': dir +'/%(title)s.%(ext)s',
        'progress_hooks': [progress_hook]
    }
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        ydl.download([url])
