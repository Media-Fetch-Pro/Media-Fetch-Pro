import time
import os

def waitFile(filePath, timeout=2):
    '''
    Wait for file to appear
    '''
    print(f"Waiting for file {filePath}")
    start = time.time()
    while not os.path.exists(filePath):
        time.sleep(0.1)
        if time.time() - start > timeout:
            raise Exception('Timeout waiting for file %s' % filePath)