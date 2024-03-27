import sys
import time
from src.core.bcolors import *

class word(object):
    def __init__(self):
        pass

    def gospel(self):
        while True:
            try:
                with open('src/scriptures/kjv.txt', 'r') as verses:
                    words = color() + verses.read()
                    for w in words:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
            except:
                neo.one()
                break

salvation = word()

if ' __name__' == '__main__':
        sys.exit(salvation())
