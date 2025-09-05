// John 3:16
package securities

import (
    "encoding/json"
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
    "time"

    "banners"
    "bcolors"
    "menus"
    "scriptures"
    "subprocess"
    "utils"
)

var Function string

var GeoAPIs = []string{
    "http://ip-api.com/json/%s",
    "https://ipapi.co/%s/json/",
    "https://freeipapi.com/api/json/%s",
}

var IpServices = []string{
    "https://ident.me",
    "https://ipinfo.io/ip",
    "https://api.ipify.org",
    "https://icanhazip.com",
    "https://ifconfig.me/ip",
    "https://checkip.amazonaws.com",
}

type stringMatcher struct {
    names  []string
    action func()
}

func Torsocks() {
    for {
        fmt.Printf("%s%s%safr%s%s torsocks(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        {[]string{"g", "t", "guide", "tutorial"}, utils.BrowseTutorials},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, menus.HelpInfoTorsocks},
        {[]string{"m", "menu"}, menus.MenuTwo},
        {[]string{"option", "options", "show option", "show options"}, menus.TorsocksOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListTorsocksFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() { AnonimityFunctions("setups") }},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, menus.HelpInfoTorsocksSetups},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run vanish", "use vanish", "exec vanish", "start vanish", "launch vanish", "exploit vanish", "execute vanish"}, func() { AnonimityFunctions("vanish") }},
        {[]string{"? 2", "info 2", "help 2", "vanish", "info vanish", "help vanish"}, menus.HelpInfoTorsocksVanish},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run status", "use status", "exec status", "start status", "launch status", "exploit status", "execute status"}, func() { AnonimityFunctions("status") }},
        {[]string{"? 3", "info 3", "help 3", "status", "info status", "help status"}, menus.HelpInfoTorsocksStatus},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run torip", "use torip", "exec torip", "start torip", "launch torip", "exploit torip", "execute torip"}, func() { AnonimityFunctions("torip") }},
        {[]string{"? 4", "info 4", "help 4", "torip", "info torip", "help torip"}, menus.HelpInfoTorsocksTorIp},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run chains", "use chains", "exec chains", "start chains", "launch chains", "exploit chains", "execute chains"}, func() { AnonimityFunctions("chains") }},
        {[]string{"? 5", "info 5", "help 5", "chains", "info chains", "help chains"}, menus.HelpInfoTorsocksChains},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run reload", "use reload", "exec reload", "start reload", "launch reload", "exploit reload", "execute reload"}, func() { AnonimityFunctions("reload") }},
        {[]string{"? 6", "info 6", "help 6", "reload", "info reload", "help reload"}, menus.HelpInfoTorsocksReload},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run exitnode", "use exitnode", "exec exitnode", "start exitnode", "launch exitnode", "exploit exitnode", "execute exitnode"}, func() { AnonimityFunctions("exitnode") }},
        {[]string{"? 7", "info 7", "help 7", "exitnode", "info exitnode", "help exitnode"}, menus.HelpInfoTorsocksExitnode},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run restore", "use restore", "exec restore", "start restore", "launch restore", "exploit restore", "execute restore"}, func() { AnonimityFunctions("restore") }},
        {[]string{"? 8", "info 8", "help 8", "restore", "info restore", "help restore"}, menus.HelpInfoTorsocksRestore},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run stop", "use stop", "exec stop", "start stop", "launch stop", "exploit stop", "execute stop"}, func() { AnonimityFunctions("stop") }},
        {[]string{"? 9", "info 9", "help 9", "stop", "info stop", "help stop"}, menus.HelpInfoTorsocksStop},

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
        "proxies":     &utils.Proxies,
        "func":        &Function,
        "funcs":       &Function,
        "module":      &Function,
        "output":      &utils.OutPutDir,
        "outputlog":   &utils.OutPutDir,
        "outputlogs":  &utils.OutPutDir,
        "function":    &Function,
        "functions":   &Function,
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
        "proxies":    &utils.Proxies,
        "func":       &Function,
        "funcs":      &Function,
        "module":     &Function,
        "output":     &utils.OutPutDir,
        "outputlog":  &utils.OutPutDir,
        "outputlogs": &utils.OutPutDir,
        "function":   &Function,
        "functions":  &Function,
    }
    defaultValues := map[string]string{
        "func":       "",
        "proxies":    "",
        "funcs":      "",
        "module":     "",
        "output":     "",
        "function":   "",
        "functions":  "",
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
    if Function == "" {
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    AnonimityFunctions(Function)
}

func AnonimityFunctions(functionName string, args ...interface{}) {
    os.MkdirAll(utils.SecLogs, os.ModePerm)

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            IFACE:       utils.IFace,
            DISTRO:      utils.Distro,
            OUTPUTLOGS:  utils.SecLogs,
            PROXIES:     utils.Proxies,
            FUNCTION:    functionName,
        }, true, false)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            fmt.Printf("%sError setting proxy: %v%s\n", bcolors.Red, err, bcolors.Endc)
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
        "setups":   func() { banners.GraphicsTorNet(); subprocess.Popen("apt-get update; apt-get install -y tor squid privoxy dnsmasq iptables isc-dhcp-client isc-dhcp-server") },
        "vanish":   func() { banners.GraphicsTorNet(); ConfigureResolv(); ConfigChangeMac(); ConfigDhclient(); ConfigDnsmasq(); ConfigSquid(); ConfigPrivoxy(); ConfigTorrc(); ConfigFirewall(); StartTorServices(); CheckServiceStatus(); CheckIP(); TorStatus(0) },
        "status":   func() { banners.GraphicsTorNet(); CheckServiceStatus(); TorStatus(0) },
        "torip":    func() { banners.GraphicsTorNet(); CheckIP(); TorStatus(0) },
        "exitnode": func() { banners.GraphicsTorNet(); TorCircuit() },
        "restore":  func() { banners.GraphicsTorNet(); ResetToDefault(false, false) },
        "chains":   func() { banners.GraphicsTorNet(); subprocess.Popen("tail -vf /var/log/privoxy/logfile") },
        "reload":   func() { banners.GraphicsTorNet(); ConfigTorrc(); ConfigFirewall(); CheckIP() },
        "stop":     func() { banners.GraphicsTorNet(); KillTor(); CheckServiceStatus(); TorStatus(0)},

        "1": func() { banners.GraphicsTorNet(); subprocess.Popen("apt-get update; apt-get install -y tor squid privoxy dnsmasq iptables isc-dhcp-client isc-dhcp-server") },
        "2": func() { banners.GraphicsTorNet(); ConfigureResolv(); ConfigChangeMac(); ConfigDhclient(); ConfigDnsmasq(); ConfigSquid(); ConfigPrivoxy(); ConfigTorrc(); ConfigFirewall(); StartTorServices(); CheckServiceStatus(); CheckIP(); TorStatus(0) },
        "3": func() { banners.GraphicsTorNet(); CheckServiceStatus(); TorStatus(0) },
        "4": func() { banners.GraphicsTorNet(); CheckIP(); TorStatus(0) },
        "5": func() { banners.GraphicsTorNet(); TorCircuit() },
        "6": func() { banners.GraphicsTorNet(); ResetToDefault(false, false) },
        "7": func() { banners.GraphicsTorNet(); subprocess.Popen("tail -vf /var/log/privoxy/logfile") },
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
        menus.ListTorsocksFunctions()
        return
    }

    lowerInput := strings.ToLower(functionName)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("\n%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, functionName, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, functionName)
    menus.ListTorsocksFunctions()
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
        subprocess.Popen("cp -r /etc/dhcp/dhclient.conf /etc/dhcp/dhclient.conf.bak_africana")
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
        subprocess.Popen("cp -r /etc/dnsmasq.conf /etc/dnsmasq.conf.bak_africana")
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
        subprocess.Popen("cp -r /etc/privoxy/config /etc/privoxy/config.bak_africana")
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
        subprocess.Popen("cp -r /etc/squid/squid.conf /etc/squid/squid.conf.bak_africana")
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
    fmt.Printf("\n%s%s[*] %sConfiguring torrc.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

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

    fmt.Printf("\n%s%s[*] %sConfiguring resolv.conf.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

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
    fmt.Printf("\n%s%s[*] %sCheck public IP address.", bcolors.Bold, bcolors.Green, bcolors.Endc)
    fmt.Printf("\n%s%s[+] %sCheck Your system geolocation.\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
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
    for _, apiTemplate := range GeoAPIs {
        apiUrl := fmt.Sprintf(apiTemplate, ip)
        startTime := time.Now()

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

        responseTime := time.Since(startTime).Round(time.Millisecond)
        now := time.Now()

        country := getString(geoData, "country", "country_name", "countryName")
        region := getString(geoData, "region", "regionName", "state_prov")
        city := getString(geoData, "city")
        isp := getString(geoData, "isp", "org", "asn", "asn_org")
        lat := getFloat(geoData, "lat", "latitude")
        lon := getFloat(geoData, "lon", "longitude", "lng")
        continent := getString(geoData, "continent", "continent_name", "continentName")

        fmt.Printf(`
{
  "ip_address": "%s",
  "geolocation": {
    "continent": "%s",
    "country": "%s",
    "region": "%s",
    "city": "%s",
    "coordinates": {
      "latitude": %.6f,
      "longitude": %.6f
    },
    "isp": "%s"
  },
  "service_provider": "%s",
  "timestamp": {
    "unix": %d,
    "iso8601": "%s",
    "human_readable": "%s",
    "timezone": "%s"
  },
  "query_info": {
    "success": true,
    "response_time": "%s"
  }
}
`, ip, continent, country, region, city, lat, lon, isp, strings.TrimPrefix(apiUrl, "http://"), now.Unix(), now.Format(time.RFC3339), now.Format("Monday, January 2, 2006 at 3:04:05 PM"), now.Format("MST"), responseTime)
        return
    }
    fmt.Printf("%sError: Could not retrieve geolocation data%s\n", bcolors.Red, bcolors.Endc)
}

func getString(data map[string]interface{}, keys ...string) string {
    for _, key := range keys {
        if value, exists := data[key]; exists {
            if str, ok := value.(string); ok {
                return str
            }
        }
    }
    return "N/A"
}

func getFloat(data map[string]interface{}, keys ...string) float64 {
    for _, key := range keys {
        if value, exists := data[key]; exists {
            if f, ok := value.(float64); ok {
                return f
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

    fmt.Println(bcolors.Green + "\nScrambling Tor Nodes" + bcolors.Endc)
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
    fmt.Printf("\n%s%s[!] %sChecking Tor network status.\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)

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
    fmt.Printf("%s%s[*] %s%sYour IP address: %s%s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Yellow, bcolors.Cyan, result["ip"], bcolors.Endc)

    if strings.Contains(result["title"], "Congratulations") {
        fmt.Printf("%s%s[*] %s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, result["title"])
        fmt.Printf("%s%s[+] %sYour connection is secured through Tor ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[!] %s%s\n", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, result["title"])
        fmt.Printf("%s%s[>] %sWarning: You are NOT using the Tor network!\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        fmt.Printf("%s%s[+] %sFor enhanced privacy, consider using Tor network ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    }

    fmt.Println()
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
        return fmt.Errorf("%s\nCan't execute %siptables-save%s. See the reason below.\n%s%s%s",
            bcolors.BrightRed, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightRed, FirewallOn, bcolors.Endc)
    }

    if strings.TrimSpace(FirewallOff) == "0" {
        fmt.Printf("%s%s[*] %sDefault rules are configured, skipping.", bcolors.Bold, bcolors.Green, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[+] %sSaving default firewall rules ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

        cmd := exec.Command("bash", "-c", "iptables-save > /etc/iptables_rules_torsocks.bak")
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to save iptables rules: %s%v", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err)
        }

        fmt.Printf("%s%s[+] %sDone saving firewall rules.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    }

    InnOutRules := `
#!/bin/bash
# Set variables
tor_uid=$(id -u debian-tor)
trans_port="9040"
dns_port="5353"
virtual_address="10.192.0.0/10"

# Define non_tor as an array (corrected)
non_tor=("127.0.0.0/8" "10.0.0.0/8" "172.16.0.0/12" "192.168.0.0/16")

# Flush current iptables rules
iptables -F
iptables -X
iptables -t nat -F
iptables -t nat -X
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT

## *nat OUTPUT (For local redirection)
# nat .onion addresses
iptables -t nat -A OUTPUT -d $virtual_address -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports $trans_port

# nat dns requests to Tor
iptables -t nat -A OUTPUT -d 127.0.0.1/32 -p udp -m udp --dport 53 -j REDIRECT --to-ports $dns_port

# Don't nat the Tor process, the loopback, or the local network
iptables -t nat -A OUTPUT -m owner --uid-owner $tor_uid -j RETURN
iptables -t nat -A OUTPUT -o lo -j RETURN

# Allow lan access for hosts in $non_tor (FIXED - process each network individually)
for lan in "${non_tor[@]}"; do
    iptables -t nat -A OUTPUT -d "$lan" -j RETURN
done

# Redirects all other pre-routing and output to Tor's TransPort
iptables -t nat -A OUTPUT -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports $trans_port

## *filter INPUT
iptables -A INPUT -m conntrack --ctstate ESTABLISHED -j ACCEPT
iptables -A INPUT -i lo -j ACCEPT

# Drop everything else
iptables -A INPUT -j DROP

## *filter FORWARD
iptables -A FORWARD -j DROP

## *filter OUTPUT
# Fix for potential kernel transproxy packet leaks
iptables -A OUTPUT -m conntrack --ctstate INVALID -j DROP
iptables -A OUTPUT -m conntrack --ctstate INVALID -j DROP
iptables -A OUTPUT -m conntrack --ctstate ESTABLISHED -j ACCEPT

# Allow Tor process output
iptables -A OUTPUT -m owner --uid-owner $tor_uid -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -m conntrack --ctstate NEW -j ACCEPT

# Allow loopback output
iptables -A OUTPUT -d 127.0.0.1/32 -o lo -j ACCEPT

# Tor transproxy magic
iptables -A OUTPUT -d 127.0.0.1/32 -p tcp -m tcp --dport $trans_port --tcp-flags FIN,SYN,RST,ACK SYN -j ACCEPT

# Allow DNS queries (CRITICAL FIX - add your DNS server IP)
iptables -A OUTPUT -p udp --dport 53 -j ACCEPT

# Drop everything else
iptables -A OUTPUT -j DROP

## Set default policies to DROP
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT DROP
`
    fmt.Printf("\n%s%s[*] %sSetting up firewall rules ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    subprocess.Popen(InnOutRules)
    fmt.Printf("%s%s[+] %sDone setting up firewall rules ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    return nil
}

func ResetToDefault(overridePass bool, resetAsChildFunc bool) error {
    if !overridePass {
        var resetConsent string
        fmt.Printf("%s\nThis will overwrite all of your existing rules %sY(do it)%s/%sN(exit)%s: ", 
            bcolors.Colors(), bcolors.Green, bcolors.Endc, bcolors.BrightRed, bcolors.Endc)
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
            return fmt.Errorf("%sError while checking existing rules; %sexiting..\n%sError message: %s%s%s", 
                bcolors.BrightRed, bcolors.BrightYellow, bcolors.Yellow, bcolors.Colors(), FirewallOn, bcolors.Endc)
        }

        if strings.TrimSpace(FirewallOff) != "0" {
            fileNameID := time.Now().Format("01_02_2006-15:04:05")
            backupFile := "/tmp/iptables_" + fileNameID + ".rules"
            cmd := exec.Command("bash", "-c", "sudo iptables-save > "+backupFile)
            if err := cmd.Run(); err != nil {
                return fmt.Errorf("%sFailed to save iptables rules: %v%s", bcolors.BrightRed, err, bcolors.Endc)
            }
            fmt.Printf("%sSaved in %s/tmp%s as %siptables_%s.rules%s\n", 
                bcolors.BrightCyan, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightRed, fileNameID, bcolors.Endc)
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
        cmd := exec.Command("bash", "-c", "sudo "+cmdStr)
        if output, err := cmd.CombinedOutput(); err != nil {
            fmt.Printf("%s%s[!] %sCommand failed: %s - Error: %v%s\n", 
                bcolors.Bold, bcolors.Red, bcolors.Endc, cmdStr, err, bcolors.Endc)
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
        return fmt.Errorf("%sSome iptables commands failed during reset%s", bcolors.BrightRed, bcolors.Endc)
    }

    if !resetAsChildFunc {
        fmt.Printf("%s Successfully reset Iptables to default :)%s\n", bcolors.BrightBlue, bcolors.Endc)
    }

    return nil
}
