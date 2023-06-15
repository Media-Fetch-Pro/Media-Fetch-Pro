# What is it
`Media Fetch PRO` is an application to download video of youtube, bilibili and more websites in future  with metadata to jellyfin, emby, plex and more softwares in future.

![](./images/screen.png)

# Document

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
```
python3 script/main.py --url https://www.youtube.com/watch?v=lyNVPxHiVyE --storage ./video --type youtube
```

# How to use
```
docker run  -itd --name video-tool-for-nas -p 8080:8080 -v ./video:/var/opt/video correctroad/video-tools:latest
```

# Roadmap
- [] Support download Video Collection
- [] Support login by cookie to download hight resolution video
- [] Better Download manager
- [] i18n
- [] Write Documents
# Architecture
![](./images/arch.png)
