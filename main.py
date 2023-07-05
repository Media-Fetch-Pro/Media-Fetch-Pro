# 接受参数-u和-s,然后输出开始下载
# -u: 指定下载的url
# -s: 指定下载的文件大小
# -t: 指定下载的线程数
# -f: 指定下载的文件名
# -d: 指定下载的目录
# -h: 帮助信息
# -v: 版本信息
# -p: 指定下载的端口号
# -l: 指定下载的日志文件
# -c: 指定下载的配置文件
# -r: 指定下载的重试次数
# -i: 指定下载的时间间隔
# -a: 指定下载的代理地址
# -b: 指定下载的代理端口号
# -x: 指定下载的代理用户名
# -y: 指定下载的代理密码
# -z: 指定下载的代理类型
# -k: 指定下载的cookie文件
# -g: 指定下载的referer
# -e: 指定下载的user-agent
# -o: 指定下载的超时时间
# -m: 指定下载的最大速度
# -n: 指定下载的最小速度
# -w: 指定下载的最大等待时间
# -q: 指定下载的最小等待时间
# -j: 指定下载的最大重试时间
# -k: 指定下载的最小重试时间
# -b: 指定下载的最大重试次数

import argparse
import os
from script.plugins.bilibili import Bilibili
from script.plugins.youtube import Youtube
from script.plugins.endDownloader import EndDownloader
from script.utils.video import generate_uuid_from_url
from script.utils.video import renameDir
from script.model.videoInfo import VideoInfo
import json

parser = argparse.ArgumentParser(description='姓名')
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
        # websites = Bilibili(Youtube(EndDownloader()))
        
        # video_info = websites.getVideoInfo(args.url)[0]
        
        # # the video_info should not is a playlist.
        # websites.downloadVideo(video_info,args.storage)
        # # TODO how to process when  video belong a playlist? 🤔
        # websites.downloadPoster(video_info,args.storage)
        websites = Bilibili(Youtube(EndDownloader()))

        # renameDir(f"{args.storage}",f"{video_info.get_title()}")
        if video_info.get_type() == "playlist":
            # TODO generate nfo
            pass
        elif video_info.type == "video":
            print("download video")

            if video_info.get_type() == "video": # episode didn't generate nfo
                # TODO generate nfo
                pass            
            websites.downloadPoster(video_info,args.storage)
            websites.downloadVideo(video_info,args.storage)
            
            if video_info.get_type() == "video": # TODO how to rename playlist is a problem
                renameDir(f"{args.storage}",f"{video_info.get_title()}")            
    elif args.type == "generateNfo":
        pass
