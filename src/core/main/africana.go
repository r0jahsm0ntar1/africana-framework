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

func Run() {
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuOne()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()
        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoKali()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoUbuntu()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoArch()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoWindows()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoAndroid()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoTermLogs()
        case "? 7", "info 7", "help 7", "info browselogs", "help browselogs", "browselogs":
            menus.HelpInfoBrowserLogs()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoClearLogs()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoUninstall()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListSetupsModules()

        case "1", "use 1", "run 1", "start 1", "run kali", "use kali", "start kali":
            utils.ClearScreen()
            banners.Banner()
            setups.KaliSetups()
            africanaFramework()
            return
        case "2", "use 2", "run 2", "start 2", "run ubuntu", "use ubuntu", "start ubuntu":
            utils.ClearScreen()
            banners.Banner()
            setups.UbuntuSetups()
            africanaFramework()
            return
        case "3", "use 3", "run 3", "start 3", "run arch", "use arch", "start arch":
            utils.ClearScreen()
            banners.Banner()
            setups.ArchSetups()
            africanaFramework()
            return
        case "4", "use 4", "run 4", "start 4", "run windows", "use windows", "start windows":
            utils.ClearScreen()
            banners.Banner()
            setups.WindowsSetups()
            menus.UpsentTools()
            africanaFramework()
            return
        case "5", "use 5", "run 5", "start 5", "run android", "use android", "start android":
            utils.ClearScreen()
            banners.Banner()
            setups.Androidinstall()
        case "6", "use 6", "run 6", "start 6", "run termlogs", "use termlogs", "start termlogs":
            utils.TermLogs()
            return
        case "7", "use 7", "run 7", "start 7", "run viewlogs", "use viewlogs", "start viewlogs":
            utils.BrowserLogs()
            return
        case "8", "use 8", "run 8", "start 8", "run clearlogs", "use clearlogs", "start clearlogs":
           utils.ClearLogs()
            return
        case "9", "use 9", "run 9", "start 9", "run uninstall", "use uninstall", "start uninstall":
            utils.ClearScreen()
            banners.Banner()
            setups.Uninstall()
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuTwo()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()
        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info anonsurf -i", "help anonsurf -i":
            menus.HelpInfoAnonsurfI()
        case "? 2", "info 2", "help 2", "info anonsurf -m", "help anonsurf -m":
            menus.HelpInfoAnonsurfM()
        case "? 3", "info 3", "help 3", "info anonsurf -r", "help anonsurf -r":
            menus.HelpInfoAnonsurfR()
        case "? 4", "info 4", "help 4", "info anonsurf -w", "help anonsurf -w":
            menus.HelpInfoAnonsurfW()
        case "? 5", "info 5", "help 5", "info anonsurf -a", "help anonsurf -a":
            menus.HelpInfoAnonsurfI()
        case "? 6", "info 6", "help 6", "info anonsurf -n", "help anonsurf -n":
            menus.HelpInfoAnonsurfN()
        case "? 7", "info 7", "help 7", "info anonsurf -s", "help anonsurf -s":
            menus.HelpInfoAnonsurfS()
        case "? 8", "info 8", "help 8", "info anonsurf -l", "help anonsurf -l":
            menus.HelpInfoAnonsurfL()
        case "? 9", "info 9", "help 9", "info anonsurf -e", "help anonsurf -e":
            menus.HelpInfoAnonsurfE()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListAnonsurfModules()

        case "1", "use 1", "run 1", "start 1", "anonsurf -i":
            securities.AnonsurfSetups()
        case "2", "use 2", "run 2", "start 2", "anonsurf -m":
            securities.AnonsurfStart()
        case "3", "use 3", "run 3", "start 3", "anonsurf -r":
            securities.AnonsurfExitnode()
        case "4", "use 4", "run 4", "start 4", "anonsurf -w":
            securities.AnonsurfStatus()
        case "5", "use 5", "run 5", "start 5", "anonsurf -a":
            securities.AnonsurfIpaddr()
        case "6", "use 6", "run 6", "start 6", "anonsurf -n":
            securities.AnonsurfRIptabls()
        case "7", "use 7", "run 7", "start 7", "anonsurf -s":
            securities.AnonsurfReload()
        case "8", "use 8", "run 8", "start 8", "anonsurf -l":
            securities.AnonsurfChains()
        case "9", "use 9", "run 9", "start 9", "anonsurf -e":
            securities.AnonsurfStop()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//3. Local Network Attack Vectors..........(Mitm, sniff)🐹
func internaltargetInput() {
    internals.InternalScan()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Networks " + bcolors.GREEN + "select your " + bcolors.RED + "Target!🎯" + bcolors.BLUE + ")" + bcolors.ENDC)
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuThree()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()
        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info discover", "help discover", "discover":
            menus.HelpInfoDiscover()
        case "? 2", "info 2", "help 2", "info portscan", "help portscan", "portscan":
            menus.HelpInfoPortScan()
        case "? 3", "info 3", "help 3", "info vulnscan", "help vulnscan", "vulnscan":
            menus.HelpInfoVulnScan()
        case "? 4", "info 4", "help 4", "info enumscan", "help enumscan", "enumscan":
            menus.HelpInfoEnumScan()
        case "? 5", "info 5", "help 5", "info smbexplo", "help smbexplo", "smbexplo":
            menus.HelpInfoExplScan()
        case "? 6", "info 6", "help 6", "info psniffer", "help psniffer", "psniffer":
            menus.HelpInfoPsniffer()
        case "? 7", "info 7", "help 7", "info responder", "help responder", "responder":
            menus.HelpInfoResponder()
        case "? 8", "info 8", "help 8", "info beefninja", "help beefninja", "beefninja":
            menus.HelpInfoBeefNinja()
        case "? 9", "info 9", "help 9", "info xsshooker", "help xsshooker", "xsshooker":
            menus.HelpInfoXssHoocker()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListInternalModules()

        case "1", "use 1", "run 1", "start 1", "run discover", "use discover", "start discover":
            internaltargetInput()
            utils.ClearScreen()
            banners.Banner()
            menus.MenuThree()
        case "2", "use 2", "run 2", "start 2", "run portscan", "use portscan", "start portscan":
            internals.NmapPortScan(userTarget)
        case "3", "use 3", "run 3", "start 3", "run vulnscan", "use vulnscan", "start vulnscan":
            internals.NmapVulnScan(userTarget)
        case "4", "use 4", "run 4", "start 4", "run enumscan", "use enumscan", "start enumscan":
            internals.SmbVulnScan(userTarget)
            internals.Enum4linux(userTarget)
            internals.EnumNxc(userTarget)
            internals.SmbMapScan(userTarget)
            internals.RpcEnumScan(userTarget)
        case "5", "use 5", "run 5", "start 5", "run smbexplo", "use smbexplo", "start smbexplo":
            internals.SmbVulnScan(userTarget)
            internals.SmbExploit(userTarget)
        case "6", "use 6", "run 6", "start 6", "run sniffer", "use sniffer", "start sniffer":
            internals.PacketSniffer(userTarget)
            internalAttackers()
        case "7", "use 7", "run 7", "start 7", "run responder", "use responder", "start responder":
            internals.PacketsResponder()
        case "8", "use 8", "run 8", "start 8", "run beefninja", "use beefninja", "start beefninja":
            internals.BeefInjector(userTarget)
        case "9", "use 9", "run 9", "start 9", "run ninjaphish", "use ninjaphish", "start ninjaphish":
            phishers.NinjaPhish(userTarget)
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuFour()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()
        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info androids", "help androids", "androids":
            menus.HelpInfoAndroids()
        case "? 2", "info 2", "help 2", "info iphones", "help iphones", "iphones":
            menus.HelpInfoIphones()
        case "? 3", "info 3", "help 3", "info windows", "help windows", "windows":
            menus.HelpInfoWindowsM()
        case "? 4", "info 4", "help 4", "info mackos", "help mackos", "mackos":
            menus.HelpInfoMackos()
        case "? 5", "info 5", "help 5", "info guineas", "help guineas", "guineas":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info websites", "help websites", "websites":
            menus.HelpInfoWebsitesM()
        case "? 7", "info 7", "help 7", "info backdoors", "help backdoors", "backdoors":
            menus.HelpInfoBackdoors()
        case "? 8", "info 8", "help 8", "info stealths", "help stealths", "stealths":
            menus.HelpInfoStealths()
        case "? 9", "info 9", "help 9", "info chosens", "help chosens", "chosens":
            menus.HelpInfoChosens()
        case "? 99", "info 99", "help 99", "info listeners", "help listeners", "listeners":
            menus.HelpInfoListeners()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMalwareModules()

        case "1", "use 1", "run 1", "start 1", "run androids", "use androids", "start androids":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourOne()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Androids" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourOne()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1", "use 1", "run 1", "start 1", "run teardroid", "use teardroid", "start teardroid":
                    butchers.TearDroid()
                case "2", "use 2", "run 2", "start 2", "run androrat", "use androrat", "start androrat":
                    butchers.AndroRat()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                     menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "2", "use 2", "run 2", "start 2", "run iphones", "use iphones", "start iphones":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourTwo()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Iphones" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourTwo()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1":
                    menus.UpsentTools()
                case "2":
                    menus.UpsentTools()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "3", "use 3", "run 3", "start 3", "run windows", "use windows", "start windows":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourThree()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Windows" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourThree()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1", "use 1", "run 1", "start 1", "run blackjack", "use blackjack", "start blackjack":
                    butchers.BlackJack()
                case "2", "use 2", "run 2", "start 2", "run shellz", "use shellz", "start shellz":
                    butchers.Shellz()
                case "3", "use 3", "run 3", "start 3", "run powerjoker", "use powerjoker", "start powerjoker":
                    butchers.PowerJoker()
                case "4", "use 4", "run 4", "start 4", "run meterpeter", "use meterpeter", "start meterpeter":
                    butchers.MeterPeter()
                case "5", "use 5", "run 5", "start 5", "run havoc", "use havoc", "start havoc":
                    butchers.Havoc()
                case "6":
                    butchers.HoaxMalware()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "4", "use 4", "run 4", "start 4", "run macos", "use macos", "start macos":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourFour()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "MackOS" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourFour()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info seashell", "help seashell", "seashell":
                    menus.HelpInfoSeaShell()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1":
                    butchers.SeaShell()
                case "2":
                    menus.UpsentTools()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "5", "use 5", "run 5", "start 5", "run linux", "use linux", "start linux":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourFive()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LinuxOS" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourFive()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1":
                    menus.UpsentTools()
                case "2":
                    menus.UpsentTools()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "6", "use 6", "run 6", "start 6", "run websites", "use websites", "start websites":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourSix()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Websites" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourSix()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1":
                    menus.UpsentTools()
                case "2":
                    menus.UpsentTools()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "7", "use 7", "run 7", "start 7", "run backdoors", "use backdoors", "start backdoors":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourSeven()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Univasals" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourSeven()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1", "use 1", "run 1", "start 1", "run noisemaker", "use noisemaker", "start noisemaker":
                    butchers.NoiseMakers()
                case "2", "use 2", "run 2", "start 2", "run codebreaker", "use codebreaker", "start codebreaker":
                    butchers.CodeBreakers()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "8", "use 8", "run 8", "start 8", "run stealths", "use stealths", "start stealths":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourEight()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Obfsications" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourEight()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1", "use 1", "run 1", "start 1", "run ghost", "use ghost", "start ghost":
                    butchers.Gh0x0st()
                case "2", "use 2", "run 2", "start 2", "run chameleon", "use chameleon", "start chameleon":
                    butchers.Chameleon()
                case "3":
                    menus.UpsentTools()
                case "4":
                    menus.UpsentTools()
                case "5":
                    menus.UpsentTools()
                case "6":
                    menus.UpsentTools()
                case "7":
                    menus.UpsentTools()
                case "8":
                    menus.UpsentTools()
                case "9":
                    menus.UpsentTools()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "9", "use 9", "run 9", "start 9", "run chosens", "use chosens", "start chosens":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuFourNine()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "ChosenOnes" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuFourNine()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    malwareGenerators()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info blackjack", "help blackjack", "blackjack":
                    menus.HelpInfoBlackJack()
                case "? 2", "info 2", "help 2", "info shellz", "help shellz", "shellz":
                    menus.HelpInfoShellz()
                case "? 3", "info 3", "help 3", "info powerjoker", "help powerjoker", "powerjoker":
                    menus.HelpInfoPowerJoker()
                case "? 4", "info 4", "help 4", "info meterpeter", "help meterpeter", "meterpeter":
                    menus.HelpInfoMeterPeter()
                case "? 5", "info 5", "help 5", "info havoc", "help havoc", "havoc":
                    menus.HelpInfoHavoc()
                case "? 6", "info 6", "help 6", "info tearndroid", "help tearndroid", "tearndroid":
                    menus.HelpInfoTearNdroid()
                case "? 7", "info 7", "help 7", "info androidrat", "help androidrat", "androidrat":
                    menus.HelpInfoAndroidRat()
                case "? 8", "info 8", "help 8", "info chameleon", "help chameleon", "chameleon":
                    menus.HelpInfoChameLeon()
                case "? 9", "info 9", "help 9", "info ghost", "help ghost", "ghost":
                    menus.HelpInfoGhost()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMalwareNineModules()

                case "1", "use 1", "run 1", "start 1", "run blackjack", "use blackjack", "start blackjack":
                    butchers.BlackJack()
                case "2", "use 2", "run 2", "start 2", "run shellz", "use shellz", "start shellz":
                    butchers.Shellz()
                case "3", "use 3", "run 3", "start 3", "run powerjoker", "use powerjoker", "start powerjoker":
                    butchers.PowerJoker()
                case "4", "use 4", "run 4", "start 4", "run meterpeter", "use meterpeter", "start meterpeter":
                    butchers.MeterPeter()
                case "5", "use 5", "run 5", "start 5", "run havoc", "use havoc", "start havoc":
                     butchers.Havoc()
                case "6", "use 6", "run 6", "start 6", "run tearndroid", "use tearndroid", "start teandroid":
                    butchers.TearDroid()
                case "7", "use 7", "run 7", "start 7", "run androidrat", "use androidrat", "start androidrat":
                    butchers.AndroRat()
                case "8", "use 8", "run 8", "start 8", "run chameleon", "use chameleon", "start chameleon":
                    butchers.Chameleon()
                case "9", "use 9", "run 9", "start 9", "run ghost", "use ghost", "start ghost":
                    butchers.Gh0x0st()
                default:
                    utils.SystemShell(userInput)
                }
            }
        case "99":
            butchers.ListenerLauncher()
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuFive()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info wifite", "help wifite", "wifite":
            menus.HelpInfoWifite()
        case "? 2", "info 2", "help 2", "info bettercap", "help bettercap", "bettercap":
            menus.HelpInfoBetterCap()
        case "? 3", "info 3", "help 3", "info wifipumpkin", "help wifipumpkin", "wifipumpkin":
            menus.HelpInfoWifiPumpkin()
        case "? 4", "info 4", "help 4":
            menus.UpsentTools()
        case "? 5", "info 5", "help 5":
            menus.UpsentTools()
        case "? 6", "info 6", "help 6", "info fluxion", "help fluxion", "fluxion":
            menus.HelpInfoFluxion()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8":
            menus.HelpInfoWifiPumpkin()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListWirelessModules()

        case "1", "use 1", "run 1", "start 1", "wifite -a", "run wifite -a", "use wifite -a", "start wifite -a":
            wireless.WifiteAuto()
        case "2", "use 2", "run 2", "start 2", "bettercap -a", "run bettercap -a", "use bettercap -a", "start bettercap -a":
            wireless.BettercapAuto()
        case "3", "use 3", "run 3", "start 3", "wifipumpkin -a", "run wifipumpkin -a", "use wifipumpkin -a", "start wifipumpkin -a":
            wireless.WifiPumpkin3Auto()
        case "4":
            menus.UpsentTools()
        case "5":
            menus.UpsentTools()
        case "6", "use 6", "run 6", "start 6", "fluxion -m", "run fluxion -m", "use fluxion -m", "start fluxion -m":
            wireless.FluxionMan()
        case "7", "use 7", "run 7", "start 7", "airgeddon -m", "run airgeddon -m", "use airgeddon -m", "start airgeddon -m":
            wireless.AirGeddon()
        case "8", "use 8", "run 8", "start 8", "wifipumpkin -m", "run wifipumpkin -m", "use wifipumpkin -m", "start wifipumpkin -m":
            wireless.WifiPumpkin3()
        case "9":
            menus.UpsentTools()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//6. Crack Hash, Pcap & Brute Passwords....(Hashcat, jo)🔐
