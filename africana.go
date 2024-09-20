package main

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
    "net/url"
    "runtime"
    "bcolors"
    "butchers"
    "wireless"
    "crackers"
    "phishers"
    "io/ioutil"
    "internals"
    "subprocess"
    "agreements"
    "scriptures"
    "securities"
    "webattackers"
)

var(
    userInput   string
    userTarget  string
    userXtarget string
    userXdomain string
    proxyURL    string
    scanner = bufio.NewScanner(os.Stdin)
)

func africana() {
    switch runtime.GOOS {
    case "windows":
        userAgreements()
        return
    default:
        if os.Geteuid() != 0 {
            utils.ClearScreen()
            banners.Banner()
            scriptures.Verse()
            fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🦊africana must be run as root. Try " + bcolors.RED + "sudo africana" + bcolors.ENDC)
            return
        }else{
        userAgreements()
        return
        }
    }
}

func userAgreements() {
    filePath := "/root/.africana/agreements/"
    content := "📜 User accepted the terms."
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        utils.ClearScreen()
        agreements.Covenant()
        for {
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🦊" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userInput := scanner.Text()
            switch strings.ToLower(userInput) {
                case "y", "yes":
                    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
                        fmt.Println("Error creating file:", err)
                        return
                    }
                    err = ioutil.WriteFile(filePath + "covenant.txt", []byte(content), os.ModePerm)
                    if err != nil {
                        fmt.Println("Error writing to file:", err)
                        return
                    }
                    utils.InitiLize()
                    genesis()
                    return
                case "n", "no":
                    os.Exit(0)
                default:
                    fmt.Println(bcolors.GREEN + "(" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
                }
            }
    } else {
        utils.InitiLize()
        genesis()
        return
    }
}

//1. Install or update africana-framework..(Start here )🩺
func systemSetups() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuOne()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Installer" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🔧" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1", "use kali":
            utils.ClearScreen()
            banners.Banner()
            setups.KaliSetups()
            africanaFramework()
            return
        case "2", "use ubuntu":
            utils.ClearScreen()
            banners.Banner()
            setups.UbuntuSetups()
            africanaFramework()
            return
        case "3", "use arch":
            utils.ClearScreen()
            banners.Banner()
            setups.ArchSetups()
            africanaFramework()
            return
        case "4", "use windows":
            utils.ClearScreen()
            banners.Banner()
            setups.WindowsSetups()
            utils.UpsentTools()
            africanaFramework()
            return
        case "5", "use android":
            utils.ClearScreen()
            banners.Banner()
            setups.Androidinstall()
        case "6", "use termlogs":
            utils.TermLogs()
        case "7", "use viewlogs":
            utils.BrowserLogs()
        case "8", "use clearlogs":
            utils.ClearLogs()
        case "9", "use uninstall":
            utils.ClearScreen()
            banners.Banner()
            setups.Uninstall()
        case "show", "help show":
            menus.HelpInfoShow()
        case "show modules", "show all":
            menus.ListSetupsModules()
        case "info", "help info":
            menus.HelpInfoSetups()
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuOne()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//2. System Security Configuration.........(Setup tor &)🎭
func anonsurfSetups() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuTwo()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Anonsurf" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1", "anonsurf install":
            utils.ClearScreen()
            securities.AnonsurfSetups()
        case "2", "anonsurf start":
            utils.ClearScreen()
            securities.AnonsurfStart()
        case "3", "anonsurf restart":
            utils.ClearScreen()
            securities.AnonsurfExitnode()
        case "4", "anonsurf status":
            utils.ClearScreen()
            securities.AnonsurfStatus()
        case "5", "anonsurf ip":
            utils.ClearScreen()
            securities.AnonsurfIpaddr()
        case "6", "anonsurf iptables":
            utils.ClearScreen()
            securities.AnonsurfRIptabls()
        case "7", "anonsurf reload":
            utils.ClearScreen()
            securities.AnonsurfReload()
        case "8", "anonsurf logs":
            utils.ClearScreen()
            securities.AnonsurfChains()
        case "9", "anonsurf stop":
            utils.ClearScreen()
            securities.AnonsurfStop()
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuTwo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//3. Local Network Attack Vectors..........(Mitm, sniff)🐹
func internaltargetInput() {
    internals.InternalScan()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Networks " + bcolors.ENDC + bcolors.ITALIC + "Select your " + bcolors.RED + "Target!🎯" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐹" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userTarget = scanner.Text()
}

