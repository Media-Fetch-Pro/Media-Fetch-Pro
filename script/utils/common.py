import time
import os
import subprocess


def waitFile(filePath, timeout=2):
    '''
    Wait for file to appear
    '''
    start = time.time()
    while not os.path.exists(filePath):
        time.sleep(0.1)
        if time.time() - start > timeout:
            raise Exception('Timeout waiting for file %s' % filePath)
        
def runShell(shell_text):
    proc = subprocess.Popen(shell_text, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, bufsize=-1)
    proc.wait()
