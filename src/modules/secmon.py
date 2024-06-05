import subprocess
from src.core.bcolors import *

class pproxy_mon(object):
    def __init__(self):
        pass

    def pproxy(self):
        print(bcolors.BLUE + "\n  ~>( " + bcolors.ENDC + bcolors.UNDERL + "Monitoring traffic through squid, privoxy, tor" + bcolors.ENDC + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[.........Your internet proxy connections routes.........] " + bcolors.ENDC)
        print(bcolors.BLUE + "[...................Running command......................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[...Launch ex. sqlmap --proxy=http://127.0.0.1:3129......] " + bcolors.ENDC)
        print(bcolors.BLUE + "[...............Your Proxy Chains Route..................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[Local " + bcolors.RED + "~>" + bcolors.GREEN + "Squid " + bcolors.RED + "~>" + bcolors.YELLOW + "3129 " + bcolors.RED + "~>" + bcolors.GREEN + "Privoxy " + bcolors.RED + "~>" + bcolors.YELLOW + "8118 " + bcolors.RED + "~>" + bcolors.GREEN + "Tor " + bcolors.RED + "~>" + bcolors.YELLOW + "9050 " + bcolors.RED + "~>" + bcolors.BLUE + "web]\n" + bcolors.ENDC)
        subprocess.Popen('tail -vf /var/log/privoxy/logfile', shell = True).wait()
sec_mon = pproxy_mon()

if ' __name__' == '__main__':
        sys.exit(pproxy())
