# Media Fetch PRO
<img src="https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/images/logo.jpg?raw=true" height="173"/></a>

[Documentation](https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/documents/document.md) • Live Demo • Discuss in [Discord](https://discord.gg/2WDm9uXZ)

![](https://img.shields.io/github/stars/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)
![](https://img.shields.io/bitbucket/issues/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)

# What is Media Fetch PRO
`Media Fetch PRO` is an application to download video of youtube, bilibili and more websites in future  with metadata to jellyfin, emby, plex and more softwares in future.

![](./images/screen.png)

# Documentation

Jellyfin

Emby

Plex

# How to Develop
backend:
```
air -c air/.air.toml
```

frontend
```
cd web
pnpm i
pnpm dev
```

python
dependencies
```shell
pip install WIP
```

fetch video info
```python
python main.py --type fetchVideoInfo --url https://www.youtube.com/watch?v=lyNVPxHiVyE --storage ./video --website youtube
```

test 
```shell
pytest .
```

# How to use
`./video` should be replace by your video folder(import in Jellyfin). for example `/home/user/Videos`.
```
docker run  -itd --name media-fetch-pro -p 7789:7789 -v ./video:/var/opt/video correctroad/media-fetch-pro:latest
```

# Roadmap
 - [x] Support download Playlist video
 - [ ] Support login by cookie to download hight resolution video
 - [ ] Better Download manager
 - [ ] i18n
 - [ ] Write Documents

# Star history
[![Star History Chart](https://api.star-history.com/svg?repos=Media-Fetch-Pro/Media-Fetch-Pro&type=Date)](https://star-history.com/#Media-Fetch-Pro/Media-Fetch-Pro&Date)