func passwordsCrackers() {
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Crackers: " + bcolors.ENDC + bcolors.BLUE + "1. " + bcolors.YELLOW + "online " + bcolors.BLUE + "2. " + bcolors.YELLOW + "offline " + bcolors.BLUE + "0. " + bcolors.YELLOW + "Go back" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            //
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListCrackersModules()

        //Online crackers
        case "1", "use 1", "run 1", "start 1", "run onlinecrackers", "use onlinecrackers", "start onlinecrackers":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuSixOne()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "OnLineCrackers" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuSixOne()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    passwordsCrackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

                case "1", "use 1", "run 1", "start 1", "run hydrassh", "use hydrassh", "start hydrassh":
                    crackers.HydraSsh()
                case "2", "use 2", "run 2", "start 2", "run hydraftp", "use hydraftp", "start hydraftp":
                    crackers.HydraFtp()
                case "3", "use 3", "run 3", "start 3", "run hydrasmb", "use hydrasmb", "start hydrasmb":
                    crackers.HydraSmb()
                case "4", "use 4", "run 4", "start 4", "run hydrardp", "use hydrardp", "start hydrardp":
                    crackers.HydraRdp()
                case "5", "use 5", "run 5", "start 5", "run hydraldap", "use hydraldap", "start hydraldap":
                    crackers.HydraLdap()
                case "6", "use 6", "run 6", "start 6", "run hydrasmtp", "use hydrasmtp", "start hydrasmtp":
                    crackers.HydraSmtp()
                case "7", "use 7", "run 7", "start 7", "run hydratelnet", "use hydratelnet", "start hydratelnet":
                    crackers.HydraTelnet()
                case "8", "use 8", "run 8", "start 8", "run hydrahttps", "use hydrahttps", "start hydrahttps":
                    crackers.HydraHttps()
                case "9", "use 9", "run 9", "start 9", "run cyberbrute", "use cyberbrute", "start cyberbrute":
                    crackers.CyberBrute()
                default:
                    utils.SystemShell(userInput)
                }
            }
        //Offline crackers
        case "2", "use 2", "run 2", "start 2", "run offlinecrackers", "use offlinecrackers", "start offlinecrackers":
            utils.ClearScreen()
            banners.Banner()
            menus.MenuSixTwo()
            for {
                fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "OffLineCrackers" + bcolors.BLUE + ")" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "\n╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                scanner.Scan()
                userInput := scanner.Text()
                switch strings.ToLower(userInput) {
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuSixTwo()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    passwordsCrackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

                case "1", "use 1", "run 1", "start 1", "run aircrackng", "use aircrackng", "start aircrackng":
                    crackers.AirCrackng()
                case "2", "use 2", "run 2", "start 2", "run john", "use john", "start john":
                    crackers.JohnCrackng()
                case "8", "use 8", "run 8", "start 8", "run hashbuster", "use hashbuster", "start hashbuster":
                    crackers.HashBuster()
                default:
                    utils.SystemShell(userInput)
                }
            }
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuSeven()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info gophish", "help gophish", "gophish":
            menus.HelpInfoGophish()
        case "? 2", "info 2", "help 2", "info goodginx", "help goodginx", "goodginx":
            menus.HelpInfoGoodGinx()
        case "? 3", "info 3", "help 3", "info zphisher", "help zphisher", "zphisher":
            menus.HelpInfoZphisher()
        case "? 4", "info 4", "help 4", "info blackeye", "help blackeye", "blackeye":
            menus.HelpInfoBlackEye()
        case "? 5", "info 5", "help 5", "info advphisher", "help advphisher", "advphisher":
            menus.HelpInfoAdvnPhish()
        case "? 6", "info 6", "help 6", "info darkphish", "help darkphish", "darkphish":
            menus.HelpInfoDarkPhish()
        case "? 7", "info 7", "help 7", "info shellphish", "help shellphish", "shellphish":
            menus.HelpInfoShellPhish()
        case "? 8", "info 8", "help 8", "info setoolkit", "help setoolkit", "setoolkit":
            menus.HelpInfoSetToolkit()
        case "? 9", "info 9", "help 9", "info thc", "help thc", "thc", "info thehackerchoice", "help thehackerchoice", "thehackerchoice":
            menus.HelpInfoTheHackerChoice()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListPhishersModules()

        case "1", "use 1", "run 1", "start 1", "run gophish", "use gophish", "start gophish":
            phishers.GoPhish()
        case "2", "use 2", "run 2", "start 2", "run goodginx", "use goodginx", "start goodginx":
            phishers.GoodGinx()
        case "3", "use 3", "run 3", "start 3", "run zphisher", "use zphisher", "start zphisher":
            phishers.ZPhisher()
        case "4", "use 4", "run 4", "start 4", "run blackeye", "use blackeye", "start blackeye":
            phishers.Blackeye()
        case "5", "use 5", "run 5", "start 5", "run advphisher", "use advphisher", "start advphisher":
            phishers.AdvPhisher()
        case "6", "use 6", "run 6", "start 6", "run darkphish", "use darkphish", "start darkphish":
            phishers.Darkphish()
        case "7", "use 7", "run 7", "start 7", "run shellphish", "use shellphish", "start shellphish":
            phishers.ShellPhish()
        case "8", "use 8", "run 8", "start 8", "run setoolkit", "use setoolkit", "start setoolkit":
            phishers.SetoolKit()
        case "9", "use 9", "run 9", "start 9", "run thc",  "use thc", "start thc", "run thehackerchoice", "use thehackerchoice", "start thehackerchoice":
            phishers.Thc()
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            //
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMainModules()

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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuEight()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMainModules()

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
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightFour()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

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
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightFive()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

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
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightSix()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

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
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightSeven()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

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
                case "banner":
                    banners.Banner()
                case "sleep":
                    utils.Sleep()
                case "v", "version":
                    banners.Version()
                case "m", "menu":
                    menus.MenuEightEight()
                case "logs", "history":
                    subprocess.LogHistory()
                case "o", "junks", "outputs":
                    utils.ListJunks()
                case "? info", "info", "help info":
                    menus.HelpInfo()
                case "00", "?", "h", "help":
                    menus.HelpInfoMenuZero()
                case "e", "q", "exit", "quit":
                    os.Exit(0)
                case "0", "b", "back":
                    websitesAttackers()
                    return
                case "g", "t", "guide", "tutarial":
                    utils.BrowseTutarilas()

                case "clear logs", "clear history":
                    subprocess.ClearHistory()
                case "info set", "set", "help set":
                    menus.HelpInfoSet()
                case "clear junks", "clear outputs":
                    utils.ClearJunks()
                case "? use", "info use", "help use", "use":
                    menus.HelpInfoUse()
                case "? run", "info run", "help run", "run":
                    menus.HelpInfoRun()
                case "f", "use f", "use features", "features":
                    menus.HelpInfoFeatures()
                case "? tips", "info tips", "help tips", "tips":
                    menus.HelpInfoTips()
                case "? show", "info show", "help show", "show":
                    menus.HelpInfoShow()
                case "info list", "help list", "use list", "list":
                    menus.HelpInfoList()
                case "? start", "info start", "help start", "start":
                    menus.HelpInfoStart()

                case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
                    menus.HelpInfoSetups()
                case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
                    menus.HelpInfoAnonsurf()
                case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
                    menus.HelpInfoNetworks()
                case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
                    menus.HelpInfoMalwares()
                case "? 5", "info 5", "help 5", "info android", "help android", "android":
                    menus.HelpInfoWireless()
                case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
                    menus.HelpInfoCrackers()
                case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
                    menus.HelpInfoPhishers()
                case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
                    menus.HelpInfoWebsites()
                case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
                    menus.HelpInfoCredits()
                case "list all", "list modules", "show modules", "show all", "modules":
                    menus.ListMainModules()

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
                default:
                    utils.SystemShell(userInput)
                }
            }
        //9. Launch Heavy Automation Attacks On The Host........🤖
        case "9":
             webattackers.WafW00f(userXtarget); webattackers.WhatWeb(userXtarget); webattackers.DigRecon(userXtarget); webattackers.WhoIs(userXtarget); webattackers.DnsRecon(userXtarget); webattackers.Ashock1(userTarget); webattackers.Sublist3r(userTarget); webattackers.Nuclei(userXtarget); webattackers.OneForAll(userTarget); webattackers.SeekOlver(userTarget); webattackers.Gobuster(userTarget); webattackers.Osmedeus3(userTarget); webattackers.ParamSpider(userTarget); webattackers.KatanaAuto(userTarget); webattackers.XsserAuto(userTarget); webattackers.Osmedeus6(userTarget); webattackers.NetTacker8(userTarget); webattackers.Jok3r6(userXdomain); webattackers.Nikto(userTarget); webattackers.Uniscan(userTarget)
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            //
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMainModules()
        default:
            utils.SystemShell(userInput)
        }
    }
}

