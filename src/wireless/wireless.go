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
    Lport = "9999"
    Hport = "3333"
    Iface = "wlan0"
    Ssid = "End times ministries"

    Lhost = LhostIp
    LhostIp, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Input, Rhost, Proxy, Module, Script string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "ssid":     "End times ministries",
    "iface":    "wlan0",
    "proxy":    "",
    "lhost":    LhostIp,
    "lport":    "9999",
    "hport":    "3333",
    "module":   "",
    "output":   OutPutDir,
}

func WirelessPentest() {
    for {
        fmt.Printf("%s%safr3%s wireless(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "wireless_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
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
        "info":             menus.HelpInfoWireless,

        "m":                menus.MenuFive,
        "menu":             menus.MenuFive,

        "option":           menus.WirelessOptions,
        "options":          menus.WirelessOptions,
        "show option":      menus.WirelessOptions,
        "show options":     menus.WirelessOptions,

        "modules":          menus.ListWirelessModules,
        "show all":         menus.ListWirelessModules,
        "list all":         menus.ListWirelessModules,
        "list modules":     menus.ListWirelessModules,
        "show modules":     menus.ListWirelessModules,

        "1":                func() {WirelessPenModules("wifite")},
        "run 1":            func() {WirelessPenModules("wifite")},
        "use 1":            func() {WirelessPenModules("wifite")},
        "exec 1":           func() {WirelessPenModules("wifite")},
        "start 1":          func() {WirelessPenModules("wifite")},
        "launch 1":         func() {WirelessPenModules("wifite")},
        "exploit 1":        func() {WirelessPenModules("wifite")},
        "execute 1":        func() {WirelessPenModules("wifite")},
        "run wifite":       func() {WirelessPenModules("wifite")},
        "use wifite":       func() {WirelessPenModules("wifite")},
        "exec wifite":      func() {WirelessPenModules("wifite")},
        "start wifite":     func() {WirelessPenModules("wifite")},
        "launch wifite":    func() {WirelessPenModules("wifite")},
        "exploit wifite":   func() {WirelessPenModules("wifite")},
        "execute wifite":   func() {WirelessPenModules("wifite")},

        "? 1":              menus.HelpInfoWifite,
        "info 1":           menus.HelpInfoWifite,
        "help 1":           menus.HelpInfoWifite,
        "setups":           menus.HelpInfoWifite,
        "info setups":      menus.HelpInfoWifite,
        "help setups":      menus.HelpInfoWifite,

        "2":                func() {WirelessPenModules("fluxion")},
        "run 2":            func() {WirelessPenModules("fluxion")},
        "use 2":            func() {WirelessPenModules("fluxion")},
        "exec 2":           func() {WirelessPenModules("fluxion")},
        "start 2":          func() {WirelessPenModules("fluxion")},
        "launch 2":         func() {WirelessPenModules("fluxion")},
        "exploit 2":        func() {WirelessPenModules("fluxion")},
        "execute 2":        func() {WirelessPenModules("fluxion")},
        "run fluxion":      func() {WirelessPenModules("fluxion")},
        "use fluxion":      func() {WirelessPenModules("fluxion")},
        "exec fluxion":     func() {WirelessPenModules("fluxion")},
        "start fluxion":    func() {WirelessPenModules("fluxion")},
        "launch fluxion":   func() {WirelessPenModules("fluxion")},
        "exploit fluxion":  func() {WirelessPenModules("fluxion")},
        "execute fluxion":  func() {WirelessPenModules("fluxion")},

        "? 2":              menus.HelpInfoFluxion,
        "info 2":           menus.HelpInfoFluxion,
        "help 2":           menus.HelpInfoFluxion,
        "anonsurf":         menus.HelpInfoFluxion,
        "info anonsurf":    menus.HelpInfoFluxion,
        "help anonsurf":    menus.HelpInfoFluxion,

        "3":                    func() {WirelessPenModules("bettercap")},
        "run 3":                func() {WirelessPenModules("bettercap")},
        "use 3":                func() {WirelessPenModules("bettercap")},
        "exec 3":               func() {WirelessPenModules("bettercap")},
        "start 3":              func() {WirelessPenModules("bettercap")},
        "launch 3":             func() {WirelessPenModules("bettercap")},
        "exploit 3":            func() {WirelessPenModules("bettercap")},
        "execute 3":            func() {WirelessPenModules("bettercap")},
        "run bettercap":        func() {WirelessPenModules("bettercap")},
        "use bettercap":        func() {WirelessPenModules("bettercap")},
        "exec bettercap":       func() {WirelessPenModules("bettercap")},
        "start bettercap":      func() {WirelessPenModules("bettercap")},
        "launch bettercap":     func() {WirelessPenModules("bettercap")},
        "exploit bettercap":    func() {WirelessPenModules("bettercap")},
        "execute bettercap":    func() {WirelessPenModules("bettercap")},

        "? 3":              menus.HelpInfoBetterCap,
        "info 3":           menus.HelpInfoBetterCap,
        "help 3":           menus.HelpInfoBetterCap,
        "networks":         menus.HelpInfoBetterCap,
        "info networks":    menus.HelpInfoBetterCap,
        "help networks":    menus.HelpInfoBetterCap,

        "4":                    func() {WirelessPenModules("airgeddon")},
        "run 4":                func() {WirelessPenModules("airgeddon")},
        "use 4":                func() {WirelessPenModules("airgeddon")},
        "exec 4":               func() {WirelessPenModules("airgeddon")},
        "start 4":              func() {WirelessPenModules("airgeddon")},
        "launch 4":             func() {WirelessPenModules("airgeddon")},
        "exploit 4":            func() {WirelessPenModules("airgeddon")},
        "execute 4":            func() {WirelessPenModules("airgeddon")},
        "run airgeddon":        func() {WirelessPenModules("airgeddon")},
        "use airgeddon":        func() {WirelessPenModules("airgeddon")},
        "exec airgeddon":       func() {WirelessPenModules("airgeddon")},
        "start airgeddon":      func() {WirelessPenModules("airgeddon")},
        "launch airgeddon":     func() {WirelessPenModules("airgeddon")},
        "exploit airgeddon":    func() {WirelessPenModules("airgeddon")},
        "execute airgeddon":    func() {WirelessPenModules("airgeddon")},

        "? 4":              menus.HelpInfoAirGeddon,
        "info 4":           menus.HelpInfoAirGeddon,
        "help 4":           menus.HelpInfoAirGeddon,
        "exploits":         menus.HelpInfoAirGeddon,
        "info exploits":    menus.HelpInfoAirGeddon,
        "help exploits":    menus.HelpInfoAirGeddon,

        "5":                    func() {WirelessPenModules("wifipumpkin")},
        "run 5":                func() {WirelessPenModules("wifipumpkin")},
        "use 5":                func() {WirelessPenModules("wifipumpkin")},
        "exec 5":               func() {WirelessPenModules("wifipumpkin")},
        "start 5":              func() {WirelessPenModules("wifipumpkin")},
        "launch 5":             func() {WirelessPenModules("wifipumpkin")},
        "exploit 5":            func() {WirelessPenModules("wifipumpkin")},
        "execute 5":            func() {WirelessPenModules("wifipumpkin")},
        "run wifipumpkin":      func() {WirelessPenModules("wifipumpkin")},
        "use wifipumpkin":      func() {WirelessPenModules("wifipumpkin")},
        "exec wifipumpkin":     func() {WirelessPenModules("wifipumpkin")},
        "start wifipumpkin":    func() {WirelessPenModules("wifipumpkin")},
        "launch wifipumpkin":   func() {WirelessPenModules("wifipumpkin")},
        "exploit wifipumpkin":  func() {WirelessPenModules("wifipumpkin")},
        "execute wifipumpkin":  func() {WirelessPenModules("wifipumpkin")},

        "? 5":              menus.HelpInfoWifiPumpkin,
        "info 5":           menus.HelpInfoWifiPumpkin,
        "help 5":           menus.HelpInfoWifiPumpkin,
        "wireless":         menus.HelpInfoWifiPumpkin,
        "info wireless":    menus.HelpInfoWifiPumpkin,
        "help wireless":    menus.HelpInfoWifiPumpkin,

        "6":                    func() {WirelessPenModules("WifiPumpkin3")},
        "run 6":                func() {WirelessPenModules("WifiPumpkin3")},
        "use 6":                func() {WirelessPenModules("WifiPumpkin3")},
        "exec 6":               func() {WirelessPenModules("WifiPumpkin3")},
        "start 6":              func() {WirelessPenModules("WifiPumpkin3")},
        "launch 6":             func() {WirelessPenModules("WifiPumpkin3")},
        "exploit 6":            func() {WirelessPenModules("WifiPumpkin3")},
        "execute 6":            func() {WirelessPenModules("WifiPumpkin3")},
        "run WifiPumpkin3":     func() {WirelessPenModules("WifiPumpkin3")},
        "use WifiPumpkin3":     func() {WirelessPenModules("WifiPumpkin3")},
        "exec WifiPumpkin3":    func() {WirelessPenModules("WifiPumpkin3")},
        "start WifiPumpkin3":   func() {WirelessPenModules("WifiPumpkin3")},
        "launch WifiPumpkin3":  func() {WirelessPenModules("WifiPumpkin3")},
        "exploit WifiPumpkin3": func() {WirelessPenModules("WifiPumpkin3")},
        "execute WifiPumpkin3": func() {WirelessPenModules("WifiPumpkin3")},

        "? 6":              menus.HelpInfoWifiPumpkin3,
        "info 6":           menus.HelpInfoWifiPumpkin3,
        "help 6":           menus.HelpInfoWifiPumpkin3,
        "crackers":         menus.HelpInfoWifiPumpkin3,
        "info crackers":    menus.HelpInfoWifiPumpkin3,
        "help crackers":    menus.HelpInfoWifiPumpkin3,

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
        "ssid":     &Ssid,
        "iface":    &Iface,
        "proxy":    &Proxy,
        "lhost":    &Lhost,
        "lport":    &Lport,
        "hport":    &Hport,
        "module":   &Module,
        "output":   &OutPutDir,
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
        "ssid":     &Ssid,
        "iface":    &Iface,
        "proxy":    &Proxy,
        "lhost":    &Lhost,
        "lport":    &Lport,
        "hport":    &Hport,
        "module":   &Module,
        "output":   &OutPutDir,
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
    if Iface == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters IFACE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    WirelessPenModules(Module, Iface)
}

func WirelessPenModules(Module string, args ...interface{}) {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        utils.SetProxy(Proxy)
    }

    commands := map[string]func(){
           "wifite": func() {Wifite(Iface)},
          "fluxion": func() {Fluxion()},
        "bettercap": func() {Bettercap(Iface)},
        "airgeddon": func() {AirGeddon()},
      "wifipumpkin": func() {WifiPumpkin()},
     "WifiPumpkin3": func() {WifiPumpkin3(Iface, Ssid)}, 
    }

    if action, exists := commands[Module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sFailed to validate MODULE: %s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, Module, bcolors.ENDC)
    }
}


func AirGeddon() {
    subprocess.Popen(`airgeddon`)
}

func WifiPumpkin() {
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3`)
        utils.ClearScreen()
    }
    subprocess.Popen(`wifipumpkin3`)
}

func Fluxion() {
    subprocess.Popen(`cd %s/wireless/fluxion/; bash fluxion.sh`, ToolsDir)
}

func WifiPumpkin3(Iface string, Ssid string) {
     if Ssid == "" {
        fmt.Printf("\n%s[!] %sFailed to validate MISSING SSID: %s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, Module, Iface, bcolors.ENDC)
        return
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3; clear`)
    }
    subprocess.Popen(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxy noproxy; start"`, Iface, Ssid)
}

func Wifite(Iface string) {
    subprocess.Popen(`wifite -i %s --ignore-locks --keep-ivs -p 1339 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/WordList/rockyou.txt --require-fakeauth --nodeauth --pmkid-timeout 120`, Iface)
}

func Bettercap(Iface string) {
    subprocess.Popen(`airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"`, Iface, Iface, Iface, Iface, Iface)
}
