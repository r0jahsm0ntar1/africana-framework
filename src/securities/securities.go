// John 3:16
package securities

import (
    "io"
    "os"
    "fmt"
    "net"
    "time"
    "menus"
    "net/url"
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
    "path/filepath"
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

type ConfigBackup struct {
    OriginalPath string
    BackupPath   string
    ServiceName  string
    RestartCmd   string
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
    Timestamp    string     `json:"timestamp"`
    Status       string     `json:"status"`
    TestType     string     `json:"test_type"`
    DNSServers   []DNSServer `json:"dns_servers"`
    LeakDetected bool       `json:"leak_detected"`
    TorCheck     TorResult  `json:"tor_check"`
}

type DNSServer struct {
    IP       string `json:"ip"`
    ISP      string `json:"isp,omitempty"`
    Country  string `json:"country,omitempty"`
    ASN      string `json:"asn,omitempty"`
    ViaTor   bool   `json:"via_tor"`
}

type TorResult struct {
    IsUsingTor  bool   `json:"is_using_tor"`
    ExitIP      string `json:"exit_ip,omitempty"`
    ExitCountry string `json:"exit_country,omitempty"`
}

type stringMatcher struct {
    names  []string
    action func()
}

const (
    NOTIFY_INFO    = "INFO"
    NOTIFY_WARNING = "WARN"
    NOTIFY_ERROR   = "ERR"
    NOTIFY_SUCCESS = "OK"
    NOTIFY_DEBUG   = "DBG"
    NOTIFY_TOR     = "TOR"
    NOTIFY_DNS     = "DNS"
    NOTIFY_FW      = "FW"
    NOTIFY_PROXY   = "PROXY"
    NOTIFY_SQL     = "SQL"
    NOTIFY_NET     = "NET"
    NOTIFY_CONFIG  = "CFG"
)

func Notify(module, message string, args ...interface{}) {
    timestamp := time.Now().Format("15:04:05")
    
    moduleColors := map[string]string{
        NOTIFY_INFO:    bcolors.BrightBlue,
        NOTIFY_WARNING: bcolors.BrightYellow,
        NOTIFY_ERROR:   bcolors.BrightRed,
        NOTIFY_SUCCESS: bcolors.BrightGreen,
        NOTIFY_DEBUG:   bcolors.BrightCyan,
        NOTIFY_TOR:     bcolors.BrightMagenta,
        NOTIFY_DNS:     bcolors.BrightCyan,
        NOTIFY_FW:      bcolors.BrightRed,
        NOTIFY_PROXY:   bcolors.BrightBlue,
        NOTIFY_SQL:     bcolors.BrightYellow,
        NOTIFY_NET:     bcolors.BrightGreen,
        NOTIFY_CONFIG:  bcolors.BrightWhite,
    }
    
    color := moduleColors[module]
    if color == "" {
        color = bcolors.White
    }

    fmt.Printf("\n%s[%s%s%s] %s[%s%s%s] %s%s%s", bcolors.DarkGray, bcolors.White, timestamp, bcolors.DarkGray, bcolors.DarkGray, color, module, bcolors.DarkGray, bcolors.White, message, bcolors.Endc)
}

func NotifyShort(module, message string, args ...interface{}) {
    moduleColors := map[string]string{
        NOTIFY_INFO:    bcolors.BrightBlue,
        NOTIFY_WARNING: bcolors.BrightYellow,
        NOTIFY_ERROR:   bcolors.BrightRed,
        NOTIFY_SUCCESS: bcolors.BrightGreen,
        NOTIFY_DEBUG:   bcolors.BrightCyan,
        NOTIFY_TOR:     bcolors.BrightMagenta,
        NOTIFY_DNS:     bcolors.BrightCyan,
        NOTIFY_FW:      bcolors.BrightRed,
        NOTIFY_PROXY:   bcolors.BrightBlue,
        NOTIFY_SQL:     bcolors.BrightYellow,
        NOTIFY_NET:     bcolors.BrightGreen,
        NOTIFY_CONFIG:  bcolors.BrightWhite,
    }
    
    color := moduleColors[module]
    if color == "" {
        color = bcolors.White
    }

    fmt.Printf("\n%s[%s%s%s] %s%s%s", bcolors.DarkGray, color, module, bcolors.DarkGray, bcolors.White, message, bcolors.Endc)
}

func NotifyStatus(module, message string, success bool) {
    moduleColors := map[string]string{
        NOTIFY_INFO:    bcolors.BrightBlue,
        NOTIFY_WARNING: bcolors.BrightYellow,
        NOTIFY_ERROR:   bcolors.BrightRed,
        NOTIFY_SUCCESS: bcolors.BrightGreen,
        NOTIFY_DEBUG:   bcolors.BrightCyan,
        NOTIFY_TOR:     bcolors.BrightMagenta,
        NOTIFY_DNS:     bcolors.BrightCyan,
        NOTIFY_FW:      bcolors.BrightRed,
        NOTIFY_PROXY:   bcolors.BrightBlue,
        NOTIFY_SQL:     bcolors.BrightYellow,
        NOTIFY_NET:     bcolors.BrightGreen,
        NOTIFY_CONFIG:  bcolors.BrightWhite,
    }
    
    color := moduleColors[module]
    if color == "" {
        color = bcolors.White
    }
    
    icon := "✓"
    iconColor := bcolors.Green
    if !success {
        icon = "✗"
        iconColor = bcolors.Red
    }
    
    fmt.Printf("\n%s[%s%s%s] %s[%s%s%s] %s%s%s", bcolors.DarkGray, color, module, bcolors.DarkGray, bcolors.DarkGray, iconColor, icon, bcolors.DarkGray, bcolors.White, message, bcolors.Endc)
}

func NotifyProgress(module, message string, current, total int) {
    moduleColors := map[string]string{
        NOTIFY_INFO:    bcolors.BrightBlue,
        NOTIFY_WARNING: bcolors.BrightYellow,
        NOTIFY_ERROR:   bcolors.BrightRed,
        NOTIFY_SUCCESS: bcolors.BrightGreen,
        NOTIFY_DEBUG:   bcolors.BrightCyan,
        NOTIFY_TOR:     bcolors.BrightMagenta,
        NOTIFY_DNS:     bcolors.BrightCyan,
        NOTIFY_FW:      bcolors.BrightRed,
        NOTIFY_PROXY:   bcolors.BrightBlue,
        NOTIFY_SQL:     bcolors.BrightYellow,
        NOTIFY_NET:     bcolors.BrightGreen,
        NOTIFY_CONFIG:  bcolors.BrightWhite,
    }
    
    color := moduleColors[module]
    if color == "" {
        color = bcolors.White
    }
    
    percent := (current * 100) / total
    bar := strings.Repeat("█", percent/5)
    bar += strings.Repeat("░", 20-percent/5)
    
    fmt.Printf("\r%s[%s%s%s] %s%s %s[%s%s%s] %s%d%%%s", 
        bcolors.DarkGray, color, module, bcolors.DarkGray,
        bcolors.White, message,
        bcolors.DarkGray, bcolors.Green, bar, bcolors.DarkGray,
        bcolors.White, percent, bcolors.Endc)
}

func NotifyBoxed(module, message string, args ...interface{}) {
    moduleColors := map[string]string{
        NOTIFY_INFO:    bcolors.BrightBlue,
        NOTIFY_WARNING: bcolors.BrightYellow,
        NOTIFY_ERROR:   bcolors.BrightRed,
        NOTIFY_SUCCESS: bcolors.BrightGreen,
        NOTIFY_DEBUG:   bcolors.BrightCyan,
        NOTIFY_TOR:     bcolors.BrightMagenta,
        NOTIFY_DNS:     bcolors.BrightCyan,
        NOTIFY_FW:      bcolors.BrightRed,
        NOTIFY_PROXY:   bcolors.BrightBlue,
        NOTIFY_SQL:     bcolors.BrightYellow,
        NOTIFY_NET:     bcolors.BrightGreen,
        NOTIFY_CONFIG:  bcolors.BrightWhite,
    }
    
    color := moduleColors[module]
    if color == "" {
        color = bcolors.White
    }

    msgLen := len(message) + 6 + len(module) + 2
    if msgLen < 50 {
        msgLen = 50
    }
    border := strings.Repeat("─", msgLen)
    
    fmt.Printf("\n%s┌%s┐%s\n", bcolors.DarkGray, border, bcolors.Endc)
    fmt.Printf("%s│ %s[%s%s%s] %s%-*s %s│%s\n", bcolors.DarkGray, bcolors.DarkGray, color, module, bcolors.DarkGray, bcolors.White, msgLen-6-len(module), message, bcolors.DarkGray, bcolors.Endc)
    fmt.Printf("%s└%s┘%s\n", bcolors.DarkGray, border, bcolors.Endc)
}

func getActiveInterface() string {
    methods := []string{
        "ip route show default | awk '{print $5}'",
        "route -n | grep '^0.0.0.0' | awk '{print $8}'",
        "netstat -rn | grep '^0.0.0.0' | awk '{print $8}'",
        "ip -o link show | grep -v 'lo' | grep 'state UP' | awk -F': ' '{print $2}' | head -1",
    }

    for _, method := range methods {
        cmd := exec.Command("bash", "-c", method)
        output, err := cmd.Output()
        if err == nil {
            iface := strings.TrimSpace(string(output))
            if iface != "" && iface != "lo" {
                if err := exec.Command("ip", "link", "show", iface).Run(); err == nil {
                    return iface
                }
            }
        }
    }

    files, err := os.ReadDir("/sys/class/net")
    if err == nil {
        for _, file := range files {
            iface := file.Name()
            if iface != "lo" {
                cmd := exec.Command("ip", "addr", "show", iface)
                output, _ := cmd.Output()
                if strings.Contains(string(output), "inet ") {
                    return iface
                }
            }
        }
    }

    commonInterfaces := []string{"ens33", "ens3", "eth0", "wlan0", "enp0s3", "enp0s8"}
    for _, iface := range commonInterfaces {
        if err := exec.Command("ip", "link", "show", iface).Run(); err == nil {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Using fallback interface: %s", iface))
            return iface
        }
    }

    Notify(NOTIFY_WARNING, "No active interface found, defaulting to eth0")
    return "eth0"
}

func getTorUserID() string {
    torUsers := []string{"debian-tor", "tor", "_tor"}
    
    for _, user := range torUsers {
        cmd := exec.Command("id", "-u", user)
        output, err := cmd.Output()
        if err == nil {
            uid := strings.TrimSpace(string(output))
            if uid != "" {
                return uid
            }
        }
    }

    Notify(NOTIFY_WARNING, "Tor user not found! Defaulting to debian-tor")
    return "debian-tor"
}

func getDNSUserID() string {
    dnsUsers := []string{"dnsmasq", "nobody", "systemd-resolve"}
    
    for _, user := range dnsUsers {
        cmd := exec.Command("id", "-u", user)
        output, err := cmd.Output()
        if err == nil {
            uid := strings.TrimSpace(string(output))
            if uid != "" {
                return uid
            }
        }
    }

    Notify(NOTIFY_WARNING, "DNS user not found, using nobody (UID 65534)")
    return "65534"
}

func isTorRunning() bool {
    cmd := exec.Command("systemctl", "is-active", "tor@default.service")
    output, err := cmd.Output()
    if err != nil {
        cmd = exec.Command("systemctl", "is-active", "tor.service")
        output, err = cmd.Output()
        if err != nil {
            return false
        }
    }
    return strings.TrimSpace(string(output)) == "active"
}

func waitForService(service string, maxAttempts int) error {
    for i := 0; i < maxAttempts; i++ {
        cmd := exec.Command("systemctl", "is-active", service)
        output, err := cmd.Output()
        if err == nil && strings.TrimSpace(string(output)) == "active" {
            return nil
        }
        time.Sleep(2 * time.Second)
    }
    return fmt.Errorf("service %s not ready after %d attempts", service, maxAttempts)
}

func getSquidConfigPath() string {
    possiblePaths := []string{
        "/etc/squid/squid.conf",
        "/etc/squid3/squid.conf",
        "/usr/local/etc/squid/squid.conf",
        "/usr/local/squid/etc/squid.conf",
        "/opt/squid/etc/squid.conf",
    }

    for _, path := range possiblePaths {
        if _, err := os.Stat(path); err == nil {
            return path
        }
    }

    defaultPath := "/etc/squid/squid.conf"
    if err := os.MkdirAll("/etc/squid", 0755); err != nil {
        Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to create squid directory: %v", err))
        return defaultPath
    }
    return defaultPath
}

func checkRootPrivileges() error {
    if os.Geteuid() != 0 {
        return fmt.Errorf("this function requires root privileges. Please run with sudo or as root")
    }
    return nil
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
        Notify(NOTIFY_ERROR, "No MODULE was set. Use 'show modules' for details.")
        return
    }
    AnonimityFunctions(Function)
}

