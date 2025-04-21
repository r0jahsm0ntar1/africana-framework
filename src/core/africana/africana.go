package africana

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "sort"
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
    proxyURL, filePath, Input, Rhost, Proxy, Xhost, Yhost, Setups, Torsocks, Networks, Exploits, Wireless, Crackers, Phishers, Websites, Credits, Verses, Function string
)

var defaultValues = map[string]string{
    "module": "",
    "proxies": "",
}

type stringMatcher struct {
    names  []string
    action func()
}

func Run() {
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
    africanaFrameworAuto()
}

func africanaFrameworAuto() {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    // Command mapping with both text and numeric support
    commands := map[string]func(){
        // Text commands
        "setups":   func() { menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero()} ,
        "torsocks": func() { menus.MenuTwo(); securities.Torsocks(); menus.MenuZero() },
        "networks": func() { menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero() },
        "exploits": func() { menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero() },
        "wireless": func() { menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero() },
        "crackers": func() { menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero() },
        "phishers": func() { menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero() },
        "websites": func() { menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero() },
        "credits":  func() { credits.Creditors() },
        "verses":   func() { scriptures.ScriptureNarators() },
        
        // Numeric shortcuts (1-9)
        "1": func() { menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero() },
        "2": func() { menus.MenuTwo(); securities.Torsocks(); menus.MenuZero() },
        "3": func() { menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero() },
        "4": func() { menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero() },
        "5": func() { menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero() },
        "6": func() { menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero() },
        "7": func() { menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero() },
        "8": func() { menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero() },
        "9": func() { credits.Creditors() },
       "10": func() { scriptures.ScriptureNarators() }, // Bonus number
    }

    // Create reverse mapping for error messages
    commandNames := map[string]string{
        "1":   "setups",
        "2":   "torsocks",
        "3":   "networks",
        "4":   "exploits",
        "5":   "wireless",
        "6":   "crackers",
        "7":   "phishers",
        "8":   "websites",
        "9":   "credits",
        "10":  "verses",
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        // Try to find if input was a number not in our map
        if _, err := strconv.Atoi(Function); err == nil {
            fmt.Printf("\n%s[!] %sNumber %s is invalid. Valid numbers are:\n", bcolors.Yellow, bcolors.Endc, Function)
            // Print available number mappings
            nums := make([]int, 0, len(commandNames))
            for k := range commandNames {
                if n, err := strconv.Atoi(k); err == nil {
                    nums = append(nums, n)
                }
            }
            sort.Ints(nums)

            for _, n := range nums {
                key := strconv.Itoa(n)
                fmt.Printf(" %s%s. %s%-3s %s> %s\n", bcolors.BrightBlue, key, bcolors.Yellow, commandNames[key], bcolors.Endc, getCommandDescription(commandNames[key]))
            }
            fmt.Println()
        } else {
            fmt.Printf("\n%s[!] %sFunction %s is invalid. Use %s'help' %sfor available modules.\n", bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, bcolors.Endc)
        }
    }
}

