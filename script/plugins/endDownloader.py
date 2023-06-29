from script.plugins.baseDownloader import BaseDownloader

class EndDownloader(BaseDownloader):
    def __init__(self):
        pass
    
    def downloadVideo(self):
        pass
    
    def downloadNfo(self):
        pass

    def getVideoInfo(self,url): # 这是一个责任链模式
        raise NotImplementedError("the website is not supported")