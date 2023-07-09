import requests
from script.model.videoInfo import VideoInfo

def updateVideoStatus(video: VideoInfo):
    data = video.toDict()
    print(data)
    try:
        res = requests.post('http://127.0.0.1:7789/api/update', json=data) 
    except:
        print("updateVideoStatus error")
        return False  