func executeCommandAuto(cmd string) bool {
    commandGroups := []stringMatcher{
        // Info/Help commands
        {[]string{"? info", "h info", "help info"}, menus.HelpInfo},
        {[]string{"v", "version"}, banners.Version},
        {[]string{"s", "sleep"}, utils.Sleep},
        {[]string{"c", "clear", "clear screen", "screen clear"}, utils.ClearScreen},
        {[]string{"o", "junks", "outputs", "clear junks", "clear outputs"}, utils.ListJunks},
        {[]string{"logs", "history", "clear logs", "clear history"}, subprocess.LogHistory},

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
        {[]string{"info"}, menus.HelpInfoMain},
        {[]string{"m", "menu"}, menus.MenuZero},
        {[]string{"option", "options", "show option", "show options"}, menus.MainOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListMainFunctions},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() { menus.MenuOne(); setups.SetupsLauncher(); menus.MenuZero() }},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, menus.HelpInfoSetups},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run torsocks", "use torsocks", "exec torsocks", "start torsocks", "launch torsocks", "exploit torsocks", "execute torsocks"}, func() { menus.MenuTwo(); securities.Torsocks(); menus.MenuZero() }},
        {[]string{"? 2", "info 2", "help 2", "torsocks", "info torsocks", "help torsocks"}, menus.HelpInfoTorsocks},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run networks", "use networks", "exec networks", "start networks", "launch networks", "exploit networks", "execute networks"}, func() { menus.MenuThree(); networks.NetworksPentest(); menus.MenuZero() }},
        {[]string{"? 3", "info 3", "help 3", "networks", "info networks", "help networks"}, menus.HelpInfoNetworks},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run exploits", "use exploits", "exec exploits", "start exploits", "launch exploits", "exploit exploits", "execute exploits"}, func() { menus.MenuFour(); exploits.ExploitsPentest(); menus.MenuZero() }},
        {[]string{"? 4", "info 4", "help 4", "exploits", "info exploits", "help exploits"}, menus.HelpInfoExploits},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wireless", "use wireless", "exec wireless", "start wireless", "launch wireless", "exploit wireless", "execute wireless"}, func() { menus.MenuFive(); wireless.WirelessPentest(); menus.MenuZero() }},
        {[]string{"? 5", "info 5", "help 5", "wireless", "info wireless", "help wireless"}, menus.HelpInfoWireless},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run crackers", "use crackers", "exec crackers", "start crackers", "launch crackers", "exploit crackers", "execute crackers"}, func() { menus.MenuSix(); crackers.CrackersPentest(); menus.MenuZero() }},
        {[]string{"? 6", "info 6", "help 6", "crackers", "info crackers", "help crackers"}, menus.HelpInfoCrackers},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run phishers", "use phishers", "exec phishers", "start phishers", "launch phishers", "exploit phishers", "execute phishers"}, func() { menus.MenuSeven(); phishers.PhishingPentest(); menus.MenuZero() }},
        {[]string{"? 7", "info 7", "help 7", "phishers", "info phishers", "help phishers"}, menus.HelpInfoPhishers},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run websites", "use websites", "exec websites", "start websites", "launch websites", "exploit websites", "execute websites"}, func() { menus.MenuEight(); webattackers.WebsitesPentest(); menus.MenuZero() }},
        {[]string{"? 8", "info 8", "help 8", "websites", "info websites", "help websites"}, menus.HelpInfoWebsites},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run credits", "use credits", "exec credits", "start credits", "launch credits", "exploit credits", "execute credits"}, func() { credits.Creditors() }},
        {[]string{"? 9", "info 9", "help 9", "credits", "info credits", "help credits"}, menus.HelpInfoCredits},

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

