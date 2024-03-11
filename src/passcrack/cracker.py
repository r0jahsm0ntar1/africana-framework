import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *

class pass_killer(object):
    def __init__(self):
        pass

    rc = '/usr/share/wordlists/rockyou.txt.gz'
    if os.path.exists(rc):
        os.system('gunzip /usr/share/wordlists/rockyou.txt.gz')
    else:
        pass

    def aircracking_password(self):
        print(bcolors.ENDC + "\n    -{" + bcolors.UNDERL + bcolors.BLUE + " Give full path for your .pcap and wordlist to be used " + bcolors.ENDC + "}-\n" + bcolors.ENDC)
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(Path = {0})".format(pcap) + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(wordlist = {0})".format(wordlist) + bcolors.ENDC)
        time.sleep(3)
        os.system('clear')
        process = subprocess.Popen("aircrack-ng {0} -w {1}".format(pcap, wordlist), shell = True).wait()
        return process

    def john_password(self):
        print(bcolors.ENDC + "\n    -{" + bcolors.UNDERL + bcolors.BLUE + " Give full path for your .pcap and wordlist to be used " + bcolors.ENDC + "}-\n" + bcolors.ENDC)
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(Path = {0})".format(pcap) + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(wordlist = {0})".format(wordlist) + bcolors.ENDC)
        time.sleep(3)
        os.system('clear')
        process = subprocess.Popen("john {0} --wordlist={1}".format(pcap, wordlist), shell = True).wait()
        return process

pass_cracker = pass_killer()
if ' __name__' == '__main__':
        sys.exit(pass_cracker())
