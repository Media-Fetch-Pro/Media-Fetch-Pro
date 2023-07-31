import argparse
import os
from script.model.videoInfo import VideoInfo
import json

from script.main import FetchVideoInfo, DownloadVideo, Rename

parser = argparse.ArgumentParser(description='Process some integers.')
parser.add_argument('--url', type=str,help='video url')
parser.add_argument('--type', type=str,help='do what') # fetching info, download video, download subtitle, download poster
parser.add_argument('--storage', type=str,help='storage path')
parser.add_argument('--video-info', type=str,help='video info') # it is a json string

temp_path = os.getcwd()+"/temp"

if not(os.path.isdir(temp_path)):
    print(temp_path)
    os.mkdir(temp_path)

args = parser.parse_args()

if __name__ == "__main__":
    if args.url == None:
        exit("url is None")
    
    if args.type == None:
        exit("type is None")
        
    if args.storage == None:
        exit("storage is None")

    print("action type:",args.type)
    print("url:",args.url)
    print("storage:",args.storage)

    if args.video_info != None:
        video_info = VideoInfo()
        try:
            video_info.deserialize(json.loads(args.video_info))
        except Exception as e:
            print(e)
            exit("video_info is not a json string")
            
        print("video info:",video_info)

    if args.type == "fetchVideoInfo":        
        FetchVideoInfo(args.url)

    elif args.type == "downloadVideo":
        DownloadVideo(video_info,args.storage)

    elif args.type == "generateNfo":
        pass
        # I think i didn't do this. generate nfo when download video.

    elif args.type == "rename":
        # does playlist have title?
        Rename(video_info,args.storage)

