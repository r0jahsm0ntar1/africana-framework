//John 3:16

package afrconsole

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "setups"
    "strconv"
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
    proxyURL, filePath, Input, RHost, Proxy, Xhost, Yhost, Setups, Torsocks, Networks, Exploits, Wireless, Crackers, Phishers, Websites, Credits, Verses, Function string
)

var defaultValues = map[string]string{
    "module": "",
    "proxies": "",
}

type stringMatcher struct {
    names  []string
    action func()
}

func Start() {
    if _, err := os.Stat(agreementDir); os.IsNotExist(err) {
        utils.ClearScreen()
        agreements.Covenant()
        for {
            fmt.Printf("%s%s%safr3%s%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
                fmt.Printf("%s[!] %sChoices are (y|n|yes|no)", bcolors.BrightYellow, bcolors.Endc)
            }
        }
    } else {
        Genesis()
    }
}

func africanaAutoMode() {
    menus.MenuZero()
    for {
        fmt.Printf("%s%s%safr3%s%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    africanaFrameworAuto("")
}

func africanaFrameworAuto(Function string, args ...interface{}) {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            //
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
        "credits":  func() {credits.Creditors(); menus.MenuZero()},
        "verses":   func() {scriptures.ScriptureNarators(); menus.MenuZero()},

        "1": func() {menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()},
        "2": func() {menus.MenuTwo(); securities.Torsocks(); menus.MenuZero()},
        "3": func() {menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero()},
        "4": func() {menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero()},
        "5": func() {menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero()},
        "6": func() {menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero()},
        "7": func() {menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero()},
        "8": func() {menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero()},
        "9": func() {credits.Creditors(); menus.MenuZero()},
       "10": func() {scriptures.ScriptureNarators(); menus.MenuZero()},
    }

    textCommands := []string{
        "setups", "torsocks", "networks", "exploits", "wireless", "crackers", "phishers", "websites", "credits", "verses",
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        if num, err := strconv.Atoi(Function); err == nil {
            fmt.Printf("\n%s[!] %sNumber %d is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, num)
            menus.ListMainFunctions()
            return
        }

        lowerInput := strings.ToLower(Function)
        for _, cmd := range textCommands {
            lowerCmd := strings.ToLower(cmd)
            if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
                fmt.Printf("\n%s[!] %sModule '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
                return
            }
        }

        fmt.Printf("\n%s[!] %sModule %s is invalid. Available commands are:\n", bcolors.Yellow, bcolors.Endc, Function)
        menus.ListMainFunctions()
    }
}

func executeCommandAuto(cmd string) bool {
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

        {[]string{"info"}, menus.HelpInfoMain},
        {[]string{"m", "menu"}, menus.MenuZero},
        {[]string{"option", "options", "show option", "show options"}, menus.MainOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListMainFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() {africanaFrameworAuto("setups")}},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, menus.HelpInfoSetups},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run torsocks", "use torsocks", "exec torsocks", "start torsocks", "launch torsocks", "exploit torsocks", "execute torsocks"}, func() {africanaFrameworAuto("torsocks")}},
        {[]string{"? 2", "info 2", "help 2", "torsocks", "info torsocks", "help torsocks"}, menus.HelpInfoTorsocks},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run networks", "use networks", "exec networks", "start networks", "launch networks", "exploit networks", "execute networks"}, func() {africanaFrameworAuto("networks")}},
        {[]string{"? 3", "info 3", "help 3", "networks", "info networks", "help networks"}, menus.HelpInfoNetworks},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run exploits", "use exploits", "exec exploits", "start exploits", "launch exploits", "exploit exploits", "execute exploits"}, func() {africanaFrameworAuto("exploits")}},
        {[]string{"? 4", "info 4", "help 4", "exploits", "info exploits", "help exploits"}, menus.HelpInfoExploits},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wireless", "use wireless", "exec wireless", "start wireless", "launch wireless", "exploit wireless", "execute wireless"}, func() {africanaFrameworAuto("wireless")}},
        {[]string{"? 5", "info 5", "help 5", "wireless", "info wireless", "help wireless"}, menus.HelpInfoWireless},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run crackers", "use crackers", "exec crackers", "start crackers", "launch crackers", "exploit crackers", "execute crackers"}, func() {africanaFrameworAuto("crackers")}},
        {[]string{"? 6", "info 6", "help 6", "crackers", "info crackers", "help crackers"}, menus.HelpInfoCrackers},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run phishers", "use phishers", "exec phishers", "start phishers", "launch phishers", "exploit phishers", "execute phishers"}, func() {africanaFrameworAuto("phishers")}},
        {[]string{"? 7", "info 7", "help 7", "phishers", "info phishers", "help phishers"}, menus.HelpInfoPhishers},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run websites", "use websites", "exec websites", "start websites", "launch websites", "exploit websites", "execute websites"}, func() {africanaFrameworAuto("websites")}},
        {[]string{"? 8", "info 8", "help 8", "websites", "info websites", "help websites"}, menus.HelpInfoWebsites},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run credits", "use credits", "exec credits", "start credits", "launch credits", "exploit credits", "execute credits"}, func() {africanaFrameworAuto("credits")}},
        {[]string{"? 9", "info 9", "help 9", "credits", "info credits", "help credits"}, menus.HelpInfoCredits},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, func() {africanaManual("verses")}},
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

func africanaManualMode() {
    for {
        fmt.Printf("%s%s%safr3%s%s > %s", bcolors.Endc, bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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

        {[]string{"info"}, menus.HelpInfoMain},
        {[]string{"m", "menu"}, menus.MenuZero},
        {[]string{"option", "options", "show option", "show options"}, menus.MainOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListMainFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() {africanaManual("setups")}},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, menus.HelpInfoSetups},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run torsocks", "use torsocks", "exec torsocks", "start torsocks", "launch torsocks", "exploit torsocks", "execute torsocks"}, func() {africanaManual("torsocks")}},
        {[]string{"? 2", "info 2", "help 2", "torsocks", "info torsocks", "help torsocks"}, menus.HelpInfoTorsocks},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run networks", "use networks", "exec networks", "start networks", "launch networks", "exploit networks", "execute networks"}, func() {africanaManual("networks")}},
        {[]string{"? 3", "info 3", "help 3", "networks", "info networks", "help networks"}, menus.HelpInfoNetworks},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run exploits", "use exploits", "exec exploits", "start exploits", "launch exploits", "exploit exploits", "execute exploits"}, func() {africanaManual("exploits")}},
        {[]string{"? 4", "info 4", "help 4", "exploits", "info exploits", "help exploits"}, menus.HelpInfoExploits},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wireless", "use wireless", "exec wireless", "start wireless", "launch wireless", "exploit wireless", "execute wireless"}, func() {africanaManual("wireless")}},
        {[]string{"? 5", "info 5", "help 5", "wireless", "info wireless", "help wireless"}, menus.HelpInfoWireless},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run crackers", "use crackers", "exec crackers", "start crackers", "launch crackers", "exploit crackers", "execute crackers"}, func() {africanaManual("crackers")}},
        {[]string{"? 6", "info 6", "help 6", "crackers", "info crackers", "help crackers"}, menus.HelpInfoCrackers},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run phishers", "use phishers", "exec phishers", "start phishers", "launch phishers", "exploit phishers", "execute phishers"}, func() {africanaManual("phishers")}},
        {[]string{"? 7", "info 7", "help 7", "phishers", "info phishers", "help phishers"}, menus.HelpInfoPhishers},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run websites", "use websites", "exec websites", "start websites", "launch websites", "exploit websites", "execute websites"}, func() {africanaManual("websites")}},
        {[]string{"? 8", "info 8", "help 8", "websites", "info websites", "help websites"}, menus.HelpInfoWebsites},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run credits", "use credits", "exec credits", "start credits", "launch credits", "exploit credits", "execute credits"}, func() {africanaManual("credits")}},
        {[]string{"? 9", "info 9", "help 9", "credits", "info credits", "help credits"}, menus.HelpInfoCredits},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, func() {africanaManual("verses")}},
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
        "functions": &Function,
    }

    validKeys := make([]string, 0, len(setValues))
    for k := range setValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
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
        "setups": &Setups,
        "verses": &Verses,
        "funcs": &Function,
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
        "functions": &Function,
    }

    validKeys := make([]string, 0, len(unsetValues))
    for k := range unsetValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key]
        fmt.Printf("%s => %s\n", strings.ToUpper(key), "Null")
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
    if Function == "" {
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    africanaManual("")
}

