package setups

import(
    "os"
    "fmt"
    "bufio"
    "menus"
    "utils"
    "strings"
    "bcolors"
    "banners"
    "scriptures"
    "subprocess"
)

var(
    userInput string
    scanner = bufio.NewScanner(os.Stdin)
)

func installFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Popen(cmd)
        fmt.Println()
    }
}

func installGithubTools() {
    githubCommands := []string{
        //cd /root/.africana/; python3 -m virtualenv africana-venv --system-site-packages; echo -n "\n source ~/.africana/africana-venv/bin/activate" >> ~/.zshrc
        `cd /root/.africana/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1`,
        `python3 -m pip install --upgrade setuptools`,
        `python3 -m pip install -r /root/.africana/africana-base/requirements.txt`,
        `go install github.com/j3ssie/osmedeus@latest`,
        `go install github.com/hahwul/dalfox/v2@latest`,
        `bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)`,
    }

    for _, cmd := range githubCommands {
        subprocess.Popen(cmd)
        fmt.Println()
    }
}

func promptForUpdate() {
    for {
        fmt.Println(bcolors.RED + "(҂`_´) " + bcolors.ENDC + "Africana already installed. " + bcolors.YELLOW + "Update it? " + bcolors.RED + "(y/n)" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Installer" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🩺" + bcolors.GREEN + "> " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "y", "yes":
            subprocess.Popen(`rm -rf /root/.africana/africana-base`)
            installGithubTools()
            subprocess.Popen(`cd /root/.africana/; git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1; cd africana-framework; make; mv africana-linux /usr/local/bin/africana; rm -rf ../africana-framework`)
            fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Africana Fully Updated. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
            return
        case "n", "no":
            utils.ClearScreen()
            banners.RandomBanners()
            scriptures.Verse()
            return
        default:
            fmt.Println(bcolors.GREEN + "(" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
        }
    }
}

func KaliSetups() {
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " setups(" + bcolors.RED +  "kali/setup_func.go" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    for {
        switch strings.ToLower(userInput) {
            case "banner":
                banners.RandomBanners()
                KaliSetups()
                return
            case "sleep":
                utils.Sleep()
                KaliSetups()
                return
            case "info":
                menus.HelpInfoKali()
                KaliSetups()
                return
            case "v", "version":
                banners.Version()
                KaliSetups()
                return
            case "options", "show options":
                menus.HelpInfoExecute()
                KaliSetups()
                return
            case "m", "menu":
                menus.MenuTwo()
                KaliSetups()
                return
            case "logs", "history":
                subprocess.LogHistory()
                KaliSetups()
                return
            case "o", "junks", "outputs":
                utils.ListJunks()
                KaliSetups()
                return
            case "? info", "help info":
                menus.HelpInfo()
                KaliSetups()
                return
            case "00", "?", "h", "help":
                menus.HelpInfoMenuZero()
                KaliSetups()
                return
            case "e", "q", "exit", "quit":
                os.Exit(0)
            case "0", "b", "back":
                return
            case "g", "t", "guide", "tutarial":
                utils.BrowseTutarilas()
                KaliSetups()
                return
            case "clear logs", "clear history":
                subprocess.ClearHistory()
                KaliSetups()
                return
            case "info set", "set", "help set":
                menus.HelpInfoSet()
                KaliSetups()
                return
            case "clear junks", "clear outputs":
                utils.ClearJunks()
                KaliSetups()
                return
            case "? run", "info run", "help run":
                menus.HelpInfoRun()
                KaliSetups()
                return
            case "? start", "info start", "help start":
                menus.HelpInfoStart()
                KaliSetups()
                return
            case "? use", "info use", "help use", "use":
                menus.HelpInfoUse()
                KaliSetups()
                return
            case "f", "use f", "use features", "features":
                menus.HelpInfoFeatures()
                KaliSetups()
                return
            case "? tips", "info tips", "help tips", "tips":
                menus.HelpInfoTips()
                KaliSetups()
                return
            case "? show", "info show", "help show", "show":
                menus.HelpInfoShow()
                KaliSetups()
                return
            case "info list", "help list", "use list", "list":
                menus.HelpInfoList()
                KaliSetups()
                return
            case "? options", "info options", "help options":
                menus.HelpInfOptions()
                KaliSetups()
                return
            case "run", "start", "execute":
                filePath := "/root/.africana/africana-base/"
                if _, err := os.Stat(filePath); os.IsNotExist(err) {
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
                    foundationCommands := []string{
                        `apt-get update -y`,
                        `apt-get install zsh git curl wget -y`,
                        `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
                        `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
                        `dpkg --add-architecture i386`,
                        `apt-get update -y`,
                        `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
                        `winecfg /v win11`,
                    }
                    installFoundationTools(foundationCommands)
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing Github third party tools" + bcolors.ENDC)
                    installGithubTools()
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
                } else {
                    promptForUpdate()
                }
            default:
                utils.SystemShell(userInput)
                KaliSetups()
                return
         }
    }
}

func UbuntuSetups() {
    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " setups(" + bcolors.RED +  "ubuntu/setup_func.go" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
    scanner.Scan()
    userInput := scanner.Text()
    for {
        switch strings.ToLower(userInput) {
            case "banner":
                banners.RandomBanners()
                UbuntuSetups()
                return
            case "sleep":
                utils.Sleep()
                UbuntuSetups()
                return
            case "v", "version":
                banners.Version()
                UbuntuSetups()
                return
            case "m", "menu":
                menus.MenuTwo()
                UbuntuSetups()
                return
            case "logs", "history":
                subprocess.LogHistory()
                UbuntuSetups()
                return
            case "o", "junks", "outputs":
                utils.ListJunks()
                UbuntuSetups()
                return
            case "? info", "info", "help info":
                menus.HelpInfo()
                UbuntuSetups()
                return
            case "00", "?", "h", "help":
                menus.HelpInfoMenuZero()
                UbuntuSetups()
                return
            case "e", "q", "exit", "quit":
                os.Exit(0)
            case "0", "b", "back":
                return
            case "g", "t", "guide", "tutarial":
                utils.BrowseTutarilas()
                UbuntuSetups()
                return
            case "clear logs", "clear history":
                subprocess.ClearHistory()
                UbuntuSetups()
                return
            case "info set", "set", "help set":
                menus.HelpInfoSet()
                UbuntuSetups()
                return
            case "clear junks", "clear outputs":
                utils.ClearJunks()
                UbuntuSetups()
                return
            case "? run", "info run", "help run":
                menus.HelpInfoRun()
                UbuntuSetups()
                return
            case "? start", "info start", "help start":
                menus.HelpInfoStart()
                UbuntuSetups()
                return
            case "? use", "info use", "help use", "use":
                menus.HelpInfoUse()
                UbuntuSetups()
                return
            case "f", "use f", "use features", "features":
                menus.HelpInfoFeatures()
                UbuntuSetups()
                return
            case "? tips", "info tips", "help tips", "tips":
                menus.HelpInfoTips()
                banners.RandomBanners()
                UbuntuSetups()
                return
            case "? show", "info show", "help show", "show":
                menus.HelpInfoShow()
                UbuntuSetups()
                return
            case "info list", "help list", "use list", "list":
                menus.HelpInfoList()
                UbuntuSetups()
                return
            case "? options", "info options", "help options":
                menus.HelpInfOptions()
                UbuntuSetups()
                return
            case "run", "start", "execute":
                filePath := "/root/.africana/africana-base/"
                if _, err := os.Stat(filePath); os.IsNotExist(err) {
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
                    foundationCommands := []string{
                        `apt-get update -y`,
                        `apt-get install zsh git curl wget -y`,
                        `wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc`,
                        `echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref`,
                        `cd /etc/apt/trusted.gpg.d/; curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg`,
                        `cd /etc/apt/sources.list.d/; curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o playit-cloud.list`,
                        `dpkg --add-architecture i386`,
                        `apt-get update -y`,
                        `apt-get install -y bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
                        `winecfg /v win11`,
                    }
                    installFoundationTools(foundationCommands)
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing Github third party tools" + bcolors.ENDC)
                    installGithubTools()
                    fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
                } else {
                    promptForUpdate()
                }
            default:
                utils.SystemShell(userInput)
                UbuntuSetups()
                return
         }
    }
}

func ArchSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)" + bcolors.ENDC)
        foundationCommands := []string{
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm zsh git curl wget go`,
        }
        installFoundationTools(foundationCommands)
        BlackArchSetups()
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func BlackArchSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
        foundationCommands := []string{
            `curl -O https://blackarch.org/strap.sh`,
            `chmod +x strap.sh`,
            `./strap.sh`,
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm blackarch`,
            `pacman -S --noconfirm base-devel bc jq npm tor aha ftp ncat gcc gawk tmux mdk4 mdk3 nmap playit rlwrap squid privoxy iptables dnsmasq openssh-client libpcap-dev openssh-server powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework gobuster feroxbuster uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux netexec libssl-dev aircrack-ng cowpatty dhcpd hostapd lighttpd net-tools macchanger dsniff openssl php-cgi xterm rfkill unzip hydra wine32:i386`,
            `winecfg /v win11`,
        }
        installFoundationTools(foundationCommands)
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func WindowsSetups() {
    // Placeholder for Windows setup logic
}

func Androidinstall() {
    // Placeholder for Windows setup logic
}

func TotalUninstaller() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); !os.IsNotExist(err) {
        for {
            fmt.Printf(bcolors.UNDERL + bcolors.BOLD + "afr3" + bcolors.ENDC + " setups(" + bcolors.RED +  "uninstaller/setup_func.go" + bcolors.ENDC + ")" + bcolors.GREEN + " > " + bcolors.ENDC)
            scanner.Scan()
            userInput := scanner.Text()
            switch strings.ToLower(userInput) {
            case "banner":
                banners.RandomBanners()
                TotalUninstaller()
                return
            case "sleep":
                utils.Sleep()
                TotalUninstaller()
                return
            case "v", "version":
                banners.Version()
                TotalUninstaller()
                return
            case "m", "menu":
                menus.MenuTwo()
                TotalUninstaller()
                return
            case "logs", "history":
                subprocess.LogHistory()
                UbuntuSetups()
                return
            case "o", "junks", "outputs":
                utils.ListJunks()
                TotalUninstaller()
                return
            case "? info", "info", "help info":
                menus.HelpInfo()
                TotalUninstaller()
                return
            case "00", "?", "h", "help":
                menus.HelpInfoMenuZero()
                TotalUninstaller()
                return
            case "e", "q", "exit", "quit":
                os.Exit(0)
            case "0", "b", "back":
                return
            case "g", "t", "guide", "tutarial":
                utils.BrowseTutarilas()
                TotalUninstaller()
                return
            case "clear logs", "clear history":
                subprocess.ClearHistory()
                TotalUninstaller()
                return
            case "info set", "set", "help set":
                menus.HelpInfoSet()
                TotalUninstaller()
                return
            case "clear junks", "clear outputs":
                utils.ClearJunks()
                TotalUninstaller()
                return
            case "? run", "info run", "help run":
                menus.HelpInfoRun()
                TotalUninstaller()
                return
            case "? start", "info start", "help start":
                menus.HelpInfoStart()
                TotalUninstaller()
                return
            case "? use", "info use", "help use", "use":
                menus.HelpInfoUse()
                TotalUninstaller()
                return
            case "f", "use f", "use features", "features":
                menus.HelpInfoFeatures()
                TotalUninstaller()
                return
            case "? tips", "info tips", "help tips", "tips":
                menus.HelpInfoTips()
                banners.RandomBanners()
                TotalUninstaller()
                return
            case "? show", "info show", "help show", "show":
                menus.HelpInfoShow()
                TotalUninstaller()
                return
            case "info list", "help list", "use list", "list":
                menus.HelpInfoList()
                TotalUninstaller()
                return
            case "? options", "info options", "help options":
                menus.HelpInfOptions()
                TotalUninstaller()
                return
            case "run", "start", "execute":
            case "y", "yes":
                subprocess.Popen(`rm -rf /root/.africana/; rm -rf /usr/local/bin/africana`)
                fmt.Println(bcolors.RED + "[*] " + bcolors.ENDC + "Africana successfully uninstalled. " + bcolors.YELLOW + "Goodbye!" + bcolors.ENDC)
                return
            case "n", "no":
                utils.ClearScreen()
                banners.RandomBanners()
                scriptures.Verse()
                menus.MenuOne()
                return
            default:
                fmt.Println(bcolors.DARKCYAN + "                       Choices are (" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + ")" + bcolors.ENDC)
            }
        }
    } else {
        fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.RED + "Africana is not installed. " + bcolors.ENDC)
    }
}
