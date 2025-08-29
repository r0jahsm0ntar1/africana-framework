//John 3:16

package webcrackers

import (
    "os"
    "fmt"
    "sort"
    "menus"
    "utils"
    "os/exec"
    "banners"
    "strings"
    "bcolors"
    "strconv"
    "io/ioutil"
    "scriptures"
    "subprocess"
    "path/filepath"
    "encoding/json"
)

var (
    Function, ReconDir string
)

var defaultValues = map[string]string{
   "rhost":     utils.RHost,
   "rhosts":    utils.RHost,
   "target":    utils.RHost,
   "function":  utils.RHost,
   "proxy":     utils.Proxies,
   "proxies":   utils.Proxies,
   "wordlist":  utils.WordsList,
   "output":    utils.WebCrackersLogs,
}

type stringMatcher struct {
    names  []string
    action func()
}

func WebsitesPentest() {
    for {
        fmt.Printf("%s%s%safr%s%s websites(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        utils.Scanner.Scan()
        Input := strings.TrimSpace(utils.Scanner.Text())
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
        {[]string{"c", "cls", "clear", "cls screen", "clear screen", "screen cls", "screen clear"}, utils.ClearScreen},

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
        {[]string{"option", "options", "show option", "show options"}, func() {menus.WebsitesOptions(utils.RHost, utils.WebCrackersLogs, utils.Proxies, utils.UserName, utils.PassWord, utils.WordsList)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListWebsitesFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run netmap", "use netmap", "exec netmap", "start netmap", "launch netmap", "exploit netmap", "execute netmap"}, func() {WebPenFunctions("netmap")}},
        {[]string{"? 1", "info 1", "help 1", "netmap", "info netmap", "help netmap"}, menus.HelpInfoPortScan},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run enumscan", "use enumscan", "exec enumscan", "start enumscan", "launch enumscan", "exploit enumscan", "execute enumscan"}, func() {WebPenFunctions("enumscan")}},
        {[]string{"? 2", "info 2", "help 2", "enumscan", "info enumscan", "help enumscan"}, menus.HelpInfoEnumScan},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run dnsrecon", "use dnsrecon", "exec dnsrecon", "start dnsrecon", "launch dnsrecon", "exploit dnsrecon", "execute dnsrecon"}, func() {WebPenFunctions("dnsrecon")}},
        {[]string{"? 3", "info 3", "help 3", "dnsrecon", "info dnsrecon", "help dnsrecon"}, menus.HelpInfoDnsRecon},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run techscan", "use techscan", "exec techscan", "start techscan", "launch techscan", "exploit techscan", "execute techscan"}, func() {WebPenFunctions("techscan")}},
        {[]string{"? 4", "info 4", "help 4", "techscan", "info techscan", "help techscan"}, menus.HelpInfoTechScan},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run asetscan", "use asetscan", "exec asetscan", "start asetscan", "launch asetscan", "exploit asetscan", "execute asetscan"}, func() {WebPenFunctions("asetscan")}},
        {[]string{"? 5", "info 5", "help 5", "asetscan", "info asetscan", "help asetscan"}, menus.HelpInfoAsetScan},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run fuzzscan", "use fuzzscan", "exec fuzzscan", "start fuzzscan", "launch fuzzscan", "exploit fuzzscan", "execute fuzzscan"}, func() {WebPenFunctions("fuzzscan")}},
        {[]string{"? 6", "info 6", "help 6", "fuzzscan", "info fuzzscan", "help fuzzscan"}, menus.HelpInfoFuzzScan},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run leakscan", "use leakscan", "exec leakscan", "start leakscan", "launch leakscan", "exploit leakscan", "execute leakscan"}, func() {WebPenFunctions("leakscan")}},
        {[]string{"? 7", "info 7", "help 7", "leakscan", "info leakscan", "help leakscan"}, menus.HelpInfoLeakScan},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run vulnscan", "use vulnscan", "exec vulnscan", "start vulnscan", "launch vulnscan", "exploit vulnscan", "execute vulnscan"}, func() {WebPenFunctions("vulnscan")}},
        {[]string{"? 8", "info 8", "help 8", "vulnscan", "info vulnscan", "help vulnscan"}, menus.HelpInfoVulnScan},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run bounty", "use bounty", "exec bounty", "start bounty", "launch bounty", "exploit bounty", "execute bounty"}, func() {WebPenFunctions("bounty")}},
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

       "func": &Function,
       "module": &Function,
       "function": &Function,
       "functions": &Function,
       "rhost": &utils.RHost,
       "rhosts": &utils.RHost,
       "target": &utils.RHost,
       "wordlist": &utils.WordsList,
       "proxy": &utils.Proxies,
       "proxies": &utils.Proxies,
       "output": &utils.WebCrackersLogs,
       "ddosmode": &utils.DDosMode,
    }

    validKeys := make([]string, 0, len(setValues))
    for k := range setValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s%s%s -> %s\n", bcolors.Cyan, strings.ToUpper(key), bcolors.Endc, value)
        return
    }

    var suggestions []string
    lowerInput := strings.ToLower(key)
    for _, cmd := range validKeys {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            suggestions = append(suggestions, cmd)
        }
    }

    if len(suggestions) > 0 {
        fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

        maxWidth := 0
        for _, s := range suggestions {
            if len(s) > maxWidth {
                maxWidth = len(s)
            }
        }
        maxWidth += 3

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }

        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }

        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)
    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 3

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Println()
    }
}

