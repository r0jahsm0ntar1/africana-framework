import sys
import time
import subprocess
from src.core.bcolors import *

class phish_hack(object):
    def __init__(self):
        pass

    def phish_gophish(self):
        process = subprocess.Popen("gophish", shell = True).wait()
        time.sleep(0.03)
        return process

    def phish_goodginx(self):
        process = subprocess.Popen("evilginx2", shell = True).wait()
        time.sleep(0.03)
        return process

    def phish_setoolkit(self):
        process = subprocess.Popen("setoolkit", shell = True).wait()
        return process
    def phish_anonphisher(self):
        process = subprocess.Popen("cd /usr/local/opt/africana-framework/afric/anonphisher; bash anonphisher.sh", shell = True).wait()
        return process

    def phish_zphisher(self):
        process = subprocess.Popen("cd /usr/local/opt/africana-framework/afric/AdvPhishing; bash AdvPhishing.sh", shell = True).wait()
        time.sleep(0.03)
        return process

cred_phisher = phish_hack()
if ' __name__' == '__main__':
        sys.exit(cred_phisher())
