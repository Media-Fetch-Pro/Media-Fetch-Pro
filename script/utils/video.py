import hashlib
import os
def generate_uuid_from_url(url):
    return hashlib.md5(url.encode("utf-8")).hexdigest()

def renameDir(dir_path, new_name):     
    print("rename",dir_path, new_name)   
    try:
        os.rename(dir_path, os.path.join(os.path.dirname(dir_path), new_name))
    except Exception as e:
        print(e)
        raise Exception("This is a new error")      


def get_clear_video_url(url:str)->str:
    pass