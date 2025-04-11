package setups

import(
    "os"
    "fmt"
    "time"
    "os/exec"
    "bufio"
    "menus"
    "utils"
    "strings"
    "bcolors"
    "banners"
    "scriptures"
    "subprocess"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
    Input, Proxy, Distro, Function string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "distro": "",
    "module": "",
    "proxies": "",
    "function": "",
}

var (
    // System tools and packages
    systemTools = map[string]string{
        "jq":              "jq",
        "npm":             "npm",
        "tor":             "tor",
        "aha":             "aha",
        "ftp":             "ftp",
        "ncat":            "ncat",
        "gcc":             "gcc",
        "gawk":            "gawk",
        "tmux":            "tmux",
        "mdk4":            "mdk4",
        "mdk3":            "mdk3",
        "nmap":            "nmap",
        "playit":          "playit",
        "rlwrap":          "rlwrap",
        "squid":           "squid",
        "privoxy":         "privoxy",
        "iptables":        "iptables",
        "dnsmasq":         "dnsmasq",
        "openssh-client":  "openssh-client",
        "libpcap-dev":     "libpcap-dev",
        "openssh-server":  "openssh-server",
        "powershell":      "powershell",
        "golang-go":       "golang-go",
        "docker.io":       "docker.io",
        "python3":         "python3",
        "python3-pip":     "python3-pip",
        "build-essential": "build-essential",
        "libssl-dev":      "libssl-dev",
        "libffi-dev":      "libffi-dev",
        "python3-dev":     "python3-dev",
        "python3-venv":    "python3-venv",
        "python3-pycurl":  "python3-pycurl",
        "python3-geoip":   "python3-geoip",
        "python3-whois":   "python3-whois",
        "python3-requests": "python3-requests",
        "python3-scapy":   "python3-scapy",
        "libgeoip1t64":    "libgeoip1t64",
        "libgeoip-dev":    "libgeoip-dev",
        "aircrack-ng":     "aircrack-ng",
        "cowpatty":        "cowpatty",
        "dhcpd":           "dhcpd",
        "hostapd":         "hostapd",
        "lighttpd":        "lighttpd",
        "net-tools":       "net-tools",
        "macchanger":      "macchanger",
        "dsniff":          "dsniff",
        "openssl":         "openssl",
        "php-cgi":         "php-cgi",
        "xterm":           "xterm",
        "rfkill":          "rfkill",
        "unzip":           "unzip",
        "hydra":           "hydra",
        "wine32":          "wine32:i386",
    }

    // Security tools
    securityTools = map[string]string{
        "gophish":        "gophish",
        "wifipumpkin3":   "wifipumpkin3",
        "wifite":         "wifite",
        "airgeddon":      "airgeddon",
        "nuclei":         "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest",
        "nikto":          "nikto",
        "smbmap":         "smbmap",
        "dnsrecon":       "dnsrecon",
        "metasploit-framework": "metasploit-framework",
        "gobuster":       "gobuster",
        "feroxbuster":    "feroxbuster",
        "uniscan":        "uniscan",
        "sqlmap":         "sqlmap",
        "commix":         "commix",
        "dnsenum":        "dnsenum",
        "sslscan":        "sslscan",
        "whatweb":        "whatweb",
        "wapiti":         "wapiti",
        "xsser":          "xsser",
        "netexec":        "netexec",
    }

    // Project discovery tools
    projectDiscoveryTools = map[string]string{
        "subfinder":    "github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest",
        "dnsx":         "github.com/projectdiscovery/dnsx/cmd/dnsx@latest",
        "naabu":        "github.com/projectdiscovery/naabu/v2/cmd/naabu@latest",
        "httpx":        "github.com/projectdiscovery/httpx/cmd/httpx@latest",
        "katana":       "github.com/projectdiscovery/katana/cmd/katana@latest",
        "uncover":      "github.com/projectdiscovery/uncover/cmd/uncover@latest",
        "mapcidr":      "github.com/projectdiscovery/mapcidr/cmd/mapcidr@latest",
        "asnmap":       "github.com/projectdiscovery/asnmap/cmd/asnmap@latest",
        "tlsx":         "github.com/projectdiscovery/tlsx/cmd/tlsx@latest",
        "interactsh":   "github.com/projectdiscovery/interactsh/cmd/interactsh-client@latest",
        "notify":       "github.com/projectdiscovery/notify/cmd/notify@latest",
        "gau":          "github.com/lc/gau/v2/cmd/gau@latest",
        "waybackurls":  "github.com/tomnomnom/waybackurls@latest",
        "dalfox":       "github.com/hahwul/dalfox/v2@latest",
        "hakrawler":    "github.com/hakluke/hakrawler@latest",
        "assetfinder":  "github.com/tomnomnom/assetfinder@latest",
        "anew":         "github.com/tomnomnom/anew@latest",
        "gf":           "github.com/tomnomnom/gf@latest",
    }
)