func africanaManual(Function string, args ...interface{}) {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            //
        }
    }

    commands := map[string]func(){
        "setups":   setups.SetupsLauncher,
        "torsocks": securities.Torsocks,
        "networks": networks.NetworksPentest,
        "exploits": exploits.ExploitsPentest,
        "wireless": wireless.WirelessPentest,
        "crackers": crackers.CrackersPentest,
        "phishers": phishers.PhishingPentest,
        "websites": webattackers.WebsitesPentest,
        "credits":  credits.Creditors,
        "verses":   scriptures.ScriptureNarators,

        "1":  setups.SetupsLauncher,
        "2":  securities.Torsocks,
        "3":  networks.NetworksPentest,
        "4":  exploits.ExploitsPentest,
        "5":  wireless.WirelessPentest,
        "6":  crackers.CrackersPentest,
        "7":  phishers.PhishingPentest,
        "8":  webattackers.WebsitesPentest,
        "9":  credits.Creditors,
        "10": scriptures.ScriptureNarators,
    }

    textCommands := []string{
        "setups", "torsocks", "networks", "exploits", "wireless", "crackers", "phishers", "websites", "credits", "verses",
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        if num, err := strconv.Atoi(Function); err == nil {
            fmt.Printf("\n%s[!] %sNumber %d is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, num)
            menus.ListMainFunctions()
            return
        }

        lowerInput := strings.ToLower(Function)
        for _, cmd := range textCommands {
            lowerCmd := strings.ToLower(cmd)
            if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
                fmt.Printf("\n%s[!] %sModule '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
                return
            }
        }

        fmt.Printf("\n%s[!] %sModule %s is invalid. Available commands are:\n", bcolors.Yellow, bcolors.Endc, Function)
        menus.ListMainFunctions()
    }
}

