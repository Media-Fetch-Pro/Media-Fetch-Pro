const getCleanUrl = (url: string):string => {
    // https://www.bilibili.com/video/BV1cX4y177WM/?spm_id_from=333.1007.tianma.1-2-2.click&vd_source=5d78b061956227324f710505619f52d7 => https://www.bilibili.com/video/BV1cX4y177WM
    const reg = /(.*)\?/;
    const match = url.match(reg);
    if (match) {
        return match[1];
    }
    return url;
}

export {getCleanUrl}