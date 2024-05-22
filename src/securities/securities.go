package securities

import (
    "fmt"
    "configures"
    "subprocess"
)

func VanishSetups() {
    subprocess.Popen(`apt-get update; apt-get install -y tor squid privoxy dnsmasq iptables isc-dhcp-client isc-dhcp-server`)
}

func Vanishstart() {
    configures.Clean()
    subprocess.Popen(`cd ~/.africana/africana-base/tor-vanish/; python3 vanisher.py -m; systemctl restart tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service`)
    fmt.Println()
}

func Vanishstop() {
    subprocess.Popen(`cd ~/.africana/africana-base/tor-vanish/; python3 vanisher.py -e; systemctl stop tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service`)
    configures.Restore()
    fmt.Println()
}

func ChecktorStatus() {
    subprocess.Popen(`cd ~/.africana/africana-base/tor-vanish/; python3 vanisher.py -w`)
    fmt.Println()
}

func ChainsStatus() {
    subprocess.Popen(`tail -vf /var/log/privoxy/logfile`)
    fmt.Println()
}
