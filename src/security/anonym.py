import sys
import subprocess
from src.core.banner import *
from src.core.bcolors import *
from src.core.spinner import *
from src.modules.secmon import *
from src.configs.config import *
from src.scriptures.verses import *

class anonym(object):
    def __init__(self):
        pass

    def vanish_install(self):
        os.system('clear')
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n              ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Installing & Configuring" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[             Tor (Install tor & set proxies)             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[        Iptables (Install Iptables for firewalls)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[        Squid (Install Squid set through Privoxy)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[       Privoxy (Install Privoxy & set through tor)       ] \n" + bcolors.ENDC)
        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing tor, privoxy, squid, dnsmasq & setting things." + bcolors.ENDC)
        spinner.run_subprocess('apt-get update; apt-get install -y tor squid privoxy iptables isc-dhcp-client isc-dhcp-server'), config.configure_all()
        spinner.run_subprocess('systemctl daemon-reload; systemctl enable tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service; systemctl restart tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service')
        print("")

    def vanish_start(self):
        os.system('clear')
        subprocess.Popen(['cd src/externals/tor-vanish/; python3 vanisher.py -m'], shell = True).wait()
        config.configure_all()

    def vanish_stop(self):
        os.system('clear')
        subprocess.Popen('cd src/externals/tor-vanish/; python3 vanisher.py -e', shell = True).wait()

    def checktor_status(self):
        os.system('clear')
        subprocess.Popen('cd src/externals/tor-vanish/; python3 vanisher.py -w', shell = True).wait()

    def chains_start(self):
        os.system('clear')
        sec_mon.pproxy()

anonymous = anonym()

if ' __name__' == '__main__':
        sys.exit(anonymous())
