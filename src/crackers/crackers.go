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
    Pcap, Input, Rhost, Proxy, Module, Script, Output, Hashes, WordPass string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "proxy": "",
    "rhost": "",
    "rhosts": "",
    "module": "",
    "password": WordPass,
    "output": OutPutDir,
    "wordlist": WordList,
}

func PasswordPentest() {
    for {
        fmt.Printf("%s%safr3%s crackers(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "passwords_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
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
            executeModule()
        default:
            utils.SystemShell(Input)
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
        "info":             menus.HelpInfoMain,

        "m":                menus.MenuEight,
        "menu":             menus.MenuEight,

        "option":           menus.HelpInfOptions,
        "options":          menus.HelpInfOptions,
        "show option":      menus.HelpInfOptions,
        "show options":     menus.HelpInfOptions,

        "modules":          menus.ListMainModules,
        "show all":         menus.ListMainModules,
        "list all":         menus.ListMainModules,
        "list modules":     menus.ListMainModules,
        "show modules":     menus.ListMainModules,
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
        "proxy": &Proxy,
        "rhost": &Rhost,
        "rhosts": &Rhost,
        "module": &Module,
        "output": &OutPutDir,
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
        "proxy": &Proxy,
        "rhost": &Rhost,
        "rhosts": &Rhost,
        "module": &Module,
        "output": &OutPutDir,
        "wordlist": &WordList,
        "password": &WordPass,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
    } else {
        menus.HelpInfoSet()
    }
}

func executeModule() {
    if Module == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter MODULE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters DISTRO. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    CrackersPenModules(Module, Rhost)
}

func CrackersPenModules(module string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nMODULE => %s\n", Rhost, module)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        utils.SetProxy(Proxy)
    }

    commands := map[string]func() {
        "ssh":        func() {HydraSsh()},
        "ftp":        func() {HydraFtp()},
        "smb":        func() {HydraSmb()},
    }

    if action, exists := commands[module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid module %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, module, bcolors.DARKGREEN, bcolors.ENDC)
    }
}






//Online crackers
func HydraSsh() {
    fmt.Print(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Print(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Print(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Print(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Print(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Print(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Print(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Print(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SSHcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSsh_outfile.txt -u ssh://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SSHcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSsh_outfile.txt -u ssh://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraFtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraFtp_outfile.txt -u ftp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraFtp_outfile.txt -u ftp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSmb() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMBcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSmb_outfile.txt -u smb://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMBCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSmb_outfile.txt -u smb://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraRdp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "RDPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraRdp_outfile.txt -u rdp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "RDPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraRdp_outfile.txt -u rdp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraLdap() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "LDAPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraLdap_outfile.txt -u ldap://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "LDAPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraLdap_outfile.txt -u ldap://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSmtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSmtp_outfile.txt -u smtp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSmtp_outfile.txt -u smtp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraSnmtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SNMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraSnmtp_outfile.txt -u snmtp://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SNMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraSnmtp_outfile.txt -u snmtp://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraTelnet() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraTelnet_outfile.txt -u telnet://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "TELNETCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraTelnet_outfile.txt -u telnet://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func HydraHttps() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single  or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Input := scanner.Text()
    switch strings.ToLower(Input) {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Name := scanner.Text()
        if Name == "" {
            Name = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.afr3/output/HydraHttps_outfile.txt -u https://%s`, Name, WordPass, Rhost)
        fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + " Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        Names := scanner.Text()
        if Names == "" {
            Names = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        WordPass := scanner.Text()
        if WordPass == "" {
            WordPass = "/usr/share/WordList/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "HTTPSCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, Rhost)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.afr3/output/HydraHttps_outfile.txt -u http://%s`, Names, WordPass, Rhost)
        fmt.Println()
        return
    }
}

func CyberBrute() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "All/SCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    Rhost := scanner.Text()
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/cyberbrute; bash cyberbrute.sh %s`, Rhost)
    fmt.Println()
}

func HashBuster() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HASHCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your hash " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    fmt.Scan(&Hashes)
    subprocess.Popen(`cd /root/.afr3/africana-base/crackers/hash-buster; python3 cracker.py -t 10 %s`, Hashes)
    fmt.Println()
}

//Offline Crackers
func AirCrackng() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your .pcap " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`aircrack-ng %s -w %s`, Pcap, WordList)
    fmt.Println()
}

func JohnCrackng() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your .pcap " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    fmt.Scan(&Pcap)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "> " + bcolors.ENDC)
    scanner.Scan()
    WordList := scanner.Text()
    if WordList == "" {
        WordList = "/usr/share/WordList/rockyou.txt"
    }
    subprocess.Popen(`john %s --wordlist=%s`, Pcap, WordList)
    fmt.Println()
}
