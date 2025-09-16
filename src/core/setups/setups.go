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
    "path/filepath"
)

var (
    Function string
)

type stringMatcher struct {
    names  []string
    action func()
}

var linuxTaskMap = map[string]func(){
    "kali":    KaliSetups,
    "debian":  KaliSetups,
    "ubuntu":  UbuntuSetups,
    "fedora":  func() { fmt.Printf("%s%s[+] %sFedora distro detected.\n", bcolors.BrightGreen, bcolors.Endc) },
    "arch":    ArchSetups,
    "alpine":  func() { fmt.Printf("%s%s[+] %sAlpine distro detected.\n", bcolors.BrightGreen, bcolors.Endc) },
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
        //"github-desktop":               "github-desktop",
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
        //"playit":                       "playit",
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
        "python3-tk":                   "python3-tk",
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
        "interactsh-client":            "github.com/projectdiscovery/interactsh/cmd/interactsh-client@latest",
        "katana":                       "github.com/projectdiscovery/katana/cmd/katana@latest",
        "mapcidr":                      "github.com/projectdiscovery/mapcidr/cmd/mapcidr@latest",
        "naabu":                        "github.com/projectdiscovery/naabu/v2/cmd/naabu@latest",
        "notify":                       "github.com/projectdiscovery/notify/cmd/notify@latest",
        "subfinder":                    "github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest",
        "tlsx":                         "github.com/projectdiscovery/tlsx/cmd/tlsx@latest",
        "uncover":                      "github.com/projectdiscovery/uncover/cmd/uncover@latest",
        "waybackurls":                  "github.com/tomnomnom/waybackurls@latest",
        "nuclei":                       "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest",
        "go-winres":                    "github.com/tc-hib/go-winres@latest",
    }
)

