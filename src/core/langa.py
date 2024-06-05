import re
import os
import sys
import time
import subprocess
from signal import *

from urllib.parse import *
from src.core.menu import *
from src.guide.info import *
from src.core.system import *
from src.core.banner import *
from src.core.musker import *
from src.core.bcolors import *
from src.wireles.wifi import *
from src.modules.secmon import *
from src.security.anonym import *
from src.kitchen.malware import *
from src.phishing.phisher import *
from src.internal.scanner import *
from src.passcrack.cracker import *
from src.webattack.scanner import *
from src.scriptures.verses import *
from src.scriptures.salvation import *

class neo_start(object):
    def __init__(self):
        pass
    def sudo():
        if not os.geteuid() == 0:
            os.system('clear')
            beauty.graphics(), scriptures.verses()
            print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "The africana-framwork needs super User. Try " + bcolors.RED + "sudo africana" + bcolors.GREEN + ")" + bcolors.ENDC)
            sys.exit(0)
    sudo()
#0
    def user_agree(self):
        while True:
            try:
                os.system('clear')
                if not os.path.isfile("src/agreement/covenant.txt"):
                    with open("readme/LICENSE") as fileopen:
                        for line in fileopen:
                            print((bcolors.ENDC + line.rstrip()))
                    if os.path.exists('/root/.africana/'):
                        pass
                    elif not os.path.exists('/root/.africana/'):
                        os.system('mkdir -p /root/.africana/')
                    mega_menu.menu_zero()
                    covenant = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "Do you agree to the terms of service [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
                    covenant += " "
                    if covenant[0].lower() == "y":
                        with open("src/agreement/covenant.txt", "w") as filewrite:
                            filewrite.write("user accepted")
                        neo.one()
                        break
                    else:
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        break
                else:
                    neo.one()
                    break
            except:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                break

#1
    def install_africana(self):
        while True:
            try:
                return installer.update_system(); neo.one()
            except:
                neo.one()
                break

#2
    def vanish_tor(self):
        while True:
            try:
                os.system('clear')
                def anonymity():
                    while True:
                        try:
                            mega_menu.menu_two()
                            choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        except:
                            neo.one()
                            break
                        try:
                            if choice == '0':
                                neo.one()
                                break
                            elif choice == '1':
                                try:
                                    return anonymous.vanish_install(), anonymity()
                                except:
                                    neo.vanish_tor()
                                    break
                            elif choice == '2':
                                try:
                                    return anonymous.vanish_start(), anonymity()
                                except:
                                    neo.vanish_tor()
                                    break
                            elif choice == '3':
                                try:
                                    return anonymous.vanish_stop(), anonymity()
                                except:
                                    neo.vanish_tor()
                                    break
                            elif choice == '4':
                                try:
                                    return anonymous.chains_start(), anonymity()
                                except:
                                    neo.vanish_tor()
                                    break
                            elif choice == '5':
                                try:
                                    return anonymous.checktor_status(), anonymity()
                                except:
                                    neo.vanish_tor()
                                    break
                            else:
                                try:
                                    print("")
                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 5 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                    for w in warn:
                                        sys.stdout.write(w)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                    neo.vanish_tor()
                                    break
                                except:
                                    neo.vanish_tor()
                                    break
                        except:
                            os.system('clear')
                            beauty.graphics(), scriptures.verses()
                            break
                anonymity()
                break
            except:
                break

