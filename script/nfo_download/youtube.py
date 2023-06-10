import os

temp_path = "/home/ctrdh/video/temp/"

def readNfo(url):
    os.system(f"yt-dlp --skip-download --write-info-json -o {temp_path}temp {url}")
    os.system(f"ytdl-nfo {temp_path}temp.info.json")
    with open(f"{temp_path}temp.nfo", "r") as f:
        return f.read()
