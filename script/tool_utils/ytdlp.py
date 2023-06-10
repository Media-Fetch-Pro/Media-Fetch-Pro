import re
def extract_progress(progress_string):
    # convert \x1b[0;94m 70.9%\x1b[0m to 70
    # convert \x1b[0;94m100.0 to 100
    # if progress_string == '':
    #     return 0
    # if progress_string.index('\x1b[0;94m100.0') != -1:
    #     return 100
    # progress = progress_string.split('%')[0]
    # progress = progress.split(' ')[-1]
    # progress = progress.strip()
    # return int(float(progress))
    progress_match = re.search(r'\b(\d+\.?\d*)', progress_string)
    if progress_match:
        return int(float(progress_match.group(0)))
    
if __name__ == '__main__':
    print(extract_progress('\x1b[0;94m 70.9%\x1b[0m'))
    print(extract_progress('\x1b[0;94m100.0'))