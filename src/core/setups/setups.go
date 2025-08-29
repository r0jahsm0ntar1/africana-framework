//John 3:16

package setups

import(
    "os"
    "fmt"
    "time"
    "menus"
    "utils"
    "strconv"
    "os/exec"
    "strings"
    "bcolors"
    "banners"
    "runtime"
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

var defaultValues = map[string]string{
    "module":       "",
    "function":     "",
    "distro":       utils.Distro,
    "proxy":        utils.Proxies,
    "proxies":      utils.Proxies,
    "pyenvname":    utils.PyEnvName,
}

var (
    systemTools = map[string]string{
        "aha":                          "aha",
        "aircrack-ng":                  "aircrack-ng",
        "build-essential":              "build-essential",
        "checkinstall":                 "checkinstall",
        "cowpatty":                     "cowpatty",
        "dhcpd":                        "dhcpd",
        "dnsmasq":                      "dnsmasq",
        "docker.io":                    "docker.io",
        "dsniff":                       "dsniff",
        "ftp":                          "ftp",
        "gawk":                         "gawk",
        "gcc":                          "gcc",
        "gdb":                          "gdb",
        "github-desktop":               "github-desktop",
        "golang-go":                    "golang-go",
        "gstreamer1.0-libav":           "gstreamer1.0-libav",
        "gunzip":                       "gunzip",
        "hostapd":                      "hostapd",
        "hydra":                        "hydra",
        "iptables":                     "iptables",
        "jq":                           "jq",
        "lcov":                         "lcov",
        "libbz2-dev":                   "libbz2-dev",
        "libc6-dev":                    "libc6-dev",
        "libffi-dev":                   "libffi-dev",
        "libgdbm-compat-dev":           "libgdbm-compat-dev",
        "libgdbm-dev":                  "libgdbm-dev",
        "libgeoip-dev":                 "libgeoip-dev",
        "libgeoip1t64":                 "libgeoip1t64",
        "liblzma-dev":                  "liblzma-dev",
        "libncurses-dev":               "libncurses-dev",
        "libpcap-dev":                  "libpcap-dev",
        "libreadline-dev":              "libreadline-dev",
        "libsqlite3-dev":               "libsqlite3-dev",
        "libssl-dev":                   "libssl-dev",
        "lighttpd":                     "lighttpd",
        "macchanger":                   "macchanger",
        "mdk3":                         "mdk3",
        "mdk4":                         "mdk4",
        "mousepad":                     "mousepad",
        "ncat":                         "ncat",
        "net-tools":                    "net-tools",
        "nmap":                         "nmap",
        "npm":                          "npm",
        "openssh-client":               "openssh-client",
        "openssh-server":               "openssh-server",
        "openssl":                      "openssl",
        "php-cgi":                      "php-cgi",
        "pkg-config":                   "pkg-config",
        "playit":                       "playit",
        "powershell":                   "powershell",
        "privoxy":                      "privoxy",
        "python3":                      "python3",
        "python3-dev":                  "python3-dev",
        "python3-geoip":                "python3-geoip",
        "python3-pip":                  "python3-pip",
        "python3-pycurl":               "python3-pycurl",
        "python3-requests":             "python3-requests",
        "python3-scapy":                "python3-scapy",
        "python3-venv":                 "python3-venv",
        "python3-whois":                "python3-whois",
        "python-dev-is-python3":        "python-dev-is-python3",
        "rfkill":                       "rfkill",
        "rlwrap":                       "rlwrap",
        "squid":                        "squid",
        "tk-dev":                       "tk-dev",
        "tmux":                         "tmux",
        "tor":                          "tor",
        "uuid-dev":                     "uuid-dev",
        "wine32":                       "wine32:i386",
        "xterm":                        "xterm",
        "zlib1g-dev":                   "zlib1g-dev",
        "zsh":                          "zsh",
        "go-winres":                    "github.com/tc-hib/go-winres@latest",
    }

    securityTools = map[string]string{
        "airgeddon":                    "airgeddon",
        "commix":                       "commix",
        "dnsenum":                      "dnsenum",
        "dnsrecon":                     "dnsrecon",
        "findomain":                    "findomain",
        "feroxbuster":                  "feroxbuster",
        "gobuster":                     "gobuster",
        "gophish":                      "gophish",
        "massdns":                      "massdns",
        "metasploit-framework":         "metasploit-framework",
        "netexec":                      "netexec",
        "nikto":                        "nikto",
        "nuclei":                       "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest",
        "smbmap":                       "smbmap",
        "sqlmap":                       "sqlmap",
        "sslscan":                      "sslscan",
        "uniscan":                      "uniscan",
        "wapiti":                       "wapiti",
        "whatweb":                      "whatweb",
        "wifipumpkin3":                 "wifipumpkin3",
        "wifite":                       "wifite",
        "xsser":                        "xsser",
    }

    projectDiscoveryTools = map[string]string{
        "anew":                         "github.com/tomnomnom/anew@latest",
        "assetfinder":                  "github.com/tomnomnom/assetfinder@latest",
        "asnmap":                       "github.com/projectdiscovery/asnmap/cmd/asnmap@latest",
        "dalfox":                       "github.com/hahwul/dalfox/v2@latest",
        "dnsx":                         "github.com/projectdiscovery/dnsx/cmd/dnsx@latest",
        "gau":                          "github.com/lc/gau/v2/cmd/gau@latest",
        "gf":                           "github.com/tomnomnom/gf@latest",
        "hakrawler":                    "github.com/hakluke/hakrawler@latest",
        "httpx":                        "github.com/projectdiscovery/httpx/cmd/httpx@latest",
        "interactsh":                   "github.com/projectdiscovery/interactsh/cmd/interactsh-client@latest",
        "katana":                       "github.com/projectdiscovery/katana/cmd/katana@latest",
        "mapcidr":                      "github.com/projectdiscovery/mapcidr/cmd/mapcidr@latest",
        "naabu":                        "github.com/projectdiscovery/naabu/v2/cmd/naabu@latest",
        "notify":                       "github.com/projectdiscovery/notify/cmd/notify@latest",
        "subfinder":                    "github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest",
        "tlsx":                         "github.com/projectdiscovery/tlsx/cmd/tlsx@latest",
        "uncover":                      "github.com/projectdiscovery/uncover/cmd/uncover@latest",
        "waybackurls":                  "github.com/tomnomnom/waybackurls@latest",
    }
)

func SetupsLauncher() {
    for {
        fmt.Printf("%s%s%safr%s%s packages(%s%score/setups_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutarilas},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, menus.HelpInfoSetups},
        {[]string{"m", "menu"}, menus.MenuOne},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.SetupsOptions(utils.Distro, Function)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListSetupsFunction},
        {[]string{"distro", "distros", "list distro", "list distros", "show distro", "show distros"}, menus.ListSetupsDistros},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run kali", "use kali", "exec kali", "start kali", "launch kali", "exploit kali", "execute kali"}, func() {SetupsFunction(Function, "kali")}},
        {[]string{"? 1", "info 1", "help 1", "kali", "info kali", "help kali"}, menus.HelpInfoKali},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run ubuntu", "use ubuntu", "exec ubuntu", "start ubuntu", "launch ubuntu", "exploit ubuntu", "execute ubuntu"}, func() {SetupsFunction(Function, "ubuntu")}},
        {[]string{"? 2", "info 2", "help 2", "ubuntu", "info ubuntu", "help ubuntu"}, menus.HelpInfoUbuntu},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run arch", "use arch", "exec arch", "start arch", "launch arch", "exploit arch", "execute arch"}, func() {SetupsFunction(Function, "arch")}},
        {[]string{"? 3", "info 3", "help 3", "arch", "info arch", "help arch"}, menus.HelpInfoArch},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run macos", "use macos", "exec macos", "start macos", "launch macos", "exploit macos", "execute macos"}, func() {SetupsFunction(Function, "macos")}},
        {[]string{"? 4", "info 4", "help 4", "macos", "info macos", "help macos"}, menus.HelpInfoMacos},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run android", "use android", "exec android", "start android", "launch android", "exploit android", "execute android"}, func() {SetupsFunction(Function, "android")}},
        {[]string{"? 5", "info 5", "help 5", "android", "info android", "help android"}, menus.HelpInfoAndroid},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run windows", "use windows", "exec windows", "start windows", "launch windows", "exploit windows", "execute windows"}, func() {SetupsFunction(Function, "windows")}},
        {[]string{"? 6", "info 6", "help 6", "windows", "info windows", "help windows"}, menus.HelpInfoWindows},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run update", "use update", "exec update", "start update", "launch update", "exploit update", "execute update"}, func() {SetupsFunction("update", utils.Distro)}},
        {[]string{"? 7", "info 7", "help 7", "update", "info update", "help update"}, menus.HelpInfoUpdate},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run repair", "use repair", "exec repair", "start repair", "launch repair", "exploit repair", "execute repair"}, func() {SetupsFunction("repair", utils.Distro)}},
        {[]string{"? 8", "info 8", "help 8", "repair", "info repair", "help repair"}, menus.HelpInfoRepair},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run uninstall", "use uninstall", "exec uninstall", "start uninstall", "launch uninstall", "exploit uninstall", "execute uninstall"}, func() {SetupsFunction("uninstall", utils.Distro)}},
        {[]string{"? 9", "info 9", "help 9", "uninstall", "info uninstall", "help uninstall"}, menus.HelpInfoUninstall},

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

        "func": &Function,
        "funcs": &Function,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
        "distro": &utils.Distro,
        "proxies": &utils.Proxies,
        "pyenvname": &utils.PyEnvName,
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

        "func": &Function,
        "funcs": &Function,
        "module": &Function,
        "function": &Function,
        "functions": &Function,
        "distro": &utils.Distro,
        "proxies": &utils.Proxies,
        "pyenvname": &utils.PyEnvName,
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
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    SetupsFunction(Function, utils.Distro)
}

func SetupsFunction(Function, Distro string, args ...interface{}) {
    os.MkdirAll(utils.SetupsLogs, os.ModePerm)

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.SeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            //RHOST: utils.RHost,
            //LHOST: utils.LHost,
            DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.SetupsLogs,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            //RECONDIR: ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            TOOLSDIR: utils.ToolsDir,
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
            MODE: utils.SeMode,
            //IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            //RHOST: utils.RHost,
            //LHOST: utils.LHost,
            DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.SetupsLogs,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, true)
    }

    commands := map[string]func(){

        "0": func() {AutoSetups()},

        "1": func() {KaliSetups()},
        "2": func() {UbuntuSetups()},
        "3": func() {ArchSetups()},
        "4": func() {MacosSetups()},
        "5": func() {AndroidSetups()},
        "6": func() {WindowsSetups()},
        "7": func() {UpdateAfricana()},
        "8": func() {UpdateAfricana()},
        "9": func() {Uninstaller()},

        "auto":      func() {AutoSetups()},

        "kali":      func() {KaliSetups()},
        "ubuntu":    func() {UbuntuSetups()},
        "arch":      func() {ArchSetups()},
        "macos":     func() {MacosSetups()},
        "android":   func() {AndroidSetups()},
        "windows":   func() {WindowsSetups()},

        "install":   func() {Installer(utils.Distro)},
        "update":    func() {UpdateAfricana()},
        "repair":    func() {UpdateAfricana()},
        "uninstall": func() {Uninstaller()},

    }

    textCommands := []string{"auto", "install", "update", "repair", "uninstall"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListSetupsFunction()
        return
    }

    lowerInput := strings.ToLower(Function)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("\n%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, Function)
    menus.ListSetupsFunction()
}

func Installer(Distro string) {
    matchers := []stringMatcher{

        {[]string{"0", "auto"}, AutoSetups},
        {[]string{"1", "arch"}, ArchSetups},
        {[]string{"2", "kali"}, KaliSetups},
        {[]string{"3", "macos"}, MacosSetups},
        {[]string{"4", "ubuntu"}, UbuntuSetups},
        {[]string{"5", "android"}, AndroidSetups},
        {[]string{"6", "windows"}, WindowsSetups},
        {[]string{"7", "update"}, UpdateAfricana},
        {[]string{"8", "repair"}, UpdateAfricana},
        {[]string{"9", "uninstall"}, Uninstaller},
    }

    distroLower := strings.ToLower(utils.Distro)
    for _, m := range matchers {
        for _, name := range m.names {
            if name == distroLower {
                m.action()
                return
            }
        }
    }

    fmt.Printf("%s[!] Error: %sInvalid DISTRO %s. Use %s'help' %sfor commands.\n", bcolors.Red, bcolors.Endc, utils.Distro, bcolors.Green, bcolors.Endc)
}

func CheckTools() {
    spinner := utils.New(
        utils.WithStyle("classic"),
        utils.WithEffect("bounce"),
    )
    spinner.Start()
    missingTools := UpsentTools()
    spinner.Stop()

    totalMissing := len(missingTools["system"]) + len(missingTools["security"]) + len(missingTools["discovery"])
    if totalMissing == 0 {
        fmt.Printf("%s[+] %sAll tools are installed and ready!\n", bcolors.Green, bcolors.Endc)
        return
    }

    if len(missingTools["system"]) > 0 {
        fmt.Printf("%s[!] %s%sMissing system tools.%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["system"] {
            fmt.Printf("  %s- %s%s...\n", bcolors.Bold, bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["security"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing security tools.%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["security"] {
            fmt.Printf("  %s- %s%s...\n", bcolors.Bold, bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["discovery"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing project discovery tools.%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["discovery"] {
            fmt.Printf("  %s- %s%s...\n", bcolors.Bold, bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }
    userChoice()
}

func userChoice() {
    fmt.Printf("\n%s[?] %sLaunch setups to install missing tools? (y/n): ", bcolors.Yellow, bcolors.Endc)
    utils.Scanner.Scan()
    Input := strings.TrimSpace(utils.Scanner.Text())
    switch Input {
    case "y", "yes":
       AutoSetups()
       return
    case "n", "q", "no", "exit", "quit":
        fmt.Printf("%s[!] %sInstallation skipped. Some tools are missing ...\n", bcolors.BrightRed, bcolors.Endc)
    default:
        fmt.Printf("%s[!] %sChoices are (y|n|yes|no)", bcolors.BrightYellow, bcolors.Endc)
        userChoice()
        return
    }
}

func isInstalled(tool string) bool {
    _, err := exec.LookPath(tool)
    if err != nil {
        cmd := exec.Command("dpkg", "-s", tool)
        if cmd.Run() == nil {
            return true
        }
        return false
    }
    return true
}

func UpsentTools() map[string]map[string]string {
    missing := make(map[string]map[string]string)
    missing["system"] = make(map[string]string)
    missing["security"] = make(map[string]string)
    missing["discovery"] = make(map[string]string)

    for tool, pkg := range systemTools {
        if !isInstalled(tool) {
            time.Sleep(45 * time.Millisecond)
            missing["system"][tool] = pkg
        }
    }

    for tool, pkg := range securityTools {
        if !isInstalled(tool) {
            time.Sleep(45 * time.Millisecond)
            missing["security"][tool] = pkg
        }
    }

    for tool, pkg := range projectDiscoveryTools {
        if !isInstalled(tool) {
            time.Sleep(45 * time.Millisecond)
            missing["discovery"][tool] = pkg
        }
    }
    return missing
}

func InstallTools(tools map[string]map[string]string) {
    if len(tools["system"]) > 0 {
        fmt.Printf("\n%sInstalling system tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["system"] {
            fmt.Printf("\n%s[+]  %s%s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            subprocess.Popen("sudo apt install -y %s", pkg)
            time.Sleep(180 * time.Millisecond)
        }
    }

    if len(tools["security"]) > 0 {
        fmt.Printf("\n%sInstalling security tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["security"] {
            fmt.Printf("\n%s[+]  %s%s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            if strings.HasPrefix(pkg, "github.com") {
                subprocess.Popen("go install %s", pkg)
            } else {
                subprocess.Popen("sudo apt install -y %s", pkg)
            }
            time.Sleep(180 * time.Millisecond)
        }
    }

    if len(tools["discovery"]) > 0 {
        fmt.Printf("\n%sInstalling project discovery tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["discovery"] {
            fmt.Printf("\n%s[+]  %s%s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            subprocess.Popen("go install %s", pkg)
            time.Sleep(180 * time.Millisecond)
        }
    }

    if runtime.GOOS == "linux" {
        gzFilePath := utils.RokyPath + ".gz"
        if _, err := os.Stat(utils.RokyPath); os.IsNotExist(err) {
            if _, err := os.Stat(gzFilePath); os.IsNotExist(err) {
                return
            }
            command := fmt.Sprintf("gunzip -d -9 %s", gzFilePath)
            subprocess.Popen(command)
        }
    }
}

var linuxTaskMap = map[string]func(){
    "kali": func() {
        fmt.Printf("%s[+] %sKali distro detected ...", bcolors.BrightGreen, bcolors.Endc)
        KaliSetups()
    },
    "debian": func() {
        fmt.Printf("%s[+] %sDebian distro detected ...", bcolors.BrightGreen, bcolors.Endc)
        KaliSetups()
    },
    "ubuntu": func() {
        fmt.Printf("%s[+] %sUbuntu distro detected ...", bcolors.BrightGreen, bcolors.Endc)
        UbuntuSetups()
    },
    "fedora": func() {
        fmt.Printf("%s[+] %sFedora distro detected ...", bcolors.BrightGreen, bcolors.Endc)
    },
    "arch": func() {
        fmt.Printf("%s[+] %sArch distro detected ...", bcolors.BrightGreen, bcolors.Endc)
        ArchSetups()
    },
    "alpine": func() {
        fmt.Printf("%s[+] %sAlpine distro detected ...", bcolors.BrightGreen, bcolors.Endc)
    },
    // Add more distros as needed
}

func Venv(PyEnvName string) {
    subprocess.Popen("cd %s; python3 -m virtualenv %s --system-site-packages; echo -n '\n source %s%s/bin/activate' >> ~/.zshrc; source ~/.zshrc", utils.BaseDir, utils.PyEnvName, utils.BaseDir, utils.PyEnvName)
}

func InstallFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Popen(cmd)
    }
}

func InstallGithubTools() {
    subprocess.Popen("cd %s; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1; python3 -m pip install -r %s/requirements.txt", utils.BaseDir, utils.BaseDir)
}

func AutoSetups() {
    osName := runtime.GOOS

    switch osName {
    case "linux":
        if utils.DetectAndroid() {
            fmt.Printf("%s[+] %sAndroid Detected ...", bcolors.BrightGreen, bcolors.Endc)
            AndroidSetups()
        } else {
            distroID, err := utils.GetLinuxDistroID()
            if err != nil {
                fmt.Printf("%s[!] %sLinux OS detected but distro detection failed: %v\n", bcolors.BrightRed, bcolors.Endc, err)
            } else if task, ok := linuxTaskMap[distroID]; ok {
                task()
            } else {
                fmt.Printf("%s[!] %sUnknown or unsupported Linux distro: %s", bcolors.BrightRed, bcolors.Endc, distroID)
            }
        }
    case "windows":
        fmt.Printf("%s[+] %sWindows distro detected ...", bcolors.BrightGreen, bcolors.Endc)
    case "darwin":
        fmt.Printf("%s[+] %smacOS distro detected ...", bcolors.BrightGreen, bcolors.Endc)
    default:
        fmt.Printf("%s[!] %sUnsupported or unknown OS: %s", bcolors.BrightRed, bcolors.Endc, osName)
    }
}

func KaliSetups() {
    fmt.Printf("\n%s[>] %sInstalling africana in kali ...\n", bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()
    if _, err := os.Stat(utils.ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            "wget https://archive.kali.org/archive-keyring.gpg -O /usr/share/keyrings/kali-archive-keyring.gpg",
            "cd /etc/apt/trusted.gpg.d/; wget -qO - https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg",
            "cd /etc/apt/sources.list.d/; wget -qO - https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list",
            "dpkg --add-architecture i386",
            "apt-get update -y",
            "apt-get install zsh git curl -y",
        }
        InstallFoundationTools(foundationCommands)
        InstallTools(missingTools)
        subprocess.Popen("winecfg /v win11")
        fmt.Printf("\n%s[*] %sInstalling third party tools\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana succesfully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
        return
    }
}

func UbuntuSetups() {
    fmt.Printf("\n%s[>] %sInstalling africana in ubuntu ...\n", bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()
    if _, err := os.Stat(utils.ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            "wget https://archive.kali.org/archive-keyring.gpg -O /usr/share/keyrings/kali-archive-keyring.gpg",
            "cd /etc/apt/trusted.gpg.d/; wget -qO - https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg",
            "cd /etc/apt/sources.list.d/; wget -qO - https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list",
            "dpkg --add-architecture i386",
            "add-apt-repository multiverse",
            "apt-get update -y; apt-get install zsh* git curl ubuntu-restricted-extras gnome-shell-extension-manager gnome-shell-extensions gnome-tweaks",
        }
        InstallFoundationTools(foundationCommands)
        InstallTools(missingTools)
        subprocess.Popen("winecfg /v win11")
        fmt.Printf("\n%s[*] %sInstalling third party tools\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana succesfully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func ArchSetups() {
    fmt.Printf("\n%s[>] %sInstalling blackarch tools ...\n", bcolors.Green, bcolors.Endc)
    if _, err := os.Stat(utils.ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            "curl -O https://blackarch.org/strap.sh",
            "chmod +x strap.sh",
            "./strap.sh",
            "pacman -Syu --noconfirm",
            "pacman -S --noconfirm blackarch",
            "pacman -S --noconfirm base-devel bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f WordsListDir wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386",
            "winecfg /v win11",
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("\n%s[>] %sInstalling third party tools ...\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana fully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func UpdateAfricana() {
    fmt.Printf("\n%s[!] %sAfricana already installed. Updating it ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("cd %s; git pull .", utils.ToolsDir)
    subprocess.Popen("cd %s; git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1; cd ./africana-framework; make; cd ./build; mv ./* /usr/local/bin/afrconsole; rm -rf ../africana-framework", utils.BaseDir)
    fmt.Printf("\n%s[*] %sAfricana succesfully updated ...\n", bcolors.Green, bcolors.Endc)
    return
}

func AndroidSetups() {
    // Placeholder for Windows setup logic
}

func MacosSetups() {
}

func WindowsSetups() {
    // Placeholder for Windows setup logic
}

func Uninstaller() {
    fmt.Printf("%s\n[!] %sUninstalling africana ...\n", bcolors.BrightRed, bcolors.Endc)
    if _, err := os.Stat(utils.BaseDir); !os.IsNotExist(err) {
        subprocess.Popen("rm -rf %s; rm -rf /usr/local/bin/afrconsole", utils.BaseDir)
        fmt.Printf("%s[*] %sAfricana uninstalled. Goodbye! ...", bcolors.Green, bcolors.Endc)
        os.Exit(0)
    } else {
        fmt.Printf("%s[!] %sAfricana is not installed ...\n", bcolors.Green, bcolors.Endc)
    }
}
