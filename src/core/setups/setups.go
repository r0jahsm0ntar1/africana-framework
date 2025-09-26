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
    "kali":    LinuxSetups,
    "debian":  LinuxSetups,
    "ubuntu":  LinuxSetups,
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
            "playit":                       "playit",
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

        {[]string{"info"}, menus.HelpInfoSetups},
        {[]string{"m", "menu"}, menus.MenuOne},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.SetupsOptions(utils.Distro, Function)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListSetupsFunction},
        {[]string{"distro", "distros", "list distro", "list distros", "show distro", "show distros"}, menus.ListSetupsDistros},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run kali", "use kali", "exec kali", "start kali", "launch kali", "exploit kali", "execute kali"}, func() {SetupsFunction(Function, "kali")}},
        {[]string{"? 1", "info 1", "help 1", "kali", "info kali", "help kali"}, func() {menus.HelpInfoKali(utils.Distro, Function)}},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run ubuntu", "use ubuntu", "exec ubuntu", "start ubuntu", "launch ubuntu", "exploit ubuntu", "execute ubuntu"}, func() {SetupsFunction(Function, "ubuntu")}},
        {[]string{"? 2", "info 2", "help 2", "ubuntu", "info ubuntu", "help ubuntu"}, func() {menus.HelpInfoUbuntu(utils.Distro, Function)}},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run arch", "use arch", "exec arch", "start arch", "launch arch", "exploit arch", "execute arch"}, func() {SetupsFunction(Function, "arch")}},
        {[]string{"? 3", "info 3", "help 3", "arch", "info arch", "help arch"}, func() {menus.HelpInfoArch(utils.Distro, Function)}},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run macos", "use macos", "exec macos", "start macos", "launch macos", "exploit macos", "execute macos"}, func() {SetupsFunction(Function, "macos")}},
        {[]string{"? 4", "info 4", "help 4", "macos", "info macos", "help macos"}, func() {menus.HelpInfoMacos(utils.Distro, Function)}},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run android", "use android", "exec android", "start android", "launch android", "exploit android", "execute android"}, func() {SetupsFunction(Function, "android")}},
        {[]string{"? 5", "info 5", "help 5", "android", "info android", "help android"}, func() {menus.HelpInfoAndroid(utils.Distro, Function)}},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run windows", "use windows", "exec windows", "start windows", "launch windows", "exploit windows", "execute windows"}, func() {SetupsFunction(Function, "windows")}},
        {[]string{"? 6", "info 6", "help 6", "windows", "info windows", "help windows"}, func() {menus.HelpInfoWindows(utils.Distro, Function)}},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run update", "use update", "exec update", "start update", "launch update", "exploit update", "execute update"}, func() {SetupsFunction("update", utils.Distro)}},
        {[]string{"? 7", "info 7", "help 7", "update", "info update", "help update"}, func() {menus.HelpInfoUpdate(utils.Distro, Function)}},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run repair", "use repair", "exec repair", "start repair", "launch repair", "exploit repair", "execute repair"}, func() {SetupsFunction("repair", utils.Distro)}},
        {[]string{"? 8", "info 8", "help 8", "repair", "info repair", "help repair"}, func() {menus.HelpInfoRepair(utils.Distro, Function)}},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run uninstall", "use uninstall", "exec uninstall", "start uninstall", "launch uninstall", "exploit uninstall", "execute uninstall"}, func() {SetupsFunction("uninstall", utils.Distro)}},
        {[]string{"? 9", "info 9", "help 9", "uninstall", "info uninstall", "help uninstall"}, func() {menus.HelpInfoUninstall(utils.Distro, Function)}},

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
      "gateway": &utils.GateWay,
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
      "venvname": &utils.VenvName,
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
      "gateway": &utils.GateWay,
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
      "venvname": &utils.VenvName,
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

        "1": func() {LinuxSetups()},
        "2": func() {LinuxSetups()},
        "3": func() {ArchSetups()},
        "4": func() {MacosSetups()},
        "5": func() {AndroidSetups()},
        "6": func() {WindowsSetups()},
        "7": func() {UpdateAfricana()},
        "8": func() {UpdateAfricana()},
        "9": func() {Uninstaller()},

        "auto":      func() {AutoSetups()},

        "kali":      func() {LinuxSetups()},
        "ubuntu":    func() {LinuxSetups()},
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
            fmt.Printf("%s%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
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
        {[]string{"2", "kali"}, LinuxSetups},
        {[]string{"3", "macos"}, MacosSetups},
        {[]string{"4", "ubuntu"}, LinuxSetups},
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
                fmt.Printf("  %s- %s%s ...\n", bcolors.Bold, bcolors.Endc, tool)
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

    isAndroid := utils.DetectAndroid()
    isArchLinux := utils.IsArchLinux()
    isMacOS := runtime.GOOS == "darwin"
    isWindows := runtime.GOOS == "windows"

    for _, cat := range categories {
        if len(tools[cat.key]) > 0 {
            fmt.Printf("\n%sInstalling %s tools%s\n", bcolors.Bold, cat.name, bcolors.Endc)

            var goTools []string
            var systemPackages []string

            for tool, pkg := range tools[cat.key] {
                if strings.Contains(pkg, "@latest") {
                    goTools = append(goTools, pkg)
                } else {
                    systemPackages = append(systemPackages, getPackageName(tool, pkg))
                }
                fmt.Printf("%s%s[+] %sQueuing %s%-20s %s ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Blue, tool, bcolors.Endc)
            }

            if len(systemPackages) > 0 {
                fmt.Printf("%s%s[*] %sInstalling %d system packages individually ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, len(systemPackages))

                if !isArchLinux && !isAndroid && !isMacOS && !isWindows {
                    fmt.Printf("%s%s[*] %sChecking package manager health ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
                    subprocess.Run("sudo apt --fix-broken install -y")
                    subprocess.Run("sudo dpkg --configure -a")
                }

                for _, pkg := range systemPackages {
                    fmt.Printf("\n%s%s[>] %sInstalling%s: %s%s%s ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Blue, bcolors.Endc, bcolors.Bold, pkg, bcolors.Endc)

                    var err error
                    var maxRetries = 2

                    for attempt := 1; attempt <= maxRetries; attempt++ {
                        switch {
                        case isArchLinux:
                            err = subprocess.Run("sudo pacman -S --noconfirm %s", pkg)
                        case isAndroid:
                            err = subprocess.Run("pkg install -y %s", pkg)
                        case isWindows:
                            err = subprocess.Run("choco install -y --force %s", pkg)
                        case isMacOS:
                            err = subprocess.Run("brew install %s", pkg)
                        default:
                            err = subprocess.Run("sudo apt-get install -y %s", pkg)

                            if err != nil && attempt < maxRetries {
                                fmt.Printf("%s%s[!] %sInstallation failed, attempting to fix broken packages ...%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Endc)
                                subprocess.Run("sudo apt --fix-broken install -y")
                                subprocess.Run("sudo dpkg --configure -a")
                                time.Sleep(2 * time.Second)
                                continue
                            }
                        }

                        break
                    }

                    if err != nil {
                        fmt.Printf("%s%s[!] %sFailed to install %s: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, pkg, err, bcolors.Endc)

                        if !isArchLinux && !isAndroid && !isMacOS && !isWindows {
                            fmt.Printf("%s%s[*] %sAttempting comprehensive repair ...%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, bcolors.Endc)
                            subprocess.Run("sudo apt --fix-broken install -y")
                            subprocess.Run("sudo dpkg --configure -a")
                            subprocess.Run("sudo apt-get autoremove -y")
                            subprocess.Run("sudo apt-get autoclean -y")
                        }
                    } else {
                        fmt.Printf("%s%s[✓] %sSuccessfully installed %s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, pkg, bcolors.Endc)
                    }
                    time.Sleep(90 * time.Millisecond)
                }
            }

            if len(goTools) > 0 {
                fmt.Printf("%s%s[*] %sInstalling %d Go tools ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, len(goTools))

                if isAndroid {
                    subprocess.Run("pkg install golang -y")
                } else if isWindows {
                    subprocess.Run("choco install golang -y --force")
                } else if isMacOS {
                    subprocess.Run("brew install go")
                } else if !isArchLinux {
                    subprocess.Run("sudo apt-get install -y golang-go")
                }

                for _, pkg := range goTools {
                    fmt.Printf("\n%s%s[>] %sInstalling Go tool%s: %s%s%s ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Blue, bcolors.Endc, bcolors.Bold, pkg, bcolors.Endc)

                    var err error
                    if isAndroid {
                        err = subprocess.Run("go install %s", pkg)
                    } else if isWindows {
                        err = subprocess.Run("go install %s", pkg)
                    } else if isMacOS {
                        err = subprocess.Run("go install %s", pkg)
                    } else {
                        err = subprocess.Run("go install %s", pkg)
                    }

                    if err != nil {
                        fmt.Printf("%s%s[!] %sFailed to install Go tool %s: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, pkg, err, bcolors.Endc)
                    } else {
                        fmt.Printf("%s%s[✓] %sSuccessfully installed Go tool %s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, pkg, bcolors.Endc)
                    }
                    time.Sleep(90 * time.Millisecond)
                }
            }
        }
    }
}

func SetupGoPyEnv(VenvName string) error {
    os.Setenv("GOPATH", utils.GoPath)
    os.Setenv("PATH", os.Getenv("PATH") + ":/usr/local/go/bin:" + filepath.Join(utils.GoPath, "bin"))

    goDirs := []string{
        utils.GoPath,
        filepath.Join(utils.GoPath, "bin"),
        filepath.Join(utils.GoPath, "pkg"),
        filepath.Join(utils.GoPath, "src"),
    }

    for _, dir := range goDirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            fmt.Printf("%s%s[!] %sFailed to create directory %s: %v%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, dir, err, bcolors.Endc)
            return err
        }
    }

    os.Setenv("VIRTUAL_ENV", utils.VenvPath)
    os.Setenv("PATH", filepath.Join(utils.VenvPath, "bin") + ":" + os.Getenv("PATH"))

    if _, err := os.Stat(filepath.Join(utils.VenvPath, "bin", "python")); err == nil {
        fmt.Printf("\n%s%s[!] %sPython virtual environment already exists\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    } else {
        if err := os.MkdirAll(utils.VenvPath, 0755); err != nil {
            fmt.Printf("%s%s[!] %sFailed to create Python venv directory %s: %v%s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, utils.VenvPath, err, bcolors.Endc)
            return err
        }

        fmt.Printf("\n%s%s[+] %sCreating Python virtual environment ...", bcolors.Bold, bcolors.Green, bcolors.Endc)

        subprocess.Run("python3 -m pip install --upgrade pip; python3 -m venv %s --upgrade-deps", utils.VenvPath)

        if _, err := os.Stat(filepath.Join(utils.VenvPath, "bin", "python")); os.IsNotExist(err) {
            return fmt.Errorf("[!] Failed to create Python virtual environment at %s", utils.VenvPath)
        }

        fmt.Printf("\n%s%s[!] %sPython virtual environment created at %s%s", bcolors.Bold, bcolors.Blue, bcolors.Endc, utils.VenvPath, bcolors.Endc)
    }

    shellProfiles := []string{
        filepath.Join(utils.HomeDir, ".bashrc"),
        filepath.Join(utils.HomeDir, ".zshrc"),
    }

    envSetup := fmt.Sprintf(`
# Go environment
export GOPATH=%s
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Python virtual environment
#export VIRTUAL_ENV=%s
#export PATH=$VIRTUAL_ENV/bin:$PATH
`, utils.GoPath, utils.VenvPath)

    for _, profile := range shellProfiles {
        if err := utils.AppendToShellProfile(profile, envSetup); err != nil {
            fmt.Printf("%s%s[!] %sFailed to update %s: %v%s\n", bcolors.Yellow, bcolors.Endc, profile, err, bcolors.Endc)
        }
    }

    return nil
}

func installThirdPartyTools() {
    isWindows := runtime.GOOS == "windows"
    fmt.Printf("\n%s%s[*] %sInstalling third party tools\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    gitCloneCmd := "cd %s; git clone --depth 1 --progress https://github.com/r0jahsm0ntar1/africana-base.git --depth 1"
    if isWindows {
        subprocess.Run("cd %s; git clone --depth 1 --progress https://github.com/r0jahsm0ntar1/africana-base.git --depth 1", utils.BaseDir)
    } else {
        subprocess.Run(gitCloneCmd, utils.BaseDir)
    }

    pipInstallCmd := "%s -m pip install --retries 10 --timeout 360 -r %s/requirements.txt"
    if isWindows {
        subprocess.Run("python -m pip install --retries 10 --timeout 360 -r %s\\requirements.txt", utils.ToolsDir)
    } else {
        subprocess.Run(pipInstallCmd, utils.VenvPython, utils.ToolsDir)
    }

    fmt.Printf("\n%s%s[*] %sAfricana successfully installed.", bcolors.Bold, bcolors.Green, bcolors.Endc)
}

func baseSetup(foundationCommands []string, missingTools ...map[string]map[string]string) {
    var tools map[string]map[string]string
    if len(missingTools) > 0 {
        tools = missingTools[0]
    } else {
        tools = make(map[string]map[string]string)
    }

    isAndroid := utils.DetectAndroid()
    isMacOS := runtime.GOOS == "darwin"
    isWindows := runtime.GOOS == "windows"
    isRegularLinux := !isAndroid && !isMacOS && !isWindows

    if _, err := os.Stat(utils.ToolsDir); os.IsNotExist(err) {
        if isAndroid {
            fmt.Printf("\n%s%s[!] %sSetting up Android environment ...", bcolors.Bold, bcolors.Red, bcolors.Endc)

            for _, cmd := range foundationCommands {
                androidCmd := convertToAndroidCmd(cmd)
                fmt.Printf("\n%s%s[*] %s%s%sRunning%s: %s%s%s ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Bold, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, androidCmd, bcolors.Endc)
                subprocess.Run(androidCmd)
                time.Sleep(90 * time.Millisecond)
            }
        } else if isMacOS {
            fmt.Printf("%s%s[*] %sSetting up macOS environment ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

            for _, cmd := range foundationCommands {
                fmt.Printf("\n%s%s[*] %s%s%sRunning%s: %s%s%s ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Bold, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, cmd, bcolors.Endc)
                subprocess.Run(cmd)
                time.Sleep(90 * time.Millisecond)
            }
        } else if isWindows {
            fmt.Printf("%s%s[*] %sSetting up Windows environment ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

            for _, cmd := range foundationCommands {
                fmt.Printf("\n%s%s[*] %s%s%sRunning%s: %s%s%s ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Bold, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, cmd, bcolors.Endc)
                subprocess.Run(cmd)
                time.Sleep(90 * time.Millisecond)
            }
        } else {
            for _, cmd := range foundationCommands {
                fmt.Printf("\n%s%s[*] %s%s%sRunning%s: %s%s%s ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Bold, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, cmd, bcolors.Endc)
                subprocess.Run(cmd)
                time.Sleep(90 * time.Millisecond)
            }
        }

        if len(tools) > 0 {
            InstallTools(tools)
        }

        if isRegularLinux {
            if err := SetupGoPyEnv(utils.VenvName); err != nil {
                fmt.Printf("\n%s%s[!] %sFailed to create Python venv: %v%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err, bcolors.Endc)
            }
        }

        if !utils.IsArchLinux() && !isAndroid && !isWindows {
            subprocess.Run("winecfg /v win11")
        }

    } else {
        UpdateAfricana()
    }
}

func convertToAndroidCmd(cmd string) string {
    cmd = strings.Replace(cmd, "sudo", "", -1)
    for _, pm := range []string{"apt install", "apt-get install", "yum install", "dnf install"} {
        if strings.Contains(cmd, pm) {
            return strings.Replace(cmd, pm, "pkg install", 1)
        }
    }

    return cmd
}

func setupChocolatey() error {
    if err := subprocess.Run("choco --version"); err != nil {
        fmt.Printf("%s%s[!] %sInstalling Chocolatey ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        
        installCmd := `@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"`
        if err := subprocess.Run(installCmd); err != nil {
            return fmt.Errorf("failed to install Chocolatey: %v", err)
        }

        subprocess.Run("refreshenv")
        time.Sleep(2 * time.Second)
    }
    return nil
}

func setupBrew() error {
    if err := subprocess.Run("brew --version"); err != nil {
        fmt.Printf("%s%s[!] %sInstalling Homebrew ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)

        installCmd := `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
        if err := subprocess.Run(installCmd); err != nil {
            return fmt.Errorf("[!] failed to install Homebrew: %v", err)
        }

        subprocess.Run(`echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zshrc`)
        subprocess.Run(`eval "$(/opt/homebrew/bin/brew shellenv)"`)
    }
    return nil
}

func getFoundationCommands() []string {
    isArchLinux := utils.IsArchLinux()
    isAndroid := utils.DetectAndroid()
    isMacOS := runtime.GOOS == "darwin"
    isWindows := runtime.GOOS == "windows"

    if isAndroid {
        return []string{
            "pkg update -y",
            "pkg install root-repo x11-repo unstable-repo -y",
            "pkg install wget curl git golang screenfetch zsh* termux-api openssh python -y",
            "git clone --depth 1 --progress https://github.com/romkatv/powerlevel10k.git ~/.zsh",
            "git clone --depth 1 --progress https://github.com/zsh-users/zsh-autosuggestions ~/.zsh/zsh-autosuggestions",
            "wget https://gist.githubusercontent.com/noahbliss/4fec4f5fa2d2a2bc857cccc5d00b19b6/raw/db5ceb8b3f54b42f0474105b4a7a138ce97c0b7a/kali-zshrc -O ~/.zshrc",
            "echo 'screenfetch' >> ~/.zshrc",
            "echo 'screenfetch' >> ~/.bashrc",
            "echo 'source ~/.zsh/powerlevel10k.zsh-theme' >> ~/.zshrc",
            "echo 'source ~/.zsh/zsh-autosuggestions/zsh-autosuggestions.zsh' >> ~/.zshrc",
            "chsh -s zsh",
            "termux-setup-storage",
        }
    } else if isMacOS {
        return []string{
            "brew install git",
            "brew install python",
            "brew install go",
            "brew install wget",
            "brew install curl",
            "brew install gnu-sed",
            "brew install gnu-tar",
        }
    } else if isWindows {
        return []string{
            "choco install git -y --force",
            "choco install python3 -y --force",
            "choco install golang -y --force",
            "choco install wget -y --force",
            "choco install curl -y --force",
            "choco install 7zip -y --force",
            "choco install vscode -y --force",
            "choco install sysinternals -y --force",
        }
    } else if isArchLinux {
        return []string{
            "sudo pacman -Syu --noconfirm",
            "sudo pacman -S --noconfirm base-devel git curl zsh go python python-pip wget",
        }
    }
    return []string{
        "sudo apt-get -o Acquire::Retries=5 -o Acquire::http::Timeout='30' -o Acquire::https::Timeout='30' -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confnew' update -y",
        "sudo apt-get -o Acquire::Retries=5 -o Acquire::http::Timeout='30' -o Acquire::https::Timeout='30' -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confnew' install curl wget gpg apt-transport-https -y",
        "sudo mkdir -p /usr/share/keyrings /etc/apt/sources.list.d /etc/apt/trusted.gpg.d",
        "sudo curl -fsSL https://archive.kali.org/archive-keyring.gpg -o /usr/share/keyrings/kali-archive-keyring.gpg",
        "curl -SsL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/playit.gpg >/dev/null",
        "echo 'deb [signed-by=/etc/apt/trusted.gpg.d/playit.gpg] https://playit-cloud.github.io/ppa/data ./' | sudo tee /etc/apt/sources.list.d/playit-cloud.list",
        "curl -fsSL https://packages.microsoft.com/keys/microsoft.asc | sudo gpg --dearmor -o /usr/share/keyrings/microsoft.gpg",
        "echo 'deb [arch=amd64,armhf,arm64 signed-by=/usr/share/keyrings/microsoft.gpg] https://packages.microsoft.com/repos/microsoft-debian-bullseye-prod bullseye main' | sudo tee /etc/apt/sources.list.d/microsoft.list",
        "echo 'Package: powershell\nPin: origin packages.microsoft.com\nPin-Priority: 1001' | sudo tee /etc/apt/preferences.d/powershell-pin",
        "sudo dpkg --add-architecture i386",
        "sudo apt-get -o Acquire::Retries=5 -o Acquire::http::Timeout='30' -o Acquire::https::Timeout='30' -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confnew' update -y",
    }
}

func AutoSetups() {
    switch runtime.GOOS {
    case "android":
        fmt.Printf("\n%s%s[!] %sAndroid distro detected ...", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        AndroidSetups()
    case "linux":
        if utils.DetectAndroid() {
            fmt.Printf("\n%s%s[!] %sAndroid distro detected.", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
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

func LinuxSetups() {
    fmt.Printf("%s%s[*] %sSetting up Linux environment ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    foundationCommands := getFoundationCommands()
    missingTools := UpsentTools()
    baseSetup(foundationCommands, missingTools)
    installThirdPartyTools()
}

func ArchSetups() {
    fmt.Printf("\n%s%s[*] %sSetting up Arch Linux environment ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    foundationCommands := getFoundationCommands()
    missingTools := UpsentTools()
    baseSetup(foundationCommands, missingTools)
    installThirdPartyTools()
}

func AndroidSetups() {
    fmt.Printf("\n%s%s[*] %sInstalling africana in android ...", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    if !utils.DetectAndroid() {
        fmt.Printf("\n%s%s[!] %sThis setup is intended for Android devices only.", bcolors.Bold, bcolors.Red, bcolors.Endc)
        return
    }

    if utils.IsNethunter() {
        fmt.Printf("\n%s%s[+] %sNetHunter is already installed ...", bcolors.Bold, bcolors.Green, bcolors.Endc)
        fmt.Printf("\n%s%s[?] %sWould you like to update NetHunter? (y/N): ", bcolors.Bold, bcolors.Cyan, bcolors.Endc)
        utils.Scanner.Scan()
        Input := strings.TrimSpace(utils.Scanner.Text())

        switch Input {
        case "y", "yes":
            installNetHunter()
            return
        case "n", "q", "no", "exit", "quit":
            fmt.Printf("\n%s%s[!] %sNetHunter update skipped ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
            return
        default:
            fmt.Printf("\n%s%s[!] %sNetHunter update skipped ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        }
    }

    fmt.Printf("\n%s%s[!] %sNetHunter not detected. Please install NetHunter first ...", bcolors.Bold, bcolors.Red, bcolors.Endc)
    fmt.Printf("\n%s%s[?] %sWould you like to install NetHunter? (y/N): ", bcolors.Bold, bcolors.Cyan, bcolors.Endc)

    utils.Scanner.Scan()
    Input := strings.TrimSpace(utils.Scanner.Text())

    switch Input {
    case "y", "yes":
        installNetHunter()
        return
    case "n", "q", "no", "exit", "quit":
        fmt.Printf("\n%s%s[>] %sInstalling android foundation tools ...", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        foundationCommands := getFoundationCommands()
        baseSetup(foundationCommands)
        fmt.Printf("\n%s%s[!] %sNetHunter installation skipped. Some tools may not work.\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
        return
    default:
        fmt.Printf("\n%s%s[!] %sNetHunter installation skipped. Some tools may not work.\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    }

}

func MacosSetups() {
    fmt.Printf("%s%s[*] %sSetting up macOS environment ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    if err := setupBrew(); err != nil {
        fmt.Printf("%s%s[!] %s%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        return
    }

    foundationCommands := getFoundationCommands()
    missingTools := UpsentTools()
    baseSetup(foundationCommands, missingTools)
    installThirdPartyTools()
}

func WindowsSetups() {
    fmt.Printf("%s%s[*] %sSetting up Windows environment ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    if err := setupChocolatey(); err != nil {
        fmt.Printf("%s%s[!] %s%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        return
    }

    foundationCommands := getFoundationCommands()
    missingTools := UpsentTools()
    baseSetup(foundationCommands, missingTools)
    installThirdPartyTools()
}

func installNetHunter() {
    fmt.Printf("\n%s%s[>] %sInstalling NetHunter ...", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    foundationCommands := getFoundationCommands()
    baseSetup(foundationCommands)

    fmt.Printf("\n%s%s[>] %sDownloading NetHunter installer ...\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    installerScript := `
#!/data/data/com.termux/files/usr/bin/bash

################################################################################
#                                                                              #
#     Termux NetHunter Installer.                                              #
#                                                                              #
#     Installs Kali NetHunter in Termux.                                       #
#                                                                              #
#     Copyright (C) 2023-2025  Jore <https://github.com/jorexdeveloper>        #
#                                                                              #
#     This program is free software: you can redistribute it and/or modify     #
#     it under the terms of the GNU General Public License as published by     #
#     the Free Software Foundation, either version 3 of the License, or        #
#     (at your option) any later version.                                      #
#                                                                              #
#     This program is distributed in the hope that it will be useful,          #
#     but WITHOUT ANY WARRANTY; without even the implied warranty of           #
#     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the            #
#     GNU General Public License for more details.                             #
#                                                                              #
#     You should have received a copy of the GNU General Public License        #
#     along with this program.  If not, see <https://www.gnu.org/licenses/>.   #
#                                                                              #
################################################################################
# shellcheck disable=SC2034,SC2155

# ATTENTION!!! CHANGE BELOW FUNTIONS FOR DISTRO DEPENDENT ACTIONS!!!

################################################################################
# Called before any safety checks                                              #
# New Variables: AUTHOR GITHUB LOG_FILE ACTION_INSTALL ACTION_CONFIGURE        #
#                ROOTFS_DIRECTORY COLOR_SUPPORT all_available_colors           #
################################################################################
pre_check_actions() {
	return
}

################################################################################
# Called before printing intro                                                 #
# New Variables: none                                                          #
################################################################################
distro_banner() {
	local spaces=''
	for ((i = $((($(stty size | cut -d ' ' -f2) - 49) / 2)); i > 0; i--)); do
		spaces+=' '
	done
	msg -a "${spaces}${B}.............."
	msg -a "${spaces}${B}            ..,;:ccc,."
	msg -a "${spaces}${B}          ......''';lxO."
	msg -a "${spaces}${B}.....''''..........,:ld;"
	msg -a "${spaces}${B}           .';;;:::;,,.x,"
	msg -a "${spaces}${B}      ..'''.            0Xxoc:,.  ..."
	msg -a "${spaces}${B}  ....                ,ONkc;,;cokOdc',."
	msg -a "${spaces}${B} .                   OMo           ':${R}dd${B}o."
	msg -a "${spaces}${B}                    dMc               :OO;"
	msg -a "${spaces}${B}                    0M.                 .:o."
	msg -a "${spaces}${B}                    ;Wd"
	msg -a "${spaces}${B}                     ;XO,"
	msg -a "${spaces}${B}                       ,d0Odlc;,.."
	msg -a "${spaces}${B}                           ..',;:cdOOd::,."
	msg -a "${spaces}${B}                                    .:d;.':;."
	msg -a "${spaces}${B}                                       'd,  .'"
	msg -a "${spaces}${C}${DISTRO_NAME}${B}                           ;l   .."
	msg -a "${spaces}${Y}    ${VERSION_NAME}${B}                                    .o"
	msg -a "${spaces}${B}                                            c"
	msg -a "${spaces}${B}                                            .'"
	msg -a "${spaces}${B}                                             ."
}

################################################################################
# Called after checking architecture and required pkgs                         #
# New Variables: SYS_ARCH LIB_GCC_PATH                                         #
################################################################################
post_check_actions() {
	return
}

################################################################################
# Called after checking for rootfs directory                                   #
# New Variables: KEEP_ROOTFS_DIRECTORY                                         #
################################################################################
pre_install_actions() {
	if [ -z "${KEEP_ROOTFS_DIRECTORY}" ]; then
		choose -d2 -t "Select your prefered installation." \
			"full - GUI + All ${DISTRO_NAME} Packages" \
			"mini - Essential Packages Only" \
			"nano - Essential Packages++"
		SELECTED_INSTALLATION=${?}
		case "${SELECTED_INSTALLATION}" in
			1)
				SELECTED_INSTALLATION="full"
				GUI_INSTALLED=true
				;;
			3) SELECTED_INSTALLATION="nano" ;;
			*) SELECTED_INSTALLATION="mini" ;;
		esac
		msg "Okay then, I shall install a '${Y}${SELECTED_INSTALLATION}${C}' rootfs."
		ARCHIVE_NAME="kali-nethunter-rootfs-${SELECTED_INSTALLATION/mini/minimal}-${SYS_ARCH}.tar.xz"
	fi
}

################################################################################
# Called after extracting rootfs                                               #
# New Variables: KEEP_ROOTFS_ARCHIVE                                           #
################################################################################
post_install_actions() {
	return
}

################################################################################
# Called before making configurations                                          #
# New Variables: none                                                          #
################################################################################
pre_config_actions() {
	mkdir -p "${ROOTFS_DIRECTORY}/etc" >>"${LOG_FILE}" 2>&1 && echo "${ROOTFS_DIRECTORY}" >"${ROOTFS_DIRECTORY}/etc/debian_chroot"
}

################################################################################
# Called after configurations                                                  #
# New Variables: none                                                          #
################################################################################
post_config_actions() {
	# Fix environment variables on login or su. (#17 fix)
	# local fix="session  required  pam_env.so readenv=1"
	# for f in su su-l system-local-login system-remote-login; do
	# 	if [ -f "${ROOTFS_DIRECTORY}/etc/pam.d/${f}" ] && ! grep -q "${fix}" "${ROOTFS_DIRECTORY}/etc/pam.d/${f}" >>"${LOG_FILE}" 2>&1; then
	# 		echo "${fix}" >>"${ROOTFS_DIRECTORY}/etc/pam.d/${f}"
	# 	fi
	# done
	# execute distro specific command for locale generation
	if [ -f "${ROOTFS_DIRECTORY}/etc/locale.gen" ] && [ -x "${ROOTFS_DIRECTORY}/sbin/dpkg-reconfigure" ]; then
		msg -t "Hold on while I generate the locales for you."
		sed -i -E 's/#[[:space:]]?(en_US.UTF-8[[:space:]]+UTF-8)/\1/g' "${ROOTFS_DIRECTORY}/etc/locale.gen"
		if distro_exec DEBIAN_FRONTEND=noninteractive /sbin/dpkg-reconfigure locales >>"${LOG_FILE}" 2>&1; then
			msg -s "Done, the locales are ready!"
		else
			msg -e "I failed to generate the locales."
		fi
	fi
}

################################################################################
# Called before complete message                                               #
# New Variables: none                                                          #
################################################################################
pre_complete_actions() {
	if ! ${GUI_INSTALLED:-false} && [ "${SELECTED_INSTALLATION}" != "full" ] && ask -y -- -t "Should I set up the GUI now?"; then
		set_up_gui && set_up_browser && GUI_INSTALLED=true
	fi
}

################################################################################
# Called after complete message                                                #
# New Variables: none                                                          #
################################################################################
post_complete_actions() {
	return
}

################################################################################
# Local Functions                                                              #
################################################################################

# Sets up the GUI
set_up_gui() {
	local available_desktops=(
		"E17" "GNOME" "i3" "KDE" "LXDE" "MATE" "Xfce"
	)
	local -A xstartups=(
		["e17"]="enlightenment_start" ["gnome"]="gnome-session" ["i3"]="i3" ["kde"]="startplasma-x11" ["lxde"]="startlxde" ["mate"]="mate-session" ["xfce"]="startxfce4"
	)
	choose -d7 -t "Select your prefered Desktop Environment." \
		"${available_desktops[@]}"
	selected_desktop="${available_desktops[$((${?} - 1))]}"
	msg "Okay then, I shall install the '${Y}${selected_desktop}${C}' Desktop."
	msg -t "The installation is going to take very long."
	msg "Lemme me acquire the '${Y}Termux wake lock${C}'."
	if [ -x "$(command -v termux-wake-lock)" ]; then
		if termux-wake-lock >>"${LOG_FILE}" 2>&1; then
			msg -s "Great, the Termux wake lock is now activated."
		else
			msg -e "I have failed to set up the Termux wake lock."
			msg "Keep Termux open during the installation."
		fi
	else
		msg -e "I could not find the '${Y}termux-wake-lock${R}' command."
		msg "Keep Termux open during the installation."
	fi
	msg -t "Lemme first upgrade the packages in ${DISTRO_NAME}."
	msg "This won't take long."
	if distro_exec apt update && distro_exec apt full-upgrade; then
		msg -s "Done, all the ${DISTRO_NAME} packages are upgraded."
		msg -t "Now lemme install the GUI in ${DISTRO_NAME}."
		msg "This will take very long."
		if distro_exec apt install -y tigervnc-standalone-server dbus-x11 kali-desktop-"${selected_desktop,,}"; then
			msg -s "Finally, the GUI is now installed in ${DISTRO_NAME}."
			msg -t "Now lemme add the xstartup script for VNC."
			if {
				local xstartup="$(
					cat 2>>"${LOG_FILE}" <<-EOF
						#!/usr/bin/bash
						unset SESSION_MANAGER
						unset DBUS_SESSION_BUS_ADDRESS
						export XDG_RUNTIME_DIR=\${TMPDIR:-/tmp}/runtime-"\${USER:-root}"
						export SHELL="\${SHELL:-/bin/sh}"
						if [ -r ~/.Xresources ]; then
						    xrdb ~/.Xresources
						fi
						exec ${xstartups["${selected_desktop,,}"]}
					EOF
				)"
				mkdir -p "${ROOTFS_DIRECTORY}/root/.vnc"
				echo "${xstartup}" >"${ROOTFS_DIRECTORY}/root/.vnc/xstartup"
				chmod 744 "${ROOTFS_DIRECTORY}/root/.vnc/xstartup"
				if [ "${DEFAULT_LOGIN}" != "root" ]; then
					mkdir -p "${ROOTFS_DIRECTORY}/home/${DEFAULT_LOGIN}/.vnc"
					echo "${xstartup}" >"${ROOTFS_DIRECTORY}/home/${DEFAULT_LOGIN}/.vnc/xstartup"
					chmod 744 "${ROOTFS_DIRECTORY}/home/${DEFAULT_LOGIN}/.vnc/xstartup"
				fi
			} 2>>"${LOG_FILE}"; then
				msg -s "Done, xstartup script added successfully!"
			else
				msg -e "I failed to add the xstartup script."
			fi
		else
			msg -qm0 "I have failed to install the GUI in ${DISTRO_NAME}."
		fi
	else
		msg -qm0 "I have failed to upgrade the packages in ${DISTRO_NAME}."
	fi
}

# Sets up the Browser
set_up_browser() {
	local available_browsers=(
		"Chromium" "Firefox ESR" "Both Browsers"
	)
	choose -d2 -t "Select your prefered Browser." \
		"${available_browsers[@]}"
	local selected_browser="${available_browsers[$((${?} - 1))]}"
	local selected_browsers verb
	if [ "${selected_browser}" = "${available_browsers[-1]}" ]; then
		selected_browsers=("${available_browsers[@]:0:${#available_browsers[@]}-1}")
		selected_browsers=("${selected_browsers[@]// /-}")
		verb="are"
	else
		selected_browsers=("${selected_browser// /-}")
		verb="is"
	fi
	msg "Okay then, I shall install '${Y}${selected_browser}${C}'."
	if distro_exec apt install -y "${selected_browsers[@],,}"; then
		if [ "${selected_browsers[0]}" = "${available_browsers[0]}" ]; then
			sed -Ei 's/^(Exec=.*chromium).*(%U)$/\1 --no-sandbox \2/' "${ROOTFS_DIRECTORY}/usr/share/applications/chromium.desktop"
		fi
		msg "Done, ${selected_browser} ${verb} now installed in ${DISTRO_NAME}."
	else
		msg -e "I have failed to install ${selected_browser} in ${DISTRO_NAME}."
	fi
}

DISTRO_NAME="Kali NetHunter"
PROGRAM_NAME="$(basename "${0}")"
DISTRO_REPOSITORY="termux-nethunter"
VERSION_NAME="2025.3"
KERNEL_RELEASE="$(uname -r)"

SHASUM_CMD=sha256sum
TRUSTED_SHASUMS="$(
	cat <<-EOF
		8dd42a9c8eb6cb7efcb169a6824b2cdc61ff0f999e87b30effa11832c528916e  kali-nethunter-rootfs-minimal-arm64.tar.xz
		709f131a7b8ca25073553b8ac8065cf9f9d113e764d1f5f4c03c54cb47fc4475  kali-nethunter-rootfs-minimal-armhf.tar.xz
		771f511202c28074a1756859ac8211bed9d85a1cf4eddba19416b12e05492d24  kali-nethunter-rootfs-nano-arm64.tar.xz
		ae1c75b78dd1c70f37fd748561a5272015a1ae054335d78de9f0a6ed49dc1bdb  kali-nethunter-rootfs-nano-armhf.tar.xz
		b7c60dd5a1db33b399afcecc40be39415f5593f7302b6573aece1265dae44d73  kali-nethunter-rootfs-full-arm64.tar.xz
		11ee09de068493a6f7a2c8f6b1e0d5a18cb3cc511f25aca7db99e1ede82c0e15  kali-nethunter-rootfs-full-armhf.tar.xz
	EOF
)"

ARCHIVE_STRIP_DIRS=1 # directories stripped by tar when extracting rootfs archive
BASE_URL="https://kali.download/nethunter-images/kali-${VERSION_NAME}/rootfs"
TERMUX_FILES_DIR="/data/data/com.termux/files"

DISTRO_SHORTCUT="${TERMUX_FILES_DIR}/usr/bin/nh"
DISTRO_LAUNCHER="${TERMUX_FILES_DIR}/usr/bin/nethunter"

DEFAULT_ROOTFS_DIR="${TERMUX_FILES_DIR}/kali"
DEFAULT_LOGIN="kali"

# WARNING!!! DO NOT CHANGE BELOW!!!

# Check in script's directory for template
distro_template="$(realpath "$(dirname "${0}")")/termux-distro.sh"
# shellcheck disable=SC1090
if [ -f "${distro_template}" ] && [ -r "${distro_template}" ]; then
	source "${distro_template}" "${@}"
elif curl -fsSLO "https://raw.githubusercontent.com/jorexdeveloper/termux-distro/main/termux-distro.sh" 2>"/dev/null" && [ -f "${distro_template}" ]; then
	source "${distro_template}"
else
	echo "You need an active internet connection to run this script."
fi
`
    installerFile := filepath.Join(utils.SetupsLogs, "nethunter-installer.sh")
    os.WriteFile(installerFile, []byte(installerScript), 0755)

    if err := subprocess.Run(installerFile); err != nil {
        fmt.Printf("%s%s[!] %sNetHunter installation failed: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        return
    }
    fmt.Printf("%s%s[+] %sNetHunter installed successfully!", bcolors.Bold, bcolors.Blue, bcolors.Endc)
}

func UpdateAfricana() {
    fmt.Printf("\n%s%s[!] %sAfricana already installed. Updating ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    subprocess.Run(fmt.Sprintf("cd %s && git pull", utils.ToolsDir))
    fmt.Printf("\n%s%s[+] %sSuccessfully updated Base-tools.\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)

    fmt.Printf("\n%s%s[*] %sUpdating africana console ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)

    commands := []string{
        fmt.Sprintf("cd %s", utils.BaseDir),
        "rm -rf africana-framework",
        "git clone --depth 1 https://github.com/r0jahsm0ntar1/africana-framework",
        "cd africana-framework && make",
        "cd build && mv ./* /usr/local/bin/afrconsole",
        "cd ../.. && rm -rf africana-framework",
    }

    subprocess.Run(strings.Join(commands, " && "))
    fmt.Printf("\n%s%s[*] %sSuccessfully updated africana console ...\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
}

func Uninstaller() {
    fmt.Printf("\n%s%s[!] %sUninstalling africana...\n", bcolors.Bold, bcolors.Red, bcolors.Endc)
 
    if _, err := os.Stat(utils.BaseDir); os.IsNotExist(err) {
        fmt.Printf("\n%s%s[!] %sAfricana is not installed.\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
        return
    }

    subprocess.Run(fmt.Sprintf("rm -rf %s", utils.BaseDir))
    subprocess.Run("rm -f /usr/local/bin/afrconsole")

    fmt.Printf("\n%s%s[*] %sAfricana uninstalled. Goodbye!\n", bcolors.Bold, bcolors.Green, bcolors.Endc)
    os.Exit(0)
}
