package setups

import(
    "os"
    "fmt"
    "time"
    "bufio"
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
    PyEnvName = "africana-venv"
    scanner = bufio.NewScanner(os.Stdin)
    Input, Proxy, Distro, Function string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "distro": "",
    "module": "",
    "proxies": "",
    "function": "",
    "pyenvname": PyEnvName,
}

type stringMatcher struct {
    names  []string
    action func()
}

var (
    // System tools and packages
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
        //"gnome-shell-extension-manager": "gnome-shell-extension-manager",
        //"gnome-shell-extensions":       "gnome-shell-extensions",
        //"gnome-tweaks":                 "gnome-tweaks",
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
        //"ubuntu-restricted-extras":      "ubuntu-restricted-extras",
        "uuid-dev":                     "uuid-dev",
        "wine32":                       "wine32:i386",
        "xterm":                        "xterm",
        "zlib1g-dev":                   "zlib1g-dev",
        "zsh":                          "zsh",
    }

    // Security tools
    securityTools = map[string]string{
        "airgeddon":                    "airgeddon",
        "commix":                       "commix",
        "dnsenum":                      "dnsenum",
        "dnsrecon":                     "dnsrecon",
        "feroxbuster":                  "feroxbuster",
        "gobuster":                     "gobuster",
        "gophish":                      "gophish",
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

    // Project discovery tools
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
        fmt.Printf("%s%safr3%s packages(%s%score/setups_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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

        //History/Junk commands
        {[]string{"histo", "history", "show history", "log", "logs", "show log", "show logs"}, subprocess.ShowHistory},
        {[]string{"c junk", "c junks", "c output", "c outputs", "clear junk", "clear junks", "clear output", "clear outputs"}, utils.ClearJunks},
        {[]string{"c log", "c logs", "c history", "c histories", "clear log", "clear logs", "clear history", "clear histories"}, subprocess.ClearHistory},
        {[]string{"junk", "junks", "output", "outputs", "show junk", "show junks", "show output", "show outputs", "l junk", "l junks", "l output", "l outputs", "list junk", "list junks", "list output", "list outputs"}, utils.ListJunks},

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
        {[]string{"info"}, menus.HelpInfoSetups},
        {[]string{"m", "menu"}, menus.MenuOne},
        {[]string{"option", "options", "show option", "show options"}, menus.SetupsOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListSetupsFunction},
        {[]string{"distro", "distros", "list distro", "list distros", "show distro", "show distros"}, menus.ListSetupsDistros},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run kali", "use kali", "exec kali", "start kali", "launch kali", "exploit kali", "execute kali"}, func() { autoExecuteFunc("kali", "install") }},
        {[]string{"? 1", "info 1", "help 1", "kali", "info kali", "help kali"}, menus.HelpInfoKali},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run ubuntu", "use ubuntu", "exec ubuntu", "start ubuntu", "launch ubuntu", "exploit ubuntu", "execute ubuntu"}, func() { autoExecuteFunc("ubuntu", "install") }},
        {[]string{"? 2", "info 2", "help 2", "ubuntu", "info ubuntu", "help ubuntu"}, menus.HelpInfoUbuntu},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run arch", "use arch", "exec arch", "start arch", "launch arch", "exploit arch", "execute arch"}, func() { autoExecuteFunc("arch", "install") }},
        {[]string{"? 3", "info 3", "help 3", "arch", "info arch", "help arch"}, menus.HelpInfoArch},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run macos", "use macos", "exec macos", "start macos", "launch macos", "exploit macos", "execute macos"}, func() { autoExecuteFunc("macos", "install") }},
        {[]string{"? 4", "info 4", "help 4", "macos", "info macos", "help macos"}, menus.HelpInfoMacos},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run android", "use android", "exec android", "start android", "launch android", "exploit android", "execute android"}, func() { autoExecuteFunc("android", "install") }},
        {[]string{"? 5", "info 5", "help 5", "android", "info android", "help android"}, menus.HelpInfoAndroid},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run windows", "use windows", "exec windows", "start windows", "launch windows", "exploit windows", "execute windows"}, func() { autoExecuteFunc("windows", "install") }},
        {[]string{"? 6", "info 6", "help 6", "windows", "info windows", "help windows"}, menus.HelpInfoWindows},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run update", "use update", "exec update", "start update", "launch update", "exploit update", "execute update"}, func() { autoExecuteFunc("update", "install") }},
        {[]string{"? 7", "info 7", "help 7", "update", "info update", "help update"}, menus.HelpInfoUpdate},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run repair", "use repair", "exec repair", "start repair", "launch repair", "exploit repair", "execute repair"}, func() { autoExecuteFunc("repair", "install") }},
        {[]string{"? 8", "info 8", "help 8", "repair", "info repair", "help repair"}, menus.HelpInfoRepair},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run uninstall", "use uninstall", "exec uninstall", "start uninstall", "launch uninstall", "exploit uninstall", "execute uninstall"}, func() { autoExecuteFunc("uninstall", "install") }},
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

        "distro": &Distro,
        "proxies": &Proxy,
        "func":  &Function,
        "module": &Function,
        "function": &Function,
        "pyenvname": &PyEnvName,

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

        "distro": &Distro,
        "proxies": &Proxy,
        "func": &Function,
        "module": &Function,
        "function": &Function,
        "pyenvname": &PyEnvName,
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
        fmt.Printf("\n%s[!] %sMissing required parameter FUNCTION. Use %s'show functions' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if Distro == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'show distros' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    SetupsFunction(Function, Distro)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    Distro = distro
    Function = function
    executeFunction()
}

func SetupsFunction(Function string, args ...interface{}) {
    fmt.Printf("\nDISTRO => %s\nFUNCTION => %s\n", Distro, Function)
    fmt.Printf("\nFunction => %s\n", Function)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    // Command mapping with direct function references
    commands := map[string]func(){
        "install": func() {Installer(Distro)},
         "update": func() {UpdateAfricana()},
         "repair": func() {UpdateAfricana()},
      "uninstall": func() {Uninstaller()},
           //"auto": ,

        // Numeric shortcuts
        "1": func() { KaliSetups()},
        "2": func() { UbuntuSetups()},
        "3": func() { ArchSetups()},
        "4": func() { MacosSetups()},
        "5": func() { AndroidSetups()},
        "6": func() { WindowsSetups()},
        "7": func() { UpdateAfricana()},
        "8": func() { UpdateAfricana()},
        "9": func() { Uninstaller()},
           //"auto": ,
    }

    // Command list for typo checking
    textCommands := []string{"install", "update", "repair", "uninstall"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    // Check if input was a number
    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("\n%s[!] %sNumber %d is invalid. Valid numbers are from 1-10.\n", bcolors.Yellow, bcolors.Endc, num)
        menus.ListSetupsFunction()
        return
    }

    // Check for similar commands
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

    distroLower := strings.ToLower(Distro)
    for _, m := range matchers {
        for _, name := range m.names {
            if name == distroLower {
                m.action()
                return
            }
        }
    }

    fmt.Printf("%s[!] Error: %sInvalid DISTRO %s. Use %s'help' %sfor commands.\n", bcolors.Red, bcolors.Endc, Distro, bcolors.Green, bcolors.Endc)
}

func CheckTools() {
    // Create spinner with custom options
    spinner := utils.New(
        utils.WithStyle("classic"),
        utils.WithEffect("bounce"),
    )
    spinner.Start()
    missingTools := UpsentTools()
    spinner.Stop()

    // Show missing tools by category
    totalMissing := len(missingTools["system"]) + len(missingTools["security"]) + len(missingTools["discovery"])
    if totalMissing == 0 {
        //fmt.Printf("%s[+] %sAll tools are installed and ready!\n", bcolors.Green, bcolors.Endc)
        return
    }

    if len(missingTools["system"]) > 0 {
        fmt.Printf("%s[!] %s%sMissing system tools.%s\n\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["system"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["security"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing security tools.%s\n\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["security"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["discovery"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing project discovery tools.%s\n\n", bcolors.BrightRed, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["discovery"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("\n%s[?] %sLaunch setups to install missing tools? (y/n): ", bcolors.Yellow, bcolors.Endc)
    response, _ := reader.ReadString('\n')

    if strings.ToLower(strings.TrimSpace(response)) == "y" || strings.ToLower(strings.TrimSpace(response)) == "yes" {
        fmt.Println(); SetupsLauncher()
        return
    } else {
        fmt.Printf("%s[!] %sInstallation skipped. Some tools are missing ...\n", bcolors.BrightRed, bcolors.Endc)
    }
}

func isInstalled(tool string) bool {
    _, err := exec.LookPath(tool)
    if err != nil {
        // For some packages that don't have direct binaries, try dpkg
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

    // Check system tools
    for tool, pkg := range systemTools {
        if !isInstalled(tool) {
            time.Sleep(50 * time.Millisecond)
            missing["system"][tool] = pkg
        }
    }

    // Check security tools
    for tool, pkg := range securityTools {
        if !isInstalled(tool) {
            time.Sleep(50 * time.Millisecond)
            missing["security"][tool] = pkg
        }
    }

    // Check project discovery tools
    for tool, pkg := range projectDiscoveryTools {
        if !isInstalled(tool) {
            time.Sleep(50 * time.Millisecond)
            missing["discovery"][tool] = pkg
        }
    }
    return missing
}

func InstallTools(tools map[string]map[string]string) {
    // Install system tools first
    if len(tools["system"]) > 0 {
        fmt.Printf("\n%sInstalling system tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["system"] {
            fmt.Printf("\n%s[+]  %s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            subprocess.Popen("sudo apt install -y %s", pkg)
            time.Sleep(200 * time.Millisecond)
        }
    }

    // Install security tools
    if len(tools["security"]) > 0 {
        fmt.Printf("\n%sInstalling security tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["security"] {
            fmt.Printf("\n%s[+]  %s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            if strings.HasPrefix(pkg, "github.com") {
                // Go-based tool
                subprocess.Popen("go install %s", pkg)
            } else {
                // System package
                subprocess.Popen("sudo apt install -y %s", pkg)
            }
            time.Sleep(200 * time.Millisecond)
        }
    }

    // Install project discovery tools
    if len(tools["discovery"]) > 0 {
        fmt.Printf("\n%sInstalling project discovery tools%s.", bcolors.Bold, bcolors.Endc)
        for tool, pkg := range tools["discovery"] {
            fmt.Printf("\n%s[+]  %s- %sInstalling %s%-20s ...", bcolors.BrightGreen, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, tool)
            subprocess.Popen("go install %s", pkg)
            time.Sleep(200 * time.Millisecond)
        }
    }

    // Handle rockyou.txt.gz for Linux
    RokyPath  = "/usr/share/wordlists/rockyou.txt"
    if runtime.GOOS == "linux" {
        gzFilePath := RokyPath + ".gz"
        if _, err := os.Stat(RokyPath); os.IsNotExist(err) {
            if _, err := os.Stat(gzFilePath); os.IsNotExist(err) {
                return
            }
            command := fmt.Sprintf("gunzip -d -9 %s", gzFilePath)
            subprocess.Popen(command)
        }
    }
}

func Venv(PyEnvName string) {
    subprocess.Popen(`cd ~/.afr3/; python3 -m virtualenv %s --system-site-packages; echo -n "\n source ~/.afr3/%s/bin/activate" >> ~/.zshrc; source ~/.zshrc`, PyEnvName, PyEnvName)
}

func InstallFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Popen(cmd)
    }
}

func InstallGithubTools() {
    githubCommands := []string{
        `cd /root/.afr3/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1`,
        `python3 -m pip install -r /root/.afr3/africana-base/requirements.txt`,
    }

    for _, cmd := range githubCommands {
        subprocess.Popen(cmd)
    }
}

func UpdateAfricana() {
    fmt.Printf("\n%s[!] %sAfricana already installed. Updating it ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd /root/.afr3/africana-base; git pull .`)
    subprocess.Popen(`cd /root/.afr3/; git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1; cd ./africana-framework; make; mv africana-linux /usr/local/bin/africana; rm -rf ../africana-framework`)
    fmt.Printf("\n%s[*] %sAfricana succesfully updated ...\n", bcolors.Green, bcolors.Endc)
    return
}

func KaliSetups() {
    fmt.Printf("\n%s[>] %sInstalling africana in kali ...\n", bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()
    if _, err := os.Stat(ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `cd /etc/apt/trusted.gpg.d/; wget -qO - https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
            `cd /etc/apt/sources.list.d/; wget -qO - https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
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
    if _, err := os.Stat(ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc`,
            `echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref`,
            `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
            `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
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
    if _, err := os.Stat(ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            `curl -O https://blackarch.org/strap.sh`,
            `chmod +x strap.sh`,
            `./strap.sh`,
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm blackarch`,
            `pacman -S --noconfirm base-devel bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f WordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("\n%s[>] %sInstalling third party tools ...\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana fully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func AndroidSetups() {
    // Placeholder for Windows setup logic
}

func MacosSetups() {
}

func WindowsSetups() {
    // Placeholder for Windows setup logic
}

func AutoSetups() {
}

func Uninstaller() {
    fmt.Printf("%s\n[!] %sUninstalling africana ...\n", bcolors.BrightRed, bcolors.Endc)
    if _, err := os.Stat(ToolsDir); !os.IsNotExist(err) {
        subprocess.Popen(`rm -rf /root/.afr3/; rm -rf /usr/local/bin/africana`)
        fmt.Printf("\n%s[*] %sAfricana uninstalled. Goodbye! ...", bcolors.Green, bcolors.Endc)
        os.Exit(0)
    } else {
        fmt.Printf("%s[!] %sAfricana is not installed ...\n", bcolors.Green, bcolors.Endc)
    }
}
