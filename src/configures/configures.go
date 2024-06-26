package configures

import (
    "fmt"
    "io/ioutil"
    "strings"
    "subprocess"
    "bcolors"
    "os"
)

func replaceStringsInFile(fileName string, replacements map[string]string) error {
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        return err
    }

    textContent := string(content)

    for old, new := range replacements {
        textContent = strings.Replace(textContent, old, new, -1)
    }

    return ioutil.WriteFile(fileName, []byte(textContent), 0644)
}

func replaceInMultipleFiles(filesToReplacements map[string]map[string]string) {
    for fileName, replacements := range filesToReplacements {
        err := replaceStringsInFile(fileName, replacements)
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error replacing strings in file %s: %v\n", fileName, err)
        } else {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Replacements completed successfully in file %s!\n", fileName)
        }
    }
}

func Changemac() {
    content := `
# Generated by africana-framework. Delete at your own risk!
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
}

func Clean() {
    subprocess.Popen(`cp -r /etc/tor/torrc /etc/tor/torrc.bak_africana; cp -r /etc/privoxy/config /etc/privoxy/config.bak_africana; cp -r /etc/squid/squid.conf /etc/squid/squid.conf.bak_africana; cp -r /etc/dhcp/dhclient.conf /etc/dhcp/dhclient.conf.bak_africana; cp -r /etc/dnsmasq.conf /etc/dnsmasq.conf.bak_africana`)
    filesToReplacements := map[string]map[string]string{
        "/etc/tor/torrc": {
            "#ORPort 9001": "DNSPort 5353",
            "#SocksPort 9050": "SocksPort 9050",
            "#ControlPort 9051": "ControlPort 9060",
            "#Nickname ididnteditheconfig": "Nickname veteran",
            "#CookieAuthentication 1": "CookieAuthentication 1",
            "#Address noname.example.com": "AutomapHostsOnResolve 1",
            "# OutboundBindAddress 10.0.0.5": "VirtualAddrNetworkIPv4 10.192.0.0/10",
            "##ORPort 127.0.0.1:9090 NoAdvertise": "TransPort 9040 IsolateClientAddr IsolateClientProtocol IsolateDestAddr IsolateDestPort",
        },
        "/etc/privoxy/config": {
            "#debug     1": "debug   1",
            "#debug     2": "debug   2",
            "#debug  1024": "debug   1024",
            "#debug  4096": "debug   4096",
            "#        forward-socks5t   /               127.0.0.1:9050 .": "forward-socks5t   /               127.0.0.1:9050 .",
        },
        "/etc/squid/squid.conf": {
            "#   cache_peer cdn.example.com      sibling   3128     0": "cache_peer 127.0.0.1 parent 8118 7 no-query no-digest",
            "http_port 3128": "http_port 3129",
        },
        "/etc/dhcp/dhclient.conf": {
            "#prepend domain-name-servers 127.0.0.1;": "prepend domain-name-servers 127.0.0.1, 1.1.1.1, 1.0.0.1, 8.8.8.8, 8.8.4.4;",
        },
        "/etc/dnsmasq.conf": {
            "#port=5353": "port=5353",
        },
    }
    replaceInMultipleFiles(filesToReplacements)
    Changemac()
}

func Restore() {
    subprocess.Popen(`rm -rf /etc/tor/torrc; mv /etc/tor/torrc.bak_africana /etc/tor/torrc; rm -rf /etc/privoxy/config; mv /etc/privoxy/config.bak_africana /etc/privoxy/config; rm -rf /etc/squid/squid.conf; mv /etc/squid/squid.conf.bak_africana /etc/squid/squid.conf; rm -rf /etc/dhcp/dhclient.conf; mv /etc/dhcp/dhclient.conf.bak_africana /etc/dhcp/dhclient.conf; rm -rf /etc/dnsmasq.conf; mv /etc/dnsmasq.conf.bak_africana /etc/dnsmasq.conf`)
}
