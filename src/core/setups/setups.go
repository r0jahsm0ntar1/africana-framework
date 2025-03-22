package setups

import(
    "os"
    "fmt"
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
    userDistro = "kali"
    userInput, userProxy, userModule string
    userCertDir, userOutPutDir, userToolsDir, userWordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "distro": userDistro,
}

func SetupsLauncher() {
    for {
        fmt.Printf("%s%safr3%s setups(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "systems_launcher.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        userInput = strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(userInput)
        if len(buildParts) == 0 {
            continue
        }

        if executeCommand(userInput) {
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
            executeModule()
        default:
            utils.SystemShell(userInput)
        }
    }
}

func executeCommand(cmd string) bool {
    commandMap := map[string]func(){

        "? info":           menus.HelpInfo,
        "help info":        menus.HelpInfo,

        "v":                banners.Version,
        "version":          banners.Version,

        "s":                utils.Sleep,
        "sleep":            utils.Sleep,

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
        "? use":            menus.HelpInfoRun,
        "h use":            menus.HelpInfoRun,
        "info use":         menus.HelpInfoRun,
        "help use":         menus.HelpInfoRun,
        "? exec":           menus.HelpInfoRun,
        "h exec":           menus.HelpInfoRun,
        "info exec":        menus.HelpInfoRun,
        "help exec":        menus.HelpInfoRun,
        "? start":          menus.HelpInfoRun,
        "h start":          menus.HelpInfoRun,
        "info start":       menus.HelpInfoRun,
        "help start":       menus.HelpInfoRun,
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

        "? options":        menus.HelpInfOptions,
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

        "modules":          menus.ListSetupsModules,
        "show all":         menus.ListSetupsModules,
        "list all":         menus.ListSetupsModules,
        "list modules":     menus.ListSetupsModules,
        "show modules":     menus.ListSetupsModules,

        "distros":          menus.ListSetupsDistros,
        "show distros":     menus.ListSetupsDistros,

        "1":                func() {Installer("kali")},
        "run 1":            func() {Installer("kali")},
        "use 1":            func() {Installer("kali")},
        "exec 1":           func() {Installer("kali")},
        "start 1":          func() {Installer("kali")},
        "launch 1":         func() {Installer("kali")},
        "exploit 1":        func() {Installer("kali")},
        "execute 1":        func() {Installer("kali")},
        "run setups":       func() {Installer("kali")},
        "use setups":       func() {Installer("kali")},
        "exec setups":      func() {Installer("kali")},
        "start setups":     func() {Installer("kali")},
        "launch setups":    func() {Installer("kali")},
        "exploit setups":   func() {Installer("kali")},
        "execute setups":   func() {Installer("kali")},

        "? 1":              menus.HelpInfoKali,
        "info 1":           menus.HelpInfoKali,
        "help 1":           menus.HelpInfoKali,
        "setups":           menus.HelpInfoKali,
        "info setups":      menus.HelpInfoKali,
        "help setups":      menus.HelpInfoKali,

        "2":                func() {Installer("arch")},
        "run 2":            func() {Installer("arch")},
        "use 2":            func() {Installer("arch")},
        "exec 2":           func() {Installer("arch")},
        "start 2":          func() {Installer("arch")},
        "launch 2":         func() {Installer("arch")},
        "exploit 2":        func() {Installer("arch")},
        "execute 2":        func() {Installer("arch")},
        "run anonsurf":     func() {Installer("arch")},
        "use anonsurf":     func() {Installer("arch")},
        "exec anonsurf":    func() {Installer("arch")},
        "start anonsurf":   func() {Installer("arch")},
        "launch anonsurf":  func() {Installer("arch")},
        "exploit anonsurf": func() {Installer("arch")},
        "execute anonsurf": func() {Installer("arch")},

        "? 2":              menus.HelpInfoArch,
        "info 2":           menus.HelpInfoArch,
        "help 2":           menus.HelpInfoArch,
        "anonsurf":         menus.HelpInfoArch,
        "info anonsurf":    menus.HelpInfoArch,
        "help anonsurf":    menus.HelpInfoArch,

        "3":                func() {Installer("ubuntu")},
        "run 3":            func() {Installer("ubuntu")},
        "use 3":            func() {Installer("ubuntu")},
        "exec 3":           func() {Installer("ubuntu")},
        "start 3":          func() {Installer("ubuntu")},
        "launch 3":         func() {Installer("ubuntu")},
        "exploit 3":        func() {Installer("ubuntu")},
        "execute 3":        func() {Installer("ubuntu")},
        "run networks":     func() {Installer("ubuntu")},
        "use networks":     func() {Installer("ubuntu")},
        "exec networks":    func() {Installer("ubuntu")},
        "start networks":   func() {Installer("ubuntu")},
        "launch networks":  func() {Installer("ubuntu")},
        "exploit networks": func() {Installer("ubuntu")},
        "execute networks": func() {Installer("ubuntu")},

        "? 3":              menus.HelpInfoUbuntu,
        "info 3":           menus.HelpInfoUbuntu,
        "help 3":           menus.HelpInfoUbuntu,
        "networks":         menus.HelpInfoUbuntu,
        "info networks":    menus.HelpInfoUbuntu,
        "help networks":    menus.HelpInfoUbuntu,

        "4":                func() {Installer("macos")},
        "run 4":            func() {Installer("macos")},
        "use 4":            func() {Installer("macos")},
        "exec 4":           func() {Installer("macos")},
        "start 4":          func() {Installer("macos")},
        "launch 4":         func() {Installer("macos")},
        "exploit 4":        func() {Installer("macos")},
        "execute 4":        func() {Installer("macos")},
        "run exploits":     func() {Installer("macos")},
        "use exploits":     func() {Installer("macos")},
        "exec exploits":    func() {Installer("macos")},
        "start exploits":   func() {Installer("macos")},
        "launch exploits":  func() {Installer("macos")},
        "exploit exploits": func() {Installer("macos")},
        "execute exploits": func() {Installer("macos")},

        "? 4":              menus.HelpInfoMacos,
        "info 4":           menus.HelpInfoMacos,
        "help 4":           menus.HelpInfoMacos,
        "exploits":         menus.HelpInfoMacos,
        "info exploits":    menus.HelpInfoMacos,
        "help exploits":    menus.HelpInfoMacos,

        "5":                func() {Installer("android")},
        "run 5":            func() {Installer("android")},
        "use 5":            func() {Installer("android")},
        "exec 5":           func() {Installer("android")},
        "start 5":          func() {Installer("android")},
        "launch 5":         func() {Installer("android")},
        "exploit 5":        func() {Installer("android")},
        "execute 5":        func() {Installer("android")},
        "run wireless":     func() {Installer("android")},
        "use wireless":     func() {Installer("android")},
        "exec wireless":    func() {Installer("android")},
        "start wireless":   func() {Installer("android")},
        "launch wireless":  func() {Installer("android")},
        "exploit wireless": func() {Installer("android")},
        "execute wireless": func() {Installer("android")},

        "? 5":              menus.HelpInfoAndroid,
        "info 5":           menus.HelpInfoAndroid,
        "help 5":           menus.HelpInfoAndroid,
        "wireless":         menus.HelpInfoAndroid,
        "info wireless":    menus.HelpInfoAndroid,
        "help wireless":    menus.HelpInfoAndroid,

        "6":                func() {Installer("windows")},
        "run 6":            func() {Installer("windows")},
        "use 6":            func() {Installer("windows")},
        "exec 6":           func() {Installer("windows")},
        "start 6":          func() {Installer("windows")},
        "launch 6":         func() {Installer("windows")},
        "exploit 6":        func() {Installer("windows")},
        "execute 6":        func() {Installer("windows")},
        "run crackers":     func() {Installer("windows")},
        "use crackers":     func() {Installer("windows")},
        "exec crackers":    func() {Installer("windows")},
        "start crackers":   func() {Installer("windows")},
        "launch crackers":  func() {Installer("windows")},
        "exploit crackers": func() {Installer("windows")},
        "execute crackers": func() {Installer("windows")},

        "? 6":              menus.HelpInfoWindows,
        "info 6":           menus.HelpInfoWindows,
        "help 6":           menus.HelpInfoWindows,
        "crackers":         menus.HelpInfoWindows,
        "info crackers":    menus.HelpInfoWindows,
        "help crackers":    menus.HelpInfoWindows,

        "7":                func() {SetupsModules("update")},
        "run 7":            func() {SetupsModules("update")},
        "use 7":            func() {SetupsModules("update")},
        "exec 7":           func() {SetupsModules("update")},
        "start 7":          func() {SetupsModules("update")},
        "launch 7":         func() {SetupsModules("update")},
        "exploit 7":        func() {SetupsModules("update")},
        "execute 7":        func() {SetupsModules("update")},
        "run phishers":     func() {SetupsModules("update")},
        "use phishers":     func() {SetupsModules("update")},
        "exec phishers":    func() {SetupsModules("update")},
        "start phishers":   func() {SetupsModules("update")},
        "launch phishers":  func() {SetupsModules("update")},
        "exploit phishers": func() {SetupsModules("update")},
        "execute phishers": func() {SetupsModules("update")},

        "? 7":              menus.HelpInfoUpdate,
        "info 7":           menus.HelpInfoUpdate,
        "help 7":           menus.HelpInfoUpdate,
        "phishers":         menus.HelpInfoUpdate,
        "info phishers":    menus.HelpInfoUpdate,
        "help phishers":    menus.HelpInfoUpdate,

        "8":                menus.UpsentTools,
        "run 8":            menus.UpsentTools,
        "use 8":            menus.UpsentTools,
        "exec 8":           menus.UpsentTools,
        "start 8":          menus.UpsentTools,
        "launch 8":         menus.UpsentTools,
        "exploit 8":        menus.UpsentTools,
        "execute 8":        menus.UpsentTools,
        "run websites":     menus.UpsentTools,
        "use websites":     menus.UpsentTools,
        "exec websites":    menus.UpsentTools,
        "start websites":   menus.UpsentTools,
        "launch websites":  menus.UpsentTools,
        "exploit websites": menus.UpsentTools,
        "execute websites": menus.UpsentTools,

         "? 8":             menus.HelpInfoAuto,
        "info 8":           menus.HelpInfoAuto,
        "help 8":           menus.HelpInfoAuto,
        "websites":         menus.HelpInfoAuto,
        "info websites":    menus.HelpInfoAuto,
        "help websites":    menus.HelpInfoAuto,

        "9":               func() {SetupsModules("uninstall")},
        "run 9":           func() {SetupsModules("uninstall")},
        "use 9":           func() {SetupsModules("uninstall")},
        "exec 9":          func() {SetupsModules("uninstall")},
        "start 9":         func() {SetupsModules("uninstall")},
        "launch 9":        func() {SetupsModules("uninstall")},
        "exploit 9":       func() {SetupsModules("uninstall")},
        "execute 9":       func() {SetupsModules("uninstall")},
        "run credits":     func() {SetupsModules("uninstall")},
        "use credits":     func() {SetupsModules("uninstall")},
        "exec credits":    func() {SetupsModules("uninstall")},
        "start credits":   func() {SetupsModules("uninstall")},
        "launch credits":  func() {SetupsModules("uninstall")},
        "exploit credits": func() {SetupsModules("uninstall")},
        "execute credits": func() {SetupsModules("uninstall")},

        "? 9":              menus.HelpInfoUninstaller,
        "info 9":           menus.HelpInfoUninstaller,
        "help 9":           menus.HelpInfoUninstaller,
        "credits":          menus.HelpInfoUninstaller,
        "info credits":     menus.HelpInfoUninstaller,
        "help credits":     menus.HelpInfoUninstaller,

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
    if action, exists := commandMap[cmd]; exists {
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
        "distro": &userDistro,
        "module": &userModule,
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
        "distro": &userDistro,
        "module": &userModule,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
    } else {
        menus.HelpInfoSet()
    }
}

func executeModule() {
    if userModule == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter MODULE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    if userDistro == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    SetupsModules(userModule, userDistro)
}

func SetupsModules(module string, args ...interface{}) {
    fmt.Printf("\nDISTRO => %s\nMODULE => %s\n", userDistro, module)
    if userProxy != "" {
        fmt.Printf("PROXIES => %s\n", userProxy)
        utils.SetProxy(userProxy)
    }

    commands := map[string]func() {
        "install": func() {Installer(userDistro)},
         "update": func() {UpdateAfricana()},
         "repair": func() {UpdateAfricana()},
      "uninstall": func() {Uninstaller()},
           //"auto": ,
    }

    if action, exists := commands[module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid module %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, module, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

func Installer(userDistro string) {
    switch userDistro {
    case "1", "kali":
        KaliSetups()
    case "2", "arch":
        ArchSetups()
    case "3", "ubuntu":
        UbuntuSetups()
    case "4", "macos":
        MacosSetups()
    case "5", "android":
        AndroidSetups()
    case "6", "windows":
        WindowsSetups()
    default:
        fmt.Printf("%s[!] Error: %sInvalid DISTRO %s. Use %s'help' %sfor commands.\n", bcolors.RED, bcolors.ENDC, userDistro, bcolors.DARKGREEN, bcolors.ENDC)
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
    fmt.Printf("%s[!] %safricana already installed. Updating it ...\n", bcolors.YELLOW, bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base; git pull .`)
    subprocess.Popen(`cd /root/.afr3/; git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1; cd ./africana-framework; make; mv africana-linux /usr/local/bin/africana; rm -rf ../africana-framework`)
    fmt.Printf("%s[*] %sAfricana succesfully updated ...\n", bcolors.GREEN, bcolors.ENDC)
    return
}

func KaliSetups() {
    fmt.Printf("%s[>] %sInstalling africana in kali\n", bcolors.YELLOW, bcolors.ENDC)
    if _, err := os.Stat(userToolsDir); os.IsNotExist(err) {
        fmt.Printf("%s[>] %sInstalling foundation tools\n", bcolors.YELLOW, bcolors.ENDC)
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
            `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
            `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f userWordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("%s[>] %sInstalling third party tools\n", bcolors.YELLOW, bcolors.ENDC)
        InstallGithubTools()
        fmt.Printf("%s[*] %sAfricana succesfully installed ...\n", bcolors.GREEN, bcolors.ENDC)
    } else {
        UpdateAfricana()
        return
    }
}

func UbuntuSetups() {
    fmt.Printf("%s[>] %sInstalling africana in ubuntu\n", bcolors.YELLOW, bcolors.ENDC)
    if _, err := os.Stat(userToolsDir); os.IsNotExist(err) {
        fmt.Printf("%s[>] %sInstalling foundation tools\n", bcolors.YELLOW, bcolors.ENDC)
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc`,
            `echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref`,
            `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
            `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
            `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f userWordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("%s[>] %sInstalling third party tools\n", bcolors.YELLOW, bcolors.ENDC)
        InstallGithubTools()
        fmt.Printf("%s[*] %sAfricana succesfully installed ...\n", bcolors.GREEN, bcolors.ENDC)
    } else {
        UpdateAfricana()
    }
}

func ArchSetups() {
    fmt.Printf("%s[>] %sInstalling africana in arch\n", bcolors.YELLOW, bcolors.ENDC)
    if _, err := os.Stat(userToolsDir); os.IsNotExist(err) {
        fmt.Printf("%s[>] %sInstalling foundation tools\n", bcolors.YELLOW, bcolors.ENDC)
        foundationCommands := []string{
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm zsh git curl wget go`,
        }
        InstallFoundationTools(foundationCommands)
        BlackArchSetups()
        fmt.Printf("%s[>] %sInstalling third party tools\n", bcolors.YELLOW, bcolors.ENDC)
        InstallGithubTools()
        fmt.Printf("%s[*] %sAfricana succesfully installed ...\n", bcolors.GREEN, bcolors.ENDC)
    } else {
        UpdateAfricana()
    }
}

func BlackArchSetups() {
    fmt.Printf("%s[>] %sInstalling blackarch tools ...\n", bcolors.YELLOW, bcolors.ENDC)
    if _, err := os.Stat(userToolsDir); os.IsNotExist(err) {
        fmt.Printf("%s[>] %sInstalling foundation tools ...\n", bcolors.YELLOW, bcolors.ENDC)
        foundationCommands := []string{
            `curl -O https://blackarch.org/strap.sh`,
            `chmod +x strap.sh`,
            `./strap.sh`,
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm blackarch`,
            `pacman -S --noconfirm base-devel bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f userWordList wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        InstallFoundationTools(foundationCommands)
        fmt.Printf("%s[>] %sInstalling third party tools ...\n", bcolors.YELLOW, bcolors.ENDC)
        InstallGithubTools()
        fmt.Printf("%s[*] %sAfricana fully installed ...\n", bcolors.GREEN, bcolors.ENDC)
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
    fmt.Printf("%s[!] %sUninstalling africana ...\n", bcolors.RED, bcolors.ENDC)
    if _, err := os.Stat(userToolsDir); !os.IsNotExist(err) {
        subprocess.Popen(`rm -rf /root/.afr3/; rm -rf /usr/local/bin/africana`)
        fmt.Printf("%s[*] %sAfricana uninstalled. Goodbye! ...", bcolors.GREEN, bcolors.ENDC)
        os.Exit(0)
    } else {
        fmt.Printf("%s[!] %sAfricana is not installed ...\n", bcolors.YELLOW, bcolors.ENDC)
    }
}
