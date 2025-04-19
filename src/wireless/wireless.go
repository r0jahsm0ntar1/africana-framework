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

type stringMatcher struct {
    names  []string
    action func()
}

func WirelessPentest() {
    for {
        fmt.Printf("%s%safr3%s wireless(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        {[]string{"info"}, menus.HelpInfoWireless},
        {[]string{"m", "menu"}, menus.MenuFive},
        {[]string{"option", "options", "show option", "show options"}, menus.WirelessOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListWirelessFunctions},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run wifite", "use wifite", "exec wifite", "start wifite", "launch wifite", "exploit wifite", "execute wifite"}, func() { WirelessPenFunctions("wifite") }},
        {[]string{"? 1", "info 1", "help 1", "wifite", "info wifite", "help wifite"}, menus.HelpInfoWifite},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run fluxion", "use fluxion", "exec fluxion", "start fluxion", "launch fluxion", "exploit fluxion", "execute fluxion"}, func() { WirelessPenFunctions("kali", "install") }},
        {[]string{"? 2", "info 2", "help 2", "fluxion", "info fluxion", "help fluxion"}, menus.HelpInfoFluxion},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run bettercap", "use bettercap", "exec bettercap", "start bettercap", "launch bettercap", "exploit bettercap", "execute bettercap"}, func() { WirelessPenFunctions("bettercap", "install") }},
        {[]string{"? 3", "info 3", "help 3", "bettercap", "info bettercap", "help bettercap"}, menus.HelpInfoBetterCap},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run airgeddon", "use airgeddon", "exec airgeddon", "start airgeddon", "launch airgeddon", "exploit airgeddon", "execute airgeddon"}, func() { WirelessPenFunctions("airgeddon", "install") }},
        {[]string{"? 4", "info 4", "help 4", "airgeddon", "info airgeddon", "help airgeddon"}, menus.HelpInfoAirGeddon},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wifipumpkin", "use wifipumpkin", "exec wifipumpkin", "start wifipumpkin", "launch wifipumpkin", "exploit wifipumpkin", "execute wifipumpkin"}, func() { WirelessPenFunctions("wifipumpkin", "install") }},
        {[]string{"? 5", "info 5", "help 5", "wifipumpkin", "info wifipumpkin", "help wifipumpkin"}, menus.HelpInfoWifiPumpkin},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run wifipumpkin3", "use wifipumpkin3", "exec wifipumpkin3", "start wifipumpkin3", "launch wifipumpkin3", "exploit wifipumpkin3", "execute wifipumpkin3"}, func() { WirelessPenFunctions("wifipumpkin3", "install") }},
        {[]string{"? 6", "info 6", "help 6", "wifipumpkin3", "info wifipumpkin3", "help wifipumpkin3"}, menus.HelpInfoWifiPumpkin},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run upsent", "use upsent", "exec upsent", "start upsent", "launch upsent", "exploit upsent", "execute upsent"}, func() { WirelessPenFunctions("upsent", "install") }},
        {[]string{"? 7", "info 7", "help 7", "upsent", "info upsent", "help upsent"}, menus.UpsentTools},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run upsent2", "use upsent2", "exec upsent2", "start upsent2", "launch upsent2", "exploit upsent2", "execute upsent2"}, func() { WirelessPenFunctions("upsent2", "install") }},
        {[]string{"? 8", "info 8", "help 8", "upsent2", "info upsent2", "help upsent2"}, menus.UpsentTools},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run upsent3", "use upsent3", "exec upsent3", "start upsent3", "launch upsent3", "exploit upsent3", "execute upsent3"}, func() { WirelessPenFunctions("upsent3", "install") }},
        {[]string{"? 9", "info 9", "help 9", "upsent3", "info upsent3", "help upsent3"}, menus.UpsentTools},

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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if Iface == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters IFACE. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        fmt.Printf("\n%s[!] %sFailed to validate Function: %s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.Green, Function, bcolors.Endc)
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
        fmt.Printf("\n%s[!] %sFailed to validate MISSING SSID: %s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.Green, Function, Iface, bcolors.Endc)
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