func AnonimityFunctions(functionName string, args ...interface{}) {
    rootFunctions := []string{"vanish", "2", "stop", "9", "restore", "8", "reload", "6", "exitnode", "7"}
    for _, fn := range rootFunctions {
        if functionName == fn {
            if err := checkRootPrivileges(); err != nil {
                Notify(NOTIFY_ERROR, err.Error())
                return
            }
            break
        }
    }

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            IFACE:       utils.IFace,
            DISTRO:      utils.Distro,
            OUTPUTLOGS:  utils.SecLogs,
            PROXIES:     utils.Proxies,
            FUNCTION:    functionName,
        }, true, false)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to set proxy: %v", err))
        }
    } else {
        menus.PrintSelected(menus.PrintOptions{
            IFACE:      utils.IFace,
            DISTRO:     utils.Distro,
            OUTPUTLOGS: utils.SecLogs,
            FUNCTION:   functionName,
        }, true, false)
    }

    activeInterface := getActiveInterface()
    Notify(NOTIFY_INFO, fmt.Sprintf("Using interface: %s", activeInterface))

    commands := map[string]func(){
        "setups": func() { 
            banners.GraphicsTorNet()
            setups.CheckMissing() 
        },
        "vanish": func() { 
            banners.GraphicsTorNet()

            if err := CreateConfigBackups(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to create backups: %v", err))
            }
            
            if err := configureTorNetwork(activeInterface); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure Tor network: %v", err))
            }
        },
        "status": func() { 
            banners.GraphicsTorNet()
            CheckServiceStatus()
            TorStatus(0) 
        },
        "torip": func() { 
            banners.GraphicsTorNet()
            CheckIP()
            CheckLeakTest()
            TorStatus(0) 
        },
        "exitnode": func() { 
            banners.GraphicsTorNet()
            TorCircuit() 
        },
        "restore": func() { 
            banners.GraphicsTorNet()
            if err := ResetToDefault(false, false); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to reset: %v", err))
            }
        },
        "chains": func() { 
            banners.GraphicsTorNet()
            subprocess.Run("tail -vf /var/log/privoxy/logfile") 
        },
        "reload": func() { 
            banners.GraphicsTorNet()
            if err := ConfigTorrc(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure torrc: %v", err))
            }
            if err := ConfigFirewall(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure firewall: %v", err))
            }
            CheckIP() 
        },
        "stop": func() { 
            banners.GraphicsTorNet()
            if err := KillTor(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to stop Tor: %v", err))
            }
            CheckServiceStatus()
            TorStatus(0)
        },
        "1": func() { 
            banners.GraphicsTorNet()
            setups.CheckMissing() 
        },
        "2": func() { 
            banners.GraphicsTorNet()
            if err := configureTorNetwork(activeInterface); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure Tor network: %v", err))
            }
        },
        "3": func() { 
            banners.GraphicsTorNet()
            CheckServiceStatus()
            TorStatus(0) 
        },
        "4": func() { 
            banners.GraphicsTorNet()
            CheckIP()
            CheckLeakTest()
            TorStatus(0) 
        },
        "5": func() { 
            banners.GraphicsTorNet()
            TorCircuit() 
        },
        "6": func() { 
            banners.GraphicsTorNet()
            if err := ResetToDefault(false, false); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to reset: %v", err))
            }
        },
        "7": func() { 
            banners.GraphicsTorNet()
            subprocess.Run("tail -vf /var/log/privoxy/logfile") 
        },
        "8": func() { 
            banners.GraphicsTorNet()
            if err := ConfigTorrc(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure torrc: %v", err))
            }
            if err := ConfigFirewall(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to configure firewall: %v", err))
            }
            CheckIP() 
        },
        "9": func() { 
            banners.GraphicsTorNet()
            if err := KillTor(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to stop Tor: %v", err))
            }
            CheckServiceStatus()
            TorStatus(0)
        },
    }

    textCommands := []string{"setups", "vanish", "status", "exitnode", "torip", "restore", "chains", "reload", "stop"}

    if action, exists := commands[functionName]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(functionName); err == nil {
        Notify(NOTIFY_ERROR, fmt.Sprintf("Number %d is invalid. Valid numbers are 1-10.", num))
        menus.ListTorsocksFunctions(Function)
        return
    }

    lowerInput := strings.ToLower(functionName)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Function '%s' is invalid. Did you mean '%s'?", functionName, cmd))
            return
        }
    }

    Notify(NOTIFY_ERROR, fmt.Sprintf("Module '%s' is invalid.", functionName))
    menus.ListTorsocksFunctions(Function)
}

func configureTorNetwork(interfaceName string) error {
    Notify(NOTIFY_TOR, fmt.Sprintf("Configuring Tor network on interface: %s", interfaceName))

    if err := ConfigureResolv(); err != nil {
        return fmt.Errorf("failed to configure resolv.conf: %v", err)
    }

    if err := ConfigChangeMac(interfaceName); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("MAC changer issue (non-critical): %v", err))
        if err := ChangeMacDirect(interfaceName); err != nil {
            Notify(NOTIFY_WARNING, "MAC randomization disabled (non-critical)")
        }
    }

    if err := ConfigDhclient(); err != nil {
        return fmt.Errorf("failed to configure dhclient: %v", err)
    }

    if err := ConfigDnsmasq(); err != nil {
        return fmt.Errorf("failed to configure dnsmasq: %v", err)
    }

    if err := ConfigSquid(); err != nil {
        return fmt.Errorf("failed to configure squid: %v", err)
    }

    if err := ConfigPrivoxy(); err != nil {
        return fmt.Errorf("failed to configure privoxy: %v", err)
    }

    if err := ConfigTorrc(); err != nil {
        return fmt.Errorf("failed to configure torrc: %v", err)
    }

    if err := StartTorServices(); err != nil {
        return fmt.Errorf("failed to start Tor services: %v", err)
    }

    if err := ConfigFirewall(); err != nil {
        return fmt.Errorf("failed to configure firewall: %v", err)
    }

    if err := CheckServiceStatus(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Service status check warning: %v", err))
    }

    CheckIP()
    CheckLeakTest()
    TorStatus(0)

    return nil
}