#3
    def scann_internal(self):
        while True:
            try:
                return internal_scanner.bettercap_discover()
            except:
                neo.one()
                break

    def attack_internal(self):
        while True:
            try:
                os.system('clear')
                neo.scann_internal()
                print(bcolors.BLUE + "\n~>( " + bcolors.DARKCYAN + "Select a target from the table to be " + bcolors.RED + "Attacked! 🎯" + bcolors.BLUE + ")<~\n" + bcolors.ENDC)
                try:
                    host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    neo.one()
                    break
                os.system('clear')
                def neo_attack():
                    while True:
                        try:
                            mega_menu.menu_three()
                            print(bcolors.BLUE + "~>( " + bcolors.RED + "Ready to attack " + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
                            choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        except:
                            neo.one()
                            break
                        try:
                            if choice == '0':
                                neo.one()
                                break
                            elif choice == '1':
                                try:
                                    return neo.attack_internal()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '2':
                                try:
                                    return internal_scanner.nmap_pscanner(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '3':
                                try:
                                    return internal_scanner.nmap_vulnscanner(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '4':
                                try:
                                    return internal_scanner.smb_enumuration(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '5':
                                try:
                                    return internal_scanner.smb_exploit(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '6':
                                try:
                                    return internal_scanner.packets_sniffer(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '7':
                                try:
                                    return internal_scanner.packets_responder(), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '8':
                                try:
                                    return internal_scanner.beefxss_bettercap(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '9':
                                try:
                                    return internal_scanner.toxssin_xss(), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            else:
                                try:
                                    print("")
                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                    for w in warn:
                                        sys.stdout.write(w)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                    os.system('clear')
                                    attack_internal.neo_attack()
                                    break
                                except:
                                    attack_internal.neo_attack()
                                    break
                        except:
                            neo_attack()
                            break
                neo_attack()
                break
            except:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                break

#4
    def rat_kitchen(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_four()
                try:
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    neo.one()
                    break
                try:
                    if choice == '0':
                        neo.one()
                        break
                    elif choice == '1':
                        try:
                            return rat.blackjack(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '2':
                        try:
                            return rat.shellz(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '3':
                        try:
                            return rat.powerjoker(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '4':
                        try:
                            return rat.meterpeter(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '5':
                        try:
                            return rat.havoc(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '6':
                        try:
                            return rat.teardroid(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '7':
                        try:
                            return rat.androrat(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '8':
                        try:
                            return rat.chameleon(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '9':
                        try:
                            return rat.chameleon_1(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    else:
                        try:
                            print("")
                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                            neo.rat_kitchen()
                            break
                        except:
                            neo.rat_kitchen()
                            break
                except:
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    break
            except:
                break

#5
    def attack_wifi(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_five()
                try:
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    neo.one()
                    break
                try:
                    if choice == '0':
                        neo.one()
                        break
                    elif choice == '1':
                        try:
                            return wifi_killer.wifi_auto_attack_wifite(), neo.attack_wifi()
                        except:
                            neo.attack_wifi()
                            break
                    elif choice == '2':
                        try:
                            return wifi_killer.wifi_auto_attack_bettercap(), neo.attack_wifi()
                        except:
                            neo.attack_wifi()
                            break
                    elif choice == '3':
                        try:
                            return wifi_killer.wifi_auto_attack_wifipumpkin3(), neo.attack_wifi()
                        except:
                            neo.attack_wifi()
                            break
                    elif choice == '4':
                        try:
                            pass
                                #return wifi_killer.wifi_attack_airgeddon(), neo.attack_wifi()
                        except:
                                #neo.attack_wifi()
                                #break
                            pass
                    elif choice == '5':
                        try:
                            pass
                                #return wifi_killer.wifi_attack_airgeddon(), neo.attack_wifi()
                        except:
                                #neo.attack_wifi()
                                #break
                            pass
                    elif choice == '6':
                        try:
                            return wifi_killer.wifi_attack_airgeddon(), neo.attack_wifi()
                        except:
                            neo.attack_wifi()
                            break
                    elif choice == '7':
                        try:
                            return wifi_killer.wifi_attack_wifipumpkin3(), neo.attack_wifi()
                        except:
                            neo.attack_wifi()
                            break
                    elif choice == '8':
                        try:
                            pass
                                #return wifi_killer.wifi_attack_airgeddon(), neo.attack_wifi()
                        except:
                                #neo.attack_wifi()
                                #break
                            pass
                    elif choice == '9':
                        try:
                            pass
                                #return wifi_killer.wifi_attack_airgeddon(), neo.attack_wifi()
                        except:
                                #neo.attack_wifi()
                                #break
                            pass
                    else:
                        try:
                            print("")
                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                            neo.attack_wifi()
                            break
                        except:
                            neo.attack_wifi()
                            break
                except:
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    break
            except:
                break

#6
    def crack_passwords(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_six()
                try:
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    neo.one()
                    break
                try:
                    if choice == '0':
                        neo.one()
                        break
                    elif choice == '1':
                        try:
                            os.system('clear')
                            def automated_online():
                                mega_menu.menu_six_one()
                                while True:
                                    try:
                                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                    except:
                                        neo.crack_passwords()
                                        break
                                    if choice == '0':
                                        return neo.crack_passwords()
                                    elif choice == '1':
                                        return pass_cracker.hydra_ssh(), automated_online()
                                    elif choice == '2':
                                        return pass_cracker.hydra_ftp(), automated_online()
                                    elif choice == '3':
                                        return pass_cracker.hydra_smb(), automated_online()
                                    elif choice == '4':
                                        return pass_cracker.hydra_rdp(), automated_online()
                                    elif choice == '5':
                                        return pass_cracker.hydra_ldap(), automated_online()
                                    elif choice == '6':
                                        return pass_cracker.hydra_smtp(), automated_online()
                                    elif choice == '7':
                                        return pass_cracker.hydra_telnet(), automated_online()
                                    elif choice == '8':
                                        return pass_cracker.hydra_https(), automated_online()
                                    elif choice == '9':
                                        return pass_cracker.cyberbrute_all(), automated_online()
                                    else:
                                        try:
                                            print("")
                                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                            for w in warn:
                                                sys.stdout.write(w)
                                                sys.stdout.flush()
                                                time.sleep(0.09)
                                            os.system('clear')
                                            automated_online()
                                            break
                                        except:
                                            os.system('clear')
                                            automated_online()
                                            break
                            automated_online()
                            break
                        except:
                            neo.crack_passwords()
                            break
                    elif choice == '2':
                        try:
                            os.system('clear')
                            def automated_offline():
                                mega_menu.menu_six_two()
                                while True:
                                    try:
                                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                    except:
                                        neo.crack_passwords()
                                        break
                                    if choice == '0':
                                        return neo.crack_passwords()
                                    elif choice == '1':
                                        return pass_cracker.aircrack_ng(), automated_offline()
                                    elif choice == '2':
                                        return pass_cracker.john_cracker(), automated_offline()
                                    elif choice == '3':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '4':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '5':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '6':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '7':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '8':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    elif choice == '9':
                                        return pass_cracker.hash_buster(), automated_offline()
                                    else:
                                        try:
                                            print("")
                                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                            for w in warn:
                                                sys.stdout.write(w)
                                                sys.stdout.flush()
                                                time.sleep(0.09)
                                            os.system('clear')
                                            automated_offline()
                                            break
                                        except:
                                            os.system('clear')
                                            automated_offline()
                                            break
                            automated_offline()
                            break
                        except:
                            neo.crack_passwords()
                            break
                    else:
                        try:
                            print("")
                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 2 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                            neo.crack_passwords()
                            break
                        except:
                            neo.crack_passwords()
                            break
                except:
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    break
            except:
                break

#7
    def phish_creds(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_seven()
                try:
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    neo.one()
                    break
                try:
                    if choice == '0':
                        neo.one()
                        break
                    elif choice == '1':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_gophish(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '2':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_goodginx(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '3':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_zphisher(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '4':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_setoolkit(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '5':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_anonphisher(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '6':
                        try:
                            os.system('clear')
                            return cred_phisher.phish_cyberphish(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    else:
                        try:
                            print("")
                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 6 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                            neo.phish_creds()
                            break
                        except:
                            neo.phish_creds()
                            break
                except:
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    break
            except:
                break

#8
    def attack_websites(self):
        while True:
            try:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                print(bcolors.BLUE + "\n Enter Your Target: ~> " + bcolors.RED + "Either HTTP(S)//: HOSTNAME or IP 🎯\n" + bcolors.ENDC)
                try:
                    url = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                    def get_proxy():
                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Use proxies for connections [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        choice += " "
                        if choice[0].lower() == 'y':
                            try:
                                proxy_config.prompt_for_proxy()
                                command_runner = CommandRunner(proxy_config)
                                proxy_config.set_proxies()
                            except:
                                warn = bcolors.GREEN + "(" + bcolors.RED + "Supported proxies are " + bcolors.BLUE + "~>" + bcolors.DARKCYAN + " SOCKS4, SOCKS5, HTTP or HTTPS" + bcolors.GREEN + ")#" + bcolors.ENDC
                                for w in warn:
                                    sys.stdout.write(w)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("")
                                get_proxy()
                        elif choice[0].lower() == 'n':
                            pass
                        else:
                            print(bcolors.BLUE + "      (" + bcolors.RED + "No choice selected. Pleas type ~> " + bcolors.DARKCYAN + "[y/n]" + bcolors.BLUE + ")" + bcolors.ENDC)
                            get_proxy()
                    get_proxy()
                except:
                    neo.one()
                    break
                def url_maker(url):
                    if not re.match(r'http(s?)\:', url):
                        url = 'http://' + url
                    parsed = urlsplit(url)
                    host = parsed.netloc
                    if host.startswith('www.'):
                        host = host[4:]
                    return url
                host = url_maker(url)
                spiders = scanners(host = '')
                os.system('clear')
                def user_nuke(self):
                    mega_menu.menu_eight()
                    print(bcolors.BLUE + "~>( " + bcolors.RED + "Ready to attack " + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
                    while True:
                        try:
                            choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        except:
                            neo.one()
                            break
                        try:
                            if choice == '0':
                                neo.one()
                                break
                            #Start passive web enumeration
                            elif choice == '1':
                                try:
                                    os.system('clear')
                                    url = host
                                    xhost = url.replace("https://", "").replace("http://", "").replace("www.", "")
                                    return spiders.wafw00f(host), spiders.whatweb(host), spiders.dnsrecon(xhost), spiders.nuclei(xhost), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Gather e-mails & subdomain namesfrom public sources
                            elif choice == '2':
                                try:
                                    os.system('clear')
                                    url = host
                                    xhost = url.replace("https://", "").replace("http://", "").replace("www.", "")
                                    return spiders.harvester(xhost), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Start Host Root File Bruteforcer
                            elif choice == '3':
                                try:
                                    os.system('clear')
                                    return spiders.gobuster(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Start SQL, XSS & SSRF Detection & Eploitation
                            elif choice == '4':
                                try:
                                    os.system('clear')
                                    def web_injector(self):
                                        mega_menu.menu_eight_four()
                                        print(bcolors.BLUE + "  ~>( " + bcolors.RED + "Launch Injection againist" + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
                                        while True:
                                            try:
                                                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                            except:
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            if choice == '0':
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            elif choice == '1':
                                                os.system('clear')
                                                return spiders.sqli_auto_sqlmap(host), web_injector(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return spiders.xss_auto_commix(host), spiders.xss_auto_xsser(host), spiders.xss_auto_katana(host), web_injector(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return spiders.sqli_man_sqlmap(host), web_injector(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return spiders.xss_man_commix(host), spiders.xss_man_xsser(host), web_injector(self)
                                            else:
                                                try:
                                                    print("")
                                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 4 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    os.system('clear')
                                                    web_injector(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    web_injector(self)
                                                    break
                                    web_injector(self)
                                    break
                                except:
                                    return user_nuke(self)
                                    break
                            #Netattacker Insane Web Reconing Next Generation Engine
                            elif choice == '5':
                                try:
                                    os.system('clear')
                                    def owasp(self):
                                        mega_menu.menu_eight_five()
                                        print(bcolors.BLUE + "~>( " + bcolors.RED + "Launch Netattacker againist" + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)

                                        while True:
                                            try:
                                                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                            except:
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            if choice == '0':
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            elif choice == '1':
                                                os.system('clear')
                                                return spiders.owasp_1(host), owasp(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return spiders.owasp_2(host), owasp(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return spiders.owasp_3(host), owasp(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return spiders.owasp_4(host), owasp(self)
                                            elif choice == '5':
                                                os.system('clear')
                                                return spiders.owasp_5(host), owasp(self)
                                            elif choice == '6':
                                                os.system('clear')
                                                return spiders.owasp_6(host), owasp(self)
                                            elif choice == '7':
                                                os.system('clear')
                                                return spiders.owasp_7(host), owasp(self)
                                            elif choice == '8':
                                                os.system('clear')
                                                return spiders.owasp_8(host), owasp(self)
                                            elif choice == '9':
                                                os.system('clear')
                                                return spiders.owasp_9(host), owasp(self)
                                            else:
                                                try:
                                                    print("")
                                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    os.system('clear')
                                                    owasp(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    owasp(self)
                                                    break
                                    owasp(self)
                                    break
                                except:
                                    return owasp(self)
                                    break
                            #Jok3r v3.5 Insane Reconing Next Generation Engine
                            elif choice == '6':
                                try:
                                    os.system('clear')
                                    def jok3r(self):
                                        mega_menu.menu_eight_six()
                                        print(bcolors.BLUE + "~>( " + bcolors.RED + "Launch Jok3r againist" + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)

                                        while True:
                                            try:
                                                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                            except:
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            if choice == '0':
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            elif choice == '1':
                                                os.system('clear')
                                                return spiders.jok3r_1(host), jok3r(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return spiders.jok3r_2(host), jok3r(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return spiders.jok3r_3(host), jok3r(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return spiders.jok3r_4(host), jok3r(self)
                                            elif choice == '5':
                                                os.system('clear')
                                                return spiders.jok3r_5(host), jok3r(self)
                                            elif choice == '6':
                                                os.system('clear')
                                                return spiders.jok3r_6(host), jok3r(self)
                                            elif choice == '7':
                                                os.system('clear')
                                                return spiders.jok3r_7(host), jok3r(self)
                                            elif choice == '8':
                                                os.system('clear')
                                                return spiders.jok3r_8(host), jok3r(self)
                                            elif choice == '9':
                                                os.system('clear')
                                                return spiders.jok3r_9(host), jok3r(self)
                                            else:
                                                try:
                                                    print("")
                                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    os.system('clear')
                                                    jok3r(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    jok3r(self)
                                                    break
                                    jok3r(self)
                                    break
                                except:
                                    return jok3r(self)
                                    break
                            #Osmedeus Next Generation A Workflow Engine
                            elif choice == '7':
                                try:
                                    os.system('clear')
                                    def osmedeus(self):
                                        mega_menu.menu_eight_seven()
                                        print(bcolors.BLUE + "~>( " + bcolors.RED + "Launch Osmedeus on" + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)

                                        while True:
                                            try:
                                                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                            except:
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            if choice == '0':
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            elif choice == '1':
                                                os.system('clear')
                                                return spiders.osmedeus_1(), osmedeus(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return spiders.osmedeus_2(host), osmedeus(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return spiders.osmedeus_3(host), osmedeus(self)
                                            elif choice == '4':
                                                return spiders.osmedeus_4(host), osmedeus(self)
                                            elif choice == '5':
                                                os.system('clear')
                                                return spiders.osmedeus_5(host), osmedeus(self)
                                            elif choice == '6':
                                                os.system('clear')
                                                return spiders.osmedeus_6(host), osmedeus(self)
                                            elif choice == '7':
                                                os.system('clear')
                                                return spiders.osmedeus_7(host), osmedeus(self)
                                            elif choice == '8':
                                                return spiders.osmedeus_8(), osmedeus(self)
                                            elif choice == '9':
                                                os.system('clear')
                                                return spiders.osmedeus_9(), osmedeus(self)
                                            else:
                                                try:
                                                    print("")
                                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    os.system('clear')
                                                    osmedeus(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    osmedeus(self)
                                                    break
                                    osmedeus(self)
                                    break
                                except:
                                    return user_nuke(self)
                                    break
                            # Ufonet Next Generation DDOS Tool Main Menu
                            elif choice == '8':
                                try:
                                    os.system('clear')
                                    def ddos(self):
                                        mega_menu.menu_eight_eight()
                                        print(bcolors.BLUE + "~>( " + bcolors.RED + "Launch DDOS againist" + bcolors.BLUE + " )" + bcolors.BLUE + " ~> " + bcolors.BLUE + "( " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " )<~\n" + bcolors.ENDC)

                                        while True:
                                            try:
                                                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                                            except:
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            if choice == '0':
                                                os.system('clear')
                                                return user_nuke(self)
                                                break
                                            elif choice == '1':
                                                os.system('clear')
                                                return spiders.ufonet_1(), ddos(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return spiders.ufonet_2(), ddos(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return spiders.ufonet_3(host), ddos(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return spiders.ufonet_4(host), ddos(self)
                                            elif choice == '5':
                                                os.system('clear')
                                                return spiders.ufonet_5(host), ddos(self)
                                            elif choice == '6':
                                                os.system('clear')
                                                return spiders.ufonet_6(host), ddos(self)
                                            elif choice == '7':
                                                os.system('clear')
                                                return spiders.ufonet_7(), ddos(self)
                                            elif choice == '8':
                                                os.system('clear')
                                                return spiders.ufonet_8(), ddos(self)
                                            elif choice == '9':
                                                os.system('clear')
                                                return spiders.ufonet_9(host), ddos(self)
                                            else:
                                                try:
                                                    print("")
                                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    os.system('clear')
                                                    ddos(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    ddos(self)
                                                    break
                                    ddos(self)
                                    break
                                except:
                                    return ddos(self)
                                    break
                            #Launch A Heavy Automation Attacks On The Target 
                            elif choice == '9':
                                try:
                                    os.system('clear')
                                    url = host
                                    zhost = url.replace("http://", "https://").replace("www.", "")
                                    xhost = url.replace("https://", "").replace("http://", "").replace("www.", "")
                                    return spiders.dnsrecon(xhost), spiders.wafw00f(host), spiders.whatweb(host), spiders.nuclei(xhost), spiders.seekolver(xhost), spiders.gobuster(host), spiders.osmedeus_3(host), spiders.param_spider(host), spiders.ssl_scan(zhost), spiders.nuclei(host), spiders.sqli_auto_sqlmap(host), spiders.xss_auto_commix(host), spiders.xss_auto_katana(host), spiders.xss_auto_xsser(host), spiders.bbot(host), spiders.nikto(host), spiders.uniscan(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            else:
                                try:
                                    print("")
                                    warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                                    for w in warn:
                                        sys.stdout.write(w)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                    attack_websites.user_nuke(self)
                                    break
                                except:
                                    os.system('clear')
                                    user_nuke(self)
                                    break
                        except:
                            os.system('clear')
                            beauty.graphics(), scriptures.verses()
                            break
                user_nuke(self)
                break
            except:
                break

#9
    def about_me(self):
        while True:
            try:
                os.system('clear')
                guide_info.guide(), mega_menu.menu_nine()
                salvation =  bcolors.DARKCYAN + " For God so loved the world, that He gave." + bcolors.ENDC + color() + "  <~[ John 3:16 ]" + bcolors.ENDC
                for s in salvation:
                    sys.stdout.write(s)
                    sys.stdout.flush()
                    time.sleep(0.09)
                print("")
            except:
                neo.one()
                break
            try:
                choice = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '0':
                    neo.one()
                    break
                else:
                    neo.one()
                    break
            except:
                neo.one()
                break
#99
    def get_saved(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_nine_tisa()
                salvation.gospel()
                break
            except:
                neo.one()
                break
            try:
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '0':
                    neo.one()
                    break
                else:
                    neo.one()
                    break
            except:
                neo.one()
                break

# (all)
    def one(self):
        while True:
            try:
                os.system('clear')
                mega_menu.menu_one()
                try:
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.ENDC + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                except:
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    break
                try:
                    if choice == '0':
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        break
                    elif choice == '1':
                        try:
                            return neo.install_africana()
                        except:
                            break
                    elif choice == '2':
                        try:
                            return neo.vanish_tor()
                        except:
                            break
                    elif choice == '3':
                        try:
                            return neo.attack_internal()
                        except:
                            break
                    elif choice == '4':
                        try:
                            return neo.rat_kitchen()
                        except:
                            break
                    elif choice == '5':
                        try:
                            return neo.attack_wifi()
                        except:
                            break
                    elif choice == '6':
                        try:
                            return neo.crack_passwords()
                        except:
                            break
                    elif choice == '7':
                        try:
                            return neo.phish_creds()
                        except:
                            break
                    elif choice == '8':
                        try:
                            return neo.attack_websites()
                        except:
                            break
                    elif choice == '9':
                        try:
                            return neo.about_me()
                        except:
                            break
                    elif choice == '99':
                        try:
                            return neo.get_saved()
                        except:
                            break
                    else:
                        try:
                            print("")
                            warn = bcolors.BLUE + "~>( " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + ")<~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                        except:
                            neo.one()
                            break
                except:
                    neo.one()
                    break
            except:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                break

neo = neo_start()
neo.user_agree()

if ' __name__' == '__main__':
        sys.exit(neo.user_agree())
