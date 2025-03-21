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
    "internals"
    "subprocess"
    "agreements"
    "scriptures"
    "securities"
    "webattackers"
)

var(
    proxyURL    string
    filePath    string
    userInput   string
    userRhost   string
    userXhost   string
    userYhost   string
    scanner   = bufio.NewScanner(os.Stdin)
    agreementDir = utils.GetAgreementPath()
)

func Run() {
    utils.ClearScreen()
    if _, err := os.Stat(agreementDir); os.IsNotExist(err) {
        utils.ClearScreen()
        agreements.Covenant()
        for {
            fmt.Printf("%s%safr3%s%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
            scanner.Scan()
            userInput := strings.TrimSpace(strings.ToLower(scanner.Text()))
            switch userInput {
            case "y", "yes":
                utils.UserSealing()
                utils.ClearScreen()
                utils.InitiLize()
                Genesis()
                return
            case "n", "q", "no", "exit", "quit":
                os.Exit(0)
            default:
                fmt.Println("Choices are (y|n|yes|no)")
            }
        }
    } else {
        utils.InitiLize()
        Genesis()
    }
}

func africanaAutomation() {
    menus.MenuZero()
    for {
        fmt.Printf("%s%safr3%s%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        userInput := strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(userInput)

        if len(buildParts) == 0 {
            continue
        }

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

            "run":              menus.HelpInfoRun,
            "use":              menus.HelpInfoRun,
            "exec":             menus.HelpInfoRun,
            "start":            menus.HelpInfoRun,
            "launch":           menus.HelpInfoRun,
            "exploit":          menus.HelpInfoRun,
            "execute":          menus.HelpInfoRun,

            //Chameleons//
            "info":             menus.HelpInfoMain,

            "m":                menus.MenuZero,
            "menu":             menus.MenuZero,

            "option":           menus.HelpInfOptions,
            "options":          menus.HelpInfOptions,
            "show option":      menus.HelpInfOptions,
            "show options":     menus.HelpInfOptions,

            "modules":          menus.ListMainModules,
            "show all":         menus.ListMainModules,
            "list all":         menus.ListMainModules,
            "list modules":     menus.ListMainModules,
            "show modules":     menus.ListMainModules,

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

            "? 1":              menus.HelpInfoPmanager,
            "info 1":           menus.HelpInfoPmanager,
            "help 1":           menus.HelpInfoPmanager,
            "setups":           menus.HelpInfoPmanager,
            "info setups":      menus.HelpInfoPmanager,
            "help setups":      menus.HelpInfoPmanager,

            "2":                func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "run 2":            func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "use 2":            func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "exec 2":           func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "start 2":          func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "launch 2":         func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "exploit 2":        func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "execute 2":        func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "run anonsurf":     func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "use anonsurf":     func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "exec anonsurf":    func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "start anonsurf":   func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "launch anonsurf":  func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "exploit anonsurf": func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
            "execute anonsurf": func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},

            "? 2":              menus.HelpInfoAnonsurf,
            "info 2":           menus.HelpInfoAnonsurf,
            "help 2":           menus.HelpInfoAnonsurf,
            "anonsurf":         menus.HelpInfoAnonsurf,
            "info anonsurf":    menus.HelpInfoAnonsurf,
            "help anonsurf":    menus.HelpInfoAnonsurf,

            "3":                func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "run 3":            func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "use 3":            func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "exec 3":           func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "start 3":          func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "launch 3":         func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "exploit 3":        func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "execute 3":        func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "run networks":     func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "use networks":     func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "exec networks":    func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "start networks":   func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "launch networks":  func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "exploit networks": func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
            "execute networks": func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},

            "? 3":              menus.HelpInfoNetworks,
            "info 3":           menus.HelpInfoNetworks,
            "help 3":           menus.HelpInfoNetworks,
            "networks":         menus.HelpInfoNetworks,
            "info networks":    menus.HelpInfoNetworks,
            "help networks":    menus.HelpInfoNetworks,

            "4":                func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "run 4":            func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "use 4":            func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "exec 4":           func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "start 4":          func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "launch 4":         func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "exploit 4":        func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "execute 4":        func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "run exploits":     func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "use exploits":     func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "exec exploits":    func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "start exploits":   func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "launch exploits":  func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "exploit exploits": func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
            "execute exploits": func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},

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

            "6":                func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "run 6":            func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "use 6":            func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "exec 6":           func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "start 6":          func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "launch 6":         func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "exploit 6":        func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "execute 6":        func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "run crackers":     func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "use crackers":     func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "exec crackers":    func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "start crackers":   func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "launch crackers":  func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "exploit crackers": func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
            "execute crackers": func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},

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

            "8":                func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "run 8":            func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "use 8":            func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "exec 8":           func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "start 8":          func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "launch 8":         func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "exploit 8":        func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "execute 8":        func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "run websites":     func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "use websites":     func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "exec websites":    func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "start websites":   func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "launch websites":  func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "exploit websites": func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
            "execute websites": func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},

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

        if action, exists := commandMap[userInput]; exists {
            action()
            continue
        }
        switch buildParts[0] {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            //
        default:
            utils.SystemShell(userInput)
        }
    }
}

func africanaFramework() {
}

func Genesis() {
    if len(os.Args) < 2 {
        scriptures.Verse(); banners.GraphicsLarge(); africanaAutomation()
        return
    }

    commands := map[string]func(){
        "-v":               banners.Version,
        "--version":        banners.Version,
        "-u":               setups.UpdateAfricana,
        "--update":         setups.UpdateAfricana,
        "-0":               func() {scriptures.Verse(); banners.GraphicsLarge(); africanaAutomation()},
        "-a":               func() {scriptures.Verse(); banners.GraphicsLarge(); africanaAutomation()},
        "--auto":           func() {scriptures.Verse(); banners.GraphicsLarge(); africanaAutomation()},
        "-1":               func() {menus.MenuFour(); setups.SetupsLauncher(); menus.MenuZero()},
        "-i":               func() {menus.MenuFour(); setups.SetupsLauncher(); menus.MenuZero()},
        "--install":        func() {menus.MenuFour(); setups.SetupsLauncher(); menus.MenuZero()},
        "-2":               func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
        "-t":               func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
        "--anonsurf":       func() {menus.MenuTwo(); securities.Anonsurf(); menus.MenuZero()},
        "-3":               func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
        "-n":               func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
        "--networks":       func() {menus.MenuThree(); internals.NetworkPentest(); menus.MenuZero()},
        "-4":               func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
        "-m":               func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
        "--exploits":       func() {menus.MenuFour(); exploits.MalwarePentest(); menus.MenuZero()},
        "-5":               func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
        "-w":               func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
        "--wireless":       func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
        "-6":               func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
        "-p":               func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
        "--crackers":       func() {menus.MenuSix(); crackers.PasswordPentest(); menus.MenuZero()},
        "-7":               func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
        "-f":               func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
        "--phishers":       func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
        "-8":               func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
        "-x":               func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
        "--websites":       func() {menus.MenuEight(); webattackers.WebPentest(); menus.MenuZero()},
        "-9":               credits.Creditors,
        "-c":               credits.Creditors,
        "--credits":        credits.Creditors,
        "-99":              scriptures.ScriptureNarators,
        "-b":               scriptures.ScriptureNarators,
        "--scriptures":     scriptures.ScriptureNarators,
        "-g":               func() {banners.RandomBanners(); utils.BrowseTutarilas()},
        "--guide":          func() {banners.RandomBanners(); utils.BrowseTutarilas()},
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
