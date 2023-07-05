import argparse
import os
from script.plugins.bilibili import Bilibili
from script.plugins.youtube import Youtube
from script.plugins.endDownloader import EndDownloader
from script.utils.video import generate_uuid_from_url
from script.utils.video import renameDir
from script.model.videoInfo import VideoInfo
import json

parser = argparse.ArgumentParser(description='Process some integers.')
parser.add_argument('--url', type=str,help='video url')
parser.add_argument('--type', type=str,help='do what') # fetching info, download video, download subtitle, download poster
parser.add_argument('--storage', type=str,help='storage path')
parser.add_argument('--website', type=str,help='video website') # this option is not necessary. I will delete it in future.
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
        
    if args.video_info != None:
        video_info = VideoInfo()
        try:
            video_info.deserialize(json.loads(args.video_info))
        except Exception as e:
            print(e)
            exit("video_info is not a json string")
    
    if args.type == "fetchVideoInfo":
        # 我觉得这里做个责任链模式比较好，一个个传下去，谁能解析就谁来解析
        websites = Bilibili(Youtube(EndDownloader()))
        # to print result for debug
        print(json.dumps(list(map(lambda x:x.serialize(),websites.getVideoInfo(args.url))),indent=4, separators=(',', ': ')))
            
    elif args.type == "downloadVideo":
        # 判断下载路径是否是一个目录

        print("download "+args.storage)
        try:
            args.storage = args.storage + "/" + generate_uuid_from_url(args.url)
            if not(os.path.isdir(args.storage)):
                os.mkdir(args.storage)
        except Exception as e:
            print(e)
        
        print("storage path is "+args.storage)

        # TODO how to process when  video belong a playlist? 🤔
        websites = Bilibili(Youtube(EndDownloader()))

        # renameDir(f"{args.storage}",f"{video_info.get_title()}")
        if video_info.get_type() == "playlist":
            # TODO generate nfo
            # this is generate a tvshow.nfo🤔 it is very very hard.
            pass
        
        elif video_info.type == "video":

            if video_info.get_type() == "video": # episode didn't generate nfo
                websites.downloadNfo(video_info,args.storage)

            websites.downloadPoster(video_info,args.storage)
            websites.downloadVideo(video_info,args.storage)
            
            if video_info.get_type() == "video": 
                # TODO it is a problem how to rename playlist🤔
                renameDir(f"{args.storage}",f"{video_info.get_title()}")        

    elif args.type == "generateNfo":
        pass
        # I think i didn't do this. generate nfo when download video.
