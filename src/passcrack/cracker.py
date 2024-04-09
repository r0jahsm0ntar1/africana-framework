import sys
import time
import subprocess
from src.core.banner import *
from src.core.bcolors import *

class pass_killer(object):
    def __init__(self):
        pass

    rc = '/usr/share/wordlists/rockyou.txt.gz'
    if os.path.exists(rc):
        os.system('gunzip /usr/share/wordlists/rockyou.txt.gz')
    else:
        pass

#Offline crackers
    def aircrack_ng(self):
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your .pcap" + bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.RED + "Pcap:" + bcolors.YELLOW + "%s" %(pcap) + bcolors.GREEN + ")# " + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your wordlist" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        process = subprocess.Popen("aircrack-ng %s -w %s" %(pcap, wordlist), shell = True).wait()
        return process

    def john_cracker(self):
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your .pcap" + bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.RED + "Pcap:" + bcolors.YELLOW + "%s" %(pcap) + bcolors.GREEN + ")# " + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your wordlist" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        process = subprocess.Popen("john %s --wordlist=%s" %(pcap, wordlist), shell = True).wait()
        return process

#Online crackers
    def hydra_ssh(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SSH password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_ssh_outfile.txt -u ssh://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_ftp(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing FTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_ftp_outfile.txt -u ftp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_smb(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SMB password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_smb_outfile.txt -u smb://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_rdp(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing RDP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_rdp_outfile.txt -u rdp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_ldap(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing LDAP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_ldap_outfile.txt -u ldap://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_smtp(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SMTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_smtp_outfile.txt -u smtp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_snmtp(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SNMTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_snmtp_outfile.txt -u snmtp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_telnet(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing TELNET password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_telnet_outfile.txt -u telnet://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_https(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing HTTP/S password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o /root/.africana/hydra_https_outfile.txt -u https://%s '%(host), encoding='utf8', shell = True).wait()
        return process

#Auto cracking
    def cyberbrute_all(host):
        host = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing all services" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('cd src/externals/cyberbrute; bash cyberbrute.sh %s' %(host), shell = True).wait()
        print("")
        return process

#Hashes cracking
    def hash_buster(self):
        hashes = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Path to Hashes" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing Hashes" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(hashes) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('cd src/externals/hash-buster; python3 cracker.py -t 10 %s' %(hashes), shell = True).wait()
        print("")
        return process

pass_cracker = pass_killer()
if ' __name__' == '__main__':
        sys.exit(pass_cracker())
