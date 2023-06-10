import requests


def updateVideoStatus(url,status,percent):
    requests.post('http://localhost:5000/api/video/updateStatus', data={'url': url, 'status': status, 'percent': percent})