//Manual mode
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
        // Info/Help commands
        {[]string{"? info", "h info", "help info"}, menus.HelpInfo},
        {[]string{"v", "version"}, banners.Version},
        {[]string{"s", "sleep"}, utils.Sleep},
        {[]string{"c", "clear", "clear screen", "screen clear"}, utils.ClearScreen},
        {[]string{"o", "junks", "outputs", "clear junks", "clear outputs"}, utils.ListJunks},
        {[]string{"logs", "history", "clear logs", "clear history"}, subprocess.LogHistory},

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
        {[]string{"info"}, menus.HelpInfoMain},
        {[]string{"m", "menu"}, menus.MenuZero},
        {[]string{"option", "options", "show option", "show options"}, menus.MainOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListMainFunctions},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run setups", "use setups", "exec setups", "start setups", "launch setups", "exploit setups", "execute setups"}, func() { setups.SetupsLauncher() }},
        {[]string{"? 1", "info 1", "help 1", "setups", "info setups", "help setups"}, menus.HelpInfoSetups},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run torsocks", "use torsocks", "exec torsocks", "start torsocks", "launch torsocks", "exploit torsocks", "execute torsocks"}, func() { securities.Torsocks() }},
        {[]string{"? 2", "info 2", "help 2", "torsocks", "info torsocks", "help torsocks"}, menus.HelpInfoTorsocks},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run networks", "use networks", "exec networks", "start networks", "launch networks", "exploit networks", "execute networks"}, func() { networks.NetworksPentest() }},
        {[]string{"? 3", "info 3", "help 3", "networks", "info networks", "help networks"}, menus.HelpInfoNetworks},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run exploits", "use exploits", "exec exploits", "start exploits", "launch exploits", "exploit exploits", "execute exploits"}, func() { exploits.ExploitsPentest() }},
        {[]string{"? 4", "info 4", "help 4", "exploits", "info exploits", "help exploits"}, menus.HelpInfoExploits},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wireless", "use wireless", "exec wireless", "start wireless", "launch wireless", "exploit wireless", "execute wireless"}, func() { wireless.WirelessPentest() }},
        {[]string{"? 5", "info 5", "help 5", "wireless", "info wireless", "help wireless"}, menus.HelpInfoWireless},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run crackers", "use crackers", "exec crackers", "start crackers", "launch crackers", "exploit crackers", "execute crackers"}, func() { crackers.CrackersPentest() }},
        {[]string{"? 6", "info 6", "help 6", "crackers", "info crackers", "help crackers"}, menus.HelpInfoCrackers},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run phishers", "use phishers", "exec phishers", "start phishers", "launch phishers", "exploit phishers", "execute phishers"}, func() { phishers.PhishingPentest() }},
        {[]string{"? 7", "info 7", "help 7", "phishers", "info phishers", "help phishers"}, menus.HelpInfoPhishers},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run websites", "use websites", "exec websites", "start websites", "launch websites", "exploit websites", "execute websites"}, func() { webattackers.WebsitesPentest() }},
        {[]string{"? 8", "info 8", "help 8", "websites", "info websites", "help websites"}, menus.HelpInfoWebsites},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run credits", "use credits", "exec credits", "start credits", "launch credits", "exploit credits", "execute credits"}, func() { credits.Creditors() }},
        {[]string{"? 9", "info 9", "help 9", "credits", "info credits", "help credits"}, menus.HelpInfoCredits},

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
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    africanaManual()
}

