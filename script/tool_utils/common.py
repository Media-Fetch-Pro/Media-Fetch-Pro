import time
import os

def waitFile(filePath, timeout=60):
    '''
    Wait for file to appear
    '''
    start = time.time()
    while not os.path.exists(filePath):
        time.sleep(0.1)
        if time.time() - start > timeout:
            raise Exception('Timeout waiting for file %s' % filePath)