#!/usr/bin/env python3
import os
import sys

class africana(object):
      def __init__(self):
        os.chdir('/usr/local/opt/africana-framework/')
        os.system('python3 africana.py')
main = africana()

if ' __name__' == '__main__':
        sys.exit(main())