func africanaManual() {
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    // Command mapping with both text and numeric support
    commands := map[string]func(){
        // Text commands (raw names without spaces)
        "setups":   func() { setups.SetupsLauncher() },
        "torsocks": func() { securities.Torsocks() },
        "networks": func() { networks.NetworksPentest() },
        "exploits": func() { exploits.ExploitsPentest() },
        "wireless": func() { wireless.WirelessPentest() },
        "crackers": func() { crackers.CrackersPentest() },
        "phishers": func() { phishers.PhishingPentest() },
        "websites": func() { webattackers.WebsitesPentest() },
        "credits":  func() { credits.Creditors() },
        "verses":   func() { scriptures.ScriptureNarators() },
        
        // Numeric shortcuts (1-9)
        "1": func() { setups.SetupsLauncher() },
        "2": func() { securities.Torsocks() },
        "3": func() { networks.NetworksPentest() },
        "4": func() { exploits.ExploitsPentest() },
        "5": func() { wireless.WirelessPentest() },
        "6": func() { crackers.CrackersPentest() },
        "7": func() { phishers.PhishingPentest() },
        "8": func() { webattackers.WebsitesPentest() },
        "9": func() { credits.Creditors() },
        "10": func() { scriptures.ScriptureNarators() },
    }

    // Display names with spaces for formatting (only used for printing)
    prettyNames := map[string]string{
        "setups":   "  setups",
        "torsocks": "torsocks",
        "networks": "networks",
        "exploits": "exploits",
        "wireless": "wireless",
        "crackers": "crackers",
        "phishers": "phishers",
        "websites": "websites",
        "credits":  " credits",
        "verses":   " verses",
    }

    // Create reverse mapping for error messages
    commandNames := map[string]string{
        "1":   "setups",
        "2":   "torsocks",
        "3":   "networks",
        "4":   "exploits",
        "5":   "wireless",
        "6":   "crackers",
        "7":   "phishers",
        "8":   "websites",
        "9":   "credits",
        "10":  "verses",
    }

    // Create text command list for typo checking (raw names)
    textCommands := []string{
        "setups", "torsocks", "networks", "exploits", "wireless",
        "crackers", "phishers", "websites", "credits", "verses",
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        // Check if input was a number not in our map
        if _, err := strconv.Atoi(Function); err == nil {
            fmt.Printf("\n%s[!] %sNumber %s is invalid. Valid numbers are:\n", 
                bcolors.Yellow, bcolors.Endc, Function)
            
            // Print available number mappings
            nums := make([]int, 0, len(commandNames))
            for k := range commandNames {
                if n, err := strconv.Atoi(k); err == nil {
                    nums = append(nums, n)
                }
            }
            sort.Ints(nums)
            
            for _, n := range nums {
                key := strconv.Itoa(n)
                cmd := commandNames[key]
                fmt.Printf("  %s%s. %s%-3s %s> %s\n",
                    bcolors.BrightBlue, key,
                    bcolors.Yellow, prettyNames[cmd],
                    bcolors.Endc, getCommandDescription(cmd))
            }
            fmt.Println()
        } else {
            // Check if input was a text command with typo
            foundSimilar := false
            lowerInput := strings.ToLower(Function)
            
            for _, cmd := range textCommands {
                lowerCmd := strings.ToLower(cmd)
                if strings.HasPrefix(lowerCmd, lowerInput) || 
                   strings.Contains(lowerCmd, lowerInput) ||
                   levenshtein(lowerInput, lowerCmd) <= 2 {
                    fmt.Printf("\n%s[!] %sCommand %s is invalid. Did you mean %s'%s'%s?\n", 
                        bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, cmd, bcolors.Endc)
                    foundSimilar = true
                    break
                }
            }

            if !foundSimilar {
                fmt.Printf("\n%s[!] %sCommand %s is invalid. Available commands are:\n", 
                    bcolors.Yellow, bcolors.Endc, Function)
                
                // Print available text commands with beautiful alignment
                for _, cmd := range textCommands {
                    fmt.Printf(" %s%-3s %s> %s\n",
                        bcolors.Yellow, prettyNames[cmd],
                        bcolors.Endc, getCommandDescription(cmd))
                }
                fmt.Println()
            }
        }
    }
}

// Helper function to get descriptions
func getCommandDescription(cmd string) string {
    descriptions := map[string]string{
        "setups":   "Install/Update framework components",
        "torsocks": "Configure Tor and proxy settings",
        "networks": "Network penetration testing tools",
        "exploits": "Exploit development framework",
        "wireless": "Wireless network attacks",
        "crackers": "Password cracking utilities", 
        "phishers": "Phishing campaign tools",
        "websites": "Web application testing",
        "credits":  "Developer credits and info",
        "verses":   "Biblical scripture integration",
    }
    return descriptions[cmd]
}

// Levenshtein distance function for fuzzy matching
func levenshtein(a, b string) int {
    la := len(a)
    lb := len(b)
    d := make([][]int, la+1)
    for i := 0; i <= la; i++ {
        d[i] = make([]int, lb+1)
        d[i][0] = i
    }
    for j := 1; j <= lb; j++ {
        d[0][j] = j
    }
    for i := 1; i <= la; i++ {
        for j := 1; j <= lb; j++ {
            cost := 0
            if a[i-1] != b[j-1] {
                cost = 1
            }
            d[i][j] = min(
                d[i-1][j]+1,
                d[i][j-1]+1,
                d[i-1][j-1]+cost,
            )
        }
    }
    return d[la][lb]
}

func min(nums ...int) int {
    m := nums[0]
    for _, n := range nums {
        if n < m {
            m = n
        }
    }
    return m
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
    scriptures.Verse()
    utils.InitiLize()
    setups.CheckTools()
    banners.GraphicsLarge()
    banners.Version()
    africanaManualMode()
}

