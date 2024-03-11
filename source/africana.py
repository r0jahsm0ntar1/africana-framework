#!/usr/bin/env python3
import os
import sys
import time
import subprocess

def check_os():
    if os.name == "nt":
        operating_system = "windows"
    if os.name == "posix":
        operating_system = "posix"
    return operating_system

if check_os() == "posix":
    class bcolors:
        PURPLE = '\033[95m'
        CYAN = '\033[96m'
        DARKCYAN = '\033[36m'
        BLUE = '\033[94m'
        GREEN = '\033[92m'
        YELLOW = '\033[93m'
        RED = '\033[91m'
        BOLD = '\033[1m'
        UNDERL = '\033[4m'
        ENDC = '\033[0m'
        backBlack = '\033[40m'
        backRed = '\033[41m'
        backGreen = '\033[42m'
        backYellow = '\033[43m'
        backBlue = '\033[44m'
        backMagenta = '\033[45m'
        backCyan = '\033[46m'
        backWhite = '\033[47m'

        def disable(self):
            self.PURPLE = ''
            self.CYAN = ''
            self.BLUE = ''
            self.GREEN = ''
            self.YELLOW = ''
            self.RED = ''
            self.ENDC = ''
            self.BOLD = ''
            self.UNDERL = ''
            self.backBlack = ''
            self.backRed = ''
            self.backGreen = ''
            self.backYellow = ''
            self.backBlue = ''
            self.backMagenta = ''
            self.backCyan = ''
            self.backWhite = ''
            self.DARKCYAN = ''

else:
    class bcolors:
        PURPLE = ''
        CYAN = ''
        DARKCYAN = ''
        BLUE = ''
        GREEN = ''
        YELLOW = ''
        RED = ''
        BOLD = ''
        UNDERL = ''
        ENDC = ''
        backBlack = ''
        backRed = ''
        backGreen = ''
        backYellow = ''
        backBlue = ''
        backMagenta = ''
        backCyan = ''
        backWhite = ''

        def disable(self):
            self.PURPLE = ''
            self.CYAN = ''
            self.BLUE = ''
            self.GREEN = ''
            self.YELLOW = ''
            self.RED = ''
            self.ENDC = ''
            self.BOLD = ''
            self.UNDERL = ''
            self.backBlack = ''
            self.backRed = ''
            self.backGreen = ''
            self.backYellow = ''
            self.backBlue = ''
            self.backMagenta = ''
            self.backCyan = ''
            self.backWhite = ''
            self.DARKCYAN = ''

class main(object):
    def __init__(self):
        pass

    def intro_banner(self):
        print(bcolors.GREEN + bcolors.BOLD + " \nHey U! Jesus Is @ the door Knocking. Do Something. " + bcolors.ENDC)
        print(bcolors.GREEN + bcolors.BOLD + """
          wake up, Christian
   Lord God Jesus Christ L❤️.VE'S you
       follow the white Pigeon.
         knock, knock, knock,
             Man Of God.""" + bcolors.ENDC)
        start = r'''
                       (`.         ,-,
                        ` `.    ,;' /
                         `.  ,'/ .'
                          `. X /.'
                .-;--''--.._` ` (
              .'            /   `
             ,           ` '   Q '
             ,         ,   `._    \
          ,.|         '     `-.;_' '
          :  . `  ;    `  ` --,.._;
           ' `    ,   )   .'
              `._ ,  '   /_
                 ; ,''-,;' ``-     [By: r0jahsm0ntar1]
                  ``-..__``--`  [God is True ❤️. John 3:16]'''

        for s in start:
            sys.stdout.write(s)
            sys.stdout.flush()
            time.sleep(0.001)

    def attacker_banner(self):
        print(bcolors.GREEN + bcolors.BOLD + """

                            .,'
                        .''.'
                        .' .'
            _.ood0Pp._ ,'  `.~ .q?00doo._
        .od00Pd0000Pdb._. . _:db?000b?000bo.
     .?000Pd0000Pd0000PdbMb?0000b?000b?0000b.
    .d0000Pd0000Pd0000Pd0000b?0000b?000b?0000b.
    d0000Pd0000Pd00000Pd0000b?00000b?0000b?000b.
    00000Pd0000Pd0000Pd00000b?00000b?0000b?0000b
    ?0000b?0000b?0000b?00Pd0[Praise be to Jesus]
    ?0000b?0000b?0000b?00000Pd00000Pd0000Pd000P
    `?0000b?0000b?0000b?0000Pd0000Pd0000Pd000P'
     `?000b?0000b?000b?0000Pd000Pd0000Pd000P
        `~?00b?000b?000b?000Pd00Pd000Pd00P'
            `~?0b?0b?000b?0Pd0Pd000PdP~'         [Christian]""" + bcolors.ENDC)

    def closser_banner(self):
        close = r'''
        ____                 _
       / __ \___  __________(_)___   _____
      / /_/ / _ \/ ___/ ___/ / __ ` / __  \ 
     / ____/  __/ /  (__  ) / /_/ // / /  /
    /_/    \___/_/In Serving Jesus.Defeat the Devil By: Fasting & Praying.'''
        for c in close:
            sys.stdout.write(c)
            sys.stdout.flush()
            time.sleep(0.01)

