from dataclasses import dataclass,asdict
from typing import List
from typing import Union, Dict, Any, Optional
from script.utils.video import generate_uuid_from_url


@dataclass
class VideoInfo():
    id: str = None
    title: str = None
    author: str = None
    url: str = None
    content: str = None
    publish_time: int = None
    thumbnail: str = None # the field may didn't used
    tags: str = None
    
    status: str = None
    reason: str = None
    percent: int = None
    
    size: int = None
    already_download_size: int = None
    
    type: str = None
    source: str = None

    parent: str = None
    length: int = None
    episode: int = None
    children: List[str] = None
    
    start_download_time: int = None
    download_speed: int = None
    end_download_time: int = None
    
    def __init__(self): # how to differentiate between init and init video info in downloader ?🤔
        self.set_children([])
        self.set_percent(0)
        self.set_size(0)
        self.start_download_time = -1
        
    def get_id(self) -> Union[str,None]:
        return self.id
    
    def set_title(self,title:str) -> 'VideoInfo':
        self.title = title
        return self
    
    def get_title(self) -> Union[str,None]:
        return self.title

    def set_url(self,url:str) -> 'VideoInfo':
        self.url = url
        self.id = generate_uuid_from_url(url)
        return self
    
    def get_url(self) -> Union[str,None]:
        return self.url
    
    # status, percent, size, type, children, author, source, content, episode, parent, length get and set method
    def set_status(self,status:str) -> 'VideoInfo':
        self.status = status
        return self

    def get_status(self) -> Union[str,None]:
        return self.status
    
    def set_percent(self,percent:int) -> 'VideoInfo':
        self.percent = percent
        return self
    
    def get_percent(self) -> Union[int,None]:
        return self.percent
    
    def set_size(self,size:int) -> 'VideoInfo':
        self.size = size
        return self
    
    def get_size(self) -> Union[int,None]:
        return self.size
    
    def set_type(self,type:str) -> 'VideoInfo':
        self.type = type
        return self 
    
    def get_type(self) -> Union[str,None]:
        return self.type
    
    def set_children(self,children:List[str]) -> 'VideoInfo':
        self.children = children
        return self 
    
    def get_children(self) -> Union[List[str],None]:
        return self.children
    
    def set_author(self,author:str) -> 'VideoInfo':
        self.author = author
        return self 
    
    def get_author(self) -> Union[str,None]:
        return self.author
    
    def set_source(self,source:str) -> 'VideoInfo':
        self.source = source
        return self
    
    def get_source(self) -> Union[str,None]:
        return self.source
    
    def set_content(self,content:str) -> 'VideoInfo':
        self.content = content
        return self
    
    def get_content(self) -> Union[str,None]:
        return self.content
    
    def set_episode(self,episode:int) -> 'VideoInfo':
        self.episode = episode
        return self
    
    def get_episode(self) -> Union[int,None]:
        return self.episode
    
    def set_parent(self,parent:str) -> 'VideoInfo':
        self.parent = parent
        return self
    
    def get_parent(self) -> Union[int,None]:
        return self.parent
    
    def set_length(self,length:int) -> 'VideoInfo':
        self.length = length
        return self
    
    def get_length(self) -> Union[str,None]:
        return self.length
    
    def set_start_download_time(self,start_download_time:int) -> 'VideoInfo':
        self.start_download_time = start_download_time
        return self
    
    def get_start_download_time(self) -> Union[int,None]:
        return self.start_download_time
    
    def serialize(self) -> 'Dict':
        return asdict(self)

    def deserialize(self, data: Dict[str, Any]) -> 'VideoInfo':
        for key, value in data.items():
            # key to lower. Because these field in go is upper
            key = key.lower()
            setattr(self, key, value)
            
        
        # setattr(self, 'start_download_time', data['StartDownloadTime'])
        # set id from url
        self.set_url(data['url'])
        return self
    
    def toDict(self) -> 'Dict':
        return asdict(self)