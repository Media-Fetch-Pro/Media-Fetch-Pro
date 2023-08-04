# Media Fetch PRO
<img src="https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/images/logo.jpg?raw=true" height="173"/></a>

[Documentation](https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/documents/document.md) • Live Demo • Discuss in [Discord](https://discord.gg/wDB2u4y2uR)

English | [简体中文](./README.zh.md)

![](https://img.shields.io/github/stars/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)
![](https://img.shields.io/bitbucket/issues/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)

# What is Media Fetch PRO
`Media Fetch PRO` is an application to download video of youtube, bilibili and more websites in future  with metadata to jellyfin, emby, plex and more softwares in future.

![](./images/screen.png)

# Documentation

[Jellyfin](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#jellyfin)

[Emby](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#emby)

[Plex](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#plex)

# How to use
`./video` should be replace by your video folder(import in Jellyfin). for example `/home/user/Videos`.
```
docker run  -itd --name media-fetch-pro -p 7789:7789 -v ./video:/var/opt/video correctroad/media-fetch-pro:latest
```

# Roadmap
 - [x] Support download Youtube Playlist video
 - [ ] Support login by cookie to download hight resolution video
 - [ ] Better Download manager
 - [x] i18n
 - [ ] Write Documents

# Star history
[![Star History Chart](https://api.star-history.com/svg?repos=Media-Fetch-Pro/Media-Fetch-Pro&type=Date)](https://star-history.com/#Media-Fetch-Pro/Media-Fetch-Pro&Date)

# How to Develop
[Development Guide](./documents/development.md)
