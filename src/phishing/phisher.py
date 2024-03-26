import sys
import time
import subprocess
from src.core.bcolors import *

class phish_hack(object):
    def __init__(self):
        pass

    def phish_gophish(self):
        process = subprocess.Popen('gophish', shell = True).wait()
        time.sleep(0.03)
        return process

    def phish_goodginx(self):
        process = subprocess.Popen('evilginx2', shell = True).wait()
        time.sleep(0.03)
        return process

    def phish_setoolkit(self):
        process = os.system('cd ./src/externals/set/; python3 setoolkit')
        return process

    def phish_anonphisher(self):
        process = subprocess.Popen("cd ./src/externals/anonphisher; bash anonphisher.sh", shell = True).wait()
        return process

    def phish_zphisher(self):
        process = subprocess.Popen("cd ./src/externals/AdvPhishing; bash AdvPhishing.sh", shell = True).wait()
        time.sleep(0.03)
        return process

    def phish_cyberphish(self):
        process = subprocess.Popen("cd ./src/externals/cyberphish; python3 cyberphish.py", shell = True).wait()
        time.sleep(0.03)
        return process

cred_phisher = phish_hack()

if ' __name__' == '__main__':
        sys.exit(cred_phisher())
