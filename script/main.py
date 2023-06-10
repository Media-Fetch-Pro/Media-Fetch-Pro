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
import nfo_download.youtube as youtube_nfo
import video_download.youtube as youtube_video
import os

temp_path = "/home/ctrdh/video/temp"


parser = argparse.ArgumentParser(description='姓名')
parser.add_argument('--url', type=str,help='视频链接')
parser.add_argument('--storage', type=str,help='地址')
parser.add_argument('--type', type=str,help='判断')

args = parser.parse_args()

if __name__ == "__main__":
    # 判断下载路径是否是一个目录
    if os.path.isdir(args.storage):
        if(args.type == "youtube"):
            with open(f"{args.storage}/temp.nfo","w") as f:
                f.write(youtube_nfo.readNfo(args.url))
            youtube_video.downloadVideo(args.url,args.storage)
