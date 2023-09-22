import pytest
from script.main import DownloadVideo
from script.model.videoInfo import VideoInfo
from script.utils.video import generate_uuid_from_url
from pathlib import Path

# @pytest.mark.server(url='/api/update', response=[{'data': 'update video status'}], method='POST')
# @pytest.mark.server_settings(port=7789)
def test_download_video(tmpdir):
    tmpdir_path = str(tmpdir)
    # single video
    video_info = VideoInfo() # TODO THE Video info and url should provide by fixtures.
    video_info.set_url("https://www.youtube.com/watch?v=4OygeexwWe0")
    video_info.set_type("video")
    video_info.set_source("youtube")
    
    DownloadVideo(video_info,tmpdir_path)
    
    assert len(list( Path(tmpdir).iterdir() )) == 1
    ## there are 3 files in the folder: video, poster, nfo
    video_dir = tmpdir / generate_uuid_from_url("https://www.youtube.com/watch?v=4OygeexwWe0")
    assert len(list( Path(video_dir).iterdir() )) == 3
    