func ChangeMacDirect(interfaceName string) error {
    if interfaceName == "" {
        interfaceName = getActiveInterface()
    }

    if err := utils.LocateTool("macchanger"); err != nil {
        return fmt.Errorf("macchanger not found: %v", err)
    }

    commands := []string{
        fmt.Sprintf("sudo ip link set %s down", interfaceName),
        fmt.Sprintf("sudo macchanger -r %s", interfaceName),
        fmt.Sprintf("sudo ip link set %s up", interfaceName),
    }
    
    for _, cmdStr := range commands {
        cmd := exec.Command("bash", "-c", cmdStr)
        output, err := cmd.CombinedOutput()
        if err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("MAC changer command failed: %s", cmdStr))
            return fmt.Errorf("failed to change MAC: %v", err)
        }
        Notify(NOTIFY_INFO, string(output))
        time.Sleep(200 * time.Millisecond)
    }
    
    Notify(NOTIFY_SUCCESS, fmt.Sprintf("MAC address changed for interface %s", interfaceName))
    return nil
}

func ConfigChangeMac(interfaceName string) error {
    if err := utils.LocateTool("macchanger"); err != nil {
        Notify(NOTIFY_WARNING, "macchanger not found, MAC address randomization disabled")
        return nil
    }

    if interfaceName == "" {
        interfaceName = getActiveInterface()
    }

    cmd := exec.Command("ip", "link", "show", interfaceName)
    if err := cmd.Run(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Interface %s not found, using fallback", interfaceName))
        interfaceName = getActiveInterface()
        if interfaceName == "eth0" {
            if err := exec.Command("ip", "link", "show", "eth0").Run(); err != nil {
                Notify(NOTIFY_WARNING, "No valid interface found, MAC changer will be skipped")
                return nil
            }
        }
    }

    content := fmt.Sprintf(`
# Generated by africana-framework. Delete at your own risk!
[Unit]
Description=changes mac for %%I
Wants=network.target
Before=network.target
BindsTo=sys-subsystem-net-devices-%%i.device
After=sys-subsystem-net-devices-%%i.device
[Service]
Type=oneshot
ExecStart=/usr/bin/macchanger -r %%I
RemainAfterExit=yes
[Install]
WantedBy=multi-user.target
`)
    
    filePath := "/etc/systemd/system/changemac@.service"
    
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to create macchanger service: %v", err))
            return nil
        }
        Notify(NOTIFY_SUCCESS, fmt.Sprintf("Created macchanger systemd service: %s", filePath))

        exec.Command("sudo", "systemctl", "daemon-reload").Run()
    }

    specificService := fmt.Sprintf("/etc/systemd/system/changemac@%s.service", interfaceName)
    if _, err := os.Stat(specificService); os.IsNotExist(err) {
        if err := exec.Command("ln", "-sf", filePath, specificService).Run(); err != nil {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Could not create service link: %v", err))
        }
    }
    
    return nil
}

func ConfigDhclient() error {
    if err := utils.LocateTool("dhclient"); err != nil {
        Notify(NOTIFY_WARNING, "dhclient not found, skipping configuration")
        return nil
    }

    filePath := "/etc/dhcp/dhclient.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if _, err := os.Stat("/etc/dhcp/dhclient.conf"); err == nil {
            subprocess.Run("cp -r /etc/dhcp/dhclient.conf /etc/dhcp/dhclient.conf.bak_africana")
            
            if err := os.MkdirAll("/etc/dhcp", 0755); err != nil {
                return fmt.Errorf("failed to create dhcp directory: %v", err)
            }
            
            filesToReplacements := map[string]map[string]string{
                "/etc/dhcp/dhclient.conf": {
                    "#prepend domain-name-servers 127.0.0.1;": "prepend domain-name-servers 127.0.0.1, 1.1.1.1, 1.0.0.1, 8.8.8.8, 8.8.4.4;",
                },
            }
            utils.Editors(filesToReplacements)
            Notify(NOTIFY_SUCCESS, "Dhclient configured")
        } else {
            content := `# Generated by africana-framework
prepend domain-name-servers 127.0.0.1, 1.1.1.1, 8.8.8.8;
`
            if err := os.MkdirAll("/etc/dhcp", 0755); err != nil {
                return fmt.Errorf("failed to create dhcp directory: %v", err)
            }
            if err := os.WriteFile("/etc/dhcp/dhclient.conf", []byte(content), 0644); err != nil {
                return fmt.Errorf("failed to write dhclient.conf: %v", err)
            }
        }
    }
    return nil
}

func ConfigDnsmasq() error {
    if err := utils.LocateTool("dnsmasq"); err != nil {
        Notify(NOTIFY_WARNING, "dnsmasq not found, skipping configuration")
        return nil
    }

    filePath := "/etc/dnsmasq.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if _, err := os.Stat("/etc/dnsmasq.conf"); err == nil {
            subprocess.Run("cp -r /etc/dnsmasq.conf /etc/dnsmasq.conf.bak_africana")
        }
        
        content := `# Generated by africana-framework. Delete at your own risk!
port=5353
listen-address=127.0.0.1
bind-interfaces
no-resolv
no-poll
server=127.0.0.1#9053
cache-size=1000
`
        if err := os.WriteFile("/etc/dnsmasq.conf", []byte(content), 0644); err != nil {
            return fmt.Errorf("failed to write dnsmasq.conf: %v", err)
        }
        Notify(NOTIFY_DNS, "Dnsmasq configured to use Tor for DNS")
    }
    return nil
}

func ConfigPrivoxy() error {
    if err := utils.LocateTool("privoxy"); err != nil {
        Notify(NOTIFY_WARNING, "privoxy not found, skipping configuration")
        return nil
    }

    possiblePaths := []string{
        "/etc/privoxy/config",
        "/usr/local/etc/privoxy/config",
        "/opt/privoxy/etc/config",
    }
    
    configPath := ""
    for _, path := range possiblePaths {
        if _, err := os.Stat(path); err == nil {
            configPath = path
            break
        }
    }
    
    if configPath == "" {
        if err := os.MkdirAll("/etc/privoxy", 0755); err != nil {
            return fmt.Errorf("failed to create privoxy directory: %v", err)
        }
        configPath = "/etc/privoxy/config"
    }

    filePath := configPath + ".bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if _, err := os.Stat(configPath); err == nil {
            subprocess.Run("cp -r " + configPath + " " + filePath)
        }
        
        content := `# Generated by africana-framework. Delete at your own risk!
listen-address 127.0.0.1:8118
forward-socks5t / 127.0.0.1:9050 .
debug 1
debug 2
debug 1024
debug 4096
`
        if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
            return fmt.Errorf("failed to write privoxy config: %v", err)
        }
        Notify(NOTIFY_PROXY, "Privoxy configured to use Tor (127.0.0.1:9050)")
    }
    return nil
}

func ConfigSquid() error {
    if err := utils.LocateTool("squid"); err != nil {
        Notify(NOTIFY_WARNING, "squid not found, skipping configuration")
        return nil
    }

    configPath := getSquidConfigPath()
    filePath := configPath + ".bak_africana"
    
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if _, err := os.Stat(configPath); err == nil {
            subprocess.Run("cp -r " + configPath + " " + filePath)
        }
        
        content := `http_port 3129
cache_peer 127.0.0.1 parent 8118 0 no-query default
never_direct allow all

# ACLs
acl localnet src 127.0.0.0/8
acl localnet src 10.0.0.0/8
acl localnet src 172.16.0.0/12
acl localnet src 192.168.0.0/16
acl SSL_ports port 443
acl Safe_ports port 80
acl Safe_ports port 443
acl Safe_ports port 8080
acl CONNECT method CONNECT

# Allow local networks
http_access allow localnet
http_access allow localhost

# Deny rest
http_access deny all

# Settings
shutdown_lifetime 0 seconds
forwarded_for off
request_header_access X-Forwarded-For deny all
`
        if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
            return fmt.Errorf("failed to create squid directory: %v", err)
        }
        
        if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
            return fmt.Errorf("failed to write squid config: %v", err)
        }
        Notify(NOTIFY_PROXY, "Squid configured to use Privoxy (127.0.0.1:8118)")
    }
    return nil
}

func ConfigTorrc() error {
    if err := utils.LocateTool("tor"); err != nil {
        return fmt.Errorf("tor not found: %v", err)
    }

    possiblePaths := []string{
        "/etc/tor/torrc",
        "/usr/local/etc/tor/torrc",
        "/opt/tor/etc/torrc",
    }
    
    torrcPath := ""
    for _, path := range possiblePaths {
        if _, err := os.Stat(path); err == nil {
            torrcPath = path
            break
        }
    }
    
    if torrcPath == "" {
        if err := os.MkdirAll("/etc/tor", 0755); err != nil {
            return fmt.Errorf("failed to create tor directory: %v", err)
        }
        torrcPath = "/etc/tor/torrc"
    }

    Notify(NOTIFY_TOR, fmt.Sprintf("Configuring torrc at %s", torrcPath))

    if _, err := os.Stat(torrcPath); err == nil {
        if err := exec.Command("cp", torrcPath, torrcPath+".bak_torsocks").Run(); err != nil {
            return fmt.Errorf("failed to create backup: %v", err)
        }
    }

    content := `# Generated by africana-framework. Delete at your own risk!
VirtualAddrNetwork 10.192.0.0/10
AutomapHostsOnResolve 1
CookieAuthentication 1
TransPort 9040
SocksPort 9050
DNSPort 5353
`

    if err := os.WriteFile(torrcPath, []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to write torrc: %v", err)
    }

    Notify(NOTIFY_TOR, "Torrc configured successfully")
    return nil
}

