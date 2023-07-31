import pytest

from script.main import FetchVideoInfo, DownloadVideo, Rename
class TestVideoDownload:
    def test_fetch_bilibili_video_info(self):
        # to mock a http service
        
        
        fixture_storage = ""
        # single video
        video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV1ZW411u7V3")
        assert len(video_info_array) == 1
        
        # playlist
        video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV1SW4y1v7WN")
        assert len(video_info_array) == 3
        
        # playlist with 3 episode
        video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV15V411t7uW")
        assert len(video_info_array) == 4

    
    def test_fetch_youtube_video_info(self):
        assert 1

    def test_download_video(self):
        print("test_download_video")
        assert 1