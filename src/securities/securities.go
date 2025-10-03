// John 3:16
package securities

import (
    "io"
    "os"
    "fmt"
    "net"
    "time"
    "menus"
    "utils"
    "regexp"
    "setups"
    "strconv"
    "strings"
    "banners"
    "bcolors"
    "os/exec"
    "net/http"
    "scriptures"
    "subprocess"
    "encoding/json"
)

var Function string

var GeoAPIs = []string{
    "http://ip-api.com/json/%s",
    "https://ipinfo.io/%s/json",
    "https://freegeoip.app/json/%s",
}

var IpServices = []string{
    "https://ident.me",
    "https://ipinfo.io/ip",
    "https://api.ipify.org",
    "https://icanhazip.com",
    "https://ifconfig.me/ip",
    "https://checkip.amazonaws.com",
}

type IPGeoInfo struct {
    IP        string  `json:"ip"`
    ISP       string  `json:"isp"`
    Country   string  `json:"country"`
    Region    string  `json:"region"`
    City      string  `json:"city"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    ASN       string  `json:"asn"`
    Continent string  `json:"continent"`
}

type DNSLeakTestResult struct {
    Timestamp  string     `json:"timestamp"`
    Status     string     `json:"status"`
    TestType   string     `json:"test_type"`
    DNSServers []DNSServer `json:"dns_servers"`
    LeakDetected bool     `json:"leak_detected"`
    TorCheck   TorResult  `json:"tor_check"`
}

type DNSServer struct {
    IP       string `json:"ip"`
    ISP      string `json:"isp,omitempty"`
    Country  string `json:"country,omitempty"`
    ASN      string `json:"asn,omitempty"`
    ViaTor   bool   `json:"via_tor"`
}

type TorResult struct {
    IsUsingTor bool   `json:"is_using_tor"`
    ExitIP     string `json:"exit_ip,omitempty"`
    ExitCountry string `json:"exit_country,omitempty"`
}


type stringMatcher struct {
    names  []string
    action func()
}

func Torsocks() {
    for {
        Input, err := utils.DisplayPrompt(subprocess.Version, "torsocks", Function)
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
        {[]string{"g", "t", "guide", "tutorial"}, utils.BrowseTutorials},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, func() {menus.HelpInfoTorsocks(Function)}},
        {[]string{"m", "menu"}, menus.MenuTwo},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.TorsocksOptions(Function)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, func() {menus.ListTorsocksFunctions(Function)}},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() { AnonimityFunctions("setups") }},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, func() {menus.HelpInfoTorsocksSetups(Function)}},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run vanish", "use vanish", "exec vanish", "start vanish", "launch vanish", "exploit vanish", "execute vanish"}, func() { AnonimityFunctions("vanish") }},
        {[]string{"? 2", "info 2", "help 2", "vanish", "info vanish", "help vanish"}, func() {menus.HelpInfoTorsocksVanish(Function)}},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run status", "use status", "exec status", "start status", "launch status", "exploit status", "execute status"}, func() { AnonimityFunctions("status") }},
        {[]string{"? 3", "info 3", "help 3", "status", "info status", "help status"}, func() {menus.HelpInfoTorsocksStatus(Function)}},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run torip", "use torip", "exec torip", "start torip", "launch torip", "exploit torip", "execute torip"}, func() { AnonimityFunctions("torip") }},
        {[]string{"? 4", "info 4", "help 4", "torip", "info torip", "help torip"}, func() {menus.HelpInfoTorsocksTorIp(Function)}},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run chains", "use chains", "exec chains", "start chains", "launch chains", "exploit chains", "execute chains"}, func() { AnonimityFunctions("chains") }},
        {[]string{"? 5", "info 5", "help 5", "chains", "info chains", "help chains"}, func() {menus.HelpInfoTorsocksChains(Function)}},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run reload", "use reload", "exec reload", "start reload", "launch reload", "exploit reload", "execute reload"}, func() { AnonimityFunctions("reload") }},
        {[]string{"? 6", "info 6", "help 6", "reload", "info reload", "help reload"}, func() {menus.HelpInfoTorsocksReload(Function)}},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run exitnode", "use exitnode", "exec exitnode", "start exitnode", "launch exitnode", "exploit exitnode", "execute exitnode"}, func() { AnonimityFunctions("exitnode") }},
        {[]string{"? 7", "info 7", "help 7", "exitnode", "info exitnode", "help exitnode"}, func() {menus.HelpInfoTorsocksExitnode(Function)}},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run restore", "use restore", "exec restore", "start restore", "launch restore", "exploit restore", "execute restore"}, func() { AnonimityFunctions("restore") }},
        {[]string{"? 8", "info 8", "help 8", "restore", "info restore", "help restore"}, func() {menus.HelpInfoTorsocksRestore(Function)}},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run stop", "use stop", "exec stop", "start stop", "launch stop", "exploit stop", "execute stop"}, func() { AnonimityFunctions("stop") }},
        {[]string{"? 9", "info 9", "help 9", "stop", "info stop", "help stop"}, func() {menus.HelpInfoTorsocksStop(Function)}},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, scriptures.ScriptureNarrators},
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
      "output": &utils.SecLogs,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.GateWay,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "outputlog": &utils.SecLogs,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "outputlogs": &utils.SecLogs,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "venvname": &utils.VenvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "obfuscator": &utils.Obfuscator,
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
      "output": &utils.SecLogs,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.GateWay,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "outputlog": &utils.SecLogs,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "outputlogs": &utils.SecLogs,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "venvname": &utils.VenvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "obfuscator": &utils.Obfuscator,
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
            fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Printf("%s", "\n")
    }
}

func executeFunction() {
    if Function == "" {
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    AnonimityFunctions(Function)
}

func AnonimityFunctions(functionName string, args ...interface{}) {
    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            IFACE:       utils.IFace,
            DISTRO:      utils.Distro,
            OUTPUTLOGS:  utils.SecLogs,
            PROXIES:     utils.Proxies,
            FUNCTION:    functionName,
        }, true, false)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            //
        }
    } else {
        menus.PrintSelected(menus.PrintOptions{
            IFACE:      utils.IFace,
            DISTRO:     utils.Distro,
            OUTPUTLOGS: utils.SecLogs,
            FUNCTION:   functionName,
        }, true, false)
    }

    commands := map[string]func(){
        "setups":   func() { banners.GraphicsTorNet(); setups.CheckMissing() },
        "vanish":   func() { banners.GraphicsTorNet(); ConfigureResolv(); ConfigChangeMac(); ConfigDhclient(); ConfigDnsmasq(); ConfigSquid(); ConfigPrivoxy(); ConfigTorrc(); ConfigFirewall(); StartTorServices(); CheckServiceStatus(); CheckIP(); CheckLeakTest(); TorStatus(0) },
        "status":   func() { banners.GraphicsTorNet(); CheckServiceStatus(); TorStatus(0) },
        "torip":    func() { banners.GraphicsTorNet(); CheckIP(); CheckLeakTest(); TorStatus(0) },
        "exitnode": func() { banners.GraphicsTorNet(); TorCircuit() },
        "restore":  func() { banners.GraphicsTorNet(); ResetToDefault(false, false) },
        "chains":   func() { banners.GraphicsTorNet(); subprocess.Run("tail -vf /var/log/privoxy/logfile") },
        "reload":   func() { banners.GraphicsTorNet(); ConfigTorrc(); ConfigFirewall(); CheckIP() },
        "stop":     func() { banners.GraphicsTorNet(); KillTor(); CheckServiceStatus(); TorStatus(0)},

        "1": func() { banners.GraphicsTorNet(); setups.CheckMissing() },
        "2": func() { banners.GraphicsTorNet(); ConfigureResolv(); ConfigChangeMac(); ConfigDhclient(); ConfigDnsmasq(); ConfigSquid(); ConfigPrivoxy(); ConfigTorrc(); ConfigFirewall(); StartTorServices(); CheckServiceStatus(); CheckIP(); CheckLeakTest(); TorStatus(0) },
        "3": func() { banners.GraphicsTorNet(); CheckServiceStatus(); TorStatus(0) },
        "4": func() { banners.GraphicsTorNet(); CheckIP(); CheckLeakTest(); TorStatus(0) },
        "5": func() { banners.GraphicsTorNet(); TorCircuit() },
        "6": func() { banners.GraphicsTorNet(); ResetToDefault(false, false) },
        "7": func() { banners.GraphicsTorNet(); subprocess.Run("tail -vf /var/log/privoxy/logfile") },
        "8": func() { banners.GraphicsTorNet(); ConfigTorrc(); ConfigFirewall(); CheckIP() },
        "9": func() { banners.GraphicsTorNet(); KillTor(); CheckServiceStatus(); TorStatus(0)},
    }

    textCommands := []string{"setups", "vanish", "status", "exitnode", "torip", "restore", "chains", "reload", "stop"}

    if action, exists := commands[functionName]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(functionName); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListTorsocksFunctions(Function)
        return
    }

    lowerInput := strings.ToLower(functionName)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, functionName, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, functionName)
    menus.ListTorsocksFunctions(Function)
}

func ConfigChangeMac() {
    if err := utils.LocateTool("macchanger"); err != nil {
        fmt.Printf("%sError: macchanger not found: %v%s\n", bcolors.Red, err, bcolors.Endc)
        return
    }
    content := `
# Generated by africana-framework. Delete at your own risk!
[Unit]
Description=changes mac for %I
Wants=network.target
Before=network.target
BindsTo=sys-subsystem-net-devices-%i.device
After=sys-subsystem-net-devices-%i.device
[Service]
Type=oneshot
ExecStart=/usr/bin/macchanger -r %I
RemainAfterExit=yes
[Install]
WantedBy=multi-user.target
`
    filePath := "/etc/systemd/system/changemac@.service"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        err = os.WriteFile(filePath, []byte(content), 0644)
        if err != nil {
            fmt.Printf("%s[!] %sError writing to file: %s ...", bcolors.Red, bcolors.Endc, err)
            return
        }
        fmt.Printf("%s[*] %sCreated macchanger systemd service: %s ...", bcolors.Green, bcolors.Endc, filePath)
    }// else {
        //fmt.Printf("%s[!] %sService file already exists: %s ...", bcolors.Yellow, bcolors.Endc, filePath)
    //}
}

func ConfigDhclient() {
    if err := utils.LocateTool("dhclient"); err != nil {
        fmt.Printf("%sError: dhclient not found: %v%s\n", bcolors.Red, err, bcolors.Endc)
        return
    }

    filePath := "/etc/dhcp/dhclient.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("cp -r /etc/dhcp/dhclient.conf /etc/dhcp/dhclient.conf.bak_africana")
        filesToReplacements := map[string]map[string]string{
            "/etc/dhcp/dhclient.conf": {
                "#prepend domain-name-servers 127.0.0.1;": "prepend domain-name-servers 127.0.0.1, 1.1.1.1, 1.0.0.1, 8.8.8.8, 8.8.4.4;",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func ConfigDnsmasq() {
    if err := utils.LocateTool("dnsmasq"); err != nil {
        fmt.Printf("%sError: dnsmasq not found: %v%s\n", bcolors.Red, err, bcolors.Endc)
        return
    }

    filePath := "/etc/dnsmasq.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("cp -r /etc/dnsmasq.conf /etc/dnsmasq.conf.bak_africana")
        filesToReplacements := map[string]map[string]string{
            "/etc/dnsmasq.conf": {
                "#port=5353": "port=5353",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func ConfigPrivoxy() {
    if err := utils.LocateTool("privoxy"); err != nil {
        fmt.Printf("%sError: privoxy not found: %v%s\n", bcolors.Red, err, bcolors.Endc)
        return
    }

    filePath := "/etc/privoxy/config.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("cp -r /etc/privoxy/config /etc/privoxy/config.bak_africana")
        filesToReplacements := map[string]map[string]string{
            "/etc/privoxy/config": {
                "#debug     1": "debug   1",
                "#debug     2": "debug   2",
                "#debug  1024": "debug   1024",
                "#debug  4096": "debug   4096",
                "#        forward-socks5t   /               127.0.0.1:9050 .": "forward-socks5t   /               127.0.0.1:9050 .",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func ConfigSquid() {
    if err := utils.LocateTool("squid"); err != nil {
        fmt.Printf("%sError: squid not found: %v%s\n", bcolors.Red, err, bcolors.Endc)
        return
    }
    filePath := "/etc/squid/squid.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Run("cp -r /etc/squid/squid.conf /etc/squid/squid.conf.bak_africana")
        filesToReplacements := map[string]map[string]string{
            "/etc/squid/squid.conf": {
                "http_port 3128": "http_port 3129\nnever_direct allow all\nshutdown_lifetime 0 seconds\ncache_peer localhost parent 8118 7 no-digest no-query",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func ConfigTorrc() error {
    if err := utils.LocateTool("tor"); err != nil {
        return err
    }

    if _, err := os.Stat("/etc/resolv.conf"); os.IsNotExist(err) {
        if err := CreateResolvConf(); err != nil {
            return err
        }
    }

    return configureTorrc()
}

func CreateResolvConf() error {
    fmt.Printf("\n%s%s[!] %sMissing resolv.conf. Create it? [Y/n]: ", bcolors.Bold, bcolors.Yellow, bcolors.Endc)

    utils.Scanner.Scan()
    input := strings.TrimSpace(strings.ToLower(utils.Scanner.Text()))

    if input == "n" || input == "no" {
        return fmt.Errorf("resolv.conf required but user declined creation")
    }

    content := `# Generated by africana-framework. Delete at your own risk!
nameserver 127.0.0.1    # Local DNS (dnsmasq/Tor)
nameserver 1.1.1.1      # Cloudflare DNS
nameserver 1.0.0.1      # Cloudflare DNS
nameserver 8.8.8.8      # Google DNS
nameserver 8.8.4.4      # Google DNS
options rotate          # Rotate between nameservers
options timeout:1       # 1 second timeout
options attempts:2      # 2 attempts per nameserver
`
    if _, err := os.Stat("/etc/resolv.conf"); err == nil {
        backupCmd := exec.Command("cp", "/etc/resolv.conf", "/etc/resolv.conf.backup_africana")
        if err := backupCmd.Run(); err != nil {
            fmt.Printf("%s%s[!] %sFailed to create backup: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err, bcolors.Endc)
        }
    }

    if err := os.WriteFile("/etc/resolv.conf", []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to create resolv.conf: %v", err)
    }

    if err := os.Chmod("/etc/resolv.conf", 0644); err != nil {
        fmt.Printf("%s%s[!] %sWarning: Failed to set permissions: %v%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, err, bcolors.Endc)
    }

    fmt.Printf("%s%s[+] %sresolv.conf created successfully%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    return nil
}

func configureTorrc() error {
    content := `# Generated by africana-framework. Delete at your own risk!
VirtualAddrNetwork 10.192.0.0/10
AutomapHostsOnResolve 1
CookieAuthentication 1
TransPort 9040
SocksPort 9050
DNSPort 5353
`
    filePath := "/etc/tor/torrc"
    fmt.Printf("%s%s[*] %sConfiguring torrc.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if _, err := os.Stat(filePath); err == nil {
        if err := exec.Command("cp", filePath, filePath+".bak_torsocks").Run(); err != nil {
            return fmt.Errorf("failed to create backup: %v", err)
        }
    }

    if err := os.MkdirAll("/etc/tor", 0755); err != nil {
        return fmt.Errorf("failed to create tor directory: %v", err)
    }

    if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to write torrc: %v", err)
    }

    fmt.Printf("%s%s[+] %sTorrc configured successfully", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    return nil
}

func ConfigureResolv() error {
    resolvContent := `# Generated by africana-framework. Delete at your own risk!
nameserver 127.0.0.1    # Local DNS (dnsmasq/Tor)
nameserver 1.1.1.1      # Cloudflare DNS
nameserver 1.0.0.1      # Cloudflare DNS
nameserver 8.8.8.8      # Google DNS
nameserver 8.8.4.4      # Google DNS
options rotate          # Rotate between nameservers
options timeout:1       # 1 second timeout
options attempts:2      # 2 attempts per nameserver
`
    content, err := os.ReadFile("/etc/resolv.conf")
    if err != nil {
        return fmt.Errorf("failed to read resolv.conf: %v", err)
    }

    if strings.Contains(string(content), "Generated by africana-framework") {
        fmt.Printf("%s%s[!] %sResolv.conf already configured.\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        return nil
    }

    fmt.Printf("%s%s[*] %sConfiguring resolv.conf.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if err := exec.Command("cp", "/etc/resolv.conf", "/etc/resolv.conf.backup_torsocks").Run(); err != nil {
        return fmt.Errorf("failed to create backup: %v", err)
    }

    if err := os.WriteFile("/etc/resolv.conf", []byte(resolvContent), 0644); err != nil {
        return fmt.Errorf("failed to write resolv.conf: %v", err)
    }

    fmt.Printf("%s%s[+] %sDone configuring Resolv.conf ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    return nil
}

func IsServiceActive(service string) (bool, error) {
    cmd := exec.Command("systemctl", "is-active", service)
    output, err := cmd.Output()
    if err != nil {
        if _, ok := err.(*exec.ExitError); ok {
            return false, nil
        }
        return false, err
    }
    return strings.TrimSpace(string(output)) == "active", nil
}

func CheckIptables() string {
    cmd := exec.Command("iptables", "-t", "nat", "-L", "-n")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("\n%s[!] %sError while checking the iptables details ...\n", bcolors.BrightRed, bcolors.Endc)
    }
    return string(output)
}

func CheckIP() {
    fmt.Printf("%s%s[*] %sGetting your public IP address ...", bcolors.Bold, bcolors.Green, bcolors.Endc)
    fmt.Printf("\n%s%s[+] %sGetting your system geolocation ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    ip := getPublicIP()
    if ip == "" {
        fmt.Printf("%s%s[!] %sError: Could not retrieve public IP address.\n", bcolors.Bold, bcolors.Red, bcolors.Endc)
        return
    }

    getIPGeolocation(ip)
}

func getPublicIP() string {
    for _, url := range IpServices {
        resp, err := http.Get(url)
        if err != nil {
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            continue
        }

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            continue
        }
        ip := strings.TrimSpace(string(body))

        if isValidIP(ip) {
            return ip
        }
    }
    return ""
}

func getIPGeolocation(ip string) {
    info, err := getIPInfo(ip)
    if err != nil {
        fmt.Printf("%sError: Could not retrieve geolocation data%s\n", bcolors.Red, bcolors.Endc)
        return
    }

    now := time.Now()
    startTime := time.Now()
    responseTime := time.Since(startTime).Round(time.Millisecond)

    fmt.Printf(`%s
{
  "ip_address": "%s%s%s",
  "geolocation": {
    "continent": "%s%s%s",
    "country": "%s%s%s",
    "region": "%s%s%s",
    "city": "%s%s%s",
    "coordinates": {
      "latitude": %s%.6f%s,
      "longitude": %s%.6f%s
    },
    "isp": "%s%s%s"
  },
  "service_provider": "%s%s%s",
  "timestamp": {
    "unix": %s%d%s,
    "iso8601": "%s%s%s",
    "human_readable": "%s%s%s",
    "timezone": "%s%s%s"
  },
  "query_info": {
    "success": %strue%s,
    "response_time": "%s%s%s"
  }
}
%s`, bcolors.Endc, bcolors.Cyan, info.IP, bcolors.Endc, bcolors.Cyan, info.Continent, bcolors.Endc, bcolors.Cyan, info.Country, bcolors.Endc, bcolors.Cyan, info.Region, bcolors.Endc, bcolors.Cyan, info.City, bcolors.Endc, bcolors.Cyan, info.Latitude, bcolors.Endc, bcolors.Cyan, info.Longitude, bcolors.Endc, bcolors.Cyan, info.ISP, bcolors.Endc, bcolors.Cyan, info.ISP, bcolors.Endc, bcolors.Cyan, now.Unix(), bcolors.Endc, bcolors.Cyan, now.Format(time.RFC3339), bcolors.Endc, bcolors.Cyan, now.Format("Monday, January 2, 2006 at 3:04:05 PM"), bcolors.Endc, bcolors.Cyan, now.Format("MST"), bcolors.Endc, bcolors.Cyan, bcolors.Endc, bcolors.Cyan, responseTime.String(), bcolors.Endc, bcolors.Endc)
}

func getString(data map[string]interface{}, keys ...string) string {
    for _, key := range keys {
        if value, exists := data[key]; exists {
            if str, ok := value.(string); ok {
                return str
            }
        }
    }
    return "Unknown"
}

func getFloat(data map[string]interface{}, keys ...string) float64 {
    for _, key := range keys {
        if value, exists := data[key]; exists {
            switch v := value.(type) {
            case float64:
                return v
            case float32:
                return float64(v)
            case int:
                return float64(v)
            case string:
                if f, err := strconv.ParseFloat(v, 64); err == nil {
                    return f
                }
            }
        }
    }
    return 0.0
}

func isValidIP(ip string) bool {
    return net.ParseIP(ip) != nil && !strings.Contains(ip, "<")
}

func SystemIP() (string, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    for _, service := range IpServices {
        fmt.Printf("%sTrying %s...%s\n", bcolors.BrightYellow, service, bcolors.Endc)
        resp, err := client.Get(service)
        if err == nil {
            defer resp.Body.Close()
            body, err := io.ReadAll(resp.Body)
            if err == nil {
                ip := strings.TrimSpace(string(body))
                fmt.Printf("%sSuccess! Got IP: %s%s\n", bcolors.BrightGreen, ip, bcolors.Endc)
                return ip, nil
            }
        }
        fmt.Printf("%s%s failed: %v%s\n", bcolors.BrightRed, service, err, bcolors.Endc)
    }
    return "", fmt.Errorf("%sAll IP services failed%s", bcolors.BrightRed, bcolors.Endc)
}

func StartServices() error {
    fmt.Printf("\n%s%s[*] %sStarting %smacchanger, dnsmasq, privoxy, squid, tor.%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Cyan, bcolors.Endc)

    services := []string{
        "tor@default.service",
        "dnsmasq.service",
        "privoxy.service",
        "squid.service",
        "changemac@eth0.service",
    }

    var startedServices []string
    var failedServices []string

    for _, service := range services {
        fmt.Printf("    %s[!] %sStarted: %s%s%s\n", bcolors.Bold, bcolors.Endc, bcolors.Cyan, service, bcolors.Endc)

        cmd := exec.Command("systemctl", "restart", service)
        if err := cmd.Run(); err != nil {
            fmt.Printf("%s%s[!] %sFailed to start %s: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, service, err)
            failedServices = append(failedServices, service)
        } else {
            startedServices = append(startedServices, service)
            time.Sleep(500 * time.Millisecond)
        }
    }

    fmt.Printf("%s%s[+] %sVerifying services ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    time.Sleep(1 * time.Second)

    var runningServices []string
    var notRunningServices []string

    for _, service := range startedServices {
        cmd := exec.Command("systemctl", "is-active", service)
        output, err := cmd.Output()
        if err == nil && strings.TrimSpace(string(output)) == "active" {
            runningServices = append(runningServices, service)
        } else {
            notRunningServices = append(notRunningServices, service)
        }
    }

    if len(runningServices) > 0 {
        fmt.Printf("%s[!] %sRunning services: %s%s%s\n", bcolors.Yellow, bcolors.Endc, bcolors.Cyan, strings.Join(runningServices, ", "), bcolors.Endc)
    }

    if len(notRunningServices) > 0 {
        fmt.Printf("%s%s[!] %sStarted but not running: %s%s%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Yellow, strings.Join(notRunningServices, ", "), bcolors.Endc)
    }

    if len(failedServices) > 0 {
        fmt.Printf("%s[!] %sFailed to start: %s%s%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.Red, strings.Join(failedServices, ", "), bcolors.Endc)
    }

    fmt.Printf("%s%s[*] %sEnabling services to start on boot ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    for _, service := range runningServices {
        exec.Command("systemctl", "enable", service).Run()
    }

    if len(failedServices) > 0 {
        return fmt.Errorf("failed to start %d services: %s", len(failedServices), strings.Join(failedServices, ", "))
    }

    if len(notRunningServices) > 0 {
        return fmt.Errorf("%d services started but not running: %s", len(notRunningServices), strings.Join(notRunningServices, ", "))
    }

    fmt.Printf("%s%s[*] %sAll services started successfully ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    return nil
}

func TorCircuit() error {
    ip, err := SystemIP()
    if err != nil {
        return fmt.Errorf("%sFailed to get IP address: %v%s", bcolors.BrightRed, err, bcolors.Endc)
    }

    if !strings.Contains(CheckIptables(), "torsocks") {
        return fmt.Errorf("\n%s[!] Torsocks not started yet ...%s", bcolors.Colors(), bcolors.Endc)
    }

    if err := exec.Command("service", "tor", "reload").Run(); err != nil {
        return fmt.Errorf("%s[!] Failed to reload Tor service: %v%s", bcolors.BrightRed, err, bcolors.Endc)
    }

    fmt.Printf("%s%s[!] %sScrambling Tor Nodes\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    fmt.Printf("%s%s[*] %sYour new IP appears to be: %s%s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Colors(), ip, bcolors.Endc)
    return nil
}

func KillTor() error {
    defer StopServices()

    trigger := 0
    fmt.Printf("\n%s%s[*] %sReverting to default resolv.conf ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if _, err := os.Stat("/etc/resolv.conf.backup_torsocks"); err == nil {
        trigger++
        if err := exec.Command("mv", "/etc/resolv.conf.backup_torsocks", "/etc/resolv.conf").Run(); err != nil {
            return fmt.Errorf("failed to restore resolv.conf: %v", err)
        }
    }

    fmt.Printf("%s%s[*] %sDropping torrc file & restoring default ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    if _, err := os.Stat("/etc/tor/torrc.bak_torsocks"); err == nil {
        trigger++
        if err := exec.Command("mv", "/etc/tor/torrc.bak_torsocks", "/etc/tor/torrc").Run(); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to restore torrc: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        }
    }

    trigger++
    fmt.Printf("%s%s[*] %sRestoring Default Iptables rules ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if _, err := os.Stat("/etc/iptables_rules_torsocks.bak"); err == nil {
        cmd := exec.Command("bash", "-c", "sudo iptables-restore < /etc/iptables_rules_torsocks.bak")
        if output, err := cmd.CombinedOutput(); err != nil {
            fmt.Printf("%s%s[!] %sFailed to restore iptables from backup: %v\nOutput: %s%s\n", 
                bcolors.Bold, bcolors.Red, bcolors.Endc, err, string(output), bcolors.Endc)
            fmt.Printf("%s%s[*] %sFalling back to default reset...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
            if err := ResetToDefault(true, true); err != nil {
                return fmt.Errorf("%s%s[!] %sFailed to reset to default: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
            }
        } else {
            fmt.Printf("%s%s[*] %sSuccessfully restored from backup%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Endc)
            os.Remove("/etc/iptables_rules_torsocks.bak")
        }
    } else {
        fmt.Printf("%s%s[*] %sNo backup found, resetting to default rules...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        if err := ResetToDefault(true, true); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to reset to default: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        }
    }
    time.Sleep(1 * time.Second)

    if trigger == 0 {
        return fmt.Errorf("%s%sNo instances of torsocks have been executed%s", bcolors.BrightRed, bcolors.Bold, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[*] %sEverything completed... \n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("%s%s[*] %sPreparing to stop all services... \n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    return nil
}


func StopServices() error {
    fmt.Printf("%s%s[+] %sKilling: %smacchanger, dnsmasq, privoxy, squid, tor.%s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Cyan, bcolors.Endc)
    services := []string{
        "changemac@eth0.service",
        "dnsmasq.service",
        "squid.service",
        "privoxy.service",
        "tor@default.service",
    }

    var stoppedServices []string
    var failedServices []string

    for _, service := range services {
        cmd := exec.Command("systemctl", "stop", service)
        if err := cmd.Run(); err != nil {
            fmt.Printf("%s[!] %sFailed to stop %s: %v%s\n", bcolors.BrightRed, bcolors.Endc, service, err, bcolors.Endc)
            failedServices = append(failedServices, service)
        } else {
            stoppedServices = append(stoppedServices, service)
        }
        time.Sleep(100 * time.Millisecond)
    }

    if len(stoppedServices) > 0 {
        fmt.Printf("%s%s[>] %sStopped: %s%s%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Cyan, strings.Join(stoppedServices, ", "), bcolors.Endc)
    }

    if len(failedServices) > 0 {
        fmt.Printf("%s%s[!] %sFailed to stop: %s%s%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Yellow, strings.Join(failedServices, ", "), bcolors.Endc)
    }

    if len(failedServices) > 0 {
        return fmt.Errorf("failed to stop %d services", len(failedServices))
    }
    return nil
}

func StartTorServices() error {
    fmt.Printf("%s%s[*] %sStarting Tor services ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    services := []string{
        "changemac@eth0.service",
        "dnsmasq.service",
        "squid.service",
        "privoxy.service",
        "tor@default.service",
    }

    for _, service := range services {
        exec.Command("sudo", "systemctl", "stop", service).Run()
        time.Sleep(100 * time.Millisecond)

        cmd := exec.Command("sudo", "systemctl", "restart", service)
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("failed to start %s: %v", service, err)
        }

        time.Sleep(1 * time.Second)
        checkCmd := exec.Command("sudo", "systemctl", "is-active", service)
        output, err := checkCmd.Output()
        if err != nil || strings.TrimSpace(string(output)) != "active" {
            return fmt.Errorf("%s%s[!] %sService %s failed to start ...\n", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Cyan, service, bcolors.Endc)
        }
    }
    fmt.Printf("%s%s[+] %sTor services Up and running...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    return nil
}

func CheckServiceStatus() {
    services := []string{
        "dnsmasq.service",
        "squid.service",
        "privoxy.service",
        "tor@default.service",
        "changemac@eth0.service",
    }

    fmt.Printf("%s%s[*] %sChecking service status ...\n\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    for _, service := range services {
        cmd := exec.Command("systemctl", "is-active", service)
        output, _ := cmd.Output()
        status := strings.TrimSpace(string(output))

        var statusColor string
        if status == "active" {
            statusColor = bcolors.Green
        } else {
            statusColor = bcolors.Red
        }
        fmt.Printf("    %s[+] %s%s%s%s: %s%s%s\n", bcolors.Bold, bcolors.Endc, statusColor, status, bcolors.Endc, bcolors.Cyan, service, bcolors.Endc)
    }
    fmt.Printf("\n%s%s[+] %sCompleted Checking service status ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
}

func TorStatus(retryCount int) {
    fmt.Printf("\n%s%s[!] %sChecking Tor network status ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)

    result, err := CheckTorStatus()
    if err != nil {
        if retryCount < 3 {
            fmt.Printf("%s%s[!] %sNetwork issue detected, retrying in 9 seconds ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
            time.Sleep(9 * time.Second)
            TorStatus(retryCount + 1)
            return
        }
        fmt.Printf("%s%s[!] %sFailed after multiple attempts: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        fmt.Printf("%s%s[>] %sCheck your internet connection and try again ...\n", bcolors.Bold, bcolors.BrightYellow, bcolors.Endc)
        return
    }
    displayTorResult(result)
}

func displayTorResult(result map[string]string) {
    fmt.Printf("%s%s[>] %s%sYour IP address: %s%s%s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Yellow, bcolors.Cyan, result["ip"], bcolors.Endc)

    if strings.Contains(result["title"], "Congratulations") {
        fmt.Printf("%s%s[*] %s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, result["title"])
        fmt.Printf("%s%s[+] %sYour connection is secured through Tor ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[!] %s%s..\n", bcolors.Bold, bcolors.Red, bcolors.Endc, result["title"])
        fmt.Printf("%s%s[>] %sWarning: You are NOT using the Tor network!\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        fmt.Printf("%s%s[+] %sFor enhanced privacy, consider using Tor network ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    }

    fmt.Printf("%s", "\n")
}

func CheckTorStatus() (map[string]string, error) {
    client := &http.Client{Timeout: 15 * time.Second}
    resp, err := client.Get("https://check.torproject.org")
    if err != nil {
        return nil, fmt.Errorf("network error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("server returned status: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %v", err)
    }

    result := map[string]string{
        "ip":    extractIP(string(body)),
        "title": extractTitle(string(body)),
    }

    return result, nil
}

func extractIP(body string) string {
    ipRegex := regexp.MustCompile(`IP address[^0-9]*([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
    if matches := ipRegex.FindStringSubmatch(body); len(matches) > 1 {
        return matches[1]
    }

    fallbackRegex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
    if ip := fallbackRegex.FindString(body); ip != "" {
        return ip
    }
    return "[!] Not detected"
}

func extractTitle(body string) string {
    titleRegex := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`)
    if matches := titleRegex.FindStringSubmatch(body); len(matches) > 1 {
        return strings.TrimSpace(matches[1])
    }
    return "Unknown Status"
}

func CheckDefaultRules() (string, string) {
    cmd := exec.Command("bash", "-c", "iptables-save | grep '^\\-' | wc -l")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", string(output)
    }
    return string(output), ""
}

func ConfigFirewall() error {
    fmt.Printf("\n%s%s[*] %sBacking up Iptables.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    FirewallOff, FirewallOn := CheckDefaultRules()
    if FirewallOn != "" {
        return fmt.Errorf("%s\nCan't execute %siptables-save%s. See the reason below.\n%s%s%s", bcolors.BrightRed, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightRed, FirewallOn, bcolors.Endc)
    }

    if strings.TrimSpace(FirewallOff) == "0" {
        fmt.Printf("%s%s[*] %sDefault rules are configured, skipping backup.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[+] %sSaving default firewall rules ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

        cmd := exec.Command("bash", "-c", "sudo iptables-save > /etc/iptables_rules_torsocks.bak")
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to save iptables rules: %v%s", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err, bcolors.Endc)
        }

        fmt.Printf("%s%s[+] %sDone saving firewall rules.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("\n%s%s[*] %sSetting up firewall rules ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if err := SetTorIptablesRules(); err != nil {
        return fmt.Errorf("%s%s[!] %sFailed to set up firewall rules: %v%s", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err, bcolors.Endc)
    }

    fmt.Printf("%s%s[+] %sDone setting up firewall rules ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    return nil
}

func SetTorIptablesRules() error {
    cmd := exec.Command("id", "-u", "debian-tor")
    uidBytes, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("[!] Failed to get tor UID: %v", err)
    }
    torUID := strings.TrimSpace(string(uidBytes))

    transPort := "9040"
    dnsPort := "5353"
    virtAddr := "10.192.0.0/10"

    commands := []string{
        "iptables -P INPUT ACCEPT",
        "iptables -P FORWARD ACCEPT",
        "iptables -P OUTPUT ACCEPT",
        "iptables -F",
        "iptables -X",
        "iptables -Z",
        "iptables -t nat -F",
        "iptables -t nat -X",
        "iptables -t mangle -F",
        "iptables -t mangle -X",
        "iptables -t raw -F",
        "iptables -t raw -X",

        "iptables -t nat -A OUTPUT -d " + virtAddr + " -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports " + transPort,

        "iptables -t nat -A OUTPUT -d 127.0.0.1/32 -p udp -m udp --dport 53 -j REDIRECT --to-ports " + dnsPort + " -m comment --comment \"neutron_triggered\"",

        "iptables -t nat -A OUTPUT -m owner --uid-owner " + torUID + " -j RETURN",
        "iptables -t nat -A OUTPUT -o lo -j RETURN",

        "iptables -t nat -A OUTPUT -d 127.0.0.0/8 -j RETURN",
        "iptables -t nat -A OUTPUT -d 10.0.0.0/8 -j RETURN",
        "iptables -t nat -A OUTPUT -d 172.16.0.0/12 -j RETURN",
        "iptables -t nat -A OUTPUT -d 192.168.0.0/16 -j RETURN",

        "iptables -t nat -A OUTPUT -d 0.0.0.0/8 -j RETURN",
        "iptables -t nat -A OUTPUT -d 100.64.0.0/10 -j RETURN",
        "iptables -t nat -A OUTPUT -d 169.254.0.0/16 -j RETURN",
        "iptables -t nat -A OUTPUT -d 192.0.0.0/24 -j RETURN",
        "iptables -t nat -A OUTPUT -d 192.0.2.0/24 -j RETURN",
        "iptables -t nat -A OUTPUT -d 192.88.99.0/24 -j RETURN",
        "iptables -t nat -A OUTPUT -d 198.18.0.0/15 -j RETURN",
        "iptables -t nat -A OUTPUT -d 198.51.100.0/24 -j RETURN",
        "iptables -t nat -A OUTPUT -d 203.0.113.0/24 -j RETURN",
        "iptables -t nat -A OUTPUT -d 224.0.0.0/4 -j RETURN",
        "iptables -t nat -A OUTPUT -d 240.0.0.0/4 -j RETURN",
        "iptables -t nat -A OUTPUT -d 255.255.255.255/32 -j RETURN",

        "iptables -t nat -A OUTPUT -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports " + transPort,

        "iptables -A INPUT -m state --state ESTABLISHED -j ACCEPT",
        "iptables -A INPUT -i lo -j ACCEPT",
        "iptables -A INPUT -j DROP",

        "iptables -A FORWARD -j DROP",

        "iptables -A OUTPUT -m conntrack --ctstate INVALID -j DROP",
        "iptables -A OUTPUT -m state --state INVALID -j DROP",

        "iptables -A OUTPUT -m state --state ESTABLISHED -j ACCEPT",

        "iptables -A OUTPUT -m owner --uid-owner " + torUID + " -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -m state --state NEW -j ACCEPT",

        "iptables -A OUTPUT -d 127.0.0.1/32 -o lo -j ACCEPT",

        "iptables -A OUTPUT -d 127.0.0.1/32 -p tcp -m tcp --dport " + transPort + " --tcp-flags FIN,SYN,RST,ACK SYN -j ACCEPT",

        "iptables -A OUTPUT -j DROP",

        "iptables -P INPUT DROP",
        "iptables -P FORWARD DROP",
        "iptables -P OUTPUT DROP",
    }
    
    fmt.Printf("%s%s[*] %sSetting up Tor iptables rules...%s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Endc)

    var lastError error
    successCount := 0

    for i, cmdStr := range commands {
        cmd := exec.Command("bash", "-c", "sudo " + cmdStr)
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Printf("%s%s[!] %sFailed command %d: %s - Error: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, i + 1, cmdStr, err, bcolors.Endc)
            if string(output) != "" {
                fmt.Printf("%sOutput: %s%s\n", bcolors.Yellow, string(output), bcolors.Endc)
            }
            lastError = err
        } else {
            successCount++
        }
        time.Sleep(50 * time.Millisecond)
    }
    
    fmt.Printf("%s%s[*] %sCompleted: %d/%d commands executed successfully%s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, successCount, len(commands), bcolors.Endc)

    if lastError != nil {
        return fmt.Errorf("%s%s[!] %sSome iptables commands failed during Tor setup%s", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Endc)
    }

    fmt.Printf("%s%s[+] %sTor iptables rules configured successfully!%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Endc)
    return nil
}

func ResetToDefault(overridePass bool, resetAsChildFunc bool) error {
    if !overridePass {
        var resetConsent string
        fmt.Printf("%s\nThis will overwrite all of your existing rules %sY(do it)%s/%sN(exit)%s: ", bcolors.Colors(), bcolors.Green, bcolors.Endc, bcolors.BrightRed, bcolors.Endc)
        fmt.Scanln(&resetConsent)

        if strings.ToLower(resetConsent) == "n" {
            fmt.Printf("%sCopy that..\n%s", bcolors.BrightRed, bcolors.Endc)
            return nil
        } else if strings.ToLower(resetConsent) != "y" {
            return fmt.Errorf("invalid input, aborting reset")
        }

        time.Sleep(1 * time.Second)
        fmt.Printf("%sBacking up current rules, just in case..%s\n", bcolors.Magenta, bcolors.Endc)

        FirewallOff, FirewallOn := CheckDefaultRules()
        if FirewallOn != "" {
            return fmt.Errorf("%sError while checking existing rules; %sexiting..\n%sError message: %s%s%s", bcolors.BrightRed, bcolors.BrightYellow, bcolors.Yellow, bcolors.Colors(), FirewallOn, bcolors.Endc)
        }

        if strings.TrimSpace(FirewallOff) != "0" {
            fileNameID := time.Now().Format("01_02_2006-15:04:05")
            backupFile := "/tmp/iptables_" + fileNameID + ".rules"
            cmd := exec.Command("bash", "-c", "sudo iptables-save > "+backupFile)
            if err := cmd.Run(); err != nil {
                return fmt.Errorf("%sFailed to save iptables rules: %v%s", bcolors.BrightRed, err, bcolors.Endc)
            }
            fmt.Printf("%sSaved in %s/tmp%s as %siptables_%s.rules%s\n", bcolors.BrightCyan, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightRed, fileNameID, bcolors.Endc)
        } else {
            fmt.Printf("%s Default rules are set, backup not required :)%s\n", bcolors.BrightYellow, bcolors.Endc)
        }

        fmt.Printf("%s%sResetting Iptables%s\n", bcolors.BrightYellow, bcolors.Bold, bcolors.Endc)
    }

    commands := []string{
        "iptables -F",
        "iptables -X",
        "iptables -t nat -F",
        "iptables -t nat -X",
        "iptables -t mangle -F",
        "iptables -t mangle -X",
        "iptables -t raw -F",
        "iptables -t raw -X",
        "iptables -P INPUT ACCEPT",
        "iptables -P OUTPUT ACCEPT",
        "iptables -P FORWARD ACCEPT",
    }

    fmt.Printf("%s%s[*] %sFlushing iptables rules...%s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Endc)

    var lastError error
    for _, cmdStr := range commands {
        cmd := exec.Command("bash", "-c", "sudo " + cmdStr)
        if output, err := cmd.CombinedOutput(); err != nil {
            fmt.Printf("%s%s[!] %sCommand failed: %s - Error: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, cmdStr, err, bcolors.Endc)
            if string(output) != "" {
                fmt.Printf("%sOutput: %s%s\n", bcolors.Yellow, string(output), bcolors.Endc)
            }
            lastError = err
        }
        time.Sleep(100 * time.Millisecond)
    }

    time.Sleep(500 * time.Millisecond)

    checkCmd := exec.Command("bash", "-c", "sudo iptables -L -n | head -10")
    if output, err := checkCmd.CombinedOutput(); err == nil {
        lines := strings.Count(string(output), "\n")
        if lines <= 8 {
            fmt.Printf("%s%s[*] %sIptables successfully reset to default%s\n", 
                bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Endc)
        } else {
            fmt.Printf("%s%s[!] %sIptables may not have been fully reset%s\n", 
                bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Endc)
        }
    }

    if lastError != nil {
        return fmt.Errorf("%s%s[!] %sSome iptables commands failed during reset ...", bcolors.Bold, bcolors.BrightRed, bcolors.Endc)
    }

    if !resetAsChildFunc {
        fmt.Printf("%s%s[*] %sSuccessfully reset Iptables to default ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    }

    return nil
}

func checkTorUsage() (TorResult, error) {
    result := TorResult{}

    cmd := exec.Command("timeout", "10", "torify", "curl", "-s", "https://check.torproject.org/api/ip")
    output, err := cmd.Output()
    if err != nil {
        return result, fmt.Errorf("tor check failed: %v", err)
    }

    var torResponse struct {
        IP      string `json:"IP"`
        IsTor   bool   `json:"IsTor"`
        Country string `json:"Country"`
    }

    if err := json.Unmarshal(output, &torResponse); err != nil {
        return result, fmt.Errorf("[!] Failed to parse tor response: %v", err)
    }

    result.IsUsingTor = torResponse.IsTor
    result.ExitIP = torResponse.IP
    result.ExitCountry = torResponse.Country

    return result, nil
}

func performDNSLeakTests() ([]DNSServer, bool) {
    var servers []DNSServer
    leakDetected := false

    testDomains := []string{
        "google.com",
        "facebook.com",
        "amazon.com",
        "microsoft.com",
    }

    dnsProviders := []struct {
        name string
        ip   string
    }{
        {"Google DNS", "8.8.8.8"},
        {"Cloudflare", "1.1.1.1"},
        {"OpenDNS", "208.67.222.222"},
        {"Quad9", "9.9.9.9"},
    }

    uniqueServers := make(map[string]DNSServer)

    for _, domain := range testDomains {
        for _, provider := range dnsProviders {
            server, err := testDNSResolution(domain, provider.ip)
            if err == nil && server.IP != "" {
                if !isTorExitNode(server.IP) {
                    server.ViaTor = false
                    leakDetected = true
                } else {
                    server.ViaTor = true
                }
                uniqueServers[server.IP] = server
            }
            time.Sleep(100 * time.Millisecond)
        }
    }

    for _, server := range uniqueServers {
        servers = append(servers, server)
    }

    return servers, leakDetected
}

func testDNSResolution(domain, dnsServer string) (DNSServer, error) {
    server := DNSServer{IP: dnsServer}

    cmd := exec.Command("dig", "+short", domain, "@"+dnsServer)
    output, err := cmd.Output()
    if err != nil {
        return server, err
    }

    if strings.TrimSpace(string(output)) == "" {
        return server, fmt.Errorf("no DNS response from %s", dnsServer)
    }

    ipInfo, err := getIPInfo(dnsServer)
    if err == nil {
        server.ISP = ipInfo.ISP
        server.Country = ipInfo.Country
        server.ASN = ipInfo.ASN
    } else {
        server.ISP = "Unknown"
        server.Country = "Unknown"
        server.ASN = "Unknown"
    }

    return server, nil
}

func getIPInfo(ip string) (IPGeoInfo, error) {
    info := IPGeoInfo{IP: ip}

    for _, apiTemplate := range GeoAPIs {
        apiUrl := fmt.Sprintf(apiTemplate, ip)
        
        resp, err := http.Get(apiUrl)
        if err != nil {
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            continue
        }

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            continue
        }

        var geoData map[string]interface{}
        if err := json.Unmarshal(body, &geoData); err != nil {
            continue
        }

        info.Country = getString(geoData, "country", "country_name", "countryName")
        info.Region = getString(geoData, "region", "regionName", "state_prov")
        info.City = getString(geoData, "city")
        info.ISP = getString(geoData, "isp", "org", "asn", "asn_org")
        info.Latitude = getFloat(geoData, "lat", "latitude")
        info.Longitude = getFloat(geoData, "lon", "longitude", "lng")
        info.Continent = getString(geoData, "continent", "continent_name", "continentName")

        if strings.Contains(info.ISP, "AS") {
            parts := strings.Split(info.ISP, " ")
            if len(parts) > 0 {
                info.ASN = parts[0]
            }
        }
        
        return info, nil
    }

    return info, fmt.Errorf("could not retrieve geolocation data")
}

func isTorExitNode(ip string) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "53"), 2*time.Second)
    if err != nil {
        return true
    }
    conn.Close()
    return false
}

func testSystemDNS() ([]DNSServer, bool) {
    var servers []DNSServer
    leakDetected := false

    dnsServers, err := getSystemDNSServers()
    if err != nil {
        return servers, true
    }

    for _, dnsIP := range dnsServers {
        server := DNSServer{
            IP:     dnsIP,
            ViaTor: isTorExitNode(dnsIP),
        }

        if !server.ViaTor {
            leakDetected = true
            ipInfo, err := getIPInfo(dnsIP)
            if err == nil {
                server.ISP = ipInfo.ISP
                server.Country = ipInfo.Country
                server.ASN = ipInfo.ASN
            } else {
                server.ISP = "Unknown"
                server.Country = "Unknown"
                server.ASN = "Unknown"
            }
        }

        servers = append(servers, server)
    }

    return servers, leakDetected
}

func getSystemDNSServers() ([]string, error) {
    content, err := os.ReadFile("/etc/resolv.conf")
    if err != nil {
        return nil, err
    }

    var servers []string
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "nameserver") {
            parts := strings.Fields(line)
            if len(parts) >= 2 {
                servers = append(servers, parts[1])
            }
        }
    }

    return servers, nil
}

func TestDNSLeak() (string, error) {
    result := DNSLeakTestResult{
        Timestamp: time.Now().Format(time.RFC3339),
        Status:    "completed",
        TestType:  "comprehensive",
    }

    torResult, err := checkTorUsage()
    if err != nil {
        result.Status = "partial"
        result.TorCheck.IsUsingTor = false
    } else {
        result.TorCheck = torResult
    }

    dnsServers, leakDetected := performDNSLeakTests()
    result.DNSServers = dnsServers
    result.LeakDetected = leakDetected

    coloredJSON := formatDNSLeakResultColored(result)
    return coloredJSON, nil
}

func formatDNSLeakResultColored(result DNSLeakTestResult) string {
    dnsServersStr := "null"
    if result.DNSServers != nil && len(result.DNSServers) > 0 {
        dnsServersStr = formatDNSServersColored(result.DNSServers)
    }

    return fmt.Sprintf(`%s
{
  "timestamp": "%s%s%s",
  "status": "%s%s%s",
  "test_type": "%s%s%s",
  "dns_servers": %s,
  "leak_detected": %s%t%s,
  "tor_check": {
    "is_using_tor": %s%t%s,
    "exit_ip": "%s%s%s"%s
  }
}
%s`, bcolors.Endc, bcolors.Cyan, result.Timestamp, bcolors.Endc, bcolors.Cyan, result.Status, bcolors.Endc, bcolors.Cyan, result.TestType, bcolors.Endc, dnsServersStr, bcolors.Cyan, result.LeakDetected, bcolors.Endc, bcolors.Cyan, result.TorCheck.IsUsingTor, bcolors.Endc, bcolors.Cyan, result.TorCheck.ExitIP, bcolors.Endc, formatTorCheckCountry(result.TorCheck), bcolors.Endc)
}

func formatTorCheckCountry(torCheck TorResult) string {
    if torCheck.ExitCountry != "" {
        return fmt.Sprintf(`,
    "exit_country": "%s%s%s"`, bcolors.Cyan, torCheck.ExitCountry, bcolors.Endc)
    }
    return ""
}

func formatDNSServersColored(servers []DNSServer) string {
    var builder strings.Builder
    builder.WriteString("[\n")

    for i, server := range servers {
        builder.WriteString(fmt.Sprintf(`    {
      "ip": "%s%s%s",
      "isp": "%s%s%s",
      "country": "%s%s%s",
      "asn": "%s%s%s",
      "via_tor": %s%t%s
    }`, bcolors.Cyan, server.IP, bcolors.Endc, bcolors.Cyan, server.ISP, bcolors.Endc, bcolors.Cyan, server.Country, bcolors.Endc, bcolors.Cyan, server.ASN, bcolors.Endc, bcolors.Cyan, server.ViaTor, bcolors.Endc))

        if i < len(servers)-1 {
            builder.WriteString(",\n")
        }
    }

    builder.WriteString("\n  ]")
    return builder.String()
}

func CheckLeakTest() (*DNSLeakTestResult, error) {
    fmt.Printf("\n%s%s[*] %sTesting DNS leaks ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    fmt.Printf("%s%s[+] %sPleas be patient. Geting results ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    result, err := TestDNSLeak()
    if err != nil {
        return nil, fmt.Errorf("%s%s[!] %sDNS leak test failed: %v", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
    }

    fmt.Printf("%s\n", result)

    cleanJSON := stripColors(result)
    var leakResult DNSLeakTestResult
    if err := json.Unmarshal([]byte(cleanJSON), &leakResult); err != nil {
        return nil, fmt.Errorf("%s%s[!] %sFailed to parse DNS leak results: %v", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
    }

    if leakResult.LeakDetected {
        fmt.Printf("%s%s[!] %s%s%sDNS LEAK DETECTED! %sFound %d non-Tor DNS servers!", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Yellow, bcolors.Blink, bcolors.Endc, len(leakResult.DNSServers))
        return &leakResult, fmt.Errorf("%s%s[!] %sDNS leak detected", bcolors.Bold, bcolors.Red, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[*] %sCongratualion. No DNS leaks detected.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        return &leakResult, nil
    }
}

func stripColors(text string) string {
    re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return re.ReplaceAllString(text, "")
}
