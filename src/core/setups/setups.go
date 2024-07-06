package setups

import (
    "os"
    "fmt"
    "menus"
    "utils"
    "bcolors"
    "banners"
    "scriptures"
    "subprocess"
)

var(
    userInput string
)

func KaliSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)" + bcolors.ENDC)
        subprocess.Popen(`apt-get update -y; apt-get install zsh git curl wget -y; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o /etc/apt/sources.list.d/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux aha set playit libssl-dev gcc hydra wine32:i386`)
        fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        subprocess.Popen(`cd /root/.africana/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1; cd ./africana-base; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; rm -rf main.zip;cd ..; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; pip3 install --upgrade setuptools; pip3 install -r /root/.africana/africana-base/requirements.txt; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)`)
        fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.DARKCYAN + "🧬Africana already installed. " + bcolors.YELLOW + "Update it? " + bcolors.RED + "(y/n)\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
            case "y", "Y", "yes", "Yes", "YES":
                subprocess.Popen(`cd /root/.africana/africana-base; git pull .; git clone https://github.com/r0jahsm0ntar1/africana-framework; cd africana-framework; go build africana; mv africana /usr/local/bin; rm -rf ../africana-framework`)
                fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Updated. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
                return
            case "n", "N", "no", "No", "NO":
                utils.ClearScreen(); banners.Banner(); scriptures.Verse()
                return
            default:
                fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
        }
    }
}

func UbuntuSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)" + bcolors.ENDC)
        subprocess.Popen(`apt-get update -y; apt-get install zsh git curl wget -y; wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc; echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref; mkdir -p /etc/apt/trusted.gpg.d/; cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o /etc/apt/sources.list.d/playit-cloud.list; dpkg --add-architecture i386; apt-get update -y; apt-get install -y tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux aha set playit libssl-dev gcc hydra wine32:i386`)
        fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        subprocess.Popen(`cd /root/.africana/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1; cd ./africana-base; git clone https://github.com/ScRiPt1337/Teardroid-phprat; cd ./Teardroid-phprat; pip3 install -r requirements.txt; wget https://github.com/ScRiPt1337/Teardroidv4_api/archive/refs/heads/main.zip; unzip main.zip; rm -rf main.zip;cd ..; git clone https://github.com/devanshbatham/paramspider; cd paramspider; pip3 install .; cd ..; git clone https://github.com/TermuxHackz/anonphisher; cd ./anonphisher; chmod 777 *; cd ..; pip3 install --upgrade setuptools; pip3 install -r /root/.africana/africana-base/requirements.txt; go install github.com/j3ssie/osmedeus@latest; go install github.com/hahwul/dalfox/v2@latest; bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)`)
        fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.DARKCYAN + "🧬Africana already installed. " + bcolors.YELLOW + "Update it? " + bcolors.RED + "(y/n)\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
            case "y", "Y", "yes", "Yes", "YES":
                subprocess.Popen(`cd /root/.africana/africana-base; git pull .; git clone https://github.com/r0jahsm0ntar1/africana-framework; cd africana-framework; go build africana; mv africana /usr/local/bin; rm -rf ../africana-framework`)
                fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Updated. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
                return
            case "n", "N", "no", "No", "NO":
                utils.ClearScreen(); banners.Banner(); scriptures.Verse()
                return
            default:
                fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
        }
    }
}

func ArchaLSetups() {

}

func WindowsSetup() {

}

func RemoveSetups() {
    fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.DARKCYAN + "🧬Are you sure you want to " + bcolors.RED + "UNINSTALL! " + bcolors.YELLOW + "africana? " + bcolors.RED + "(y/n)\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
        case "y", "Y", "yes", "Yes", "YES":
            subprocess.Popen(`rm -rf /root/.africana/; rm -rf /usr/local/bin/africana`)
            fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully. " + bcolors.YELLOW + "Uninstalled!" + bcolors.ENDC)
            return
        case "n", "N", "no", "No", "NO":
            menus.MenuOne()
            return
        default:
            fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
    }
}
