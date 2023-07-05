from script.plugins.baseDownloader import BaseDownloader

# this downloader is for download bilibili playlist nfo. not for video🤔
class BilibiliPlayList(BaseDownloader):
    def isSupport(self,url):
        # 🤔 how to judge url is a playlist. I think it very very hard.
        return False
