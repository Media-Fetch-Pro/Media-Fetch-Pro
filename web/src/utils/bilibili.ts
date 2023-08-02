const getCleanUrl = (url: string):string => {
    const reg = /(?<=\?).*$/g
    const match = url.match(reg)
    if (match) {
        return url.replace(match[0], '')
    }
    return url
}

export {getCleanUrl}