func handleUnsetCommand(parts []string) {
    if len(parts) < 2 {
        menus.HelpInfoSet()
        return
    }
    key := parts[1]
    unsetValues := map[string]*string{

       "func": &Function,
       "module": &Function,
       "function": &Function,
       "functions": &Function,
       "rhost": &utils.RHost,
       "rhosts": &utils.RHost,
       "target": &utils.RHost,
       "wordlist": &utils.WordsList,
       "proxy": &utils.Proxies,
       "proxies": &utils.Proxies,
       "output": &utils.WebCrackersLogs,
       "ddosmode": &utils.DDosMode,
    }

    validKeys := make([]string, 0, len(unsetValues))
    for k := range unsetValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key]
        fmt.Printf("%s -> %s\n", strings.ToUpper(key), "Null")
        return
    }

    var suggestions []string
    lowerInput := strings.ToLower(key)
    for _, cmd := range validKeys {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            suggestions = append(suggestions, cmd)
        }
    }

    if len(suggestions) > 0 {
        fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

        maxWidth := 0
        for _, s := range suggestions {
            if len(s) > maxWidth {
                maxWidth = len(s)
            }
        }
        maxWidth += 3

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }
        
        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }

        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 3

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Println()
    }
}

func executeFunction() {
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    WebPenFunctions(Function, utils.RHost)
}

func formatRHost(rawHost string) (cleanedHost, fullURL string) {
    prefixes := []string{"http://", "https://", "www."}
    cleanedHost = rawHost
    for _, prefix := range prefixes {
        cleanedHost = strings.TrimPrefix(cleanedHost, prefix)
    }
    cleanedHost = strings.TrimSpace(cleanedHost)

    if !strings.HasPrefix(rawHost, "http://") && !strings.HasPrefix(rawHost, "https://") {
        fullURL = "http://" + cleanedHost
    } else {
        fullURL = rawHost
    }

    return cleanedHost, fullURL
}