func Genesis() {
    if len(os.Args) < 2 {
        launchDefaultMode()
        return
    }

    commandMap := createCommandMap()

    if action, exists := commandMap[os.Args[1]]; exists {
        action()
    } else {
        menus.HelpInfoMenuMain()
    }
}

func launchDefaultMode() {
    scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsTinny(); banners.GraphicsIntro(); africanaManualMode()
}

func launchAutoModeWithMenu(menuFunc func(), actionFunc func()) {
    scriptures.Verse(); utils.InitiLize(); setups.CheckTools(); banners.GraphicsLarge()
    if menuFunc != nil {
        menuFunc()
    }
    if actionFunc != nil {
        actionFunc()
    }
    africanaAutoMode()
}

func createCommandMap() map[string]func() {
    return map[string]func(){
        "-v":         banners.Version,
        "--version":  banners.Version,

        "-u":         setups.UpdateAfricana,
        "--update":   setups.UpdateAfricana,

        "-0":         func() {launchAutoModeWithMenu(nil, nil)},
        "-a":         func() {launchAutoModeWithMenu(nil, nil)},
        "--auto":     func() {launchAutoModeWithMenu(nil, nil)},

        "-1":         func() {launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher)},
        "-i":         func() {launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher)},
        "--install":  func() {launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher)},

        "-2":         func() {launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks)},
        "-t":         func() {launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks)},
        "--torsocks": func() {launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks)},

        "-3":         func() {launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest)},
        "-n":         func() {launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest)},
        "--networks": func() {launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest)},

        "-4":         func() {launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest)},
        "-e":         func() {launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest)},
        "--exploits": func() {launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest)},

        "-5":         func() {launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest)},
        "-w":         func() {launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest)},
        "--wireless": func() {launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest)},

        "-6":         func() {launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest)},
        "-c":         func() {launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest)},
        "--crackers": func() {launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest)},

        "-7":         func() {launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest)},
        "-p":         func() {launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest)},
        "--phishers": func() {launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest)},

        "-8":         func() {launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest)},
        "-x":         func() {launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest)},
        "--websites": func() {launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest)},

        "-9":         func() {credits.Creditors(); africanaAutoMode()},
        "-k":         func() {credits.Creditors(); africanaAutoMode()},
        "--credits":  func() {credits.Creditors(); africanaAutoMode()},

        "-10":        func() {scriptures.ScriptureNarators(); africanaAutoMode()},
        "-s":         func() {scriptures.ScriptureNarators(); africanaAutoMode()},
        "--verses":   func() {scriptures.ScriptureNarators(); africanaAutoMode()},

        "-g":         func() {banners.RandomBanners(); utils.BrowseTutarilas()},
        "--guide":    func() {banners.RandomBanners(); utils.BrowseTutarilas()},

        "-q":         func() {scriptures.Verse(); utils.InitiLize(); banners.Version(); africanaManualMode()},
        "--quite":    func() {scriptures.Verse(); utils.InitiLize(); banners.Version(); africanaManualMode()},

        "-00":       menus.HelpInfoMenuMain,
        "-h":        menus.HelpInfoMenuMain,
        "?":         menus.HelpInfoMenuMain,
        "--help":    menus.HelpInfoMenuMain,
    }
}