class scanners(main):
    def __init__(self, host):
        self.host = host
        main.attacker_banner(self)

    def nmap(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] Nmap Scanner Has Began to find Open Ports .." + bcolors.ENDC)
        process = subprocess.Popen("nmap -p- {0}".format(host), shell = True).wait()
        return process

    def dnsrec(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] Dnsrecon has started Scanning for subdomains .." + bcolors.ENDC)
        time.sleep(0.1)
        process = subprocess.Popen("dnsrecon -a -d {0}".format(host), shell = True).wait()
        return process
    def whatweb(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] WhatWeb Scanner has begun to find running technology .." + bcolors.ENDC)
        time.sleep(0.1)
        process = subprocess.Popen("whatweb -a 3 -v {0}".format(host), shell = True).wait()
        return process
    def nuclei(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] Nuclei Scanner has begun Vulnerbility Scanning .." + bcolors.ENDC)
        time.sleep(0.1)
        process = subprocess.Popen("nuclei -target {0}".format(host), shell = True).wait()
        return process
    def nikto(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] Nikto Scanner has started Vulnerbility Scanning .." + bcolors.ENDC)
        time.sleep(0.1)
        process = subprocess.Popen("nikto -C all -host {0}".format(host), shell = True).wait()
        return process
    def ferox(self, host):
        print(bcolors.BLUE + bcolors.BOLD + "\n[+] FeroxBuster Scanner has started Discovering host internal files .." + bcolors.ENDC)
        time.sleep(0.1)
        process = subprocess.Popen("feroxbuster --url http://{0}".format(host), shell = True).wait()
        return process

class Options():
    def menu(self):
        print(bcolors.BLUE + bcolors.BOLD +"""
[---]       Choose what to do from the menu below      [---]"""+ bcolors.ENDC)
        print(bcolors.BLUE + """
1) Port Scanning                      5) Start Nikto Scaning
2) Dns Reconning                      6) Fero File Searching
3) Web Technologies                   7) Automation Scanning
4) Nuclei Vuln Scanning               8) Exit Africana Tool."""+ bcolors.ENDC)

os.system('clear')
shamura = main()
shamura.intro_banner()

host = input(bcolors.RED + bcolors.BOLD + "\n\n[+] What is your host to Attack:?... ~$ "+ bcolors.ENDC)
os.system('clear')
spiders = scanners(host = '')

def select():
    optio = Options()
    optio.menu()
    choice = input(bcolors.RED + bcolors.BOLD + "\n[+] What is Your Choice from the above table ?: " + bcolors.ENDC)
    while True:
        try:
            if choice == '1':
                os.system('clear')
                return  spiders.nmap(host), select()
            elif choice == '2':
                os.system('clear')
                return spiders.dnsrec(host), select()
            elif choice == '3':
                os.system('clear')
                return spiders.whatweb(host), select()
            elif choice == '4':
                os.system('clear')
                return spiders.nuclei(host), select()
            elif choice == '5':
                os.system('clear')
                return spiders.nikto(host), select()
            elif choice == '6':
                os.system('clear')
                return spiders.ferox(host), select()
            elif choice == '7':
                os.system('clear')
                return  spiders.nmap(host), spiders.dnsrec(host), spiders.whatweb(host), spiders.nuclei(host), spiders.nikto(host), spiders.ferox(host), select()
            elif choice == '8':
                os.system('clear')
                print(bcolors.RED + bcolors.BOLD + "\n[-] Exiting The Engine Bye and Get Saved if Not .." + bcolors.ENDC)
                shamura.closser_banner()
                break
            else:
                print(bcolors.RED + bcolors.BOLD + "\n[-] Critical Error has happened" + bcolors.ENDC)
                break
        except:
            print(bcolors.RED + bcolors.BOLD + "\n[-] Critical Error has happened" + bcolors.ENDC)
            break
select()

if __name__ == '__main__':
        main()
