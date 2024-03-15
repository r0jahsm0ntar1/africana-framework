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
        print(bcolors.RED + "+---------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.RED + "[ DEALINGS IN THE SOFTWARE.                               ] " + bcolors.ENDC)
        print(bcolors.RED + "[ The Africana-Framework is designed purely for Good  and ] " + bcolors.ENDC)
        print(bcolors.RED + "[ not evil. If you are planning on using this tool for    ] " + bcolors.ENDC)
        print(bcolors.RED + "[ malicious purposes that are not authorized by the       ] " + bcolors.ENDC)
        print(bcolors.RED + "[ company you are performing assessments for, you are     ] " + bcolors.ENDC)
        print(bcolors.RED + "[ violating the terms of service and license of this      ] " + bcolors.ENDC)
        print(bcolors.RED + "[ toolset. By hitting yes (only one time), you agree to   ] " + bcolors.ENDC)
        print(bcolors.RED + "[ the terms of service and that you will only use this    ] " + bcolors.ENDC)
        print(bcolors.RED + "[ tool for lawful purposes only.                          ] " + bcolors.ENDC)
        print(bcolors.RED + "+---------------------------------------------------------+ \n" + bcolors.ENDC)

    def menu_one(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install or update africana-framework ~> " + bcolors.DARKCYAN + "(Start here )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. System Security Configuration        ~> " + bcolors.DARKCYAN + "(Setup tor &)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Local Network Attack Vectors         ~> " + bcolors.DARKCYAN + "(Mitm, sniff)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Generate Undetectable Malware        ~> " + bcolors.DARKCYAN + "(C2 & shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. WiFi Attack Vectors                  ~> " + bcolors.DARKCYAN + "(Wifite, air)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Crack .Pcap & BruteForce Passwords   ~> " + bcolors.DARKCYAN + "(Hashcat, jo)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Social-Engineering Attacks           ~> " + bcolors.DARKCYAN + "(Gophish, gi)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Website Attack Vectors               ~> " + bcolors.DARKCYAN + "(Osmedeus, j)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Help, Credits, Tricks and About      ~> " + bcolors.BLUE + "   💡99. " + bcolors.BLUE + "    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit africana-framework              ~> " + bcolors.DARKCYAN + "(Inspiration)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.GREEN + bcolors.BOLD + "                                              ¯\_(ツ)_/¯ " + bcolors.ENDC)

    def menu_two(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                 Install & Setup Tor      " + bcolors.RED + "(start here)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.            Start anonymizing through tor             ]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.            Stop tor & restore all iptables           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.   Chains (local ~> squid ~> privoxy ~> tor ~> net)   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                  Check if using tor                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.             Exit and go back fo main menu            ] \n" + bcolors.ENDC)

    def menu_three(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                 🐹Discover Targets                  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.            Port discovery on the target             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                Vuln' Scann the Target               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.            S.M.B enumration on the target           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.               Exploit S.M.B on target               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                Sniff Internal Packets               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.           Lunch Responder with IPv6 Support         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.        Start Beefxss & Bettercap For (M.I.B)        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.         Bruteforce Password on given services       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go to main menu               ] \n" + bcolors.ENDC)

    def menu_three_nine(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                    Bruteforce SMB                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                    Bruteforce SSH                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    Bruteforce FTP                    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                    Bruteforce HTTP                   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                 Exit & Go To Main Menu               ] \n" + bcolors.ENDC)

    def menu_four(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Shellz                  ~>       " + bcolors.RED + "(All Distro R.A.T  )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Shakamura               ~>       " + bcolors.RED + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. PowerJoker              ~>       " + bcolors.RED + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. MeterPeter              ~>       " + bcolors.RED + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Havoc C2                ~>       " + bcolors.RED + "(Windows Rev Shells)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Teardroid               ~>       " + bcolors.RED + "(Android 4 > 13 Rat)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. AndroRAT                ~>       " + bcolors.RED + "(Android 4 > 10 Rat)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go To Main Menu                ] \n" + bcolors.ENDC)

    def menu_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                       Wifite              " + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                      Bettercap            " + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    Wifipumpkin3  " + bcolors.RED + "(Automated Phishing)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                     Airgeddon                " + bcolors.RED + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                    wifiPumpkin3              " + bcolors.RED + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                     Air-Crackng  (offline)" + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                        John      (offline)" + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_seven(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                Gophish(Browser Gui)  " + bcolors.RED + "(All Templetes )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                Good Ginx (Advanced)  " + bcolors.RED + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.                    AdvPhishing       " + bcolors.RED + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                     Setoolkit        " + bcolors.RED + "(Clones Website)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5.                    Anonphisher       " + bcolors.RED + "(OTP Bypass    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9.                       To Add                         ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                Exit & Go To Main Menu                ] \n" + bcolors.ENDC)

    def menu_eight(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Start Open Ports Discovery & Subdomain Enumration    ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Recon for Wafs & Running Web Technologies            ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Gather e-mails & subdomain namesfrom public sources  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Start Bruteforcing Host's Root Files                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Start SQL, XSS & SSRF Detection & Eploitation        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu     ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Osmedeus Next Generation A Workflow Engine Main Menu ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Ufonet Next Generation DDOS Tool Main Menu           ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Insane & Heavy Automation Attacks On The Host ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit Web Scanner & Go To Main Menu                   ] \n" + bcolors.ENDC)

    def menu_eight_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.             Sql injection with sqlmap     " + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.  Xss injection with Commix, Dalfox, Xsser " + bcolors.RED + "(Automated)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.             Sql injection with Sqlmap        " + bcolors.RED + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.             Xss injection with Xsser         " + bcolors.RED + "(Manual)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)

    def menu_eight_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a numberfrom the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Download list of " + bcolors.RED + "Bots " + bcolors.BLUE + " ~>      " + bcolors.BLUE+ "from" + bcolors.YELLOW + " Community " + bcolors.BLUE + "server ]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Test if bots are alive ~>                            ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch palantir        ~>         " + bcolors.RED + "(palantir 3.14    )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Launch socking_waves   ~>         " + bcolors.RED + "(instant-knockout!)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Launch xcom-1          ~>         " + bcolors.RED + "(only DDoS        )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Launch xcom-2          ~>         " + bcolors.RED + "(only DoS         )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Launch ufonet-gui      ~>         " + bcolors.RED + "(gui on browser   )" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start Grider           ~> " + bcolors.RED + "(python3 ufonet --grider &)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Armageddon!     ~> " + bcolors.RED + "(with ALL!)    (Take care.)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.               Exit & Go To Main Menu                 ] \n" + bcolors.ENDC)

    def menu_nine(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n+---------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "| I am Rojahs Montari a Devoted Christian & Pentester     | " + bcolors.ENDC)
        print(bcolors.BLUE + "+---------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "[ A    F    R    I    C    A    N    A     [  Framework.  ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ *    *    *    *    *    *    *    *     *              ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    |    |    |    |    |  0. Exit & Back ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    |    |    |    |    +- 1. Install or U] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    |    |    |    +------ 2. System Sec. ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    |    |    +----------- 3. Local Networ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    |    +---------------- 4. C2 & shells ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    |    +--------------------- 5. Wifi Attack ] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    |    + ------------------------- 6. Password Cra] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    |    + ------------------------------ 7. Social Engee] " + bcolors.ENDC)
        print(bcolors.BLUE + "|    +------------------------------------ 8. Website Atta] " + bcolors.ENDC)
        print(bcolors.BLUE + "+----------------------------------------- 9. About Author] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ Gmail  : rojahsmontari@gmail.com                        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ YouTube: https://youtube.com/@RojahsMontari             ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ Tip    : Defeat the devil by fasting & praying.         ] " + bcolors.ENDC)
        print(bcolors.BLUE +  "[                                ¯\_(ツ)_/¯               ] " + bcolors.ENDC)
        print(bcolors.BLUE + "(Press)(Enter)(To go BACK)+-------------------------------+ \n" + bcolors.ENDC)

    def menu_nine_tisa(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n(Genesis)+------------------------------------+(Revelation) " + bcolors.ENDC)
        print(bcolors.BLUE + "| I am Rojahs Montari a Devoted Christian & Pentester     | " + bcolors.ENDC)
        print(bcolors.BLUE + "+---------------------------------------------------------+ " + bcolors.ENDC)
        print(bcolors.BLUE + "[   J    E    S    U    S    C    H    R     I    S   T   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[   *    *    *    *    *    *    *    *    *     *   *   ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[       Tip: Defeat the devil by fasting & praying.       ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[                       ¯\_(ツ)_/¯                        ] " + bcolors.ENDC)
        print(bcolors.BLUE + "+---------------------------------------------------------+ \n" + bcolors.ENDC)

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
                if not os.path.isfile("src/agreement/covenant"):
                    with open("readme/LICENSE") as fileopen:
                        for line in fileopen:
                            print((bcolors.ENDC + line.rstrip()))
                    mega_menu.menu_zero()
                    covenant = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "Do you agree to the terms of service [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
                    covenant += " "
                    if covenant[0].lower() == "y":
                        with open("src/agreement/covenant", "w") as filewrite:
                            filewrite.write("user accepted")
                        os.chdir("reports")
                        neo.one()
                        break
                    else:
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        break
                else:
                    os.chdir("reports")
                    neo.one()
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
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0. to 5. " + bcolors.ENDC +  " ]<~" + bcolors.ENDC
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
                print(bcolors.ENDC + "\n~>[" + bcolors.BLUE + " Select a targetfrom the table above to be" + bcolors.RED + " Attacked! " + bcolors.ENDC + "]<~\n" + bcolors.ENDC)
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
                            print(bcolors.BLUE + "         ~>[ " + bcolors.RED + "ready to attack " + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                    os.system('clear')
                                    return internal_scanner.nmap_pscanner(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '3':
                                try:
                                    os.system('clear')
                                    return internal_scanner.nmap_vulnscanner(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '4':
                                try:
                                    os.system('clear')
                                    return internal_scanner.smb_enumuration(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '5':
                                try:
                                    os.system('clear')
                                    return internal_scanner.smb_exploit(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '6':
                                try:
                                    os.system('clear')
                                    return internal_scanner.packets_sniffer(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '7':
                                try:
                                    return internal_scanner.packets_responder(), attack_internal.neo_attack()
                                except:
                                    os.system('clear')
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '8':
                                try:
                                    os.system('clear')
                                    return internal_scanner.beefxss_bettercap(host), attack_internal.neo_attack()
                                except:
                                    attack_internal.neo_attack()
                                    break
                            elif choice == '9':
                                try:
                                    os.system('clear')
                                    def bruteforce(self):
                                        mega_menu.menu_three_nine()
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Password Bruteforce attacks againist" + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                                return bruteforce(self)
                                            elif choice == '2':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '3':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '4':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '5':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '6':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '7':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '8':
                                                os.system('clear')
                                                return bruteforce(self)
                                            elif choice == '9':
                                                os.system('clear')
                                                return bruteforce(self)
                                            else:
                                                try:
                                                    print("\n")
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.ENDC +  " ]<~" + bcolors.ENDC
                                                    for w in warn:
                                                        sys.stdout.write(w)
                                                        sys.stdout.flush()
                                                        time.sleep(0.09)
                                                    bruteforce(self)
                                                    break
                                                except:
                                                    os.system('clear')
                                                    bruteforce(self)
                                                    break
                                    bruteforce(self)
                                    break
                                except:
                                    return user_nuke(self)
                                    break
                            else:
                                try:
                                    print("\n")
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.ENDC +  " ]<~" + bcolors.ENDC
                                    for w in warn:
                                        sys.stdout.write(w)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                    attack_internal.neo_attack()
                                    break
                                except:
                                    os.system('clear')
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
                            return rat.shakamura(), neo.rat_kitchen()
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
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                            return pass_cracker.aircracking_password(), neo.crack_passwords()
                        except:
                            neo.crack_passwords()
                            break
                    elif choice == '2':
                        try:
                            return pass_cracker.john_password(), neo.crack_passwords()
                        except:
                            neo.crack_passwords()
                            break
                    else:
                        try:
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.ENDC + "]<~" + bcolors.ENDC
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
                            return cred_phisher.phish_gophish(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '2':
                        try:
                            return cred_phisher.phish_goodginx(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '3':
                        try:
                            return cred_phisher.phish_zphisher(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '4':
                        try:
                            return cred_phisher.phish_setoolkit(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    elif choice == '5':
                        try:
                            return cred_phisher.phish_anonphisher(), neo.phish_creds()
                        except:
                            neo.phish_creds()
                            break
                    else:
                        try:
                            print("\n")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                print(bcolors.ENDC + "\n~{ " + bcolors.BLUE + "Enter Your Target ( " + bcolors.RED + "Either HTTP(S)//: HostName OR IP" + bcolors.BLUE + " ) " + bcolors.ENDC + "]<~\n" + bcolors.ENDC)
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
                    print(bcolors.BLUE + "         ~>[ " + bcolors.RED + "ready to attack " + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                            #Start Open Ports Discovery & Subdomain Enumration
                            elif choice == '1':
                                try:
                                    os.system('clear')
                                    url = host
                                    xhost = url.replace("https://", "").replace("http://", "").replace("www.", "")
                                    return spiders.seeker(xhost), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            # Recon for Waf & Running Web Technology Detection
                            elif choice == '2':
                                try:
                                    os.system('clear')
                                    url = host
                                    xhost = url.replace("https://", "").replace("http://", "").replace("www.", "")
                                    return spiders.wafw00f(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Gather e-mails & subdomain namesfrom public sources
                            elif choice == '3':
                                try:
                                    os.system('clear')
                                    return spiders.harvester(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Start Host Root File Bruteforcer
                            elif choice == '4':
                                try:
                                    os.system('clear')
                                    return spiders.gobuster(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            #Start SQL, XSS & SSRF Detection & Eploitation
                            elif choice == '5':
                                try:
                                    os.system('clear')
                                    def web_injector(self):
                                        mega_menu.menu_eight_five()
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Sql injection attacks againist" + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
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
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 4" + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                            #Jok3r v3.5 Insane Reconing Next Generation Engine
                            elif choice == '6':
                                try:
                                    os.system('clear')
                                    def jok3r(self):
                                        mega_menu.menu_eight_six()
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Launch an attack againist" + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Launch Osmedeus tool againist" + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                                        print(bcolors.BLUE + "  ~>[ " + bcolors.RED + "Launch DDOS attacks againist" + bcolors.BLUE + " }" + bcolors.BLUE + " ~> " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)

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
                                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                                    return spiders.wafw00f(host), spiders.seeker(xhost), spiders.whatweb(host), spiders.gobuster(host), spiders.osmedeus_3(host), spiders.param_spider(host), spiders.ssl_scan(zhost), spiders.nuclei(host), spiders.sqli_auto_sqlmap(host), spiders.xss_auto_commix(host), spiders.xss_auto_katana(host), spiders.xss_auto_xsser(host), spiders.bbot(host), spiders.nikto(host), spiders.uniscan(host), attack_websites.user_nuke(self)
                                except:
                                    user_nuke(self)
                                    break
                            else:
                                try:
                                    print("\n")
                                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + " ]<~" + bcolors.ENDC
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
                salvation =  bcolors.DARKCYAN + " For God so loved the world, that He gave." + bcolors.ENDC + color() + bcolors.BOLD + " <~[ John 3:16" + bcolors.ENDC
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
                    choice = input(bcolors.GREEN + "(" + bcolors.ENDC + bcolors.BOLD + "africana:" + bcolors.ENDC + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
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
