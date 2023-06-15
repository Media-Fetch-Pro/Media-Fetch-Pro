import requests


def updateVideoStatus(id,url,title,status,percent,size):
    data = {
        "id": id,
        "url": url, 
        "title": title,
        "status": status, 
        "percent": percent,
        "size": size
    }
    res = requests.post('http://127.0.0.1:7789/api/update', json=data)   
