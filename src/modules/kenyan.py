#!/usr/bin/env python3
import os
import sys
import time

class africana(object):
    def __init__(self):
        africs = '/root/.africana/'
        if os.path.exists(africs):
            pass
        else:
            os.mkdir('/root/.africana')
        os.system('cd /usr/local/opt/africana-framework; python3 africana.py')
main = africana()

if ' __name__' == '__main__':
        sys.exit(main())