func SetupsLauncher() {
    for {
        fmt.Printf("%s%safr3%s packages(%score/setups_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Red, Function, bcolors.Endc, bcolors.Green, bcolors.Endc)
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
    "info":             menus.HelpInfoSetups,

    "m":                menus.MenuOne,
    "menu":             menus.MenuOne,

    "option":           menus.SetupsOptions,
    "options":          menus.SetupsOptions,
    "show option":      menus.SetupsOptions,
    "show options":     menus.SetupsOptions,

    "show all":         menus.ListSetupsFunction,
    "list all":         menus.ListSetupsFunction,

    "func":             menus.ListSetupsFunction,
    "funcs":            menus.ListSetupsFunction,
    "functions":        menus.ListSetupsFunction,
    "show func":        menus.ListSetupsFunction,
    "list funcs":       menus.ListSetupsFunction,
    "show funcs":       menus.ListSetupsFunction,
    "show function":    menus.ListSetupsFunction,
    "list function":    menus.ListSetupsFunction,
    "list functions":   menus.ListSetupsFunction,
    "show functions":   menus.ListSetupsFunction,

    "module":           menus.ListSetupsFunction,
    "modules":          menus.ListSetupsFunction,
    "list module":      menus.ListSetupsFunction,
    "show module":      menus.ListSetupsFunction,
    "list modules":     menus.ListSetupsFunction,
    "show modules":     menus.ListSetupsFunction,

    "distro":           menus.ListSetupsDistros,
    "distros":          menus.ListSetupsDistros,
    "list distro":      menus.ListSetupsDistros,
    "list distros":     menus.ListSetupsDistros,
    "show distro":      menus.ListSetupsDistros,
    "show distros":     menus.ListSetupsDistros,

    "1":                func() {autoExecuteFunc("arch", "install")},
    "run 1":            func() {autoExecuteFunc("arch", "install")},
    "use 1":            func() {autoExecuteFunc("arch", "install")},
    "exec 1":           func() {autoExecuteFunc("arch", "install")},
    "start 1":          func() {autoExecuteFunc("arch", "install")},
    "launch 1":         func() {autoExecuteFunc("arch", "install")},
    "exploit 1":        func() {autoExecuteFunc("arch", "install")},
    "execute 1":        func() {autoExecuteFunc("arch", "install")},
    "run arch":         func() {autoExecuteFunc("arch", "install")},
    "use arch":         func() {autoExecuteFunc("arch", "install")},
    "exec arch":        func() {autoExecuteFunc("arch", "install")},
    "start arch":       func() {autoExecuteFunc("arch", "install")},
    "launch arch":      func() {autoExecuteFunc("arch", "install")},
    "exploit arch":     func() {autoExecuteFunc("arch", "install")},
    "execute arch":     func() {autoExecuteFunc("arch", "install")},

    "? 1":              menus.HelpInfoArch,
    "info 1":           menus.HelpInfoArch,
    "help 1":           menus.HelpInfoArch,
    "arch":             menus.HelpInfoArch,
    "info arch":        menus.HelpInfoArch,
    "help arch":        menus.HelpInfoArch,

    "2":                func() {autoExecuteFunc("kali", "install")},
    "run 2":            func() {autoExecuteFunc("kali", "install")},
    "use 2":            func() {autoExecuteFunc("kali", "install")},
    "exec 2":           func() {autoExecuteFunc("kali", "install")},
    "start 2":          func() {autoExecuteFunc("kali", "install")},
    "launch 2":         func() {autoExecuteFunc("kali", "install")},
    "exploit 2":        func() {autoExecuteFunc("kali", "install")},
    "execute 2":        func() {autoExecuteFunc("kali", "install")},
    "run kali":         func() {autoExecuteFunc("kali", "install")},
    "use kali":         func() {autoExecuteFunc("kali", "install")},
    "exec kali":        func() {autoExecuteFunc("kali", "install")},
    "start kali":       func() {autoExecuteFunc("kali", "install")},
    "launch kali":      func() {autoExecuteFunc("kali", "install")},
    "exploit kali":     func() {autoExecuteFunc("kali", "install")},
    "execute kali":     func() {autoExecuteFunc("kali", "install")},

    "? 2":              menus.HelpInfoKali,
    "info 2":           menus.HelpInfoKali,
    "help 2":           menus.HelpInfoKali,
    "kali":             menus.HelpInfoKali,
    "info kali":        menus.HelpInfoKali,
    "help kali":        menus.HelpInfoKali,

    "3":                func() {autoExecuteFunc("macos", "install")},
    "run 3":            func() {autoExecuteFunc("macos", "install")},
    "use 3":            func() {autoExecuteFunc("macos", "install")},
    "exec 3":           func() {autoExecuteFunc("macos", "install")},
    "start 3":          func() {autoExecuteFunc("macos", "install")},
    "launch 3":         func() {autoExecuteFunc("macos", "install")},
    "exploit 3":        func() {autoExecuteFunc("macos", "install")},
    "execute 3":        func() {autoExecuteFunc("macos", "install")},
    "run macos":        func() {autoExecuteFunc("macos", "install")},
    "use macos":        func() {autoExecuteFunc("macos", "install")},
    "exec macos":       func() {autoExecuteFunc("macos", "install")},
    "start macos":      func() {autoExecuteFunc("macos", "install")},
    "launch macos":     func() {autoExecuteFunc("macos", "install")},
    "exploit macos":    func() {autoExecuteFunc("macos", "install")},
    "execute macos":    func() {autoExecuteFunc("macos", "install")},

    "? 3":              menus.HelpInfoMacos,
    "info 3":           menus.HelpInfoMacos,
    "help 3":           menus.HelpInfoMacos,
    "macos":         menus.HelpInfoMacos,
    "info macos":    menus.HelpInfoMacos,
    "help macos":    menus.HelpInfoMacos,

    "4":                func() {autoExecuteFunc("ubuntu", "install")},
    "run 4":            func() {autoExecuteFunc("ubuntu", "install")},
    "use 4":            func() {autoExecuteFunc("ubuntu", "install")},
    "exec 4":           func() {autoExecuteFunc("ubuntu", "install")},
    "start 4":          func() {autoExecuteFunc("ubuntu", "install")},
    "launch 4":         func() {autoExecuteFunc("ubuntu", "install")},
    "exploit 4":        func() {autoExecuteFunc("ubuntu", "install")},
    "execute 4":        func() {autoExecuteFunc("ubuntu", "install")},
    "run ubuntu":       func() {autoExecuteFunc("ubuntu", "install")},
    "use ubuntu":       func() {autoExecuteFunc("ubuntu", "install")},
    "exec ubuntu":      func() {autoExecuteFunc("ubuntu", "install")},
    "start ubuntu":     func() {autoExecuteFunc("ubuntu", "install")},
    "launch ubuntu":    func() {autoExecuteFunc("ubuntu", "install")},
    "exploit ubuntu":   func() {autoExecuteFunc("ubuntu", "install")},
    "execute ubuntu":   func() {autoExecuteFunc("ubuntu", "install")},

    "? 4":              menus.HelpInfoUbuntu,
    "info 4":           menus.HelpInfoUbuntu,
    "help 4":           menus.HelpInfoUbuntu,
    "ubuntu":           menus.HelpInfoUbuntu,
    "info ubuntu":      menus.HelpInfoUbuntu,
    "help ubuntu":      menus.HelpInfoUbuntu,

    "5":                func() {autoExecuteFunc("android", "install")},
    "run 5":            func() {autoExecuteFunc("android", "install")},
    "use 5":            func() {autoExecuteFunc("android", "install")},
    "exec 5":           func() {autoExecuteFunc("android", "install")},
    "start 5":          func() {autoExecuteFunc("android", "install")},
    "launch 5":         func() {autoExecuteFunc("android", "install")},
    "exploit 5":        func() {autoExecuteFunc("android", "install")},
    "execute 5":        func() {autoExecuteFunc("android", "install")},
    "run android":      func() {autoExecuteFunc("android", "install")},
    "use android":      func() {autoExecuteFunc("android", "install")},
    "exec android":     func() {autoExecuteFunc("android", "install")},
    "start android":    func() {autoExecuteFunc("android", "install")},
    "launch android":   func() {autoExecuteFunc("android", "install")},
    "exploit android":  func() {autoExecuteFunc("android", "install")},
    "execute android":  func() {autoExecuteFunc("android", "install")},

    "? 5":              menus.HelpInfoAndroid,
    "info 5":           menus.HelpInfoAndroid,
    "help 5":           menus.HelpInfoAndroid,
    "android":          menus.HelpInfoAndroid,
    "info android":     menus.HelpInfoAndroid,
    "help android":     menus.HelpInfoAndroid,

    "6":                func() {autoExecuteFunc("windows", "install")},
    "run 6":            func() {autoExecuteFunc("windows", "install")},
    "use 6":            func() {autoExecuteFunc("windows", "install")},
    "exec 6":           func() {autoExecuteFunc("windows", "install")},
    "start 6":          func() {autoExecuteFunc("windows", "install")},
    "launch 6":         func() {autoExecuteFunc("windows", "install")},
    "exploit 6":        func() {autoExecuteFunc("windows", "install")},
    "execute 6":        func() {autoExecuteFunc("windows", "install")},
    "run windows":      func() {autoExecuteFunc("windows", "install")},
    "use windows":      func() {autoExecuteFunc("windows", "install")},
    "exec windows":     func() {autoExecuteFunc("windows", "install")},
    "start windows":    func() {autoExecuteFunc("windows", "install")},
    "launch windows":   func() {autoExecuteFunc("windows", "install")},
    "exploit windows":  func() {autoExecuteFunc("windows", "install")},
    "execute windows":  func() {autoExecuteFunc("windows", "install")},

    "? 6":              menus.HelpInfoWindows,
    "info 6":           menus.HelpInfoWindows,
    "help 6":           menus.HelpInfoWindows,
    "windows":           menus.HelpInfoWindows,
    "info windows":      menus.HelpInfoWindows,
    "help windows":      menus.HelpInfoWindows,

    "7":                func() {autoExecuteFunc(Distro, "update")},
    "run 7":            func() {autoExecuteFunc(Distro, "update")},
    "use 7":            func() {autoExecuteFunc(Distro, "update")},
    "exec 7":           func() {autoExecuteFunc(Distro, "update")},
    "start 7":          func() {autoExecuteFunc(Distro, "update")},
    "launch 7":         func() {autoExecuteFunc(Distro, "update")},
    "exploit 7":        func() {autoExecuteFunc(Distro, "update")},
    "execute 7":        func() {autoExecuteFunc(Distro, "update")},
    "run update":     func() {autoExecuteFunc(Distro, "update")},
    "use update":     func() {autoExecuteFunc(Distro, "update")},
    "exec update":    func() {autoExecuteFunc(Distro, "update")},
    "start update":   func() {autoExecuteFunc(Distro, "update")},
    "launch update":  func() {autoExecuteFunc(Distro, "update")},
    "exploit update": func() {autoExecuteFunc(Distro, "update")},
    "execute update": func() {autoExecuteFunc(Distro, "update")},

    "? 7":              menus.HelpInfoUpdate,
    "info 7":           menus.HelpInfoUpdate,
    "help 7":           menus.HelpInfoUpdate,
    "update":             menus.HelpInfoUpdate,
    "info update":        menus.HelpInfoUpdate,
    "help update":        menus.HelpInfoUpdate,

    "8":                func() {autoExecuteFunc(Distro, "reapair")},
    "run 8":            func() {autoExecuteFunc(Distro, "reapair")},
    "use 8":            func() {autoExecuteFunc(Distro, "reapair")},
    "exec 8":           func() {autoExecuteFunc(Distro, "reapair")},
    "start 8":          func() {autoExecuteFunc(Distro, "reapair")},
    "launch 8":         func() {autoExecuteFunc(Distro, "reapair")},
    "exploit 8":        func() {autoExecuteFunc(Distro, "reapair")},
    "execute 8":        func() {autoExecuteFunc(Distro, "reapair")},
    "run reapair":     func() {autoExecuteFunc(Distro, "reapair")},
    "use reapair":     func() {autoExecuteFunc(Distro, "reapair")},
    "exec reapair":    func() {autoExecuteFunc(Distro, "reapair")},
    "start reapair":   func() {autoExecuteFunc(Distro, "reapair")},
    "launch reapair":  func() {autoExecuteFunc(Distro, "reapair")},
    "exploit reapair": func() {autoExecuteFunc(Distro, "reapair")},
    "execute reapair": func() {autoExecuteFunc(Distro, "reapair")},

     "? 8":             menus.HelpInfoRepair,
    "info 8":           menus.HelpInfoRepair,
    "help 8":           menus.HelpInfoRepair,
    "reapair":         menus.HelpInfoRepair,
    "info reapair":    menus.HelpInfoRepair,
    "help reapair":    menus.HelpInfoRepair,

    "9":               func() {autoExecuteFunc(Distro, "uninstall")},
    "run 9":           func() {autoExecuteFunc(Distro, "uninstall")},
    "use 9":           func() {autoExecuteFunc(Distro, "uninstall")},
    "exec 9":          func() {autoExecuteFunc(Distro, "uninstall")},
    "start 9":         func() {autoExecuteFunc(Distro, "uninstall")},
    "launch 9":        func() {autoExecuteFunc(Distro, "uninstall")},
    "exploit 9":       func() {autoExecuteFunc(Distro, "uninstall")},
    "execute 9":       func() {autoExecuteFunc(Distro, "uninstall")},
    "run uninstall":     func() {autoExecuteFunc(Distro, "uninstall")},
    "use uninstall":     func() {autoExecuteFunc(Distro, "uninstall")},
    "exec uninstall":    func() {autoExecuteFunc(Distro, "uninstall")},
    "start uninstall":   func() {autoExecuteFunc(Distro, "uninstall")},
    "launch uninstall":  func() {autoExecuteFunc(Distro, "uninstall")},
    "exploit uninstall": func() {autoExecuteFunc(Distro, "uninstall")},
    "execute uninstall": func() {autoExecuteFunc(Distro, "uninstall")},

    "? 9":              menus.HelpInfoUninstaller,
    "info 9":           menus.HelpInfoUninstaller,
    "help 9":           menus.HelpInfoUninstaller,
    "uninstall":          menus.HelpInfoUninstaller,
    "info uninstall":     menus.HelpInfoUninstaller,
    "help uninstall":     menus.HelpInfoUninstaller,

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

        "distro": &Distro,
        "proxies": &Proxy,
        "func":  &Function,
        "module": &Function,
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

        "distro": &Distro,
        "proxies": &Proxy,
        "func": &Function,
        "module": &Function,
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
        fmt.Printf("\n%s[!] %sMissing required parameter FUNCTION. Use %s'show functions' %sfor details.\n", bcolors.Red, bcolors.Endc, bcolors.Green, bcolors.Endc)
        return
    }
    if Distro == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'show distros' %sfor details.\n", bcolors.Red, bcolors.Endc, bcolors.Green, bcolors.Endc)
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
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    commands := map[string]func() {
        "install": func() {Installer(Distro)},
         "update": func() {UpdateAfricana()},
         "repair": func() {UpdateAfricana()},
      "uninstall": func() {Uninstaller()},
           //"auto": ,
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid FUNCTION %s. Use %s'help' %sfor available Functions.\n", bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, bcolors.Endc)
    }
}

func Installer(Distro string) {
    switch Distro {
    case "1", "arch":
        ArchSetups()
        return
    case "2", "kali":
        KaliSetups()
        return
    case "3", "macos":
        MacosSetups()
        return
    case "4", "ubuntu":
        UbuntuSetups()
        return
    case "5", "android":
        AndroidSetups()
        return
    case "6", "windows":
        WindowsSetups()
        return
    case "7", "update":
        UpdateAfricana()
        return
    case "8", "repair":
        UpdateAfricana()
        return
    case "9", "uninstall":
        Uninstaller()
        return
    default:
        fmt.Printf("%s[!] Error: %sInvalid DISTRO %s. Use %s'help' %sfor commands.\n", bcolors.Red, bcolors.Endc, Distro, bcolors.Green, bcolors.Endc)
    }
}

func CheckTools() {
    // Create spinner with custom options
    spinner := utils.New(
        utils.WithStyle("classic"),
        utils.WithText("[+] starting the africana framework console ..."),
        utils.WithEffect("bounce"),
    )
    spinner.Start()
    time.Sleep(300 * time.Millisecond)
    missingTools := UpsentTools()
    time.Sleep(500 * time.Millisecond)
    spinner.Stop()

    // Show missing tools by category
    totalMissing := len(missingTools["system"]) + len(missingTools["security"]) + len(missingTools["discovery"])
    if totalMissing == 0 {
        //fmt.Printf("%s[+] %sAll tools are installed and ready!\n", bcolors.Green, bcolors.Endc)
        return
    }

    if len(missingTools["system"]) > 0 {
        fmt.Printf("%s[!] %s%sMissing system tools.%s\n\n", bcolors.Red, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["system"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["security"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing security tools.%s\n\n", bcolors.Red, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["security"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    if len(missingTools["discovery"]) > 0 {
        fmt.Printf("\n%s[!] %s%sMissing project discovery tools.%s\n\n", bcolors.Red, bcolors.Endc, bcolors.Bold, bcolors.Endc)
        for tool := range missingTools["discovery"] {
            fmt.Printf("  - %s%s ...\n", bcolors.Endc, tool)
            time.Sleep(90 * time.Millisecond)
        }
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("\n%s[?] %sLaunch setups to install missing tools? (y/n): ", bcolors.Yellow, bcolors.Endc)
    response, _ := reader.ReadString('\n')

    if strings.ToLower(strings.TrimSpace(response)) == "y" || strings.ToLower(strings.TrimSpace(response)) == "yes" {
        SetupsLauncher()
        return
    } else {
        fmt.Printf("%s[!] %sInstallation skipped. Some tools are missing ...\n", bcolors.Red, bcolors.Endc)
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

func installTools(tools map[string]map[string]string) {
    // Install system tools first
    if len(tools["system"]) > 0 {
        fmt.Printf("\n🔧 Installing system tools:")
        for tool, pkg := range tools["system"] {
            fmt.Printf("  - Installing %-20s", tool)
            subprocess.Popen("sudo apt install -y %s", pkg)
            time.Sleep(200 * time.Millisecond)
        }
    }

    // Install security tools
    if len(tools["security"]) > 0 {
        fmt.Printf("\n🔧 Installing security tools:")
        for tool, pkg := range tools["security"] {
            fmt.Printf("  - Installing %-20s", tool)
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
        fmt.Printf("\n🔧 Installing project discovery tools:")
        for tool, pkg := range tools["discovery"] {
            fmt.Printf("  - Installing %-20s", tool)
            subprocess.Popen("go install %s", pkg)
            time.Sleep(200 * time.Millisecond)
        }
    }
}

func InstallFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Popen(cmd)
    }
}

func InstallGithubTools() {
    githubCommands := []string{
        //cd /root/.afr3/; python3 -m virtualenv africana-venv --system-site-packages; echo -n "\n source ~/.afr3/africana-venv/bin/activate" >> ~/.zshrc
        `cd /root/.afr3/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1`,
        `python3 -m pip install -r /root/.afr3/africana-base/requirements.txt`,
    }

    for _, cmd := range githubCommands {
        subprocess.Popen(cmd)
    }
}

func UpdateAfricana() {
    fmt.Printf("\n%s[!] %safricana already installed. Updating it ...\n", bcolors.Green, bcolors.Endc)
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
            `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
            `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
            `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f WordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        installTools(missingTools)
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
            `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f WordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("\n%s[>] %sInstalling third party tools\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana succesfully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func ArchSetups() {
    fmt.Printf("\n%s[>] %sInstalling africana in arch\n", bcolors.Green, bcolors.Endc)
    if _, err := os.Stat(ToolsDir); os.IsNotExist(err) {
        foundationCommands := []string{
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm zsh git curl wget go`,
        }
        InstallFoundationTools(foundationCommands)
        BlackArchSetups()
        fmt.Printf("\n%s[>] %sInstalling third party tools\n", bcolors.Green, bcolors.Endc)
        InstallGithubTools()
        fmt.Printf("\n%s[*] %sAfricana succesfully installed ...\n", bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func BlackArchSetups() {
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
    fmt.Printf("%s\n[!] %sUninstalling africana ...\n", bcolors.Red, bcolors.Endc)
    if _, err := os.Stat(ToolsDir); !os.IsNotExist(err) {
        subprocess.Popen(`rm -rf /root/.afr3/; rm -rf /usr/local/bin/africana`)
        fmt.Printf("\n%s[*] %sAfricana uninstalled. Goodbye! ...", bcolors.Green, bcolors.Endc)
        os.Exit(0)
    } else {
        fmt.Printf("%s[!] %sAfricana is not installed ...\n", bcolors.Green, bcolors.Endc)
    }
}
