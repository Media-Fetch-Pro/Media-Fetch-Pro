import re
def extract_progress(text):
    # convert \x1b[0;94m 70.9%\x1b[0m to 70
    # convert \x1b[0;94m100.0 to 100
    pattern = r'\x1b\[\d+(?:;\d+)*m'  # Matches escape sequences like \x1b[0;94m or \x1b[0m
    cleaned_text = re.sub(pattern, '', text)
    return int(float(cleaned_text.split('%')[0]))
    
if __name__ == '__main__':
    print(extract_progress('\x1b[0;94m 70.9%\x1b[0m'))
    print(extract_progress('\x1b[0;94m100.0'))