func launchAutoModeWithMenu(menuFunc func(), actionFunc func()) {
    scriptures.Verse()
    utils.InitiLize()
    setups.CheckTools()
    banners.GraphicsLarge()
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
        // Version commands
        "-v":        banners.Version,
        "--version": banners.Version,
        
        // Update commands
        "-u":        setups.UpdateAfricana,
        "--update":  setups.UpdateAfricana,
        
        // Auto mode commands
        "-0":        func() { launchAutoModeWithMenu(nil, nil) },
        "-a":        func() { launchAutoModeWithMenu(nil, nil) },
        "--auto":    func() { launchAutoModeWithMenu(nil, nil) },
        
        // Installation commands
        "-1":        func() { launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher) },
        "-i":        func() { launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher) },
        "--install": func() { launchAutoModeWithMenu(menus.MenuOne, setups.SetupsLauncher) },
        
        // Torsocks commands
        "-2":        func() { launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks) },
        "-t":        func() { launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks) },
        "--torsocks": func() { launchAutoModeWithMenu(menus.MenuTwo, securities.Torsocks) },
        
        // Networks commands
        "-3":        func() { launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest) },
        "-n":        func() { launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest) },
        "--networks": func() { launchAutoModeWithMenu(menus.MenuThree, networks.NetworksPentest) },
        
        // Exploits commands
        "-4":        func() { launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest) },
        "-e":        func() { launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest) },
        "--exploits": func() { launchAutoModeWithMenu(menus.MenuFour, exploits.ExploitsPentest) },
        
        // Wireless commands
        "-5":        func() { launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest) },
        "-w":        func() { launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest) },
        "--wireless": func() { launchAutoModeWithMenu(menus.MenuFive, wireless.WirelessPentest) },
        
        // Crackers commands
        "-6":        func() { launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest) },
        "-p":        func() { launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest) },
        "--crackers": func() { launchAutoModeWithMenu(menus.MenuSix, crackers.CrackersPentest) },
        
        // Phishers commands
        "-7":        func() { launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest) },
        "-f":        func() { launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest) },
        "--phishers": func() { launchAutoModeWithMenu(menus.MenuSeven, phishers.PhishingPentest) },
        
        // Websites commands
        "-8":        func() { launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest) },
        "-x":        func() { launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest) },
        "--websites": func() { launchAutoModeWithMenu(menus.MenuEight, webattackers.WebsitesPentest) },
        
        // Credits commands
        "-9":        func() { credits.Creditors(); africanaAutoMode() },
        "-c":        func() { credits.Creditors(); africanaAutoMode() },
        "--credits": func() { credits.Creditors(); africanaAutoMode() },
        
        // Verses commands
        "-10":       func() { scriptures.ScriptureNarators(); africanaAutoMode() },
        "-s":        func() { scriptures.ScriptureNarators(); africanaAutoMode() },
        "--verses":  func() { scriptures.ScriptureNarators(); africanaAutoMode() },
        
        // Guide commands
        "-g":        func() { banners.RandomBanners(); utils.BrowseTutarilas() },
        "--guide":   func() { banners.RandomBanners(); utils.BrowseTutarilas() },
        
        // Quiet mode commands
        "-q":        func() { 
            scriptures.Verse()
            utils.InitiLize()
            banners.GraphicsTinny()
            banners.Version()
            africanaManualMode()
        },
        "--quite":   func() { 
            scriptures.Verse()
            utils.InitiLize()
            banners.GraphicsTinny()
            banners.Version()
            africanaManualMode()
        },
        
        // Help commands
        "-00":       menus.HelpInfoMenuMain,
        "-h":        menus.HelpInfoMenuMain,
        "?":         menus.HelpInfoMenuMain,
        "--help":    menus.HelpInfoMenuMain,
    }
}
