# Development Environment Setup
## backend:
```
air -c air/.air.toml
```

## frontend
```
cd web
pnpm i
pnpm dev
```

## python
dependencies
```shell
pip install -r requirements.txt
```

fetch video info
```python
python main.py --type fetchVideoInfo --url https://www.youtube.com/watch?v=lyNVPxHiVyE --storage ./video --website youtube
```

test 
```shell
pytest .
```

# Program Explanation
## Video Status
Unstart: a video to add the download queue
Fetching: the video is fetching the video information
Pending: the video is waiting for the download
Downloading: the video is downloading video、poster、nfo
Finished: the video is downloaded. waiting for the rename folder
Completed: the video folder is renamed.
Failed: the video is failed in above any step.