## Media Fetch PRO
![Media Fetch PRO Logo](https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/images/logo.jpg?raw=true)

[Documentation](https://github.com/CorrectRoadH/Media-Fetch-Pro/blob/main/documents/document.md) • Live-Demo • Diskussion auf [Discord](https://discord.gg/wDB2u4y2uR)

[Englisch](./README.md) | [简体中文](./README.zh.md) | German

![GitHub Sterne](https://img.shields.io/github/stars/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)
![Bitbucket Probleme](https://img.shields.io/bitbucket/issues/Media-Fetch-Pro/Media-Fetch-Pro?style=for-the-badge)

# Was ist Media Fetch PRO
`Media Fetch PRO` ist eine Anwendung zum Herunterladen von Videos von YouTube, Bilibili und zukünftig weiteren Websites, mit Metadaten für Jellyfin, Emby, Plex und zukünftig weiteren Softwares.

![Screenshot von Media Fetch PRO](./images/screen.png)

# Dokumentation

[Jellyfin](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#jellyfin)

[Emby](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#emby)

[Plex](https://github.com/Media-Fetch-Pro/Media-Fetch-Pro/blob/main/documents/document.md#plex)

# Nutzung
`./video` sollte durch Ihren Videoordner (Import in Jellyfin) ersetzt werden, zum Beispiel `/home/user/Videos`.
```
docker run  -itd --name media-fetch-pro -p 7789:7789 -v ./video:/var/opt/video correctroad/media-fetch-pro:latest
```

# Fahrplan
 - [x] Unterstützung für das Herunterladen von YouTube-Playlisten
 - [ ] Unterstützung für das Einloggen mit Cookies, um Videos in hoher Auflösung herunterzuladen
 - [ ] Besserer Download-Manager
 - [x] i18n
 - [ ] Dokumente schreiben

# Sternenhistorie
[![Sternenverlauf Diagramm](https://api.star-history.com/svg?repos=Media-Fetch-Pro/Media-Fetch-Pro&type=Date)](https://star-history.com/#Media-Fetch-Pro/Media-Fetch-Pro&Date)

# Entwicklung
[Entwicklungsleitfaden](./documents/development.md)