func internalAttackers() {
    banners.Banner()
    menus.MenuThree()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Networks" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐹" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1", "discover targets":
            internaltargetInput()
            utils.ClearScreen()
            banners.Banner()
            menus.MenuThree()
        case "2", "run portscan":
            internals.NmapPortScan(userTarget)
        case "3", "run vulnscan":
            internals.NmapVulnScan(userTarget)
        case "4", "run enumscan":
            internals.SmbVulnScan(userTarget)
            internals.Enum4linux(userTarget)
            internals.EnumNxc(userTarget)
            internals.SmbMapScan(userTarget)
            internals.RpcEnumScan(userTarget)
        case "5", "run exploit":
            internals.SmbVulnScan(userTarget)
            internals.SmbExploit(userTarget)
        case "6", "run sniffer":
            internals.PacketSniffer(userTarget)
            internalAttackers()
            return
        case "7", "run responder":
            internals.PacketsResponder()
        case "8", "run beefninja":
            internals.BeefInjector(userTarget)
        case "9":
            internals.RpcEnumScan(userTarget)
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuThree()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//4. Generate Undetectable Backdoors.......(C2 & shells)🐭
func malwareGenerators() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuFour()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Malwares" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourOne()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Androids" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    butchers.TearDroid()
                case "2":
                    butchers.AndroRat()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                     utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFour()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "2":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourTwo()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Iphones" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    utils.UpsentTools()
                case "2":
                    utils.UpsentTools()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourTwo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "3":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourThree()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Windows" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    butchers.BlackJack()
                case "2":
                    butchers.Shellz()
                case "3":
                    butchers.PowerJoker()
                case "4":
                    butchers.MeterPeter()
                case "5":
                     butchers.Havoc()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourThree()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "4":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourFour()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "MackOS" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    utils.UpsentTools()
                case "2":
                    utils.UpsentTools()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourFour()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "5": 
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourFive()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LinuxOS" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    utils.UpsentTools()
                case "2":
                    utils.UpsentTools()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourFive()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "6":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourSix()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Websites" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    utils.UpsentTools()
                case "2":
                    utils.UpsentTools()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourSix()
                case "00", "?", "h", "help":
                    menus.MenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "7":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourSeven()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Univasals" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    butchers.NoiseMakers()
                case "2":
                    butchers.CodeBreakers()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourSeven()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "8":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourEight()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Obfsications" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    butchers.Gh0x0st()
                case "2":
                    butchers.Chameleon()
                case "3":
                    utils.UpsentTools()
                case "4":
                    utils.UpsentTools()
                case "5":
                    utils.UpsentTools()
                case "6":
                    utils.UpsentTools()
                case "7":
                    utils.UpsentTools()
                case "8":
                    utils.UpsentTools()
                case "9":
                    utils.UpsentTools()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourEight()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "9":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourNine()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "ChosenOnes" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "1":
                    butchers.BlackJack()
                case "2":
                    butchers.Shellz()
                case "3":
                    butchers.PowerJoker()
                case "4":
                    butchers.MeterPeter()
                case "5":
                     butchers.Havoc()
                case "6":
                    butchers.TearDroid()
                case "7":
                    butchers.AndroRat()
                case "8":
                    butchers.Chameleon()
                case "9":
                    butchers.Gh0x0st()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourNine()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "show", "show modules":
            menus.ListMalwareModules()
        case "info", "help info":
            menus.HelpInfoMalwares()
        case "info blackjack":
            menus.HelpInfoBlackJack()
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuFour()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//5. WiFi Attack Vectors...................(Wifite, air)🍍
func wirelessAttackers() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuFive()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Wireless" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1":
            wireless.WifiteAuto()
        case "2":
            wireless.BettercapAuto()
        case "3":
            wireless.WifiPumpkin3Auto()
        case "4":
            utils.UpsentTools()
        case "5":
            utils.UpsentTools()
        case "6":
            wireless.FluxionMan()
        case "7":
            wireless.AirGeddon()
        case "8":
            wireless.WifiPumpkin3()
        case "9":
            utils.UpsentTools()
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuFive()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//6. Crack Hash, Pcap & Brute Passwords....(Hashcat, jo)🔐
func passwordsCrackers() {
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.RED + "Cracker " + bcolors.BLUE + "1. " + bcolors.YELLOW + "Online " + bcolors.BLUE + "2. " + bcolors.YELLOW + "Offline " + bcolors.ENDC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        //Online crackers
        case "1":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuSixOne()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "OnLineCrackers" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    passwordsCrackers()
                    return
                case "1":
                    crackers.HydraSsh()
                case "2":
                    crackers.HydraFtp()
                case "3":
                    crackers.HydraSmb()
                case "4":
                    crackers.HydraRdp()
                case "5":
                    crackers.HydraLdap()
                case "6":
                    crackers.HydraSmtp()
                case "7":
                    crackers.HydraTelnet()
                case "8":
                    crackers.HydraHttps()
                case "9":
                    crackers.CyberBrute()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuSixOne()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //Offline crackers
        case "2":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuSixTwo()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "OffLineCrackers" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    passwordsCrackers()
                    return
                case "1":
                    crackers.AirCrackng()
                case "2":
                    crackers.JohnCrackng()
                case "8":
                    crackers.HashBuster()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuSixTwo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "99", "m", "menu":
            menus.MenuSix()
        case "00", "?", "h", "help":
            menus.HelpInfoSix()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//7. Social-Engineering Attacks............(Gophish, gi)🪝
func credsPhishers() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuSeven()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Phishers" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🪝" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "1":
            phishers.GoPhish()
        case "2":
            phishers.GoodGinx()
        case "3":
            phishers.ZPhisher()
        case "4":
            phishers.Blackeye()
        case "5":
            phishers.AdvPhisher()
        case "6":
            phishers.Darkphish()
        case "7":
            phishers.CyberPhish()
        case "8":
            phishers.SetoolKit()
        case "9":
            internaltargetInput()
            phishers.NinjaPhish(userTarget)
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuSeven()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//8. Website Attack Vectors................(Proxy, Conf)🎭
func askForProxy() *url.URL {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Enter proxy URL🛰️ Eg:" + bcolors.YELLOW + "(http://localhost:80)" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)

    scanner.Scan()
    proxyStr := scanner.Text()
    proxyStr = strings.TrimSpace(proxyStr)

    proxyURL, err := url.Parse(proxyStr)
    if err != nil {
        fmt.Println(bcolors.RED + "Invalid URL format. Please try again." + bcolors.ENDC)
        return askForProxy()
    }

    if proxyURL.Scheme != "http" && proxyURL.Scheme != "https" && proxyURL.Scheme != "socks5" && proxyURL.Scheme != "socks4" {
        fmt.Println(bcolors.RED + "Invalid scheme. Only http, https, socks5, or socks4 are allowed." + bcolors.ENDC)
        return askForProxy()
    }

    return proxyURL
}

func setProxyEnv(proxyURL *url.URL) error {
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return err
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return err
    }
    // SOCKS5 proxy is not directly supported by environment variables in all tools,
    // but you can use tools like `proxychains` to enable SOCKS5 support.
    return nil
}

//8. Website Attack Vectors................(userTarget, set)🎭
func websiteUserTarget() {
    utils.ClearScreen()
    banners.Banner()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "Enter Target:" + bcolors.DARKCYAN + "Either 📡HTTP(S)//: HOSTNAME or IP🌍" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🎯" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userTarget = scanner.Text()

    prefixes := []string{"http://", "https://", "www."}
    userXtarget = userTarget
    for _, prefix := range prefixes {
        userXtarget = strings.TrimPrefix(userXtarget, prefix)
    }

    userXtarget = strings.TrimSpace(userXtarget)

    if !strings.HasPrefix(userTarget, "http://") && !strings.HasPrefix(userTarget, "https://") {
        userXdomain = "http://" + userTarget
    } else {
        userXdomain = userTarget
    }

    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "Use proxies:" + bcolors.DARKCYAN + "Eg: SOCK4, SOCKS5, HTTP or HTTPS " + bcolors.YELLOW + "(y/n)" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
        case "y", "yes":
            proxyURL := askForProxy()

            if err := setProxyEnv(proxyURL); err != nil {
                fmt.Println("Error setting proxy environment variables:", err)
                return
            }
            return
        case "n", "no":
            return
        default:
            utils.SystemShell(userInput)
        }
    }
}