func ConfigureResolv() error {
    resolvContent := `# Generated by africana-framework - DNS forced through Tor
nameserver 127.0.0.1    # Local DNS (dnsmasq/Tor)
nameserver 127.0.0.53   # systemd-resolved (fallback)
options rotate
options timeout:1
options attempts:2
`
    content, err := os.ReadFile("/etc/resolv.conf")
    if err != nil {
        return CreateResolvConf()
    }

    if strings.Contains(string(content), "DNS forced through Tor") {
        Notify(NOTIFY_DNS, "Resolv.conf already configured for Tor DNS")
        return nil
    }

    Notify(NOTIFY_DNS, "Configuring resolv.conf to force DNS through Tor")
    if err := exec.Command("cp", "/etc/resolv.conf", "/etc/resolv.conf.backup_torsocks").Run(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to create backup: %v", err))
    }

    if err := os.WriteFile("/etc/resolv.conf", []byte(resolvContent), 0644); err != nil {
        return fmt.Errorf("failed to write resolv.conf: %v", err)
    }

    if _, err := os.Stat("/run/systemd/resolve/resolv.conf"); err == nil {
        exec.Command("sudo", "systemctl", "restart", "systemd-resolved").Run()
        Notify(NOTIFY_DNS, "Restarted systemd-resolved")
    }

    Notify(NOTIFY_SUCCESS, "Resolv.conf configured to use Tor DNS")
    return nil
}

func CreateResolvConf() error {
    Notify(NOTIFY_CONFIG, "Missing resolv.conf. Creating default...")

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
    if err := os.MkdirAll("/etc", 0755); err != nil {
        return fmt.Errorf("failed to create /etc directory: %v", err)
    }

    if err := os.WriteFile("/etc/resolv.conf", []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to create resolv.conf: %v", err)
    }

    Notify(NOTIFY_SUCCESS, "resolv.conf created successfully")
    return nil
}

func StartTorServices() error {
    Notify(NOTIFY_TOR, "Starting Tor services...")
    
    activeInterface := getActiveInterface()
    Notify(NOTIFY_INFO, fmt.Sprintf("Active interface detected: %s", activeInterface))

    macchangerInstalled := true
    if err := utils.LocateTool("macchanger"); err != nil {
        macchangerInstalled = false
        Notify(NOTIFY_WARNING, "macchanger not found, skipping MAC address randomization")
    }
    
    services := []string{
        "dnsmasq.service",
        "tor@default.service",
        "privoxy.service",
        "squid.service",
    }

    if macchangerInstalled {
        macService := fmt.Sprintf("changemac@%s.service", activeInterface)
        if _, err := os.Stat("/etc/systemd/system/changemac@.service"); err == nil {
            services = append(services, macService)
        } else {
            Notify(NOTIFY_WARNING, "macchanger service not configured, skipping")
        }
    }

    for _, service := range services {
        if err := exec.Command("sudo", "systemctl", "stop", service).Run(); err != nil {
            if !strings.Contains(err.Error(), "not-found") {
                Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to stop %s: %v", service, err))
            }
        }
        time.Sleep(100 * time.Millisecond)
    }
    for _, service := range services {
        cmd := exec.Command("systemctl", "list-unit-files", service)
        output, _ := cmd.Output()
        if !strings.Contains(string(output), service) {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Service %s not found, skipping", service))
            continue
        }

        cmd = exec.Command("sudo", "systemctl", "restart", service)
        if err := cmd.Run(); err != nil {
            if strings.Contains(service, "changemac") {
                Notify(NOTIFY_WARNING, fmt.Sprintf("MAC changer service failed (non-critical): %v", err))
                continue
            }
            return fmt.Errorf("failed to start %s: %v", service, err)
        }

        if !strings.Contains(service, "changemac") {
            if err := waitForService(service, 10); err != nil {
                return fmt.Errorf("%s%s[!] %sService %s failed to start ...\n", 
                    bcolors.Bold, bcolors.Red, bcolors.Endc, service)
            }
        }
        time.Sleep(500 * time.Millisecond)
    }
    
    Notify(NOTIFY_TOR, "Waiting for Tor to fully initialize (10 seconds)...")
    time.Sleep(10 * time.Second)
    
    if !isTorRunning() {
        return fmt.Errorf("Tor is not running properly after startup")
    }
    
    Notify(NOTIFY_SUCCESS, "Tor services Up and running...")
    return nil
}

func CheckServiceStatus() error {
    activeInterface := getActiveInterface()
    
    services := []string{
        "dnsmasq.service",
        "squid.service",
        "privoxy.service",
        "tor@default.service",
        fmt.Sprintf("changemac@%s.service", activeInterface),
    }

    Notify(NOTIFY_INFO, "Checking service status...")

    var failedServices []string
    for _, service := range services {
        cmd := exec.Command("systemctl", "is-active", service)
        output, _ := cmd.Output()
        status := strings.TrimSpace(string(output))

        if status == "active" {
            NotifyStatus(NOTIFY_INFO, fmt.Sprintf("%s is running ...", service), true)
        } else {
            NotifyStatus(NOTIFY_ERROR, fmt.Sprintf("%s is not running ...", service), false)
            failedServices = append(failedServices, service)
        }
    }

    if len(failedServices) > 0 {
        return fmt.Errorf("Following services are not running: %s ...", strings.Join(failedServices, ", "))
    }
    return nil
}

func ConfigFirewall() error {
    Notify(NOTIFY_FW, "Backing up Iptables")

    FirewallOff, FirewallOn := CheckDefaultRules()
    if FirewallOn != "" {
        return fmt.Errorf("%s\nCan't execute %siptables-save%s. See the reason below.\n%s%s%s", 
            bcolors.BrightRed, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightRed, FirewallOn, bcolors.Endc)
    }

    if strings.TrimSpace(FirewallOff) == "0" {
        Notify(NOTIFY_FW, "Default rules are configured, skipping backup")
    } else {
        Notify(NOTIFY_FW, "Saving default firewall rules...")

        cmd := exec.Command("bash", "-c", "sudo iptables-save > /etc/iptables_rules_torsocks.bak")
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to save iptables rules: %v%s", 
                bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err, bcolors.Endc)
        }

        Notify(NOTIFY_SUCCESS, "Done saving firewall rules")
    }

    Notify(NOTIFY_FW, "Setting up firewall rules...")

    Notify(NOTIFY_SUCCESS, "Done setting up firewall rules")
    return nil
}

