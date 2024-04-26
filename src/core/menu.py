from src.core.banner import *
from src.core.bcolors import *
from src.scriptures.verses import *

class main_menu(object):
    def __init_(self):
        pass

    def menu_zero(self):
        print(bcolors.RED + "\n[ The 🍄Africana-Framework is designed purely for Good &..] " + bcolors.ENDC)
        print(bcolors.RED + "[ not evil. If you are planning on using this tool for..🐝] " + bcolors.ENDC)
        print(bcolors.RED + "[ malicious purposes that are not authorized by the.......] " + bcolors.ENDC)
        print(bcolors.RED + "[ 🥝company you are performing assessments for, you are...] " + bcolors.ENDC)
        print(bcolors.RED + "[ violating the terms of service and license of this......] " + bcolors.ENDC)
        print(bcolors.RED + "[ toolset. By hitting yes (only one time), you agree to...] " + bcolors.ENDC)
        print(bcolors.RED + "[ the terms of service and that you will only use this....] " + bcolors.ENDC)
        print(bcolors.RED + "[ tool for lawful purposes only.........................🥥] \n" + bcolors.ENDC)

    def menu_one(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ 🍄" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install or update africana-framework.." + bcolors.DARKCYAN + "(Start here )" + bcolors.BLUE + "🐞] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. System Security Configuration........." + bcolors.DARKCYAN + "(Setup tor &)" + bcolors.BLUE + "🐈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Local Network Attack Vectors.........." + bcolors.DARKCYAN + "(Mitm, sniff)" + bcolors.BLUE + "🐹] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Generate Undetectable Backdoors......." + bcolors.DARKCYAN + "(C2 & shells)" + bcolors.BLUE + "🐭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. WiFi Attack Vectors..................." + bcolors.DARKCYAN + "(Wifite, air)" + bcolors.BLUE + "🦝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Crack Hash, Pcap & Brute Passwords...." + bcolors.DARKCYAN + "(Hashcat, jo)" + bcolors.BLUE + "🐙] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Social-Engineering Attacks............" + bcolors.DARKCYAN + "(Gophish, gi)" + bcolors.BLUE + "🧪] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Website Attack Vectors................" + bcolors.DARKCYAN + "(Osmedeus, j)" + bcolors.BLUE + "🌍] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Help, Credits, Tricks and About......." + bcolors.DARKCYAN + "(🕊️  " + bcolors.RED + "︻╦╤─" + bcolors.GREEN + "JC❤️sU" + bcolors.DARKCYAN + ")" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit africana-framework..............." + bcolors.DARKCYAN + "(Try option " + bcolors.YELLOW + "99" + bcolors.DARKCYAN + ")" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "                                              ¯\_(ツ)_/¯                                                 " + bcolors.ENDC)

    def menu_two(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ 🎭" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install & Setup " + bcolors.RED + "Tor" + bcolors.BLUE + "......................" + bcolors.DARKCYAN + "(start here)" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Start anonymizing through tor......................🧅]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Stop tor & restore all iptables....................🐝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Chains....." + bcolors.GREEN + "local 🐌 " + bcolors.GREEN + "squid 🐙 " + bcolors.GREEN + "privoxy 🎭 " + bcolors.GREEN + "tor 🧄 " + bcolors.GREEN + "WEB 🌍" + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Check if using " + bcolors.GREEN + "tor 🧄 " + bcolors.BLUE + ".............................🦨] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit and go back fo main menu........................] \n" + bcolors.ENDC)

    def menu_three(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Discover Targets..................................🐹] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Port discovery on the target......................🐾] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Vuln' Scann the Target............................🦉] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. S.M.B enumration on the target....................🪳] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Exploit S.M.B on target...........................🐞] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Sniff Internal Packets............................🪰] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Lunch Responder with IPv6 Support.................🍈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start Beefxss & Bettercap For (M.I.B).............🐝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Get Reverse shell through snifing (XSS Injection).🍇] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go to main menu..............................] \n" + bcolors.ENDC)

    def menu_four(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Blackjack......................" + bcolors.DARKCYAN + "(All Distro...R.A.T)" + bcolors.BLUE + "🐚] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. ShellzGen......................" + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🦐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. PowerJoker....................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🍐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. MeterPeter....................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Havoc C2......................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Teardroid......................" + bcolors.DARKCYAN + "(Android 4 > 13 Rat)" + bcolors.BLUE + "🥙] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. AndroidRAT....................." + bcolors.DARKCYAN + "(Android 4 > 10 Rat)" + bcolors.BLUE + "🐭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Chameleon......................" + bcolors.DARKCYAN + "(Obfuscation .Ps1 S)" + bcolors.BLUE + "🍤] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Gh0x0st........................" + bcolors.DARKCYAN + "(Obfuscation .Ps1 S)" + bcolors.BLUE + "🍈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ 📶" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ......................Wifite............" + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🎯] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. ....................Bettercap..........." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🧄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. ...................Wifipumpkin3" + bcolors.DARKCYAN + "(Automated Phishing)" + bcolors.BLUE + "🍍] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. ....................Airgeddon.............." + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. ...................wifiPumpkin3............" + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐚] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)

    def menu_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Online......................." + bcolors.DARKCYAN + "(automated_Bruteforce)" + bcolors.BLUE + "🍋] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Offline................." + bcolors.DARKCYAN + "(automated/ maunal/ Hashes)" + bcolors.BLUE + "🥭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_six_one(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n      ~>[ 🔐" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ...................Bruteforce " + bcolors.DARKCYAN + "SSH" + bcolors.BLUE + "..................🫑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. ...................Bruteforce " + bcolors.DARKCYAN + "FTP" + bcolors.BLUE + "..................🔑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. ...................Bruteforce " + bcolors.DARKCYAN + "SMB" + bcolors.BLUE + "..................🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. ...................Bruteforce " + bcolors.DARKCYAN + "RDP" + bcolors.BLUE + "..................🍒] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. ...................Bruteforce " + bcolors.DARKCYAN + "LDAP" + bcolors.BLUE + ".................🧉] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. ...................Bruteforce " + bcolors.DARKCYAN + "SMTP" + bcolors.BLUE + ".................🌈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. ...................Bruteforce " + bcolors.DARKCYAN + "Telnet" + bcolors.BLUE + "...............🐚] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. ...................Bruteforce " + bcolors.DARKCYAN + "HTTP/S" + bcolors.BLUE + "...............🪀] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. ...................Bruteforce " + bcolors.DARKCYAN + "All/SS" + bcolors.BLUE + "...............🩴] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)

    def menu_six_two(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n      ~>[ 🔐" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ....................Aircrack_ng........." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🫑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. .......................John............." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🥭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. ......................To Add.......................🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. ....................Hash-Buster....................🫑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)

    def menu_seven(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n      ~>[ 🍄" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ..............Gophish.(Browser Gui)." + bcolors.DARKCYAN + "(All Templetes)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. ...............Good Ginx (Advanced)." + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🍹] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. ...................AdvPhishing......" + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🦠] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. ....................Setoolkit......." + bcolors.DARKCYAN + "(Web Cloning..)" + bcolors.BLUE + "🧉] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. ....................Anonphisher....." + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. ....................Cyberphish......" + bcolors.DARKCYAN + "(phish temps..)" + bcolors.BLUE + "🫑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. ...............Exit & Go To Main Menu................] \n" + bcolors.ENDC)

    def menu_eight(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ 🦟" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Start Passive Web recon & Subdomain Enumration.....🌍] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Gather e-mails & subdomain namesfrom public sources🪰] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Start Bruteforcing Host's Root Files...............🚀] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Start SQL, XSS & SSRF Detection & Eploitation......🐌] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Launch OWASP Nettacker project MainMenu............🦣] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu...👊] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Osmedeus Next Generation Workflow Engine Main Menu.🍈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Ufonet Next Generation DDOS Tool Main Menu.........🦠] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Heavy Automation Attacks On The Host........🍄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit Web Scanner & Go To Main Menu...................] \n" + bcolors.ENDC)

    def menu_eight_four(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n      ~>[ " + bcolors.ENDC + bcolors.UNDERL + "💉Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. ............Sql injection with sqlmap..." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🐞] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. .......Xss injection with Dalfox, Xsser." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🪰] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. ............Sql injection with Sqlmap......" + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐛] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. ............Xss injection with Xsser......." + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐌] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. .............Exit & Go To Main Menu..................] \n" + bcolors.ENDC)

    def menu_eight_five(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Start Nettacker port & web content discovery.......📡] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Recon, find and scan subdomains....................🐾] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch admin_scan to find admin panel..............🦨] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Start insane information gathering on host.........🧄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks...........🧭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Run CVE scans on the target host...................🍹] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)......🍄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Automate all modules & security checks on target...🥑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. View scanned Nettacker report list.................🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_eight_six(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Install all jok3r tools..." + bcolors.DARKCYAN + "(Pleas start here if not)" + bcolors.BLUE + "🃏] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Update all the tools in the toolbox................🍄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Show all the tools in the toolbox..................🍒] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Show supported products for all services...........🍵] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks...........🧭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Run all & intense security checks against an URL...🦠] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)......🌈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. View the full resultsfrom the security checks......🧄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Clean database & delete results....................🥑] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_eight_seven(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Update Osmedeus & Run diagnostics to check config..🍄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Start a simple scan with other flow................🦠] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch vuln and directory scan on domains..........🧭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Scan list of targets (Full path of target needed)..🥐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Cloud - Run scan in Distributed Cloud mode.........🌏] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Performing Full web vuln & secret scan on host.....🏹] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Update the vulnerability database before attacking.📜] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start web UI server................................🪐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. View scanned osmedeus report list..................🌈] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_eight_eight(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Download list of " + bcolors.RED + "Bots " + bcolors.BLUE + "from" + bcolors.YELLOW + " Community " + bcolors.BLUE + "server.......🍄.]" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Test if bots are alive............................📡.] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Launch palantir..................." + bcolors.DARKCYAN + "(Palantir 3.14..)" + bcolors.BLUE + "🧭] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Launch socking_waves.............." + bcolors.DARKCYAN + "(Knockout!......)" + bcolors.BLUE + "🦠] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 5. Launch xcom-1....................." + bcolors.DARKCYAN + "(Only DDoS......)" + bcolors.BLUE + "🥐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 6. Launch xcom-2....................." + bcolors.DARKCYAN + "(Only DoS.......)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 7. Launch ufonet-gui................." + bcolors.DARKCYAN + "(Gui on browser.)" + bcolors.BLUE + "🧄] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 8. Start Grider.............." + bcolors.DARKCYAN + "........(Grider.........)" + bcolors.BLUE + "🪐] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 9. Launch Armageddon!........" + bcolors.DARKCYAN + "........(Launch ALL!!...)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)

    def menu_nine(self):
        beauty.graphics(), scriptures.verses()
        print(bcolors.ENDC + "\n[ I am Rojahs Montari a Devoted Christian & Pentester.....] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ One might describe me as an erudite.....................] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ & perspicacious individual, a connoisseur of cybernetic ] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ security and digital fortification. My acumen in........] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ devising the 🦠 africana-framework bespeaks a sagacious.] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ approach to ethological hacking and vulnerability.......] " + bcolors.ENDC)
        print(bcolors.ENDC + "[ reconnaissance. As a virtuoso of cryptographic endeavors] \n" + bcolors.ENDC)
        print(bcolors.RED  + "[ My pedagogical content disseminates esoteric knowledge,.] " + bcolors.ENDC)
        print(bcolors.RED  + "[ fostering a coterie of aspirant ethical hackers. My.....] " + bcolors.ENDC)
        print(bcolors.RED  + "[ prolific contributions to the technological milieu......] " + bcolors.ENDC)
        print(bcolors.RED  + "[ underscore a quintessential commitment to advancing.....] " + bcolors.ENDC)
        print(bcolors.RED  + "[ cybersecurity paradigms.................................] " + bcolors.ENDC)
        print(bcolors.RED  + "[ Email.......💬: rojahsmontari@gmail.com.................] " + bcolors.ENDC)
        print(bcolors.RED  + "[ YouTube.....💎: https://youtube.com/@RojahsMontari......] " + bcolors.ENDC)
        print(bcolors.RED  + "[ Life Tip....📚: " + bcolors.BLUE + "Defeat the devil by " + bcolors.YELLOW + "fasting & praying" + bcolors.RED  + "...] " + bcolors.ENDC)
        print(bcolors.GREEN + "                                         ¯\_(ツ)_/¯"          + bcolors.ENDC)
        print(bcolors.BLUE + "[ ................." + bcolors.DARKCYAN + "Press Enter To go BACK." + bcolors.BLUE + "................] \n" + bcolors.ENDC)

    def menu_nine_tisa(self):
        print(bcolors.GREEN + """
            __ 🕊              _____ _____     _     _
         __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
        |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
        |_____|___|___|___|___|_____|__|__|_| |_|___|_| 🕊
""" + bcolors.ENDC)
        print(bcolors.BLUE + "                ~>[ 📜" + bcolors.ENDC + "The Ten Commandments." + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 1. You shall have no other gods before Me................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 2. You shall make no idols...............................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 3. You shall not take the name of the Lord your God in...] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[  . vein..................................................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 4. Keep the Sabbath day holy.............................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 5. Honor your father and your mother.....................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 6. You shall not murder..................................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 7. You shall not commit adultery.........................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 8. You shall not steal...................................] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 9. You shall not bear false witness against your eighbor.] " + bcolors.ENDC)
        print(bcolors.YELLOW + "[ 10 You shall not covet...................................] \n" + bcolors.ENDC)
        print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Tip:🕊️ " + bcolors.RED + "︻╦╤─ " + bcolors.DARKCYAN + "Defeat the devil by fasting & praying" + bcolors.GREEN + ") \n" + bcolors.ENDC)

mega_menu = main_menu()

if ' __name__' == '__main__':
        sys.exit(mega_menu())
