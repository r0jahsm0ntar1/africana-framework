package africana

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "setups"
    "credits"
    "strings"
    "banners"
    "bcolors"
    "exploits"
    "wireless"
    "crackers"
    "phishers"
    "networks"
    "subprocess"
    "agreements"
    "scriptures"
    "securities"
    "webattackers"
)

var(
    scanner = bufio.NewScanner(os.Stdin)
    agreementDir = utils.GetAgreementPath()
    proxyURL, filePath, Input, Rhost, Proxy, Xhost, Yhost, Setups, Torsocks, Networks, Exploits, Wireless, Crackers, Phishers, Websites, Credits, Verses, Function string
)

var defaultValues = map[string]string{
    "module": "",
    "proxies": "",
}

func Run() {
    if _, err := os.Stat(agreementDir); os.IsNotExist(err) {
        utils.ClearScreen()
        agreements.Covenant()
        for {
            fmt.Printf("%s%s%safr3%s%s > %s", bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
            scanner.Scan()
            Input := strings.TrimSpace(scanner.Text())
            switch Input {
            case "y", "yes":
                utils.Sealing()
                utils.ClearScreen()
                Genesis()
                return
            case "n", "q", "no", "exit", "quit":
                os.Exit(0)
            default:
                fmt.Printf("Choices are (y|n|yes|no)")
            }
        }
    } else {
        Genesis()
    }
}

//Automation mode
func africanaAutoMode() {
    menus.MenuZero()
    for {
        fmt.Printf("%s%s%safr3%s%s > %s", bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        Input = strings.TrimSpace(scanner.Text())
        buildParts := strings.Fields(strings.ToLower(Input))
        if len(buildParts) == 0 {
            continue
        }

        if executeCommandAuto(Input) {
            continue
        }

        switch buildParts[0] {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            //
        case "set":
            handleSetCommand(buildParts)
            executeFunctionAuto()
        case "unset", "delete":
            handleUnsetCommand(buildParts)
        case "run", "start", "launch", "exploit", "execute":
            executeFunctionAuto()
        default:
            utils.SystemShell(Input)
        }
    }
}

func executeFunctionAuto() {
    if Function == "" {
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    africanaFrameworAuto()
}

func autoExecuteFuncAuto() {
    //Distro = distro
    //Function = function
    executeFunctionAuto()
}

func africanaFrameworAuto() {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }
    commands := map[string]func(){
        "setups":   func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
        "torsocks": func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
        "networks": func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
        "exploits": func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
        "wireless": func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
        "crackers": func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
        "phishers": func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
        "websites": func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
        "credits":  func() {credits.Creditors()},
        "verses":   func() {scriptures.ScriptureNarators()},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sFunction %s is invalid. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, Function, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

func executeCommandAuto(cmd string) bool {
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

    //"run":              menus.HelpInfoRun,
    //"use":              menus.HelpInfoRun,
    //"exec":             menus.HelpInfoRun,
    //"start":            menus.HelpInfoRun,
    //"launch":           menus.HelpInfoRun,
    //"exploit":          menus.HelpInfoRun,
    //"execute":          menus.HelpInfoRun,

    //Chameleons//
    "info":             menus.HelpInfoMain,

    "m":                menus.MenuZero,
    "menu":             menus.MenuZero,

    "option":           menus.MainOptions,
    "options":          menus.MainOptions,
    "show option":      menus.MainOptions,
    "show options":     menus.MainOptions,

    "show all":         menus.ListMainFunctions,
    "list all":         menus.ListMainFunctions,

    "func":             menus.ListMainFunctions,
    "funcs":            menus.ListMainFunctions,
    "functions":        menus.ListMainFunctions,
    "show func":        menus.ListMainFunctions,
    "list funcs":       menus.ListMainFunctions,
    "show funcs":       menus.ListMainFunctions,
    "show function":    menus.ListMainFunctions,
    "list function":    menus.ListMainFunctions,
    "list functions":   menus.ListMainFunctions,
    "show functions":   menus.ListMainFunctions,

    "module":           menus.ListMainFunctions,
    "modules":          menus.ListMainFunctions,
    "list module":      menus.ListMainFunctions,
    "show module":      menus.ListMainFunctions,
    "list modules":     menus.ListMainFunctions,
    "show modules":     menus.ListMainFunctions,

    "1":                func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "run 1":            func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "use 1":            func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "exec 1":           func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "start 1":          func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "launch 1":         func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "exploit 1":        func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "execute 1":        func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "run setups":       func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "use setups":       func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "exec setups":      func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "start setups":     func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "launch setups":    func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "exploit setups":   func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
    "execute setups":   func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},

    "? 1":              menus.HelpInfoSetups,
    "info 1":           menus.HelpInfoSetups,
    "help 1":           menus.HelpInfoSetups,
    "setups":           menus.HelpInfoSetups,
    "info setups":      menus.HelpInfoSetups,
    "help setups":      menus.HelpInfoSetups,

    "2":                func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "run 2":            func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "use 2":            func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "exec 2":           func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "start 2":          func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "launch 2":         func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "exploit 2":        func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "execute 2":        func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "run torsocks":     func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "use torsocks":     func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "exec torsocks":    func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "start torsocks":   func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "launch torsocks":  func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "exploit torsocks": func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
    "execute torsocks": func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},

    "? 2":              menus.HelpInfoTorsocks,
    "info 2":           menus.HelpInfoTorsocks,
    "help 2":           menus.HelpInfoTorsocks,
    "torsocks":         menus.HelpInfoTorsocks,
    "info torsocks":    menus.HelpInfoTorsocks,
    "help torsocks":    menus.HelpInfoTorsocks,

    "3":                func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "run 3":            func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "use 3":            func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "exec 3":           func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "start 3":          func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "launch 3":         func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "exploit 3":        func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "execute 3":        func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "run networks":     func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "use networks":     func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "exec networks":    func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "start networks":   func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "launch networks":  func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "exploit networks": func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
    "execute networks": func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},

    "? 3":              menus.HelpInfoNetworks,
    "info 3":           menus.HelpInfoNetworks,
    "help 3":           menus.HelpInfoNetworks,
    "networks":         menus.HelpInfoNetworks,
    "info networks":    menus.HelpInfoNetworks,
    "help networks":    menus.HelpInfoNetworks,

    "4":                func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "run 4":            func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "use 4":            func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "exec 4":           func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "start 4":          func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "launch 4":         func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "exploit 4":        func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "execute 4":        func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "run exploits":     func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "use exploits":     func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "exec exploits":    func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "start exploits":   func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "launch exploits":  func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "exploit exploits": func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
    "execute exploits": func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},

    "? 4":              menus.HelpInfoExploits,
    "info 4":           menus.HelpInfoExploits,
    "help 4":           menus.HelpInfoExploits,
    "exploits":         menus.HelpInfoExploits,
    "info exploits":    menus.HelpInfoExploits,
    "help exploits":    menus.HelpInfoExploits,

    "5":                func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "run 5":            func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "use 5":            func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "exec 5":           func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "start 5":          func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "launch 5":         func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "exploit 5":        func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "execute 5":        func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "run wireless":     func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "use wireless":     func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "exec wireless":    func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "start wireless":   func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "launch wireless":  func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "exploit wireless": func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
    "execute wireless": func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},

    "? 5":              menus.HelpInfoWireless,
    "info 5":           menus.HelpInfoWireless,
    "help 5":           menus.HelpInfoWireless,
    "wireless":         menus.HelpInfoWireless,
    "info wireless":    menus.HelpInfoWireless,
    "help wireless":    menus.HelpInfoWireless,

    "6":                func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "run 6":            func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "use 6":            func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "exec 6":           func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "start 6":          func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "launch 6":         func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "exploit 6":        func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "execute 6":        func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "run crackers":     func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "use crackers":     func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "exec crackers":    func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "start crackers":   func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "launch crackers":  func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "exploit crackers": func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
    "execute crackers": func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},

    "? 6":              menus.HelpInfoCrackers,
    "info 6":           menus.HelpInfoCrackers,
    "help 6":           menus.HelpInfoCrackers,
    "crackers":         menus.HelpInfoCrackers,
    "info crackers":    menus.HelpInfoCrackers,
    "help crackers":    menus.HelpInfoCrackers,

    "7":                func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "run 7":            func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "use 7":            func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "exec 7":           func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "start 7":          func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "launch 7":         func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "exploit 7":        func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "execute 7":        func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "run phishers":     func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "use phishers":     func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "exec phishers":    func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "start phishers":   func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "launch phishers":  func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "exploit phishers": func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
    "execute phishers": func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},

    "? 7":              menus.HelpInfoPhishers,
    "info 7":           menus.HelpInfoPhishers,
    "help 7":           menus.HelpInfoPhishers,
    "phishers":         menus.HelpInfoPhishers,
    "info phishers":    menus.HelpInfoPhishers,
    "help phishers":    menus.HelpInfoPhishers,

    "8":                func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "run 8":            func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "use 8":            func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "exec 8":           func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "start 8":          func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "launch 8":         func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "exploit 8":        func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "execute 8":        func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "run websites":     func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "use websites":     func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "exec websites":    func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "start websites":   func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "launch websites":  func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "exploit websites": func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
    "execute websites": func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},

    "? 8":              menus.HelpInfoWebsites,
    "info 8":           menus.HelpInfoWebsites,
    "help 8":           menus.HelpInfoWebsites,
    "websites":         menus.HelpInfoWebsites,
    "info websites":    menus.HelpInfoWebsites,
    "help websites":    menus.HelpInfoWebsites,

    "9":               credits.Creditors,
    "run 9":           credits.Creditors,
    "use 9":           credits.Creditors,
    "exec 9":          credits.Creditors,
    "start 9":         credits.Creditors,
    "launch 9":        credits.Creditors,
    "exploit 9":       credits.Creditors,
    "execute 9":       credits.Creditors,
    "run credits":     credits.Creditors,
    "use credits":     credits.Creditors,
    "exec credits":    credits.Creditors,
    "start credits":   credits.Creditors,
    "launch credits":  credits.Creditors,
    "exploit credits": credits.Creditors,
    "execute credits": credits.Creditors,

    "? 9":              menus.HelpInfoCredits,
    "info 9":           menus.HelpInfoCredits,
    "help 9":           menus.HelpInfoCredits,
    "credits":          menus.HelpInfoCredits,
    "info credits":     menus.HelpInfoCredits,
    "help credits":     menus.HelpInfoCredits,

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


