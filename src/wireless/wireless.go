package wireless

import(
    "os"
    "fmt"
    "menus"
    "utils"
    "bufio"
    "strings"
    "banners"
    "bcolors"
    "scriptures"
    "subprocess"
)


var (
    Mode = "auto"
    Lport = "9999"
    Hport = "3333"
    Iface = "wlan0"
    Ssid = "End times ministries"

    Lhost = LhostIp
    LhostIp, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Input, Rhost, Proxy, Function, Script string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    
    "proxies": "",
    "mode": "auto",
    "function": "",
    "lport": "9999",
    "hport": "3333",
    "iface": "wlan0",
    "ssid": "End times ministries",

    "lhost": LhostIp,
    "output": OutPutDir,
}

func WirelessPentest() {
    for {
        fmt.Printf("%s%safr3%s wireless(%ssrc/pentest_%s.fn%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, Function, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
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
    "info":             menus.HelpInfoWireless,

    "m":                menus.MenuFive,
    "menu":             menus.MenuFive,

    "option":           menus.WirelessOptions,
    "options":          menus.WirelessOptions,
    "show option":      menus.WirelessOptions,
    "show options":     menus.WirelessOptions,

    "show all":         menus.ListWirelessFunctions,
    "list all":         menus.ListWirelessFunctions,

    "func":             menus.ListWirelessFunctions,
    "funcs":            menus.ListWirelessFunctions,
    "functions":        menus.ListWirelessFunctions,
    "show func":        menus.ListWirelessFunctions,
    "list funcs":       menus.ListWirelessFunctions,
    "show funcs":       menus.ListWirelessFunctions,
    "show function":    menus.ListWirelessFunctions,
    "list function":    menus.ListWirelessFunctions,
    "list functions":   menus.ListWirelessFunctions,
    "show functions":   menus.ListWirelessFunctions,

    "module":           menus.ListWirelessFunctions,
    "modules":          menus.ListWirelessFunctions,
    "list module":      menus.ListWirelessFunctions,
    "show module":      menus.ListWirelessFunctions,
    "list modules":     menus.ListWirelessFunctions,
    "show modules":     menus.ListWirelessFunctions,

    "1":                func() {WirelessPenFunctions("wifite")},
    "run 1":            func() {WirelessPenFunctions("wifite")},
    "use 1":            func() {WirelessPenFunctions("wifite")},
    "exec 1":           func() {WirelessPenFunctions("wifite")},
    "start 1":          func() {WirelessPenFunctions("wifite")},
    "launch 1":         func() {WirelessPenFunctions("wifite")},
    "exploit 1":        func() {WirelessPenFunctions("wifite")},
    "execute 1":        func() {WirelessPenFunctions("wifite")},
    "run wifite":       func() {WirelessPenFunctions("wifite")},
    "use wifite":       func() {WirelessPenFunctions("wifite")},
    "exec wifite":      func() {WirelessPenFunctions("wifite")},
    "start wifite":     func() {WirelessPenFunctions("wifite")},
    "launch wifite":    func() {WirelessPenFunctions("wifite")},
    "exploit wifite":   func() {WirelessPenFunctions("wifite")},
    "execute wifite":   func() {WirelessPenFunctions("wifite")},

    "? 1":              menus.HelpInfoWifite,
    "info 1":           menus.HelpInfoWifite,
    "help 1":           menus.HelpInfoWifite,
    "setups":           menus.HelpInfoWifite,
    "info setups":      menus.HelpInfoWifite,
    "help setups":      menus.HelpInfoWifite,

    "2":                func() {WirelessPenFunctions("fluxion")},
    "run 2":            func() {WirelessPenFunctions("fluxion")},
    "use 2":            func() {WirelessPenFunctions("fluxion")},
    "exec 2":           func() {WirelessPenFunctions("fluxion")},
    "start 2":          func() {WirelessPenFunctions("fluxion")},
    "launch 2":         func() {WirelessPenFunctions("fluxion")},
    "exploit 2":        func() {WirelessPenFunctions("fluxion")},
    "execute 2":        func() {WirelessPenFunctions("fluxion")},
    "run fluxion":      func() {WirelessPenFunctions("fluxion")},
    "use fluxion":      func() {WirelessPenFunctions("fluxion")},
    "exec fluxion":     func() {WirelessPenFunctions("fluxion")},
    "start fluxion":    func() {WirelessPenFunctions("fluxion")},
    "launch fluxion":   func() {WirelessPenFunctions("fluxion")},
    "exploit fluxion":  func() {WirelessPenFunctions("fluxion")},
    "execute fluxion":  func() {WirelessPenFunctions("fluxion")},

    "? 2":              menus.HelpInfoFluxion,
    "info 2":           menus.HelpInfoFluxion,
    "help 2":           menus.HelpInfoFluxion,
    "torsocks":         menus.HelpInfoFluxion,
    "info torsocks":    menus.HelpInfoFluxion,
    "help torsocks":    menus.HelpInfoFluxion,

    "3":                    func() {WirelessPenFunctions("bettercap")},
    "run 3":                func() {WirelessPenFunctions("bettercap")},
    "use 3":                func() {WirelessPenFunctions("bettercap")},
    "exec 3":               func() {WirelessPenFunctions("bettercap")},
    "start 3":              func() {WirelessPenFunctions("bettercap")},
    "launch 3":             func() {WirelessPenFunctions("bettercap")},
    "exploit 3":            func() {WirelessPenFunctions("bettercap")},
    "execute 3":            func() {WirelessPenFunctions("bettercap")},
    "run bettercap":        func() {WirelessPenFunctions("bettercap")},
    "use bettercap":        func() {WirelessPenFunctions("bettercap")},
    "exec bettercap":       func() {WirelessPenFunctions("bettercap")},
    "start bettercap":      func() {WirelessPenFunctions("bettercap")},
    "launch bettercap":     func() {WirelessPenFunctions("bettercap")},
    "exploit bettercap":    func() {WirelessPenFunctions("bettercap")},
    "execute bettercap":    func() {WirelessPenFunctions("bettercap")},

    "? 3":              menus.HelpInfoBetterCap,
    "info 3":           menus.HelpInfoBetterCap,
    "help 3":           menus.HelpInfoBetterCap,
    "networks":         menus.HelpInfoBetterCap,
    "info networks":    menus.HelpInfoBetterCap,
    "help networks":    menus.HelpInfoBetterCap,

    "4":                    func() {WirelessPenFunctions("airgeddon")},
    "run 4":                func() {WirelessPenFunctions("airgeddon")},
    "use 4":                func() {WirelessPenFunctions("airgeddon")},
    "exec 4":               func() {WirelessPenFunctions("airgeddon")},
    "start 4":              func() {WirelessPenFunctions("airgeddon")},
    "launch 4":             func() {WirelessPenFunctions("airgeddon")},
    "exploit 4":            func() {WirelessPenFunctions("airgeddon")},
    "execute 4":            func() {WirelessPenFunctions("airgeddon")},
    "run airgeddon":        func() {WirelessPenFunctions("airgeddon")},
    "use airgeddon":        func() {WirelessPenFunctions("airgeddon")},
    "exec airgeddon":       func() {WirelessPenFunctions("airgeddon")},
    "start airgeddon":      func() {WirelessPenFunctions("airgeddon")},
    "launch airgeddon":     func() {WirelessPenFunctions("airgeddon")},
    "exploit airgeddon":    func() {WirelessPenFunctions("airgeddon")},
    "execute airgeddon":    func() {WirelessPenFunctions("airgeddon")},

    "? 4":              menus.HelpInfoAirGeddon,
    "info 4":           menus.HelpInfoAirGeddon,
    "help 4":           menus.HelpInfoAirGeddon,
    "exploits":         menus.HelpInfoAirGeddon,
    "info exploits":    menus.HelpInfoAirGeddon,
    "help exploits":    menus.HelpInfoAirGeddon,

    "5":                    func() {WirelessPenFunctions("wifipumpkin")},
    "run 5":                func() {WirelessPenFunctions("wifipumpkin")},
    "use 5":                func() {WirelessPenFunctions("wifipumpkin")},
    "exec 5":               func() {WirelessPenFunctions("wifipumpkin")},
    "start 5":              func() {WirelessPenFunctions("wifipumpkin")},
    "launch 5":             func() {WirelessPenFunctions("wifipumpkin")},
    "exploit 5":            func() {WirelessPenFunctions("wifipumpkin")},
    "execute 5":            func() {WirelessPenFunctions("wifipumpkin")},
    "run wifipumpkin":      func() {WirelessPenFunctions("wifipumpkin")},
    "use wifipumpkin":      func() {WirelessPenFunctions("wifipumpkin")},
    "exec wifipumpkin":     func() {WirelessPenFunctions("wifipumpkin")},
    "start wifipumpkin":    func() {WirelessPenFunctions("wifipumpkin")},
    "launch wifipumpkin":   func() {WirelessPenFunctions("wifipumpkin")},
    "exploit wifipumpkin":  func() {WirelessPenFunctions("wifipumpkin")},
    "execute wifipumpkin":  func() {WirelessPenFunctions("wifipumpkin")},

    "? 5":              menus.HelpInfoWifiPumpkin,
    "info 5":           menus.HelpInfoWifiPumpkin,
    "help 5":           menus.HelpInfoWifiPumpkin,
    "wireless":         menus.HelpInfoWifiPumpkin,
    "info wireless":    menus.HelpInfoWifiPumpkin,
    "help wireless":    menus.HelpInfoWifiPumpkin,

    "6":                    func() {WirelessPenFunctions("WifiPumpkin3")},
    "run 6":                func() {WirelessPenFunctions("WifiPumpkin3")},
    "use 6":                func() {WirelessPenFunctions("WifiPumpkin3")},
    "exec 6":               func() {WirelessPenFunctions("WifiPumpkin3")},
    "start 6":              func() {WirelessPenFunctions("WifiPumpkin3")},
    "launch 6":             func() {WirelessPenFunctions("WifiPumpkin3")},
    "exploit 6":            func() {WirelessPenFunctions("WifiPumpkin3")},
    "execute 6":            func() {WirelessPenFunctions("WifiPumpkin3")},
    "run WifiPumpkin3":     func() {WirelessPenFunctions("WifiPumpkin3")},
    "use WifiPumpkin3":     func() {WirelessPenFunctions("WifiPumpkin3")},
    "exec WifiPumpkin3":    func() {WirelessPenFunctions("WifiPumpkin3")},
    "start WifiPumpkin3":   func() {WirelessPenFunctions("WifiPumpkin3")},
    "launch WifiPumpkin3":  func() {WirelessPenFunctions("WifiPumpkin3")},
    "exploit WifiPumpkin3": func() {WirelessPenFunctions("WifiPumpkin3")},
    "execute WifiPumpkin3": func() {WirelessPenFunctions("WifiPumpkin3")},

    "? 6":              menus.HelpInfoWifiPumpkin,
    "info 6":           menus.HelpInfoWifiPumpkin,
    "help 6":           menus.HelpInfoWifiPumpkin,
    "crackers":         menus.HelpInfoWifiPumpkin,
    "info crackers":    menus.HelpInfoWifiPumpkin,
    "help crackers":    menus.HelpInfoWifiPumpkin,

    "7":                func() {menus.UpsentTools()},
    "run 7":            func() {menus.UpsentTools()},
    "use 7":            func() {menus.UpsentTools()},
    "exec 7":           func() {menus.UpsentTools()},
    "start 7":          func() {menus.UpsentTools()},
    "launch 7":         func() {menus.UpsentTools()},
    "exploit 7":        func() {menus.UpsentTools()},
    "execute 7":        func() {menus.UpsentTools()},
    "run phishers":     func() {menus.UpsentTools()},
    "use phishers":     func() {menus.UpsentTools()},
    "exec phishers":    func() {menus.UpsentTools()},
    "start phishers":   func() {menus.UpsentTools()},
    "launch phishers":  func() {menus.UpsentTools()},
    "exploit phishers": func() {menus.UpsentTools()},
    "execute phishers": func() {menus.UpsentTools()},

    "? 7":              menus.UpsentTools,
    "info 7":           menus.UpsentTools,
    "help 7":           menus.UpsentTools,
    "phishers":         menus.UpsentTools,
    "info phishers":    menus.UpsentTools,
    "help phishers":    menus.UpsentTools,

    "8":                func() {menus.UpsentTools()},
    "run 8":            func() {menus.UpsentTools()},
    "use 8":            func() {menus.UpsentTools()},
    "exec 8":           func() {menus.UpsentTools()},
    "start 8":          func() {menus.UpsentTools()},
    "launch 8":         func() {menus.UpsentTools()},
    "exploit 8":        func() {menus.UpsentTools()},
    "execute 8":        func() {menus.UpsentTools()},
    "run websites":     func() {menus.UpsentTools()},
    "use websites":     func() {menus.UpsentTools()},
    "exec websites":    func() {menus.UpsentTools()},
    "start websites":   func() {menus.UpsentTools()},
    "launch websites":  func() {menus.UpsentTools()},
    "exploit websites": func() {menus.UpsentTools()},
    "execute websites": func() {menus.UpsentTools()},

    "? 8":              menus.UpsentTools,
    "info 8":           menus.UpsentTools,
    "help 8":           menus.UpsentTools,
    "websites":         menus.UpsentTools,
    "info websites":    menus.UpsentTools,
    "help websites":    menus.UpsentTools,

    "9":               menus.UpsentTools,
    "run 9":           menus.UpsentTools,
    "use 9":           menus.UpsentTools,
    "exec 9":          menus.UpsentTools,
    "start 9":         menus.UpsentTools,
    "launch 9":        menus.UpsentTools,
    "exploit 9":       menus.UpsentTools,
    "execute 9":       menus.UpsentTools,
    "run credits":     menus.UpsentTools,
    "use credits":     menus.UpsentTools,
    "exec credits":    menus.UpsentTools,
    "start credits":   menus.UpsentTools,
    "launch credits":  menus.UpsentTools,
    "exploit credits": menus.UpsentTools,
    "execute credits": menus.UpsentTools,

    "? 9":              menus.UpsentTools,
    "info 9":           menus.UpsentTools,
    "help 9":           menus.UpsentTools,
    "credits":          menus.UpsentTools,
    "info credits":     menus.UpsentTools,
    "help credits":     menus.UpsentTools,

    "99":               scriptures.ScriptureNarators,
    "run 99":           scriptures.ScriptureNarators,
    "use 99":           scriptures.ScriptureNarators,
    "exec 99":          scriptures.ScriptureNarators,
    "start 99":         scriptures.ScriptureNarators,
    "launch 99":        scriptures.ScriptureNarators,
    "exploit 99":       scriptures.ScriptureNarators,
    "execute 99":       scriptures.ScriptureNarators,
    "run verses":       scriptures.ScriptureNarators,
    "use verses":       scriptures.ScriptureNarators,
    "exec verses":      scriptures.ScriptureNarators,
    "start verses":     scriptures.ScriptureNarators,
    "launch verses":    scriptures.ScriptureNarators,
    "exploit verses":   scriptures.ScriptureNarators,
    "execute verses":   scriptures.ScriptureNarators,

    "? 99":             menus.HelpInfoVerses,
    "verses":           menus.HelpInfoVerses,
    "info 99":          menus.HelpInfoVerses,
    "help 99":          menus.HelpInfoVerses,
    "info verses":      menus.HelpInfoVerses,
    "help verses":      menus.HelpInfoVerses,

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

        "mode": &Mode,
        "ssid": &Ssid,
        "iface": &Iface,
        "lhost": &Lhost,
        "lport": &Lport,
        "hport": &Hport,
        "proxies": &Proxy,
        "func": &Function,
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

        "mode": &Mode,
        "ssid": &Ssid,
        "iface": &Iface,
        "lhost": &Lhost,
        "lport": &Lport,
        "hport": &Hport,
        "proxies": &Proxy,
        "func": &Function,
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
    if Iface == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters IFACE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    WirelessPenFunctions(Function, Iface)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
}

func WirelessPenFunctions(Function string, args ...interface{}) {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    commands := map[string]func(){
           "wifite": func() {Wifite(Iface)},
          "fluxion": func() {Fluxion()},
        "bettercap": func() {Bettercap(Iface)},
        "airgeddon": func() {AirGeddon()},
      "wifipumpkin": func() {WifiPumpkin()},
     "WifiPumpkin3": func() {WifiPumpkin3(Iface, Ssid)},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sFailed to validate Function: %s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, Function, bcolors.ENDC)
    }
}


func AirGeddon() {
    subprocess.Popen(`airgeddon`)
}

func Fluxion() {
    subprocess.Popen(`cd %s/wireless/fluxion/; bash fluxion.sh`, ToolsDir)
}

func WifiPumpkin() {
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3`)
        utils.ClearScreen()
    }
    subprocess.Popen(`wifipumpkin3`)
}

func WifiPumpkin3(Iface string, Ssid string) {
     if Ssid == "" {
        fmt.Printf("\n%s[!] %sFailed to validate MISSING SSID: %s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, Function, Iface, bcolors.ENDC)
        return
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3; clear`)
    }
    subprocess.Popen(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxy noproxy; start"`, Iface, Ssid)
}

func Wifite(Iface string) {
    subprocess.Popen(`wifite -i %s --ignore-locks --keep-ivs -p 1339 -mac --random-mac -v -inf --bully --pmkid --dic %s --require-fakeauth --nodeauth --pmkid-timeout 120`, Iface, RokyPath)
}

func Bettercap(Iface string) {
    subprocess.Popen(`airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"`, Iface, Iface, Iface, Iface, Iface)
}
