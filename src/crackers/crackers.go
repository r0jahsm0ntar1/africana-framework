package crackers

import(
    "os"
    "fmt"
    "utils"
    "bufio"
    "menus"
    "banners"
    "strings"
    "bcolors"
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

func CrackersPentest() {
    for {
        fmt.Printf("%s%safr3%s crackers(%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Underl, bcolors.Bold, bcolors.Endc, bcolors.Red, Function, bcolors.Endc, bcolors.Green, bcolors.Endc)
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
    "info":             menus.HelpInfoCrackers,

    "m":                menus.MenuSix,
    "menu":             menus.MenuSix,

    "option":           menus.CrackersOptions,
    "options":          menus.CrackersOptions,
    "show option":      menus.CrackersOptions,
    "show options":     menus.CrackersOptions,

    "show all":         menus.ListCrackersFunctions,
    "list all":         menus.ListCrackersFunctions,

    "func":             menus.ListCrackersFunctions,
    "funcs":            menus.ListCrackersFunctions,
    "functions":        menus.ListCrackersFunctions,
    "show func":        menus.ListCrackersFunctions,
    "list funcs":       menus.ListCrackersFunctions,
    "show funcs":       menus.ListCrackersFunctions,
    "show function":    menus.ListCrackersFunctions,
    "list function":    menus.ListCrackersFunctions,
    "list functions":   menus.ListCrackersFunctions,
    "show functions":   menus.ListCrackersFunctions,

    "module":           menus.ListCrackersFunctions,
    "modules":          menus.ListCrackersFunctions,
    "list module":      menus.ListCrackersFunctions,
    "show module":      menus.ListCrackersFunctions,
    "list modules":     menus.ListCrackersFunctions,
    "show modules":     menus.ListCrackersFunctions,

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
        fmt.Printf("\n%s[!] %sMissing required parameter Function. Use %s'help' %sfor details.\n", bcolors.Red, bcolors.Endc, bcolors.Green, bcolors.Endc)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'help' %sfor details.\n", bcolors.Red, bcolors.Endc, bcolors.Green, bcolors.Endc)
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
    fmt.Print(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Print(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Print(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Print(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Print(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Print(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Print(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Print(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SSHcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPcracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "FTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMBCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "RDPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "LDAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "SNMTPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "TELNETCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "use " + bcolors.Yellow + "1. " + bcolors.Red + "Single  or " + bcolors.Yellow + "2. " + bcolors.Red + "Wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "admin" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + " Names to Crack " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HTTPSCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
        fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
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
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "All/SCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "TARGET " + bcolors.Magenta + "To attack" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    Rhost := scanner.Text()
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/cyberbrute; bash cyberbrute.sh %s`, Rhost)
    fmt.Println()
}

func HashBuster() {
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "HASHCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Full path to your hash " + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Hashes)
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/hash-buster; python3 cracker.py -t 10 %s`, Hashes)
    fmt.Println()
}

//Offline Crackers
func AirCrackng() {
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Full path to your .pcap " + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`aircrack-ng %s -w %s`, Pcap, WordList)
    fmt.Println()
}

func JohnCrackng() {
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Full path to your .pcap " + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.Blue + "\n╭─(" + bcolors.Endc + "africana:" + bcolors.Cyan + "framework:" + bcolors.Grey + bcolors.Italic + "PCAPCracker " + bcolors.Endc + bcolors.Italic + "set " + bcolors.Red + "Path to Pass wordlist " + bcolors.Magenta + "default " + bcolors.Endc + "= " + bcolors.Yellow + bcolors.Italic + "Rockyou.txt" + bcolors.Endc + bcolors.Blue + ")" + bcolors.Endc)
    fmt.Printf(bcolors.Blue + "\n╰─🐼" + bcolors.Green + "> " + bcolors.Endc)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`john %s --wordlist=%s`, Pcap, WordList)
    fmt.Println()
}
