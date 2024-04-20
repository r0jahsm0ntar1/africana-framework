import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *
from src.core.spinner import *
from src.scriptures.verses import *

class update(object):
    def __init__(self):
        pass
    def definepath(self):
        return os.getcwd()

    def update_system(self):
        os.system('clear')
        beauty.graphics(), scriptures.verses()
        print(bcolors.BLUE + "\n        ~>[ 🍄" + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        print(bcolors.BLUE + "[ 1. Kali-Linux..........................." + bcolors.DARKCYAN + "(It is Stable)" + bcolors.BLUE + "🍒] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 2. Ubuntu-Linux.......................................🥝] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 3. Arch-Linux/ Black-Arch-Linux/ Manjaro-Linux..........] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 4. Uninstall-Africana...................................] " + bcolors.ENDC)
        print(bcolors.BLUE + "[ 0. Exit.................................................] \n" + bcolors.ENDC)
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
                    scriptures.verses()
                    print(bcolors.BLUE + "\n[          🍄Installing africana on kali-linux            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                    Pleas be patient                     ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[      Installer will copy core files to your system      ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[                            &                            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + "[              All necessary foundation tools             ] " + bcolors.ENDC)

                    africs = '/usr/local/opt/africana-framework'
                    if os.path.exists(africs):
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        print(bcolors.BLUE + "[              Africana detected in your system           ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                 Installer runs full updates             ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                              &                          ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[       Upgrades for all necessary foundation tools       ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "\n[" + bcolors.GREEN + "✔" + bcolors.BLUE + "] " + bcolors.GREEN + "Africana-fri already installed in Ur system. " + bcolors.RED + "Update It?" + bcolors.ENDC)
                        choice = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "[y/n]" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        choice += " "
                        if choice[0].lower() == "y":
                            try:
                                print(bcolors.BLUE + "\n[" + bcolors.GREEN + "✔" + bcolors.BLUE + "] " + bcolors.GREEN + "Updating all Project Discovery tools & testing them 4U" + bcolors.ENDC)
                                spinner.run_subprocesu('apt-get update -y')
                                spinner.run_subprocesux('cd /usr/local/opt/africana-framework; git pull; cd src/externals/Teardroid-phprat; git pull; cd ../AdvPhishing; git pull; cd ../anonphisher; git pull')
                                time.sleep(0.09)
                                africs = '/root/go'
                                if os.path.exists(africs):
                                    spinner.run_subprocesx('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc')
                                else:
                                    pass
                                print(bcolors.BLUE + "[                 Necessary tools updateded               ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[      Anonymous tools updateded (Tor, privoxy, squid)    ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[          Wifi pentesting tools updateded (wifite)       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[        Local pentesting tools updateded (bettercap)     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[  Web pentest tools updateded (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                africana = bcolors.BLUE + "\n[" + bcolors.GREEN + "✔" + bcolors.BLUE + "] " + bcolors.GREEN + "Everything is set. Type 'africana' to start africana-fr " + bcolors.BLUE + "]" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                            except:
                                africana = print(bcolors.BLUE + "\n[" + bcolors.GREEN + "✘" + bcolors.BLUE + "] " + bcolors.GREEN + "▄︻̷̿┻̿═━一 An Error detected pleas run the setup again" + bcolors.ENDC)
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                        else:
                            neo.one()
                            break

                    else:
                        print("")
                        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Making Necessary dirs & Copying africana to your system" + bcolors.ENDC)
                        spinner.run_subprocess('mkdir -p /usr/local/opt; cp -r %s /usr/local/opt/africana-framework; ln -s /usr/local/opt/africana-framework/src/modules/kenyan.py /usr/local/bin/africana; chmod +x /usr/local/bin/africana' %(installer.definepath()))
                        print("")
                        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing Essential tools: nc, zsh, wget, curl, python3" + bcolors.ENDC)
                        spinner.run_subprocesx('apt-get update -y; apt-get install zsh git curl -y; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -SL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg;apt-get install curl wget git zip -y; curl -SL -o /etc/apt/sources.list.d/playit-cloud.list https://playit-cloud.github.io/ppa/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux aha set playit libssl-dev gcc hydra wine32:i386; npm install -g kickthemout')
                        time.sleep(0.3)

                        if os.path.exists('~/.zshrc.bak_africana'):
                            pass
                        elif not os.path.exists('~/.zshrc.bak_africana'):
                            print("")
                            print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Configuring python3 environment at ~/.africs/bin/activate" + bcolors.ENDC)
                            spinner.run_subprocesx('cd ~/; cp .zshrc .zshrc.bak_africana; pip3 install --upgrade setuptools; pip3 install virtualenv; python3 -m virtualenv .africs --system-site-packages; echo -n "\n# Generated by africana-framework. Delete at your own risk!\nsource ~/.africs/bin/activate" >> ~/.zshrc')
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

                        git = '/usr/local/opt/africana-framework/'
                        if os.path.exists(git):
                            try:
                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing and Configuring python3 necessary requirements" + bcolors.ENDC)
                                spinner.run_subprocesx('zsh -c "source ~/.zshrc; cd /usr/local/opt/africana-framework; pip3 install -r requirements.txt"')

                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing and Configuring some github third party tools" + bcolors.ENDC)
                                try:
                                    spinner.run_subprocesx('cd /usr/local/opt/africana-framework/src/externals/androrat; apt install ./zipalign_*_amd64.deb; cd ..; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; rm -rf main.zip; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; git clone https://github.com/Ignitetch/AdvPhishing.git; cd AdvPhishing/; chmod 777 *; ./Linux-Setup.sh; cd ..;git clone https://github.com/TermuxHackz/anonphisher; cd ./anonphisher; chmod 777 *; bash anonphisher.sh')
                                except:
                                    pass

                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing Project Discovery tools & testing them" + bcolors.ENDC)
                                spinner.run_subprocesx('go install github.com/projectdiscovery/pdtm/cmd/pdtm@latest; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)')
                                africs = '/root/go'
                                if os.path.exists(africs):
                                    os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc; update-grub2')
                                else:
                                    print("")
                                    print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.RED + "'./pdtm -ia; source ~/.zshrc'. Command failed to execute" + bcolors.ENDC)
                                    time.sleep(0.3)
                                    spinner.run_subprocess('update-grub2')

                                    print("")
                                    print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                    print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Everything is set. Type 'africana' to launch africana" + bcolors.ENDC)
                                    break
                            except:
                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Setup failed. Pleas retrie again and check your internet connection." + bcolors.ENDC)
                                print("")
                                break
                        else:
                            print("")
                            print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                            print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Everything is set. Type 'africana' to launch africana" + bcolors.ENDC)
                            break

#UbuntuLinux
                elif choice == '2':
                    os.system('clear')
                    print(bcolors.BLUE + bcolors.BOLD + "\n[            Installing africana on Ubuntu-linux          ] " + bcolors.ENDC)
                    print(bcolors.BLUE + bcolors.BOLD + "[                    Pleas be patient                     ] " + bcolors.ENDC)
                    print(bcolors.BLUE + bcolors.BOLD + "[      Installer will copy core files to your system      ] " + bcolors.ENDC)
                    print(bcolors.BLUE + bcolors.BOLD + "[                            &                            ] " + bcolors.ENDC)
                    print(bcolors.BLUE + bcolors.BOLD + "[              All necessary foundation tools             ] " + bcolors.ENDC)

                    africs = '/usr/local/opt/africana-framework'
                    if os.path.exists(africs):
                        os.system('clear')
                        print(bcolors.BLUE + "[              Africana detected in your system           ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                 Installer runs full updates             ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[                              &                          ] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[       Upgrades for all necessary foundation tools       ] \n" + bcolors.ENDC)
                        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Africana-fri already installed in Ur system. Update It?" + bcolors.ENDC)
                        print("")
                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.RED + "[y/n]" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        choice += " "
                        if choice[0].lower() == "y":
                            try:
                                print("")
                                spinner.run_subprocess('apt-get update -y')
                                print("")
                                spinner.run_subprocess('cd /usr/local/opt/africana-framework; git pull; cd src/externals/Teardroid-phprat; git pull; cd ../AdvPhishing; git pull; cd ../anonphisher; git pull')
                                time.sleep(0.09)
                                print("")
                                print(bcolors.BLUE + "[              Setting up project discovery tools         ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                   Pleas be patient as the               ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                         Installer                       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                          Configs                        ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[                  All golang necessary tools             ] \n" + bcolors.ENDC)
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Updating all Project Discovery tools & testing them 4U" + bcolors.ENDC)
                                print("")
                                africs = '/root/go'
                                if os.path.exists(africs):
                                    spinner.run_subprocesx('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc')
                                else:
                                    pass

                                print(bcolors.BLUE + "[                 Necessary tools updateded               ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[      Anonymous tools updateded (Tor, privoxy, squid)    ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[          Wifi pentesting tools updateded (wifite)       ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[        Local pentesting tools updateded (bettercap)     ] " + bcolors.ENDC)
                                print(bcolors.BLUE + "[  Web pentest tools updateded (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                africana = bcolors.BLUE + "[ " + bcolors.GREEN + "Everything is set. Type 'africana' to start africana-fr " + bcolors.BLUE + "]" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                            except:
                                africana = bcolors.BLUE + "\n[ " + bcolors.GREEN + "An Error detected pleas run the setup again " + bcolors.BLUE + "]" + bcolors.ENDC
                                for a in africana:
                                    sys.stdout.write(a)
                                    sys.stdout.flush()
                                    time.sleep(0.09)
                                break
                        else:
                            break

                    else:
                        print("")
                        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Making Necessary dirs & Copying africana to your system" + bcolors.ENDC)
                        spinner.run_subprocess('mkdir -p /usr/local/opt; cp -r %s /usr/local/opt/africana-framework; ln -s /usr/local/opt/africana-framework/src/modules/kenyan.py /usr/local/bin/africana; chmod +x /usr/local/bin/africana' %(installer.definepath()))
                        print("")
                        print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing Essential tools: nc, zsh, wget, curl, python3" + bcolors.ENDC)
                        spinner.run_subprocesx('mkdir -p /usr/local/opt; cp -rf ../../africana-framework /usr/local/opt/; ln -s /usr/local/opt/africana-framework/src/modules/kenyan.py /usr/local/bin/africana; chmod +x /usr/local/bin/africana; echo -n "deb https://http.kali.org/kali kali-rolling main non-free contrib" >> /etc/apt/sources.list.d/kali.list; apt-get update -y; apt-get install wget zip tmux zsh git curl -y; wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc; echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -SL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg; curl -SL -o /etc/apt/sources.list.d/playit-cloud.list https://playit-cloud.github.io/ppa/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables openssh-client openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libpcap-dev npm libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser set util-linux playit hydra libssl-dev aha gcc wine32:i386; npm install -g kickthemout')
                        time.sleep(0.3)

                        if os.path.exists('~/.zshrc.bak_africana'):
                            pass
                        elif not os.path.exists('~/.zshrc.bak_africana'):
                            print("")
                            print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Configuring python3 environment at ~/.africs/bin/activate" + bcolors.ENDC)
                            spinner.run_subprocesx('cd ~/; cp .zshrc .zshrc.bak_africana; pip3 install --upgrade setuptools; pip3 install virtualenv; python3 -m virtualenv .africs --system-site-packages; echo -n "\n# Generated by africana-framework. Delete at your own risk!\nsource ~/.africs/bin/activate" >> ~/.zshrc')
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

                        git = '/usr/local/opt/africana-framework/'
                        if os.path.exists(git):
                            try:
                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing and Configuring python3 necessary requirements" + bcolors.ENDC)
                                spinner.run_subprocesx('zsh -c "source ~/.zshrc; cd /usr/local/opt/africana-framework; pip3 install -r requirements.txt"')

                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing and Configuring some github third party tools" + bcolors.ENDC)
                                try:
                                    spinner.run_subprocesx('cd /usr/local/opt/africana-framework/src/externals/androrat; apt install ./zipalign_*_amd64.deb; cd ..; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; rm -rf main.zip; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; git clone https://github.com/Ignitetch/AdvPhishing.git; cd AdvPhishing/; chmod 777 *; ./Linux-Setup.sh; cd ..;git clone https://github.com/TermuxHackz/anonphisher; cd ./anonphisher; chmod 777 *; bash anonphisher.sh')
                                except:
                                    pass

                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Installing Project Discovery tools & testing them" + bcolors.ENDC)
                                spinner.run_subprocesx('go install github.com/projectdiscovery/pdtm/cmd/pdtm@latest; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)')
                                africs = '/root/go'
                                if os.path.exists(africs):
                                    os.system('cd ~/go/bin; ./pdtm -ia; source ~/.zshrc; update-grub2')
                                else:
                                    print("")
                                    print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.RED + "'./pdtm -ia; source ~/.zshrc'. Command failed to execute" + bcolors.ENDC)
                                    time.sleep(0.3)
                                    spinner.run_subprocess('update-grub2')

                                    print("")
                                    print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                                    print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                                    print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Everything is set. Type 'africana' to launch africana" + bcolors.ENDC)
                                    break
                            except:
                                print("")
                                print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Setup failed. Pleas retrie again and check your internet connection." + bcolors.ENDC)
                                print("")
                                break
                        else:
                            print("")
                            print(bcolors.BLUE + "[                Necessary tools installed                ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[     Anonymous tools installed (Tor, privoxy, squid )    ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[          Wifi pentesting tools installed (wifite)       ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[        Local pentesting tools installed (bettercap)     ] " + bcolors.ENDC)
                            print(bcolors.BLUE + "[  Web pentest tools installed (nuclei, nikto, osmedeus)  ] \n" + bcolors.ENDC)
                            print(bcolors.BOLD + "[" + bcolors.BLUE + "+" + bcolors.ENDC + bcolors.BOLD + "] " + bcolors.GREEN + "Everything is set. Type 'africana' to launch africana" + bcolors.ENDC)
                            break

## Uninstaller
                elif choice == '4':
                    try:
                        os.system('clear')
                        beauty.graphics(), scriptures.verses()
                        print("")
                        print(bcolors.BLUE + "[ ..............Just Incase of any bug....................] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[........Pleas email me @:rojahsmontari@gmail.com ........] " + bcolors.ENDC)
                        print(bcolors.BLUE + "[..Are U sure you want to completely uninstall africana?..]\n" + bcolors.ENDC)
                        choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "Do you agree to the terms of service [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
                        choice += " "
                        if choice[0].lower() == "y":
                            spinner.run_subprocess('systemctl --no-pager status privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service -l')
                            print("")
                            spinner.run_subprocess('systemctl stop privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service; systemctl disable privoxy.service tor@default.service dnsmasq.service squid.service changemac@eth0.service')
                            if os.path.exists('~/.zshrc.bak_africana'):
                                spinner.run_subprocess('cd ~/; rm -rf .zshrc; mv .zshrc.bak_africana .zshrc')
                            else:
                                pass
                            if os.path.exists('/etc/sysctl.conf.bak_africana'):
                                spinner.run_subprocess('cd /etc/; rm -rf sysctl.conf; mv sysctl.conf.bak_africana sysctl.conf')
                            else:
                                pass
                            if os.path.exists('/etc/tor/torrc.bak_africana'):
                                spinner.run_subprocess('cd /etc/tor/; rm -rf torrc; mv torrc.bak_africana torrc')
                            else:
                                pass
                            if os.path.exists('/lib/systemd/system/tor@default.service.bak_africana'):
                                spinner.run_subprocess('cd /lib/systemd/system/; rm -rf tor@default.service; mv tor@default.service.bak_africana tor@default.service')
                            else:
                                pass
                            if os.path.exists('/etc/systemd/system/changemac@.service'):
                                spinner.run_subprocess('rm -rf /etc/systemd/system/changemac@.bak_africana')
                            else:
                                pass
                            if os.path.exists('/etc/systemd/system/changemac@.bak_africana'):
                                spinner.run_subprocess('rm -rf /etc/systemd/system/changemac@.bak_africana')
                            else:
                                pass
                            if os.path.exists('/etc/squid/squid.conf.bak_africana'):
                                spinner.run_subprocess('cd /etc/squid/; rm -rf squid.conf; mv squid.conf.bak_africana squid.conf')
                            else:
                                pass
                            if os.path.exists('/etc/privoxy/privoxy.bak_africana'):
                                spinner.run_subprocess('cd /etc/privoxy/; rm -rf config; mv privoxy.bak_africana config')
                            else:
                                pass
                            if os.path.exists('/etc/dhcp/dhclient.conf.bak_africana'):
                                spinner.run_subprocess('cd /etc/dhcp/; rm -rf dhclient.conf; mv dhclient.conf.bak_africana dhclient.conf')
                            else:
                                pass
                            if os.path.exists('~/.zshrc.bak_africana'):
                                spinner.run_subprocess('cd ~/; rm -rf .zshrc; mv .zshrc.bak_africana .zshrc')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.bak_africana'):
                                spinner.run_subprocess('cd /etc/default/; rm -rf grub; mv grub.bak_africana grub')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.d/kali-themes.cfg.bak_africana'):
                                spinner.run_subprocess('cd /etc/default/grub.d/; rm -rf kali-themes.cfg; mv kali-themes.cfg.bak_africana kali-themes.cfg')
                            else:
                                pass
                            if os.path.exists('/etc/default/grub.d/06_dark_theme.bak_africana'):
                                spinner.run_subprocess('cd /etc/default/grub.d/; rm -rf 06_dark_theme; mv 06_dark_theme.bak_africana 06_dark_theme')
                            else:
                                pass
                            africs = '/usr/local/opt/africana-framework'
                            if os.path.exists(africs):
                                spinner.run_subprocess('rm -rf /usr/local/opt/africana-framework; rm -rf /usr/local/bin/africana')
                            else:
                                pass
                            africana = bcolors.ENDC + "              ~>[ " + bcolors.YELLOW + "africana-framework uninstalled. " + bcolors.ENDC +  "]<~" + bcolors.ENDC
                            for a in africana:
                                sys.stdout.write(a)
                                sys.stdout.flush()
                                time.sleep(0.03)
                            break
                        elif choice[0].lower() == "n":
                            installer.update_system()
                            break
                        else:
                            print("")
                            warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 1 " + bcolors.BLUE +  " ]<~" + bcolors.ENDC
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
                        print("")
                        warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 4 " + bcolors.BLUE +  "]<~" + bcolors.ENDC
                        for w in warn:
                            sys.stdout.write(w)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        installer.update_system()
                        pass
                    except:
                        neo.one()
                        break
            except:
                neo.one()
                break

installer = update()
if ' __name__' == '__main__':
        sys.exit(installer())
