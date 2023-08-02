import pytest
from script.main import FetchVideoInfo

@pytest.mark.server(url='/api/update', response=[{'data': 'update video status'}], method='POST')
@pytest.mark.server_settings(port=7789)
def test_fetch_bilibili_video_info():        
    # single video
    video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV1ZW411u7V3")
    assert len(video_info_array) == 1
    
    # playlist
    video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV1SW4y1v7WN")
    assert len(video_info_array) == 3
    
    # playlist with 3 episode
    video_info_array = FetchVideoInfo("https://www.bilibili.com/video/BV15V411t7uW")
    assert len(video_info_array) == 4
