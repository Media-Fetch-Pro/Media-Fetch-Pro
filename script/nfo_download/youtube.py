import os
def readNfo(url):
    os.system(f"yt-dlp --skip-download --write-info-json -o temp {url}")
    os.system(f"ytdl-nfo temp.info.json")
    with open("temp.nfo", "r") as f:
        return f.read()