//Manual mode
func africanaManualMode() {
    for {
        fmt.Printf("%s%s%safr3%s%s > %s", bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
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
            //
        case "set":
            handleSetCommand(buildParts)
            executeFunction()
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

    //"run":              menus.HelpInfoRun,
    //"use":              menus.HelpInfoRun,
    //"exec":             menus.HelpInfoRun,
    //"start":            menus.HelpInfoRun,
    //"launch":           menus.HelpInfoRun,
    //"exploit":          menus.HelpInfoRun,
    //"execute":          menus.HelpInfoRun,

    //Chameleons//
    "info":             menus.HelpInfoMain,

    "m":                menus.MenuZero,
    "menu":             menus.MenuZero,

    "option":           menus.MainOptions,
    "options":          menus.MainOptions,
    "show option":      menus.MainOptions,
    "show options":     menus.MainOptions,

    "show all":         menus.ListMainFunctions,
    "list all":         menus.ListMainFunctions,

    "func":             menus.ListMainFunctions,
    "funcs":            menus.ListMainFunctions,
    "functions":        menus.ListMainFunctions,
    "show func":        menus.ListMainFunctions,
    "list funcs":       menus.ListMainFunctions,
    "show funcs":       menus.ListMainFunctions,
    "show function":    menus.ListMainFunctions,
    "list function":    menus.ListMainFunctions,
    "list functions":   menus.ListMainFunctions,
    "show functions":   menus.ListMainFunctions,

    "module":           menus.ListMainFunctions,
    "modules":          menus.ListMainFunctions,
    "list module":      menus.ListMainFunctions,
    "show module":      menus.ListMainFunctions,
    "list modules":     menus.ListMainFunctions,
    "show modules":     menus.ListMainFunctions,

    "1":                func() {setups.SetupsLauncher()},
    "run 1":            func() {setups.SetupsLauncher()},
    "use 1":            func() {setups.SetupsLauncher()},
    "exec 1":           func() {setups.SetupsLauncher()},
    "start 1":          func() {setups.SetupsLauncher()},
    "launch 1":         func() {setups.SetupsLauncher()},
    "exploit 1":        func() {setups.SetupsLauncher()},
    "execute 1":        func() {setups.SetupsLauncher()},
    "run setups":       func() {setups.SetupsLauncher()},
    "use setups":       func() {setups.SetupsLauncher()},
    "exec setups":      func() {setups.SetupsLauncher()},
    "start setups":     func() {setups.SetupsLauncher()},
    "launch setups":    func() {setups.SetupsLauncher()},
    "exploit setups":   func() {setups.SetupsLauncher()},
    "execute setups":   func() {setups.SetupsLauncher()},

    "? 1":              menus.HelpInfoSetups,
    "info 1":           menus.HelpInfoSetups,
    "help 1":           menus.HelpInfoSetups,
    "setups":           menus.HelpInfoSetups,
    "info setups":      menus.HelpInfoSetups,
    "help setups":      menus.HelpInfoSetups,

    "2":                func() {securities.Torsocks()},
    "run 2":            func() {securities.Torsocks()},
    "use 2":            func() {securities.Torsocks()},
    "exec 2":           func() {securities.Torsocks()},
    "start 2":          func() {securities.Torsocks()},
    "launch 2":         func() {securities.Torsocks()},
    "exploit 2":        func() {securities.Torsocks()},
    "execute 2":        func() {securities.Torsocks()},
    "run torsocks":     func() {securities.Torsocks()},
    "use torsocks":     func() {securities.Torsocks()},
    "exec torsocks":    func() {securities.Torsocks()},
    "start torsocks":   func() {securities.Torsocks()},
    "launch torsocks":  func() {securities.Torsocks()},
    "exploit torsocks": func() {securities.Torsocks()},
    "execute torsocks": func() {securities.Torsocks()},

    "? 2":              menus.HelpInfoTorsocks,
    "info 2":           menus.HelpInfoTorsocks,
    "help 2":           menus.HelpInfoTorsocks,
    "torsocks":         menus.HelpInfoTorsocks,
    "info torsocks":    menus.HelpInfoTorsocks,
    "help torsocks":    menus.HelpInfoTorsocks,

    "3":                func() {networks.NetworksPentest()},
    "run 3":            func() {networks.NetworksPentest()},
    "use 3":            func() {networks.NetworksPentest()},
    "exec 3":           func() {networks.NetworksPentest()},
    "start 3":          func() {networks.NetworksPentest()},
    "launch 3":         func() {networks.NetworksPentest()},
    "exploit 3":        func() {networks.NetworksPentest()},
    "execute 3":        func() {networks.NetworksPentest()},
    "run networks":     func() {networks.NetworksPentest()},
    "use networks":     func() {networks.NetworksPentest()},
    "exec networks":    func() {networks.NetworksPentest()},
    "start networks":   func() {networks.NetworksPentest()},
    "launch networks":  func() {networks.NetworksPentest()},
    "exploit networks": func() {networks.NetworksPentest()},
    "execute networks": func() {networks.NetworksPentest()},

    "? 3":              menus.HelpInfoNetworks,
    "info 3":           menus.HelpInfoNetworks,
    "help 3":           menus.HelpInfoNetworks,
    "networks":         menus.HelpInfoNetworks,
    "info networks":    menus.HelpInfoNetworks,
    "help networks":    menus.HelpInfoNetworks,

    "4":                func() {exploits.ExploitsPentest()},
    "run 4":            func() {exploits.ExploitsPentest()},
    "use 4":            func() {exploits.ExploitsPentest()},
    "exec 4":           func() {exploits.ExploitsPentest()},
    "start 4":          func() {exploits.ExploitsPentest()},
    "launch 4":         func() {exploits.ExploitsPentest()},
    "exploit 4":        func() {exploits.ExploitsPentest()},
    "execute 4":        func() {exploits.ExploitsPentest()},
    "run exploits":     func() {exploits.ExploitsPentest()},
    "use exploits":     func() {exploits.ExploitsPentest()},
    "exec exploits":    func() {exploits.ExploitsPentest()},
    "start exploits":   func() {exploits.ExploitsPentest()},
    "launch exploits":  func() {exploits.ExploitsPentest()},
    "exploit exploits": func() {exploits.ExploitsPentest()},
    "execute exploits": func() {exploits.ExploitsPentest()},

    "? 4":              menus.HelpInfoExploits,
    "info 4":           menus.HelpInfoExploits,
    "help 4":           menus.HelpInfoExploits,
    "exploits":         menus.HelpInfoExploits,
    "info exploits":    menus.HelpInfoExploits,
    "help exploits":    menus.HelpInfoExploits,

    "5":                func() {wireless.WirelessPentest()},
    "run 5":            func() {wireless.WirelessPentest()},
    "use 5":            func() {wireless.WirelessPentest()},
    "exec 5":           func() {wireless.WirelessPentest()},
    "start 5":          func() {wireless.WirelessPentest()},
    "launch 5":         func() {wireless.WirelessPentest()},
    "exploit 5":        func() {wireless.WirelessPentest()},
    "execute 5":        func() {wireless.WirelessPentest()},
    "run wireless":     func() {wireless.WirelessPentest()},
    "use wireless":     func() {wireless.WirelessPentest()},
    "exec wireless":    func() {wireless.WirelessPentest()},
    "start wireless":   func() {wireless.WirelessPentest()},
    "launch wireless":  func() {wireless.WirelessPentest()},
    "exploit wireless": func() {wireless.WirelessPentest()},
    "execute wireless": func() {wireless.WirelessPentest()},

    "? 5":              menus.HelpInfoWireless,
    "info 5":           menus.HelpInfoWireless,
    "help 5":           menus.HelpInfoWireless,
    "wireless":         menus.HelpInfoWireless,
    "info wireless":    menus.HelpInfoWireless,
    "help wireless":    menus.HelpInfoWireless,

    "6":                func() {crackers.CrackersPentest()},
    "run 6":            func() {crackers.CrackersPentest()},
    "use 6":            func() {crackers.CrackersPentest()},
    "exec 6":           func() {crackers.CrackersPentest()},
    "start 6":          func() {crackers.CrackersPentest()},
    "launch 6":         func() {crackers.CrackersPentest()},
    "exploit 6":        func() {crackers.CrackersPentest()},
    "execute 6":        func() {crackers.CrackersPentest()},
    "run crackers":     func() {crackers.CrackersPentest()},
    "use crackers":     func() {crackers.CrackersPentest()},
    "exec crackers":    func() {crackers.CrackersPentest()},
    "start crackers":   func() {crackers.CrackersPentest()},
    "launch crackers":  func() {crackers.CrackersPentest()},
    "exploit crackers": func() {crackers.CrackersPentest()},
    "execute crackers": func() {crackers.CrackersPentest()},

    "? 6":              menus.HelpInfoCrackers,
    "info 6":           menus.HelpInfoCrackers,
    "help 6":           menus.HelpInfoCrackers,
    "crackers":         menus.HelpInfoCrackers,
    "info crackers":    menus.HelpInfoCrackers,
    "help crackers":    menus.HelpInfoCrackers,

    "7":                func() {phishers.PhishingPentest()},
    "run 7":            func() {phishers.PhishingPentest()},
    "use 7":            func() {phishers.PhishingPentest()},
    "exec 7":           func() {phishers.PhishingPentest()},
    "start 7":          func() {phishers.PhishingPentest()},
    "launch 7":         func() {phishers.PhishingPentest()},
    "exploit 7":        func() {phishers.PhishingPentest()},
    "execute 7":        func() {phishers.PhishingPentest()},
    "run phishers":     func() {phishers.PhishingPentest()},
    "use phishers":     func() {phishers.PhishingPentest()},
    "exec phishers":    func() {phishers.PhishingPentest()},
    "start phishers":   func() {phishers.PhishingPentest()},
    "launch phishers":  func() {phishers.PhishingPentest()},
    "exploit phishers": func() {phishers.PhishingPentest()},
    "execute phishers": func() {phishers.PhishingPentest()},

    "? 7":              menus.HelpInfoPhishers,
    "info 7":           menus.HelpInfoPhishers,
    "help 7":           menus.HelpInfoPhishers,
    "phishers":         menus.HelpInfoPhishers,
    "info phishers":    menus.HelpInfoPhishers,
    "help phishers":    menus.HelpInfoPhishers,

    "8":                func() {webattackers.WebsitesPentest()},
    "run 8":            func() {webattackers.WebsitesPentest()},
    "use 8":            func() {webattackers.WebsitesPentest()},
    "exec 8":           func() {webattackers.WebsitesPentest()},
    "start 8":          func() {webattackers.WebsitesPentest()},
    "launch 8":         func() {webattackers.WebsitesPentest()},
    "exploit 8":        func() {webattackers.WebsitesPentest()},
    "execute 8":        func() {webattackers.WebsitesPentest()},
    "run websites":     func() {webattackers.WebsitesPentest()},
    "use websites":     func() {webattackers.WebsitesPentest()},
    "exec websites":    func() {webattackers.WebsitesPentest()},
    "start websites":   func() {webattackers.WebsitesPentest()},
    "launch websites":  func() {webattackers.WebsitesPentest()},
    "exploit websites": func() {webattackers.WebsitesPentest()},
    "execute websites": func() {webattackers.WebsitesPentest()},

    "? 8":              menus.HelpInfoWebsites,
    "info 8":           menus.HelpInfoWebsites,
    "help 8":           menus.HelpInfoWebsites,
    "websites":         menus.HelpInfoWebsites,
    "info websites":    menus.HelpInfoWebsites,
    "help websites":    menus.HelpInfoWebsites,

    "9":               credits.Creditors,
    "run 9":           credits.Creditors,
    "use 9":           credits.Creditors,
    "exec 9":          credits.Creditors,
    "start 9":         credits.Creditors,
    "launch 9":        credits.Creditors,
    "exploit 9":       credits.Creditors,
    "execute 9":       credits.Creditors,
    "run credits":     credits.Creditors,
    "use credits":     credits.Creditors,
    "exec credits":    credits.Creditors,
    "start credits":   credits.Creditors,
    "launch credits":  credits.Creditors,
    "exploit credits": credits.Creditors,
    "execute credits": credits.Creditors,

    "? 9":              menus.HelpInfoCredits,
    "info 9":           menus.HelpInfoCredits,
    "help 9":           menus.HelpInfoCredits,
    "credits":          menus.HelpInfoCredits,
    "info credits":     menus.HelpInfoCredits,
    "help credits":     menus.HelpInfoCredits,

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

        "func": &Function,
        "setups": &Setups,
        "verses": &Verses,
        "credits": &Credits,
        "module": &Function,
        "function": &Function,
        "torsocks": &Torsocks,
        "networks": &Networks,
        "exploits": &Exploits,
        "wireless": &Wireless,
        "crackers": &Crackers,
        "phishers": &Phishers,
        "websites": &Websites,

    }
    if ptr, exists := setValues[key]; exists {
        *ptr = value
        //fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
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

        "func": &Function,
        "setups": &Setups,
        "verses": &Verses,
        "credits": &Credits,
        "module": &Function,
        "function": &Function,
        "torsocks": &Torsocks,
        "networks": &Networks,
        "exploits": &Exploits,
        "wireless": &Wireless,
        "crackers": &Crackers,
        "phishers": &Phishers,
        "websites": &Websites,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        //fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
    } else {
        menus.HelpInfoSet()
    }
}

func executeFunction() {
    if Function == "" {
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    africanaManual()
}

// Helper modules
func autoExecuteFunc() {
    //Distro = distro
    //Function = function
    executeFunction()
}

func africanaManual() {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }
    commands := map[string]func(){
        "setups":   func() {setups.SetupsLauncher()},
        "torsocks": func() {securities.Torsocks()},
        "networks": func() {networks.NetworksPentest()},
        "exploits": func() {exploits.ExploitsPentest()},
        "wireless": func() {wireless.WirelessPentest()},
        "crackers": func() {crackers.CrackersPentest()},
        "phishers": func() {phishers.PhishingPentest()},
        "websites": func() {webattackers.WebsitesPentest()},
        "credits":  func() {credits.Creditors()},
        "verses":   func() {scriptures.ScriptureNarators()},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sFunction %s is invalid. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, Function, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

func Genesis() {
    if len(os.Args) < 2 {
        utils.InitiLize()
        scriptures.Verse()
        setups.CheckTools()
        banners.GraphicsLarge()
        banners.Version()
        africanaManualMode()
        return
    }

    commands := map[string]func(){
        "-v":               banners.Version,
        "--version":        banners.Version,
        "-u":               setups.UpdateAfricana,
        "--update":         setups.UpdateAfricana,
        "-0":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); africanaAutoMode()},
        "-a":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); africanaAutoMode()},
        "--auto":           func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); africanaAutoMode()},
        "-1":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuOne(); setups.SetupsLauncher(); africanaAutoMode()},
        "-i":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuOne(); setups.SetupsLauncher(); africanaAutoMode()},
        "--install":        func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuOne(); setups.SetupsLauncher(); africanaAutoMode()},
        "-2":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuTwo(); securities.Torsocks(); africanaAutoMode()},
        "-t":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuTwo(); securities.Torsocks(); africanaAutoMode()},
        "--torsocks":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuTwo(); securities.Torsocks(); africanaAutoMode()},
        "-3":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuThree(); networks.NetworksPentest(); africanaAutoMode()},
        "-n":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuThree(); networks.NetworksPentest(); africanaAutoMode()},
        "--networks":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuThree(); networks.NetworksPentest(); africanaAutoMode()},
        "-4":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFour(); exploits.ExploitsPentest(); africanaAutoMode()},
        "-e":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFour(); exploits.ExploitsPentest(); africanaAutoMode()},
        "--exploits":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFour(); exploits.ExploitsPentest(); africanaAutoMode()},
        "-5":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFive(); wireless.WirelessPentest(); africanaAutoMode()},
        "-w":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFive(); wireless.WirelessPentest(); africanaAutoMode()},
        "--wireless":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuFive(); wireless.WirelessPentest(); africanaAutoMode()},
        "-6":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSix(); crackers.CrackersPentest(); africanaAutoMode()},
        "-p":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSix(); crackers.CrackersPentest(); africanaAutoMode()},
        "--crackers":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSix(); crackers.CrackersPentest(); africanaAutoMode()},
        "-7":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSeven(); phishers.PhishingPentest(); africanaAutoMode()},
        "-f":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSeven(); phishers.PhishingPentest(); africanaAutoMode()},
        "--phishers":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuSeven(); phishers.PhishingPentest(); africanaAutoMode()},
        "-8":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuEight(); webattackers.WebsitesPentest(); africanaAutoMode()},
        "-x":               func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuEight(); webattackers.WebsitesPentest(); africanaAutoMode()},
        "--websites":       func() {scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge(); menus.MenuEight(); webattackers.WebsitesPentest(); africanaAutoMode()},
        "-9":               func() {credits.Creditors(); africanaAutoMode()},
        "-c":               func() {credits.Creditors(); africanaAutoMode()},
        "--credits":        func() {credits.Creditors(); africanaAutoMode()},
        "-99":              func() {scriptures.ScriptureNarators(); africanaAutoMode()},
        "-b":               func() {scriptures.ScriptureNarators(); africanaAutoMode()},
        "--scriptures":     func() {scriptures.ScriptureNarators(); africanaAutoMode()},
        "-g":               func() {banners.RandomBanners(); utils.BrowseTutarilas(); africanaAutoMode()},
        "--guide":          func() {banners.RandomBanners(); utils.BrowseTutarilas(); africanaAutoMode()},
        "-00":              menus.HelpInfoMenuMain,
        "-h":               menus.HelpInfoMenuMain,
        "?":                menus.HelpInfoMenuMain,
        "--help":           menus.HelpInfoMenuMain,
    }

    if action, exists := commands[os.Args[1]]; exists {
        action()
    } else {
        menus.HelpInfoMenuMain()
    }
}
