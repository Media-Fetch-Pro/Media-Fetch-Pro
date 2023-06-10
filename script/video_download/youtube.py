import yt_dlp
import api.request as request
def progress_hook(d):
    request.updateVideoStatus(d['filename'],d['status'],d['_percent_str'])

def downloadVideo(url,dir):
    ydl_opts = {
        'outtmpl': dir +'%(title)s.mp4',
        'progress_hooks': [progress_hook]
    }
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        ydl.download([url])
