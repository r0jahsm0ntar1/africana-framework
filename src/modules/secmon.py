import subprocess
from src.core.bcolors import *

class pproxy_mon(object):
    def __init__(self):
        pass

    def pproxy(self):
        print(bcolors.RED + "\n                     Part of africana-framework" + bcolors.ENDC)
        print(bcolors.BLUE + "               Your internet proxy connections routes\n" + bcolors.ENDC)
        print(bcolors.BLUE + "  Running command " + bcolors.RED + "-> " + bcolors.YELLOW + "tail -vf /var/log/privoxy/logfile " + bcolors.BLUE + "To see Your Logs" + bcolors.ENDC)
        print(bcolors.BLUE + " Launch attack using port 3129 " + bcolors.RED + "ex. " + bcolors.GREEN + "sqlmap --proxy=http://127.0.0.1:3129" + bcolors.ENDC)
        print(bcolors.BLUE + "\n                     -[ Your Proxy Chains ]-\n" + bcolors.ENDC)
        print(bcolors.GREEN + "  ( Local " + bcolors.RED + "> " + bcolors.GREEN + "Squid " + bcolors.YELLOW + "3129 " + bcolors.RED + "> " + bcolors.GREEN + "Privoxy " + bcolors.RED + "> " + bcolors.YELLOW + "8118 " + bcolors.RED + "> " + bcolors.GREEN + "Tor " + bcolors.YELLOW + "9050 " + bcolors.RED + "> " + bcolors.GREEN + "web )\n" + bcolors.ENDC)
        subprocess.Popen('tail -vf /var/log/privoxy/logfile', shell = True).wait()
sec_mon = pproxy_mon()

if ' __name__' == '__main__':
        sys.exit(pproxy())
