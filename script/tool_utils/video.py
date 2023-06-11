import hashlib
import os
def generate_uuid_from_url(url):
    return hashlib.md5(url.encode("utf-8")).hexdigest()

def renameDir(dir_path, new_name):
    os.rename(dir_path, os.path.join(os.path.dirname(dir_path), new_name))