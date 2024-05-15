package securities

import (
    "fmt"
    "subprocess"
)

func VanishSetups() {
    subprocess.Popen(`apt-get update; apt-get install -y tor squid privoxy iptables isc-dhcp-client isc-dhcp-server; systemctl daemon-reload; systemctl enable tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service; systemctl restart tor@default.service privoxy.service squid.service dnsmasq.service changemac@eth0.service`)
}

func Vanishstart() {
    subprocess.Popen(`cd ~/.africana/africana-base/tor-vanish/; python3 vanisher.py -m`)
    fmt.Println()
}

func Vanishstop() {
    subprocess.Popen(`cd ~/.africana/africana-base/tor-vanish/; python3 vanisher.py -e`)
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