//8. Website Attack Vectors................(Attack, Phase)🌍
func websitesAttackers() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuEight()

    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Websites" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🪲" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        //1. Start Passive Web recon & SubuserXdomain Enumration.....🌍
        case "1":
             webattackers.WafW00f(userXtarget); webattackers.WhatWeb(userXtarget); webattackers.DigRecon(userXtarget); webattackers.WhoIs(userXtarget); webattackers.DnsRecon(userXtarget); webattackers.Ashock1(userTarget); webattackers.Sublist3r(userTarget); webattackers.OneForAll(userTarget); webattackers.Nuclei(userXtarget)
        //2. Gather e-mails & subuserXdomain namesfrom public sources🪰
        case "2":
            webattackers.TheHarvester(userTarget)
        //3. Start Bruteforcing Host's Root Files...............🦑
        case "3":
            webattackers. Gobuster(userTarget)
        //4. Start SQL, XSS & SSRF Detection & Eploitation......💉
        case "4":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuEightFour()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Injections" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─💉" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget)
                    webattackers.XsserAuto(userTarget)
                    webattackers.KatanaAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan()
                    webattackers.XsserMan()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightFour()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 4" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //5. Launch OWASP Nettacker project MainMenu............🦣
        case "5": 
            utils.ClearScreen()
            banners.Banner()
            menus.MenuEightFive()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Nettacker" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🦣" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.NetTacker1(userTarget)
                case "2":
                    webattackers.NetTacker2(userTarget)
                case "3":
                    webattackers.NetTacker3(userTarget)
                case "4":
                    webattackers.NetTacker4(userTarget)
                case "5":
                    webattackers.NetTacker5(userTarget)
                case "6":
                    webattackers.NetTacker6(userTarget)
                case "7":
                    webattackers.NetTacker7(userTarget)
                case "8":
                    webattackers.NetTacker8(userTarget)
                case "9":
                    webattackers.NetTacker9()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightFive()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu...🃏
        case "6":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuEightSix()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Jok3r" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🃏" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Jok3r1()
                case "2":
                    webattackers.Jok3r2()
                case "3":
                    webattackers.Jok3r3()
                case "4":
                    webattackers.Jok3r4()
                case "5":
                    webattackers.Jok3r5(userXdomain)
                case "6":
                    webattackers.Jok3r6(userXdomain)
                case "7":
                    webattackers.Jok3r7(userXdomain)
                case "8":
                    webattackers.Jok3r8()
                case "9":
                    webattackers.Jok3r9()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightSix()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //7. Osmedeus Next Generation Workflow Engine Main Menu.🐨
        case "7":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuEightSeven()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Osmedeus" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐨" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Osmedeus1()
                case "2":
                    webattackers.Osmedeus2(userTarget)
                case "3":
                    webattackers.Osmedeus3(userTarget)
                case "4":
                    webattackers.Osmedeus4()
                case "5":
                    webattackers.Osmedeus5(userTarget)
                case "6":
                    webattackers.Osmedeus6(userTarget)
                case "7":
                    webattackers.Osmedeus7(userTarget)
                case "8":
                    webattackers.Osmedeus8()
                case "9":
                    webattackers.Osmedeus9()
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightSeven()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //8. Ufonet Next Generation DDOS Tool Main Menu.........🦠
        case "8":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuEightEight()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Ufonet" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🦠" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Ufonet1()
                case "2":
                    webattackers.Ufonet2()
                case "3":
                    webattackers.Ufonet3(userTarget)
                case "4":
                    webattackers.Ufonet4(userTarget)
                case "5":
                    webattackers.Ufonet5(userTarget)
                case "6":
                    webattackers.Ufonet6(userTarget)
                case "7":
                    webattackers.Ufonet7()
                case "8":
                    webattackers.Ufonet8()
                case "9":
                    webattackers.Ufonet9(userTarget)
                case "set", "help set":
                    menus.HelpInfoSet()
                case "use", "help use":
                    menus.HelpInfoUse()
                case "run", "help run":
                    menus.HelpInfoRun()
                case "banner":
                    banners.Banner()
                case "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightEight()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //9. Launch Heavy Automation Attacks On The Host........🤖
        case "9":
             webattackers.WafW00f(userXtarget); webattackers.WhatWeb(userXtarget); webattackers.DigRecon(userXtarget); webattackers.WhoIs(userXtarget); webattackers.DnsRecon(userXtarget); webattackers.Ashock1(userTarget); webattackers.Sublist3r(userTarget); webattackers.Nuclei(userXtarget); webattackers.OneForAll(userTarget); webattackers.SeekOlver(userTarget); webattackers.Gobuster(userTarget); webattackers.Osmedeus3(userTarget); webattackers.ParamSpider(userTarget); webattackers.KatanaAuto(userTarget); webattackers.XsserAuto(userTarget); webattackers.Osmedeus6(userTarget); webattackers.NetTacker8(userTarget); webattackers.Jok3r6(userXdomain); webattackers.Nikto(userTarget); webattackers.Uniscan(userTarget)
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuEight()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//9. Help, Credits, Tricks and About.......(🕊  ︻╦╤─JC❤sU)
func creditsGivers() {
    utils.ClearScreen()
    banners.Banner()
    scriptures.Verse()
    credits.Contributors()
    credits.Developer()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Enter '0' 'exit' " + bcolors.DARKCYAN + "or" + bcolors.DARKGREY + bcolors.ITALIC + "'EXIT' " + bcolors.ENDC + "2 go back" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─📚" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            //
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//0. Exit africana-framework...............(Try option 00)
func scriptureNarators() {
    utils.ClearScreen()
    banners.Banner()
    scriptures.Narration()
    scriptures.CommandMents()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Enter '0' 'exit' " + bcolors.DARKCYAN + "or" + bcolors.DARKGREY + bcolors.ITALIC + "'EXIT' " + bcolors.ENDC + "2 go back" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─✍🏼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "set", "help set":
            menus.HelpInfoSet()
        case "use", "help use":
            menus.HelpInfoUse()
        case "run", "help run":
            menus.HelpInfoRun()
        case "banner":
            banners.Banner()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuZero()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//Africana-Framework ......................(The rolling 9)
func africanaFramework() {
    utils.ClearScreen()
    banners.Banner()
    menus.MenuZero()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🦊" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "0", "e", "q", "exit", "quit":
            os.Exit(0)
        case "b", "back":
        case "1", "use setups":
            systemSetups()
            return
        case "2", "use anonsurf":
            anonsurfSetups()
            return
        case "3", "use networks":
            internaltargetInput()
            internalAttackers()
            return
        case "4", "use malwares":
            malwareGenerators()
            return
        case "5", "use wireless":
            wirelessAttackers()
            return
        case "6", "use crackers":
            passwordsCrackers()
            return
        case "7", "use phishers":
            credsPhishers()
            return
        case "8", "use websites":
            websiteUserTarget()
            websitesAttackers()
            return
        case "9", "use credits":
            creditsGivers()
            return
        case "99", "use verses":
            scriptureNarators()
        case "banner":
            banners.Banner()
        case "history":
            subprocess.History()
        case "version":
            banners.Version()
        case "m", "menu":
            menus.MenuZero()
        case "info", "help info":
            menus.HelpInfoMain()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "clear history":
            subprocess.ClearHistory()
        case "bash", "use bash":
            utils.SystemShell(userInput)
        case "set", "info set", "help set":
            menus.HelpInfoSet()
        case "use", "info use", "help use":
            menus.HelpInfoUse()
        case "run", "info run", "help run":
            menus.HelpInfoRun()
        case "show", "info show", "help show":
            menus.HelpInfoShow()
        case "setups", "info setups", "help setups":
            menus.HelpInfoSetups()
        case "l", "list", "modules", "list modules", "show modules", "show all":
            menus.ListMainModules()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//...........................................(Main runner..)
func genesis() {
    if len(os.Args) < 2 {
        africanaFramework()
        return
    }
    command := os.Args[1]
    switch command {
    case "-v", "--version":
        banners.Banner()
        return
    case "-u", "--update":
        systemSetups()
        return
    case "-0", "-a", "--auto":
        africanaFramework()
        return
    case "-1", "-i", "--install":
        systemSetups()
        return
    case "-2", "-t", "--anonsurf":
        anonsurfSetups()
        return
    case "-3", "-n", "--networks":
        internaltargetInput()
        internalAttackers()
        return
    case "-4", "-m", "--malwares":
        malwareGenerators()
        return
    case "-5", "-w", "--wireless":
        wirelessAttackers()
        return
    case "-6", "-p", "--crackers":
        passwordsCrackers()
        return
    case "-7", "-f", "--phishers":
        credsPhishers()
        return
    case "-8", "-x", "--websites":
        websiteUserTarget()
        websitesAttackers()
        return
    case "-9", "-c", "--credits":
        creditsGivers()
        return
    case "-99", "-b", "--scriptures":
        scriptureNarators()
        return
    case "-00", "-h", "?", "--help":
        menus.HelpInfoMain()
        return
    default:
        menus.HelpInfoMain()
    }
}

func main() {
     africana()
}
