import os
def downloadVideo(url,dir):
    os.system(f"yt-dlp -o \"{dir}/%(title)s.mp4\" {url} &")
