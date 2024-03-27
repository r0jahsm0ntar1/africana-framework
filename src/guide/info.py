import sys
import time
from src.core.bcolors import *

class guide_inf(object):
    def __init__(self):
        pass

    def guide(self):
        with open('src/guide/guide.txt', 'r') as guide:
            for line in guide:
                sys.stdout.write(line)
                sys.stdout.flush()
                time.sleep(0.00003)

guide_info = guide_inf()
if ' __name__' == '__main__':
        sys.exit(guide_info())
