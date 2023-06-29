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
    try:
        res = requests.post('http://127.0.0.1:7789/api/update', json=data) 
    except:
        print("updateVideoStatus error")
        return False  
