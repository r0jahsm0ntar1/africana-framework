package crackers

import(
    "os"
    "fmt"
    "utils"
    "credits"
    "bufio"
    "menus"
    "banners"
    "strings"
    "bcolors"
    "scriptures"
    "subprocess"
)

var(
    scanner = bufio.NewScanner(os.Stdin)
    Pcap, Input, Rhost, Proxy, Function, Script, Output, Hashes, WordPass string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{

    "proxies": "",
    "rhost": "",
    "rhosts": "",
    "function": "",
    "output": OutPutDir,
    "password": WordPass,
    "wordlist": WordList,
}

type stringMatcher struct {
    names  []string
    action func()
}

func CrackersPentest() {
    for {
        fmt.Printf("%s%safr3%s crackers(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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
        {[]string{"info"}, menus.HelpInfoCrackers},
        {[]string{"m", "menu"}, menus.MenuSix},
        {[]string{"option", "options", "show option", "show options"}, menus.CrackersOptions},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListCrackersFunctions},

        // Commands executions
        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run ssh", "use ssh", "exec ssh", "start ssh", "launch ssh", "exploit ssh", "execute ssh"}, func() { HydraSsh() }},
        {[]string{"? 1", "info 1", "help 1", "ssh", "info ssh", "help ssh"}, menus.CrackersOptions},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run ftp", "use ftp", "exec ftp", "start ftp", "launch ftp", "exploit ftp", "execute ftp"}, func() { HydraFtp() }},
        {[]string{"? 2", "info 2", "help 2", "ftp", "info ftp", "help ftp"}, menus.CrackersOptions},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run networks", "use networks", "exec networks", "start networks", "launch networks", "exploit networks", "execute networks"}, func() { HydraSmb() }},
        {[]string{"? 3", "info 3", "help 3", "networks", "info networks", "help networks"}, menus.CrackersOptions},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run exploits", "use exploits", "exec exploits", "start exploits", "launch exploits", "exploit exploits", "execute exploits"}, func() { menus.MenuZero() }},
        {[]string{"? 4", "info 4", "help 4", "exploits", "info exploits", "help exploits"}, menus.CrackersOptions},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run wireless", "use wireless", "exec wireless", "start wireless", "launch wireless", "exploit wireless", "execute wireless"}, func() { menus.MenuZero() }},
        {[]string{"? 5", "info 5", "help 5", "wireless", "info wireless", "help wireless"}, menus.CrackersOptions},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run crackers", "use crackers", "exec crackers", "start crackers", "launch crackers", "exploit crackers", "execute crackers"}, func() { menus.MenuZero() }},
        {[]string{"? 6", "info 6", "help 6", "crackers", "info crackers", "help crackers"}, menus.CrackersOptions},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run phishers", "use phishers", "exec phishers", "start phishers", "launch phishers", "exploit phishers", "execute phishers"}, func() { menus.MenuZero() }},
        {[]string{"? 7", "info 7", "help 7", "phishers", "info phishers", "help phishers"}, menus.CrackersOptions},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run websites", "use websites", "exec websites", "start websites", "launch websites", "exploit websites", "execute websites"}, func() { menus.MenuZero() }},
        {[]string{"? 8", "info 8", "help 8", "websites", "info websites", "help websites"}, menus.CrackersOptions},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run credits", "use credits", "exec credits", "start credits", "launch credits", "exploit credits", "execute credits"}, func() { credits.Creditors() }},
        {[]string{"? 9", "info 9", "help 9", "credits", "info credits", "help credits"}, menus.CrackersOptions},

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

        "rhost": &Rhost,
        "rhosts": &Rhost,
        "proxies": &Proxy,
        "func": &Function,
        "module": &Function,
        "output": &OutPutDir,
        "function": &Function,
        "wordlist": &WordList,
        "password": &WordPass,

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
        "proxies": &Proxy,
        "func": &Function,
        "module": &Function,
        "output": &OutPutDir,
        "function": &Function,
        "wordlist": &WordList,
        "password": &WordPass,
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
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    CrackersPenFunctions(Function, Rhost)
}

// Helper functions
func autoExecuteFunc(distro string, function string) {
    //Distro = distro
    //Function = function
    executeFunction()
}

func CrackersPenFunctions(Function string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nFunction => %s\n", Rhost, Function)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        if err := utils.SetProxy(Proxy); err != nil {
            // Error already printed by SetProxy
        }
    }

    commands := map[string]func() {
        "ssh":        func() {HydraSsh()},
        "ftp":        func() {HydraFtp()},
        "smb":        func() {HydraSmb()},
    }

    if action, exists := commands[Function]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid Function %s. Use %s'help' %sfor available Functions.\n", bcolors.Yellow, bcolors.Endc, Function, bcolors.Green, bcolors.Endc)
    }
}


//Online crackers
func HydraSsh() {
    fmt.Print(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Print(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Print(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Print(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Print(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Print(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Print(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Print(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SSHcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSsh_outfile.txt -u ssh://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SSHcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSsh_outfile.txt -u ssh://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraFtp() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "FTPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraFtp_outfile.txt -u ftp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "FTPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraFtp_outfile.txt -u ftp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSmb() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SMBcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSmb_outfile.txt -u smb://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SMBCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSmb_outfile.txt -u smb://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraRdp() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "RDPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraRdp_outfile.txt -u rdp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "RDPCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraRdp_outfile.txt -u rdp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraLdap() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "LDAPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraLdap_outfile.txt -u ldap://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "LDAPCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraLdap_outfile.txt -u ldap://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSmtp() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SMTPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSmtp_outfile.txt -u smtp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SMTPCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSmtp_outfile.txt -u smtp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSnmtp() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SNMTPCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSnmtp_outfile.txt -u snmtp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "SNMTPCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSnmtp_outfile.txt -u snmtp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraTelnet() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "FTPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraTelnet_outfile.txt -u telnet://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "TELNETCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraTelnet_outfile.txt -u telnet://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraHttps() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.BrightRed + "Single  or " + bcolors.Yellow + "2. " + bcolors.BrightRed + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "FTPcracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraHttps_outfile.txt -u https://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Endc + "\n[" + bcolors.Cyan + "HTTPSCracker has began" + bcolors.Endc + "] [" + bcolors.Green + "Target" + bcolors.Endc + ": " + bcolors.Yellow + "%s" + bcolors.Endc + "]\n\n" + bcolors.Endc, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraHttps_outfile.txt -u http://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func CyberBrute() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "All/SCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/cyberbrute; bash cyberbrute.sh %s`, Rhost)
    fmt.Println()
}

func HashBuster() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HASHCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Full path to your hash " + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Hashes)
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/hash-buster; python3 cracker.py -t 10 %s`, Hashes)
    fmt.Println()
}

//Offline Crackers
func AirCrackng() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Full path to your .pcap " + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`aircrack-ng %s -w %s`, Pcap, WordList)
    fmt.Println()
}

func JohnCrackng() {
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Full path to your .pcap " + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.BrightBlue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.BrightRed + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.BrightBlue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.BrightBlue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`john %s --wordlist=%s`, Pcap, WordList)
    fmt.Println()
}