func SetTorIptablesRules() error {
    torUID := getTorUserID()
    dnsmasqUID := getDNSUserID()
    
    transPort := "9040"
    dnsPort := "5353"
    virtAddr := "10.192.0.0/10"

    Notify(NOTIFY_FW, "Setting up Tor iptables rules...")

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

        fmt.Sprintf("iptables -t nat -A OUTPUT -m owner --uid-owner %s -j RETURN", torUID),
        fmt.Sprintf("iptables -A OUTPUT -m owner --uid-owner %s -j ACCEPT", torUID),

        fmt.Sprintf("iptables -A OUTPUT -p udp --dport 53 -m owner --uid-owner %s -j ACCEPT", dnsmasqUID),
        fmt.Sprintf("iptables -A OUTPUT -p tcp --dport 53 -m owner --uid-owner %s -j ACCEPT", dnsmasqUID),

        "iptables -A INPUT -i lo -j ACCEPT",
        "iptables -A OUTPUT -o lo -j ACCEPT",
        "iptables -A FORWARD -i lo -j ACCEPT",
        "iptables -A FORWARD -o lo -j ACCEPT",

        "iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT",
        "iptables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT",
        "iptables -A FORWARD -m state --state ESTABLISHED,RELATED -j ACCEPT",

        fmt.Sprintf("iptables -t nat -A OUTPUT -p tcp --dport %s -j RETURN", transPort),
        "iptables -t nat -A OUTPUT -p tcp --dport 9050 -j RETURN",
        "iptables -t nat -A OUTPUT -p tcp --dport 9051 -j RETURN",
        "iptables -t nat -A OUTPUT -p tcp --dport 5353 -j RETURN",
        "iptables -t nat -A OUTPUT -p udp --dport 5353 -j RETURN",

        fmt.Sprintf("iptables -A OUTPUT -p tcp -m owner --uid-owner %s -j ACCEPT", torUID),

        fmt.Sprintf("iptables -A OUTPUT -p udp --dport 53 -m owner --uid-owner %s -j ACCEPT", torUID),
        fmt.Sprintf("iptables -A OUTPUT -p tcp --dport 53 -m owner --uid-owner %s -j ACCEPT", torUID),

        "iptables -t nat -A OUTPUT -d 127.0.0.0/8 -j RETURN",
        "iptables -t nat -A OUTPUT -d 10.0.0.0/8 -j RETURN",
        "iptables -t nat -A OUTPUT -d 172.16.0.0/12 -j RETURN",
        "iptables -t nat -A OUTPUT -d 192.168.0.0/16 -j RETURN",
        "iptables -t nat -A OUTPUT -d 169.254.0.0/16 -j RETURN",
        "iptables -t nat -A OUTPUT -d " + virtAddr + " -j RETURN",
        "iptables -t nat -A OUTPUT -d 224.0.0.0/4 -j RETURN",
        "iptables -t nat -A OUTPUT -d 240.0.0.0/4 -j RETURN",

        "iptables -t nat -A OUTPUT -p udp --dport 53 -j REDIRECT --to-ports " + dnsPort,
        "iptables -t nat -A OUTPUT -p tcp --dport 53 -j REDIRECT --to-ports " + dnsPort,

        "iptables -t nat -A OUTPUT -p tcp --syn -j REDIRECT --to-ports " + transPort,

        "iptables -A OUTPUT -p udp -j ACCEPT",
        "iptables -A INPUT -p udp -j ACCEPT",

        "iptables -A OUTPUT -p icmp -j ACCEPT",
        "iptables -A INPUT -p icmp -j ACCEPT",

        "iptables -A OUTPUT -m state --state INVALID -j DROP",
        "iptables -A FORWARD -m state --state INVALID -j DROP",

        "iptables -A INPUT -p udp --sport 53 -j ACCEPT",
        "iptables -A INPUT -p tcp --sport 53 -j ACCEPT",

        "iptables -A OUTPUT -p tcp --dport 443 -j ACCEPT",
        "iptables -A OUTPUT -p tcp --dport 80 -j ACCEPT",
        "iptables -A OUTPUT -p tcp --dport 9030 -j ACCEPT",
        "iptables -A OUTPUT -p tcp --dport 9001 -j ACCEPT",
        "iptables -A OUTPUT -p tcp --dport 9040 -j ACCEPT",

        "iptables -P INPUT DROP",
        "iptables -P FORWARD DROP",
        "iptables -P OUTPUT DROP",
    }

    if _, err := os.Stat("/proc/net/if_inet6"); err == nil {
        Notify(NOTIFY_FW, "IPv6 detected, adding IPv6 rules")
        
        ipv6Commands := []string{
            "ip6tables -P INPUT ACCEPT",
            "ip6tables -P FORWARD ACCEPT",
            "ip6tables -P OUTPUT ACCEPT",
            "ip6tables -F",
            "ip6tables -X",
            "ip6tables -t nat -F",
            "ip6tables -t nat -X",

            fmt.Sprintf("ip6tables -A OUTPUT -m owner --uid-owner %s -j ACCEPT", torUID),

            fmt.Sprintf("ip6tables -A OUTPUT -p udp --dport 53 -m owner --uid-owner %s -j ACCEPT", dnsmasqUID),
            fmt.Sprintf("ip6tables -A OUTPUT -p tcp --dport 53 -m owner --uid-owner %s -j ACCEPT", dnsmasqUID),

            "ip6tables -A INPUT -i lo -j ACCEPT",
            "ip6tables -A OUTPUT -o lo -j ACCEPT",

            "ip6tables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT",
            "ip6tables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT",

            fmt.Sprintf("ip6tables -t nat -A OUTPUT -p tcp --dport %s -j RETURN", transPort),
            "ip6tables -t nat -A OUTPUT -p tcp --dport 9050 -j RETURN",
            "ip6tables -t nat -A OUTPUT -p tcp --dport 9051 -j RETURN",
            "ip6tables -t nat -A OUTPUT -p udp --dport 5353 -j RETURN",

            "ip6tables -t nat -A OUTPUT -d ::1/128 -j RETURN",
            "ip6tables -t nat -A OUTPUT -d fe80::/64 -j RETURN",
            "ip6tables -t nat -A OUTPUT -d fc00::/7 -j RETURN",
            "ip6tables -t nat -A OUTPUT -d ff00::/8 -j RETURN",

            "ip6tables -t nat -A OUTPUT -p udp --dport 53 -j REDIRECT --to-ports " + dnsPort,
            "ip6tables -t nat -A OUTPUT -p tcp --dport 53 -j REDIRECT --to-ports " + dnsPort,

            "ip6tables -t nat -A OUTPUT -p tcp --syn -j REDIRECT --to-ports " + transPort,

            "ip6tables -A OUTPUT -p udp -j ACCEPT",
            "ip6tables -A INPUT -p udp -j ACCEPT",

            "ip6tables -A OUTPUT -p icmpv6 -j ACCEPT",
            "ip6tables -A INPUT -p icmpv6 -j ACCEPT",

            "ip6tables -A INPUT -m state --state INVALID -j DROP",
            "ip6tables -A OUTPUT -m state --state INVALID -j DROP",

            "ip6tables -A INPUT -p udp --sport 53 -j ACCEPT",
            "ip6tables -A INPUT -p tcp --sport 53 -j ACCEPT",

            "ip6tables -P INPUT DROP",
            "ip6tables -P FORWARD DROP",
            "ip6tables -P OUTPUT DROP",
        }
        commands = append(commands, ipv6Commands...)
    }

    var lastError error
    successCount := 0
    totalCommands := len(commands)

    for i, cmdStr := range commands {
        if strings.TrimSpace(cmdStr) == "" {
            continue
        }
        
        cmd := exec.Command("bash", "-c", "sudo " + cmdStr)
        output, err := cmd.CombinedOutput()
        if err != nil {
            if !strings.Contains(string(output), "No chain/target/match") && 
               !strings.Contains(string(output), "Bad rule") &&
               !strings.Contains(string(output), "No such file") &&
               !strings.Contains(string(output), "Protocol not supported") &&
               !strings.Contains(string(output), "No such process") {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed command %d/%d: %s", i+1, totalCommands, cmdStr))
                if string(output) != "" {
                    fmt.Printf("%sOutput: %s%s\n", bcolors.Yellow, string(output), bcolors.Endc)
                }
                lastError = err
            }
        } else {
            successCount++
        }
        
        if i%10 == 0 {
            NotifyProgress(NOTIFY_FW, "Configuring iptables rules", i, totalCommands)
        }
        time.Sleep(30 * time.Millisecond)
    }

    fmt.Printf("\n")
    Notify(NOTIFY_INFO, fmt.Sprintf("Completed: %d/%d commands executed successfully", successCount, totalCommands))

    if lastError != nil {
        return fmt.Errorf("some iptables commands failed during Tor setup")
    }

    successMsg := `
    ALL TRAFFIC IS NOW FORCED THROUGH TOR!
    TCP  ✅ Redirected through Tor
    UDP  ✅ Allowed (Tor doesn't support UDP)
    DNS  ✅ Redirected through Tor
    ICMP ✅ Allowed for testing
    IPv6 ✅ Protected against leaks`

    fmt.Printf("%s%s%s\n", bcolors.Green, successMsg, bcolors.Endc)

    return nil
}

func verifyIptablesRules() error {
    cmd := exec.Command("iptables", "-t", "nat", "-L", "OUTPUT", "-n")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to verify NAT rules: %v", err)
    }
    
    if !strings.Contains(string(output), "REDIRECT") {
        return fmt.Errorf("REDIRECT rules not found in iptables")
    }

    if !strings.Contains(string(output), "9040") {
        return fmt.Errorf("Tor TransPort (9040) not found in rules")
    }

    if !strings.Contains(string(output), "5353") {
        return fmt.Errorf("DNS redirect (5353) not found in rules")
    }
    
    return nil
}

func getNonLoopbackInterfaces() []string {
    var interfaces []string
    
    files, err := os.ReadDir("/sys/class/net")
    if err != nil {
        return []string{"eth0", "wlan0"}
    }
    
    for _, file := range files {
        iface := file.Name()
        if iface != "lo" {
            interfaces = append(interfaces, iface)
        }
    }
    
    return interfaces
}

func saveFirewallRules() error {
    backupFile := "/etc/iptables_rules_original_" + time.Now().Format("2006-01-02_15-04-05") + ".bak"
    cmd := exec.Command("bash", "-c", "sudo iptables-save > "+backupFile)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to save iptables rules: %v", err)
    }
    Notify(NOTIFY_SUCCESS, fmt.Sprintf("Firewall rules saved to: %s", backupFile))
    return nil
}

func RestoreIptablesDefault() error {
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
    }

    Notify(NOTIFY_FW, "Restoring default iptables rules...")
    
    var lastError error
    for _, cmdStr := range commands {
        cmd := exec.Command("bash", "-c", "sudo " + cmdStr)
        if err := cmd.Run(); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed: %s", cmdStr))
            lastError = err
        }
        time.Sleep(50 * time.Millisecond)
    }
    
    if lastError != nil {
        return fmt.Errorf("failed to restore default rules")
    }
    
    Notify(NOTIFY_SUCCESS, "Default iptables rules restored")
    return nil
}

func CheckDefaultRules() (string, string) {
    cmd := exec.Command("bash", "-c", "iptables-save | grep '^\\-' | wc -l")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", string(output)
    }
    return string(output), ""
}

func ResetToDefault(overridePass bool, resetAsChildFunc bool) error {
    if !overridePass {
        if err := saveFirewallRules(); err != nil {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Could not save rules: %v", err))
        }
        
        fmt.Printf("%s\nThis will overwrite all of your existing rules %sY(do it)%s/%sN(exit)%s: ", 
            bcolors.Colors(), bcolors.Green, bcolors.Endc, bcolors.BrightRed, bcolors.Endc)
        
        var resetConsent string
        fmt.Scanln(&resetConsent)

        if strings.ToLower(resetConsent) == "n" {
            fmt.Printf("%sCopy that..\n%s", bcolors.BrightRed, bcolors.Endc)
            return nil
        } else if strings.ToLower(resetConsent) != "y" {
            return fmt.Errorf("invalid input, aborting reset")
        }
    }

    if err := RestoreIptablesDefault(); err != nil {
        Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to restore iptables: %v", err))
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

        for _, cmdStr := range commands {
            cmd := exec.Command("bash", "-c", "sudo " + cmdStr)
            if err := cmd.Run(); err != nil {
                Notify(NOTIFY_ERROR, fmt.Sprintf("Failed: %s", cmdStr))
            }
            time.Sleep(50 * time.Millisecond)
        }
    }

    if !resetAsChildFunc {
        Notify(NOTIFY_SUCCESS, "Successfully reset Iptables to default")
    }

    return nil
}

