package securities

import(
    "os"
    "fmt"
    "time"
    "bufio"
    "utils"
    "menus"
    "regexp"
    "banners"
    "strings"
    "bcolors"
    "os/exec"
    "net/http"
    "io/ioutil"
    "subprocess"
)

var(
    userInput string
    scanner = bufio.NewScanner(os.Stdin)
)

func chameLeons() {
    if _, err := exec.LookPath("macchanger"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install macchanger'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }
    content := `# Generated by africana-framework. Delete at your own risk!!
[Unit]
Description=changes mac for %I
Wants=network.target
Before=network.target
BindsTo=sys-subsystem-net-devices-%i.device
After=sys-subsystem-net-devices-%i.device
[Service]
Type=oneshot
ExecStart=/usr/bin/macchanger -r %I
RemainAfterExit=yes
[Install]
WantedBy=multi-user.target
`
    filePath := "/etc/systemd/system/changemac@.service"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        err = ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }
    configDhclient()
    configDnsmasq()
    configSquid()
    configPrivoxy()
}

func configDhclient() {
    if _, err := exec.LookPath("dhclient"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install isc-dhcp-client isc-dhcp-server'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }
    filePath := "/etc/dhcp/dhclient.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp -r /etc/dhcp/dhclient.conf /etc/dhcp/dhclient.conf.bak_africana`)
        filesToReplacements := map[string]map[string]string{
            "/etc/dhcp/dhclient.conf": {
                "#prepend domain-name-servers 127.0.0.1;": "prepend domain-name-servers 127.0.0.1, 1.1.1.1, 1.0.0.1, 8.8.8.8, 8.8.4.4;",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func configDnsmasq() {
    if _, err := exec.LookPath("dnsmasq"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install dnsmasq'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }
    filePath := "/etc/dnsmasq.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp -r /etc/dnsmasq.conf /etc/dnsmasq.conf.bak_africana`)
        filesToReplacements := map[string]map[string]string{
            "/etc/dnsmasq.conf": {
                "#port=5353": "port=5353",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func configPrivoxy() {
    if _, err := exec.LookPath("privoxy"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install privoxy'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }
    filePath := "/etc/privoxy/privoxy.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp -r /etc/privoxy/config /etc/privoxy/privoxy.conf.bak_africana`)
        filesToReplacements := map[string]map[string]string{
            "/etc/privoxy/config": {
                "#debug     1": "debug   1",
                "#debug     2": "debug   2",
                "#debug  1024": "debug   1024",
                "#debug  4096": "debug   4096",
                "#        forward-socks5t   /               127.0.0.1:9050 .": "forward-socks5t   /               127.0.0.1:9050 .",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func configSquid() {
    if _, err := exec.LookPath("squid"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install squid'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }
    filePath := "/etc/squid/squid.conf.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`cp -r /etc/squid/squid.conf /etc/squid/squid.conf.bak_africana`)
        filesToReplacements := map[string]map[string]string{
            "/etc/squid/squid.conf": {
                "http_port 3128": "http_port 3129\nnever_direct allow all\nshutdown_lifetime 0 seconds\ncache_peer localhost parent 8118 7 no-digest no-query",
            },
        }
        utils.Editors(filesToReplacements)
    }
}

func iP() string {
    resp, err := http.Get("https://ifconfig.me")
    if err != nil {
        resp, err = http.Get("https://ipinfo.io/ip")
        if err != nil {
            fmt.Printf("\n%s\n", err)
            fmt.Println(bcolors.Colors() + "\nSorry, can't fetch the details.\nEither the site's down or something's up with your internet configuration.\nYou may find the solution here :)\nhttps://github.com/Feliz-SZK/Linux-Decoded/blob/master/Fix%20temporary%20failure%20in%20name%20resolution.md" + bcolors.ENDC)
            os.Exit(1)
        }
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}

func flaG() string {
    cmd := exec.Command("iptables", "-t", "nat", "-L", "-n")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("\n%sencountered some hiccups while checking the iptables details%s\n", bcolors.RED, bcolors.ENDC)
    }
    return string(output)
}

func cypheR(backoff int, resolvSwitch int) int {
    if backoff > 5 {
        fmt.Printf("\n%sExceeded number of retries, terminating to prevent memory corruption.%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }

    if _, err := os.Stat("/etc/resolv.conf"); os.IsNotExist(err) {
        fmt.Printf("\n%sresolv.conf file is missing, %syou want me to manually create it for you? %sY/N: %s", bcolors.BLUE, bcolors.BLUE, bcolors.GREEN, bcolors.ENDC)
        var askin string
        fmt.Scanln(&askin)
        if strings.ToLower(askin) == "y" {
            file, err := os.Create("/etc/resolv.conf")
            if err != nil {
                fmt.Printf("%ssomething's wrong, can't write the file.%s\n%s", bcolors.RED, bcolors.ENDC, err)
                os.Exit(1)
            }
            defer file.Close()
            file.WriteString("# Generated by africana-framework. Delete at your own risk!!\nnameserver 127.0.0.1\nnameserver 1.1.1.1\nnameserver 1.0.0.1\nnameserver 8.8.8.8\nnameserver 8.8.4.4")
            resolvSwitch++
            fmt.Printf("%sDone, saved with local DNS.%s\n", bcolors.Colors(), bcolors.ENDC)
        } else if strings.ToLower(askin) == "n" {
            fmt.Printf("%sRoger that, terminating....%s\n", bcolors.GREEN, bcolors.ENDC)
            os.Exit(1)
        } else {
            backoff++
            return cypheR(backoff, resolvSwitch)
        }
    }
    return resolvSwitch
}

func resolvConfig(rSwitch int) int {
    resolvString := "# Generated by africana-framework. Delete at your own risk!!\nnameserver 127.0.0.1\nnameserver 1.1.1.1\nnameserver 1.0.0.1\nnameserver 8.8.8.8\nnameserver 8.8.4.4"
    if rSwitch == 0 {
        file, err := os.Open("/etc/resolv.conf")
        if err != nil {
            fmt.Printf("%s\n", err)
            os.Exit(1)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        found := false
        for scanner.Scan() {
            if strings.Contains(scanner.Text(), resolvString) {
                found = true
                break
            }
        }

        if !found {
            fmt.Println(bcolors.BLUE + "(" + bcolors.ORANGE + "Configuring resolv.conf.." + bcolors.ENDC)
            time.Sleep(400 * time.Millisecond)
            cmd := exec.Command("cp", "/etc/resolv.conf", "/etc/resolv.conf.backup_anonsurf")
            cmd.Run()
            file, err := os.Create("/etc/resolv.conf")
            if err != nil {
                fmt.Printf("%s\n", err)
                os.Exit(1)
            }
            defer file.Close()
            file.WriteString(resolvString + "\n")
            fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
        } else {
            fmt.Println(bcolors.GREEN + bcolors.BOLD + "Configuring resolv.conf" + bcolors.ENDC)
            time.Sleep(400 * time.Millisecond)
            fmt.Println(bcolors.Colors() + bcolors.BOLD + " :) Already Configured" + bcolors.ENDC)
            time.Sleep(500 * time.Millisecond)
        }
    } else {
        fmt.Println(bcolors.GREEN + bcolors.BOLD + "Configuring resolv.conf" + bcolors.ENDC)
        time.Sleep(400 * time.Millisecond)
        fmt.Println(bcolors.Colors() + " :) Already Configured" + bcolors.ENDC)
    }
    return 0
}

func confiGure() {
    if strings.Contains(flaG(), "anonsurf") {
        fmt.Print(bcolors.DARKCYAN + "\n              Anonsurf: " + bcolors.ENDC + "is already running....\n" + bcolors.ENDC)
        os.Exit(1)
    }

    rSwitch := cypheR(0, 0)

    if _, err := exec.LookPath("tor"); err != nil {
        fmt.Printf("\n%sTor isn't installed, install it with 'sudo apt install tor'%s\n", bcolors.RED, bcolors.ENDC)
        os.Exit(1)
    }

    if _, err := os.Stat("/etc/tor/torrc"); os.IsNotExist(err) {
        fmt.Printf("%sNo torrc file is configured.....%s%sConfiguring:)%s\n", bcolors.RED, bcolors.ENDC, bcolors.GREEN, bcolors.ENDC)
        file, err := os.Create("/etc/tor/torrc")
        if err != nil {
            fmt.Printf("%sFailed to write the torrc file%s \n %s", bcolors.RED, bcolors.ENDC, err)
            os.Exit(1)
        }
        defer file.Close()
        for _, element := range torrString {
            file.WriteString(element + "\n")
        }
        fmt.Printf("%sDone....%s\n", bcolors.CYAN, bcolors.ENDC)
    } else {
        fmt.Println(bcolors.BLUE + "\n(" + bcolors.ORANGE + "Configuring Torrc.." + bcolors.ENDC)
        time.Sleep(400 * time.Millisecond)
        cmd := exec.Command("cp", "/etc/tor/torrc", "/etc/tor/torrc.bak_anonsurf")
        cmd.Run()
        file, err := os.Create("/etc/tor/torrc")
        if err != nil {
            fmt.Printf("%s\n", err)
            os.Exit(1)
        }
        defer file.Close()
        for _, element := range torrString {
            file.WriteString(element + "\n")
        }
        fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    }

    resolvConfig(rSwitch)
    fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Configuring dnsmasq, privoxy, squid..\n" + bcolors.ENDC)
    chameLeons()
    fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Starting macchanger, dnsmasq, privoxy, squid..\n" + bcolors.ENDC)
    subprocess.Popen(`systemctl restart changemac@eth0.service dnsmasq.service tor@default.service privoxy.service squid.service`)
    fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
}

func termiNate() {
    trigger := 0
    if _, err := os.Stat("/etc/resolv.conf.backup_anonsurf"); err == nil {
        trigger++
        cmd := exec.Command("mv", "/etc/resolv.conf.backup_anonsurf", "/etc/resolv.conf")
        cmd.Stdin = strings.NewReader("yes\n")
        cmd.Run()
        fmt.Print(bcolors.BLUE + "\n(" + bcolors.ORANGE + "Reverting to default resolv.conf..\n" + bcolors.ENDC)
        time.Sleep(500 * time.Millisecond)
        fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    }

    if _, err := os.Stat("/etc/tor/torrc.bak_anonsurf"); err == nil {
        trigger++
        cmd := exec.Command("mv", "/etc/tor/torrc.bak_anonsurf", "/etc/tor/torrc")
        cmd.Stdin = strings.NewReader("yes\n")
        cmd.Run()
        fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Dropping torrc file & restoring default..\n" + bcolors.ENDC)
        time.Sleep(500 * time.Millisecond)
        fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    }

    if strings.Contains(flaG(), "anonsurf") {
        trigger++
        fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Restoring Default Iptables rules..\n" + bcolors.ENDC)
        if _, err := os.Stat("/etc/iptables_rules_anonsurf.bak"); err == nil {
            cmd := exec.Command("iptables-restore", "<", "/etc/iptables_rules_anonsurf.bak")
            cmd.Run()
            os.Remove("/etc/iptables_rules_anonsurf.bak")
        } else {
            resetToDefault(true, true)
        }
        time.Sleep(1 * time.Second)
    }

    if trigger == 0 {
        fmt.Printf("\n%s%sNo instances of anonsurf have been executed%s\n", bcolors.RED, bcolors.BOLD, bcolors.ENDC)
        fmt.Printf("%s%s[Exiting...]%s\n", bcolors.GREEN, bcolors.BOLD, bcolors.ENDC)
        os.Exit(1)
    } else {
        fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    }
    fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Stoping macchanger, dnsmasq, privoxy, squid..\n" + bcolors.ENDC)
    subprocess.Popen(`systemctl stop changemac@eth0.service dnsmasq.service squid.service privoxy.service tor@default.service`)
    fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
}

func torCircuit() {
    if !strings.Contains(flaG(), "anonsurf") {
        fmt.Printf("\n%sYou gotta start anonsurf first%s\n", bcolors.Colors(), bcolors.ENDC)
        os.Exit(1)
    }
    exec.Command("service", "tor", "reload").Run()
    fmt.Println(bcolors.GREEN + "\nScrambling Tor Nodes" + bcolors.ENDC)
    time.Sleep(400 * time.Millisecond)
    fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    time.Sleep(400 * time.Millisecond)
    fmt.Printf("%sYour new IP appears to be: %s%s%s\n", bcolors.GREEN, bcolors.ENDC, bcolors.Colors(), bcolors.BOLD, iP())
}

func torStatus(Timer float64) {
    fmt.Print(bcolors.DARKGREEN + "\nConfarming if your system is secure....\n" + bcolors.ENDC)
    resp, err := http.Get("https://check.torproject.org")
    if err != nil {
        Timer++
        if Timer > 2 {
            fmt.Println(bcolors.RED + "\nUnable to get network details, " + bcolors.ORANGE + "Check Internet Connections & " + bcolors.CYAN + "retry....\n" + bcolors.ENDC)
            fmt.Println(bcolors.RED + "Error:" + bcolors.DARKCYAN, err, bcolors.ENDC)
            fmt.Println()
            return
        }
        fmt.Println(bcolors.GREEN + "\nHaving trouble fetching exit-node details,     " + bcolors.CYAN + "retrying...." + bcolors.ENDC)
        time.Sleep(9 * time.Second)
        torStatus(Timer)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    title := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`).FindStringSubmatch(string(body))[1]
    ip := regexp.MustCompile(`[0-9]+(?:\.[0-9]+){3}`).FindString(string(body))

    fmt.Printf("\nYour IP address is: %s%s%s%s\n", bcolors.Colors(), bcolors.BOLD, ip, bcolors.ENDC)
    if strings.Contains(title, "Congratulations") {
        fmt.Printf("%sCongratulations, you're using Tor :)%s\n", bcolors.Colors(), bcolors.ENDC)
    } else {
        fmt.Printf("%s%s%s\n", bcolors.Colors(), strings.TrimSpace(title), bcolors.ENDC)
    }
}

func checkDefaultRules() (string, string) {
    cmd := exec.Command(`iptables-save | grep '^\-' | wc -l`)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", string(output)
    }
    return string(output), ""
}

func configureFirewall() {
    fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Backing up Iptables..\n" + bcolors.ENDC)
    firewallGreen, firewallRed := checkDefaultRules()
    if firewallRed != "" {
        fmt.Printf("%s\nCan't execute %siptables-save%s. See the reason below.\n%s%s%s\n", bcolors.RED, bcolors.BLUE, bcolors.ENDC, bcolors.RED, firewallRed, bcolors.ENDC)
        os.Exit(1)
    }
    if strings.TrimSpace(firewallGreen) == "0" {
        fmt.Printf(" %sDefault rules are configured, skipping..%s\n", bcolors.BLUE, bcolors.ENDC)
    } else {
        cmd := exec.Command("iptables-save", ">", "/etc/iptables_rules_anonsurf.bak")
        cmd.Run()
        fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")\n" + bcolors.ENDC)
    }

    innOutRules := `
### Set variables
# The UID that Tor runs as (varies from system to system)

_tor_uid=$(id -u debian-tor) #Debian/Ubuntu

# Tor's TransPort
_trans_port="9040"

# Tor's DNSPort
_dns_port="5353"

# Tor's VirtualAddrNetworkIPv4
_virt_addr="10.192.0.0/10"

# LAN destinations that shouldn't be routed through Tor
_non_tor="127.0.0.0/8 10.0.0.0/8 172.16.0.0/12 192.168.0.0/16"

# Other IANA reserved blocks (These are not processed by tor and dropped by default)
_resv_iana="0.0.0.0/8 100.64.0.0/10 169.254.0.0/16 192.0.0.0/24 192.0.2.0/24 192.88.99.0/24 198.18.0.0/15 198.51.100.0/24 203.0.113.0/24 224.0.0.0/4 240.0.0.0/4 255.255.255.255/32"

# Flushing existing Iptables Chains/Firewall rules #
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT

iptables -F
iptables -X
iptables -Z

iptables -t nat -F
iptables -t nat -X
iptables -t mangle -F
iptables -t mangle -X
iptables -t raw -F
iptables -t raw -X

### *nat OUTPUT (For local redirection)
# nat .onion addresses
iptables -t nat -A OUTPUT -d $_virt_addr -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports $_trans_port

# nat dns requests to Tor
iptables -t nat -A OUTPUT -d 127.0.0.1/32 -p udp -m udp --dport 53 -j REDIRECT --to-ports $_dns_port -m comment --comment "anonsurf_triggered"

# Don't nat the Tor process, the loopback, or the local network
iptables -t nat -A OUTPUT -m owner --uid-owner $_tor_uid -j RETURN
iptables -t nat -A OUTPUT -o lo -j RETURN

# Allow lan access for hosts in $_non_tor and $_resv_ina
# This is to make sure that this local addresses don't get dropped.
for _lan in $_non_tor; do
  iptables -t nat -A OUTPUT -d $_lan -j RETURN
done

for _iana in $_resv_iana; do
  iptables -t nat -A OUTPUT -d $_iana -j RETURN
done

# Redirect all other pre-routing and output to Tor's TransPort
iptables -t nat -A OUTPUT -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -j REDIRECT --to-ports $_trans_port

### *filter INPUT
iptables -A INPUT -m state --state ESTABLISHED -j ACCEPT
iptables -A INPUT -i lo -j ACCEPT

# Log & Drop everything else. Uncomment to enable logging
#iptables -A INPUT -j LOG --log-prefix "Dropped INPUT packet: " --log-level 7 --log-uid
iptables -A INPUT -j DROP

### *filter FORWARD
iptables -A FORWARD -j DROP

### Fix for possible kernel packet-leak as discussed in,
### https://lists.torproject.org/pipermail/tor-talk/2014-March/032507.html
### uncomment below lines to log dropped packets

iptables -A OUTPUT -m conntrack --ctstate INVALID -j DROP
# iptables -A OUTPUT -m state --state INVALID -j LOG --log-prefix "Transproxy state leak blocked: " --log-uid
iptables -A OUTPUT -m state --state INVALID -j DROP

### *filter OUTPUT
iptables -A OUTPUT -m state --state ESTABLISHED -j ACCEPT

# Allow Tor process output
iptables -A OUTPUT -m owner --uid-owner $_tor_uid -p tcp -m tcp --tcp-flags FIN,SYN,RST,ACK SYN -m state --state NEW -j ACCEPT

# Allow loopback output
iptables -A OUTPUT -d 127.0.0.1/32 -o lo -j ACCEPT

# Tor transproxy magic
iptables -A OUTPUT -d 127.0.0.1/32 -p tcp -m tcp --dport $_trans_port --tcp-flags FIN,SYN,RST,ACK SYN -j ACCEPT

# Drop everything else.
iptables -A OUTPUT -j DROP

### Set default policies to DROP
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT DROP
`
    fmt.Print(bcolors.BLUE + "(" + bcolors.ORANGE + "Setting up firewall rules..\n" + bcolors.ENDC)
    cmd := exec.Command("/bin/bash", "-c", innOutRules)
    cmd.Run()
    fmt.Print("                                                    " + bcolors.GREEN + "Done " + bcolors.ORANGE + "✔" + bcolors.BLUE +")" + bcolors.ENDC)
}

func resetToDefault(overidePass bool, resetAsChildFunc bool) {
    var resetTrigger int
    if !overidePass {
        if resetTrigger > 7 {
            fmt.Printf("%sexiting to prevent memory corruption.%s\n", bcolors.RED, bcolors.ENDC)
            os.Exit(1)
        }
        var resetConsent string
        fmt.Printf("%s\nThis will overwrite all of your existing rules %sY(do it)%s/%sN(exit)%s: ", bcolors.Colors(), bcolors.GREEN, bcolors.ENDC, bcolors.RED, bcolors.ENDC)
        fmt.Scanln(&resetConsent)
        if strings.ToLower(resetConsent) == "y" {
            //continue
        } else if strings.ToLower(resetConsent) == "n" {
            fmt.Printf("%sCopy that..\n%s", bcolors.RED, bcolors.ENDC)
            os.Exit(1)
        } else {
            resetTrigger++
            resetToDefault(overidePass, resetAsChildFunc)
            return
        }

        time.Sleep(1 * time.Second)
        fmt.Printf("%sBacking up current rules, just in case..%s\n", bcolors.PURPLE, bcolors.ENDC)

        defaultCheckOne, defaultCheckNeo := checkDefaultRules()
        if defaultCheckNeo != "" {
            fmt.Printf("%sError while checking existing rules; %sexiting..\n%sError message: %s%s%s\n", bcolors.RED, bcolors.ORANGE, bcolors.YELLOW, bcolors.Colors(), defaultCheckNeo, bcolors.ENDC)
            os.Exit(1)
        }
        if strings.TrimSpace(defaultCheckOne) != "0" {
            fileNameID := time.Now().Format("01_02_2006-15:04:05")
            cmd := exec.Command("sudo", "iptables-save", ">", "/tmp/iptables_"+fileNameID+".rules")
            cmd.Run()
            fmt.Printf("%sSaved in %s/tmp%s as %siptables_%s.rules%s\n\n", bcolors.CYAN, bcolors.BLUE, bcolors.ENDC, bcolors.RED, fileNameID, bcolors.ENDC)
        } else {
            fmt.Printf("%s Default rules are set, backup not required :)%s\n", bcolors.ORANGE, bcolors.ENDC)
        }
        fmt.Printf("%s%sResetting Iptables%s\n", bcolors.ORANGE, bcolors.BOLD, bcolors.ENDC)
    }
    iptablesRules := `
# Accepting all traffic first#
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT

# Flushing All Iptables Chains/Firewall rules #
iptables -F

# Deleting all Iptables Chains #
iptables -X

# Flushing all counters too #
iptables -Z
# Flush and delete all nat and  mangle #
iptables -t nat -F
iptables -t nat -X
iptables -t mangle -F
iptables -t mangle -X
iptables -t raw -F
iptables -t raw -X
`
    cmd := exec.Command("/bin/bash", "-c", iptablesRules)
    output, err := cmd.CombinedOutput()
    time.Sleep(500 * time.Millisecond)
    if err != nil {
        fmt.Printf("%sCan't reset Iptables%s\n", bcolors.Colors(), bcolors.ENDC)
        fmt.Printf("%s%s%s\n", bcolors.Colors(), string(output), bcolors.ENDC)
        os.Exit(1)
    }
    if !resetAsChildFunc {
        fmt.Printf("%s Successfully reset Iptables to default :)%s\n", bcolors.BLUE, bcolors.ENDC)
    }
}

var torrString = []string{
    "# Generated by africana-framework. Delete at your own risk!!\n",
    "VirtualAddrNetwork 10.192.0.0/10",
    "AutomapHostsOnResolve 1",
    "CookieAuthentication 1",
    "TransPort 9040",
    "SocksPort 9050",
    "DNSPort 5353",
}

func AnonsurfSetups() {
    banners.GraphicsTorNet()
    fmt.Println()
    subprocess.Popen(`apt-get update; apt-get install -y tor squid privoxy dnsmasq iptables isc-dhcp-client isc-dhcp-server`)
    fmt.Println()
}

func AnonsurfStart() {
    banners.GraphicsTorNet()
    fmt.Println()
    confiGure()
    configureFirewall()
    torStatus(0)
}

func AnonsurfExitnode() {
    banners.GraphicsTorNet()
    fmt.Println()
    torCircuit()
}

func AnonsurfStatus() {
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " anonsurf(" + bcolors.RED +  "src/securities/show_status.fn" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    for {
        switch strings.ToLower(userInput) {
            case "banner":
                banners.RandomBanners()
                AnonsurfStatus()
                return
            case "sleep":
                utils.Sleep()
                AnonsurfStatus()
                return
            case "info":
                menus.HelpInfoAnonsurf()
                AnonsurfStatus()
                return
            case "v", "version":
                banners.Version()
                AnonsurfStatus()
                return
            case "options", "show options":
                menus.HelpInfoExecute()
                AnonsurfStatus()
                return
            case "m", "menu":
                menus.MenuTwo()
                AnonsurfStatus()
                return
            case "logs", "history":
                subprocess.LogHistory()
                AnonsurfStatus()
                return
            case "o", "junks", "outputs":
                utils.ListJunks()
                AnonsurfStatus()
                return
            case "? info", "help info":
                menus.HelpInfo()
                AnonsurfStatus()
                return
            case "00", "?", "h", "help":
                menus.HelpInfoMenuZero()
                AnonsurfStatus()
                return
            case "e", "q", "exit", "quit":
                os.Exit(0)
            case "0", "b", "back":
                return
            case "g", "t", "guide", "tutarial":
                utils.BrowseTutarilas()
                AnonsurfStatus()
                return
            case "clear logs", "clear history":
                subprocess.ClearHistory()
                AnonsurfStatus()
                return
            case "info set", "set", "help set":
                menus.HelpInfoSet()
                AnonsurfStatus()
                return
            case "clear junks", "clear outputs":
                utils.ClearJunks()
                AnonsurfStatus()
                return
            case "? run", "info run", "help run":
                menus.HelpInfoRun()
                AnonsurfStatus()
                return
            case "? start", "info start", "help start":
                menus.HelpInfoStart()
                AnonsurfStatus()
                return
            case "? use", "info use", "help use", "use":
                menus.HelpInfoUse()
                AnonsurfStatus()
                return
            case "f", "use f", "use features", "features":
                menus.HelpInfoFeatures()
                AnonsurfStatus()
                return
            case "? tips", "info tips", "help tips", "tips":
                menus.HelpInfoTips()
                AnonsurfStatus()
                return
            case "? show", "info show", "help show", "show":
                menus.HelpInfoShow()
                AnonsurfStatus()
                return
            case "info list", "help list", "use list", "list":
                menus.HelpInfoList()
                AnonsurfStatus()
                return
            case "? options", "info options", "help options":
                menus.HelpInfOptions()
                AnonsurfStatus()
                return
            case "run", "start", "execute":
                banners.GraphicsTorNet()
                fmt.Println()
                subprocess.Popen(`systemctl --no-pager -l status changemac@eth0.service dnsmasq.service squid.service privoxy.service tor@default.service`)
                fmt.Println()
                AnonsurfStatus()
                return
            default:
                utils.SystemShell(userInput)
                AnonsurfStatus()
                return
         }
    }
}

func AnonsurfIpaddr() {
    banners.GraphicsTorNet()
    fmt.Println()
    torStatus(0)
}

func AnonsurfRIptabls() {
    banners.GraphicsTorNet()
    fmt.Println()
    resetToDefault(false, false)
}

func AnonsurfReload() {
    banners.GraphicsTorNet()
    fmt.Println()
    confiGure()
    configureFirewall()
    torStatus(0)
}

func AnonsurfChains() {
    banners.GraphicsTorNet()
    fmt.Println()
    subprocess.Popen(`tail -vf /var/log/privoxy/logfile`)
    fmt.Println()
}

func AnonsurfStop() {
    banners.GraphicsTorNet()
    fmt.Println()
    termiNate()
}
