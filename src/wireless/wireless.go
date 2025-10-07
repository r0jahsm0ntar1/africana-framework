//John 3:16

package wireless

import(
    "os"
    "fmt"
    "menus"
    "utils"
    "strings"
    "banners"
    "bcolors"
    "strconv"
    "scriptures"
    "subprocess"
)

var (
    Function string
)

type stringMatcher struct {
    names  []string
    action func()
}

func WirelessPentest() {
    for {
        Input, err := utils.DisplayPrompt(subprocess.Version, "wireless", Function)
        if err != nil {
            if err.Error() == "interrupted" {
                fmt.Printf("%s", "\n")
                break
            }
            break
        }

        Input = strings.TrimSpace(Input)
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
            utils.SystemShell(strings.ToLower(Input))
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
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutorials},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, menus.HelpInfoWireless},
        {[]string{"m", "menu"}, menus.MenuFive},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.WirelessOptions(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListWirelessFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run wifite", "use wifite", "exec wifite", "start wifite", "launch wifite", "exploit wifite", "execute wifite"}, func() {WirelessPenFunctions("wifite")}},
        {[]string{"? 1", "info 1", "help 1", "wifite", "info wifite", "help wifite"}, func() {menus.HelpInfoWifite(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run fluxion", "use fluxion", "exec fluxion", "start fluxion", "launch fluxion", "exploit fluxion", "execute fluxion"}, func() {WirelessPenFunctions("fluxion")}},
        {[]string{"? 2", "info 2", "help 2", "fluxion", "info fluxion", "help fluxion"},func() { menus.HelpInfoFluxion(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run bettercap", "use bettercap", "exec bettercap", "start bettercap", "launch bettercap", "exploit bettercap", "execute bettercap"}, func() {WirelessPenFunctions("bettercap")}},
        {[]string{"? 3", "info 3", "help 3", "bettercap", "info bettercap", "help bettercap"}, func() {menus.HelpInfoBetterCap(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run airgeddon", "use airgeddon", "exec airgeddon", "start airgeddon", "launch airgeddon", "exploit airgeddon", "execute airgeddon"}, func() {WirelessPenFunctions("airgeddon")}},
        {[]string{"? 4", "info 4", "help 4", "airgeddon", "info airgeddon", "help airgeddon"}, func() {menus.HelpInfoAirGeddon(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wifipumpkin", "use wifipumpkin", "exec wifipumpkin", "start wifipumpkin", "launch wifipumpkin", "exploit wifipumpkin", "execute wifipumpkin"}, func() {WirelessPenFunctions("wifipumpkin")}},
        {[]string{"? 5", "info 5", "help 5", "wifipumpkin", "info wifipumpkin", "help wifipumpkin"}, func() {menus.HelpInfoWifiPumpkin(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run wifipumpkin3", "use wifipumpkin3", "exec wifipumpkin3", "start wifipumpkin3", "launch wifipumpkin3", "exploit wifipumpkin3", "execute wifipumpkin3"}, func() {WirelessPenFunctions("wifipumpkin3")}},
        {[]string{"? 6", "info 6", "help 6", "wifipumpkin3", "info wifipumpkin3", "help wifipumpkin3"}, func() {menus.HelpInfoWifiPumpkin(utils.WlanIFace, utils.LHost, utils.WirelessLogs)}},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run upsent", "use upsent", "exec upsent", "start upsent", "launch upsent", "exploit upsent", "execute upsent"}, func() {WirelessPenFunctions("upsent")}},
        {[]string{"? 7", "info 7", "help 7", "upsent", "info upsent", "help upsent"}, menus.UpsentTools},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run upsent2", "use upsent2", "exec upsent2", "start upsent2", "launch upsent2", "exploit upsent2", "execute upsent2"}, func() {WirelessPenFunctions("upsent2")}},
        {[]string{"? 8", "info 8", "help 8", "upsent2", "info upsent2", "help upsent2"}, menus.UpsentTools},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run upsent3", "use upsent3", "exec upsent3", "start upsent3", "launch upsent3", "exploit upsent3", "execute upsent3"}, func() {WirelessPenFunctions("upsent3")}},
        {[]string{"? 9", "info 9", "help 9", "upsent3", "info upsent3", "help upsent3"}, menus.UpsentTools},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, scriptures.ScriptureNarrators},
        {[]string{"? 10", "verses", "info 10", "help 10", "info verses", "help verses"}, menus.HelpInfoVerses},
    }

    cmdLower := strings.ToLower(cmd)
    for _, group := range commandGroups {
        for _, name := range group.names {
            if name == cmdLower {
                if err := subprocess.ValidateToolsDir(); err != nil {
                    fmt.Printf("%s%s[!] %sError: %s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
                    fmt.Printf("%s%s[>] %sPlease launch the installation process to avoid errors ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
                    return true
                }
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
      "ssid": &utils.Ssid,
      "mode": &utils.NeMode,
      "function": &Function,
      "lhost": &utils.LHost,
      "lport": &utils.LPort,
      "hport": &utils.HPort,
      "rhost": &utils.RHost,
      "rhosts": &utils.RHost,
      "functions": &Function,
      "target": &utils.RHost,
      "distro": &utils.Distro,
      "targets": &utils.RHost,
      "proxy": &utils.Proxies,
      "script": &utils.Script,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.GateWay,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "venvname": &utils.VenvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "output": &utils.WirelessLogs,
      "obfuscator": &utils.Obfuscator,
      "outputlog": &utils.WirelessLogs,
      "outputlogs": &utils.WirelessLogs,
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
        maxWidth += 1

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }

        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Printf("%s", "\n")
        }
        fmt.Printf("%s", "\n")
        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)
    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 1

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Printf("%s", "\n")
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
      "ssid": &utils.Ssid,
      "mode": &utils.NeMode,
      "function": &Function,
      "lhost": &utils.LHost,
      "lport": &utils.LPort,
      "hport": &utils.HPort,
      "rhost": &utils.RHost,
      "rhosts": &utils.RHost,
      "functions": &Function,
      "target": &utils.RHost,
      "distro": &utils.Distro,
      "targets": &utils.RHost,
      "proxy": &utils.Proxies,
      "script": &utils.Script,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.GateWay,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "venvname": &utils.VenvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "output": &utils.WirelessLogs,
      "obfuscator": &utils.Obfuscator,
      "outputlog": &utils.WirelessLogs,
      "outputlogs": &utils.WirelessLogs,
    }

    validKeys := make([]string, 0, len(unsetValues))
    for k := range unsetValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = utils.DefaultValues[key]
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
        maxWidth += 1

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }

        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Printf("%s", "\n")
        }
        fmt.Printf("%s", "\n")
        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 1

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Printf("%s", "\n")
    }
}

func executeFunction() {
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if utils.WlanIFace == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters IFACE. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    WirelessPenFunctions(Function, utils.WlanIFace)
}

func WirelessPenFunctions(Function string, args ...interface{}) {
    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.WiMode,
            IFACE: utils.WlanIFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            //RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.WirelessLogs,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
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
            MODE: utils.WiMode,
            IFACE: utils.WlanIFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            //RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.WirelessLogs,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, true)
    }

    commands := map[string]func(){

        "1": func() {Wifite(utils.WlanIFace)},
        "2": func() {Fluxion()},
        "3": func() {Bettercap(utils.WlanIFace)},
        "4": func() {AirGeddon()},
        "5": func() {WifiPumpkin()},
        "6": func() {WifiPumpkin3(utils.WlanIFace, utils.Ssid)},

        "wifite":       func() {Wifite(utils.WlanIFace)},
        "fluxion":      func() {Fluxion()},
        "bettercap":    func() {Bettercap(utils.WlanIFace)},
        "airgeddon":    func() {AirGeddon()},
        "wifipumpkin":  func() {WifiPumpkin()},
        "WifiPumpkin3": func() {WifiPumpkin3(utils.WlanIFace, utils.Ssid)},

    }

    textCommands := []string{"wifite", "fluxion", "bettercap", "airgeddon", "wifipumpkin"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListWirelessFunctions()
        return
    }

    lowerInput := strings.ToLower(Function)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, Function)
    menus.ListWirelessFunctions()
}

func AirGeddon() {
    Tools := []string{"airgeddon"}
    for _, Tool := range Tools {
        if err := utils.LocateTool(Tool); err != nil {
            return
        }
    }
    subprocess.Run("airgeddon")
}

func Fluxion() {
    Tools := []string{"iptables"}
    for _, Tool := range Tools {
        if err := utils.LocateTool(Tool); err != nil {
            return
        }
    }
    subprocess.Run("cd %s/fluxion/; bash ./fluxion.sh", utils.WirelessTools)
    return
}

func WifiPumpkin() {
    Tools := []string{"wifipumpkin3", "iptables"}
    for _, Tool := range Tools {
        if err := utils.LocateTool(Tool); err != nil {
            return
        }
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("wifipumpkin3")
        utils.ClearScreen()
    }
    subprocess.Run("wifipumpkin3")
    return
}

func WifiPumpkin3(WlanIFace, Ssid string) {
    Tools := []string{"wifipumpkin3", "iptables"}
    for _, Tool := range Tools {
        if err := utils.LocateTool(Tool); err != nil {
            return
        }
    }
    if utils.Ssid == "" {
        fmt.Printf("\n%s[!] %sFailed to validate MISSING SSID: %s%s Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.Green, Function, utils.WlanIFace, bcolors.Endc)
        return
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("wifipumpkin3; clear")
    }
    subprocess.Run(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxies noproxy; start"`, utils.WlanIFace, utils.Ssid)
    return
}

func Wifite(WlanIFace string) {
    subprocess.Run("cd %s/wifite/; %s ./wifite.py -i %s --ignore-locks --keep-ivs -p 1339 -mac --random-mac -v -inf --bully --pmkid --dic %s --require-fakeauth --nodeauth --pmkid-timeout 120", utils.WirelessTools, utils.VenvPython, utils.WlanIFace, utils.RokyPath)
    return
}

func Bettercap(WlanIFace string) {
    Tools := []string{"bettercap"}
    for _, Tool := range Tools {
        if err := utils.LocateTool(Tool); err != nil {
            return
        }
    }
    subprocess.Run("airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file %s/handshakes; wifi.deauth all'", utils.WlanIFace, utils.WlanIFace, utils.WlanIFace, utils.WlanIFace, utils.WlanIFace, utils.WirelessLogs)
    return
}