//0. Go  back.........................................(Try option 00)
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            //
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            africanaFramework()
            return
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            utils.ClearJunks()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            subprocess.ClearHistory()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info kali", "help kali", "kali":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info ubuntu", "help ubuntu", "ubuntu":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info arch", "help arch", "arch":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info windows", "help windows", "windows":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info android", "help android", "android":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info termlogs", "help termlogs", "tremlogs":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info viewlogs", "help viewlogs", "viewlogs":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info clearlogs", "help clearlogs", "clearlogs":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info uninstall", "help uninstall", "uninstall":
            menus.HelpInfoCredits()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMainModules()
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
        case "banner":
            banners.Banner()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuZero()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            //
        case "g", "t", "guide", "tutarial":
            utils.BrowseTutarilas()

        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run", "run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start", "start":
            menus.HelpInfoStart()

        case "? 1", "info 1", "help 1", "info setups", "help setups", "setups":
            menus.HelpInfoSetups()
        case "? 2", "info 2", "help 2", "info anonsurf", "help anonsurf", "anonsurf":
            menus.HelpInfoAnonsurf()
        case "? 3", "info 3", "help 3", "info networks", "help networks", "networks":
            menus.HelpInfoNetworks()
        case "? 4", "info 4", "help 4", "info malwares", "help malwares", "malwares":
            menus.HelpInfoMalwares()
        case "? 5", "info 5", "help 5", "info wireless", "help wireless", "wireless":
            menus.HelpInfoWireless()
        case "? 6", "info 6", "help 6", "info crackers", "help crackers", "crackers":
            menus.HelpInfoCrackers()
        case "? 7", "info 7", "help 7", "info phishers", "help phishers", "phishers":
            menus.HelpInfoPhishers()
        case "? 8", "info 8", "help 8", "info websites", "help websites", "websites":
            menus.HelpInfoWebsites()
        case "? 9", "info 9", "help 9", "info credits", "help credits", "credits":
            menus.HelpInfoCredits()
        case "? 99", "info 99", "info verses", "help verses", "verses":
            menus.HelpInfoVerses()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.ListMainModules()

        case "1", "use 1", "run 1", "start 1", "run setups", "use setups", "start setups":
            systemSetups()
            return
        case "2", "use 2", "run 2", "start 2", "run anonsurf", "use anonsurf", "start anonsurf":
            anonsurfSetups()
            return
        case "3", "use 3", "run 3", "start 3", "run networks", "use networks", "start networks":
            internaltargetInput()
            internalAttackers()
            return
        case "4", "use 4", "run 4", "start 4", "run malwares", "use malwares", "start malwares":
            malwareGenerators()
            return
        case "5", "use 5", "run 5", "start 5", "run wireless", "use wireless", "start wireless":
            wirelessAttackers()
            return
        case "6", "use 6", "run 6", "start 6", "run crackers", "use crackers", "start crackers":
            passwordsCrackers()
            return
        case "7", "use 7", "run 7", "start 7", "run phishers", "use phishers", "start phishers":
            credsPhishers()
            return
        case "8", "use 8", "run 8", "start 8", "run websites", "use websites", "start websites":
            websiteUserTarget()
            websitesAttackers()
            return
        case "9", "use 9", "run 9", "start 9", "run credits", "use credits", "start credits":
            creditsGivers()
            return
        case "99", "use 99", "run verses", "use verses", "start verses":
            scriptureNarators()
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
        banners.Version()
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
    case "-g", "--guide":
        utils.BrowseTutarilas()
    case "-00", "-h", "?", "--help":
        banners.Banner()
        menus.HelpInfoMenuMain()
        return
    default:
        menus.HelpInfoMenuMain()
    }
}