func StopServices() error {
    activeInterface := getActiveInterface()
    Notify(NOTIFY_INFO, fmt.Sprintf("Killing: macchanger, dnsmasq, privoxy, squid, tor"))
    
    services := []string{
        fmt.Sprintf("changemac@%s.service", activeInterface),
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
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to stop %s: %v", service, err))
            failedServices = append(failedServices, service)
        } else {
            stoppedServices = append(stoppedServices, service)
        }
        time.Sleep(100 * time.Millisecond)
    }

    if len(stoppedServices) > 0 {
        Notify(NOTIFY_SUCCESS, fmt.Sprintf("Stopped: %s", strings.Join(stoppedServices, ", ")))
    }

    if len(failedServices) > 0 {
        Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to stop: %s", strings.Join(failedServices, ", ")))
    }

    if len(failedServices) > 0 {
        return fmt.Errorf("failed to stop %d services", len(failedServices))
    }
    return nil
}

func CheckIP() {
    Notify(NOTIFY_NET, "Getting your public IP address...")
    Notify(NOTIFY_NET, "Getting your system geolocation...")
    
    ip := getPublicIP()
    if ip == "" {
        Notify(NOTIFY_ERROR, "Could not retrieve public IP address.")
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
        Notify(NOTIFY_ERROR, "Could not retrieve geolocation data")
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

func CheckIptables() string {
    cmd := exec.Command("iptables", "-t", "nat", "-L", "-n")
    output, err := cmd.CombinedOutput()
    if err != nil {
        Notify(NOTIFY_ERROR, "Error while checking the iptables details...")
    }
    return string(output)
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

func CheckLeakTest() (*DNSLeakTestResult, error) {
    Notify(NOTIFY_DNS, "Testing DNS leaks...")

    if err := FixDNSLeaks(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Auto-fix failed: %v", err))
    }

    Notify(NOTIFY_DNS, "Please be patient. Getting results...")

    result, err := TestDNSLeak()
    if err != nil {
        return nil, fmt.Errorf("%s%s[!] %sDNS leak test failed: %v", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
    }

    fmt.Printf("%s\n", result)

    cleanJSON := stripColors(result)
    var leakResult DNSLeakTestResult
    if err := json.Unmarshal([]byte(cleanJSON), &leakResult); err != nil {
        return nil, fmt.Errorf("%s%s[!] %sFailed to parse DNS leak results: %v", 
            bcolors.Bold, bcolors.Red, bcolors.Endc, err)
    }

    if leakResult.LeakDetected {
        Notify(NOTIFY_ERROR, fmt.Sprintf("DNS LEAK DETECTED! Found %d non-Tor DNS servers!", len(leakResult.DNSServers)))

        for _, server := range leakResult.DNSServers {
            if !server.ViaTor {
                Notify(NOTIFY_WARNING, fmt.Sprintf("Leaked DNS: %s (ISP: %s, Country: %s)", 
                    server.IP, server.ISP, server.Country))
            }
        }
        
        Notify(NOTIFY_INFO, "Attempting to fix DNS leaks...")
        if err := FixDNSLeaks(); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to fix DNS leaks: %v", err))
        } else {
            Notify(NOTIFY_SUCCESS, "DNS leaks fixed! Please run the test again.")
        }
        
        return &leakResult, fmt.Errorf("%s%s[!] %sDNS leak detected", bcolors.Bold, bcolors.Red, bcolors.Endc)
    } else {
        Notify(NOTIFY_SUCCESS, "No DNS leaks detected. All DNS queries are going through Tor.")
        return &leakResult, nil
    }
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

    testMethods := []struct {
        name string
        dns  string
    }{
        {"System DNS", ""},
        {"Google DNS", "8.8.8.8"},
        {"Cloudflare DNS", "1.1.1.1"},
        {"OpenDNS", "208.67.222.222"},
    }

    testDomains := []string{
        "google.com",
        "facebook.com",
        "amazon.com",
        "microsoft.com",
    }

    systemDNS, err := getSystemDNSServers()
    if err == nil {
        for _, dnsIP := range systemDNS {
            isTor := isDNSThroughTor(dnsIP)
            server := DNSServer{
                IP:     dnsIP,
                ViaTor: isTor,
            }
            ipInfo, err := getIPInfo(dnsIP)
            if err == nil {
                server.ISP = ipInfo.ISP
                server.Country = ipInfo.Country
                server.ASN = ipInfo.ASN
            }
            
            servers = append(servers, server)
            if !isTor {
                leakDetected = true
                Notify(NOTIFY_WARNING, fmt.Sprintf("DNS leak detected: %s (ISP: %s)", dnsIP, server.ISP))
            }
        }
    }

    for _, method := range testMethods {
        if method.dns == "" {
            continue
        }
        
        for _, domain := range testDomains {
            server, err := testDNSResolution(domain, method.dns)
            if err == nil && server.IP != "" {
                isTor := isTorExitNode(server.IP)
                server.ViaTor = isTor
                
                if !isTor {
                    leakDetected = true
                    Notify(NOTIFY_WARNING, fmt.Sprintf("DNS leak detected via %s: %s", method.name, server.IP))
                }
                exists := false
                for _, s := range servers {
                    if s.IP == server.IP {
                        exists = true
                        break
                    }
                }
                if !exists {
                    servers = append(servers, server)
                }
            }
            time.Sleep(100 * time.Millisecond)
        }
    }

    return servers, leakDetected
}

func isDNSThroughTor(dnsIP string) bool {
    if dnsIP == "127.0.0.1" || dnsIP == "::1" {
        return true
    }
    if isTorExitNode(dnsIP) {
        return true
    }
    cmd := exec.Command("timeout", "5", "dig", "+short", "google.com", "@"+dnsIP, "-p", "5353")
    output, err := cmd.Output()
    if err == nil && len(strings.TrimSpace(string(output))) > 0 {
        return true
    }

    if strings.HasPrefix(dnsIP, "10.192.") {
        return true
    }
    
    return false
}

func FixDNSLeaks() error {
    Notify(NOTIFY_DNS, "Fixing DNS leaks...")
    if err := ConfigDnsmasq(); err != nil {
        return fmt.Errorf("failed to configure dnsmasq: %v", err)
    }

    if err := ConfigureResolv(); err != nil {
        return fmt.Errorf("failed to configure resolv.conf: %v", err)
    }

    services := []string{
        "systemd-resolved",
        "dnsmasq",
        "tor@default",
    }
    
    for _, service := range services {
        if err := exec.Command("sudo", "systemctl", "restart", service).Run(); err != nil {
            Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to restart %s: %v", service, err))
        }
        time.Sleep(500 * time.Millisecond)
    }

    cmd := exec.Command("dig", "+short", "google.com", "@127.0.0.1", "-p", "5353")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("DNS through Tor failed: %v", err)
    }
    
    if len(strings.TrimSpace(string(output))) == 0 {
        return fmt.Errorf("DNS through Tor returned empty response")
    }
    
    Notify(NOTIFY_SUCCESS, "DNS leaks fixed successfully")
    return nil
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

func isTorExitNode(ip string) bool {
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get("https://check.torproject.org/torbulkexitlist")
    if err != nil {
        return false
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return false
    }

    return strings.Contains(string(body), ip)
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

func stripColors(text string) string {
    re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return re.ReplaceAllString(text, "")
}

// ============= TOR STATUS FUNCTIONS =============

func TorStatus(retryCount int) {
    Notify(NOTIFY_TOR, "Checking Tor network status...")

    if retryCount >= 3 {
        Notify(NOTIFY_ERROR, "Failed to check Tor status after 3 attempts")
        // Try to diagnose the issue
        if err := testTorConnectivity(); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Diagnosis: %v", err))
        }
        return
    }

    // First test connectivity
    if err := testTorConnectivity(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Connectivity issue: %v, retrying...", err))
        time.Sleep(9 * time.Second)
        TorStatus(retryCount + 1)
        return
    }

    result, err := CheckTorStatus()
    if err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Status check failed: %v, retrying...", err))
        time.Sleep(9 * time.Second)
        TorStatus(retryCount + 1)
        return
    }
    displayTorResult(result)
}

func CheckTorStatus() (map[string]string, error) {
    // Try multiple methods to check Tor status
    
    // Method 1: Check Tor's control port
    cmd := exec.Command("timeout", "5", "torify", "curl", "-s", "https://check.torproject.org/api/ip")
    output, err := cmd.Output()
    if err == nil {
        var result map[string]interface{}
        if err := json.Unmarshal(output, &result); err == nil {
            if isTor, ok := result["IsTor"].(bool); ok && isTor {
                return map[string]string{
                    "ip":    result["IP"].(string),
                    "title": "Congratulations. You are using Tor.",
                }, nil
            }
        }
    }
    
    // Method 2: Try direct HTTP with proxy
    client := &http.Client{
        Timeout: 15 * time.Second,
        Transport: &http.Transport{
            Proxy: http.ProxyURL(&url.URL{
                Scheme: "socks5",
                Host:   "127.0.0.1:9050",
            }),
        },
    }
    
    resp, err := client.Get("https://check.torproject.org/api/ip")
    if err == nil {
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)
        if err == nil {
            var result map[string]interface{}
            if err := json.Unmarshal(body, &result); err == nil {
                if isTor, ok := result["IsTor"].(bool); ok && isTor {
                    return map[string]string{
                        "ip":    result["IP"].(string),
                        "title": "Congratulations. You are using Tor.",
                    }, nil
                }
            }
        }
    }
    
    // Method 3: Check if Tor is running locally
    if isTorRunning() {
        // Get IP from Tor via SOCKS
        cmd = exec.Command("timeout", "5", "torsocks", "curl", "-s", "https://check.torproject.org/api/ip")
        output, err = cmd.Output()
        if err == nil {
            var result map[string]interface{}
            if err := json.Unmarshal(output, &result); err == nil {
                if isTor, ok := result["IsTor"].(bool); ok && isTor {
                    return map[string]string{
                        "ip":    result["IP"].(string),
                        "title": "Congratulations. You are using Tor.",
                    }, nil
                }
            }
        }
    }
    
    // Method 4: Get IP from local Tor
    ip := getPublicIP()
    if ip != "" {
        return map[string]string{
            "ip":    ip,
            "title": "Unknown - Tor may not be working properly",
        }, nil
    }
    
    return nil, fmt.Errorf("could not verify Tor status")
}

func testTorConnectivity() error {
    Notify(NOTIFY_TOR, "Testing Tor connectivity ...")
    
    // Test 1: Check if Tor is running
    if !isTorRunning() {
        return fmt.Errorf("Tor service is not running ...")
    }
    
    // Test 2: Test SOCKS proxy
    conn, err := net.DialTimeout("tcp", "127.0.0.1:9050", 5*time.Second)
    if err != nil {
        return fmt.Errorf("Tor SOCKS proxy not responding: %v", err)
    }
    conn.Close()
    
    // Test 3: Test DNS through Tor
    cmd := exec.Command("timeout", "5", "dig", "+short", "google.com", "@127.0.0.1", "-p", "5353")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("DNS through Tor failed: %v", err)
    }
    if len(strings.TrimSpace(string(output))) == 0 {
        return fmt.Errorf("DNS through Tor returned empty response")
    }
    
    Notify(NOTIFY_SUCCESS, "Tor connectivity test passed")
    return nil
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

func displayTorResult(result map[string]string) {
    fmt.Printf("\n%s%s[>] %s%sYour IP address: %s%s%s", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Yellow, bcolors.Cyan, result["ip"], bcolors.Endc)

    if strings.Contains(result["title"], "Congratulations") {
        Notify(NOTIFY_SUCCESS, result["title"])
        Notify(NOTIFY_SUCCESS, "Your connection is secured through Tor...")
    } else {
        Notify(NOTIFY_ERROR, result["title"])
        Notify(NOTIFY_WARNING, "Warning: You are NOT using the Tor network!")
        Notify(NOTIFY_INFO, "For enhanced privacy, consider using Tor network...")
    }

    fmt.Printf("%s", "\n")
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

    Notify(NOTIFY_TOR, "Scrambling Tor Nodes")
    Notify(NOTIFY_TOR, fmt.Sprintf("Your new IP appears to be: %s", ip))
    return nil
}

// ============= ADDITIONAL FUNCTIONS =============

func StartServices() error {
    Notify(NOTIFY_INFO, fmt.Sprintf("Starting macchanger, dnsmasq, privoxy, squid, tor"))

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
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to start %s: %v", service, err))
            failedServices = append(failedServices, service)
        } else {
            startedServices = append(startedServices, service)
            time.Sleep(500 * time.Millisecond)
        }
    }

    Notify(NOTIFY_INFO, "Verifying services...")
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
        Notify(NOTIFY_INFO, fmt.Sprintf("Running services: %s", strings.Join(runningServices, ", ")))
    }

    if len(notRunningServices) > 0 {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Started but not running: %s", strings.Join(notRunningServices, ", ")))
    }

    if len(failedServices) > 0 {
        Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to start: %s", strings.Join(failedServices, ", ")))
    }

    Notify(NOTIFY_INFO, "Enabling services to start on boot...")
    for _, service := range runningServices {
        exec.Command("systemctl", "enable", service).Run()
    }

    if len(failedServices) > 0 {
        return fmt.Errorf("failed to start %d services: %s", len(failedServices), strings.Join(failedServices, ", "))
    }

    if len(notRunningServices) > 0 {
        return fmt.Errorf("%d services started but not running: %s", len(notRunningServices), strings.Join(notRunningServices, ", "))
    }

    Notify(NOTIFY_SUCCESS, "All services started successfully...")
    return nil
}