func WebPenFunctions(Function string, args ...interface{}) {
    ReconDir = filepath.Join(utils.WebCrackersLogs, fmt.Sprintf("%s-%s", utils.RHost, utils.Date))
    os.MkdirAll(ReconDir, os.ModePerm)
    //cleaned, full := formatRHost(utils.RHost)

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.WeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            //OUTPUTLOGS: utils.OutPut,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            RECONDIR: ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            //
        }
    } else {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.WeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            //OUTPUTLOGS: utils.OutPut,
            FUNCTION: Function,
            RECONDIR: ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)
    }

    commands := map[string]func(){

        "1": func() {PortScan(utils.RHost, ReconDir)},
        "2": func() {EnumScan(utils.RHost, ReconDir)},
        "3": func() {DnsRecon(utils.RHost, ReconDir)},
        "4": func() {TechScan(utils.RHost, ReconDir)},
        "5": func() {AsetScan(utils.RHost, ReconDir)},
        "6": func() {FuzzScan(utils.RHost, ReconDir)},
        "7": func() {LeakScan(utils.RHost, ReconDir)},
        "8": func() {VulnScan(utils.RHost, ReconDir)},
        "9": func() {AutoScan(utils.RHost, ReconDir)},
       "10": func() {DDosAttack(utils.RHost, utils.DDosMode)},

        "netmap":   func() {PortScan(utils.RHost, ReconDir)},
        "enumscan": func() {EnumScan(utils.RHost, ReconDir)},
        "dnsrecon": func() {DnsRecon(utils.RHost, ReconDir)},
        "techscan": func() {TechScan(utils.RHost, ReconDir)},
        "asetscan": func() {AsetScan(utils.RHost, ReconDir)},
        "fuzzscan": func() {FuzzScan(utils.RHost, ReconDir)},
        "leakscan": func() {LeakScan(utils.RHost, ReconDir)},
        "vulnscan": func() {VulnScan(utils.RHost, ReconDir)},
        "bounty":   func() {AutoScan(utils.RHost, ReconDir)},
        "ddos":     func() {DDosAttack(utils.RHost, utils.DDosMode)},
    }

    textCommands := []string{"netmap", "enumscan", "dnsrecon", "techscan", "asetscan", "fuzzscan", "leakscan", "vulnscan", "bounty", "ddos"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
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

func EnumScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming full enumeration scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("subfinder -nW -d %s -o %s/subfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    ScanCRT(utils.RHost, ReconDir)
    subprocess.Popen("assetfinder %s > %s/assetfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("findomain --target %s -u %s/findomain_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("amass enum -r 1.1.1.1 -d %s -o %s/amass_%s.txt", utils.RHost, ReconDir, utils.RHost)
}

func DnsRecon(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming DNS reconnaissance scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("dnsx -l %s -o %s/dnsx_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("asnmap -d %s -o %s/asnmap_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("cdncheck -l %s -o %s/cdncheck_%s.txt", utils.RHost, ReconDir, utils.RHost)
}

func PortScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
 
    fmt.Printf("\n%s[*] %sPerforming full port scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("naabu -nmap-cli 'nmap --script firewall-bypass -v -sV' -host %s -o %s/naabu_%s.txt", utils.RHost, ReconDir, utils.RHost)
}

func TechScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming tech detection scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("httpx -title -status-code -tech-detect -follow-redirects -u %s -o %s/httpx_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("aquatone -out %s/aquatone_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("httprobe -l %s -o %s/httprobe_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("gowitness -l %s -o %s/gowitness_%s.txt", utils.RHost, ReconDir, utils.RHost)
}

func FuzzScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming fuzz scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("gobuster dir -u %s -w %s -x 7z,zip,tar,gz,bz2,sql,bak,backup,old,db,json,xml,conf,config,asp,aspx,php,jsp -t 50 -k -o %s/gobuster_%s.txt", utils.RHost, utils.WordsList, ReconDir, utils.RHost)
}

func LeakScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming leak scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("gitleaks detect -v --report-path %s/gitleaks_%s.json", utils.OutPutDir, utils.RHost)
    subprocess.Popen("trufflehog filesystem --no-update --json --output %s/trufflehog_%s.json", utils.OutPutDir, utils.RHost)
    subprocess.Popen("semgrep -l --output %s/semgrep_%s.json", utils.OutPutDir, utils.RHost)
}

func AsetScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sPerforming asset scan ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("waybackurls %s > %s/waybackurls_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("gau %s > %s/gau_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("urlfinder -l %s -o %s/urlfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    subprocess.Popen("gospider -s %s -o %s/gospider_%s", utils.RHost, ReconDir, utils.RHost)
}

func VulnScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[!] %sPerforming vuln scan ...%s", bcolors.Yellow, bcolors.Endc, bcolors.Endc)
    httpxOutput := filepath.Join(ReconDir, fmt.Sprintf("httpx_%s.txt", utils.RHost))
    sortedHttpxOutput := filepath.Join(ReconDir, fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    katanaFinalOutPut := filepath.Join(ReconDir, fmt.Sprintf("katana_%s.txt", utils.RHost))

    fmt.Printf("\n%s[*] %sScanning host using httpx ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("httpx -u %s -title -tech-detect -status-code -o %s", utils.RHost, httpxOutput)
    SortHttpx(RHost, ReconDir)

    fmt.Printf("\n%s[*] %sResource scanning using katana ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("katana -kf all -d 45 -r 1.1.1.1 -list %s -o %s", sortedHttpxOutput, katanaFinalOutPut)
    
    fmt.Printf("\n%s[*] %sVulnerability assessment using nuclei ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("nuclei -list %s -o %s/nuclei_%s.txt", sortedHttpxOutput, ReconDir, utils.RHost)

    fmt.Printf("\n%s[*] %sPort scanning using naabu ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("naabu -nmap-cli 'nmap -o %s/nmap_%s.txt --script firewall-bypass -v -sV' -list %s -o %s/naabu_%s.txt", ReconDir, utils.RHost, sortedHttpxOutput, ReconDir, utils.RHost)

    fmt.Printf("\n%s[*] %sXss scanning using dalfox ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("dalfox file --only-poc r --ignore-return 302,404,403 --skip-bav %s -o %s/dalfox_%s.txt", katanaFinalOutPut, ReconDir, utils.RHost)
}

func AutoScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting bug bounty for %s%s%s ...", bcolors.Yellow, bcolors.Endc, bcolors.BrightGreen, utils.RHost, bcolors.Endc)
    EnumScan(RHost, ReconDir); CombineSu(RHost, ReconDir); Intrusive(RHost, ReconDir)
    fmt.Printf("%s[*] %sWorkflow Results saved at %s%s%s\n", bcolors.Green, bcolors.Endc, bcolors.BrightGreen, ReconDir, bcolors.Endc)
}


func ScanCRT(RHost, ReconDir string) error {
    fmt.Printf("\n%s[*] %sScanning for subdomains with crt.sh ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)

    crtOutput := filepath.Join(ReconDir, fmt.Sprintf("crt_%s.txt", utils.RHost))
    crtURL := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", utils.RHost)

    cmd := exec.Command("curl", "-s", crtURL)
    response, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to execute curl command: %v", err)
    }

    var data []map[string]interface{}
    if err := json.Unmarshal(response, &data); err != nil {
        return fmt.Errorf("failed to parse JSON response: %v", err)
    }

    uniqueSubdomains := make(map[string]struct{})
    for _, entry := range data {
        if nameValue, ok := entry["name_value"].(string); ok && nameValue != "" {
            nameValue = strings.Replace(nameValue, "*.", "", -1)
            for _, domain := range strings.Split(nameValue, "\n") {
                domain = strings.TrimSpace(domain)
                if domain != "" {
                    uniqueSubdomains[domain] = struct{}{}
                }
            }
        }
    }

    writeUniqueSubdomainsToFile(crtOutput, uniqueSubdomains)

    content, err := os.ReadFile(crtOutput)
    if err != nil {
        return fmt.Errorf("failed to read output file: %v", err)
    }
    fmt.Print(string(content))

    return nil
}

func writeUniqueSubdomainsToFile(filename string, uniqueSubdomains map[string]struct{}) {
    subdomains := make([]string, 0, len(uniqueSubdomains))
    for subdomain := range uniqueSubdomains {
        subdomains = append(subdomains, subdomain)
    }
    sort.Strings(subdomains)

    err := ioutil.WriteFile(filename, []byte(strings.Join(subdomains, "\n")+"\n"), 0644)
    utils.CheckErr(err)
}

func CombineSu(RHost, ReconDir string) {
    subfinderOutput := filepath.Join(ReconDir, fmt.Sprintf("subfinder_%s.txt", utils.RHost))
    amassOutput := filepath.Join(ReconDir, fmt.Sprintf("amass_%s.txt", utils.RHost))
    assetfinderOutput := filepath.Join(ReconDir, fmt.Sprintf("assetfinder_%s.txt", utils.RHost))
    findomainOutput := filepath.Join(ReconDir, fmt.Sprintf("findomain_%s.txt", utils.RHost))
    crtOutput := filepath.Join(ReconDir, fmt.Sprintf("crt_%s.txt", utils.RHost))
    subdomainsOutput := filepath.Join(ReconDir, fmt.Sprintf("subdomains_%s.txt", utils.RHost))

    subdomainFiles := []string{subfinderOutput, amassOutput, assetfinderOutput, findomainOutput, crtOutput}
    uniqueSubdomains := make(map[string]struct{})

    for _, file := range subdomainFiles {
        if utils.FileExists(file) {
            lines, err := ioutil.ReadFile(file)
            utils.CheckErr(err)
            for _, line := range strings.Split(string(lines), "\n") {
                line = strings.TrimSpace(line)
                if line != "" {
                    uniqueSubdomains[line] = struct{}{}
                }
            }
        } else {
            fmt.Printf("File not found: %s\n", file)
        }
    }

    writeUniqueSubdomainsToFile(subdomainsOutput, uniqueSubdomains)
}

func SortHttpx(RHost, ReconDir string) {
    sortedHttpxOutput := filepath.Join(ReconDir, fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    httpxOutput := filepath.Join(ReconDir, fmt.Sprintf("httpx_%s.txt", utils.RHost))

    linesBytes, err := ioutil.ReadFile(httpxOutput)
    utils.CheckErr(err)

    lines := strings.Split(string(linesBytes), "\n")

    strippedLines := make([]string, 0, len(lines))
    for _, line := range lines {
        strippedLines = append(strippedLines, utils.StripANSI(line))
    }

    strippedBracketsLines := make([]string, 0, len(strippedLines))
    for _, line := range strippedLines {
        strippedBracketsLines = append(strippedBracketsLines, utils.StripBrackets(line))
    }

    sortedLines := make([]string, 0, len(strippedBracketsLines))
    for _, line := range strippedBracketsLines {
        if len(line) > 0 {
            sortedLines = append(sortedLines, line)
        }
    }

    sort.Slice(sortedLines, func(i, j int) bool {
        iStatusCode, _ := strconv.Atoi(strings.Fields(sortedLines[i])[1])
        jStatusCode, _ := strconv.Atoi(strings.Fields(sortedLines[j])[1])
        return iStatusCode < jStatusCode
    })

    withUrls := make([]string, 0, len(sortedLines))
    for _, line := range sortedLines {
        url := strings.Fields(line)[0]
        url = strings.Replace(url, "https://", "", 1)
        url = strings.Replace(url, "http://", "", 1)
        withUrls = append(withUrls, url)
    }

    err = ioutil.WriteFile(sortedHttpxOutput, []byte(strings.Join(withUrls, "\n")+"\n"), 0644)
    utils.CheckErr(err)
}

func Intrusive(RHost, ReconDir string) {
    sortedHttpxOutput := filepath.Join(ReconDir, fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    katanaFinalOutPut := filepath.Join(ReconDir, fmt.Sprintf("katana_%s.txt", utils.RHost))

    fmt.Sprintf("\n%s[*] %sScanning host using httpx ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("httpx -list %s/subdomains_%s.txt -title -tech-detect -status-code -o %s/httpx_%s.txt", ReconDir, utils.RHost, ReconDir, utils.RHost)
    SortHttpx(RHost, ReconDir)

    fmt.Printf("\n%s[*] %sResource scanning using katana ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("katana -kf all -d 45 -r 1.1.1.1 -list %s -o %s", sortedHttpxOutput, katanaFinalOutPut)

    fmt.Printf("\n%s[*] %sVulnerability assessment using nuclei ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("nuclei -list %s -o %s/nuclei_%s.txt", sortedHttpxOutput, ReconDir, utils.RHost)

    fmt.Printf("\n%s[*] %sPort scanning using naabu ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("naabu -nmap-cli 'nmap -o %s/nmap_%s.txt -sV' -list %s -o %s/naabu_%s.txt", ReconDir, utils.RHost, sortedHttpxOutput, ReconDir, utils.RHost)

    fmt.Printf("\n%s[*] %sXss scanning using dalfox ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    subprocess.Popen("dalfox file --only-poc r --ignore-return 302,404,403 --skip-bav %s -o %s/dalfox_%s.txt", katanaFinalOutPut, ReconDir, utils.RHost)
}

func DDosAttack(RHost, DDosMode string) {
    fmt.Sprintf("\n%s[*] %sSending DDos packets ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    switch strings.ToLower(utils.DDosMode) {
    case "1", "light":
        fmt.Printf("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -r '100' --threads '100' --nuke '10000' -a '%s'", utils.WeBountyTools, utils.RHost)
        return
    case "2", "simple":
        fmt.Printf("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -r '100' --threads '100' --spray '1000' --smurf '1000' --tachyon '1000' --monlist '1000' --fraggle '1000' --sniper '1000' -a '%s'", utils.WeBountyTools, utils.RHost)
        return
    case "3", "hard":
        fmt.Printf("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -r '100' --threads '100' --loic '1000' --loris '1000' -a '%s'", utils.WeBountyTools, utils.RHost)
        return
    case "4", "harder":
        fmt.Printf("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -r '100' --threads '100' --loic '1000' --loris '1000' --ufosyn '1000' --xmas '1000' --nuke '1000' --ufoack '1000' --uforst '1000' --droper '1000' --overlap '1000' --pinger '1000' --ufoudp '1000' -a '%s'", utils.WeBountyTools, utils.RHost)
        return
    case "5", "update":
        fmt.Printf("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet --download-zombies", utils.WeBountyTools)
        return
    case "6", "test":
        fmt.Printf("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -t botnet/zombies.txt", utils.WeBountyTools)
        return
    case "7", "grider":
        fmt.Printf("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet --grider", utils.WeBountyTools)
        return
    case "8", "gui":
        fmt.Printf("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet --gui", utils.WeBountyTools)
        return
    case "9", "intense":
        fmt.Printf("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.Green, bcolors.Endc)
        subprocess.Popen("cd %s/ufonet/; python3 ufonet -r '100' --threads '100' --loic '1000' --loris '1000' --ufosyn '1000' --spray '1000' --smurf '1000' --xmas '1000' --nuke '1000' --tachyon '1000' --monlist '1000' --fraggle '1000' --sniper '1000' --ufoack '1000' --uforst '1000' --droper '1000' --overlap '1000' --pinger '1000' --ufoudp '1000' -a '%s'", utils.WeBountyTools, utils.RHost)
        return
    default:
        fmt.Printf("\n%s[!] %sError: Invalid DDOSMODE. %s'%s' %sTry %s'show ddosmode' %stype", bcolors.BrightRed, bcolors.Endc, bcolors.Yellow, bcolors.Endc, bcolors.Green, utils.DDosMode, bcolors.Endc)
    }
}



//Intense bug bounty hunt
func FindIPs(RHost, ResolversFile, ReconDir string) {
    fmt.Printf("\n%s[*] %sFinding IPs for subdomains ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)

    subdomainsOutput := filepath.Join(ReconDir, fmt.Sprintf("subdomains_%s.txt", utils.RHost))
    ipsOutput := filepath.Join(ReconDir, fmt.Sprintf("%s.ips.txt", utils.RHost))

    subprocess.Popen("massdns -r %s -t A -o S -w %s %s", ResolversFile, ipsOutput, subdomainsOutput)
}

func Sublist3r(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/sublist3r/; python3 sublist3r.py -v -d %s", utils.WeBountyTools, utils.RHost)
}

func Ashock1(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/ashok/; python3 ashok.py --headers %s", utils.WeBountyTools, utils.RHost)
}

func OneForAll(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/oneforall/; python3 oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s ", utils.WeBountyTools, utils.RHost)
}

func SeekOlver(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s", utils.WeBountyTools, utils.RHost)
}

func ParamSpider(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming urls mine ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/paramspider/; python3 paramspider.py -s -d %s", utils.WeBountyTools, utils.RHost)
}

func SslScan(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming ssl scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("sslscan --show-certificate --show-sigs --tlsall --verbose %s", utils.WeBountyTools, utils.RHost)
}

func Gobuster(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming dir recon ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("gobuster dir vhost --debug --no-error --random-agent -w %s -e -u %s", utils.WordsList, utils.WeBountyTools, utils.RHost)
}

func Nuclei(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("nuclei -u %s", utils.WeBountyTools, utils.RHost)
}

func Nikto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("nikto -ask no -Cgidirs all -Display 3 -host %s", utils.WeBountyTools, utils.RHost)
}

func Bbot(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s", utils.RHost)
}

func Uniscan(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("uniscan -qweds -u %s", utils.WeBountyTools, utils.RHost)
}

func SqlmapAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s", utils.RHost)
}

func SqlmapMan() {
    fmt.Printf("\n%s[*] %sPerforming man sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard")
}

func CommixAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s", utils.RHost)
}

func CommixMan() {
    fmt.Printf("\n%s[*] %sPerforming man command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard")
}

func KatanaAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav", utils.RHost)
}

func XsserAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("xsser -c 100 --Cl -u %s", utils.RHost)
}

func XsserMan() {
    fmt.Printf("\n%s[*] %sPerforming man xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("xsser -u --wizard")
}

func NetTacker1(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s", utils.WeBountyTools, utils.RHost)
}

func NetTacker2(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming domain scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m subdomain_scan", utils.WeBountyTools, utils.RHost)
}

func NetTacker3(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming admin finder scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m scan", utils.WeBountyTools, utils.RHost)
}

func NetTacker4(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming info gathering ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m information_gathering -s", utils.WeBountyTools, utils.RHost)
    fmt.Println()
}

func NetTacker5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m server_version_vuln", utils.WeBountyTools, utils.RHost)
}

func NetTacker6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming CVE scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m cve", utils.WeBountyTools, utils.RHost)
}

func NetTacker7(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s -m high_severity", utils.WeBountyTools, utils.RHost)
}

func NetTacker8(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming intrusive scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py -i %s --profile all/results.txt", utils.WeBountyTools, utils.RHost)
}

func NetTacker9() {
    fmt.Printf("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/nettacker/; python3 nettacker.py --start-api --api-access-key africana", utils.WeBountyTools, utils.RHost)
}

func Jok3r1() {
    fmt.Printf("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; bash install-all.sh", utils.WeBountyTools, utils.RHost)
}

func Jok3r2() {
    fmt.Printf("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py toolbox --update-all --auto", utils.WeBountyTools, utils.RHost)
}

func Jok3r3() {
    fmt.Printf("\n%s[*] %sShowing tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py toolbox --show-all", utils.WeBountyTools, utils.RHost)
}

func Jok3r4() {
    fmt.Printf("\n%s[*] %sShowing supported products ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py info --services", utils.WeBountyTools, utils.RHost)
}

func Jok3r5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming avery fast-scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast", utils.WeBountyTools, utils.RHost)
}

func Jok3r6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming security checks on " + bcolors.BrightRed + "\nTarget: " + bcolors.Yellow + "%s \n" + bcolors.Endc, utils.RHost)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast", utils.WeBountyTools, utils.RHost)
}

func Jok3r7(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast", utils.WeBountyTools, utils.RHost)
}

func Jok3r8() {
    fmt.Printf("\n%s[*] %sShowing results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -")
}

func Jok3r9() {
    fmt.Printf("\n%s[*] %sCleaning results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s/jok3r/; python3 jok3r.py db 'mission -d default'")
}

func Osmedeus1() {
    fmt.Printf("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean")
}

func Osmedeus2(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming simple scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus --nv scan -t %s", utils.RHost)
}

func Osmedeus3(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus --nv scan -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus4(Targe string) {
    fmt.Printf("\n%s[*] %sPerforming bulk scan on " + bcolors.BrightYellow + "Targets " + bcolors.Endc + "= " + bcolors.Green + bcolors.Italic + "%s \n" + bcolors.Endc, utils.RHost)
    subprocess.Popen("osmedeus scan -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming cloud scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s", utils.RHost)
}

func Osmedeus6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus7(RHost string) {
    fmt.Printf("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus scan --update-vuln -t %s", utils.RHost)
}

func Osmedeus8(Port string) {
    fmt.Printf("\n%s[*] %sStarted Osmedeusweb UI http://127.0.0.1:%s%s", bcolors.Green, bcolors.Endc, Port, bcolors.Endc)
    subprocess.Popen("osmedeus server --port %s", Port)
}

func Osmedeus9() {
    fmt.Printf("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen("osmedeus report list")
}