func getPackageName(tool, pkg string) string {
    if utils.DetectAndroid() {
        androidPackages := map[string]string{
            "golang-go":                    "golang",
            "python3":                      "python",
            "python3-pip":                  "python-pip",
            "python3-dev":                  "python",
            "python3-venv":                 "python",
            "git":                          "git",
            "curl":                         "curl",
            "wget":                         "wget",
            "nmap":                         "nmap",
            "aircrack-ng":                  "aircrack-ng",
            "hydra":                        "hydra",
            "sqlmap":                       "sqlmap",
            "nikto":                        "nikto",
            "gobuster":                     "gobuster",
            "net-tools":                    "net-tools",
            "dnsutils":                     "dnsutils",
            "build-essential":              "clang",
            "gcc":                          "clang",
            "make":                         "make",
            "libpcap-dev":                  "libpcap",
            "libssl-dev":                   "openssl",
            "zlib1g-dev":                   "zlib",
        }
        if androidPkg, exists := androidPackages[pkg]; exists {
            return androidPkg
        }
        return tool
    }

    if utils.IsArchLinux() {
        archPackages := map[string]string{

            "aha":                          "aha",
            "aircrack-ng":                  "aircrack-ng",
            "build-essential":              "base-devel",
            "checkinstall":                 "checkinstall",
            "cowpatty":                     "cowpatty",
            "dhcpd":                        "dhcp",
            "dnsmasq":                      "dnsmasq",
            "docker.io":                    "docker",
            "dsniff":                       "dsniff",
            "ftp":                          "inetutils",
            "gawk":                         "gawk",
            "gcc":                          "gcc",
            "gdb":                          "gdb",
            //"github-desktop":               "github-desktop",
            "golang-go":                    "go",
            "gstreamer1.0-libav":           "gst-libav",
            "gunzip":                       "gzip",
            "hostapd":                      "hostapd",
            "hydra":                        "hydra",
            "iptables":                     "iptables",
            "jq":                           "jq",
            "lcov":                         "lcov",
            "libbz2-dev":                   "bzip2",
            "libc6-dev":                    "glibc",
            "libffi-dev":                   "libffi",
            "libgdbm-compat-dev":           "gdbm",
            "libgdbm-dev":                  "gdbm",
            "libgeoip-dev":                 "geoip",
            "libgeoip1t64":                 "geoip",
            "liblzma-dev":                  "xz",
            "libncurses-dev":               "ncurses",
            "libpcap-dev":                  "libpcap",
            "libreadline-dev":              "readline",
            "libsqlite3-dev":               "sqlite",
            "libssl-dev":                   "openssl",
            "lighttpd":                     "lighttpd",
            "macchanger":                   "macchanger",
            "mdk3":                         "mdk3",
            "mdk4":                         "mdk4",
            "mousepad":                     "mousepad",
            "ncat":                         "nmap",
            "net-tools":                    "net-tools",
            "nmap":                         "nmap",
            "npm":                          "npm",
            "openssh-client":               "openssh",
            "openssh-server":               "openssh",
            "openssl":                      "openssl",
            "php-cgi":                      "php",
            "pkg-config":                   "pkgconf",
            //"playit":                       "playit",
            "powershell":                   "powershell",
            "privoxy":                      "privoxy",
            "python3":                      "python",
            "python3-dev":                  "python",
            "python3-geoip":                "python-geoip",
            "python3-pip":                  "python-pip",
            "python3-pycurl":               "python-pycurl",
            "python3-requests":             "python-requests",
            "python3-scapy":                "python-scapy",
            "python3-venv":                 "python",
            "python3-whois":                "python-whois",
            "python-dev-is-python3":        "python",
            "python3-tk":                   "python3-tk",
            "rfkill":                       "util-linux",
            "rlwrap":                       "rlwrap",
            "squid":                        "squid",
            "tk-dev":                       "tk",
            "tmux":                         "tmux",
            "tor":                          "tor",
            "uuid-dev":                     "util-linux",
            "wine32":                       "wine",
            "xterm":                        "xterm",
            "zlib1g-dev":                   "zlib",
            "zsh":                          "zsh",
                        /*Security tools*/
            "airgeddon":                    "airgeddon",
            "commix":                       "commix",
            "dnsenum":                      "dnsenum",
            "dnsrecon":                     "dnsrecon",
            "findomain":                    "findomain",
            "feroxbuster":                  "feroxbuster",
            "gobuster":                     "gobuster",
            "gophish":                      "gophish",
            "massdns":                      "massdns",
            "metasploit-framework":         "metasploit",
            "netexec":                      "netexec",
            "nikto":                        "nikto",
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
        
        if archPkg, exists := archPackages[pkg]; exists {
            return archPkg
        }
        return tool
    }
    return pkg
}

func SetupsLauncher() {
    for {
        fmt.Printf("%s%s%safr%s%s packages(%s%score/setups_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underline, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutorials},
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
      "iface": &utils.IFace,
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
      "interface": &utils.IFace,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.Gateway,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "output": &utils.SetupsLogs,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "pyenvname": &utils.PyEnvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "outputlog": &utils.SetupsLogs,
      "obfuscator": &utils.Obfuscator,
      "outputlogs": &utils.SetupsLogs,
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
        fmt.Printf("%s%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

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
            fmt.Println()
        }
        fmt.Println()
        return
    }

    fmt.Printf("%s%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)
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
      "ssid": &utils.Ssid,
      "iface": &utils.IFace,
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
      "interface": &utils.IFace,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.Gateway,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "output": &utils.SetupsLogs,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "pyenvname": &utils.PyEnvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "outputlog": &utils.SetupsLogs,
      "obfuscator": &utils.Obfuscator,
      "outputlogs": &utils.SetupsLogs,
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
        fmt.Printf("%s%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

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
            fmt.Println()
        }
        fmt.Println()
        return
    }

    fmt.Printf("%s%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

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
        fmt.Println()
    }
}

func executeFunction() {
    if Function == ""{
        fmt.Printf("\n%s%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    SetupsFunction(Function, utils.Distro)
}

func SetupsFunction(Function, Distro string, args ...interface{}) {
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
        }, true, true)

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
        fmt.Printf("%s%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListSetupsFunction()
        return
    }

    lowerInput := strings.ToLower(Function)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("\n%s%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, Function)
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

    fmt.Printf("%s%s[!] Error: %sInvalid DISTRO %s. Use %s'help' %sfor commands.\n", bcolors.Red, bcolors.Endc, utils.Distro, bcolors.Green, bcolors.Endc)
}

func CheckMissing() {
    spinner := utils.New(utils.WithStyle("circle"), utils.WithEffect("bounce"), utils.WithClearOnStop(true), utils.WithText("[+] Setup is Checking all missing tools ..."))
    spinner.Start()
    missingTools := UpsentTools()
    spinner.Stop()

    if len(missingTools["system"]) + len(missingTools["security"]) + len(missingTools["discovery"]) == 0 {
        fmt.Printf("%s%s[+] %sAll tools are installed and ready!\n", bcolors.Green, bcolors.Endc)
        return
    }

    printMissingTools(missingTools)
    userChoice()
}

func CheckTools() {
    spinner := utils.New(utils.WithStyle("classic"), utils.WithEffect("bounce"))
    spinner.Start()
    missingTools := UpsentTools()
    spinner.Stop()

    if len(missingTools["system"]) + len(missingTools["security"]) + len(missingTools["discovery"]) == 0 {
        fmt.Printf("%s%s[+] %sAll tools are installed and ready!\n", bcolors.Green, bcolors.Endc)
        return
    }

    printMissingTools(missingTools)
    userChoice()
}

func printMissingTools(missingTools map[string]map[string]string) {
    categories := []struct{
        key, name string
    }{
        {"system", "system"},
        {"security", "security"}, 
        {"discovery", "project discovery"},
    }

    for _, cat := range categories {
        if len(missingTools[cat.key]) > 0 {
            fmt.Printf("\n%s%s[!] %s%sMissing %s tools.%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Bold, cat.name, bcolors.Endc)
            for tool := range missingTools[cat.key] {
                fmt.Printf("  %s- %s%s...\n", bcolors.Bold, bcolors.Endc, tool)
                time.Sleep(90 * time.Millisecond)
            }
        }
    }
}

func userChoice() {
    fmt.Printf("\n%s%s[?] %sInstall missing tools? (y/n): ", bcolors.Bold, bcolors.Green, bcolors.Endc)
    utils.Scanner.Scan()
    input := strings.ToLower(strings.TrimSpace(utils.Scanner.Text()))

    switch input {
    case "y", "yes":
        AutoSetups()
    case "n", "q", "no", "exit", "quit":
        fmt.Printf("%s%s[!] %sInstallation skipped. Some tools are missing ...\n",  bcolors.Bold, bcolors.BrightRed, bcolors.Endc)
    default:
        fmt.Printf("%s%s[!] %sChoices are (y|n|yes|no)", bcolors.Bold, bcolors.Red, bcolors.Endc)
        userChoice()
    }
}

func isInstalled(tool string) bool {
    if _, err := exec.LookPath(tool); err == nil {
        return true
    }
    if utils.IsArchLinux() {
        cmd := exec.Command("pacman", "-Q", tool)
        return cmd.Run() == nil
    }
    cmd := exec.Command("dpkg", "-s", tool)
    return cmd.Run() == nil
}

func UpsentTools() map[string]map[string]string {
    missing := map[string]map[string]string{
        "system":    {},
        "security":  {},
        "discovery": {},
    }

    toolSets := []struct{
        category string
        tools    map[string]string
    }{
        {"system", systemTools},
        {"security", securityTools},
        {"discovery", projectDiscoveryTools},
    }

    for _, set := range toolSets {
        for tool, pkg := range set.tools {
            if !isInstalled(tool) {
                time.Sleep(45 * time.Millisecond)
                missing[set.category][tool] = pkg
            }
        }
    }

    return missing
}

func InstallTools(tools map[string]map[string]string) {
    categories := []struct{
        key, name string
    }{
        {"system", "system"},
        {"security", "security"},
        {"discovery", "project discovery"},
    }

    for _, cat := range categories {
        if len(tools[cat.key]) > 0 {
            fmt.Printf("\n%sInstalling %s tools%s.", bcolors.Bold, cat.name, bcolors.Endc)
            for tool, pkg := range tools[cat.key] {
                fmt.Printf("\n%s%s[+]  %s%s- %sInstalling %s%-20s ...", bcolors.Bold, bcolors.Green, bcolors.Bold, bcolors.Endc, bcolors.Blue, bcolors.Endc, tool)
                actualPkg := getPackageName(tool, pkg)

                if cat.key == "security" && strings.HasPrefix(pkg, "@latest") {
                    subprocess.Run("go install %s", pkg)
                } else if cat.key == "discovery" {
                    subprocess.Run("go install %s", pkg)
                } else {
                    if utils.IsArchLinux() {
                        subprocess.Run("sudo pacman -S --noconfirm %s", actualPkg)
                    } else {
                        subprocess.Run("sudo apt install -y %s", actualPkg)
                    }
                }
                time.Sleep(180 * time.Millisecond)
            }
        }
    }
}

func SetupGoEnvironment(PyEnvName string) {

    os.Setenv("GOPATH", utils.GoPath)
    os.Setenv("PATH", os.Getenv("PATH") + ":/usr/local/go/bin:" + filepath.Join(utils.GoPath, "bin"))

    os.Setenv("VIRTUAL_ENV", utils.PythonVenv)
    os.Setenv("PATH", filepath.Join(utils.PythonVenv, "bin") + ":" + os.Getenv("PATH"))

    goDirs := []string{
        utils.GoPath,
        filepath.Join(utils.GoPath, "bin"),
        filepath.Join(utils.GoPath, "pkg"),
        filepath.Join(utils.GoPath, "src"),
    }

    for _, dir := range goDirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            fmt.Printf("%s%s[!] %sFailed to create directory %s: %v%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, dir, err, bcolors.Endc)
        }
    }

    if err := os.MkdirAll(utils.PythonVenv, 0755); err != nil {
        fmt.Printf("%s%s[!] %sFailed to create Python venv directory %s: %v%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, utils.PythonVenv, err, bcolors.Endc)
    }

    shellProfiles := []string{
        filepath.Join(utils.HomeDir, ".bashrc"),
        filepath.Join(utils.HomeDir, ".zshrc"),
    }

    envSetup := fmt.Sprintf(`
# Go environment
export GOPATH=%s
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Python AFR environment
#export AFR_VENV=%s
#export PATH=$AFR_VENV/bin:$PATH
#source $AFR_VENV/bin/activate 2>/dev/null
`, utils.GoPath, utils.PythonVenv)

    for _, profile := range shellProfiles {
        if err := appendToShellProfile(profile, envSetup); err != nil {
            fmt.Printf("%s%s[!] %sFailed to update %s: %v%s\n", bcolors.Yellow, bcolors.Endc, profile, err, bcolors.Endc)
        }
    }
}

func appendToShellProfile(profilePath, content string) error {
    if _, err := os.Stat(profilePath); os.IsNotExist(err) {
        if err := os.WriteFile(profilePath, []byte(content), 0644); err != nil {
            return err
        }
        return nil
    }

    f, err := os.OpenFile(profilePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(content)
    return err
}

func InstallFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Run(cmd)
    }
}

func InstallGithubTools() {
    subprocess.Run("cd %s; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1", utils.BaseDir)
    //subprocess.Run("%s install -r %s/africana-base/requirements.txt", venvPip, utils.BaseDir)
    subprocess.Run("%s -m pip install -r %s/requirements.txt", utils.VenvPython, utils.ToolsDir)
}

func CreatePythonVenv() error {
    venvPath := filepath.Join(utils.BaseDir, utils.PyEnvName)

    if _, err := os.Stat(filepath.Join(venvPath, "bin", "python")); err == nil {
        fmt.Printf("%s%s[+] %sPython virtual environment already exists\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        return nil
    }

    fmt.Printf("%s%s[+] %sCreating Python virtual environment...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    subprocess.Run("python3 -m venv %s --upgrade-deps", venvPath)

    fmt.Printf("%s%s[+] %sPython virtual environment created at %s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, venvPath, bcolors.Endc)
    return nil
}

func baseLinuxSetup(missingTools map[string]map[string]string, foundationCommands []string) {
    if _, err := os.Stat(utils.ToolsDir); os.IsNotExist(err) {
        SetupGoEnvironment(utils.PyEnvName)

        if err := CreatePythonVenv(); err != nil {
            fmt.Printf("%s%s[!] %sFailed to create Python venv: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err, bcolors.Endc)
        }

        InstallFoundationTools(foundationCommands)
        InstallTools(missingTools)

        if !utils.IsArchLinux() {
            subprocess.Run("winecfg /v win11")
        }
        fmt.Printf("\n\n%s%s[*] %sInstalling third party tools\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        InstallGithubTools()

        fmt.Printf("\n%s%s[*] %sAfricana successfully installed.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    } else {
        UpdateAfricana()
    }
}

func AutoSetups() {
    switch runtime.GOOS {
    case "linux":
        if utils.DetectAndroid() {
            fmt.Printf("%s%s[+] %sAndroid Detected.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
            AndroidSetups()
        } else if distroID, err := utils.GetLinuxDistroID(); err == nil {
            if task, ok := linuxTaskMap[distroID]; ok {
                task()
            } else {
                fmt.Printf("%s%s[!] %sUnsupported Linux distro: %s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, distroID)
            }
        } else {
            fmt.Printf("%s%s[!] %sLinux distro detection failed.\n", bcolors.Bold, bcolors.Red, bcolors.Endc)
        }
    case "windows":
        fmt.Printf("%s%s[+] %sWindows detected.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        WindowsSetups()
    case "darwin":
        fmt.Printf("%s%s[+] %smacOS detected.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        MacosSetups()
    default:
        fmt.Printf("%s%s[!] %sUnsupported OS: %s\n", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, runtime.GOOS)
    }
}

func KaliSetups() {
    fmt.Printf("\n%s%s[>] %sInstalling africana in kali.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()

    commands := []string{
        "wget https://archive.kali.org/archive-keyring.gpg -O /usr/share/keyrings/kali-archive-keyring.gpg",
        "cd /etc/apt/trusted.gpg.d/; wget -qO - https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg",
        "cd /etc/apt/sources.list.d/; wget -qO - https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list",
        "dpkg --add-architecture i386",
        "apt-get update -y",
        "apt-get install zsh git curl -y",
    }
    baseLinuxSetup(missingTools, commands)
}

func UbuntuSetups() {
    fmt.Printf("\n%s%s[>] %sInstalling africana in ubuntu.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()
    commands := []string{
        "wget https://archive.kali.org/archive-keyring.gpg -O /usr/share/keyrings/kali-archive-keyring.gpg",
        "cd /etc/apt/trusted.gpg.d/; wget -qO - https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg",
        "cd /etc/apt/sources.list.d/; wget -qO - https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list",
        "dpkg --add-architecture i386",
        "add-apt-repository multiverse",
        "apt-get update -y; apt-get install zsh* git curl ubuntu-restricted-extras gnome-shell-extension-manager gnome-shell-extensions gnome-tweaks",
    }
    baseLinuxSetup(missingTools, commands)
}

func ArchSetups() {
    fmt.Printf("\n%s%s[>] %sSetting up Arch Linux.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    missingTools := UpsentTools()
    commands := []string{
        "sudo pacman -Syu --noconfirm",
        "sudo pacman -S --noconfirm base-devel git curl zsh go",
    }
    baseLinuxSetup(missingTools, commands)
}

func UpdateAfricana() {
    fmt.Printf("\n%s%s[!] %sAfricana already installed. Updating ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s; git pull .", utils.ToolsDir)
    fmt.Printf("\n%s%s[+] %sSuccesfully updated Base-tools.\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    fmt.Printf("\n%s%s[*] %sUpdating africana console ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    subprocess.Run("cd %s; git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1; cd ./africana-framework; make; cd ./build; mv ./* /usr/local/bin/afrconsole; rm -rf ../africana-framework", utils.BaseDir)
    fmt.Printf("\n%s%s[*] %sSuccesfully updated africana console ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
}

func Uninstaller() {
    fmt.Printf("\n%s%s[!] %sUninstalling africana.", bcolors.Bold, bcolors.Red, bcolors.Endc)
    if _, err := os.Stat(utils.BaseDir); !os.IsNotExist(err) {
        subprocess.Run("rm -rf %s; rm -rf /usr/local/bin/afrconsole", utils.BaseDir)
        fmt.Printf("\n%s%s[*] %sAfricana uninstalled. Goodbye!\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        os.Exit(0)
    } else {
        fmt.Printf("\n%s%s[!] %sAfricana is not installed.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    }
}

func AndroidSetups() {
    fmt.Printf("\n%s%s[+] %sAndroid setup initiated.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    if !utils.DetectAndroid() {
        fmt.Printf("%s%s[!] %sThis setup is intended for Android devices only.\n", bcolors.Bold, bcolors.Red, bcolors.Endc)
        return
    }

    missingTools := UpsentTools()

    foundationCommands := []string{
        "pkg update -y",
        "pkg install root-repo x11-repo unstable-repo -y",
        "pkg update -y",
        "pkg upgrade -y",
        "pkg install wget curl git python golang screenfetch zsh* -y",

        "git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/.zsh",
        "git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions ~/.zsh/zsh-autosuggestions",

        "echo 'screenfetch' >> ~/.zshrc",
        "echo 'screenfetch' >> ~/.bashrc",

        "echo 'source ~/.zsh/powerlevel10k.zsh-theme' >> ~/.zshrc",
        "echo 'source ~/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh' >> ~/.zshrc",

        "chsh -s zsh",
        "termux-setup-storage",
        "wget -O install-nethunter-termux https://offs.ec/2MceZWr",
        "chmod +x install-nethunter-termux",
        "./install-nethunter-termux",
        "nethunter",
        "sudo apt update; sudo apt full-upgrade -y",
    }

    baseLinuxSetup(missingTools, foundationCommands)

    fmt.Printf("\n%s%s[+] %sAndroid setup completed successfully!\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    fmt.Printf("%s%s[!] %sNote: Some tools may require root access or additional setup.\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
}

func MacosSetups() {
    fmt.Printf("\n%s%s[+] %smacOS setup placeholder.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
}

func WindowsSetups() {
    fmt.Printf("\n%s%s[+] %sWindows setup placeholder.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
}