func ConfigProxyEnv() {
    filePath := "/etc/profile.d/proxy.sh"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        proxyContent := `# Generated by africana-framework
export http_proxy="http://127.0.0.1:3129"
export https_proxy="http://127.0.0.1:3129"
export ftp_proxy="http://127.0.0.1:3129"
export all_proxy="socks5://127.0.0.1:9050"
export no_proxy="localhost,127.0.0.1"
`
        err := os.WriteFile(filePath, []byte(proxyContent), 0644)
        if err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to write proxy environment: %v", err))
            return
        }
        Notify(NOTIFY_SUCCESS, "Proxy environment variables configured")
    }
}

func WaitForDNS() error {
    maxAttempts := 10
    for i := 0; i < maxAttempts; i++ {
        cmd := exec.Command("dig", "+short", "google.com", "@127.0.0.1", "-p", "5353")
        output, err := cmd.Output()
        if err == nil && len(strings.TrimSpace(string(output))) > 0 {
            Notify(NOTIFY_SUCCESS, fmt.Sprintf("DNS is ready after %d attempts", i+1))
            return nil
        }
        Notify(NOTIFY_INFO, fmt.Sprintf("Waiting for DNS to be ready... (%d/%d)", i+1, maxAttempts))
        time.Sleep(2 * time.Second)
    }
    return fmt.Errorf("DNS service not ready after %d attempts", maxAttempts)
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

func CheckDNSmasq() error {
    cmd := exec.Command("dig", "+short", "google.com", "@127.0.0.1", "-p", "5353")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("dnsmasq not responding: %v", err)
    }
    if len(strings.TrimSpace(string(output))) == 0 {
        return fmt.Errorf("dnsmasq returned empty response")
    }
    Notify(NOTIFY_SUCCESS, "DNSmasq is working properly")
    return nil
}

func RestoreOriginalConfigs() error {
    var restoredFiles []string
    var failedFiles []string

    // Define all configuration files to restore
    backups := []ConfigBackup{
        {
            OriginalPath: "/etc/resolv.conf",
            BackupPath:   "/etc/resolv.conf.backup_torsocks",
            ServiceName:  "systemd-resolved",
            RestartCmd:   "systemctl restart systemd-resolved",
        },
        {
            OriginalPath: "/etc/tor/torrc",
            BackupPath:   "/etc/tor/torrc.bak_torsocks",
            ServiceName:  "tor",
            RestartCmd:   "systemctl restart tor",
        },
        {
            OriginalPath: "/etc/privoxy/config",
            BackupPath:   "/etc/privoxy/config.bak_africana",
            ServiceName:  "privoxy",
            RestartCmd:   "systemctl restart privoxy",
        },
        {
            OriginalPath: "/etc/squid/squid.conf",
            BackupPath:   "/etc/squid/squid.conf.bak_africana",
            ServiceName:  "squid",
            RestartCmd:   "systemctl restart squid",
        },
        {
            OriginalPath: "/etc/dnsmasq.conf",
            BackupPath:   "/etc/dnsmasq.conf.bak_africana",
            ServiceName:  "dnsmasq",
            RestartCmd:   "systemctl restart dnsmasq",
        },
        {
            OriginalPath: "/etc/dhcp/dhclient.conf",
            BackupPath:   "/etc/dhcp/dhclient.conf.bak_africana",
            ServiceName:  "dhclient",
            RestartCmd:   "systemctl restart networking",
        },
    }

    for _, backup := range backups {
        Notify(NOTIFY_CONFIG, fmt.Sprintf("Restoring %s...", filepath.Base(backup.OriginalPath)))

        if err := restoreSingleConfig(backup); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed: %v", err))
            failedFiles = append(failedFiles, filepath.Base(backup.OriginalPath))
        } else {
            Notify(NOTIFY_SUCCESS, fmt.Sprintf("Restored %s successfully", filepath.Base(backup.OriginalPath)))
            restoredFiles = append(restoredFiles, filepath.Base(backup.OriginalPath))
        }
        time.Sleep(200 * time.Millisecond)
    }

    // Remove proxy environment variables
    if err := removeProxyEnv(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to remove proxy env: %v", err))
    }

    // Restart services that were modified
    if err := restartModifiedServices(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Failed to restart services: %v", err))
    }

    // Print summary
    fmt.Printf(`
  %s✅ Restored%s : %s%d%s files
  %s❌ Failed%s   : %s%d%s files

`,
        bcolors.Green, bcolors.Endc, bcolors.Cyan, len(restoredFiles), bcolors.Endc,
        bcolors.Red, bcolors.Endc, bcolors.Cyan, len(failedFiles), bcolors.Endc)

    if len(failedFiles) > 0 {
        return fmt.Errorf("failed to restore %d files", len(failedFiles))
    }

    Notify(NOTIFY_SUCCESS, "Original configurations restored successfully")
    return nil
}

