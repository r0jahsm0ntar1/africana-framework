import re
import os
import sys
import time
import subprocess

from signal import *
from src.guide.info import *
from urllib.parse import *
from src.core.system import *
from src.core.banner import *
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

class main_menu(object):
    def __init_(self):
        pass

    def menu_zero(self):
        print(bcolors.RED + "\n[ The Africana-Framework is designed purely for Good  and ] " + bcolors.ENDC)
        print(bcolors.RED + "[ not evil. If you are planning on using this tool for    ] " + bcolors.ENDC)
        print(bcolors.RED + "[ malicious purposes that are not authorized by the       ] " + bcolors.ENDC)
        print(bcolors.RED + "[ company you are performing assessments for, you are     ] " + bcolors.ENDC)
        print(bcolors.RED + "[ violating the terms of service and license of this      ] " + bcolors.ENDC)
        print(bcolors.RED + "[ toolset. By hitting yes (only one time), you agree to   ] " + bcolors.ENDC)
        print(bcolors.RED + "[ the terms of service and that you will only use this    ] " + bcolors.ENDC)
        print(bcolors.RED + "[ tool for lawful purposes only.                          ] \n" + bcolors.ENDC)

    def menu_one(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install or update africana-framework ~> " + bcolors.DARKCYAN + "(Start here )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. System Security Configuration        ~> " + bcolors.DARKCYAN + "(Setup tor &)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Local Network Attack Vectors         ~> " + bcolors.DARKCYAN + "(Mitm, sniff)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Generate Undetectable Malware        ~> " + bcolors.DARKCYAN + "(C2 & shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. WiFi Attack Vectors                  ~> " + bcolors.DARKCYAN + "(Wifite, air)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Crack Hash, Pcap & Brute Passwords   ~> " + bcolors.DARKCYAN + "(Hashcat, jo)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Social-Engineering Attacks           ~> " + bcolors.DARKCYAN + "(Gophish, gi)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Website Attack Vectors               ~> " + bcolors.DARKCYAN + "(Osmedeus, j)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Help, Credits, Tricks and About      ~> " + bcolors.BLUE + "   💡99. " + bcolors.BLUE + "    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit africana-framework              ~> " + bcolors.DARKCYAN + "(Inspiration)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.GREEN + bcolors.BOLD + "                                              ¯\_(ツ)_/¯ " + bcolors.ENDC)

    def menu_two(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                 Install & Setup Tor🎭    " + bcolors.DARKCYAN + "(start here)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.            Start anonymizing through tor             ]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.            Stop tor & restore all iptables           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.   Chains (local ~> squid ~> privoxy ~> tor ~> net)   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                  Check if using tor                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.             Exit and go back fo main menu            ] \n" + bcolors.ENDC)

    def menu_three(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                 🐹Discover Targets                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.            Port discovery on the target             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                Vuln' Scann the Target               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.            S.M.B enumration on the target           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.               Exploit S.M.B on target               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                Sniff Internal Packets               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.           Lunch Responder with IPv6 Support         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.         Start Beefxss & Bettercap For (M.I.B)       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.         Kick people out of the network (DDOS)       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go to main menu               ] \n" + bcolors.ENDC)

    def menu_four(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ShellzGen               ~>       " + bcolors.DARKCYAN + "(All Distro   R.A.T)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Blackjack               ~>       " + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. PowerJoker              ~>       " + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. MeterPeter              ~>       " + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Havoc C2                ~>       " + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Teardroid               ~>       " + bcolors.DARKCYAN + "(Android 4 > 13 Rat)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. AndroRAT                ~>       " + bcolors.DARKCYAN + "(Android 4 > 10 Rat)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go To Main Menu                ] \n" + bcolors.ENDC)

    def menu_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                       Wifite              " + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                      Bettercap            " + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    Wifipumpkin3  " + bcolors.RED + "(Automated Phishing)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                     Airgeddon                " + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                    wifiPumpkin3              " + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.        online (automated/ manual/ Bruteforce)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.          offline (automated/ maunal/ Hashes)         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_six_one(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                    Bruteforce " + bcolors.DARKCYAN + "SSH" + bcolors.BLUE + "                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                    Bruteforce " + bcolors.DARKCYAN + "FTP" + bcolors.BLUE + "                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    Bruteforce " + bcolors.DARKCYAN + "SMB" + bcolors.BLUE + "                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                    Bruteforce " + bcolors.DARKCYAN + "RDP" + bcolors.BLUE + "                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                    Bruteforce " + bcolors.DARKCYAN + "LDAP" + bcolors.BLUE + "                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                    Bruteforce " + bcolors.DARKCYAN + "SMTP" + bcolors.BLUE + "                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                    Bruteforce " + bcolors.DARKCYAN + "Telnet" + bcolors.BLUE + "                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                    Bruteforce " + bcolors.DARKCYAN + "HTTP/S" + bcolors.BLUE + "                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                    Bruteforce " + bcolors.DARKCYAN + "All/SS" + bcolors.BLUE + "                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                 Exit & Go To Main Menu               ] \n" + bcolors.ENDC)

    def menu_six_two(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                     Aircrack_ng  (offline)" + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                        John      (offline)" + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                     Hash-Buster  (offline)           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_seven(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                Gophish(Browser Gui)  " + bcolors.DARKCYAN + "(All Templetes )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                Good Ginx (Advanced)  " + bcolors.DARKCYAN + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    AdvPhishing       " + bcolors.DARKCYAN + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                     Setoolkit        " + bcolors.DARKCYAN + "(Clones Website)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                    Anonphisher       " + bcolors.DARKCYAN + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                     Cyberphish " + bcolors.DARKCYAN + "(Gene phishing temps) " + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go To Main Menu                ] \n" + bcolors.ENDC)

    def menu_eight(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Start Passive Web recon & Subdomain Enumration       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Gather e-mails & subdomain namesfrom public sources  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Start Bruteforcing Host's Root Files                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Start SQL, XSS & SSRF Detection & Eploitation        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Launch OWASP Nettacker project MainMenu              ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu     ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Osmedeus Next Generation A Workflow Engine Main Menu ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Ufonet Next Generation DDOS Tool Main Menu           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Insane & Heavy Automation Attacks On The Host ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit Web Scanner & Go To Main Menu                   ] \n" + bcolors.ENDC)

    def menu_eight_four(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.             Sql injection with sqlmap     " + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.  Xss injection with Commix, Dalfox, Xsser " + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.             Sql injection with Sqlmap        " + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.             Xss injection with Xsser         " + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)

    def menu_eight_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Start Nettacker port & web content discovery         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Recon, find and scan subdomains                      ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch admin_scan to find admin panel                ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Start insane information gathering on host           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Run CVE scans on the target host                     ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Automate all modules & security checks on target     ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. View scanned Nettacker report list                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu                               ] \n" + bcolors.ENDC)

    def menu_eight_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install all jok3r tools ~> " + bcolors.RED + "(Pleas start here if not )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Update all the tools in the toolbox                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Show all the tools in the toolbox                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Show supported products for all services             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Run all & intense security checks against an URL     ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. View the full resultsfrom the security checks        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Clean database & delete results                      ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu                               ] \n" + bcolors.ENDC)

    def menu_eight_seven(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Update Osmedeus & Run diagnostics to check config.   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Start a simple scan with other flow                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Directly run on vuln scan & directory scan on domains] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Scan list of targets (Full path of target needed..)  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Cloud - Run scan in Distributed Cloud mode           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Performing Full web vuln & secret scan on host       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Update the vulnerability database before attacking   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start web UI server                                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. View scanned osmedeus report list                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu                               ] \n" + bcolors.ENDC)

    def menu_eight_eight(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Download list of " + bcolors.RED + "Bots " + bcolors.BLUE + " ~>      " + bcolors.BLUE+ "from" + bcolors.YELLOW + " Community " + bcolors.BLUE + "server ]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Test if bots are alive ~>                            ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch palantir        ~>         " + bcolors.DARKCYAN + "(palantir 3.14    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Launch socking_waves   ~>         " + bcolors.DARKCYAN + "(instant-knockout!)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Launch xcom-1          ~>         " + bcolors.DARKCYAN + "(only DDoS        )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Launch xcom-2          ~>         " + bcolors.DARKCYAN + "(only DoS         )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Launch ufonet-gui      ~>         " + bcolors.DARKCYAN + "(gui on browser   )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start Grider           ~> " + bcolors.DARKCYAN + "(python3 ufonet --grider &)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Armageddon!     ~> " + bcolors.DARKCYAN + "(with ALL!)    (Take care.)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_nine(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n+---------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "[   I am Rojahs Montari a Devoted Christian & Pentester   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[                                                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ Email       ~>: rojahsmontari@gmail.com                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ YouTube     ~>: https://youtube.com/@RojahsMontari      ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ Life Tip    ~>: Defeat the devil by fasting & praying   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[                               ¯\_(ツ)_/ ¯               ]  " + bcolors.ENDC)
        print(bcolors.BLUE + "+(Press Enter To go BACK)---------------------------------+ \n" + bcolors.ENDC)

    def menu_nine_tisa(self):
        print(color() + """
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|""" + bcolors.ENDC)
        print(bcolors.PURPLE + "\n     J    E    S    U    S    C    H    R     I    S   T     " + bcolors.ENDC)
        print(bcolors.BOLD + color() + "     *    *    *    *    *    *    *    *    *     *   *     \n" + bcolors.ENDC)
        print(bcolors.BLUE + "\n                ~>[ " + bcolors.ENDC + bcolors.BOLD + "The Ten Commandments." + bcolors.ENDC + bcolors.BLUE + " ]<~        " + bcolors.ENDC)
        print(bcolors.BLUE + "+----------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "| 1. You shall have no other gods before Me.               | " + bcolors.ENDC)
        print(bcolors.BLUE + "+ 2. You shall make no idols.                              + " + bcolors.ENDC)
        print(bcolors.BLUE + "| 3. You shall not take the name of the Lord your God in   | " + bcolors.ENDC)
        print(bcolors.BLUE + "+    vein.                                                 + " + bcolors.ENDC)
        print(bcolors.BLUE + "| 4. Keep the Sabbath day holy.                            | " + bcolors.ENDC)
        print(bcolors.BLUE + "+ 5. Honor your father and your mother.                    + " + bcolors.ENDC)
        print(bcolors.BLUE + "| 6. You shall not murder.                                 | " + bcolors.ENDC)
        print(bcolors.BLUE + "+ 7. You shall not commit adultery.                        + " + bcolors.ENDC)
        print(bcolors.BLUE + "| 8. You shall not steal.                                  | " + bcolors.ENDC)
        print(bcolors.BLUE + "+ 9. You shall not bear false witness against your eighbor.+ " + bcolors.ENDC)
        print(bcolors.BLUE + "| 10 You shall not covet.                                  | " + bcolors.ENDC)
        print(bcolors.BLUE + "+----------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "\n     ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Tip: Defeat the devil by fasting & praying." + bcolors.ENDC + bcolors.BLUE + " ]<~        " + bcolors.ENDC)
        print(bcolors.GREEN + "                         ¯\_(ツ)_/¯                       \n" + bcolors.ENDC)

mega_menu = main_menu()

def sudo():
    if not os.geteuid() == 0:
        os.system('clear')
        beauty.graphics(), scriptures.verses()
        sys.exit(1)
sudo ()

class neo_start(object):
    def __init__(self):
        pass
#0
    def user_agree(self):
        while True:
            try:
                os.system('clear')
                if not os.path.isfile("src/agreement/covenant.txt"):
                    with open("readme/LICENSE") as fileopen:
                        for line in fileopen:
                            print((bcolors.ENDC + line.rstrip()))
                    mega_menu.menu_zero()
                    covenant = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "Do you agree to the terms of service [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
                    covenant += " "
                    if covenant[0].lower() == "y":
                        with open("src/agreement/covenant.txt", "w") as filewrite:
                            filewrite.write("user accepted")
                        os.chdir("reports")
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
                                    print("\n")
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 5 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                print(bcolors.BLUE + "\n~>[ Select a target from the table above to be" + bcolors.RED + " Attacked! " + bcolors.BLUE + "]<~\n" + bcolors.ENDC)
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
                            print(bcolors.BLUE + "~>[ " + bcolors.RED + "Ready to attack " + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                    return internal_scanner.kickthemout_dos(), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            else:
                                try:
                                    print("\n")
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
                                    for w in warn:
                                        sys.stdout.write(w)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
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
                            return rat.shellz(), neo.rat_kitchen()
                        except:
                            neo.rat_kitchen()
                            break
                    elif choice == '2':
                        try:
                            return rat.blackjack(), neo.rat_kitchen()
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
                    else:
                        try:
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                            print("\n")
                                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
                                            for w in warn:
                                                sys.stdout.write(w)
                                                sys.stdout.flush()
                                                time.sleep(0.09)
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
                                            print("\n")
                                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 2 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 6 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                print(bcolors.BLUE + "\n  Enter Your Target: ~> " + bcolors.RED + "Either HTTP(S)//: HOSTNAME or IP \n" + bcolors.ENDC)
                try:
                    url = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
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
                    print(bcolors.BLUE + "~>[ " + bcolors.RED + "Ready to attack " + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Launch Injection againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 4 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launch Netattacker againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launch Jok3r againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launch Osmedeus on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launch DDOS againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                                    return spiders.dnsrecon(xhost), spiders.wafw00f(host), spiders.whatweb(host), spiders.nuclei(xhost), spiders.seeker(xhost), spiders.gobuster(host), spiders.osmedeus_3(host), spiders.param_spider(host), spiders.ssl_scan(zhost), spiders.nuclei(host), spiders.sqli_auto_sqlmap(host), spiders.xss_auto_commix(host), spiders.xss_auto_katana(host), spiders.xss_auto_xsser(host), spiders.bbot(host), spiders.nikto(host), spiders.uniscan(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            else:
                                try:
                                    print("\n")
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
                print("\n")
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
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC
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
        sys.exit(neo())
