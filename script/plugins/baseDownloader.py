class BaseDownloader():
    def __init__(self, next):
        self.next = next
    
    def downloadVideo(self):
        pass
    
    def downloadNfo(self):
        pass
    
    def isSupport(self,url):
        pass
    
    def getVideoInfo(self,url): # 这是一个责任链模式，如果
        pass
    