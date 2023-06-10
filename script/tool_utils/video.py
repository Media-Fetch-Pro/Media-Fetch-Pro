import hashlib
def generate_uuid_from_url(url):
    return hashlib.md5(url.encode("utf-8")).hexdigest()