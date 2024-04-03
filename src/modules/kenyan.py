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
        os.system('cd /usr/local/opt/africana-framework; script -q -c "python3 africana.py" %s%s%s%s' %(africs, "africana-", time.strftime("%Y%m%d-%H%M%S"), ".log"))
main = africana()

if ' __name__' == '__main__':
        sys.exit(main())
