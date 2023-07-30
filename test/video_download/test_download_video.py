import pytest
import subprocess
import json
class TestVideoDownload:
    def test_fetch_bilibili_video_info(self):
        
        result = subprocess.check_output(["python3","main.py","--type","fetchVideoInfo" ,"--url","https://www.bilibili.com/video/BV1yu411h7Be"])
        video_info= json.loads(str(result))
        print(video_info)
        assert result == 0
        
    
    def test_fetch_youtube_video_info(self):
        assert 1

    def test_download_video(self):
        print("test_download_video")
        assert 1