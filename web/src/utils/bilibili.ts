const getCleanUrl = (url: string):string => {
    // https://www.bilibili.com/video/BV1cX4y177WM/?spm_id_from=333.1007.tianma.1-2-2.click&vd_source=5d78b061956227324f710505619f52d7 => https://www.bilibili.com/video/BV1cX4y177WM
    // https://www.youtube.com/watch?v=IYYUvfSBXD0 => https://www.youtube.com/watch?v=IYYUvfSBXD0
    // https://www.youtube.com/watch?v=4OygeexwWe0&list=PL-xubgenClPjBv1705B8XWwdTWJUGcFtn => https://www.youtube.com/watch?v=4OygeexwWe0&list=PL-xubgenClPjBv1705B8XWwdTWJUGcFtn
    const bilibiliRegex = /https:\/\/www\.bilibili\.com\/video\/[A-Za-z0-9]+\/?/;
    const youtubeRegex = /https:\/\/www\.youtube\.com\/watch\?v=[A-Za-z0-9]+/;
    if (bilibiliRegex.test(url)){
        return bilibiliRegex.exec(url)![0];
    }else if (youtubeRegex.test(url)){
        return url;
    }else{
        return url;
    }
}

export {getCleanUrl}