import requests


def updateVideoStatus(url,status,percent):
    data = {"url": url, "status": status, "percent": percent}
    res = requests.post('http://127.0.0.1:8080/update', json=data)   
