import pytest
from script.main import FetchVideoInfo, DownloadVideo, Rename
from script.model.videoInfo import VideoInfo
from script.utils.video import generate_uuid_from_url
from pathlib import Path

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


# def test_fetch_youtube_video_info(self):
#     assert 1

@pytest.mark.server(url='/api/update', response=[{'data': 'update video status'}], method='POST')
@pytest.mark.server_settings(port=7789)
def test_download_video(tmpdir):
    tmpdir_path = str(tmpdir)
    # single video
    video_info = VideoInfo() # TODO THE Video info and url should provide by fixtures.
    video_info.set_url("https://www.bilibili.com/video/BV1ZW411u7V3")
    video_info.set_type("video")
    video_info.set_source("bilibili")
    
    DownloadVideo(video_info,tmpdir_path)
    
    assert len(list( Path(tmpdir).iterdir() )) == 1
    ## there are 3 files in the folder: video, poster, nfo
    video_dir = tmpdir / generate_uuid_from_url("https://www.bilibili.com/video/BV1ZW411u7V3")
    assert len(list( Path(video_dir).iterdir() )) == 3
    
    parentId = generate_uuid_from_url("https://www.bilibili.com/video/BV1SW4y1v7WN")
    # playlist with 2 episode
    video_info = VideoInfo()
    video_info.set_type("playlist")
    video_info.set_length(2)
    video_info.set_url("https://www.bilibili.com/video/BV1SW4y1v7WN")
    video_info.set_parent(parentId)
    DownloadVideo(video_info,tmpdir_path)

    video_info = VideoInfo()
    video_info.set_type("episode")
    video_info.set_episode(1)
    video_info.set_url("https://www.bilibili.com/video/BV1SW4y1v7WN?p=1")
    video_info.set_parent(parentId)
    DownloadVideo(video_info,tmpdir_path)

    video_info = VideoInfo()
    video_info.set_type("episode")
    video_info.set_episode(2)
    video_info.set_url("https://www.bilibili.com/video/BV1SW4y1v7WN?p=2")
    video_info.set_parent(parentId)
    DownloadVideo(video_info,tmpdir_path)

    assert len(list( Path(tmpdir).iterdir() )) == 2
    ## there are 3 files in the folder: video，video2, poster, nfo
    video_dir = tmpdir / generate_uuid_from_url("https://www.bilibili.com/video/BV1SW4y1v7WN")
    assert len(list( Path(video_dir).iterdir() )) == 4

    # # playlist with 3 episode
    # video_info = VideoInfo()
    # DownloadVideo(video_info,fixture_storage)
