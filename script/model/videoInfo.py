from dataclasses import dataclass
from dataclasses_json import dataclass_json
from typing import List
from typing import Union, Dict, Any, Optional
from script.utils.video import generate_uuid_from_url


@dataclass_json
@dataclass
class VideoInfo():
    id: str = None
    title: str = None
    url: str = None
    status: str = None
    percent: int = None
    size: int = None
    type: str = None
    children: List[str] = None
    author: str = None
    source: str = None
    content: str = None
    episode: str = None
    parent: str = None
    length: str = None
    start_download_time: int = None
    
    # id should generate from md5 url. So we don't need set_id method
    # def set_id(self,id:str) -> 'VideoInfo':
    #     self.id = id
    #     return self
    
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
    
    def set_episode(self,episode:str) -> 'VideoInfo':
        self.episode = episode
        return self
    
    def get_episode(self) -> Union[str,None]:
        return self.episode
    
    def set_parent(self,parent:str) -> 'VideoInfo':
        self.parent = parent
        return self
    
    def get_parent(self) -> Union[str,None]:
        return self.parent
    
    def set_length(self,length:str) -> 'VideoInfo':
        self.length = length
        return self
    
    def get_length(self) -> Union[str,None]:
        return self.length
    
    def set_start_download_time(self,start_download_time:int) -> 'VideoInfo':
        self.start_download_time = start_download_time
        return self
    
    def get_start_download_time(self) -> Union[int,None]:
        return self.start_download_time