func restoreSingleConfig(backup ConfigBackup) error {
    // Check if backup exists
    if _, err := os.Stat(backup.BackupPath); os.IsNotExist(err) {
        // Backup doesn't exist, try to restore from system default
        return restoreFromSystemDefault(backup.OriginalPath)
    }

    // Read backup content
    backupContent, err := os.ReadFile(backup.BackupPath)
    if err != nil {
        return fmt.Errorf("failed to read backup: %v", err)
    }

    // Check if original exists and is different
    if _, err := os.Stat(backup.OriginalPath); err == nil {
        originalContent, err := os.ReadFile(backup.OriginalPath)
        if err == nil && string(originalContent) == string(backupContent) {
            // Files are identical, no need to restore
            return nil
        }
    }

    // Restore the backup
    if err := os.WriteFile(backup.OriginalPath, backupContent, 0644); err != nil {
        return fmt.Errorf("failed to restore file: %v", err)
    }

    // Remove backup after successful restore
    if err := os.Remove(backup.BackupPath); err != nil {
        // Non-critical error, just log it
        Notify(NOTIFY_WARNING, fmt.Sprintf("Could not remove backup: %v", err))
    }

    return nil
}

func restoreFromSystemDefault(filePath string) error {
    // Try to get default configuration from system
    var defaultContent string

    switch filePath {
    case "/etc/resolv.conf":
        defaultContent = `# Generated by systemd-resolved
nameserver 8.8.8.8
nameserver 8.8.4.4
`
    case "/etc/tor/torrc":
        defaultContent = `## Configuration file for Tor
## See 'man tor' for more options

SocksPort 9050
`
    case "/etc/privoxy/config":
        defaultContent = `# Privoxy configuration
listen-address 127.0.0.1:8118
forward-socks5t / 127.0.0.1:9050 .
`
    case "/etc/squid/squid.conf":
        defaultContent = `# Squid configuration
http_port 3128
http_access allow localnet
http_access allow localhost
http_access deny all
`
    case "/etc/dnsmasq.conf":
        defaultContent = `# Dnsmasq configuration
port=53
listen-address=127.0.0.1
`
    case "/etc/dhcp/dhclient.conf":
        defaultContent = `# Dhclient configuration
`
    default:
        return fmt.Errorf("no default template for %s", filePath)
    }

    if err := os.WriteFile(filePath, []byte(defaultContent), 0644); err != nil {
        return fmt.Errorf("failed to write default config: %v", err)
    }

    return nil
}

func removeProxyEnv() error {
    proxyFiles := []string{
        "/etc/profile.d/proxy.sh",
        "/etc/profile.d/proxy.csh",
        "/etc/environment",
    }

    for _, file := range proxyFiles {
        if _, err := os.Stat(file); err == nil {
            content, err := os.ReadFile(file)
            if err == nil {
                // Remove proxy lines
                lines := strings.Split(string(content), "\n")
                var newLines []string
                for _, line := range lines {
                    if !strings.Contains(line, "proxy") && 
                       !strings.Contains(line, "PROXY") &&
                       !strings.Contains(line, "socks") &&
                       !strings.Contains(line, "Socks") {
                        newLines = append(newLines, line)
                    }
                }
                if len(newLines) != len(lines) {
                    if err := os.WriteFile(file, []byte(strings.Join(newLines, "\n")), 0644); err != nil {
                        return fmt.Errorf("failed to remove proxy from %s: %v", file, err)
                    }
                }
            }
        }
    }
    return nil
}

func restartModifiedServices() error {
    services := []string{
        "systemd-resolved",
        "tor",
        "privoxy",
        "squid",
        "dnsmasq",
        "networking",
    }

    Notify(NOTIFY_INFO, "Restarting modified services...")

    var restarted []string
    var failed []string

    for _, service := range services {
        // Check if service exists
        cmd := exec.Command("systemctl", "list-unit-files", service+".service")
        output, _ := cmd.Output()
        if !strings.Contains(string(output), service+".service") {
            continue
        }

        cmd = exec.Command("systemctl", "restart", service)
        if err := cmd.Run(); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to restart %s: %v", service, err))
            failed = append(failed, service)
        } else {
            Notify(NOTIFY_SUCCESS, fmt.Sprintf("Restarted %s", service))
            restarted = append(restarted, service)
        }
        time.Sleep(100 * time.Millisecond)
    }

    if len(failed) > 0 {
        return fmt.Errorf("failed to restart: %s", strings.Join(failed, ", "))
    }
    return nil
}

// ============= CHECK CURRENT CONFIGURATION STATUS =============

func CheckConfigStatus() error {

    configs := map[string]string{
        "/etc/resolv.conf":       "DNS Configuration",
        "/etc/tor/torrc":         "Tor Configuration",
        "/etc/privoxy/config":    "Privoxy Configuration",
        "/etc/squid/squid.conf":  "Squid Configuration",
        "/etc/dnsmasq.conf":      "Dnsmasq Configuration",
    }

    Notify(NOTIFY_INFO, "Checking configuration status...")

    for path, name := range configs {
        if _, err := os.Stat(path); err == nil {
            content, err := os.ReadFile(path)
            if err == nil {
                // Check if it's our config
                isAfricana := strings.Contains(string(content), "africana-framework")
                isTorsocks := strings.Contains(string(content), "torsocks")
                
                status := "Default"
                color := bcolors.Green
                if isAfricana || isTorsocks {
                    status = "Modified by Tor"
                    color = bcolors.Yellow
                }
                
                fmt.Printf("  %s•%s %-25s : %s%s%s\n", 
                    bcolors.Cyan, bcolors.Endc, 
                    name, 
                    color, status, bcolors.Endc)
            }
        } else {
            fmt.Printf("  %s•%s %-25s : %s[NOT FOUND]%s\n", 
                bcolors.Cyan, bcolors.Endc, 
                name, 
                bcolors.Red, bcolors.Endc)
        }
    }

    // Check iptables status
    cmd := exec.Command("iptables", "-L", "-n")
    output, err := cmd.Output()
    if err == nil {
        isTorRedirect := strings.Contains(string(output), "REDIRECT") && 
                         strings.Contains(string(output), "9040")
        if isTorRedirect {
            fmt.Printf("  %s•%s %-25s : %s%s%s\n", 
                bcolors.Cyan, bcolors.Endc,
                "Firewall Rules",
                bcolors.Yellow, "Modified for Tor", bcolors.Endc)
        } else {
            fmt.Printf("  %s•%s %-25s : %s%s%s\n", 
                bcolors.Cyan, bcolors.Endc,
                "Firewall Rules",
                bcolors.Green, "Default", bcolors.Endc)
        }
    }

    fmt.Printf("\n")
    return nil
}

// ============= UPDATE KillTor to use RestoreOriginalConfigs =============

func KillTor() error {
    defer StopServices()

    trigger := 0
    Notify(NOTIFY_CONFIG, "Reverting to default configurations...")

    // Restore all original configurations
    if err := RestoreOriginalConfigs(); err != nil {
        Notify(NOTIFY_WARNING, fmt.Sprintf("Some configurations could not be restored: %v", err))
    }

    trigger++
    Notify(NOTIFY_FW, "Restoring Default Iptables rules...")

    if _, err := os.Stat("/etc/iptables_rules_torsocks.bak"); err == nil {
        cmd := exec.Command("bash", "-c", "sudo iptables-restore < /etc/iptables_rules_torsocks.bak")
        if output, err := cmd.CombinedOutput(); err != nil {
            Notify(NOTIFY_ERROR, fmt.Sprintf("Failed to restore iptables from backup: %v\nOutput: %s", err, string(output)))
            Notify(NOTIFY_WARNING, "Falling back to default reset...")
            if err := ResetToDefault(true, true); err != nil {
                return fmt.Errorf("%s%s[!] %sFailed to reset to default: %v\n", 
                    bcolors.Bold, bcolors.Red, bcolors.Endc, err)
            }
        } else {
            Notify(NOTIFY_SUCCESS, "Successfully restored from backup")
            os.Remove("/etc/iptables_rules_torsocks.bak")
        }
    } else {
        Notify(NOTIFY_WARNING, "No backup found, resetting to default rules...")
        if err := ResetToDefault(true, true); err != nil {
            return fmt.Errorf("%s%s[!] %sFailed to reset to default: %v\n", 
                bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        }
    }
    time.Sleep(1 * time.Second)

    if trigger == 0 {
        return fmt.Errorf("%s%sNo instances of torsocks have been executed%s", 
            bcolors.BrightRed, bcolors.Bold, bcolors.Endc)
    } else {
        Notify(NOTIFY_SUCCESS, "Everything completed...")
    }

    Notify(NOTIFY_INFO, "Preparing to stop all services...")

    return nil
}

func CreateConfigBackups() error {
    Notify(NOTIFY_CONFIG, "Creating backups of current configurations...")

    configs := []string{
        "/etc/resolv.conf",
        "/etc/tor/torrc",
        "/etc/privoxy/config",
        "/etc/squid/squid.conf",
        "/etc/dnsmasq.conf",
        "/etc/dhcp/dhclient.conf",
    }

    for _, config := range configs {
        if _, err := os.Stat(config); err == nil {
            backup := config + ".backup_torsocks"
            if _, err := os.Stat(backup); os.IsNotExist(err) {
                if err := exec.Command("cp", config, backup).Run(); err != nil {
                    return fmt.Errorf("failed to backup %s: %v", config, err)
                }
                Notify(NOTIFY_SUCCESS, fmt.Sprintf("Backed up %s", filepath.Base(config)))
            }
        }
    }
    return nil
}
