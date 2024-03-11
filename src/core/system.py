import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *
from src.scriptures.verses import *

class update(object):
    def __init__(self):
        pass

    def update_system(self):
        os.system('clear')
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n          ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1.                      KaliLinux        " + bcolors.RED + "(It is Stable) " + bcolors.BLUE + "] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2.                     UbuntuLinux                      ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3.         ArchLinux/ BlackArchLinux/ ManjaroLinux      ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4.                   Uninstall Africana                 ] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0.                         Exit                         ] \n" + bcolors.ENDC)
        try:
            choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        except:
            pass
        while True:
            try:
                if choice == '0':
                    neo.one()
                    break
                elif choice == '1':
                    os.system('clear')
                    print("\n")
                    beauty.graphics(), scriptures.verses()
                    print(bcolors.BLUE + "[            Installing africana on kali-linux            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                    Pleas be patient                     ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[      Installer will copy core files to your system      ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[              All necessary foundation tools             ] \n" + bcolors.ENDC)
                    print(bcolors.ENDC + "      ~>[ " + bcolors.RED + bcolors.UNDERL + " apt-get install -y git zsh python3 python3-pip. " + bcolors.ENDC + "]<~\n" + bcolors.ENDC)

                    afric = '/usr/local/opt/africana-framework'
                    if os.path.exists(afric):
                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[              Africana detected in your system           ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                 Installer runs full updates             ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[       Upgrades for all necessary foundation tools       ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "               {" + bcolors.RED + bcolors.UNDERL + " git pull & apt-get upgrade -y.. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Africana is already installed in your system. Updating It. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        os.system('apt-get update -y')
                        print("\n")
                        os.system('cd /usr/local/opt/africana-framework; git pull; cd src/externals/Teardroid-phprat; git pull; cd ../AdvPhishing; git pull; cd ../ufonet; git pull; cd ../anonphisher; git pull')
                        time.sleep(0.09)

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[              Setting up project discovery tools         ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                         Installer                       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                          Configs                        ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                  All golang necessary tools             ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Updating all Project Discovery tools & testing them for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        afric = '/root/go'
                        if os.path.exists(afric):
                            os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc')
                        else:
                            pass

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print("\n")
                        print(bcolors.BLUE + "[                 Necessary tools updateded               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[      Anonymous tools updateded (Tor, privoxy, squid)    ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[          Wifi pentesting tools updateded (wifite)       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[        Local pentesting tools updateded (bettercap)     ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[  Web pentest tools updateded (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Everything is set. Type 'africana' to launch the framework " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        break

                    else:
                        os.system('mkdir -p /usr/local/opt; cp -rf ../../africana-framework /usr/local/opt/; ln -s /usr/local/opt/africana-framework/src/modules/kenyan.py /usr/local/bin/africana; chmod +x /usr/local/bin/africana; apt-get update -y; apt-get install zsh git curl -y; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -SL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg;apt-get install curl wget git zip -y; curl -SL -o /etc/apt/sources.list.d/playit-cloud.list https://playit-cloud.github.io/ppa/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables tmux openssh-client openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser villain set playit wine32:i386')

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[           Setting up Go, Python env @ '~/.afric'        ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                        Setup runs                       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[              Source @ ~/.afric/bin/activate             ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/; pip3 install virtualenv; virtualenv " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Creating Go, Python env @ '~/.afric' & activating it for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        if os.path.exists('~/.zshrc.bak_africana'):
                            pass
                        elif not os.path.exists('~/.zshrc.bak_africana'):
                            os.system('cd ~/; cp .zshrc .zshrc.bak_africana; pip3 install --upgrade setuptools; pip3 install virtualenv; python3 -m virtualenv .afric --system-site-packages; echo -n "\n# Generated by africana-framework. Delete at your own risk!\nsource ~/.afric/bin/activate" >> ~/.zshrc')
                        else:
                            pass

                        if os.path.exists('/etc/sysctl.conf.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/sysctl.conf.bak_africana'):
                            infile = '/etc/sysctl.conf'
                            outfile = '/etc/sysctl.conf'
                            delete_list = ['#net.ipv4.ip_forward=1']
                            fin = open(infile)
                            os.system('mv /etc/sysctl.conf /etc/sysctl.conf.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'net.ipv4.ip_forward=1')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass

                        if os.path.exists('/etc/default/grub.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/default/grub.bak_africana'):
                            infile = '/etc/default/grub'
                            outfile = '/etc/default/grub'
                            delete_list = ['GRUB_CMDLINE_LINUX_DEFAULT="quiet"']
                            delete_list1 = ['#GRUB_TERMINAL=console']
                            fin = open(infile)
                            os.system('mv /etc/default/grub /etc/default/grub.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'GRUB_CMDLINE_LINUX_DEFAULT=""')
                                for word in delete_list1:
                                    line = line.replace(word, 'GRUB_TERMINAL=console')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass

                        if os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                            infile = '/etc/default/grub.d/kali-themes.cfg'
                            outfile = '/etc/default/grub.d/kali-themes.cfg'
                            delete_list = ['if ! echo "$GRUB_CMDLINE_LINUX_DEFAULT" | grep -q splash; then']
                            delete_list1 = ['    GRUB_CMDLINE_LINUX_DEFAULT="$GRUB_CMDLINE_LINUX_DEFAULT splash"']
                            delete_list2 = ['fi']
                            fin = open(infile)
                            os.system('mv /etc/default/grub.d/kali-themes.cfg /etc/default/grub.d/kali-themes.cfg.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, '#if ! echo "$GRUB_CMDLINE_LINUX_DEFAULT" | grep -q splash; then')
                                for word in delete_list1:
                                    line = line.replace(word, '#    GRUB_CMDLINE_LINUX_DEFAULT="$GRUB_CMDLINE_LINUX_DEFAULT splash"')
                                for word in delete_list2:
                                    line = line.replace(word, '#fi')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass
                        if os.path.exists('/etc/grub.d/06_dark_theme.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/grub.d/06_dark_theme.bak_africana'):
                            infile = '/etc/grub.d/06_dark_theme'
                            outfile = '/etc/grub.d/06_dark_theme'
                            delete_list = ['echo "set menu_color_normal=green/black"']
                            delete_list1 = ['echo "set menu_color_highlight=white/black"']
                            delete_list2 = ['echo "set color_normal=white/black"']
                            delete_list3 = ['echo "set color_highlight=green/black"']
                            fin = open(infile)
                            os.system('mv /etc/grub.d/06_dark_theme /etc/grub.d/06_dark_theme.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'echo "set menu_color_normal=green/black"')
                                for word in delete_list1:
                                    line = line.replace(word, 'echo "set menu_color_highlight=white/black"')
                                for word in delete_list2:
                                    line = line.replace(word, 'echo "set color_normal=white/black"')
                                for word in delete_list3:
                                    line = line.replace(word, 'echo "set color_highlight=green/black"')
                                fout.write(line)
                            fin.close()
                            fout.close()
                            os.system('chmod +x /etc/grub.d/06_dark_theme')
                        else:
                            pass

                        git = '/usr/local/opt/africana-framework/externals'
                        if os.path.exists(git):
                            try:
                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[                 Virtualenv installed & Set              ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                   Go env installed & Set                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                  sysctl.conf already set                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                /etc/default/grub configured             ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                 kali-themes.cfg already set             ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "     {" + bcolors.RED + bcolors.UNDERL + " git clone. & cd ~/; pip3 install -r requirements.txt " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "  ~>[ " + bcolors.YELLOW + " Installing africana-framework' s Python3 requirements. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('zsh -c "source ~/.zshrc; cd /usr/local/opt/africana-framework; pip3 install -r requirements.txt"')

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[                Pip3 core tools installed eg,            ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 1.                         bbot                         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 2.                         pipx                         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 3.                         pproxy                       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 4.                         colorama (etc.)              ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "     { " + bcolors.RED + bcolors.UNDERL + "git clone. & cd ~/; pip3 install -r requirements.txt " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + " ~>[  " + bcolors.YELLOW + "Installing and setting up third party tools from github. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('cd /usr/local/opt/africana-framework/src/externals/; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; cd ..; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; git clone https://github.com/Ignitetch/AdvPhishing.git; cd AdvPhishing/; chmod 777 *; ./Linux-Setup.sh; cd ..;git clone https://github.com/TermuxHackz/anonphisher; cd ./anonphisher; chmod 777 *; bash anonphisher.sh')

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[              Setting up project discovery tools         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                   Pleas Be Patient as the               ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                      Installer runs                     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[              Sets up all golang necessary tools         ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Updating all Project Discovery tools & testing them for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('go install github.com/projectdiscovery/pdtm/cmd/pdtm@latest; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)')
                                afric = '/root/go'
                                if os.path.exists(afric):
                                    os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc; update-grub2')
                                else:
                                    africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + "cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. Command failed to execute" + bcolors.ENDC + "]<~" + bcolors.ENDC
                                    for a in africana:
                                        sys.stdout.write(a)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                print("\n")
                                os.system('update-grub2')
                                pass

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print("\n")
                                print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Everything is set. Type 'africana' to launch the framework " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                            except:
                                break

#UbuntuLinux

                elif choice == '2':
                    os.system('clear')
                    print("\n")
                    beauty.graphics(), scriptures.verses()
                    print(bcolors.BLUE + "[            Installing africana on kali-linux            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                    Pleas be patient                     ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[      Installer will copy core files to your system      ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[              All necessary foundation tools             ] \n" + bcolors.ENDC)
                    print(bcolors.ENDC + "      ~>[ " + bcolors.RED + bcolors.UNDERL + " apt-get install -y git zsh python3 python3-pip. " + bcolors.ENDC + "]<~\n" + bcolors.ENDC)

                    afric = '/usr/local/opt/africana-framework'
                    if os.path.exists(afric):
                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[              Africana detected in your system           ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                 Installer runs full updates             ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[       Upgrades for all necessary foundation tools       ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "               {" + bcolors.RED + bcolors.UNDERL + " git pull & apt-get upgrade -y.. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Africana is already installed in your system. Updating It. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        os.system('apt-get update -y')
                        print("\n")
                        os.system('cd /usr/local/opt/africana-framework; git pull; cd src/externals/Teardroid-phprat; git pull; cd ../AdvPhishing; git pull; cd ../ufonet; git pull; cd ../anonphisher; git pull')
                        time.sleep(0.09)

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[              Setting up project discovery tools         ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                         Installer                       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                          Configs                        ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                  All golang necessary tools             ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Updating all Project Discovery tools & testing them for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        afric = '/root/go'
                        if os.path.exists(afric):
                            os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc')
                        else:
                            pass

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print("\n")
                        print(bcolors.BLUE + "[                 Necessary tools updateded               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[      Anonymous tools updateded (Tor, privoxy, squid)    ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[          Wifi pentesting tools updateded (wifite)       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[        Local pentesting tools updateded (bettercap)     ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[  Web pentest tools updateded (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Everything is set. Type 'africana' to launch the framework " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        break

                    else:
                        os.system('mkdir -p /usr/local/opt; cp -rf ../../africana-framework /usr/local/opt/; ln -s /usr/local/opt/africana-framework/src/modules/kenyan.py /usr/local/bin/africana; chmod +x /usr/local/bin/africana; echo -n "deb https://http.kali.org/kali kali-rolling main non-free contrib" >> /etc/apt/sources.list.d/kali.list; apt-get update -y; apt-get install wget zip tmux zsh git curl -y; wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc; echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -SL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg; curl -SL -o /etc/apt/sources.list.d/playit-cloud.list https://playit-cloud.github.io/ppa/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables openssh-client openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser villain set playit wine32:i386')

                        os.system('clear')
                        print("\n")
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[           Setting up Go, Python env @ '~/.afric'        ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                        Setup runs                       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[              Source @ ~/.afric/bin/activate             ] \n" + bcolors.ENDC)
                        print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/; pip3 install virtualenv; virtualenv " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                        africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Creating Go, Python env @ '~/.afric' & activating it for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                        for a in africana:
                            sys.stdout.write(a)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        print("\n")
                        if os.path.exists('~/.zshrc.bak_africana'):
                            pass
                        elif not os.path.exists('~/.zshrc.bak_africana'):
                            os.system('cd ~/; cp .zshrc .zshrc.bak_africana; pip3 install --upgrade setuptools; pip3 install virtualenv; python3 -m virtualenv .afric --system-site-packages; echo -n "\n# Generated by africana-framework. Delete at your own risk!\nsource ~/.afric/bin/activate" >> ~/.zshrc')
                        else:
                            pass

                        if os.path.exists('/etc/sysctl.conf.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/sysctl.conf.bak_africana'):
                            infile = '/etc/sysctl.conf'
                            outfile = '/etc/sysctl.conf'
                            delete_list = ['#net.ipv4.ip_forward=1']
                            fin = open(infile)
                            os.system('mv /etc/sysctl.conf /etc/sysctl.conf.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'net.ipv4.ip_forward=1')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass

                        if os.path.exists('/etc/default/grub.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/default/grub.bak_africana'):
                            infile = '/etc/default/grub'
                            outfile = '/etc/default/grub'
                            delete_list = ['GRUB_CMDLINE_LINUX_DEFAULT="quiet"']
                            delete_list1 = ['#GRUB_TERMINAL=console']
                            fin = open(infile)
                            os.system('mv /etc/default/grub /etc/default/grub.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'GRUB_CMDLINE_LINUX_DEFAULT=""')
                                for word in delete_list1:
                                    line = line.replace(word, 'GRUB_TERMINAL=console')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass

                        if os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                            infile = '/etc/default/grub.d/kali-themes.cfg'
                            outfile = '/etc/default/grub.d/kali-themes.cfg'
                            delete_list = ['if ! echo "$GRUB_CMDLINE_LINUX_DEFAULT" | grep -q splash; then']
                            delete_list1 = ['    GRUB_CMDLINE_LINUX_DEFAULT="$GRUB_CMDLINE_LINUX_DEFAULT splash"']
                            delete_list2 = ['fi']
                            fin = open(infile)
                            os.system('mv /etc/default/grub.d/kali-themes.cfg /etc/default/grub.d/kali-themes.cfg.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, '#if ! echo "$GRUB_CMDLINE_LINUX_DEFAULT" | grep -q splash; then')
                                for word in delete_list1:
                                    line = line.replace(word, '#    GRUB_CMDLINE_LINUX_DEFAULT="$GRUB_CMDLINE_LINUX_DEFAULT splash"')
                                for word in delete_list2:
                                    line = line.replace(word, '#fi')
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass
                        if os.path.exists('/etc/grub.d/06_dark_theme.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/grub.d/06_dark_theme.bak_africana'):
                            infile = '/etc/grub.d/06_dark_theme'
                            outfile = '/etc/grub.d/06_dark_theme'
                            delete_list = ['echo "set menu_color_normal=green/black"']
                            delete_list1 = ['echo "set menu_color_highlight=white/black"']
                            delete_list2 = ['echo "set color_normal=white/black"']
                            delete_list3 = ['echo "set color_highlight=green/black"']
                            fin = open(infile)
                            os.system('mv /etc/grub.d/06_dark_theme /etc/grub.d/06_dark_theme.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, 'echo "set menu_color_normal=green/black"')
                                for word in delete_list1:
                                    line = line.replace(word, 'echo "set menu_color_highlight=white/black"')
                                for word in delete_list2:
                                    line = line.replace(word, 'echo "set color_normal=white/black"')
                                for word in delete_list3:
                                    line = line.replace(word, 'echo "set color_highlight=green/black"')
                                fout.write(line)
                            fin.close()
                            fout.close()
                            os.system('chmod +x /etc/grub.d/06_dark_theme')
                        else:
                            pass

                        git = '/usr/local/opt/africana-framework/externals'
                        if os.path.exists(git):
                            try:
                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[                 Virtualenv installed & Set              ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                   Go env installed & Set                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                  sysctl.conf already set                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                /etc/default/grub configured             ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                 kali-themes.cfg already set             ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "     {" + bcolors.RED + bcolors.UNDERL + " git clone. & cd ~/; pip3 install -r requirements.txt " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "  ~>[ " + bcolors.YELLOW + " Installing africana-framework' s Python3 requirements. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('zsh -c "source ~/.zshrc; cd /usr/local/opt/africana-framework; pip3 install -r requirements.txt"')

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[                Pip3 core tools installed eg,            ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 1.                         bbot                         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 2.                         pipx                         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 3.                         pproxy                       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[ 4.                         colorama (etc.)              ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "     { " + bcolors.RED + bcolors.UNDERL + "git clone. & cd ~/; pip3 install -r requirements.txt " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + " ~>[  " + bcolors.YELLOW + "Installing and setting up third party tools from github. " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('cd /usr/local/opt/africana-framework/src/externals/; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; cd ..; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; git clone https://github.com/Ignitetch/AdvPhishing.git; cd AdvPhishing/; chmod 777 *; ./Linux-Setup.sh; cd ..;git clone https://github.com/TermuxHackz/anonphisher; cd ./anonphisher; chmod 777 *; bash anonphisher.sh')

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print(bcolors.BLUE + "[              Setting up project discovery tools         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                   Pleas Be Patient as the               ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                      Installer runs                     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[              Sets up all golang necessary tools         ] \n" + bcolors.ENDC)
                                print(bcolors.ENDC + "          {" + bcolors.RED + bcolors.UNDERL + " cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. " + bcolors.ENDC + "}\n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Updating all Project Discovery tools & testing them for you " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                print("\n")
                                os.system('go install github.com/projectdiscovery/pdtm/cmd/pdtm@latest; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)')
                                afric = '/root/go'
                                if os.path.exists(afric):
                                    os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc; update-grub2')
                                else:
                                    africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + "cd ~/go/bin; ./pdtm -ia; source ~/.zshrc. Command failed to execute" + bcolors.ENDC + "]<~" + bcolors.ENDC
                                    for a in africana:
                                        sys.stdout.write(a)
                                        sys.stdout.flush()
                                        time.sleep(0.09)
                                print("\n")
                                os.system('update-grub2')
                                pass

                                os.system('clear')
                                print("\n")
                                beauty.graphics(), scriptures.verses()
                                print("\n")
                                print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                africana = bcolors.ENDC + "~>[ " + bcolors.YELLOW + " Everything is set. Type 'africana' to launch the framework " + bcolors.ENDC + "]<~" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                            except:
                                break
## Uninstaller
                elif choice == '4':
                    try:
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        print("\n")
                        print(bcolors.BLUE + "[   Incase of any bug email me @: " + bcolors.YELLOW + "rojahsmontari@gmail.com " + bcolors.BLUE + "    ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[  Are you sure you want to completely uninstall africana?    ]\n" + bcolors.ENDC)
                        print(bcolors.ENDC + "         ~>[  " + bcolors.BLUE + "Type: " + bcolors.RED + "1. " + bcolors.GREEN + "Accept " + bcolors.RED + "0. " + bcolors.GREEN + "Reject " + bcolors.YELLOW + "& " + bcolors.BLUE + "Go Back To Menu." + bcolors.ENDC + " ]~\n" + bcolors.ENDC)

                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        if choice == '1':
                            os.system('systemctl --no-pager status privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service -l')
                            print("\n")
                            os.system('systemctl stop privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service; systemctl disable privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service')
                            if os.path.exists('~/.zshrc.bak_africana'):
                                os.system('cd ~/; rm -rf .zshrc; mv .zshrc.bak_africana .zshrc')
                            else:
                                pass
                            if os.path.exists('/etc/sysctl.conf.bak_africana'):
                                os.system('cd /etc/; rm -rf sysctl.conf; mv sysctl.conf.bak_africana sysctl.conf')
                            else:
                                pass
                            if os.path.exists('/etc/tor/torrc.bak_africana'):
                                os.system('cd /etc/tor/; rm -rf torrc; mv torrc.bak_africana torrc')
                            else:
                                pass
                            if os.path.exists('/lib/systemd/system/tor@default.service.bak_africana'):
                                os.system('cd /lib/systemd/system/; rm -rf tor@default.service; mv tor@default.service.bak_africana tor@default.service')
                            else:
                                pass
                            if os.path.exists('/etc/systemd/system/changemac@.service'):
                                os.system('rm -rf /etc/systemd/system/changemac@.bak_africana')
                            else:
                                pass
                            if os.path.exists('/etc/systemd/system/changemac@.bak_africana'):
                                os.system('rm -rf /etc/systemd/system/changemac@.bak_africana')
                            else:
                                pass
                            if os.path.exists('/etc/squid/squid.conf.bak_africana'):
                                os.system('cd /etc/squid/; rm -rf squid.conf; mv squid.conf.bak_africana squid.conf')
                            else:
                                pass
                            if os.path.exists('/etc/privoxy/privoxy.bak_africana'):
                                os.system('cd /etc/privoxy/; rm -rf config; mv privoxy.bak_africana config')
                            else:
                                pass
                            if os.path.exists('/etc/dhcp/dhclient.conf.bak_africana'):
                                os.system('cd /etc/dhcp/; rm -rf dhclient.conf; mv dhclient.conf.bak_africana dhclient.conf')
                            else:
                                pass
                            if os.path.exists('~/.zshrc.bak_africana'):
                                os.system('cd ~/; rm -rf .zshrc; mv .zshrc.bak_africana .zshrc')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.bak_africana'):
                                os.system('cd /etc/default/; rm -rf grub; mv grub.bak_africana grub')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                                os.system('cd /etc/default/grub.d/; rm -rf kali-themes.cfg; mv kali-themes.cfg.bak_africana kali-themes.cfg')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.d/06_dark_theme.bak_africana'):
                                os.system('cd /etc/default/grub.d/; rm -rf 06_dark_theme; mv 06_dark_theme.bak_africana 06_dark_theme')
                            else:
                                pass
                            afric = '/usr/local/opt/africana-framework'
                            if os.path.exists(afric):
                                os.system('rm -rf /usr/local/opt/africana-framework; rm -rf /usr/local/bin/africana')
                            else:
                                pass
                            os.system('clear')
                            beauty.graphics(), scriptures.verses()
                            africana = bcolors.ENDC + "              ~{ " + bcolors.YELLOW + "africana-framework uninstalled. " + bcolors.ENDC +  "}~" + bcolors.ENDC
                            for a in africana:
                                sys.stdout.write(a)
                                sys.stdout.flush()
                                time.sleep(0.03)
                            break
                        elif choice == '0':
                            installer.update_system()
                            break
                        else:
                            print("\n")
                            warn = bcolors.ENDC + "   ~{ " + bcolors.RED + "Poor choice of selection!. Please select int -> " + bcolors.DARKCYAN + "0. or 1. " + bcolors.ENDC +  "}~" + bcolors.ENDC
                            for w in warn:
                                sys.stdout.write(w)
                                sys.stdout.flush()
                                time.sleep(0.09)
                            installer.update_system()
                            break
                    except:
                        neo.one()
                        break
                else:
                    try:
                        print("\n")
                        warn = bcolors.ENDC + "   ~{ " + bcolors.RED + "Poor choice of selection!. Please select int -> " + bcolors.DARKCYAN + "0. to 4. " + bcolors.ENDC +  "}~" + bcolors.ENDC
                        for w in warn:
                            sys.stdout.write(w)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        update_system()
                        pass
                    except:
                        os.system('clear')
                        installer.update_system()
                        break
            except:
                neo.one()
                break

installer = update()
if ' __name__' == '__main__':
        sys.exit(installer())
