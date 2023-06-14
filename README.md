# What is it
it is an application to download youtube and bilibili video with metadata to jellyfin,emby and plex.

# Architecture
![](./images/arch.png)

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